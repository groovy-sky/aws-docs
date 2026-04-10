This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application WorkerConfiguration

The configuration of a worker. For more information,
see [Supported worker configurations](../../../emr/latest/emr-serverless-userguide/app-behavior.md#worker-configs).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cpu" : String,
  "Disk" : String,
  "DiskType" : String,
  "Memory" : String
}

```

### YAML

```yaml

  Cpu: String
  Disk: String
  DiskType: String
  Memory: String

```

## Properties

`Cpu`

The CPU requirements of the worker configuration. Each worker can have 1, 2, 4, 8, or 16 vCPUs.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(vCPU|vcpu|VCPU)?$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Disk`

The disk requirements of the worker configuration.

_Required_: No

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(GB|gb|gB|Gb)$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DiskType`

The disk type for every worker instance of the work type. Shuffle optimized disks have higher
performance characteristics and are better for shuffle heavy workloads. Default is `STANDARD`.

_Required_: No

_Type_: String

_Pattern_: `^(SHUFFLE_OPTIMIZED|[Ss]huffle_[Oo]ptimized|STANDARD|[Ss]tandard)$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Memory`

The memory requirements of the worker configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]*(\s)?(GB|gb|gB|Gb)?$`

_Minimum_: `1`

_Maximum_: `15`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WorkerTypeSpecificationInput

All content copied from https://docs.aws.amazon.com/.
