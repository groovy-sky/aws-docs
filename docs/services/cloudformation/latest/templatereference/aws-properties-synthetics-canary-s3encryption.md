This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary S3Encryption

A structure that contains the configuration
of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3
or later. For more information, see
[Encrypting canary artifacts](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-artifact-encryption.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionMode" : String,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  EncryptionMode: String
  KmsKeyArn: String

```

## Properties

`EncryptionMode`

The encryption method to use
for artifacts created by this canary. Specify `SSE_S3` to use
server-side encryption (SSE) with an Amazon S3-managed
key. Specify `SSE-KMS` to use server-side encryption with a customer-managed AWS KMS key.

If you omit this parameter, an
AWS-managed AWS KMS key is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the customer-managed AWS KMS key to use, if you specify `SSE-KMS`
for `EncryptionMode`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RunConfig

Schedule

All content copied from https://docs.aws.amazon.com/.
