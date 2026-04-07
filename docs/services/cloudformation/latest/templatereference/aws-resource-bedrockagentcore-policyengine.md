This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::PolicyEngine

Specifies a policy engine for Amazon Bedrock AgentCore. A policy engine provides Cedar-based authorization to control what actions your AI agents can perform.

For more information, see [Control agent actions with Amazon Bedrock AgentCore policy engines](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/policy-engine.html).

See the **Properties** section below for descriptions of both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::PolicyEngine",
  "Properties" : {
      "Description" : String,
      "EncryptionKeyArn" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::PolicyEngine
Properties:
  Description: String
  EncryptionKeyArn: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

A human-readable description of the policy engine's purpose and scope. Limited to 4,096 characters, this helps administrators understand the policy engine's role in the overall governance strategy.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The ARN of the KMS key used to encrypt the policy engine data.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The customer-assigned immutable name for the policy engine. This human-readable identifier must be unique within the account and cannot exceed 48 characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]*$`

_Minimum_: `1`

_Maximum_: `48`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the policy engine.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-policyengine-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the policy engine. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:policy-engine/MyPolicyEngine-a1b2c3d4e5`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the policy engine was created.

`PolicyEngineArn`

The Amazon Resource Name (ARN) of the policy engine.

`PolicyEngineId`

The unique identifier of the policy engine.

`Status`

The current status of the policy engine.

`StatusReasons`

Additional information about the current status of the policy engine.

`UpdatedAt`

The timestamp when the policy engine was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PolicyDefinition

Tag
