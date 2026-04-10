This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::BrowserCustom BrowserNetworkConfiguration

The network configuration for a browser. This structure defines how the browser connects to the network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NetworkMode" : String,
  "VpcConfig" : VpcConfig
}

```

### YAML

```yaml

  NetworkMode: String
  VpcConfig:
    VpcConfig

```

## Properties

`NetworkMode`

The network mode for the browser. This field specifies how the browser connects to the network.

_Required_: Yes

_Type_: String

_Allowed values_: `PUBLIC | VPC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfig`

The network mode configuration for the AgentCore Runtime.

_Required_: No

_Type_: [VpcConfig](aws-properties-bedrockagentcore-browsercustom-vpcconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::BrowserCustom

BrowserSigning

All content copied from https://docs.aws.amazon.com/.
