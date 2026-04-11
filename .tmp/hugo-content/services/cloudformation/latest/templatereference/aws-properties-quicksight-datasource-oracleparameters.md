This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource OracleParameters

Oracle parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Database" : String,
  "Host" : String,
  "Port" : Number,
  "UseServiceName" : Boolean
}

```

### YAML

```yaml

  Database: String
  Host: String
  Port: Number
  UseServiceName: Boolean

```

## Properties

`Database`

Database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

Host.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

Port.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseServiceName`

A Boolean value that indicates whether the `Database` uses a service name or an SID. If this value is left blank, the default value is `SID`. If this value is set to `false`, the value is `SID`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuthParameters

PostgreSqlParameters

All content copied from https://docs.aws.amazon.com/.
