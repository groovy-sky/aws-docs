This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DonutCenterOptions

The label options of the label that is displayed in the center of a donut chart. This option isn't available for pie charts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelVisibility" : String
}

```

### YAML

```yaml

  LabelVisibility: String

```

## Properties

`LabelVisibility`

Determines the visibility of the label in a donut chart. In the Quick Sight console, this option is called `'Show total'`.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DimensionField

DonutOptions

All content copied from https://docs.aws.amazon.com/.
