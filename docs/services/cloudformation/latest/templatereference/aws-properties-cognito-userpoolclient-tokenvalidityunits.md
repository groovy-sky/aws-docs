This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolClient TokenValidityUnits

The units that validity times are represented in. The default unit for refresh tokens
is days, and the default for ID and access tokens are hours.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "IdToken" : String,
  "RefreshToken" : String
}

```

### YAML

```yaml

  AccessToken: String
  IdToken: String
  RefreshToken: String

```

## Properties

`AccessToken`

A time unit for the value that you set in the `AccessTokenValidity`
parameter. The default `AccessTokenValidity` time unit is `hours`.
`AccessTokenValidity` duration can range from five minutes to one
day.

_Required_: No

_Type_: String

_Allowed values_: `seconds | minutes | hours | days`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdToken`

A time unit for the value that you set in the `IdTokenValidity` parameter.
The default `IdTokenValidity` time unit is `hours`.
`IdTokenValidity` duration can range from five minutes to one day.

_Required_: No

_Type_: String

_Allowed values_: `seconds | minutes | hours | days`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

A time unit for the value that you set in the `RefreshTokenValidity`
parameter. The default `RefreshTokenValidity` time unit is `days`.
`RefreshTokenValidity` duration can range from 60 minutes to 10
years.

_Required_: No

_Type_: String

_Allowed values_: `seconds | minutes | hours | days`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RefreshTokenRotation

AWS::Cognito::UserPoolDomain
