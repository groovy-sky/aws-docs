This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthLake::FHIRDatastore KmsEncryptionConfig

The customer-managed-key(CMK) used when creating a Data Store. If a customer owned key is not specified, an
Amazon owned key will be used for encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CmkType" : String,
  "KmsKeyId" : String
}

```

### YAML

```yaml

  CmkType: String
  KmsKeyId: String

```

## Properties

`CmkType`

The type of customer-managed-key(CMK) used for encryption. The two types of supported CMKs are customer owned
CMKs and Amazon owned CMKs. For more information on CMK types, see [KmsEncryptionConfig](../../../../reference/healthlake/latest/apireference/api-kmsencryptionconfig.md#HealthLake-Type-KmsEncryptionConfig-CmkType).

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED_KMS_KEY | AWS_OWNED_KMS_KEY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Key Management Service (KMS) encryption key id/alias used to encrypt the data store contents at
rest.

_Required_: No

_Type_: String

_Pattern_: `(arn:aws((-us-gov)|(-iso)|(-iso-b)|(-cn))?:kms:)?([a-z]{2}-[a-z]+(-[a-z]+)?-\d:)?(\d{12}:)?(((key/)?[a-zA-Z0-9-_]+)|(alias/[a-zA-Z0-9:/_-]+))`

_Minimum_: `1`

_Maximum_: `400`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdentityProviderConfiguration

PreloadDataConfig

All content copied from https://docs.aws.amazon.com/.
