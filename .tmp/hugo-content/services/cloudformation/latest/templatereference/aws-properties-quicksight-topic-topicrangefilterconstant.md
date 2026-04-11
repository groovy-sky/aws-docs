This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicRangeFilterConstant

A constant value that is used in a range filter to specify the endpoints of the range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConstantType" : String,
  "RangeConstant" : RangeConstant
}

```

### YAML

```yaml

  ConstantType: String
  RangeConstant:
    RangeConstant

```

## Properties

`ConstantType`

The data type of the constant value that is used in a range filter. Valid values for this structure are `RANGE`.

_Required_: No

_Type_: String

_Allowed values_: `SINGULAR | RANGE | COLLECTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeConstant`

The value of the constant that is used to specify the endpoints of a range filter.

_Required_: No

_Type_: [RangeConstant](aws-properties-quicksight-topic-rangeconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicNumericRangeFilter

TopicRelativeDateFilter

All content copied from https://docs.aws.amazon.com/.
