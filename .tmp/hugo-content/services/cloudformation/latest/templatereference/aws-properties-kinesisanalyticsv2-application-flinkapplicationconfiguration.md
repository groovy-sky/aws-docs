This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application FlinkApplicationConfiguration

Describes configuration parameters for a Managed Service for Apache Flink application or a Studio notebook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CheckpointConfiguration" : CheckpointConfiguration,
  "MonitoringConfiguration" : MonitoringConfiguration,
  "ParallelismConfiguration" : ParallelismConfiguration
}

```

### YAML

```yaml

  CheckpointConfiguration:
    CheckpointConfiguration
  MonitoringConfiguration:
    MonitoringConfiguration
  ParallelismConfiguration:
    ParallelismConfiguration

```

## Properties

`CheckpointConfiguration`

Describes an application's checkpointing configuration. Checkpointing is the process
of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://ci.apache.org/projects/flink/flink-docs-release-1.8/concepts/programming-model.html) in the [Apache Flink\
Documentation](https://ci.apache.org/projects/flink/flink-docs-release-1.8).

_Required_: No

_Type_: [CheckpointConfiguration](aws-properties-kinesisanalyticsv2-application-checkpointconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitoringConfiguration`

Describes configuration parameters for Amazon CloudWatch logging for an
application.

_Required_: No

_Type_: [MonitoringConfiguration](aws-properties-kinesisanalyticsv2-application-monitoringconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelismConfiguration`

Describes parameters for how an application executes multiple tasks simultaneously.

_Required_: No

_Type_: [ParallelismConfiguration](aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [FlinkApplicationConfiguration](../../../managed-flink/latest/apiv2/api-flinkapplicationconfiguration.md) in the _Amazon Kinesis_
_Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentProperties

FlinkRunConfiguration

All content copied from https://docs.aws.amazon.com/.
