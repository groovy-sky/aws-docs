This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster TieredStorageConfig

Defines the configuration for managed tier checkpointing in a HyperPod cluster.
Managed tier checkpointing uses multiple storage tiers, including cluster CPU memory, to
provide faster checkpoint operations and improved fault tolerance for large-scale model
training. The system automatically saves checkpoints at high frequency to memory and
periodically persists them to durable storage, like Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceMemoryAllocationPercentage" : Integer,
  "Mode" : String
}

```

### YAML

```yaml

  InstanceMemoryAllocationPercentage: Integer
  Mode: String

```

## Properties

`InstanceMemoryAllocationPercentage`

The percentage (int) of cluster memory to allocate for checkpointing.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Specifies whether managed tier checkpointing is enabled or disabled for the HyperPod
cluster. When set to `Enable`, the system installs a memory management daemon
that provides disaggregated memory as a service for checkpoint storage. When set to
`Disable`, the feature is turned off and the memory management daemon is
removed from the cluster.

_Required_: Yes

_Type_: String

_Allowed values_: `Enable | Disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

VpcConfig
