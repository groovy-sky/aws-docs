This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationFSxOpenZFS MountOptions

Represents the mount options that are available for DataSync to access a
Network File System (NFS) location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Version" : String
}

```

### YAML

```yaml

  Version: String

```

## Properties

`Version`

The specific NFS version that you want DataSync to use to mount your NFS
share. If the server refuses to use the version specified, the sync will fail. If you
don't specify a version, DataSync defaults to `AUTOMATIC`. That is,
DataSync automatically selects a version based on negotiation with the NFS
server.

You can specify the following NFS versions:

- **[NFSv3](https://tools.ietf.org/html/rfc1813)**: Stateless protocol version that allows for asynchronous writes on
the server.

- **[NFSv4.0](https://tools.ietf.org/html/rfc3530)**: Stateful, firewall-friendly protocol version that supports
delegations and pseudo file systems.

- **[NFSv4.1](https://tools.ietf.org/html/rfc5661)**: Stateful protocol version that supports sessions, directory
delegations, and parallel data processing. Version 4.1 also includes all
features available in version 4.0.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | NFS3 | NFS4_0 | NFS4_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataSync::LocationFSxOpenZFS

NFS

All content copied from https://docs.aws.amazon.com/.
