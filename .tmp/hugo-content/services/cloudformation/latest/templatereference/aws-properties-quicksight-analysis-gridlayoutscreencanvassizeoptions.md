This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis GridLayoutScreenCanvasSizeOptions

The options that determine the sizing of the canvas used in a grid layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OptimizedViewPortWidth" : String,
  "ResizeOption" : String
}

```

### YAML

```yaml

  OptimizedViewPortWidth: String
  ResizeOption: String

```

## Properties

`OptimizedViewPortWidth`

The width that the view port will be optimized for when the layout renders.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResizeOption`

This value determines the layout behavior when the viewport is resized.

- `FIXED`: A fixed width will be used when optimizing the layout. In
the Quick Sight console, this option is called `Classic`.

- `RESPONSIVE`: The width of the canvas will be responsive and
optimized to the view port. In the Quick Sight console, this option is called `Tiled`.

_Required_: Yes

_Type_: String

_Allowed values_: `FIXED | RESPONSIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GridLayoutElement

GrowthRateComputation

All content copied from https://docs.aws.amazon.com/.
