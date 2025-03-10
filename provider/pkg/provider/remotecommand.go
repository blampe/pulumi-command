// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/retry"

	"golang.org/x/crypto/ssh"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type remoteconnection struct {
	User       string  `pulumi:"user,optional"`
	Password   *string `pulumi:"password,optional"`
	Host       string  `pulumi:"host"`
	Port       int     `pulumi:"port,optional"`
	PrivateKey *string `pulumi:"privateKey,optional"`
}

// Generate an ssh config from a connection specification.
func (con remoteconnection) SShConfig() (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User:            con.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if con.PrivateKey != nil {
		signer, err := ssh.ParsePrivateKey([]byte(*con.PrivateKey))
		if err != nil {
			return nil, err
		}
		config.Auth = append(config.Auth, ssh.PublicKeys(signer))
	}
	if con.Password != nil {
		config.Auth = append(config.Auth, ssh.Password(*con.Password))
		config.Auth = append(config.Auth, ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
			for i := range questions {
				answers[i] = *con.Password
			}
			return answers, err
		}))
	}

	return config, nil
}

// Dial a ssh client connection from a ssh client configuration, retrying as necessary.
func (con remoteconnection) Dial(ctx context.Context, config *ssh.ClientConfig) (*ssh.Client, error) {
	var client *ssh.Client
	var err error
	_, _, err = retry.Until(ctx, retry.Acceptor{
		Accept: func(try int, nextRetryTime time.Duration) (bool, interface{}, error) {
			client, err = ssh.Dial("tcp",
				net.JoinHostPort(con.Host, fmt.Sprintf("%d", con.Port)),
				config)
			if err != nil {
				if try > 10 {
					return true, nil, err
				}
				return false, nil, nil
			}
			return true, nil, nil
		},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

type remotecommand struct {
	// Input
	Connection  remoteconnection   `pulumi:"connection"`
	Interpreter *[]string          `pulumi:"interpreter,optional"`
	Dir         *string            `pulumi:"dir,optional"`
	Environment *map[string]string `pulumi:"environment,optional"`
	Triggers    *[]interface{}     `pulumi:"triggers,optional"`
	Create      string             `pulumi:"create"`
	Delete      *string            `pulumi:"delete,optional"`
	Stdin       *string            `pulumi:"stdin,optional"`

	// Output
	Stdout string `pulumi:"stdout"`
	Stderr string `pulumi:"stderr"`
}

// RunCreate executes the create command, sets Stdout and Stderr, and returns a unique
// ID for the command execution
func (c *remotecommand) RunCreate(ctx context.Context, host *provider.HostClient, urn resource.URN) (string, error) {
	stdout, stderr, id, err := c.run(ctx, c.Create, host, urn)
	c.Stdout = stdout
	c.Stderr = stderr
	return id, err
}

// RunDelete executes the delete command
func (c *remotecommand) RunDelete(ctx context.Context, host *provider.HostClient, urn resource.URN) error {
	if c.Delete == nil {
		return nil
	}
	_, _, _, err := c.run(ctx, *c.Delete, host, urn)
	return err
}

func (c *remotecommand) run(ctx context.Context, cmd string, host *provider.HostClient, urn resource.URN) (string, string, string, error) {
	config, err := c.Connection.SShConfig()
	if err != nil {
		return "", "", "", err
	}

	client, err := c.Connection.Dial(ctx, config)
	if err != nil {
		return "", "", "", err
	}

	session, err := client.NewSession()
	if err != nil {
		return "", "", "", err
	}
	defer session.Close()

	if c.Environment != nil {
		for k, v := range *c.Environment {
			session.Setenv(k, v)
		}
	}

	if c.Stdin != nil && len(*c.Stdin) > 0 {
		session.Stdin = strings.NewReader(*c.Stdin)
	}

	id, err := resource.NewUniqueHex("", 8, 0)
	if err != nil {
		return "", "", "", err
	}

	stdoutr, stdoutw, err := os.Pipe()
	if err != nil {
		return "", "", "", err
	}
	stderrr, stderrw, err := os.Pipe()
	if err != nil {
		return "", "", "", err
	}
	session.Stdout = stdoutw
	session.Stderr = stderrw

	var stdoutbuf bytes.Buffer
	var stderrbuf bytes.Buffer

	stdouttee := io.TeeReader(stdoutr, &stdoutbuf)
	stderrtee := io.TeeReader(stderrr, &stderrbuf)

	stdoutch := make(chan struct{})
	stderrch := make(chan struct{})
	go copyOutput(ctx, host, urn, stdouttee, stdoutch)
	go copyOutput(ctx, host, urn, stderrtee, stderrch)

	err = session.Run(cmd)

	stdoutw.Close()
	stderrw.Close()

	<-stdoutch
	<-stderrch

	return stdoutbuf.String(), stderrbuf.String(), id, err
}
