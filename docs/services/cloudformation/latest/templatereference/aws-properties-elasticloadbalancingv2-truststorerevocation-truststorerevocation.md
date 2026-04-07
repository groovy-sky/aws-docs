This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TrustStoreRevocation TrustStoreRevocation

Information about a revocation file in use by a trust store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NumberOfRevokedEntries" : Integer,
  "RevocationId" : String,
  "RevocationType" : String,
  "TrustStoreArn" : String
}

```

### YAML

```yaml

  NumberOfRevokedEntries: Integer
  RevocationId: String
  RevocationType: String
  TrustStoreArn: String

```

## Properties

`NumberOfRevokedEntries`

The number of revoked certificates.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevocationId`

The revocation ID of the revocation file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevocationType`

The type of revocation file.

_Required_: No

_Type_: String

_Allowed values_: `CRL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreArn`

The Amazon Resource Name (ARN) of the trust store.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RevocationContent

Next
