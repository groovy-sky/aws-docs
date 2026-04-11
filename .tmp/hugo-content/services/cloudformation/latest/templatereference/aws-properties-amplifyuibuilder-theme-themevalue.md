This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Theme ThemeValue

The `ThemeValue` property specifies the configuration of a theme's
properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Children" : [ ThemeValues, ... ],
  "Value" : String
}

```

### YAML

```yaml

  Children:
    - ThemeValues
  Value: String

```

## Properties

`Children`

A list of key-value pairs that define the theme's properties.

_Required_: No

_Type_: Array of [ThemeValues](aws-properties-amplifyuibuilder-theme-themevalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a theme property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AmplifyUIBuilder::Theme

ThemeValues

All content copied from https://docs.aws.amazon.com/.
