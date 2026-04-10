This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolClient

The `AWS::Cognito::UserPoolClient` resource specifies an Amazon Cognito
user pool client.

###### Note

If you don't specify a value for a parameter, Amazon Cognito sets it to a default
value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolClient",
  "Properties" : {
      "AccessTokenValidity" : Integer,
      "AllowedOAuthFlows" : [ String, ... ],
      "AllowedOAuthFlowsUserPoolClient" : Boolean,
      "AllowedOAuthScopes" : [ String, ... ],
      "AnalyticsConfiguration" : AnalyticsConfiguration,
      "AuthSessionValidity" : Integer,
      "CallbackURLs" : [ String, ... ],
      "ClientName" : String,
      "DefaultRedirectURI" : String,
      "EnablePropagateAdditionalUserContextData" : Boolean,
      "EnableTokenRevocation" : Boolean,
      "ExplicitAuthFlows" : [ String, ... ],
      "GenerateSecret" : Boolean,
      "IdTokenValidity" : Integer,
      "LogoutURLs" : [ String, ... ],
      "PreventUserExistenceErrors" : String,
      "ReadAttributes" : [ String, ... ],
      "RefreshTokenRotation" : RefreshTokenRotation,
      "RefreshTokenValidity" : Integer,
      "SupportedIdentityProviders" : [ String, ... ],
      "TokenValidityUnits" : TokenValidityUnits,
      "UserPoolId" : String,
      "WriteAttributes" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolClient
Properties:
  AccessTokenValidity: Integer
  AllowedOAuthFlows:
    - String
  AllowedOAuthFlowsUserPoolClient: Boolean
  AllowedOAuthScopes:
    - String
  AnalyticsConfiguration:
    AnalyticsConfiguration
  AuthSessionValidity: Integer
  CallbackURLs:
    - String
  ClientName: String
  DefaultRedirectURI: String
  EnablePropagateAdditionalUserContextData: Boolean
  EnableTokenRevocation: Boolean
  ExplicitAuthFlows:
    - String
  GenerateSecret: Boolean
  IdTokenValidity: Integer
  LogoutURLs:
    - String
  PreventUserExistenceErrors: String
  ReadAttributes:
    - String
  RefreshTokenRotation:
    RefreshTokenRotation
  RefreshTokenValidity: Integer
  SupportedIdentityProviders:
    - String
  TokenValidityUnits:
    TokenValidityUnits
  UserPoolId: String
  WriteAttributes:
    - String

```

## Properties

`AccessTokenValidity`

The access token time limit. After this limit expires, your user can't use
their access token. To specify the time unit for `AccessTokenValidity` as
`seconds`, `minutes`, `hours`, or `days`,
set a `TokenValidityUnits` value in your API request.

For example, when you set `AccessTokenValidity` to `10` and
`TokenValidityUnits` to `hours`, your user can authorize access with
their access token for 10 hours.

The default time unit for `AccessTokenValidity` in an API request is hours.
_Valid range_ is displayed below in seconds.

If you don't specify otherwise in the configuration of your app client, your access
tokens are valid for one hour.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOAuthFlows`

The OAuth grant types that you want your app client to generate for clients in managed
login authentication. To create an app client that generates client credentials grants,
you must add `client_credentials` as the only allowed OAuth flow.

code

Use a code grant flow, which provides an authorization code as the
response. This code can be exchanged for access tokens with the
`/oauth2/token` endpoint.

implicit

Issue the access token, and the ID token when scopes like
`openid` and `profile` are requested, directly to
your user.

client\_credentials

Issue the access token from the `/oauth2/token` endpoint
directly to a non-person user, authorized by a combination of the client ID
and client secret.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOAuthFlowsUserPoolClient`

Set to `true` to use OAuth 2.0 authorization server features in your app client.

This parameter must have a value of `true` before you can configure
the following features in your app client.

- `CallBackURLs`: Callback URLs.

- `LogoutURLs`: Sign-out redirect URLs.

- `AllowedOAuthScopes`: OAuth 2.0 scopes.

- `AllowedOAuthFlows`: Support for authorization code, implicit, and client credentials OAuth 2.0 grants.

To use authorization server features, configure one of these features in the Amazon Cognito console or set
`AllowedOAuthFlowsUserPoolClient` to `true` in a `CreateUserPoolClient` or
`UpdateUserPoolClient` API request. If you don't set a value for
`AllowedOAuthFlowsUserPoolClient` in a request with the AWS CLI or SDKs, it defaults
to `false`. When `false`, only SDK-based API sign-in is permitted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOAuthScopes`

The OAuth, OpenID Connect (OIDC), and custom scopes that you want to permit your app
client to authorize access with. Scopes govern access control to user pool self-service
API operations, user data from the `userInfo` endpoint, and third-party APIs.
Scope values include `phone`, `email`, `openid`, and
`profile`. The `aws.cognito.signin.user.admin` scope
authorizes user self-service operations. Custom scopes with resource servers authorize
access to external APIs.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnalyticsConfiguration`

The user pool analytics configuration for collecting metrics and sending them to your
Amazon Pinpoint campaign.

In AWS Regions where Amazon Pinpoint isn't available, user pools might not have access to
analytics or might be configurable with campaigns in the US East (N. Virginia) Region.
For more information, see [Using Amazon Pinpoint analytics](../../../cognito/latest/developerguide/cognito-user-pools-pinpoint-integration.md).

_Required_: No

_Type_: [AnalyticsConfiguration](aws-properties-cognito-userpoolclient-analyticsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthSessionValidity`

Amazon Cognito creates a session token for each API request in an authentication flow. `AuthSessionValidity` is the duration,
in minutes, of that session token. Your user pool native user must respond to each authentication challenge before the session expires.

_Required_: No

_Type_: Integer

_Minimum_: `3`

_Maximum_: `15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CallbackURLs`

A list of allowed redirect, or callback, URLs for managed login authentication. These
URLs are the paths where you want to send your users' browsers after they complete
authentication with managed login or a third-party IdP. Typically, callback URLs are the
home of an application that uses OAuth or OIDC libraries to process authentication
outcomes.

A redirect URI must meet the following requirements:

- Be an absolute URI.

- Be registered with the authorization server. Amazon Cognito doesn't accept
authorization requests with `redirect_uri` values that aren't in
the list of `CallbackURLs` that you provide in this parameter.

- Not include a fragment component.

See [OAuth 2.0 -\
Redirection Endpoint](https://tools.ietf.org/html/rfc6749).

Amazon Cognito requires HTTPS over HTTP except for callback URLs to
`http://localhost`, `http://127.0.0.1` and
`http://[::1]`. These callback URLs are for testing purposes only. You
can specify custom TCP ports for your callback URLs.

App callback URLs such as `myapp://example` are also supported.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientName`

A friendly name for the app client that you want to create.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRedirectURI`

The default redirect URI. In app clients with one assigned IdP, replaces
`redirect_uri` in authentication requests. Must be in the
`CallbackURLs` list.

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePropagateAdditionalUserContextData`

When `true`, your application can include additional
`UserContextData` in authentication requests. This data includes the IP
address, and contributes to analysis by threat protection features. For more information
about propagation of user context data, see [Adding session data to API requests](../../../cognito/latest/developerguide/cognito-user-pool-settings-adaptive-authentication.md#user-pool-settings-adaptive-authentication-device-fingerprint). If you don’t include this parameter,
you can't send the source IP address to Amazon Cognito threat protection features. You can only
activate `EnablePropagateAdditionalUserContextData` in an app client that has
a client secret.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableTokenRevocation`

Activates or deactivates token revocation.

If you don't include this parameter, token revocation is automatically activated for
the new user pool client.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExplicitAuthFlows`

The [authentication flows](../../../cognito/latest/developerguide/amazon-cognito-user-pools-authentication-flow-methods.md) that you want your user pool client to support. For each app
client in your user pool, you can sign in your users with any combination of one or more flows, including with
a user name and Secure Remote Password (SRP), a user name and password, or a custom authentication process that
you define with Lambda functions.

###### Note

If you don't specify a value for `ExplicitAuthFlows`, your app client supports
`ALLOW_REFRESH_TOKEN_AUTH`, `ALLOW_USER_SRP_AUTH`, and `ALLOW_CUSTOM_AUTH`.

The values for authentication flow options include the following.

- `ALLOW_USER_AUTH`: Enable selection-based sign-in
with `USER_AUTH`. This setting covers username-password,
secure remote password (SRP), passwordless, and passkey authentication.
This authentiation flow can do username-password and SRP authentication
without other `ExplicitAuthFlows` permitting them. For example
users can complete an SRP challenge through `USER_AUTH`
without the flow `USER_SRP_AUTH` being active for the app
client. This flow doesn't include `CUSTOM_AUTH`.

To activate this setting, your user pool must be in the [Essentials tier](../../../cognito/latest/developerguide/feature-plans-features-essentials.md) or higher.

- `ALLOW_ADMIN_USER_PASSWORD_AUTH`: Enable admin based user password
authentication flow `ADMIN_USER_PASSWORD_AUTH`. This setting replaces
the `ADMIN_NO_SRP_AUTH` setting. With this authentication flow, your app
passes a user name and password to Amazon Cognito in the request, instead of using the Secure
Remote Password (SRP) protocol to securely transmit the password.

- `ALLOW_CUSTOM_AUTH`: Enable Lambda trigger based
authentication.

- `ALLOW_USER_PASSWORD_AUTH`: Enable user password-based
authentication. In this flow, Amazon Cognito receives the password in the request instead
of using the SRP protocol to verify passwords.

- `ALLOW_USER_SRP_AUTH`: Enable SRP-based authentication.

- `ALLOW_REFRESH_TOKEN_AUTH`: Enable authflow to refresh
tokens.

In some environments, you will see the values `ADMIN_NO_SRP_AUTH`, `CUSTOM_AUTH_FLOW_ONLY`, or `USER_PASSWORD_AUTH`.
You can't assign these legacy `ExplicitAuthFlows` values to user pool clients at the same time as values that begin with `ALLOW_`,
like `ALLOW_USER_SRP_AUTH`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GenerateSecret`

When `true`, generates a client secret for the app client. Client secrets
are used with server-side and machine-to-machine applications. Client secrets are
automatically generated; you can't specify a secret value. For more information,
see [App client types](../../../cognito/latest/developerguide/user-pool-settings-client-apps.md#user-pool-settings-client-app-client-types).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdTokenValidity`

The ID token time limit. After this limit expires, your user can't use
their ID token. To specify the time unit for `IdTokenValidity` as
`seconds`, `minutes`, `hours`, or `days`,
set a `TokenValidityUnits` value in your API request.

For example, when you set `IdTokenValidity` as `10` and
`TokenValidityUnits` as `hours`, your user can authenticate their
session with their ID token for 10 hours.

The default time unit for `IdTokenValidity` in an API request is hours.
_Valid range_ is displayed below in seconds.

If you don't specify otherwise in the configuration of your app client, your ID
tokens are valid for one hour.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogoutURLs`

A list of allowed logout URLs for managed login authentication. When you pass
`logout_uri` and `client_id` parameters to
`/logout`, Amazon Cognito signs out your user and redirects them to the logout
URL. This parameter describes the URLs that you want to be the permitted targets of
`logout_uri`. A typical use of these URLs is when a user selects "Sign
out" and you redirect them to your public homepage. For more information, see [Logout\
endpoint](../../../cognito/latest/developerguide/logout-endpoint.md).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreventUserExistenceErrors`

Errors and responses that you want Amazon Cognito APIs to return during
authentication, account confirmation, and password recovery when the user doesn't exist
in the user pool. When set to `ENABLED` and the user doesn't exist,
authentication returns an error indicating either the username or password was
incorrect. Account confirmation and password recovery return a response indicating a
code was sent to a simulated destination. When set to `LEGACY`, those APIs
return a `UserNotFoundException` exception if the user doesn't exist in the
user pool.

Valid values include:

- `ENABLED` \- This prevents user existence-related errors.

- `LEGACY` \- This represents the early behavior of Amazon Cognito
where user existence related errors aren't prevented.

Defaults to `LEGACY` when you don't provide a value.

_Required_: No

_Type_: String

_Allowed values_: `LEGACY | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadAttributes`

The list of user attributes that you want your app client to have read access to.
After your user authenticates in your app, their access token authorizes them to read
their own attribute value for any attribute in this list. An example of this kind of
activity is when your user selects a link to view their profile information.

When you don't specify the `ReadAttributes` for your app client, your app
can read the values of `email_verified`, `phone_number_verified`,
and the Standard attributes of your user pool. When your user pool app client has read
access to these default attributes, `ReadAttributes` doesn't return any
information. Amazon Cognito only populates `ReadAttributes` in the API
response if you have specified your own custom set of read attributes.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshTokenRotation`

The configuration of your app client for refresh token rotation. When enabled, your
app client issues new ID, access, and refresh tokens when users renew their sessions
with refresh tokens. When disabled, token refresh issues only ID and access
tokens.

_Required_: No

_Type_: [RefreshTokenRotation](aws-properties-cognito-userpoolclient-refreshtokenrotation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshTokenValidity`

The refresh token time limit. After this limit expires, your user can't use
their refresh token. To specify the time unit for `RefreshTokenValidity` as
`seconds`, `minutes`, `hours`, or `days`,
set a `TokenValidityUnits` value in your API request.

For example, when you set `RefreshTokenValidity` as `10` and
`TokenValidityUnits` as `days`, your user can refresh their session
and retrieve new access and ID tokens for 10 days.

The default time unit for `RefreshTokenValidity` in an API request is days.
You can't set `RefreshTokenValidity` to 0. If you do, Amazon Cognito overrides the
value with the default value of 30 days. _Valid range_ is displayed below
in seconds.

If you don't specify otherwise in the configuration of your app client, your refresh
tokens are valid for 30 days.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `315360000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedIdentityProviders`

A list of provider names for the identity providers (IdPs) that are supported on this
client. The following are supported: `COGNITO`, `Facebook`,
`Google`, `SignInWithApple`, and `LoginWithAmazon`.
You can also specify the names that you configured for the SAML and OIDC IdPs in your
user pool, for example `MySAMLIdP` or `MyOIDCIdP`.

This parameter sets the IdPs that [managed \
login](../../../cognito/latest/developerguide/cognito-user-pools-managed-login.md) will display on the login page for your app client. The removal of
`COGNITO` from this list doesn't prevent authentication operations
for local users with the user pools API in an AWS SDK. The only way to prevent
SDK-based authentication is to block access with a [AWS WAF rule](../../../cognito/latest/developerguide/user-pool-waf.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenValidityUnits`

The units that validity times are represented in. The default unit for refresh tokens
is days, and the default for ID and access tokens are hours.

_Required_: No

_Type_: [TokenValidityUnits](aws-properties-cognito-userpoolclient-tokenvalidityunits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool where you want to create an app client.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WriteAttributes`

The list of user attributes that you want your app client to have write access to.
After your user authenticates in your app, their access token authorizes them to set or
modify their own attribute value for any attribute in this list.

When you don't specify the `WriteAttributes` for your app client, your app
can write the values of the Standard attributes of your user pool. When your user pool
has write access to these default attributes, `WriteAttributes` doesn't
return any information. Amazon Cognito only populates `WriteAttributes` in
the API response if you have specified your own custom set of write attributes.

If your app client allows users to sign in through an IdP, this array must include all
attributes that you have mapped to IdP attributes. Amazon Cognito updates mapped
attributes when users sign in to your application through an IdP. If your app client
does not have write access to a mapped attribute, Amazon Cognito throws an error when it
tries to update the attribute. For more information, see [Specifying IdP Attribute Mappings for Your user pool](../../../cognito/latest/developerguide/cognito-user-pools-specifying-attribute-mapping.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Cognito user pool client ID, such as
`1h57kf5cpq17m0eml12EXAMPLE`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`ClientId`

The ID of the app client, for example `1example23456789`.

## Examples

### Creating a new app client for a user pool

The following example creates an app client with analytics, custom scopes,
OAuth services, and third-party identity providers. An app client with this
configuration is used with managed login.

#### JSON

```json

{
    "UserPoolClient": {
        "Properties": {
            "AccessTokenValidity": 30,
            "AllowedOAuthFlows": [
                "code",
                "implicit"
            ],
            "AllowedOAuthFlowsUserPoolClient": true,
            "AllowedOAuthScopes": [
                "openid",
                "profile",
                "email",
                "myapi.example.com/international.read",
                "myapi.example.com/domestic.read"
            ],
            "AnalyticsConfiguration": {
                "ApplicationArn": "arn:aws:mobiletargeting:us-west-2:123456789012:apps/d70b2ba36a8c4dc5a04a0451aEXAMPLE",
                "UserDataShared": true
            },
            "AuthSessionValidity": 9,
            "CallbackURLs": [
                "https://www.example.com",
                "myapp://example"
            ],
            "ClientName": "my-test-app-client",
            "DefaultRedirectURI": "https://www.example.com",
            "EnablePropagateAdditionalUserContextData": true,
            "EnableTokenRevocation": true,
            "ExplicitAuthFlows": [
                "ALLOW_USER_AUTH",
                "ALLOW_ADMIN_USER_PASSWORD_AUTH",
                "ALLOW_USER_PASSWORD_AUTH",
                "ALLOW_REFRESH_TOKEN_AUTH"
            ],
            "GenerateSecret": true,
            "IdTokenValidity": 30,
            "LogoutURLs": [
                "https://www.example.com/logout",
                "http://localhost:8080"
            ],
            "PreventUserExistenceErrors": "ENABLED",
            "ReadAttributes": [
                "email",
                "phone_number",
                "oidc:profile",
                "custom:department"
            ],
            "RefreshTokenValidity": 10,
            "SupportedIdentityProviders": [
                "Google",
                "LoginWithAmazon",
                "SignInWithApple",
                "Facebook",
                "MYSSO"
            ],
            "TokenValidityUnits": {
                "AccessToken": "minutes",
                "IdToken": "minutes",
                "RefreshToken": "days"
            },
            "UserPoolId": "us-west-2_EXAMPLE",
            "WriteAttributes": [
                "email",
                "oidc:profile",
                "custom:department"
            ]
        },
        "Type": "AWS::Cognito::UserPoolClient"
    }
}
```

#### YAML

```yaml

UserPoolClient:
    Type: AWS::Cognito::UserPoolClient
    Properties:
      AccessTokenValidity: 30
      AllowedOAuthFlows:
        - code
        - implicit
      AllowedOAuthFlowsUserPoolClient: true
      AllowedOAuthScopes:
        - openid
        - profile
        - email
        - myapi.example.com/international.read
        - myapi.example.com/domestic.read
      AnalyticsConfiguration:
        ApplicationArn: arn:aws:mobiletargeting:us-west-2:123456789012:apps/d70b2ba36a8c4dc5a04a0451aEXAMPLE
        UserDataShared: true
      AuthSessionValidity: 9
      CallbackURLs:
        - https://www.example.com
        - myapp://example
      ClientName: my-test-app-client
      DefaultRedirectURI: https://www.example.com
      EnablePropagateAdditionalUserContextData: true
      EnableTokenRevocation: true
      ExplicitAuthFlows:
        - ALLOW_USER_AUTH
        - ALLOW_ADMIN_USER_PASSWORD_AUTH
        - ALLOW_USER_PASSWORD_AUTH
        - ALLOW_REFRESH_TOKEN_AUTH
      GenerateSecret: true
      IdTokenValidity: 30
      LogoutURLs:
        - https://www.example.com/logout
        - http://localhost:8080
      PreventUserExistenceErrors: ENABLED
      ReadAttributes:
        - email
        - phone_number
        - oidc:profile
        - custom:department
      RefreshTokenValidity: 10
      SupportedIdentityProviders:
        - Google
        - LoginWithAmazon
        - SignInWithApple
        - Facebook
        - MYSSO
      TokenValidityUnits:
        AccessToken: minutes
        IdToken: minutes
        RefreshToken: days
      UserPoolId: us-west-2_EXAMPLE
      WriteAttributes:
        - email
        - oidc:profile
        - custom:111
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VerificationMessageTemplate

AnalyticsConfiguration

All content copied from https://docs.aws.amazon.com/.
