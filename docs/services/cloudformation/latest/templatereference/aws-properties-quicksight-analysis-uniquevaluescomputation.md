This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis UniqueValuesComputation

The unique values computation configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Category" : DimensionField,
  "ComputationId" : String,
  "Name" : String
}

```

### YAML

```yaml

  Category:
    DimensionField
  ComputationId: String
  Name: String

```

## Properties

`Category`

The category field that is used in a computation.

_Required_: No

_Type_: [DimensionField](aws-properties-quicksight-analysis-dimensionfield.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationId`

The ID for a computation.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a computation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UnaggregatedField

ValidationStrategy

All content copied from https://docs.aws.amazon.com/.
