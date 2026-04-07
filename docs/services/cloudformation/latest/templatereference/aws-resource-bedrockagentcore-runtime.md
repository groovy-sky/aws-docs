This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime

Contains information about an agent runtime. An agent runtime is the execution
environment for a Amazon Bedrock Agent.

AgentCore Runtime is a secure, serverless runtime purpose-built for deploying and
scaling dynamic AI agents and tools using any open-source framework including LangGraph,
CrewAI, and Strands Agents, any protocol, and any model.

For more information about using agent runtime in Amazon Bedrock AgentCore, see [Host agent or\
tools with Amazon Bedrock AgentCore Runtime](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agents-tools-runtime.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::Runtime",
  "Properties" : {
      "AgentRuntimeArtifact" : AgentRuntimeArtifact,
      "AgentRuntimeName" : String,
      "AuthorizerConfiguration" : AuthorizerConfiguration,
      "Description" : String,
      "EnvironmentVariables" : {Key: Value, ...},
      "FilesystemConfigurations" : [ FilesystemConfiguration, ... ],
      "LifecycleConfiguration" : LifecycleConfiguration,
      "NetworkConfiguration" : NetworkConfiguration,
      "ProtocolConfiguration" : String,
      "RequestHeaderConfiguration" : RequestHeaderConfiguration,
      "RoleArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::Runtime
Properties:
  AgentRuntimeArtifact:
    AgentRuntimeArtifact
  AgentRuntimeName: String
  AuthorizerConfiguration:
    AuthorizerConfiguration
  Description: String
  EnvironmentVariables:
    Key: Value
  FilesystemConfigurations:
    - FilesystemConfiguration
  LifecycleConfiguration:
    LifecycleConfiguration
  NetworkConfiguration:
    NetworkConfiguration
  ProtocolConfiguration: String
  RequestHeaderConfiguration:
    RequestHeaderConfiguration
  RoleArn: String
  Tags:
    Key: Value

```

## Properties

`AgentRuntimeArtifact`

The artifact of the AgentCore Runtime.

_Required_: Yes

_Type_: [AgentRuntimeArtifact](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-agentruntimeartifact.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentRuntimeName`

The name of the AgentCore Runtime.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z][a-zA-Z0-9_]{0,47}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizerConfiguration`

The authorizer configuration for the AgentCore Runtime.

_Required_: No

_Type_: [AuthorizerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-authorizerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the AgentCore Runtime.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

Environment variables to set in the AgentCore Runtime environment.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z_][a-zA-Z0-9_]*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilesystemConfigurations`

Property description not available.

_Required_: No

_Type_: Array of [FilesystemConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-filesystemconfiguration.html)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleConfiguration`

The lifecycle configuration for the AgentCore Runtime.

_Required_: No

_Type_: [LifecycleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-lifecycleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkConfiguration`

The network configuration for the AgentCore Runtime.

_Required_: Yes

_Type_: [NetworkConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-networkconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolConfiguration`

The protocol configuration for an agent runtime. This structure defines how the agent runtime communicates with clients.

_Required_: No

_Type_: String

_Allowed values_: `MCP | HTTP | A2A | AGUI`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestHeaderConfiguration`

Configuration for HTTP request headers that will be passed through to the runtime.

_Required_: No

_Type_: [RequestHeaderConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-requestheaderconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role ARN that provides permissions for the AgentCore Runtime.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[^:]+)?:iam::([0-9]{12})?:role/.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the agent.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the agent runtime. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:runtime/MyRuntime-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentRuntimeArn`

The Amazon Resource Name (ARN) of the AgentCore Runtime.

`AgentRuntimeId`

The unique identifier of the AgentCore Runtime.

`AgentRuntimeVersion`

The version of the AgentCore Runtime.

`CreatedAt`

The timestamp when the AgentCore Runtime was created.

`FailureReason`

The reason for failure if the AgentCore Runtime is in a failed state.

`LastUpdatedAt`

The timestamp when the AgentCore Runtime was last updated.

`Status`

The current status of the AgentCore Runtime.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AgentRuntimeArtifact
