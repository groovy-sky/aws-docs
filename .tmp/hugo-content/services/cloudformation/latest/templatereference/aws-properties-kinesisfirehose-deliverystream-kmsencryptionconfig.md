This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream KMSEncryptionConfig

The `KMSEncryptionConfig` property type specifies the AWS
Key Management Service (AWS KMS) encryption key that Amazon Simple Storage
Service (Amazon S3) uses to encrypt data delivered by the Amazon Kinesis Data Firehose
(Kinesis Data Firehose) stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSKMSKeyARN" : String
}

```

### YAML

```yaml

  AWSKMSKeyARN: String

```

## Properties

`AWSKMSKeyARN`

The Amazon Resource Name (ARN) of the AWS KMS encryption key that
Amazon S3 uses to encrypt data delivered by the Kinesis Data Firehose stream. The key must
belong to the same region as the destination S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*:kms:[a-zA-Z0-9\-]+:\d{12}:key/[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisStreamSourceConfiguration

MSKSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
