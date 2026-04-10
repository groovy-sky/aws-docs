This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template GradientColor

Determines the gradient color settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Stops" : [ GradientStop, ... ]
}

```

### YAML

```yaml

  Stops:
    - GradientStop

```

## Properties

`Stops`

The list of gradient color stops.

_Required_: No

_Type_: Array of [GradientStop](aws-properties-quicksight-template-gradientstop.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlobalTableBorderOptions

GradientStop

All content copied from https://docs.aws.amazon.com/.
