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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-apikeycredentialprovider-tag.html)

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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Bedrock AgentCore

ApiKeySecretArn
