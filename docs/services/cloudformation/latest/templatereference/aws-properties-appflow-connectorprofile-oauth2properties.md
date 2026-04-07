This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile OAuth2Properties

The OAuth 2.0 properties required for OAuth 2.0 authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OAuth2GrantType" : String,
  "TokenUrl" : String,
  "TokenUrlCustomProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  OAuth2GrantType: String
  TokenUrl: String
  TokenUrlCustomProperties:
    Key: Value

```

## Properties

`OAuth2GrantType`

The OAuth 2.0 grant type used by connector for OAuth 2.0 authentication.

_Required_: No

_Type_: String

_Allowed values_: `CLIENT_CREDENTIALS | AUTHORIZATION_CODE | JWT_BEARER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrl`

The token URL required for OAuth 2.0 authentication.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrlCustomProperties`

Associates your token URL with a map of properties that you define. Use this parameter to
provide any additional details that the connector requires to authenticate your
request.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\w]{1,128}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OAuth2Credentials

OAuthCredentials
