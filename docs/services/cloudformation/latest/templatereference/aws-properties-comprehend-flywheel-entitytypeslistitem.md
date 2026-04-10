This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel EntityTypesListItem

An entity type within a labeled training dataset that Amazon Comprehend uses to train a
custom entity recognizer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

An entity type within a labeled training dataset that Amazon Comprehend uses to train a
custom entity recognizer.

Entity types must not contain the following invalid characters: \\n (line break), \\\n
(escaped line break, \\r (carriage return), \\\r (escaped carriage return), \\t (tab), \\\t
(escaped tab), and , (comma).

_Required_: Yes

_Type_: String

_Pattern_: `^(?![^\n\r\t,]*\\n|\\r|\\t)[^\n\r\t,]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EntityRecognitionConfig

Tag

All content copied from https://docs.aws.amazon.com/.
