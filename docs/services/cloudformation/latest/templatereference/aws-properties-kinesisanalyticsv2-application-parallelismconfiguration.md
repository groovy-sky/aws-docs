This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application ParallelismConfiguration

Describes parameters for how a Flink-based Kinesis Data Analytics application executes
multiple tasks simultaneously. For more information about parallelism, see [Parallel Execution](https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/execution/parallel) in the [Apache Flink\
Documentation](https://nightlies.apache.org/flink/flink-docs-master).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingEnabled" : Boolean,
  "ConfigurationType" : String,
  "Parallelism" : Integer,
  "ParallelismPerKPU" : Integer
}

```

### YAML

```yaml

  AutoScalingEnabled: Boolean
  ConfigurationType: String
  Parallelism: Integer
  ParallelismPerKPU: Integer

```

## Properties

`AutoScalingEnabled`

Describes whether the Managed Service for Apache Flink service can increase the parallelism of the application in response to increased throughput.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationType`

Describes whether the application uses the default parallelism for the Managed Service for Apache Flink service. You must set this property to `CUSTOM`
in order to change your application's `AutoScalingEnabled`, `Parallelism`, or `ParallelismPerKPU` properties.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOM | DEFAULT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parallelism`

Describes the initial number of parallel tasks that a Java-based Kinesis Data
Analytics application can perform. The Kinesis Data Analytics service can increase this
number automatically if [ParallelismConfiguration:AutoScalingEnabled](../../../managed-flink/latest/apiv2/api-parallelismconfiguration.md#kinesisanalytics-Type-ParallelismConfiguration-AutoScalingEnabled.html) is set to
`true`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParallelismPerKPU`

Describes the number of parallel tasks that a Java-based Kinesis Data Analytics
application can perform per Kinesis Processing Unit (KPU) used by the application. For
more information about KPUs, see [Amazon Kinesis Data Analytics\
Pricing](https://aws.amazon.com/kinesis/data-analytics/pricing).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ParallelismConfiguration](../../../managed-flink/latest/apiv2/api-parallelismconfiguration.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringConfiguration

PropertyGroup

All content copied from https://docs.aws.amazon.com/.
