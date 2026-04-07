This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Policy

Specifies a Cedar authorization policy within an Amazon Bedrock AgentCore policy engine. A policy defines the authorization logic that controls what actions your AI agents can perform.

For more information, see [Control agent actions with Amazon Bedrock AgentCore policy engines](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/policy-engine.html).

See the **Properties** section below for descriptions of both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::BedrockAgentCore::Policy",
  "Properties" : {
      "Definition" : PolicyDefinition,
      "Description" : String,
      "Name" : String,
      "PolicyEngineId" : String,
      "ValidationMode" : String
    }
}

```

### YAML

```yaml

Type: AWS::BedrockAgentCore::Policy
Properties:
  Definition:
    PolicyDefinition
  Description: String
  Name: String
  PolicyEngineId: String
  ValidationMode: String

```

## Properties

`Definition`

The Cedar policy statement that defines the access control rules. This contains the actual policy logic used for agent behavior control and access decisions.

_Required_: Yes

_Type_: [PolicyDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-policy-policydefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A human-readable description of the policy's purpose and functionality. Limited to 4,096 characters, this helps administrators understand and manage the policy.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The customer-assigned immutable name for the policy. This human-readable identifier must be unique within the account and cannot exceed 48 characters.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]*$`

_Minimum_: `1`

_Maximum_: `48`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyEngineId`

The identifier of the policy engine that manages this policy. This establishes the policy engine context for policy evaluation and management.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_]*-[a-z0-9_]{10}$`

_Minimum_: `12`

_Maximum_: `59`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationMode`

The validation mode for the policy. Determines how Cedar analyzer validation results are handled.

_Required_: No

_Type_: String

_Allowed values_: `FAIL_ON_ANY_FINDINGS | IGNORE_ALL_FINDINGS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the policy. For example:

`arn:aws:bedrock-agentcore:us-east-1:123456789012:policy-engine/MyPolicyEngine-a1b2c3d4e5/policy/MyPolicy-f6g7h8i9j0`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the policy was created.

`PolicyArn`

The Amazon Resource Name (ARN) of the policy.

`PolicyId`

The unique identifier of the policy.

`Status`

The current status of the policy.

`StatusReasons`

Additional information about the current status of the policy.

`UpdatedAt`

The timestamp when the policy was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CedarPolicy
