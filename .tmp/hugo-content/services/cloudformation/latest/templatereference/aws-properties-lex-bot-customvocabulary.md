This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot CustomVocabulary

Specifies a custom vocabulary. A custom vocabulary is a list of
words that you expect to be used during a conversation with your
bot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomVocabularyItems" : [ CustomVocabularyItem, ... ]
}

```

### YAML

```yaml

  CustomVocabularyItems:
    - CustomVocabularyItem

```

## Properties

`CustomVocabularyItems`

Specifies a list of words that you expect to be used during a
conversation with your bot.

_Required_: Yes

_Type_: Array of [CustomVocabularyItem](aws-properties-lex-bot-customvocabularyitem.md)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomPayload

CustomVocabularyItem

All content copied from https://docs.aws.amazon.com/.
