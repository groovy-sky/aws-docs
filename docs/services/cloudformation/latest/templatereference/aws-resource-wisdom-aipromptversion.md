This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIPromptVersion

Creates an Amazon Q in Connect AI Prompt version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIPromptVersion",
  "Properties" : {
      "AIPromptId" : String,
      "AssistantId" : String,
      "ModifiedTimeSeconds" : Number
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIPromptVersion
Properties:
  AIPromptId: String
  AssistantId: String
  ModifiedTimeSeconds: Number

```

## Properties

`AIPromptId`

The identifier of the Amazon Q in Connect AI prompt.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssistantId`

The identifier of the Amazon Q in Connect assistant. Can be either the ID or the ARN.
URLs cannot contain the ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModifiedTimeSeconds`

The time the AI Prompt version was last modified in seconds.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`AIPromptArn`

The ARN of the AI prompt.

`AIPromptVersionId`

Property description not available.

`AssistantArn`

Property description not available.

`VersionNumber`

The version number for this AI Prompt version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TextFullAIPromptEditTemplateConfiguration

AWS::Wisdom::Assistant
