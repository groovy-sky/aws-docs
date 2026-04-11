This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RelationalFilterConfiguration

The relational filter configuration for the data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "FilterExpressions" : [ FilterExpression, ... ],
  "SchemaName" : String
}

```

### YAML

```yaml

  DatabaseName: String
  FilterExpressions:
    - FilterExpression
  SchemaName: String

```

## Properties

`DatabaseName`

The database name specified in the relational filter configuration for the data
source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterExpressions`

The filter expressions specified in the relational filter configuration for the data
source.

_Required_: No

_Type_: Array of [FilterExpression](aws-properties-datazone-datasource-filterexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaName`

The schema name specified in the relational filter configuration for the data
source.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftStorage

SageMakerRunConfigurationInput

All content copied from https://docs.aws.amazon.com/.
