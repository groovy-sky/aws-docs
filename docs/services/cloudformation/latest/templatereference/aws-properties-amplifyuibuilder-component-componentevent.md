This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmplifyUIBuilder::Component ComponentEvent

The `ComponentEvent` property specifies the configuration of an event. You can bind an event and a corresponding
action to a `Component` or a `ComponentChild`. A button click
is an example of an event.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "BindingEvent" : String,
  "Parameters" : ActionParameters
}

```

### YAML

```yaml

  Action: String
  BindingEvent: String
  Parameters:
    ActionParameters

```

## Properties

`Action`

The action to perform when a specific event is raised.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BindingEvent`

Binds an event to an action on a component. When you specify a `bindingEvent`,
the event is called when the action is performed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Describes information about the action.

_Required_: No

_Type_: [ActionParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplifyuibuilder-component-actionparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComponentDataConfiguration

ComponentProperty
