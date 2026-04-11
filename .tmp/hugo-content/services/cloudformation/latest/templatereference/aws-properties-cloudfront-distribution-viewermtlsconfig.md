This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution ViewerMtlsConfig

A viewer mTLS configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String,
  "TrustStoreConfig" : TrustStoreConfig
}

```

### YAML

```yaml

  Mode: String
  TrustStoreConfig:
    TrustStoreConfig

```

## Properties

`Mode`

The viewer mTLS mode.

_Required_: No

_Type_: String

_Allowed values_: `required | optional`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustStoreConfig`

The trust store configuration associated with the viewer mTLS configuration.

_Required_: No

_Type_: [TrustStoreConfig](aws-properties-cloudfront-distribution-truststoreconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ViewerCertificate

VpcOriginConfig

All content copied from https://docs.aws.amazon.com/.
