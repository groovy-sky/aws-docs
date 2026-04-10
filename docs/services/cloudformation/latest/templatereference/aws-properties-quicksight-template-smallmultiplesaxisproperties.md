This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SmallMultiplesAxisProperties

Configures the properties of a chart's axes that are used by small multiples panels.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Placement" : String,
  "Scale" : String
}

```

### YAML

```yaml

  Placement: String
  Scale: String

```

## Properties

`Placement`

Defines the placement of the axis. By default, axes are rendered `OUTSIDE` of the panels. Axes with `INDEPENDENT` scale are rendered `INSIDE` the panels.

_Required_: No

_Type_: String

_Allowed values_: `OUTSIDE | INSIDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scale`

Determines whether scale of the axes are shared or independent. The default value is `SHARED`.

_Required_: No

_Type_: String

_Allowed values_: `SHARED | INDEPENDENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SliderControlDisplayOptions

SmallMultiplesOptions

All content copied from https://docs.aws.amazon.com/.
