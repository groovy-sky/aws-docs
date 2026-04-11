This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobRateIncreaseCriteria

Contains information about criteria to meet before a job increases its rollout rate.
Specify either `numberOfNotifiedThings` or
`numberOfSucceededThings`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberOfNotifiedThings" : Integer,
  "NumberOfSucceededThings" : Integer
}

```

### YAML

```yaml

  NumberOfNotifiedThings: Integer
  NumberOfSucceededThings: Integer

```

## Properties

`NumberOfNotifiedThings`

The number of devices to receive the job notification before the rollout rate
increases.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NumberOfSucceededThings`

The number of devices to successfully run the configuration job before the rollout rate
increases.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobExponentialRolloutRate

IoTJobTimeoutConfig

All content copied from https://docs.aws.amazon.com/.
