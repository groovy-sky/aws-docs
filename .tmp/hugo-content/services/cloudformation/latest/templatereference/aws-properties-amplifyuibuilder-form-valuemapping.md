This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form ValueMapping

The `ValueMapping` property specifies the association between a complex object and a display value. Use `ValueMapping` to store
how to represent complex objects when they are displayed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayValue" : FormInputValueProperty,
  "Value" : FormInputValueProperty
}

```

### YAML

```yaml

  DisplayValue:
    FormInputValueProperty
  Value:
    FormInputValueProperty

```

## Properties

`DisplayValue`

The value to display for the complex object.

_Required_: No

_Type_: [FormInputValueProperty](aws-properties-amplifyuibuilder-form-forminputvalueproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The complex object.

_Required_: Yes

_Type_: [FormInputValueProperty](aws-properties-amplifyuibuilder-form-forminputvalueproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SectionalElement

ValueMappings

All content copied from https://docs.aws.amazon.com/.
