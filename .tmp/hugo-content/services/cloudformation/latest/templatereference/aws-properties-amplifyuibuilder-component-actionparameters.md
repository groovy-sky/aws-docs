This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ActionParameters

Represents the event action configuration for an element of a `Component` or
`ComponentChild`. Use for the workflow feature in Amplify Studio
that allows you to bind events and actions to components. `ActionParameters`
defines the action that is performed when an event occurs on the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Anchor" : ComponentProperty,
  "Fields" : {Key: Value, ...},
  "Global" : ComponentProperty,
  "Id" : ComponentProperty,
  "Model" : String,
  "State" : MutationActionSetStateParameter,
  "Target" : ComponentProperty,
  "Type" : ComponentProperty,
  "Url" : ComponentProperty
}

```

### YAML

```yaml

  Anchor:
    ComponentProperty
  Fields:
    Key: Value
  Global:
    ComponentProperty
  Id:
    ComponentProperty
  Model: String
  State:
    MutationActionSetStateParameter
  Target:
    ComponentProperty
  Type:
    ComponentProperty
  Url:
    ComponentProperty

```

## Properties

`Anchor`

The HTML anchor link to the location to open. Specify this value for a navigation
action.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fields`

A dictionary of key-value pairs mapping Amplify Studio properties to fields
in a data model. Use when the action performs an operation on an Amplify
DataStore model.

_Required_: No

_Type_: Object of [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Global`

Specifies whether the user should be signed out globally. Specify this value for an auth
sign out action.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The unique ID of the component that the `ActionParameters` apply to.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Model`

The name of the data model. Use when the action performs an operation on an Amplify DataStore model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

A key-value pair that specifies the state property name and its initial value.

_Required_: No

_Type_: [MutationActionSetStateParameter](aws-properties-amplifyuibuilder-component-mutationactionsetstateparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The element within the same component to modify when the action occurs.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of navigation action. Valid values are `url` and `anchor`.
This value is required for a navigation action.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL to the location to open. Specify this value for a navigation action.

_Required_: No

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AmplifyUIBuilder::Component

ComponentBindingPropertiesValue

All content copied from https://docs.aws.amazon.com/.
