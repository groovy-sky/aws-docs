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
VPC](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).

_Required_: No

_Type_: [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-comprehend-flywheel-vpcconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Comprehend::Flywheel

DocumentClassificationConfig
