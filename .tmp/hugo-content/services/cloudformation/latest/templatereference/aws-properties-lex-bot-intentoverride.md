This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot IntentOverride

Override settings to configure the intent state.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Slots" : [ SlotValueOverrideMap, ... ]
}

```

### YAML

```yaml

  Name: String
  Slots:
    - SlotValueOverrideMap

```

## Properties

`Name`

The name of the intent. Only required when you're switching
intents.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slots`

A map of all of the slot value overrides for the intent. The name of
the slot maps to the value of the slot. Slots that are not included in
the map aren't overridden.

_Required_: No

_Type_: Array of [SlotValueOverrideMap](aws-properties-lex-bot-slotvalueoverridemap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntentDisambiguationSettings

KendraConfiguration

All content copied from https://docs.aws.amazon.com/.
