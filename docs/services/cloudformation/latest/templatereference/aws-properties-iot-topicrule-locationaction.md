This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule LocationAction

Describes an action to send device location updates from an MQTT message to an Amazon
Location tracker resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeviceId" : String,
  "Latitude" : String,
  "Longitude" : String,
  "RoleArn" : String,
  "Timestamp" : Timestamp,
  "TrackerName" : String
}

```

### YAML

```yaml

  DeviceId: String
  Latitude: String
  Longitude: String
  RoleArn: String
  Timestamp:
    Timestamp
  TrackerName: String

```

## Properties

`DeviceId`

The unique ID of the device providing the location data.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Latitude`

A string that evaluates to a double value that represents the latitude of the device's
location.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Longitude`

A string that evaluates to a double value that represents the longitude of the device's
location.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that grants permission to write to the Amazon Location resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

The time that the location data was sampled. The default value is the time the MQTT
message was processed.

_Required_: No

_Type_: [Timestamp](aws-properties-iot-topicrule-timestamp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackerName`

The name of the tracker resource in Amazon Location in which the location is
updated.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaAction

OpenSearchAction

All content copied from https://docs.aws.amazon.com/.
