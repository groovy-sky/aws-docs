This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application FlinkRunConfiguration

Describes the starting parameters for a Managed Service for Apache Flink application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowNonRestoredState" : Boolean
}

```

### YAML

```yaml

  AllowNonRestoredState: Boolean

```

## Properties

`AllowNonRestoredState`

When restoring from a snapshot, specifies whether the runtime is allowed to skip a
state that cannot be mapped to the new program. This will happen if the program is
updated between snapshots to remove stateful parameters, and state data in the snapshot
no longer corresponds to valid application data. For more information, see [Allowing Non-Restored State](https://nightlies.apache.org/flink/flink-docs-master/docs/ops/state/savepoints) in the [Apache Flink\
documentation](https://nightlies.apache.org/flink/flink-docs-master).

###### Note

This value defaults to `false`. If you update your application without
specifying this parameter, `AllowNonRestoredState` will be set to
`false`, even if it was previously set to `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlinkApplicationConfiguration

GlueDataCatalogConfiguration

All content copied from https://docs.aws.amazon.com/.
