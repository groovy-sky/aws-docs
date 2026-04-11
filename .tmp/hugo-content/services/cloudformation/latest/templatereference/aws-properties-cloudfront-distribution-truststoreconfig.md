This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution TrustStoreConfig

A trust store configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvertiseTrustStoreCaNames" : Boolean,
  "IgnoreCertificateExpiry" : Boolean,
  "TrustStoreId" : String
}

```

### YAML

```yaml

  AdvertiseTrustStoreCaNames: Boolean
  IgnoreCertificateExpiry: Boolean
  TrustStoreId: String

```

## Properties

`AdvertiseTrustStoreCaNames`

The configuration to use to advertise trust store CA names.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnoreCertificateExpiry`

The configuration to use to ignore certificate expiration.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreId`

The trust store ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TenantConfig

ViewerCertificate

All content copied from https://docs.aws.amazon.com/.
