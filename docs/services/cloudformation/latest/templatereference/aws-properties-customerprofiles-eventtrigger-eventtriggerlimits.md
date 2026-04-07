This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger EventTriggerLimits

Defines limits controlling whether an event triggers the destination, based on
ingestion latency and the number of invocations per profile over specific time
periods.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventExpiration" : Integer,
  "Periods" : [ Period, ... ]
}

```

### YAML

```yaml

  EventExpiration: Integer
  Periods:
    - Period

```

## Properties

`EventExpiration`

Specifies that an event will only trigger the destination if it is processed within a
certain latency period.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Periods`

A list of time periods during which the limits apply.

_Required_: No

_Type_: Array of [Period](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-eventtrigger-period.html)

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventTriggerDimension

ObjectAttribute
