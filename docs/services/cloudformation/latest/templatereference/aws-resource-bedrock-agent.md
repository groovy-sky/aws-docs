This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Agent

Specifies an agent as a resource in a top-level template. Minimally, you must specify
the following properties:

- AgentName – Specify a name for the agent.

- AgentResourceRoleArn – Specify the Amazon Resource Name (ARN) of the service
role with permissions to invoke API operations on the agent. For more
information, see [Create a service role\
for Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html).

- FoundationModel – Specify the model ID of a foundation model to use when
invoking the agent. For more information, see [Supported regions and\
models for Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-supported.html).

For more information about using agents in Amazon Bedrock, see [Agents for Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/agents.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::Agent",
  "Properties" : {
      "ActionGroups" : [ AgentActionGroup, ... ],
      "AgentCollaboration" : String,
      "AgentCollaborators" : [ AgentCollaborator, ... ],
      "AgentName" : String,
      "AgentResourceRoleArn" : String,
      "AutoPrepare" : Boolean,
      "CustomerEncryptionKeyArn" : String,
      "CustomOrchestration" : CustomOrchestration,
      "Description" : String,
      "FoundationModel" : String,
      "GuardrailConfiguration" : GuardrailConfiguration,
      "IdleSessionTTLInSeconds" : Number,
      "Instruction" : String,
      "KnowledgeBases" : [ AgentKnowledgeBase, ... ],
      "MemoryConfiguration" : MemoryConfiguration,
      "OrchestrationType" : String,
      "PromptOverrideConfiguration" : PromptOverrideConfiguration,
      "SkipResourceInUseCheckOnDelete" : Boolean,
      "Tags" : {Key: Value, ...},
      "TestAliasTags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::Agent
Properties:
  ActionGroups:
    - AgentActionGroup
  AgentCollaboration: String
  AgentCollaborators:
    - AgentCollaborator
  AgentName: String
  AgentResourceRoleArn: String
  AutoPrepare: Boolean
  CustomerEncryptionKeyArn: String
  CustomOrchestration:
    CustomOrchestration
  Description: String
  FoundationModel: String
  GuardrailConfiguration:
    GuardrailConfiguration
  IdleSessionTTLInSeconds: Number
  Instruction: String
  KnowledgeBases:
    - AgentKnowledgeBase
  MemoryConfiguration:
    MemoryConfiguration
  OrchestrationType: String
  PromptOverrideConfiguration:
    PromptOverrideConfiguration
  SkipResourceInUseCheckOnDelete: Boolean
  Tags:
    Key: Value
  TestAliasTags:
    Key: Value

```

## Properties

`ActionGroups`

The action groups that belong to an agent.

_Required_: No

_Type_: Array of [AgentActionGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-agentactiongroup.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentCollaboration`

The agent's collaboration settings.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | SUPERVISOR | SUPERVISOR_ROUTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentCollaborators`

Property description not available.

_Required_: No

_Type_: Array of [AgentCollaborator](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-agentcollaborator.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentName`

The name of the agent.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentResourceRoleArn`

The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the agent.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoPrepare`

Specifies whether to automatically update the `DRAFT` version of the agent
after making changes to the agent. The `DRAFT` version can be continually
iterated upon during internal development. By default, this value is
`false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerEncryptionKeyArn`

The Amazon Resource Name (ARN) of the AWS KMS key that encrypts the agent.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomOrchestration`

Contains custom orchestration configurations for the agent.

_Required_: No

_Type_: [CustomOrchestration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-customorchestration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the agent.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FoundationModel`

The foundation model used for orchestration by the agent.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:(([0-9]{12}:custom-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}(([:][a-z0-9-]{1,63}){0,2})?/[a-z0-9]{12})|(:foundation-model/([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|([0-9]{12}:(inference-profile|application-inference-profile)/[a-zA-Z0-9-:.]+))|(([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|(([0-9a-zA-Z][_-]?)+)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardrailConfiguration`

Details about the guardrail associated with the agent.

_Required_: No

_Type_: [GuardrailConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-guardrailconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdleSessionTTLInSeconds`

The number of seconds for which Amazon Bedrock keeps information about a user's conversation with the agent.

A user interaction remains active for the amount of time specified. If no conversation occurs during this time, the session expires and Amazon Bedrock deletes any data provided before the timeout.

_Required_: No

_Type_: Number

_Minimum_: `60`

_Maximum_: `5400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Instruction`

Instructions that tell the agent what it should do and how it should interact with users.

_Required_: No

_Type_: String

_Minimum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBases`

The knowledge bases associated with the agent.

_Required_: No

_Type_: Array of [AgentKnowledgeBase](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-agentknowledgebase.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryConfiguration`

Contains memory configuration for the agent.

_Required_: No

_Type_: [MemoryConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-memoryconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrchestrationType`

Specifies the orchestration strategy for the agent.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | CUSTOM_ORCHESTRATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PromptOverrideConfiguration`

Contains configurations to override prompt templates in different parts of an agent sequence. For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html).

_Required_: No

_Type_: [PromptOverrideConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-agent-promptoverrideconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkipResourceInUseCheckOnDelete`

Specifies whether to delete the resource even if it's in use. By default, this value
is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that you can assign to a resource as key-value pairs. For more information,
see the following resources:

- [Tag naming\
limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)

- [Tagging\
best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TestAliasTags`

Metadata that you can assign to a resource as key-value pairs. For more information,
see the following resources:

- [Tag naming\
limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)

- [Tagging\
best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the agent base ID.

For example, `{ "Ref": "myAgent" }` could return the value
`"AGENT12345"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentArn`

The Amazon Resource Name (ARN) of the agent.

`AgentId`

The unique identifier of the agent.

`AgentStatus`

The status of the agent and whether it is ready for use. The following statuses are possible:

- CREATING – The agent is being created.

- PREPARING – The agent is being prepared.

- PREPARED – The agent is prepared and ready to be invoked.

- NOT\_PREPARED – The agent has been created but not yet prepared.

- FAILED – The agent API operation failed.

- UPDATING – The agent is being updated.

- DELETING – The agent is being deleted.

`AgentVersion`

The version of the agent.

`CreatedAt`

The time at which the agent was created.

`FailureReasons`

Contains reasons that the agent-related API that you invoked failed.

`PreparedAt`

The time at which the agent was last prepared.

`RecommendedActions`

Contains recommended actions to take for the agent-related API that you invoked to succeed.

`UpdatedAt`

The time at which the agent was last updated.

## Examples

### Create an agent

The following example creates an agent that orchestrates on the Anthropic
Claude v2 foundation model to help customers with IT problems. It contains one
action group and one knowledge base.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: "CFN stack for creating an agent"
Resources:
   ExampleAgentResource:
      Type: AWS::Bedrock::Agent
      Properties:
        AgentName: "Example Agent"
        AgentResourceRoleArn: "arn:aws:iam::123456789012:role/AmazonBedrockExecutionRoleForAgents_user"
        FoundationModel: "anthropic.claude-v2"
        Instruction: "You are an IT agent who solves customer's problems"
        Description: "Description is here"
        IdleSessionTTLInSeconds: 900
        CustomerEncryptionKeyArn: "arn:aws:kms:us-west-2:123456789012:key/1234abcd-ab12-34cd-56ef-abcdefg123456"
        ActionGroups:
          - ActionGroupName: "IT Action"
            Description: "Testing latest IT Management action"
            ActionGroupExecutor:
              Lambda: "arn:aws:lambda:us-west-2:123456789012:function:ItAgentLambda"
            ApiSchema:
              S3:
                S3BucketName: "apischema-s3"
                S3ObjectKey: "ApiSchema.json"
        KnowledgeBases:
          - KnowledgeBaseId: "1234567890"
            Description: "IT Knowledge Base"
            KnowledgeBaseState: ENABLED
```

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Description": "CFN stack for creating an agent",
   "Resources": {
      "ExampleAgentResource": {
         "Type": "AWS::Bedrock::Agent",
         "Properties": {
            "AgentName": "Example Agent",
            "AgentResourceRoleArn": "arn:aws:iam::123456789012:role/AmazonBedrockExecutionRoleForAgents_user",
            "FoundationModel": "anthropic.claude-v2",
            "Instruction": "You are an IT agent who solves customer's problems",
            "Description": "Description is here",
            "IdleSessionTTLInSeconds": 900,
            "CustomerEncryptionKeyArn": "arn:aws:kms:us-west-2:123456789012:key/1234abcd-ab12-34cd-56ef-abcdefg123456",
            "ActionGroups": [
               {
                  "ActionGroupName": "IT Action",
                  "Description": "Testing latest IT Management action",
                  "ActionGroupExecutor": {
                     "Lambda": "arn:aws:lambda:us-west-2:123456789012:function:ItAgentLambda"
                  },
                  "ApiSchema": {
                     "S3": {
                        "S3BucketName": "apischema-s3",
                        "S3ObjectKey": "ApiSchema.json"
                     }
                  }
               }
            ],
            "KnowledgeBases": [
               {
                  "KnowledgeBaseId": "1234567890",
                  "Description": "IT Knowledge Base",
                  "KnowledgeBaseState": "ENABLED"
               }
            ]
         }
      }
   }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Bedrock

ActionGroupExecutor
