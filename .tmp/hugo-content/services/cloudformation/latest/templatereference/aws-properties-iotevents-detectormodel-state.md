This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel State

Information that defines a state of a detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnEnter" : OnEnter,
  "OnExit" : OnExit,
  "OnInput" : OnInput,
  "StateName" : String
}

```

### YAML

```yaml

  OnEnter:
    OnEnter
  OnExit:
    OnExit
  OnInput:
    OnInput
  StateName: String

```

## Properties

`OnEnter`

When entering this state, perform these `actions` if the `condition`
is TRUE.

_Required_: No

_Type_: [OnEnter](aws-properties-iotevents-detectormodel-onenter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnExit`

When exiting this state, perform these `actions` if the specified
`condition` is `TRUE`.

_Required_: No

_Type_: [OnExit](aws-properties-iotevents-detectormodel-onexit.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnInput`

When an input is received and the `condition` is TRUE, perform the specified
`actions`.

_Required_: No

_Type_: [OnInput](aws-properties-iotevents-detectormodel-oninput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateName`

The name of the state.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sqs

Tag

All content copied from https://docs.aws.amazon.com/.
