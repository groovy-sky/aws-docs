This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiSchemaConfiguration

Configuration for API schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InlinePayload" : String,
  "S3" : S3Configuration
}

```

### YAML

```yaml

  InlinePayload: String
  S3:
    S3Configuration

```

## Properties

`InlinePayload`

The inline payload containing the API schema definition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The S3 configuration for a gateway. This structure defines how the gateway accesses files in S3.

_Required_: No

_Type_: [S3Configuration](aws-properties-bedrockagentcore-gatewaytarget-s3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKeyCredentialProvider

CredentialProvider

All content copied from https://docs.aws.amazon.com/.
