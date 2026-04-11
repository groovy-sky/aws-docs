This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingType Mqtt5Configuration

The configuration to add user-defined properties to enrich MQTT 5 messages.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PropagatingAttributes" : [ PropagatingAttribute, ... ]
}

```

### YAML

```yaml

  PropagatingAttributes:
    - PropagatingAttribute

```

## Properties

`PropagatingAttributes`

An object that represents the connection attribute, the thing attribute, and the MQTT 5 user property key.

_Required_: No

_Type_: Array of [PropagatingAttribute](aws-properties-iot-thingtype-propagatingattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::ThingType

PropagatingAttribute

All content copied from https://docs.aws.amazon.com/.
