This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataCatalogEncryptionSettings EncryptionAtRest

Specifies the encryption-at-rest configuration for the Data Catalog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogEncryptionMode" : String,
  "CatalogEncryptionServiceRole" : String,
  "SseAwsKmsKeyId" : String
}

```

### YAML

```yaml

  CatalogEncryptionMode: String
  CatalogEncryptionServiceRole: String
  SseAwsKmsKeyId: String

```

## Properties

`CatalogEncryptionMode`

The encryption-at-rest mode for encrypting Data Catalog data.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | SSE-KMS | SSE-KMS-WITH-SERVICE-ROLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CatalogEncryptionServiceRole`

The role that AWS Glue assumes to encrypt and decrypt the Data Catalog objects on the caller's behalf.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso(-[bef])?))?:iam::[0-9]{12}:role/.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SseAwsKmsKeyId`

The ID of the AWS KMS key to use for encryption at rest.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataCatalogEncryptionSettings

AWS::Glue::DataQualityRuleset

All content copied from https://docs.aws.amazon.com/.
