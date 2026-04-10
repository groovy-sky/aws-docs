This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ReferenceLineStyleConfiguration

The style configuration of the reference
line.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "Pattern" : String
}

```

### YAML

```yaml

  Color: String
  Pattern: String

```

## Properties

`Color`

The hex color of the reference line.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pattern`

The pattern type of the line style. Choose one of the following options:

- `SOLID`

- `DASHED`

- `DOTTED`

_Required_: No

_Type_: String

_Allowed values_: `SOLID | DASHED | DOTTED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineStaticDataConfiguration

ReferenceLineValueLabelConfiguration

All content copied from https://docs.aws.amazon.com/.
