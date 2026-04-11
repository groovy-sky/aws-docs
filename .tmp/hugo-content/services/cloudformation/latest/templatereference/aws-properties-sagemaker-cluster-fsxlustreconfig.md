This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster FSxLustreConfig

Configuration settings for an Amazon FSx for Lustre file system to be used with the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PerUnitStorageThroughput" : Integer,
  "SizeInGiB" : Integer
}

```

### YAML

```yaml

  PerUnitStorageThroughput: Integer
  SizeInGiB: Integer

```

## Properties

`PerUnitStorageThroughput`

The throughput capacity of the Amazon FSx for Lustre file system, measured in MB/s per TiB of
storage.

_Required_: Yes

_Type_: Integer

_Minimum_: `125`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizeInGiB`

The storage capacity of the Amazon FSx for Lustre file system, specified in gibibytes (GiB).

_Required_: Yes

_Type_: Integer

_Minimum_: `1200`

_Maximum_: `100800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentConfig

Orchestrator

All content copied from https://docs.aws.amazon.com/.
