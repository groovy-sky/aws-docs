This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule MonitoringScheduleConfig

Configures the monitoring schedule and defines the monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MonitoringJobDefinition" : MonitoringJobDefinition,
  "MonitoringJobDefinitionName" : String,
  "MonitoringType" : String,
  "ScheduleConfig" : ScheduleConfig
}

```

### YAML

```yaml

  MonitoringJobDefinition:
    MonitoringJobDefinition
  MonitoringJobDefinitionName: String
  MonitoringType: String
  ScheduleConfig:
    ScheduleConfig

```

## Properties

`MonitoringJobDefinition`

Defines the monitoring job.

_Required_: No

_Type_: [MonitoringJobDefinition](aws-properties-sagemaker-monitoringschedule-monitoringjobdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringJobDefinitionName`

The name of the monitoring job definition to schedule.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringType`

The type of the monitoring job definition to schedule.

_Required_: No

_Type_: String

_Allowed values_: `DataQuality | ModelQuality | ModelBias | ModelExplainability`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleConfig`

Configures the monitoring schedule.

_Required_: No

_Type_: [ScheduleConfig](aws-properties-sagemaker-monitoringschedule-scheduleconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringResources

NetworkConfig

All content copied from https://docs.aws.amazon.com/.
