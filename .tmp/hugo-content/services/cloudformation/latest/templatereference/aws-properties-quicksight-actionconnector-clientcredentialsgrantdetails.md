This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector ClientCredentialsGrantDetails

Configuration details for OAuth2 client credentials grant flow, including client ID, client secret, token endpoint, and optional scopes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ClientSecret" : String,
  "TokenEndpoint" : String
}

```

### YAML

```yaml

  ClientId: String
  ClientSecret: String
  TokenEndpoint: String

```

## Properties

`ClientId`

The client identifier issued to the client during the registration process with the authorization server.

_Required_: Yes

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret issued to the client during the registration process with the authorization server.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenEndpoint`

The authorization server endpoint used to obtain access tokens via the client credentials grant flow.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientCredentialsDetails

ClientCredentialsGrantMetadata

All content copied from https://docs.aws.amazon.com/.
