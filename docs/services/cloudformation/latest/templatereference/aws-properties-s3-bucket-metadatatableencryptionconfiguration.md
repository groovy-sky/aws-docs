This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket MetadataTableEncryptionConfiguration

The encryption settings for an S3 Metadata journal table or inventory table configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "SseAlgorithm" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  SseAlgorithm: String

```

## Properties

`KmsKeyArn`

If server-side encryption with AWS Key Management Service (AWS KMS) keys (SSE-KMS) is specified, you must also
specify the KMS key Amazon Resource Name (ARN). You must specify a customer-managed KMS key
that's located in the same Region as the general purpose bucket that corresponds to the metadata
table configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SseAlgorithm`

The encryption type specified for a metadata table. To specify server-side encryption with
AWS Key Management Service (AWS KMS) keys (SSE-KMS), use the `aws:kms` value. To specify server-side
encryption with Amazon S3 managed keys (SSE-S3), use the `AES256` value.

_Required_: Yes

_Type_: String

_Allowed values_: `aws:kms | AES256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetadataTableConfiguration

Metrics
