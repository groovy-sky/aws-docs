This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AnalysisError

Analysis error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Message" : String,
  "Type" : String,
  "ViolatedEntities" : [ Entity, ... ]
}

```

### YAML

```yaml

  Message: String
  Type: String
  ViolatedEntities:
    - Entity

```

## Properties

`Message`

The message associated with the analysis error.

_Required_: No

_Type_: String

_Pattern_: `\S`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the analysis error.

_Required_: No

_Type_: String

_Allowed values_: `ACCESS_DENIED | SOURCE_NOT_FOUND | DATA_SET_NOT_FOUND | INTERNAL_FAILURE | PARAMETER_VALUE_INCOMPATIBLE | PARAMETER_TYPE_INVALID | PARAMETER_NOT_FOUND | COLUMN_TYPE_MISMATCH | COLUMN_GEOGRAPHIC_ROLE_MISMATCH | COLUMN_REPLACEMENT_MISSING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViolatedEntities`

Lists the violated entities that caused the analysis error

_Required_: No

_Type_: Array of [Entity](aws-properties-quicksight-analysis-entity.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisDefinition

AnalysisSourceEntity

All content copied from https://docs.aws.amazon.com/.
