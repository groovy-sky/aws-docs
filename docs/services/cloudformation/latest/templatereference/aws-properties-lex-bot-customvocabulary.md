---
title: "AWS::Lex::Bot CustomVocabulary"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::Lex::Bot CustomVocabulary
<a name="aws-properties-lex-bot-customvocabulary"></a>

Specifies a custom vocabulary. A custom vocabulary is a list of words that you expect to be used during a conversation with your bot.

## Syntax
<a name="aws-properties-lex-bot-customvocabulary-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-lex-bot-customvocabulary-syntax.json"></a>

```
{
  "[CustomVocabularyItems](#cfn-lex-bot-customvocabulary-customvocabularyitems)" : {{[ CustomVocabularyItem, ... ]}}
}
```

### YAML
<a name="aws-properties-lex-bot-customvocabulary-syntax.yaml"></a>

```
  [CustomVocabularyItems](#cfn-lex-bot-customvocabulary-customvocabularyitems): {{
    - CustomVocabularyItem}}
```

## Properties
<a name="aws-properties-lex-bot-customvocabulary-properties"></a>

`CustomVocabularyItems`  <a name="cfn-lex-bot-customvocabulary-customvocabularyitems"></a>
Specifies a list of words that you expect to be used during a conversation with your bot.
*Required*: Yes
*Type*: Array of [CustomVocabularyItem](aws-properties-lex-bot-customvocabularyitem.md)
*Maximum*: `500`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
