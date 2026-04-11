This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormInputValueProperty

The `FormInputValueProperty` property specifies the configuration for an input field on a form. Use
`FormInputValueProperty` to specify the values to render or bind by
default.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BindingProperties" : FormInputValuePropertyBindingProperties,
  "Concat" : [ FormInputValueProperty, ... ],
  "Value" : String
}

```

### YAML

```yaml

  BindingProperties:
    FormInputValuePropertyBindingProperties
  Concat:
    - FormInputValueProperty
  Value: String

```

## Properties

`BindingProperties`

The information to bind fields to data at runtime.

_Required_: No

_Type_: [FormInputValuePropertyBindingProperties](aws-properties-amplifyuibuilder-form-forminputvaluepropertybindingproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Concat`

A list of form properties to concatenate to create the value to assign to this field
property.

_Required_: No

_Type_: Array of [FormInputValueProperty](aws-properties-amplifyuibuilder-form-forminputvalueproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to assign to the input field.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormInputBindingPropertiesValueProperties

FormInputValuePropertyBindingProperties

All content copied from https://docs.aws.amazon.com/.
