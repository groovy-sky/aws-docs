This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate ExponentialRolloutRate

Allows you to create an exponential rate of rollout for a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseRatePerMinute" : Integer,
  "IncrementFactor" : Number,
  "RateIncreaseCriteria" : RateIncreaseCriteria
}

```

### YAML

```yaml

  BaseRatePerMinute: Integer
  IncrementFactor: Number
  RateIncreaseCriteria:
    RateIncreaseCriteria

```

## Properties

`BaseRatePerMinute`

The minimum number of things that will be notified of a pending job, per minute at
the start of job rollout. This parameter allows you to define the initial rate of rollout.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IncrementFactor`

The exponential factor to increase the rate of rollout for a job.

AWS IoT Core supports up to one digit after the decimal (for example, 1.5, but
not 1.55).

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RateIncreaseCriteria`

The criteria to initiate the increase in rate of rollout for a job.

_Required_: Yes

_Type_: [RateIncreaseCriteria](aws-properties-iot-jobtemplate-rateincreasecriteria.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AbortCriteria

JobExecutionsRetryConfig

All content copied from https://docs.aws.amazon.com/.
