This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel DataSecurityConfig

Data security configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataLakeKmsKeyId" : String,
  "ModelKmsKeyId" : String,
  "VolumeKmsKeyId" : String,
  "VpcConfig" : VpcConfig
}

```

### YAML

```yaml

  DataLakeKmsKeyId: String
  ModelKmsKeyId: String
  VolumeKmsKeyId: String
  VpcConfig:
    VpcConfig

```

## Properties

`DataLakeKmsKeyId`

ID for the AWS KMS key that Amazon Comprehend uses to encrypt the data in the data lake.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelKmsKeyId`

ID for the AWS KMS key that Amazon Comprehend uses to encrypt
trained custom models. The ModelKmsKeyId can be either of the following formats:

- KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`

- Amazon Resource Name (ARN) of a KMS Key:
`"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeKmsKeyId`

ID for the AWS KMS key that Amazon Comprehend uses to encrypt the volume.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfig`

Configuration parameters for an optional private Virtual Private Cloud (VPC) containing
the resources you are using for the job. For more information, see [Amazon\
VPC](../../../vpc/latest/userguide/what-is-amazon-vpc.md).

_Required_: No

_Type_: [VpcConfig](aws-properties-comprehend-flywheel-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Comprehend::Flywheel

DocumentClassificationConfig

All content copied from https://docs.aws.amazon.com/.
