This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection OAuth2Properties

The OAuth2 properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationCodeProperties" : AuthorizationCodeProperties,
  "OAuth2ClientApplication" : OAuth2ClientApplication,
  "OAuth2Credentials" : GlueOAuth2Credentials,
  "OAuth2GrantType" : String,
  "TokenUrl" : String,
  "TokenUrlParametersMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  AuthorizationCodeProperties:
    AuthorizationCodeProperties
  OAuth2ClientApplication:
    OAuth2ClientApplication
  OAuth2Credentials:
    GlueOAuth2Credentials
  OAuth2GrantType: String
  TokenUrl: String
  TokenUrlParametersMap:
    Key: Value

```

## Properties

`AuthorizationCodeProperties`

The authorization code properties of the OAuth2 properties.

_Required_: No

_Type_: [AuthorizationCodeProperties](aws-properties-datazone-connection-authorizationcodeproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2ClientApplication`

The OAuth2 client application of the OAuth2 properties.

_Required_: No

_Type_: [OAuth2ClientApplication](aws-properties-datazone-connection-oauth2clientapplication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2Credentials`

The OAuth2 credentials of the OAuth2 properties.

_Required_: No

_Type_: [GlueOAuth2Credentials](aws-properties-datazone-connection-glueoauth2credentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2GrantType`

The OAuth2 grant type of the OAuth2 properties.

_Required_: No

_Type_: String

_Allowed values_: `AUTHORIZATION_CODE | CLIENT_CREDENTIALS | JWT_BEARER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrl`

The OAuth2 token URL of the OAuth2 properties.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrlParametersMap`

The OAuth2 token URL parameter map of the OAuth2 properties.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2ClientApplication

PhysicalConnectionRequirements

All content copied from https://docs.aws.amazon.com/.
