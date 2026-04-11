This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel DetectorModelDefinition

Information that defines how a detector operates.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InitialStateName" : String,
  "States" : [ State, ... ]
}

```

### YAML

```yaml

  InitialStateName: String
  States:
    - State

```

## Properties

`InitialStateName`

The state that is entered at the creation of each detector (instance).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`States`

Information about the states of the detector.

_Required_: Yes

_Type_: Array of [State](aws-properties-iotevents-detectormodel-state.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClearTimer

DynamoDB

All content copied from https://docs.aws.amazon.com/.
