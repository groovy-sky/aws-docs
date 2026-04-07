This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::AIGuardrailVersion

Creates an Amazon Q in Connect AI Guardrail version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Wisdom::AIGuardrailVersion",
  "Properties" : {
      "AIGuardrailId" : String,
      "AssistantId" : String,
      "ModifiedTimeSeconds" : Number
    }
}

```

### YAML

```yaml

Type: AWS::Wisdom::AIGuardrailVersion
Properties:
  AIGuardrailId: String
  AssistantId: String
  ModifiedTimeSeconds: Number

```

## Properties

`AIGuardrailId`

The ID of the AI guardrail version.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssistantId`

The ID of the AI guardrail version assistant.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModifiedTimeSeconds`

The modified time of the AI guardrail version in seconds.

_Required_: No

_Type_: Number

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`AIGuardrailArn`

The ARN of the AI guardrail version.

`AIGuardrailVersionId`

The ID of the AI guardrail version.

`AssistantArn`

The ARN of the AI guardrail version assistant.

`VersionNumber`

The version number for this AI Guardrail version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GuardrailWordConfig

AWS::Wisdom::AIPrompt
