This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource DatabricksParameters

The required parameters that are needed to connect to a Databricks data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Host" : String,
  "Port" : Number,
  "SqlEndpointPath" : String
}

```

### YAML

```yaml

  Host: String
  Port: Number
  SqlEndpointPath: String

```

## Properties

`Host`

The host name of the Databricks data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port for the Databricks data source.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlEndpointPath`

The HTTP path of the Databricks data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CredentialPair

DataSourceCredentials

All content copied from https://docs.aws.amazon.com/.
