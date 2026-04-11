This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate SyntheticDataColumnProperties

Properties that define how a specific data column should be handled during synthetic data generation, including its name, type, and role in predictive modeling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ColumnType" : String,
  "IsPredictiveValue" : Boolean
}

```

### YAML

```yaml

  ColumnName: String
  ColumnType: String
  IsPredictiveValue: Boolean

```

## Properties

`ColumnName`

The name of the data column as it appears in the dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9_](([a-z0-9_]+-)*([a-z0-9_]+))?$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnType`

The data type of the column, which determines how the synthetic data generation algorithm processes and synthesizes values for this column.

_Required_: Yes

_Type_: String

_Allowed values_: `CATEGORICAL | NUMERICAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IsPredictiveValue`

Indicates if this column contains predictive values that should be treated as target variables in machine learning models. This affects how the synthetic data generation preserves statistical relationships.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Location

SyntheticDataParameters

All content copied from https://docs.aws.amazon.com/.
