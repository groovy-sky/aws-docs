This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign SignalFetchConfig

The configuration of the signal fetch operation.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionBased" : ConditionBasedSignalFetchConfig,
  "TimeBased" : TimeBasedSignalFetchConfig
}

```

### YAML

```yaml

  ConditionBased:
    ConditionBasedSignalFetchConfig
  TimeBased:
    TimeBasedSignalFetchConfig

```

## Properties

`ConditionBased`

The configuration of a condition-based signal fetch operation.

_Required_: No

_Type_: [ConditionBasedSignalFetchConfig](aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeBased`

The configuration of a time-based signal fetch operation.

_Required_: No

_Type_: [TimeBasedSignalFetchConfig](aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Config

SignalFetchInformation

All content copied from https://docs.aws.amazon.com/.
