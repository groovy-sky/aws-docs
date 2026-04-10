This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::InstanceStorageConfig S3Config

Information about the Amazon Simple Storage Service (Amazon S3) storage type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "EncryptionConfig" : EncryptionConfig
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  EncryptionConfig:
    EncryptionConfig

```

## Properties

`BucketName`

The S3 bucket name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The S3 bucket prefix.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfig`

The Amazon S3 encryption configuration.

_Required_: No

_Type_: [EncryptionConfig](aws-properties-connect-instancestorageconfig-encryptionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisVideoStreamConfig

AWS::Connect::IntegrationAssociation

All content copied from https://docs.aws.amazon.com/.
