This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget CredentialProvider

A credential provider for gateway authentication. This structure contains the configuration for authenticating with the target endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKeyCredentialProvider" : ApiKeyCredentialProvider,
  "IamCredentialProvider" : IamCredentialProvider,
  "OauthCredentialProvider" : OAuthCredentialProvider
}

```

### YAML

```yaml

  ApiKeyCredentialProvider:
    ApiKeyCredentialProvider
  IamCredentialProvider:
    IamCredentialProvider
  OauthCredentialProvider:
    OAuthCredentialProvider

```

## Properties

`ApiKeyCredentialProvider`

The API key credential provider. This provider uses an API key to authenticate with the target endpoint.

_Required_: No

_Type_: [ApiKeyCredentialProvider](aws-properties-bedrockagentcore-gatewaytarget-apikeycredentialprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamCredentialProvider`

The IAM credential provider. This provider uses IAM authentication with SigV4 signing to access the target endpoint.

_Required_: No

_Type_: [IamCredentialProvider](aws-properties-bedrockagentcore-gatewaytarget-iamcredentialprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OauthCredentialProvider`

The OAuth credential provider. This provider uses OAuth authentication to access the target endpoint.

_Required_: No

_Type_: [OAuthCredentialProvider](aws-properties-bedrockagentcore-gatewaytarget-oauthcredentialprovider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiSchemaConfiguration

CredentialProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
