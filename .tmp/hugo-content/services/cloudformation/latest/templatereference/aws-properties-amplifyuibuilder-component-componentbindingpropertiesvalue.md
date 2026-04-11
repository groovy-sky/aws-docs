This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentBindingPropertiesValue

The `ComponentBindingPropertiesValue` property specifies the data binding configuration for a component at runtime. You can use
`ComponentBindingPropertiesValue` to add exposed properties to a component to
allow different values to be entered when a component is reused in different places in an
app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BindingProperties" : ComponentBindingPropertiesValueProperties,
  "DefaultValue" : String,
  "Type" : String
}

```

### YAML

```yaml

  BindingProperties:
    ComponentBindingPropertiesValueProperties
  DefaultValue: String
  Type: String

```

## Properties

`BindingProperties`

Describes the properties to customize with data at runtime.

_Required_: No

_Type_: [ComponentBindingPropertiesValueProperties](aws-properties-amplifyuibuilder-component-componentbindingpropertiesvalueproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value of the property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The property type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionParameters

ComponentBindingPropertiesValueProperties

All content copied from https://docs.aws.amazon.com/.
