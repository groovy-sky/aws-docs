This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection OAuth2PropertiesInput

A structure containing properties for OAuth2 in the CreateConnection request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationCodeProperties" : AuthorizationCodeProperties,
  "OAuth2ClientApplication" : OAuth2ClientApplication,
  "OAuth2Credentials" : OAuth2Credentials,
  "OAuth2GrantType" : String,
  "TokenUrl" : String,
  "TokenUrlParametersMap" : Json
}

```

### YAML

```yaml

  AuthorizationCodeProperties:
    AuthorizationCodeProperties
  OAuth2ClientApplication:
    OAuth2ClientApplication
  OAuth2Credentials:
    OAuth2Credentials
  OAuth2GrantType: String
  TokenUrl: String
  TokenUrlParametersMap: Json

```

## Properties

`AuthorizationCodeProperties`

The set of properties required for the the OAuth2 `AUTHORIZATION_CODE` grant type.

_Required_: No

_Type_: [AuthorizationCodeProperties](aws-properties-glue-connection-authorizationcodeproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2ClientApplication`

The client application type in the CreateConnection request. For example, `AWS_MANAGED` or `USER_MANAGED`.

_Required_: No

_Type_: [OAuth2ClientApplication](aws-properties-glue-connection-oauth2clientapplication.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2Credentials`

The credentials used when the authentication type is OAuth2 authentication.

_Required_: No

_Type_: [OAuth2Credentials](aws-properties-glue-connection-oauth2credentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2GrantType`

The OAuth2 grant type in the CreateConnection request. For example, `AUTHORIZATION_CODE`, `JWT_BEARER`, or `CLIENT_CREDENTIALS`.

_Required_: No

_Type_: String

_Allowed values_: `AUTHORIZATION_CODE | CLIENT_CREDENTIALS | JWT_BEARER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrl`

The URL of the provider's authentication server, to exchange an authorization code for an access token.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrlParametersMap`

A map of parameters that are added to the token `GET` request.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2Credentials

PhysicalConnectionRequirements

All content copied from https://docs.aws.amazon.com/.
