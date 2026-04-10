This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot Slot

Specifies the definition of a slot. Amazon Lex elicits slot
values from uses to fulfill the user's intent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "MultipleValuesSetting" : MultipleValuesSetting,
  "Name" : String,
  "ObfuscationSetting" : ObfuscationSetting,
  "SlotTypeName" : String,
  "SubSlotSetting" : SubSlotSetting,
  "ValueElicitationSetting" : SlotValueElicitationSetting
}

```

### YAML

```yaml

  Description: String
  MultipleValuesSetting:
    MultipleValuesSetting
  Name: String
  ObfuscationSetting:
    ObfuscationSetting
  SlotTypeName: String
  SubSlotSetting:
    SubSlotSetting
  ValueElicitationSetting:
    SlotValueElicitationSetting

```

## Properties

`Description`

The description of the slot.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultipleValuesSetting`

Indicates whether a slot can return multiple values.

_Required_: No

_Type_: [MultipleValuesSetting](aws-properties-lex-bot-multiplevaluessetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name given to the slot.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObfuscationSetting`

Determines whether the contents of the slot are obfuscated in
Amazon CloudWatch Logs logs. Use obfuscated slots to protect
information such as personally identifiable information (PII) in
logs.

_Required_: No

_Type_: [ObfuscationSetting](aws-properties-lex-bot-obfuscationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypeName`

The name of the slot type that this slot is based on. The slot type
defines the acceptable values for the slot.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubSlotSetting`

Property description not available.

_Required_: No

_Type_: [SubSlotSetting](aws-properties-lex-bot-subslotsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueElicitationSetting`

Determines the slot resolution strategy that Amazon Lex uses
to return slot type values. The field can be set to one of the
following values:

- ORIGINAL\_VALUE - Returns the value entered by the user, if the
user value is similar to a slot value.

- TOP\_RESOLUTION - If there is a resolution list for the slot,
return the first value in the resolution list as the slot type
value. If there is no resolution list, null is returned.

If you don't specify the `valueSelectionStrategy`,
the default is `ORIGINAL_VALUE`.

_Required_: Yes

_Type_: [SlotValueElicitationSetting](aws-properties-lex-bot-slotvalueelicitationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionAttribute

SlotCaptureSetting

All content copied from https://docs.aws.amazon.com/.
