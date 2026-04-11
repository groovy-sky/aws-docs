This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolClient RefreshTokenRotation

The configuration of your app client for refresh token rotation. When enabled, your
app client issues new ID, access, and refresh tokens when users renew their sessions
with refresh tokens. When disabled, token refresh issues only ID and access
tokens.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Feature" : String,
  "RetryGracePeriodSeconds" : Integer
}

```

### YAML

```yaml

  Feature: String
  RetryGracePeriodSeconds: Integer

```

## Properties

`Feature`

The state of refresh token rotation for the current app client.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryGracePeriodSeconds`

When you request a token refresh with `GetTokensFromRefreshToken`, the
original refresh token that you're rotating out can remain valid for a period of time of
up to 60 seconds. This allows for client-side retries. When
`RetryGracePeriodSeconds` is `0`, the grace period is disabled
and a successful request immediately invalidates the submitted refresh token.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalyticsConfiguration

TokenValidityUnits

All content copied from https://docs.aws.amazon.com/.
