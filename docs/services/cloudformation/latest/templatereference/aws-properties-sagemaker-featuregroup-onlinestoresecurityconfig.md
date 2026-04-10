This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup OnlineStoreSecurityConfig

The security configuration for `OnlineStore`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String
}

```

### YAML

```yaml

  KmsKeyId: String

```

## Properties

`KmsKeyId`

The AWS Key Management Service (KMS) key ARN that SageMaker Feature Store
uses to encrypt the Amazon S3 objects at rest using Amazon S3 server-side
encryption.

The caller (either user or IAM role) of `CreateFeatureGroup` must have below
permissions to the `OnlineStore` `KmsKeyId`:

- `"kms:Encrypt"`

- `"kms:Decrypt"`

- `"kms:DescribeKey"`

- `"kms:CreateGrant"`

- `"kms:RetireGrant"`

- `"kms:ReEncryptFrom"`

- `"kms:ReEncryptTo"`

- `"kms:GenerateDataKey"`

- `"kms:ListAliases"`

- `"kms:ListGrants"`

- `"kms:RevokeGrant"`

The caller (either user or IAM role) to all DataPlane operations
( `PutRecord`, `GetRecord`, `DeleteRecord`) must have the
following permissions to the `KmsKeyId`:

- `"kms:Decrypt"`

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnlineStoreConfig

S3StorageConfig

All content copied from https://docs.aws.amazon.com/.
