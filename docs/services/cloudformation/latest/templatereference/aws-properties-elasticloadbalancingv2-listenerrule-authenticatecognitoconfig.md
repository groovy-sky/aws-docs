This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule AuthenticateCognitoConfig

Specifies information required when integrating with Amazon Cognito to authenticate
users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationRequestExtraParams" : {Key: Value, ...},
  "OnUnauthenticatedRequest" : String,
  "Scope" : String,
  "SessionCookieName" : String,
  "SessionTimeout" : Integer,
  "UserPoolArn" : String,
  "UserPoolClientId" : String,
  "UserPoolDomain" : String
}

```

### YAML

```yaml

  AuthenticationRequestExtraParams:
    Key: Value
  OnUnauthenticatedRequest: String
  Scope: String
  SessionCookieName: String
  SessionTimeout: Integer
  UserPoolArn: String
  UserPoolClientId: String
  UserPoolDomain: String

```

## Properties

`AuthenticationRequestExtraParams`

The query parameters (up to 10) to include in the redirect request to the authorization
endpoint.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnUnauthenticatedRequest`

The behavior if the user is not authenticated. The following are possible values:

- deny `` \- Return an HTTP 401 Unauthorized error.

- allow `` \- Allow the request to be forwarded to the target.

- authenticate `` \- Redirect the request to the IdP authorization endpoint. This is
the default value.

_Required_: No

_Type_: String

_Allowed values_: `deny | allow | authenticate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The set of user claims to be requested from the IdP. The default is
`openid`.

To verify which scope values your IdP supports and how to separate multiple values, see
the documentation for your IdP.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionCookieName`

The name of the cookie used to maintain session information. The default is
AWSELBAuthSessionCookie.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionTimeout`

The maximum duration of the authentication session, in seconds. The default is 604800
seconds (7 days).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolArn`

The Amazon Resource Name (ARN) of the Amazon Cognito user pool.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolClientId`

The ID of the Amazon Cognito user pool client.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolDomain`

The domain prefix or fully-qualified domain name of the Amazon Cognito user pool.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

AuthenticateOidcConfig

All content copied from https://docs.aws.amazon.com/.
