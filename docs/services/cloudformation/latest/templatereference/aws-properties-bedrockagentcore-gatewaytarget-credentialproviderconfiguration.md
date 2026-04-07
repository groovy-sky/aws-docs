This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget CredentialProviderConfiguration

The configuration for a credential provider. This structure defines how the gateway authenticates with the target endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialProvider" : CredentialProvider,
  "CredentialProviderType" : String
}

```

### YAML

```yaml

  CredentialProvider:
    CredentialProvider
  CredentialProviderType: String

```

## Properties

`CredentialProvider`

The credential provider. This field contains the specific configuration for the credential provider type.

_Required_: No

_Type_: [CredentialProvider](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-gatewaytarget-credentialprovider.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialProviderType`

The type of credential provider. This field specifies which authentication method the gateway uses.

_Required_: Yes

_Type_: String

_Allowed values_: `GATEWAY_IAM_ROLE | OAUTH | API_KEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CredentialProvider

McpLambdaTargetConfiguration
