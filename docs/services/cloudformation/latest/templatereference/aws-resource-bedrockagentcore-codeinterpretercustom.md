This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::CodeInterpreterCustom

The AgentCore Code Interpreter tool enables agents to securely execute code in
isolated sandbox environments. It offers advanced configuration support and seamless
integration with popular frameworks.

For more information about using the custom code interpreter, see [Execute code and\
analyze data using Amazon Bedrock AgentCore Code Interpreter](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/code-interpreter-tool.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::CodeInterpreterCustom",
  "Properties" : {
      "Description" : String,
      "ExecutionRoleArn" : String,
      "Name" : String,
      "NetworkConfiguration" : CodeInterpreterNetworkConfiguration,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::CodeInterpreterCustom
Properties:
  Description: String
  ExecutionRoleArn: String
  Name: String
  NetworkConfiguration:
    CodeInterpreterNetworkConfiguration
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the code interpreter.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRoleArn`

The IAM role ARN that provides permissions for the code interpreter.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws(?:-cn|-us-gov|-iso(?:-[bef])?)?):iam::[0-9]{12}:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the code interpreter.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkConfiguration`

The network configuration for a code interpreter. This structure defines how the code interpreter connects to the network.

_Required_: Yes

_Type_: [CodeInterpreterNetworkConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-codeinterpretercustom-codeinterpreternetworkconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A map of tag keys and values to assign to the browser. Tags enable you to categorize your resources in different ways, for example, by purpose, owner, or environment.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the custom code interpreter. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:code-interpreter-custom/MyCodeInterpreter-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CodeInterpreterArn`

The code interpreter Amazon Resource Name (ARN).

`CodeInterpreterId`

The unique identifier of the created code interpreter.

`CreatedAt`

The timestamp when the code interpreter was created.

`FailureReason`

The reason for failure if the code interpreter is in a failed state.

`LastUpdatedAt`

The time at which the code interpreter was last updated.

`Status`

The status of the custom code interpreter.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::BedrockAgentCore::BrowserProfile

CodeInterpreterNetworkConfiguration
