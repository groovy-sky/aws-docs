This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate MaintenanceWindow

An optional configuration within the `SchedulingConfig` to setup a recurring
maintenance window with a predetermined start time and duration for the rollout of a job
document to all devices in a target group for a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DurationInMinutes" : Integer,
  "StartTime" : String
}

```

### YAML

```yaml

  DurationInMinutes: Integer
  StartTime: String

```

## Properties

`DurationInMinutes`

Displays the duration of the next maintenance window.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1430`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartTime`

Displays the start time of the next maintenance window.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobExecutionsRolloutConfig

PresignedUrlConfig

All content copied from https://docs.aws.amazon.com/.
