This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIAgentVersion

Creates and Amazon Q in Connect AI Agent version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIAgentVersion",
  "Properties" : {
      "AIAgentId" : String,
      "AssistantId" : String,
      "ModifiedTimeSeconds" : Number
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIAgentVersion
Properties:
  AIAgentId: String
  AssistantId: String
  ModifiedTimeSeconds: Number

```

## Properties

`AIAgentId`

The identifier of the AI Agent.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssistantId`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModifiedTimeSeconds`

The time the AI Agent version was last modified in seconds.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`AIAgentArn`

Property description not available.

`AIAgentVersionId`

Property description not available.

`AssistantArn`

Property description not available.

`VersionNumber`

The version number for this AI Agent version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UserInteractionConfiguration

AWS::Wisdom::AIGuardrail
