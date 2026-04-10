This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment DeploymentIoTJobConfiguration

Contains information about an AWS IoT job configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AbortConfig" : IoTJobAbortConfig,
  "JobExecutionsRolloutConfig" : IoTJobExecutionsRolloutConfig,
  "TimeoutConfig" : IoTJobTimeoutConfig
}

```

### YAML

```yaml

  AbortConfig:
    IoTJobAbortConfig
  JobExecutionsRolloutConfig:
    IoTJobExecutionsRolloutConfig
  TimeoutConfig:
    IoTJobTimeoutConfig

```

## Properties

`AbortConfig`

The stop configuration for the job. This configuration defines when and how to stop a job
rollout.

_Required_: No

_Type_: [IoTJobAbortConfig](aws-properties-greengrassv2-deployment-iotjobabortconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobExecutionsRolloutConfig`

The rollout configuration for the job. This configuration defines the rate at which the
job rolls out to the fleet of target devices.

_Required_: No

_Type_: [IoTJobExecutionsRolloutConfig](aws-properties-greengrassv2-deployment-iotjobexecutionsrolloutconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutConfig`

The timeout configuration for the job. This configuration defines the amount of time each
device has to complete the job.

_Required_: No

_Type_: [IoTJobTimeoutConfig](aws-properties-greengrassv2-deployment-iotjobtimeoutconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentConfigurationValidationPolicy

DeploymentPolicies

All content copied from https://docs.aws.amazon.com/.
