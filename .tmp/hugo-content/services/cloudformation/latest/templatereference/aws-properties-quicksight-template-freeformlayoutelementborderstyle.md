This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FreeFormLayoutElementBorderStyle

The background style configuration of a free-form layout element.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  Color: String
  Visibility: String

```

## Properties

`Color`

The border color of a free-form layout element.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}(?:[A-F0-9]{2})?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The border visibility of a free-form layout element.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FreeFormLayoutElementBackgroundStyle

FreeFormLayoutScreenCanvasSizeOptions

All content copied from https://docs.aws.amazon.com/.
