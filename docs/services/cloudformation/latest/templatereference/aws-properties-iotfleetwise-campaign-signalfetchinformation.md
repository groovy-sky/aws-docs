This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign SignalFetchInformation

Information about the signal to be fetched.

###### Important

AWS IoT FleetWise will no longer be open to new customers starting April 30, 2026.
If you would like to use AWS IoT FleetWise, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see [AWS IoT FleetWise availability change](../../../iot-fleetwise/latest/developerguide/iotfleetwise-availability-change.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "ConditionLanguageVersion" : Number,
  "FullyQualifiedName" : String,
  "SignalFetchConfig" : SignalFetchConfig
}

```

### YAML

```yaml

  Actions:
    - String
  ConditionLanguageVersion: Number
  FullyQualifiedName: String
  SignalFetchConfig:
    SignalFetchConfig

```

## Properties

`Actions`

The actions to be performed by the signal fetch.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `2048 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConditionLanguageVersion`

The version of the condition language used.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FullyQualifiedName`

The fully qualified name of the signal to be fetched.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.]+$`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignalFetchConfig`

The configuration of the signal fetch operation.

_Required_: Yes

_Type_: [SignalFetchConfig](aws-properties-iotfleetwise-campaign-signalfetchconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignalFetchConfig

SignalInformation

All content copied from https://docs.aws.amazon.com/.
