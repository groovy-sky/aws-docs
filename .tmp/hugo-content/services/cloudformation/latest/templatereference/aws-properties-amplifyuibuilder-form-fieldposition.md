This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FieldPosition

The `FieldPosition` property specifies the field position.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Below" : String,
  "Fixed" : String,
  "RightOf" : String
}

```

### YAML

```yaml

  Below: String
  Fixed: String
  RightOf: String

```

## Properties

`Below`

The field position is below the field specified by the string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fixed`

The field position is fixed and doesn't change in relation to other fields.

_Required_: No

_Type_: String

_Allowed values_: `first`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RightOf`

The field position is to the right of the field specified by the string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldInputConfig

FieldValidationConfiguration

All content copied from https://docs.aws.amazon.com/.
