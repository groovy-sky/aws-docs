This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FieldValidationConfiguration

The `FieldValidationConfiguration` property specifies the validation configuration for a field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumValues" : [ Number, ... ],
  "StrValues" : [ String, ... ],
  "Type" : String,
  "ValidationMessage" : String
}

```

### YAML

```yaml

  NumValues:
    - Number
  StrValues:
    - String
  Type: String
  ValidationMessage: String

```

## Properties

`NumValues`

The validation to perform on a number value.

_Required_: No

_Type_: Array of Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrValues`

The validation to perform on a string value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The validation to perform on an object type. ``

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationMessage`

The validation message to display.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldPosition

FileUploaderFieldConfig

All content copied from https://docs.aws.amazon.com/.
