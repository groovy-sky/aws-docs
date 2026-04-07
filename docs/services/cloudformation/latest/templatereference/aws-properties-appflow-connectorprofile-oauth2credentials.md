This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile OAuth2Credentials

The OAuth 2.0 credentials required for OAuth 2.0 authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "OAuthRequest" : ConnectorOAuthRequest,
  "RefreshToken" : String
}

```

### YAML

```yaml

  AccessToken: String
  ClientId: String
  ClientSecret: String
  OAuthRequest:
    ConnectorOAuthRequest
  RefreshToken: String

```

## Properties

`AccessToken`

The access token used to access the connector on your behalf.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The identifier for the desired client.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret used by the OAuth client to authenticate to the authorization
server.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthRequest`

Property description not available.

_Required_: No

_Type_: [ConnectorOAuthRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-connectoroauthrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

The refresh token used to refresh an expired access token.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MarketoConnectorProfileProperties

OAuth2Properties
