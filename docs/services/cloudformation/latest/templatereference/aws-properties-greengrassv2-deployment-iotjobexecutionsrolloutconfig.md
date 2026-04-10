This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobExecutionsRolloutConfig

Contains information about the rollout configuration for a job. This configuration defines
the rate at which the job deploys a configuration to a fleet of target devices.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExponentialRate" : IoTJobExponentialRolloutRate,
  "MaximumPerMinute" : Integer
}

```

### YAML

```yaml

  ExponentialRate:
    IoTJobExponentialRolloutRate
  MaximumPerMinute: Integer

```

## Properties

`ExponentialRate`

The exponential rate to increase the job rollout rate.

_Required_: No

_Type_: [IoTJobExponentialRolloutRate](aws-properties-greengrassv2-deployment-iotjobexponentialrolloutrate.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaximumPerMinute`

The maximum number of devices that receive a pending job notification, per minute.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobAbortCriteria

IoTJobExponentialRolloutRate

All content copied from https://docs.aws.amazon.com/.
