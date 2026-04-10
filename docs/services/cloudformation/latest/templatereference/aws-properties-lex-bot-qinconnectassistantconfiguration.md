This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot QInConnectAssistantConfiguration

The `QInConnectAssistantConfiguration` property type specifies Property description not available. for an [AWS::Lex::Bot](aws-resource-lex-bot.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssistantArn" : String
}

```

### YAML

```yaml

  AssistantArn: String

```

## Properties

`AssistantArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}){0,2}$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PromptSpecification

QInConnectIntentConfiguration

All content copied from https://docs.aws.amazon.com/.
