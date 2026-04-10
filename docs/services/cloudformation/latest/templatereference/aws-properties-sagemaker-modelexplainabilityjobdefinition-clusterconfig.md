This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelExplainabilityJobDefinition ClusterConfig

The configuration for the cluster resources used to run the processing job.

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

The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify
a value greater than 1. The default value is 1.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceType`

The ML compute instance type for the processing job.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeKmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to
encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeSizeInGB`

The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML
storage for your scenario.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `16384`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchTransformInput

ConstraintsResource

All content copied from https://docs.aws.amazon.com/.
