This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotValueSelectionSetting

Contains settings used by Amazon Lex to select a slot value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdvancedRecognitionSetting" : AdvancedRecognitionSetting,
  "RegexFilter" : SlotValueRegexFilter,
  "ResolutionStrategy" : String
}

```

### YAML

```yaml

  AdvancedRecognitionSetting:
    AdvancedRecognitionSetting
  RegexFilter:
    SlotValueRegexFilter
  ResolutionStrategy: String

```

## Properties

`AdvancedRecognitionSetting`

Provides settings that enable advanced recognition settings for slot
values. You can use this to enable using slot values as a custom
vocabulary for recognizing user utterances.

_Required_: No

_Type_: [AdvancedRecognitionSetting](aws-properties-lex-bot-advancedrecognitionsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexFilter`

A regular expression used to validate the value of a slot.

_Required_: No

_Type_: [SlotValueRegexFilter](aws-properties-lex-bot-slotvalueregexfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolutionStrategy`

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

_Required_: Yes

_Type_: String

_Allowed values_: `ORIGINAL_VALUE | TOP_RESOLUTION | CONCATENATION`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotValueRegexFilter

Specifications

All content copied from https://docs.aws.amazon.com/.
