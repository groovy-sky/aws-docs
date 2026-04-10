This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup S3StorageConfig

The Amazon Simple Storage (Amazon S3) location and security configuration for
`OfflineStore`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  S3Uri: String

```

## Properties

`KmsKeyId`

The AWS Key Management Service (KMS) key ARN of the key used to encrypt
any objects written into the `OfflineStore` S3 location.

The IAM `roleARN` that is passed as a parameter to
`CreateFeatureGroup` must have below permissions to the
`KmsKeyId`:

- `"kms:GenerateDataKey"`

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The S3 URI, or location in Amazon S3, of `OfflineStore`.

S3 URIs have a format similar to the following:
`s3://example-bucket/prefix/`.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnlineStoreSecurityConfig

Tag

All content copied from https://docs.aws.amazon.com/.
