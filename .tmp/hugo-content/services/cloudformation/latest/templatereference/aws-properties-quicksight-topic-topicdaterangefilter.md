This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicDateRangeFilter

A filter used to restrict data based on a range of dates or times.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Constant" : TopicRangeFilterConstant,
  "Inclusive" : Boolean
}

```

### YAML

```yaml

  Constant:
    TopicRangeFilterConstant
  Inclusive: Boolean

```

## Properties

`Constant`

The constant used in a date range filter.

_Required_: No

_Type_: [TopicRangeFilterConstant](aws-properties-quicksight-topic-topicrangefilterconstant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inclusive`

A Boolean value that indicates whether the date range filter should include the boundary values. If
set to true, the filter includes the start and end dates. If set to false, the filter
excludes them.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicConfigOptions

TopicFilter

All content copied from https://docs.aws.amazon.com/.
