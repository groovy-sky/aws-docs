This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer TableOptimizerConfiguration

Specifies configuration details of a table optimizer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "OrphanFileDeletionConfiguration" : OrphanFileDeletionConfiguration,
  "RetentionConfiguration" : RetentionConfiguration,
  "RoleArn" : String,
  "VpcConfiguration" : VpcConfiguration
}

```

### YAML

```yaml

  Enabled: Boolean
  OrphanFileDeletionConfiguration:
    OrphanFileDeletionConfiguration
  RetentionConfiguration:
    RetentionConfiguration
  RoleArn: String
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`Enabled`

Whether the table optimization is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrphanFileDeletionConfiguration`

`OrphanFileDeletionConfiguration` is a property that can be included within the
TableOptimizer resource. It controls the automatic deletion of orphaned
files - files that are not tracked by the table metadata, and older than the configured age limit.

_Required_: No

_Type_: [OrphanFileDeletionConfiguration](aws-properties-glue-tableoptimizer-orphanfiledeletionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionConfiguration`

The configuration for a snapshot retention optimizer for Apache Iceberg tables.

_Required_: No

_Type_: [RetentionConfiguration](aws-properties-glue-tableoptimizer-retentionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

A role passed by the caller which gives the service permission to update the resources associated with the optimizer on the caller's behalf.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

An object that describes the VPC configuration for a table optimizer. This configuration is necessary to perform optimization on tables that are in a customer VPC.

_Required_: No

_Type_: [VpcConfiguration](aws-properties-glue-tableoptimizer-vpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionConfiguration

VpcConfiguration

All content copied from https://docs.aws.amazon.com/.
