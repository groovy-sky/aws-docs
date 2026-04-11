This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application CheckpointConfiguration

Describes an application's checkpointing configuration. Checkpointing is the process
of persisting application state for fault tolerance. For more information, see [Checkpoints for Fault Tolerance](https://nightlies.apache.org/flink/flink-docs-master/docs/dev/datastream/fault-tolerance/checkpointing) in the [Apache Flink\
Documentation](https://nightlies.apache.org/flink/flink-docs-master).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CheckpointingEnabled" : Boolean,
  "CheckpointInterval" : Integer,
  "ConfigurationType" : String,
  "MinPauseBetweenCheckpoints" : Integer
}

```

### YAML

```yaml

  CheckpointingEnabled: Boolean
  CheckpointInterval: Integer
  ConfigurationType: String
  MinPauseBetweenCheckpoints: Integer

```

## Properties

`CheckpointingEnabled`

Describes whether checkpointing is enabled for a Managed Service for Apache Flink application.

###### Note

If `CheckpointConfiguration.ConfigurationType` is `DEFAULT`,
the application will use a `CheckpointingEnabled` value of `true`, even if this value
is set to another value using this API or in application code.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CheckpointInterval`

Describes the interval in milliseconds between checkpoint operations.

###### Note

If `CheckpointConfiguration.ConfigurationType` is `DEFAULT`,
the application will use a `CheckpointInterval` value of 60000, even if this value is set
to another value using this API or in application code.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `9223372036854775807`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationType`

Describes whether the application uses Managed Service for Apache Flink' default checkpointing behavior.
You must set this property to `CUSTOM` in order to set the
`CheckpointingEnabled`, `CheckpointInterval`, or `MinPauseBetweenCheckpoints` parameters.

###### Note

If this value is set to `DEFAULT`, the application will use the following values, even if they are set to other values using APIs or
application code:

- **CheckpointingEnabled:** true

- **CheckpointInterval:** 60000

- **MinPauseBetweenCheckpoints:** 5000

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinPauseBetweenCheckpoints`

Describes the minimum time in milliseconds after a checkpoint operation completes that
a new checkpoint operation can start. If a checkpoint operation takes longer than the
`CheckpointInterval`, the application otherwise performs continual
checkpoint operations. For more information, see [Tuning Checkpointing](https://nightlies.apache.org/flink/flink-docs-master/docs/ops/state/large_state_tuning) in the [Apache Flink\
Documentation](https://nightlies.apache.org/flink/flink-docs-master).

###### Note

If `CheckpointConfiguration.ConfigurationType` is `DEFAULT`,
the application will use a `MinPauseBetweenCheckpoints` value of 5000,
even if this value is set using this API or in application code.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `9223372036854775807`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CheckpointConfiguration](../../../managed-flink/latest/apiv2/api-checkpointconfiguration.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CatalogConfiguration

CodeContent

All content copied from https://docs.aws.amazon.com/.
