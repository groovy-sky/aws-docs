This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign ConditionBasedCollectionScheme

Information about a collection scheme that uses a simple logical expression to
recognize what data to collect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionLanguageVersion" : Integer,
  "Expression" : String,
  "MinimumTriggerIntervalMs" : Number,
  "TriggerMode" : String
}

```

### YAML

```yaml

  ConditionLanguageVersion: Integer
  Expression: String
  MinimumTriggerIntervalMs: Number
  TriggerMode: String

```

## Properties

`ConditionLanguageVersion`

Specifies the version of the conditional expression language.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Expression`

The logical expression used to recognize what data to collect. For example,
`$variable.Vehicle.OutsideAirTemperature >= 105.0`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumTriggerIntervalMs`

The minimum duration of time between two triggering events to collect data, in
milliseconds.

###### Note

If a signal changes often, you might want to collect data at a slower rate.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TriggerMode`

Whether to collect data for all triggering events ( `ALWAYS`). Specify
( `RISING_EDGE`), or specify only when the condition first evaluates to
false. For example, triggering on "AirbagDeployed"; Users aren't interested on
triggering when the airbag is already exploded; they only care about the change from not
deployed => deployed.

_Required_: No

_Type_: String

_Allowed values_: `ALWAYS | RISING_EDGE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CollectionScheme

ConditionBasedSignalFetchConfig

All content copied from https://docs.aws.amazon.com/.
