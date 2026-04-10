This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotType

Describes a slot type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompositeSlotTypeSetting" : CompositeSlotTypeSetting,
  "Description" : String,
  "ExternalSourceSetting" : ExternalSourceSetting,
  "Name" : String,
  "ParentSlotTypeSignature" : String,
  "SlotTypeValues" : [ SlotTypeValue, ... ],
  "ValueSelectionSetting" : SlotValueSelectionSetting
}

```

### YAML

```yaml

  CompositeSlotTypeSetting:
    CompositeSlotTypeSetting
  Description: String
  ExternalSourceSetting:
    ExternalSourceSetting
  Name: String
  ParentSlotTypeSignature: String
  SlotTypeValues:
    - SlotTypeValue
  ValueSelectionSetting:
    SlotValueSelectionSetting

```

## Properties

`CompositeSlotTypeSetting`

Property description not available.

_Required_: No

_Type_: [CompositeSlotTypeSetting](aws-properties-lex-bot-compositeslottypesetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the slot type. Use the description to help identify
the slot type in lists.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalSourceSetting`

Sets the type of external information used to create the slot
type.

_Required_: No

_Type_: [ExternalSourceSetting](aws-properties-lex-bot-externalsourcesetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the slot type. A slot type name must be unique withing
the account.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentSlotTypeSignature`

The built-in slot type used as a parent of this slot type. When you
define a parent slot type, the new slot type has the configuration of
the parent lot type.

Only `AMAZON.AlphaNumeric` is supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotTypeValues`

A list of SlotTypeValue objects that defines the values that the
slot type can take. Each value can have a list of synonyms, additional
values that help train the machine learning model about the values that
it resolves for the slot.

_Required_: No

_Type_: Array of [SlotTypeValue](aws-properties-lex-bot-slottypevalue.md)

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueSelectionSetting`

Determines the slot resolution strategy that Amazon Lex uses to
return slot type values. The field can be set to one of the following
values:

- `ORIGINAL_VALUE` \- Returns the value entered by the user, if the
user value is similar to the slot value.

- `TOP_RESOLUTION` \- If there is a resolution list for the slot,
return the first value in the resolution list as the slot type
value. If there is no resolution list, null is returned.

If you don't specify the `valueSelectionStrategy`, the
default is `ORIGINAL_VALUE`.

_Required_: No

_Type_: [SlotValueSelectionSetting](aws-properties-lex-bot-slotvalueselectionsetting.md)

_Allowed values_: `OriginalValue | TopResolution | Concatenation`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotResolutionImprovementSpecification

SlotTypeValue

All content copied from https://docs.aws.amazon.com/.
