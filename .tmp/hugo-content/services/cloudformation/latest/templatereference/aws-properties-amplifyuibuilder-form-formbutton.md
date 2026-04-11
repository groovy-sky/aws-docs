This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormButton

The `FormButton` property specifies the configuration for a button UI element that is a part of a form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Children" : String,
  "Excluded" : Boolean,
  "Position" : FieldPosition
}

```

### YAML

```yaml

  Children: String
  Excluded: Boolean
  Position:
    FieldPosition

```

## Properties

`Children`

Describes the button's properties.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Excluded`

Specifies whether the button is visible on the form.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

The position of the button.

_Required_: No

_Type_: [FieldPosition](aws-properties-amplifyuibuilder-form-fieldposition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FileUploaderFieldConfig

FormCTA

All content copied from https://docs.aws.amazon.com/.
