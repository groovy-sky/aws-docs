This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign DataPartitionUploadOptions

The upload options for the data partition. If upload options are specified, you must
also specify storage options. See [DataPartitionStorageOptions](../../../../reference/iot-fleetwise/latest/apireference/api-datapartitionstorageoptions.md).

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionLanguageVersion" : Integer,
  "Expression" : String
}

```

### YAML

```yaml

  ConditionLanguageVersion: Integer
  Expression: String

```

## Properties

`ConditionLanguageVersion`

The version of the condition language. Defaults to the most recent condition language
version.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Expression`

The logical expression used to recognize what data to collect. For example, ``$variable.`Vehicle.OutsideAirTemperature` >= 105.0``.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPartitionStorageOptions

MqttTopicConfig

All content copied from https://docs.aws.amazon.com/.
