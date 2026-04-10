This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign StorageMinimumTimeToLive

Information about the minimum amount of time that data will be kept.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Unit: String
  Value: Integer

```

## Properties

`Unit`

The time increment type.

_Required_: Yes

_Type_: String

_Allowed values_: `HOURS | DAYS | WEEKS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The minimum amount of time to store the data.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageMaximumSize

Tag

All content copied from https://docs.aws.amazon.com/.
