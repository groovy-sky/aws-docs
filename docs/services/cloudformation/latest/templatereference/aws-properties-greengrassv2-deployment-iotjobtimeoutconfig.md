This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobTimeoutConfig

Contains information about the timeout configuration for a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InProgressTimeoutInMinutes" : Integer
}

```

### YAML

```yaml

  InProgressTimeoutInMinutes: Integer

```

## Properties

`InProgressTimeoutInMinutes`

The amount of time, in minutes, that devices have to complete the job. The timer starts
when the job status is set to `IN_PROGRESS`. If the job status doesn't change to a
terminal state before the time expires, then the job status is set to
`TIMED_OUT`.

The timeout interval must be between 1 minute and 7 days (10080 minutes).

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IoTJobRateIncreaseCriteria

SystemResourceLimits

All content copied from https://docs.aws.amazon.com/.
