This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream EncryptionConfiguration

The `EncryptionConfiguration` property type specifies the encryption
settings that Amazon Kinesis Data Firehose (Kinesis Data Firehose) uses when delivering
data to Amazon Simple Storage Service (Amazon S3).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSEncryptionConfig" : KMSEncryptionConfig,
  "NoEncryptionConfig" : String
}

```

### YAML

```yaml

  KMSEncryptionConfig:
    KMSEncryptionConfig
  NoEncryptionConfig: String

```

## Properties

`KMSEncryptionConfig`

The AWS Key Management Service (AWS KMS) encryption
key that Amazon S3 uses to encrypt your data.

_Required_: No

_Type_: [KMSEncryptionConfig](aws-properties-kinesisfirehose-deliverystream-kmsencryptionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoEncryptionConfig`

Disables encryption. For valid values, see the `NoEncryptionConfig`
content for the [EncryptionConfiguration](../../../../reference/firehose/latest/apireference/api-encryptionconfiguration.md) data type in the _Amazon Kinesis Data Firehose_
_API Reference_.

_Required_: No

_Type_: String

_Allowed values_: `NoEncryption`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticsearchRetryOptions

ExtendedS3DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
