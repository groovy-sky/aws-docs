This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FieldInputConfig

The `FieldInputConfig` property specifies the configuration for the default input values to display for a field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultChecked" : Boolean,
  "DefaultCountryCode" : String,
  "DefaultValue" : String,
  "DescriptiveText" : String,
  "FileUploaderConfig" : FileUploaderFieldConfig,
  "IsArray" : Boolean,
  "MaxValue" : Number,
  "MinValue" : Number,
  "Name" : String,
  "Placeholder" : String,
  "ReadOnly" : Boolean,
  "Required" : Boolean,
  "Step" : Number,
  "Type" : String,
  "Value" : String,
  "ValueMappings" : ValueMappings
}

```

### YAML

```yaml

  DefaultChecked: Boolean
  DefaultCountryCode: String
  DefaultValue: String
  DescriptiveText: String
  FileUploaderConfig:
    FileUploaderFieldConfig
  IsArray: Boolean
  MaxValue: Number
  MinValue: Number
  Name: String
  Placeholder: String
  ReadOnly: Boolean
  Required: Boolean
  Step: Number
  Type: String
  Value: String
  ValueMappings:
    ValueMappings

```

## Properties

`DefaultChecked`

Specifies whether a field has a default value.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultCountryCode`

The default country code for a phone number.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value for the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DescriptiveText`

The text to display to describe the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileUploaderConfig`

The configuration for the file uploader field.

_Required_: No

_Type_: [FileUploaderFieldConfig](aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsArray`

Specifies whether to render the field as an array. This property is ignored if the
`dataSourceType` for the form is a Data Store.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxValue`

The maximum value to display for the field.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValue`

The minimum value to display for the field.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Placeholder`

The text to display as a placeholder for the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnly`

Specifies a read only field.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Required`

Specifies a field that requires input.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Step`

The stepping increment for a numeric value in a field.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The input type for the field.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueMappings`

The information to use to customize the input fields with data at runtime.

_Required_: No

_Type_: [ValueMappings](aws-properties-amplifyuibuilder-form-valuemappings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldConfig

FieldPosition

All content copied from https://docs.aws.amazon.com/.
