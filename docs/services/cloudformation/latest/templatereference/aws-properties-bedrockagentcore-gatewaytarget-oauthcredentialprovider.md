This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget OAuthCredentialProvider

An OAuth credential provider for gateway authentication. This structure contains the configuration for authenticating with the target endpoint using OAuth.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomParameters" : {Key: Value, ...},
  "DefaultReturnUrl" : String,
  "GrantType" : String,
  "ProviderArn" : String,
  "Scopes" : [ String, ... ]
}

```

### YAML

```yaml

  CustomParameters:
    Key: Value
  DefaultReturnUrl: String
  GrantType: String
  ProviderArn: String
  Scopes:
    - String

```

## Properties

`CustomParameters`

The custom parameters for the OAuth credential provider. These parameters provide additional configuration for the OAuth authentication process.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultReturnUrl`

The URL where the end user's browser is redirected after obtaining the authorization code. Generally points to the customer's application.

_Required_: No

_Type_: String

_Pattern_: `\w+:(\/?\/?)[^\s]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrantType`

Specifies the kind of credentials to use for authorization:

- `CLIENT_CREDENTIALS` \- Authorization with a client ID and secret.

- `AUTHORIZATION_CODE` \- Authorization with a token that is specific to an individual end user.

_Required_: No

_Type_: String

_Allowed values_: `AUTHORIZATION_CODE | CLIENT_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderArn`

The Amazon Resource Name (ARN) of the OAuth credential provider. This ARN identifies the provider in AWS.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:([^:]*):([^:]*):([^:]*):([0-9]{12})?:(.+)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scopes`

The OAuth scopes for the credential provider. These scopes define the level of access requested from the OAuth provider.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `64 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataConfiguration

S3Configuration

All content copied from https://docs.aws.amazon.com/.
