This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime NetworkConfiguration

SecurityConfig for the Agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkMode" : String,
  "NetworkModeConfig" : VpcConfig
}

```

### YAML

```yaml

  NetworkMode: String
  NetworkModeConfig:
    VpcConfig

```

## Properties

`NetworkMode`

The network mode for the AgentCore Runtime.

_Required_: Yes

_Type_: String

_Allowed values_: `PUBLIC | VPC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkModeConfig`

The network mode configuration for the AgentCore Runtime.

_Required_: No

_Type_: [VpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-vpcconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LifecycleConfiguration

RequestHeaderConfiguration
