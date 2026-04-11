This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::SignalCatalog Node

A general abstraction of a signal. A node can be specified as an actuator, attribute,
branch, or sensor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actuator" : Actuator,
  "Attribute" : Attribute,
  "Branch" : Branch,
  "Sensor" : Sensor
}

```

### YAML

```yaml

  Actuator:
    Actuator
  Attribute:
    Attribute
  Branch:
    Branch
  Sensor:
    Sensor

```

## Properties

`Actuator`

Information about a node specified as an actuator.

###### Note

An actuator is a digital representation of a vehicle device.

_Required_: No

_Type_: [Actuator](aws-properties-iotfleetwise-signalcatalog-actuator.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

Information about a node specified as an attribute.

###### Note

An attribute represents static information about a vehicle.

_Required_: No

_Type_: [Attribute](aws-properties-iotfleetwise-signalcatalog-attribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Branch`

Information about a node specified as a branch.

###### Note

A group of signals that are defined in a hierarchical structure.

_Required_: No

_Type_: [Branch](aws-properties-iotfleetwise-signalcatalog-branch.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sensor`

An input component that reports the environmental condition of a vehicle.

###### Note

You can collect data about fluid levels, temperatures, vibrations, or battery
voltage from sensors.

_Required_: No

_Type_: [Sensor](aws-properties-iotfleetwise-signalcatalog-sensor.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Branch

NodeCounts

All content copied from https://docs.aws.amazon.com/.
