This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::IdentitySource IdentitySourceConfiguration

A structure that contains configuration information used when creating or updating a
new identity source.

###### Note

At this time, the only valid member of this structure is a Amazon Cognito user
pool configuration.

You must specify a `userPoolArn`, and optionally, a
`ClientId`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CognitoUserPoolConfiguration" : CognitoUserPoolConfiguration,
  "OpenIdConnectConfiguration" : OpenIdConnectConfiguration
}

```

### YAML

```yaml

  CognitoUserPoolConfiguration:
    CognitoUserPoolConfiguration
  OpenIdConnectConfiguration:
    OpenIdConnectConfiguration

```

## Properties

`CognitoUserPoolConfiguration`

A structure that contains configuration information used when creating or updating an
identity source that represents a connection to an Amazon Cognito user pool used as
an identity provider for Verified Permissions.

_Required_: No

_Type_: [CognitoUserPoolConfiguration](aws-properties-verifiedpermissions-identitysource-cognitouserpoolconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenIdConnectConfiguration`

Property description not available.

_Required_: No

_Type_: [OpenIdConnectConfiguration](aws-properties-verifiedpermissions-identitysource-openidconnectconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CognitoUserPoolConfiguration

OpenIdConnectAccessTokenConfiguration

All content copied from https://docs.aws.amazon.com/.
