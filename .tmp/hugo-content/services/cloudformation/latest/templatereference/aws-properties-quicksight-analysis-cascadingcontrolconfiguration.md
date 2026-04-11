This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CascadingControlConfiguration

The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceControls" : [ CascadingControlSource, ... ]
}

```

### YAML

```yaml

  SourceControls:
    - CascadingControlSource

```

## Properties

`SourceControls`

A list of source controls that determine the values that are used in the current control.

_Required_: No

_Type_: Array of [CascadingControlSource](aws-properties-quicksight-analysis-cascadingcontrolsource.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculatedMeasureField

CascadingControlSource

All content copied from https://docs.aws.amazon.com/.
