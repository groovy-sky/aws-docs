This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FieldConfig

The `FieldConfig` property specifies the configuration information for a field in a table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Excluded" : Boolean,
  "InputType" : FieldInputConfig,
  "Label" : String,
  "Position" : FieldPosition,
  "Validations" : [ FieldValidationConfiguration, ... ]
}

```

### YAML

```yaml

  Excluded: Boolean
  InputType:
    FieldInputConfig
  Label: String
  Position:
    FieldPosition
  Validations:
    - FieldValidationConfiguration

```

## Properties

`Excluded`

Specifies whether to hide a field.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputType`

Describes the configuration for the default input value to display for a field.

_Required_: No

_Type_: [FieldInputConfig](aws-properties-amplifyuibuilder-form-fieldinputconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label for the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

Specifies the field position.

_Required_: No

_Type_: [FieldPosition](aws-properties-amplifyuibuilder-form-fieldposition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validations`

The validations to perform on the value in the field.

_Required_: No

_Type_: Array of [FieldValidationConfiguration](aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AmplifyUIBuilder::Form

FieldInputConfig

All content copied from https://docs.aws.amazon.com/.
