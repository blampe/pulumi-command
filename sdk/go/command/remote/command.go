// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package remote

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A command to run on a remote host.
// The connection is established via ssh.
type Command struct {
	pulumi.CustomResourceState

	// The parameters with which to connect to the remote host
	Connection ConnectionPtrOutput `pulumi:"connection"`
	// The command to run on create.
	Create pulumi.StringPtrOutput `pulumi:"create"`
	// The command to run on delete.
	Delete pulumi.StringPtrOutput `pulumi:"delete"`
	// Additional environment variables available to the command's process.
	Environment pulumi.StringMapOutput `pulumi:"environment"`
	// The standard error of the command's process
	Stderr pulumi.StringOutput `pulumi:"stderr"`
	// The standard output of the command's process
	Stdout pulumi.StringOutput `pulumi:"stdout"`
	// The command to run on update.
	Update pulumi.StringPtrOutput `pulumi:"update"`
}

// NewCommand registers a new resource with the given unique name, arguments, and options.
func NewCommand(ctx *pulumi.Context,
	name string, args *CommandArgs, opts ...pulumi.ResourceOption) (*Command, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connection == nil {
		return nil, errors.New("invalid value for required argument 'Connection'")
	}
	var resource Command
	err := ctx.RegisterResource("command:remote:Command", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommand gets an existing Command resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommandState, opts ...pulumi.ResourceOption) (*Command, error) {
	var resource Command
	err := ctx.ReadResource("command:remote:Command", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Command resources.
type commandState struct {
}

type CommandState struct {
}

func (CommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*commandState)(nil)).Elem()
}

type commandArgs struct {
	// The parameters with which to connect to the remote host
	Connection Connection `pulumi:"connection"`
	// The command to run on create.
	Create *string `pulumi:"create"`
	// The command to run on delete.
	Delete *string `pulumi:"delete"`
	// Additional environment variables available to the command's process.
	Environment map[string]string `pulumi:"environment"`
	// The command to run on update.
	Update *string `pulumi:"update"`
}

// The set of arguments for constructing a Command resource.
type CommandArgs struct {
	// The parameters with which to connect to the remote host
	Connection ConnectionInput
	// The command to run on create.
	Create pulumi.StringPtrInput
	// The command to run on delete.
	Delete pulumi.StringPtrInput
	// Additional environment variables available to the command's process.
	Environment pulumi.StringMapInput
	// The command to run on update.
	Update pulumi.StringPtrInput
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commandArgs)(nil)).Elem()
}

type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(ctx context.Context) CommandOutput
}

func (*Command) ElementType() reflect.Type {
	return reflect.TypeOf((*Command)(nil))
}

func (i *Command) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i *Command) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}

func (i *Command) ToCommandPtrOutput() CommandPtrOutput {
	return i.ToCommandPtrOutputWithContext(context.Background())
}

func (i *Command) ToCommandPtrOutputWithContext(ctx context.Context) CommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandPtrOutput)
}

type CommandPtrInput interface {
	pulumi.Input

	ToCommandPtrOutput() CommandPtrOutput
	ToCommandPtrOutputWithContext(ctx context.Context) CommandPtrOutput
}

type commandPtrType CommandArgs

func (*commandPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil))
}

func (i *commandPtrType) ToCommandPtrOutput() CommandPtrOutput {
	return i.ToCommandPtrOutputWithContext(context.Background())
}

func (i *commandPtrType) ToCommandPtrOutputWithContext(ctx context.Context) CommandPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandPtrOutput)
}

// CommandArrayInput is an input type that accepts CommandArray and CommandArrayOutput values.
// You can construct a concrete instance of `CommandArrayInput` via:
//
//          CommandArray{ CommandArgs{...} }
type CommandArrayInput interface {
	pulumi.Input

	ToCommandArrayOutput() CommandArrayOutput
	ToCommandArrayOutputWithContext(context.Context) CommandArrayOutput
}

type CommandArray []CommandInput

func (CommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (i CommandArray) ToCommandArrayOutput() CommandArrayOutput {
	return i.ToCommandArrayOutputWithContext(context.Background())
}

func (i CommandArray) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandArrayOutput)
}

// CommandMapInput is an input type that accepts CommandMap and CommandMapOutput values.
// You can construct a concrete instance of `CommandMapInput` via:
//
//          CommandMap{ "key": CommandArgs{...} }
type CommandMapInput interface {
	pulumi.Input

	ToCommandMapOutput() CommandMapOutput
	ToCommandMapOutputWithContext(context.Context) CommandMapOutput
}

type CommandMap map[string]CommandInput

func (CommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (i CommandMap) ToCommandMapOutput() CommandMapOutput {
	return i.ToCommandMapOutputWithContext(context.Background())
}

func (i CommandMap) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandMapOutput)
}

type CommandOutput struct {
	*pulumi.OutputState
}

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Command)(nil))
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

func (o CommandOutput) ToCommandPtrOutput() CommandPtrOutput {
	return o.ToCommandPtrOutputWithContext(context.Background())
}

func (o CommandOutput) ToCommandPtrOutputWithContext(ctx context.Context) CommandPtrOutput {
	return o.ApplyT(func(v Command) *Command {
		return &v
	}).(CommandPtrOutput)
}

type CommandPtrOutput struct {
	*pulumi.OutputState
}

func (CommandPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil))
}

func (o CommandPtrOutput) ToCommandPtrOutput() CommandPtrOutput {
	return o
}

func (o CommandPtrOutput) ToCommandPtrOutputWithContext(ctx context.Context) CommandPtrOutput {
	return o
}

type CommandArrayOutput struct{ *pulumi.OutputState }

func (CommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Command)(nil))
}

func (o CommandArrayOutput) ToCommandArrayOutput() CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) Index(i pulumi.IntInput) CommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Command {
		return vs[0].([]Command)[vs[1].(int)]
	}).(CommandOutput)
}

type CommandMapOutput struct{ *pulumi.OutputState }

func (CommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Command)(nil))
}

func (o CommandMapOutput) ToCommandMapOutput() CommandMapOutput {
	return o
}

func (o CommandMapOutput) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return o
}

func (o CommandMapOutput) MapIndex(k pulumi.StringInput) CommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Command {
		return vs[0].(map[string]Command)[vs[1].(string)]
	}).(CommandOutput)
}

func init() {
	pulumi.RegisterOutputType(CommandOutput{})
	pulumi.RegisterOutputType(CommandPtrOutput{})
	pulumi.RegisterOutputType(CommandArrayOutput{})
	pulumi.RegisterOutputType(CommandMapOutput{})
}
