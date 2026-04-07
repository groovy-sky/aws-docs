This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationSMB MountOptions

Specifies the version of the SMB protocol that DataSync uses to access your SMB
file server.

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

By default, DataSync automatically chooses an SMB protocol version based on
negotiation with your SMB file server. You also can configure DataSync to use a
specific SMB version, but we recommend doing this only if DataSync has trouble
negotiating with the SMB file server automatically.

These are the following options for configuring the SMB version:

- `AUTOMATIC` (default): DataSync and the SMB file server negotiate
the highest version of SMB that they mutually support between 2.1 and 3.1.1.

This is the recommended option. If you instead choose a specific version that your
file server doesn't support, you may get an `Operation Not Supported`
error.

- `SMB3`: Restricts the protocol negotiation to only SMB version
3.0.2.

- `SMB2`: Restricts the protocol negotiation to only SMB version 2.1.

- `SMB2_0`: Restricts the protocol negotiation to only SMB version
2.0.

- `SMB1`: Restricts the protocol negotiation to only SMB version 1.0.

###### Note

The `SMB1` option isn't available when [creating an Amazon FSx for NetApp ONTAP location](https://docs.aws.amazon.com/datasync/latest/userguide/API_CreateLocationFsxOntap.html).

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | SMB1 | SMB2_0 | SMB2 | SMB3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ManagedSecretConfig

Tag
