This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::SamplingRule SamplingRateBoost

Enable temporary sampling rate increases when you detect anomalies to improve visibility.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CooldownWindowMinutes" : Integer,
  "MaxRate" : Number
}

```

### YAML

```yaml

  CooldownWindowMinutes: Integer
  MaxRate: Number

```

## Properties

`CooldownWindowMinutes`

Sets the time window (in minutes) in which only one sampling rate boost can be triggered.
After a boost occurs, no further boosts are allowed until the next window.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRate`

Defines max temporary sampling rate to apply when a boost is triggered.
Calculated boost rate by X-Ray will be less than or equal to this max rate.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::XRay::SamplingRule

SamplingRule

All content copied from https://docs.aws.amazon.com/.
