This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer RetentionConfiguration

The configuration for a snapshot retention optimizer for Apache Iceberg tables.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IcebergConfiguration" : IcebergRetentionConfiguration
}

```

### YAML

```yaml

  IcebergConfiguration:
    IcebergRetentionConfiguration

```

## Properties

`IcebergConfiguration`

The configuration for an Iceberg snapshot retention optimizer.

_Required_: No

_Type_: [IcebergRetentionConfiguration](aws-properties-glue-tableoptimizer-icebergretentionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrphanFileDeletionConfiguration

TableOptimizerConfiguration

All content copied from https://docs.aws.amazon.com/.
