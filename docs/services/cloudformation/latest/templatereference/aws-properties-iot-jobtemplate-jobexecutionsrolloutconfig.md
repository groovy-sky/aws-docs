This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate JobExecutionsRolloutConfig

Allows you to create a staged rollout of a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExponentialRolloutRate" : ExponentialRolloutRate,
  "MaximumPerMinute" : Integer
}

```

### YAML

```yaml

  ExponentialRolloutRate:
    ExponentialRolloutRate
  MaximumPerMinute: Integer

```

## Properties

`ExponentialRolloutRate`

The rate of increase for a job rollout. This parameter allows you to define an
exponential rate for a job rollout.

_Required_: No

_Type_: [ExponentialRolloutRate](aws-properties-iot-jobtemplate-exponentialrolloutrate.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaximumPerMinute`

The maximum number of things that will be notified of a pending job, per minute.
This parameter allows you to create a staged rollout.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobExecutionsRetryConfig

MaintenanceWindow

All content copied from https://docs.aws.amazon.com/.
