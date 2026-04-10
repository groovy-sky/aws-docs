This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource TrinoParameters

The parameters that are required to connect to a Trino data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Catalog" : String,
  "Host" : String,
  "Port" : Number
}

```

### YAML

```yaml

  Catalog: String
  Host: String
  Port: Number

```

## Properties

`Catalog`

The catalog name for the Trino data source.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

The host name of the Trino data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port for the Trino data source.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TeradataParameters

VpcConnectionProperties

All content copied from https://docs.aws.amazon.com/.
