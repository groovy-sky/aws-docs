This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget McpLambdaTargetConfiguration

The Lambda configuration for a Model Context Protocol target. This structure defines how the gateway uses a Lambda function to communicate with the target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaArn" : String,
  "ToolSchema" : ToolSchema
}

```

### YAML

```yaml

  LambdaArn: String
  ToolSchema:
    ToolSchema

```

## Properties

`LambdaArn`

The Amazon Resource Name (ARN) of the Lambda function. This function is invoked by the gateway to communicate with the target.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:lambda:([a-z]{2}(-gov)?-[a-z]+-\d{1}):(\d{12}):function:([a-zA-Z0-9-_.]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?$`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolSchema`

The tool schema for the Lambda function. This schema defines the structure of the tools that the Lambda function provides.

_Required_: Yes

_Type_: [ToolSchema](aws-properties-bedrockagentcore-gatewaytarget-toolschema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IamCredentialProvider

McpServerTargetConfiguration

All content copied from https://docs.aws.amazon.com/.
