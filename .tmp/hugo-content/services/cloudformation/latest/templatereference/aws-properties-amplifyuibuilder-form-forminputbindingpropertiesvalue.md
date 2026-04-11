This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Form FormInputBindingPropertiesValue

Represents the data binding configuration for a form's input fields at runtime.You can use
`FormInputBindingPropertiesValue` to add exposed properties to a form to allow
different values to be entered when a form is reused in different places in an app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BindingProperties" : FormInputBindingPropertiesValueProperties,
  "Type" : String
}

```

### YAML

```yaml

  BindingProperties:
    FormInputBindingPropertiesValueProperties
  Type: String

```

## Properties

`BindingProperties`

Describes the properties to customize with data at runtime.

_Required_: No

_Type_: [FormInputBindingPropertiesValueProperties](aws-properties-amplifyuibuilder-form-forminputbindingpropertiesvalueproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The property type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormDataTypeConfig

FormInputBindingPropertiesValueProperties

All content copied from https://docs.aws.amazon.com/.
