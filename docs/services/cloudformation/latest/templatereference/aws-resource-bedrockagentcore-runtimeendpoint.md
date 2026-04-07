This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::RuntimeEndpoint

AgentCore Runtime is a secure, serverless runtime purpose-built for deploying and
scaling dynamic AI agents and tools using any open-source framework including LangGraph,
CrewAI, and Strands Agents, any protocol, and any model.

For more information about using agent runtime endpoints in Amazon Bedrock AgentCore,
see [AgentCore\
Runtime versioning and endpoints](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agent-runtime-versioning.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::RuntimeEndpoint",
  "Properties" : {
      "AgentRuntimeId" : String,
      "AgentRuntimeVersion" : String,
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::RuntimeEndpoint
Properties:
  AgentRuntimeId: String
  AgentRuntimeVersion: String
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`AgentRuntimeId`

The unique identifier of the AgentCore Runtime.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,99}-[a-zA-Z0-9]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AgentRuntimeVersion`

The version of the AgentCore Runtime to use for the endpoint.

_Required_: No

_Type_: String

_Pattern_: `^([1-9][0-9]{0,4})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the AgentCore Runtime endpoint.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AgentCore Runtime endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Minimum_: `1`

_Maximum_: `48`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the AgentCore Runtime endpoint.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the runtime endpoint. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:runtime/MyRuntime-a1b2c3d4e5/runtime-endpoint/MyEndpoint-f6g7h8i9j0`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentRuntimeArn`

The Amazon Resource Name (ARN) of the AgentCore Runtime.

`AgentRuntimeEndpointArn`

The Amazon Resource Name (ARN) of the AgentCore Runtime endpoint.

`CreatedAt`

The timestamp when the AgentCore Runtime endpoint was created.

`FailureReason`

The reason for failure if the AgentCore Runtime endpoint is in a failed state.

`Id`

The unique identifier of the AgentCore Runtime endpoint.

`LastUpdatedAt`

The timestamp when the AgentCore Runtime endpoint was last updated.

`LiveVersion`

The currently deployed version of the AgentCore Runtime on the endpoint.

`Status`

The current status of the AgentCore Runtime endpoint.

`TargetVersion`

The target version of the AgentCore Runtime for the endpoint.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkloadIdentityDetails

AWS::BedrockAgentCore::WorkloadIdentity
