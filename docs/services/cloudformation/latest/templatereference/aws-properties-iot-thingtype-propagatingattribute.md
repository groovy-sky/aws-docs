This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::ThingType PropagatingAttribute

An object that represents the connection attribute, the thing attribute, and the MQTT 5 user property key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectionAttribute" : String,
  "ThingAttribute" : String,
  "UserPropertyKey" : String
}

```

### YAML

```yaml

  ConnectionAttribute: String
  ThingAttribute: String
  UserPropertyKey: String

```

## Properties

`ConnectionAttribute`

The attribute associated with the connection details.

_Required_: No

_Type_: String

_Allowed values_: `iot:ClientId | iot:Thing.ThingName`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThingAttribute`

The thing attribute that is propagating for MQTT 5 message enrichment.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.,@/:#-]+`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPropertyKey`

The key of the MQTT 5 user property, which is a key-value pair.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9:$.]+`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Mqtt5Configuration

Tag

All content copied from https://docs.aws.amazon.com/.
