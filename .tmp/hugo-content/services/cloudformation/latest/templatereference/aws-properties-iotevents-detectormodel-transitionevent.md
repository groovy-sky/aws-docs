This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel TransitionEvent

Specifies the actions performed and the next state entered when a `condition`
evaluates to TRUE.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ Action, ... ],
  "Condition" : String,
  "EventName" : String,
  "NextState" : String
}

```

### YAML

```yaml

  Actions:
    - Action
  Condition: String
  EventName: String
  NextState: String

```

## Properties

`Actions`

The actions to be performed.

_Required_: No

_Type_: Array of [Action](aws-properties-iotevents-detectormodel-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

Required. A Boolean expression that when TRUE causes the actions to be performed and the
`nextState` to be entered.

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventName`

The name of the transition event.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NextState`

The next state to enter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoTEvents::Input

All content copied from https://docs.aws.amazon.com/.
