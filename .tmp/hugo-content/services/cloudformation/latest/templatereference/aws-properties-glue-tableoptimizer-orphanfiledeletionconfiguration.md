This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer OrphanFileDeletionConfiguration

Configuration for removing files that are are not tracked by the Iceberg table metadata, and are older than your configured age limit. This configuration helps optimize storage usage and costs by automatically
cleaning up files that are no longer needed by the table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IcebergConfiguration" : IcebergConfiguration
}

```

### YAML

```yaml

  IcebergConfiguration:
    IcebergConfiguration

```

## Properties

`IcebergConfiguration`

The `IcebergConfiguration` property helps optimize your Iceberg tables in AWS Glue
by allowing you to specify format-specific settings that control how data is
stored, compressed, and managed.

_Required_: No

_Type_: [IcebergConfiguration](aws-properties-glue-tableoptimizer-icebergconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergRetentionConfiguration

RetentionConfiguration

All content copied from https://docs.aws.amazon.com/.
