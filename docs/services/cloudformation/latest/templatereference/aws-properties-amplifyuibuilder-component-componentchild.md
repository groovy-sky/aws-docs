This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentChild

The `ComponentChild` property specifies a nested UI configuration within a
parent `Component`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Children" : [ ComponentChild, ... ],
  "ComponentType" : String,
  "Events" : {Key: Value, ...},
  "Name" : String,
  "Properties" : {Key: Value, ...},
  "SourceId" : String
}

```

### YAML

```yaml

  Children:
    - ComponentChild
  ComponentType: String
  Events:
    Key: Value
  Name: String
  Properties:
    Key: Value
  SourceId: String

```

## Properties

`Children`

The list of `ComponentChild` instances for this component.

_Required_: No

_Type_: Array of [ComponentChild](aws-properties-amplifyuibuilder-component-componentchild.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentType`

The type of the child component.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Events`

Describes the events that can be raised on the child component. Use for the workflow
feature in Amplify Studio that allows you to bind events and actions to
components.

_Required_: No

_Type_: Object of [ComponentEvent](aws-properties-amplifyuibuilder-component-componentevent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the child component.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

Describes the properties of the child component. You can't specify `tags` as a
valid property for `properties`.

_Required_: Yes

_Type_: Object of [ComponentProperty](aws-properties-amplifyuibuilder-component-componentproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceId`

The unique ID of the child component in its original source system, such as Figma.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentBindingPropertiesValueProperties

ComponentConditionProperty

All content copied from https://docs.aws.amazon.com/.
