This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection OAuth2Credentials

The credentials used when the authentication type is OAuth2 authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "JwtToken" : String,
  "RefreshToken" : String,
  "UserManagedClientApplicationClientSecret" : String
}

```

### YAML

```yaml

  AccessToken: String
  JwtToken: String
  RefreshToken: String
  UserManagedClientApplicationClientSecret: String

```

## Properties

`AccessToken`

The access token used when the authentication type is OAuth2.

_Required_: No

_Type_: String

_Pattern_: `^[\x20-\x7E]*$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtToken`

The JSON Web Token (JWT) used when the authentication type is OAuth2.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`

_Maximum_: `8000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

The refresh token used when the authentication type is OAuth2.

_Required_: No

_Type_: String

_Pattern_: `^[\x20-\x7E]*$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserManagedClientApplicationClientSecret`

The client application client secret if the client application is user managed.

_Required_: No

_Type_: String

_Pattern_: `^[\x20-\x7E]*$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2ClientApplication

OAuth2PropertiesInput

All content copied from https://docs.aws.amazon.com/.
