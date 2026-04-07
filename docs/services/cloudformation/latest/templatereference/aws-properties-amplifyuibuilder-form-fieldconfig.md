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

_Type_: [FieldInputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-fieldinputconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Label`

The label for the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

Specifies the field position.

_Required_: No

_Type_: [FieldPosition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-fieldposition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validations`

The validations to perform on the value in the field.

_Required_: No

_Type_: Array of [FieldValidationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-form-fieldvalidationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AmplifyUIBuilder::Form

FieldInputConfig
