This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector AuthorizationCodeGrantDetails

Configuration details for OAuth 2.0 authorization code grant flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationEndpoint" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "TokenEndpoint" : String
}

```

### YAML

```yaml

  AuthorizationEndpoint: String
  ClientId: String
  ClientSecret: String
  TokenEndpoint: String

```

## Properties

`AuthorizationEndpoint`

The authorization endpoint URL for the OAuth flow.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The client ID for the OAuth application.

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret for the OAuth application.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

The token endpoint URL for obtaining access tokens.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationCodeGrantCredentialsDetails

AuthorizationCodeGrantMetadata

All content copied from https://docs.aws.amazon.com/.
