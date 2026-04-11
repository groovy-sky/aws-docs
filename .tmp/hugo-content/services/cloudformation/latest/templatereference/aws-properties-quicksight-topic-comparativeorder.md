This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic ComparativeOrder

The order in which data is displayed for the column when
it's used in a comparative context.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SpecifedOrder" : [ String, ... ],
  "TreatUndefinedSpecifiedValues" : String,
  "UseOrdering" : String
}

```

### YAML

```yaml

  SpecifedOrder:
    - String
  TreatUndefinedSpecifiedValues: String
  UseOrdering: String

```

## Properties

`SpecifedOrder`

The list of columns to be used in the ordering.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatUndefinedSpecifiedValues`

The treat of undefined specified values. Valid values for this structure are `LEAST` and `MOST`.

_Required_: No

_Type_: String

_Allowed values_: `LEAST | MOST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseOrdering`

The ordering type for a column. Valid values for this structure are `GREATER_IS_BETTER`, `LESSER_IS_BETTER` and `SPECIFIED`.

_Required_: No

_Type_: String

_Allowed values_: `GREATER_IS_BETTER | LESSER_IS_BETTER | SPECIFIED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CollectiveConstant

CustomInstructions

All content copied from https://docs.aws.amazon.com/.
