This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot Specifications

Subslot specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SlotTypeId" : String,
  "SlotTypeName" : String,
  "ValueElicitationSetting" : SubSlotValueElicitationSetting
}

```

### YAML

```yaml

  SlotTypeId: String
  SlotTypeName: String
  ValueElicitationSetting:
    SubSlotValueElicitationSetting

```

## Properties

`SlotTypeId`

The unique identifier assigned to the slot type.

_Required_: No

_Type_: String

_Pattern_: `^((AMAZON\.)[a-zA-Z_]+?|[0-9a-zA-Z]+)$`

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypeName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueElicitationSetting`

Specifies the elicitation setting details for constituent sub slots of a composite slot.

_Required_: Yes

_Type_: [SubSlotValueElicitationSetting](aws-properties-lex-bot-subslotvalueelicitationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValueSelectionSetting

SpeechFoundationModel

All content copied from https://docs.aws.amazon.com/.
