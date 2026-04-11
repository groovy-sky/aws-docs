This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::BotVersion BotVersionLocaleDetails

The version of a bot used for a bot locale.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceBotVersion" : String
}

```

### YAML

```yaml

  SourceBotVersion: String

```

## Properties

`SourceBotVersion`

The version of a bot used for a bot locale.

_Required_: Yes

_Type_: String

_Pattern_: `^(DRAFT|[0-9]+)$`

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lex::BotVersion

BotVersionLocaleSpecification

All content copied from https://docs.aws.amazon.com/.
