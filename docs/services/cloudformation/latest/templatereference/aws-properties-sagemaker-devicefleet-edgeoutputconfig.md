This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::DeviceFleet EdgeOutputConfig

The output configuration for storing sample data collected by the fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "S3OutputLocation" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  S3OutputLocation: String

```

## Properties

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to
encrypt data on the storage volume after compilation job. If you don't provide a KMS key ID, Amazon SageMaker uses
the default KMS key for Amazon S3 for your role's account.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputLocation`

The Amazon Simple Storage (S3) bucket URI.

_Required_: Yes

_Type_: String

_Pattern_: `^s3://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SageMaker::DeviceFleet

Tag
