# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'Connection',
]

@pulumi.output_type
class Connection(dict):
    """
    Instructions for how to connect to a remote endpoint.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "privateKey":
            suggest = "private_key"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in Connection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        Connection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        Connection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 host: str,
                 password: Optional[str] = None,
                 port: Optional[float] = None,
                 private_key: Optional[str] = None,
                 user: Optional[str] = None):
        """
        Instructions for how to connect to a remote endpoint.
        :param str host: The address of the resource to connect to.
        :param str password: The password we should use for the connection.
        :param float port: The port to connect to. Defaults to 22.
        :param str private_key: The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        """
        pulumi.set(__self__, "host", host)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        The address of the resource to connect to.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        """
        The password we should use for the connection.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        The port to connect to. Defaults to 22.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[str]:
        """
        The contents of an SSH key to use for the connection. This takes preference over the password if provided.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def user(self) -> Optional[str]:
        return pulumi.get(self, "user")


