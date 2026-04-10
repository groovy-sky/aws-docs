This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobExponentialRolloutRate

Contains information about an exponential rollout rate for a configuration deployment
job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BaseRatePerMinute" : Integer,
  "IncrementFactor" : Number,
  "RateIncreaseCriteria" : IoTJobRateIncreaseCriteria
}

```

### YAML

```yaml

  BaseRatePerMinute: Integer
  IncrementFactor: Number
  RateIncreaseCriteria:
    IoTJobRateIncreaseCriteria

```

## Properties

`BaseRatePerMinute`

The minimum number of devices that receive a pending job notification, per minute, when
the job starts. This parameter defines the initial rollout rate of the job.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IncrementFactor`

The exponential factor to increase the rollout rate for the job.

This parameter supports up to one digit after the decimal (for example, you can specify
`1.5`, but not `1.55`).

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RateIncreaseCriteria`

The criteria to increase the rollout rate for the job.

_Required_: Yes

_Type_: [IoTJobRateIncreaseCriteria](aws-properties-greengrassv2-deployment-iotjobrateincreasecriteria.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobExecutionsRolloutConfig

IoTJobRateIncreaseCriteria

All content copied from https://docs.aws.amazon.com/.
