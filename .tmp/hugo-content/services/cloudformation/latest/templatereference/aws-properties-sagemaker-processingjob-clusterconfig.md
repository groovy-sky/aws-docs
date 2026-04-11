This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob ClusterConfig

Configuration for the cluster used to run a processing job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceCount" : Integer,
  "InstanceType" : String,
  "VolumeKmsKeyId" : String,
  "VolumeSizeInGB" : Integer
}

```

### YAML

```yaml

  InstanceCount: Integer
  InstanceType: String
  VolumeKmsKeyId: String
  VolumeSizeInGB: Integer

```

## Properties

`InstanceCount`

The number of ML compute instances to use in the processing job. For distributed
processing jobs, specify a value greater than 1. The default value is 1.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

The ML compute instance type for the processing job.

_Required_: Yes

_Type_: String

_Allowed values_: `ml.t3.medium | ml.t3.large | ml.t3.xlarge | ml.t3.2xlarge | ml.m4.xlarge | ml.m4.2xlarge | ml.m4.4xlarge | ml.m4.10xlarge | ml.m4.16xlarge | ml.c4.xlarge | ml.c4.2xlarge | ml.c4.4xlarge | ml.c4.8xlarge | ml.c5.xlarge | ml.c5.2xlarge | ml.c5.4xlarge | ml.c5.9xlarge | ml.c5.18xlarge | ml.m5.large | ml.m5.xlarge | ml.m5.2xlarge | ml.m5.4xlarge | ml.m5.12xlarge | ml.m5.24xlarge | ml.r5.large | ml.r5.xlarge | ml.r5.2xlarge | ml.r5.4xlarge | ml.r5.8xlarge | ml.r5.12xlarge | ml.r5.16xlarge | ml.r5.24xlarge | ml.g4dn.xlarge | ml.g4dn.2xlarge | ml.g4dn.4xlarge | ml.g4dn.8xlarge | ml.g4dn.12xlarge | ml.g4dn.16xlarge | ml.g5.xlarge | ml.g5.2xlarge | ml.g5.4xlarge | ml.g5.8xlarge | ml.g5.16xlarge | ml.g5.12xlarge | ml.g5.24xlarge | ml.g5.48xlarge | ml.r5d.large | ml.r5d.xlarge | ml.r5d.2xlarge | ml.r5d.4xlarge | ml.r5d.8xlarge | ml.r5d.12xlarge | ml.r5d.16xlarge | ml.r5d.24xlarge | ml.g6.xlarge | ml.g6.2xlarge | ml.g6.4xlarge | ml.g6.8xlarge | ml.g6.12xlarge | ml.g6.16xlarge | ml.g6.24xlarge | ml.g6.48xlarge | ml.g6e.xlarge | ml.g6e.2xlarge | ml.g6e.4xlarge | ml.g6e.8xlarge | ml.g6e.12xlarge | ml.g6e.16xlarge | ml.g6e.24xlarge | ml.g6e.48xlarge | ml.m6i.large | ml.m6i.xlarge | ml.m6i.2xlarge | ml.m6i.4xlarge | ml.m6i.8xlarge | ml.m6i.12xlarge | ml.m6i.16xlarge | ml.m6i.24xlarge | ml.m6i.32xlarge | ml.c6i.xlarge | ml.c6i.2xlarge | ml.c6i.4xlarge | ml.c6i.8xlarge | ml.c6i.12xlarge | ml.c6i.16xlarge | ml.c6i.24xlarge | ml.c6i.32xlarge | ml.m7i.large | ml.m7i.xlarge | ml.m7i.2xlarge | ml.m7i.4xlarge | ml.m7i.8xlarge | ml.m7i.12xlarge | ml.m7i.16xlarge | ml.m7i.24xlarge | ml.m7i.48xlarge | ml.c7i.large | ml.c7i.xlarge | ml.c7i.2xlarge | ml.c7i.4xlarge | ml.c7i.8xlarge | ml.c7i.12xlarge | ml.c7i.16xlarge | ml.c7i.24xlarge | ml.c7i.48xlarge | ml.r7i.large | ml.r7i.xlarge | ml.r7i.2xlarge | ml.r7i.4xlarge | ml.r7i.8xlarge | ml.r7i.12xlarge | ml.r7i.16xlarge | ml.r7i.24xlarge | ml.r7i.48xlarge`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeKmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker
uses to encrypt data on the storage volume attached to the ML compute instance(s) that
run the processing job.

###### Note

Certain Nitro-based instances include local storage, dependent on the instance
type. Local storage volumes are encrypted using a hardware module on the instance.
You can't request a `VolumeKmsKeyId` when using an instance type with
local storage.

For a list of instance types that support local instance storage, see [Instance Store Volumes](../../../ec2/latest/userguide/instancestorage.md#instance-store-volumes).

For more information about local instance storage encryption, see [SSD\
Instance Store Volumes](../../../ec2/latest/userguide/ssd-instance-store.md).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/_-]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeSizeInGB`

The size of the ML storage volume in gigabytes that you want to provision. You must
specify sufficient ML storage for your scenario.

###### Note

Certain Nitro-based instances include local storage with a fixed total size,
dependent on the instance type. When using these instances for processing, Amazon SageMaker
mounts the local instance storage instead of Amazon EBS gp2 storage. You can't request a
`VolumeSizeInGB` greater than the total size of the local instance
storage.

For a list of instance types that support local instance storage, including the
total size per instance type, see [Instance Store Volumes](../../../ec2/latest/userguide/instancestorage.md#instance-store-volumes).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AthenaDatasetDefinition

DatasetDefinition

All content copied from https://docs.aws.amazon.com/.
