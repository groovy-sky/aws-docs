This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgent

Creates an Amazon Q in Connect AI Agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIAgent",
  "Properties" : {
      "AssistantId" : String,
      "Configuration" : AIAgentConfiguration,
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIAgent
Properties:
  AssistantId: String
  Configuration:
    AIAgentConfiguration
  Description: String
  Name: String
  Tags:
    Key: Value
  Type: String

```

## Properties

`AssistantId`

The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN. URLs cannot contain the
ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$|^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}){0,2}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

Configuration for the AI Agent.

_Required_: Yes

_Type_: [AIAgentConfiguration](aws-properties-wisdom-aiagent-aiagentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AI Agent.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s_.,-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the AI Agent.

_Required_: Yes

_Type_: String

_Allowed values_: `MANUAL_SEARCH | ANSWER_RECOMMENDATION | SELF_SERVICE | EMAIL_RESPONSE | EMAIL_OVERVIEW | EMAIL_GENERATIVE_ANSWER | ORCHESTRATION | NOTE_TAKING | CASE_SUMMARIZATION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`AIAgentArn`

The Amazon Resource Name (ARN) of the AI agent.

`AIAgentId`

The identifier of the AI Agent.

`AssistantArn`

The Amazon Resource Name (ARN) of the Amazon Q in Connect assistant.

`ModifiedTimeSeconds`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Connect Wisdom

AIAgentConfiguration

All content copied from https://docs.aws.amazon.com/.
