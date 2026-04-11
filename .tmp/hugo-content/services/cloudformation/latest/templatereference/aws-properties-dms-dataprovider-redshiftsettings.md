This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider RedshiftSettings

Provides information that defines an Amazon Redshift endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "Port" : Integer,
  "ServerName" : String
}

```

### YAML

```yaml

  DatabaseName: String
  Port: Integer
  ServerName: String

```

## Properties

`DatabaseName`

The name of the Amazon Redshift data warehouse (service) that you are working
with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port number for Amazon Redshift. The default value is 5439.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The name of the Amazon Redshift cluster you are using.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PostgreSqlSettings

Settings

All content copied from https://docs.aws.amazon.com/.
