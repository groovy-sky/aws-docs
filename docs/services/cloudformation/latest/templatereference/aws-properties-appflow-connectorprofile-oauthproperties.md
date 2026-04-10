This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile OAuthProperties

The OAuth properties required for OAuth type authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthCodeUrl" : String,
  "OAuthScopes" : [ String, ... ],
  "TokenUrl" : String
}

```

### YAML

```yaml

  AuthCodeUrl: String
  OAuthScopes:
    - String
  TokenUrl: String

```

## Properties

`AuthCodeUrl`

The authorization code url required to redirect to SAP Login Page to fetch authorization
code for OAuth type authentication.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuthScopes`

The OAuth scopes required for OAuth type authentication.

_Required_: No

_Type_: Array of String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenUrl`

The token url required to fetch access/refresh tokens using authorization code and also
to refresh expired access token using refresh token.

_Required_: No

_Type_: String

_Pattern_: `^(https?)://[-a-zA-Z0-9+&amp;@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&amp;@#/%=~_|]`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuthCredentials

PardotConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
