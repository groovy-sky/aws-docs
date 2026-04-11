This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot CustomVocabularyItem

Specifies an entry in a custom vocabulary.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DisplayAs" : String,
  "Phrase" : String,
  "Weight" : Integer
}

```

### YAML

```yaml

  DisplayAs: String
  Phrase: String
  Weight: Integer

```

## Properties

`DisplayAs`

The DisplayAs value for the custom vocabulary item
from the custom vocabulary list.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phrase`

Specifies 1 - 4 words that should be recognized.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

Specifies the degree to which the phrase recognition is boosted. The
default value is 1.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomVocabulary

DataPrivacy

All content copied from https://docs.aws.amazon.com/.
