This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory

Memory allows AI agents to maintain both immediate and long-term knowledge, enabling
context-aware and personalized interactions.

For more information about using Memory in Amazon Bedrock AgentCore, see [Host agent or\
tools with Amazon Bedrock AgentCore Memory](../../../bedrock-agentcore/latest/devguide/memory-getting-started.md).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::Memory",
  "Properties" : {
      "Description" : String,
      "EncryptionKeyArn" : String,
      "EventExpiryDuration" : Integer,
      "MemoryExecutionRoleArn" : String,
      "MemoryStrategies" : [ MemoryStrategy, ... ],
      "Name" : String,
      "StreamDeliveryResources" : StreamDeliveryResources,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::Memory
Properties:
  Description: String
  EncryptionKeyArn: String
  EventExpiryDuration: Integer
  MemoryExecutionRoleArn: String
  MemoryStrategies:
    - MemoryStrategy
  Name: String
  StreamDeliveryResources:
    StreamDeliveryResources
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the memory.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The ARN of the KMS key used to encrypt the memory.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventExpiryDuration`

The number of days after which memory events will expire.

_Required_: Yes

_Type_: Integer

_Minimum_: `3`

_Maximum_: `365`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryExecutionRoleArn`

The ARN of the IAM role that provides permissions for the memory.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryStrategies`

The list of memory strategies associated with this memory.

_Required_: No

_Type_: Array of [MemoryStrategy](aws-properties-bedrockagentcore-memory-memorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the memory.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StreamDeliveryResources`

Property description not available.

_Required_: No

_Type_: [StreamDeliveryResources](aws-properties-bedrockagentcore-memory-streamdeliveryresources.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map of tag keys and values to assign to an AgentCore Memory. Tags enable you to categorize your resources in different ways, for example, by purpose, owner, or environment.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the memory. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:memory/MyMemory-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the memory was created.

`FailureReason`

The reason for failure if the memory is in a failed state.

`MemoryArn`

The Amazon Resource Name (ARN) of the memory.

`MemoryId`

The unique identifier of the memory.

`Status`

The current status of the memory.

`UpdatedAt`

The timestamp when the memory was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ToolSchema

ContentConfiguration

All content copied from https://docs.aws.amazon.com/.
