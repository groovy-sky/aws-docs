This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component MutationActionSetStateParameter

Represents the state configuration when an action modifies a property of another element
within the same component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentName" : String,
  "Property" : String,
  "Set" : ComponentProperty
}

```

### YAML

```yaml

  ComponentName: String
  Property: String
  Set:
    ComponentProperty

```

## Properties

`ComponentName`

The name of the component that is being modified.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The name of the component property to apply the state configuration to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Set`

The state configuration to assign to the property.

_Required_: Yes

_Type_: [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormBindingElement

Predicate

All content copied from https://docs.aws.amazon.com/.
