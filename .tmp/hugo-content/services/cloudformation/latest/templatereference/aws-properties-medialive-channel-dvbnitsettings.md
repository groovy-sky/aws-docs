This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel DvbNitSettings

The configuration of DVB NIT.

The parent of this entity is M2tsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkId" : Integer,
  "NetworkName" : String,
  "RepInterval" : Integer
}

```

### YAML

```yaml

  NetworkId: Integer
  NetworkName: String
  RepInterval: Integer

```

## Properties

`NetworkId`

The numeric value placed in the Network Information Table (NIT).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkName`

The network name text placed in the networkNameDescriptor inside the Network Information Table (NIT). The
maximum length is 256 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepInterval`

The number of milliseconds between instances of this table in the output transport stream.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DisabledLockingSettings

DvbSdtSettings

All content copied from https://docs.aws.amazon.com/.
