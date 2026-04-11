This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection AuthorizationCodeProperties

The set of properties required for the the OAuth2 `AUTHORIZATION_CODE` grant type workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationCode" : String,
  "RedirectUri" : String
}

```

### YAML

```yaml

  AuthorizationCode: String
  RedirectUri: String

```

## Properties

`AuthorizationCode`

An authorization code to be used in the third leg of the `AUTHORIZATION_CODE` grant workflow. This is a single-use code which becomes invalid once exchanged for an access token, thus it is acceptable to have this value as a request parameter.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectUri`

The redirect URI where the user gets redirected to by authorization server when issuing an authorization code. The URI is subsequently used when the authorization code is exchanged for an access token.

_Required_: No

_Type_: String

_Pattern_: `^(https?):\/\/[^\s/$.?#].[^\s]*$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthenticationConfigurationInput

BasicAuthenticationCredentials

All content copied from https://docs.aws.amazon.com/.
