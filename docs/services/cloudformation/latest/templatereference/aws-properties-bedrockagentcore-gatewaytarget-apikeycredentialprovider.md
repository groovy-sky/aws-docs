This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiKeyCredentialProvider

An API key credential provider for gateway authentication. This structure contains the configuration for authenticating with the target endpoint using an API key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialLocation" : String,
  "CredentialParameterName" : String,
  "CredentialPrefix" : String,
  "ProviderArn" : String
}

```

### YAML

```yaml

  CredentialLocation: String
  CredentialParameterName: String
  CredentialPrefix: String
  ProviderArn: String

```

## Properties

`CredentialLocation`

The location of the API key credential. This field specifies where in the request the API key should be placed.

_Required_: No

_Type_: String

_Allowed values_: `HEADER | QUERY_PARAMETER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialParameterName`

The name of the credential parameter for the API key. This parameter name is used when sending the API key to the target endpoint.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CredentialPrefix`

The prefix for the API key credential. This prefix is added to the API key when sending it to the target endpoint.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderArn`

The Amazon Resource Name (ARN) of the API key credential provider. This ARN identifies the provider in AWS.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:([^:]*):([^:]*):([^:]*):([0-9]{12})?:(.+)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiGatewayToolOverride

ApiSchemaConfiguration

All content copied from https://docs.aws.amazon.com/.
