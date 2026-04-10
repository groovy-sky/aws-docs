This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource DeltaSyncConfig

Describes a Delta Sync configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseTableTTL" : String,
  "DeltaSyncTableName" : String,
  "DeltaSyncTableTTL" : String
}

```

### YAML

```yaml

  BaseTableTTL: String
  DeltaSyncTableName: String
  DeltaSyncTableTTL: String

```

## Properties

`BaseTableTTL`

The number of minutes that an Item is stored in the data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeltaSyncTableName`

The Delta Sync table name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeltaSyncTableTTL`

The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsIamConfig

DynamoDBConfig

All content copied from https://docs.aws.amazon.com/.
