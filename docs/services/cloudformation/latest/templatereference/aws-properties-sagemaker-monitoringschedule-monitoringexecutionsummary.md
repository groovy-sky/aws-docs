This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule MonitoringExecutionSummary

Summary of information about the last monitoring job to run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreationTime" : String,
  "EndpointName" : String,
  "FailureReason" : String,
  "LastModifiedTime" : String,
  "MonitoringExecutionStatus" : String,
  "MonitoringScheduleName" : String,
  "ProcessingJobArn" : String,
  "ScheduledTime" : String
}

```

### YAML

```yaml

  CreationTime: String
  EndpointName: String
  FailureReason: String
  LastModifiedTime: String
  MonitoringExecutionStatus: String
  MonitoringScheduleName: String
  ProcessingJobArn: String
  ScheduledTime: String

```

## Properties

`CreationTime`

The time at which the monitoring job was created.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointName`

The name of the endpoint used to run the monitoring job.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureReason`

Contains the reason a monitoring job failed, if it failed.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastModifiedTime`

A timestamp that indicates the last time the monitoring job was modified.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringExecutionStatus`

The status of the monitoring job.

_Required_: Yes

_Type_: String

_Allowed values_: `Pending | Completed | CompletedWithViolations | InProgress | Failed | Stopping | Stopped`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringScheduleName`

The name of the monitoring schedule.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessingJobArn`

The Amazon Resource Name (ARN) of the monitoring job.

_Required_: No

_Type_: String

_Pattern_: `aws[a-z\-]*:sagemaker:[a-z0-9\-]*:[0-9]{12}:processing-job/.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledTime`

The time the monitoring job was scheduled.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MonitoringAppSpecification

MonitoringInput
