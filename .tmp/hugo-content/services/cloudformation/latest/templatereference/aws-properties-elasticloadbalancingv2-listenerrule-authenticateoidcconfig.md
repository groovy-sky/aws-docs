This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule AuthenticateOidcConfig

Specifies information required using an identity provide (IdP) that is compliant with
OpenID Connect (OIDC) to authenticate users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationRequestExtraParams" : {Key: Value, ...},
  "AuthorizationEndpoint" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "Issuer" : String,
  "OnUnauthenticatedRequest" : String,
  "Scope" : String,
  "SessionCookieName" : String,
  "SessionTimeout" : Integer,
  "TokenEndpoint" : String,
  "UseExistingClientSecret" : Boolean,
  "UserInfoEndpoint" : String
}

```

### YAML

```yaml

  AuthenticationRequestExtraParams:
    Key: Value
  AuthorizationEndpoint: String
  ClientId: String
  ClientSecret: String
  Issuer: String
  OnUnauthenticatedRequest: String
  Scope: String
  SessionCookieName: String
  SessionTimeout: Integer
  TokenEndpoint: String
  UseExistingClientSecret: Boolean
  UserInfoEndpoint: String

```

## Properties

`AuthenticationRequestExtraParams`

The query parameters (up to 10) to include in the redirect request to the authorization
endpoint.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationEndpoint`

The authorization endpoint of the IdP. This must be a full URL, including the HTTPS
protocol, the domain, and the path.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The OAuth 2.0 client identifier.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The OAuth 2.0 client secret. This parameter is required if you are creating a rule. If you
are modifying a rule, you can omit this parameter if you set
`UseExistingClientSecret` to true.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

The OIDC issuer identifier of the IdP. This must be a full URL, including the HTTPS
protocol, the domain, and the path.

_Required_: Yes

_Type_: String

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

`TokenEndpoint`

The token endpoint of the IdP. This must be a full URL, including the HTTPS protocol, the
domain, and the path.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseExistingClientSecret`

Indicates whether to use the existing client secret when modifying a rule. If you are
creating a rule, you can omit this parameter or set it to false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserInfoEndpoint`

The user info endpoint of the IdP. This must be a full URL, including the HTTPS protocol,
the domain, and the path.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticateCognitoConfig

FixedResponseConfig

All content copied from https://docs.aws.amazon.com/.
