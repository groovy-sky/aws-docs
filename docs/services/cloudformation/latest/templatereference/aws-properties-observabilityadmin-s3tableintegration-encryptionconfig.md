This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::S3TableIntegration EncryptionConfig

Defines the encryption configuration for S3 Table integrations, including the encryption
algorithm and KMS key settings.

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

The Amazon Resource Name (ARN) of the KMS key used for encryption when using
customer-managed keys.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z0-9\-]+)?:([a-zA-Z0-9\-]+):([a-z0-9\-]+)?:([0-9]{12})?:(.+)$`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseAlgorithm`

The server-side encryption algorithm used for encrypting data in the S3 Table
integration.

_Required_: Yes

_Type_: String

_Allowed values_: `AES256 | aws:kms`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ObservabilityAdmin::S3TableIntegration

LogSource
