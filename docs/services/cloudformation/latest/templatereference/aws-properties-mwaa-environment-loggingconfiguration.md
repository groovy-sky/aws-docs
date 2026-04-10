This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAA::Environment LoggingConfiguration

The type of Apache Airflow logs to send to CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DagProcessingLogs" : ModuleLoggingConfiguration,
  "SchedulerLogs" : ModuleLoggingConfiguration,
  "TaskLogs" : ModuleLoggingConfiguration,
  "WebserverLogs" : ModuleLoggingConfiguration,
  "WorkerLogs" : ModuleLoggingConfiguration
}

```

### YAML

```yaml

  DagProcessingLogs:
    ModuleLoggingConfiguration
  SchedulerLogs:
    ModuleLoggingConfiguration
  TaskLogs:
    ModuleLoggingConfiguration
  WebserverLogs:
    ModuleLoggingConfiguration
  WorkerLogs:
    ModuleLoggingConfiguration

```

## Properties

`DagProcessingLogs`

Defines the processing logs sent to CloudWatch Logs and the logging level to send.

_Required_: No

_Type_: [ModuleLoggingConfiguration](aws-properties-mwaa-environment-moduleloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchedulerLogs`

Defines the scheduler logs sent to CloudWatch Logs and the logging level to send.

_Required_: No

_Type_: [ModuleLoggingConfiguration](aws-properties-mwaa-environment-moduleloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskLogs`

Defines the task logs sent to CloudWatch Logs and the logging level to send.

_Required_: No

_Type_: [ModuleLoggingConfiguration](aws-properties-mwaa-environment-moduleloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebserverLogs`

Defines the web server logs sent to CloudWatch Logs and the logging level to send.

_Required_: No

_Type_: [ModuleLoggingConfiguration](aws-properties-mwaa-environment-moduleloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerLogs`

Defines the worker logs sent to CloudWatch Logs and the logging level to send.

_Required_: No

_Type_: [ModuleLoggingConfiguration](aws-properties-mwaa-environment-moduleloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MWAA::Environment

ModuleLoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
