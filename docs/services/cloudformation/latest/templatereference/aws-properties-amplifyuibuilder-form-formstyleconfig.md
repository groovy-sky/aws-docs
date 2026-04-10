This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormStyleConfig

The `FormStyleConfig` property specifies the configuration settings for the form's style properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TokenReference" : String,
  "Value" : String
}

```

### YAML

```yaml

  TokenReference: String
  Value: String

```

## Properties

`TokenReference`

A reference to a design token to use to bind the form's style properties to an existing
theme.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the style setting.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormStyle

SectionalElement

All content copied from https://docs.aws.amazon.com/.
