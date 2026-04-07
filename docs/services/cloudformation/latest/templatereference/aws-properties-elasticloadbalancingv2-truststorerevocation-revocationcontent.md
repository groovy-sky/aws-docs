This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TrustStoreRevocation RevocationContent

Information about a revocation file. You must specify `S3Bucket` and `S3Key`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RevocationType" : String,
  "S3Bucket" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String
}

```

### YAML

```yaml

  RevocationType: String
  S3Bucket: String
  S3Key: String
  S3ObjectVersion: String

```

## Properties

`RevocationType`

The type of revocation file.

_Required_: No

_Type_: String

_Allowed values_: `CRL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Bucket`

The Amazon S3 bucket for the revocation file.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Key`

The Amazon S3 path for the revocation file.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3ObjectVersion`

The Amazon S3 object version of the revocation file.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ElasticLoadBalancingV2::TrustStoreRevocation

TrustStoreRevocation
