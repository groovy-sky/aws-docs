This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DefaultRelativeDateTimeControlOptions

The default options that correspond to the `RelativeDateTime` filter control type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CommitMode" : String,
  "DisplayOptions" : RelativeDateTimeControlDisplayOptions
}

```

### YAML

```yaml

  CommitMode: String
  DisplayOptions:
    RelativeDateTimeControlDisplayOptions

```

## Properties

`CommitMode`

The visibility configuration of the Apply button on a `RelativeDateTimeControl`.

_Required_: No

_Type_: String

_Allowed values_: `AUTO | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayOptions`

The display options of a control.

_Required_: No

_Type_: [RelativeDateTimeControlDisplayOptions](aws-properties-quicksight-dashboard-relativedatetimecontroldisplayoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultPaginatedLayoutConfiguration

DefaultSectionBasedLayoutConfiguration

All content copied from https://docs.aws.amazon.com/.
