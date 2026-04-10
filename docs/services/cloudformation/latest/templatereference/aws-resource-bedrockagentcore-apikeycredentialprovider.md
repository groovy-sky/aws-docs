This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::ApiKeyCredentialProvider

Creates a new API key credential provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::ApiKeyCredentialProvider",
  "Properties" : {
      "ApiKey" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::ApiKeyCredentialProvider
Properties:
  ApiKey: String
  Name: String
  Tags:
    - Tag

```

## Properties

`ApiKey`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the API key credential provider.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-bedrockagentcore-apikeycredentialprovider-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedTime`

The timestamp when the API key credential provider was created.

`CredentialProviderArn`

The Amazon Resource Name (ARN) of the API key credential provider.

`LastUpdatedTime`

The timestamp when the API key credential provider was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Bedrock AgentCore

ApiKeySecretArn

All content copied from https://docs.aws.amazon.com/.
