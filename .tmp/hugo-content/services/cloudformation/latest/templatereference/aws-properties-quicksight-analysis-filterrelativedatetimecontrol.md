This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FilterRelativeDateTimeControl

A control from a date filter that is used to specify the relative date.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitMode" : String,
  "DisplayOptions" : RelativeDateTimeControlDisplayOptions,
  "FilterControlId" : String,
  "SourceFilterId" : String,
  "Title" : String
}

```

### YAML

```yaml

  CommitMode: String
  DisplayOptions:
    RelativeDateTimeControlDisplayOptions
  FilterControlId: String
  SourceFilterId: String
  Title: String

```

## Properties

`CommitMode`

The visibility configuration of the Apply button on a `FilterRelativeDateTimeControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [RelativeDateTimeControlDisplayOptions](aws-properties-quicksight-analysis-relativedatetimecontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterControlId`

The ID of the `FilterTextAreaControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFilterId`

The source filter ID of the `FilterTextAreaControl`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `FilterTextAreaControl`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterOperationTargetVisualsConfiguration

FilterScopeConfiguration

All content copied from https://docs.aws.amazon.com/.
