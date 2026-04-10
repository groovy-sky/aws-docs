This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationNFS MountOptions

Specifies the options that DataSync can use to mount your NFS file
server.

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

Specifies the NFS version that you want DataSync to use when mounting your NFS
share. If the server refuses to use the version specified, the task fails.

You can specify the following options:

- `AUTOMATIC` (default): DataSync chooses NFS version 4.1.

- `NFS3`: Stateless protocol version that allows for asynchronous writes on
the server.

- `NFSv4_0`: Stateful, firewall-friendly protocol version that supports
delegations and pseudo file systems.

- `NFSv4_1`: Stateful protocol version that supports sessions, directory
delegations, and parallel data processing. NFS version 4.1 also includes all features
available in version 4.0.

###### Note

DataSync currently only supports NFS version 3 with Amazon FSx for NetApp ONTAP locations.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | NFS3 | NFS4_0 | NFS4_1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataSync::LocationNFS

OnPremConfig

All content copied from https://docs.aws.amazon.com/.
