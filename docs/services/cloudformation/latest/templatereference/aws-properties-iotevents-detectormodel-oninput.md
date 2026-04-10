This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel OnInput

Specifies the actions performed when the `condition` evaluates to TRUE.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Events" : [ Event, ... ],
  "TransitionEvents" : [ TransitionEvent, ... ]
}

```

### YAML

```yaml

  Events:
    - Event
  TransitionEvents:
    - TransitionEvent

```

## Properties

`Events`

Specifies the actions performed when the `condition` evaluates to TRUE.

_Required_: No

_Type_: Array of [Event](aws-properties-iotevents-detectormodel-event.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitionEvents`

Specifies the actions performed, and the next state entered, when a `condition`
evaluates to TRUE.

_Required_: No

_Type_: Array of [TransitionEvent](aws-properties-iotevents-detectormodel-transitionevent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnExit

Payload

All content copied from https://docs.aws.amazon.com/.
