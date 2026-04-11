This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataAccessor DataAccessorIdcTrustedTokenIssuerConfiguration

Configuration details for IAM Identity Center Trusted Token Issuer (TTI)
authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdcTrustedTokenIssuerArn" : String
}

```

### YAML

```yaml

  IdcTrustedTokenIssuerArn: String

```

## Properties

`IdcTrustedTokenIssuerArn`

The Amazon Resource Name (ARN) of the IAM Identity Center Trusted Token Issuer that
will be used for authentication.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws:sso::[0-9]{12}:trustedTokenIssuer/(sso)?ins-[a-zA-Z0-9-.]{16}/tti-[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataAccessorAuthenticationDetail

DocumentAttribute

All content copied from https://docs.aws.amazon.com/.
