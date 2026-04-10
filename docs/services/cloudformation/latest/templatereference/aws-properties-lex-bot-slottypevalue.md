This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotTypeValue

Each slot type can have a set of values. Each
`SlotTypeValue` represents a value that the slot type can
take.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SampleValue" : SampleValue,
  "Synonyms" : [ SampleValue, ... ]
}

```

### YAML

```yaml

  SampleValue:
    SampleValue
  Synonyms:
    - SampleValue

```

## Properties

`SampleValue`

The value of the slot type entry.

_Required_: Yes

_Type_: [SampleValue](aws-properties-lex-bot-samplevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Synonyms`

Additional values related to the slot type entry.

_Required_: No

_Type_: Array of [SampleValue](aws-properties-lex-bot-samplevalue.md)

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotType

SlotValue

All content copied from https://docs.aws.amazon.com/.
