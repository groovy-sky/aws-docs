This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel OnEnter

When entering this state, perform these `actions` if the `condition`
is TRUE.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Events" : [ Event, ... ]
}

```

### YAML

```yaml

  Events:
    - Event

```

## Properties

`Events`

Specifies the actions that are performed when the state is entered and the
`condition` is `TRUE`.

_Required_: No

_Type_: Array of [Event](aws-properties-iotevents-detectormodel-event.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Lambda

OnExit

All content copied from https://docs.aws.amazon.com/.
