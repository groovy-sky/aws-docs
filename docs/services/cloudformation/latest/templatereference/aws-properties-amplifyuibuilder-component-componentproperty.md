This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentProperty

The `ComponentProperty` property specifies the configuration for all of a
component's properties. Use `ComponentProperty` to specify the values to render or
bind by default.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BindingProperties" : ComponentPropertyBindingProperties,
  "Bindings" : {Key: Value, ...},
  "CollectionBindingProperties" : ComponentPropertyBindingProperties,
  "ComponentName" : String,
  "Concat" : [ ComponentProperty, ... ],
  "Condition" : ComponentConditionProperty,
  "Configured" : Boolean,
  "DefaultValue" : String,
  "Event" : String,
  "ImportedValue" : String,
  "Model" : String,
  "Property" : String,
  "Type" : String,
  "UserAttribute" : String,
  "Value" : String
}

```

### YAML

```yaml

  BindingProperties:
    ComponentPropertyBindingProperties
  Bindings:
    Key: Value
  CollectionBindingProperties:
    ComponentPropertyBindingProperties
  ComponentName: String
  Concat:
    - ComponentProperty
  Condition:
    ComponentConditionProperty
  Configured: Boolean
  DefaultValue: String
  Event: String
  ImportedValue: String
  Model: String
  Property: String
  Type: String
  UserAttribute: String
  Value: String

```

## Properties

`BindingProperties`

The information to bind the component property to data at runtime.

_Required_: No

_Type_: [ComponentPropertyBindingProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentpropertybindingproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Bindings`

The information to bind the component property to form data.

_Required_: No

_Type_: Object of [FormBindingElement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-formbindingelement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollectionBindingProperties`

The information to bind the component property to data at runtime. Use this for collection
components.

_Required_: No

_Type_: [ComponentPropertyBindingProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentpropertybindingproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentName`

The name of the component that is affected by an event.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Concat`

A list of component properties to concatenate to create the value to assign to this
component property.

_Required_: No

_Type_: Array of [ComponentProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

The conditional expression to use to assign a value to the component property.

_Required_: No

_Type_: [ComponentConditionProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-componentconditionproperty.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configured`

Specifies whether the user configured the property in Amplify Studio after
importing it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value to assign to the component property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Event`

An event that occurs in your app. Use this for workflow data binding.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportedValue`

The default value assigned to the property when the component is imported into an
app.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The data model to use to assign a value to the component property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The name of the component's property that is affected by an event.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The component type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAttribute`

An authenticated user attribute to use to assign a value to the component property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to assign to the component property.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComponentEvent

ComponentPropertyBindingProperties
