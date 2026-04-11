This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessTrustProvider DeviceOptions

Describes the options for an AWS Verified Access device-identity based trust provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PublicSigningKeyUrl" : String,
  "TenantId" : String
}

```

### YAML

```yaml

  PublicSigningKeyUrl: String
  TenantId: String

```

## Properties

`PublicSigningKeyUrl`

The URL AWS Verified Access will use to verify the authenticity of the device tokens.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TenantId`

The ID of the tenant application with the device-identity provider.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VerifiedAccessTrustProvider

NativeApplicationOidcOptions

All content copied from https://docs.aws.amazon.com/.
