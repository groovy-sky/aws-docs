This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot Intent

Represents an action that the user wants to perform.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockAgentIntentConfiguration" : BedrockAgentIntentConfiguration,
  "Description" : String,
  "DialogCodeHook" : DialogCodeHookSetting,
  "DisplayName" : String,
  "FulfillmentCodeHook" : FulfillmentCodeHookSetting,
  "InitialResponseSetting" : InitialResponseSetting,
  "InputContexts" : [ InputContext, ... ],
  "IntentClosingSetting" : IntentClosingSetting,
  "IntentConfirmationSetting" : IntentConfirmationSetting,
  "KendraConfiguration" : KendraConfiguration,
  "Name" : String,
  "OutputContexts" : [ OutputContext, ... ],
  "ParentIntentSignature" : String,
  "QInConnectIntentConfiguration" : QInConnectIntentConfiguration,
  "QnAIntentConfiguration" : QnAIntentConfiguration,
  "SampleUtterances" : [ SampleUtterance, ... ],
  "SlotPriorities" : [ SlotPriority, ... ],
  "Slots" : [ Slot, ... ]
}

```

### YAML

```yaml

  BedrockAgentIntentConfiguration:
    BedrockAgentIntentConfiguration
  Description: String
  DialogCodeHook:
    DialogCodeHookSetting
  DisplayName: String
  FulfillmentCodeHook:
    FulfillmentCodeHookSetting
  InitialResponseSetting:
    InitialResponseSetting
  InputContexts:
    - InputContext
  IntentClosingSetting:
    IntentClosingSetting
  IntentConfirmationSetting:
    IntentConfirmationSetting
  KendraConfiguration:
    KendraConfiguration
  Name: String
  OutputContexts:
    - OutputContext
  ParentIntentSignature: String
  QInConnectIntentConfiguration:
    QInConnectIntentConfiguration
  QnAIntentConfiguration:
    QnAIntentConfiguration
  SampleUtterances:
    - SampleUtterance
  SlotPriorities:
    - SlotPriority
  Slots:
    - Slot

```

## Properties

`BedrockAgentIntentConfiguration`

Property description not available.

_Required_: No

_Type_: [BedrockAgentIntentConfiguration](aws-properties-lex-bot-bedrockagentintentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the intent. Use the description to help identify
the intent in lists.

_Required_: No

_Type_: String

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DialogCodeHook`

Specifies that Amazon Lex invokes the alias Lambda function for each user input. You can invoke this Lambda
function to personalize user interaction.

_Required_: No

_Type_: [DialogCodeHookSetting](aws-properties-lex-bot-dialogcodehooksetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FulfillmentCodeHook`

Specifies that Amazon Lex invokes the alias Lambda function when the intent is ready for fulfillment. You can invoke
this function to complete the bot's transaction with the user.

_Required_: No

_Type_: [FulfillmentCodeHookSetting](aws-properties-lex-bot-fulfillmentcodehooksetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitialResponseSetting`

Configuration setting for a response sent to the user before Amazon Lex starts eliciting slots.

_Required_: No

_Type_: [InitialResponseSetting](aws-properties-lex-bot-initialresponsesetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputContexts`

A list of contexts that must be active for this intent to be
considered by Amazon Lex.

_Required_: No

_Type_: Array of [InputContext](aws-properties-lex-bot-inputcontext.md)

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntentClosingSetting`

Sets the response that Amazon Lex sends to the user when the
intent is closed.

_Required_: No

_Type_: [IntentClosingSetting](aws-properties-lex-bot-intentclosingsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntentConfirmationSetting`

Provides prompts that Amazon Lex sends to the user to
confirm the completion of an intent. If the user answers "no," the
settings contain a statement that is sent to the user to end the
intent.

_Required_: No

_Type_: [IntentConfirmationSetting](aws-properties-lex-bot-intentconfirmationsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KendraConfiguration`

Provides configuration information for the `AMAZON.KendraSearchIntent`
intent. When you use this intent, Amazon Lex searches the specified Amazon Kendra
index and returns documents from the index that match the user's
utterance.

_Required_: No

_Type_: [KendraConfiguration](aws-properties-lex-bot-kendraconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the intent. Intent names must be unique within the
locale that contains the intent and can't match the name of any
built-in intent.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputContexts`

A list of contexts that the intent activates when it is
fulfilled.

_Required_: No

_Type_: Array of [OutputContext](aws-properties-lex-bot-outputcontext.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentIntentSignature`

A unique identifier for the built-in intent to base this intent
on.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QInConnectIntentConfiguration`

Property description not available.

_Required_: No

_Type_: [QInConnectIntentConfiguration](aws-properties-lex-bot-qinconnectintentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QnAIntentConfiguration`

Property description not available.

_Required_: No

_Type_: [QnAIntentConfiguration](aws-properties-lex-bot-qnaintentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleUtterances`

A list of utterances that a user might say to signal the
intent.

_Required_: No

_Type_: Array of [SampleUtterance](aws-properties-lex-bot-sampleutterance.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlotPriorities`

Indicates the priority for slots. Amazon Lex prompts the
user for slot values in priority order.

_Required_: No

_Type_: Array of [SlotPriority](aws-properties-lex-bot-slotpriority.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slots`

A list of slots that the intent requires for fulfillment.

_Required_: No

_Type_: Array of [Slot](aws-properties-lex-bot-slot.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputContext

IntentClosingSetting

All content copied from https://docs.aws.amazon.com/.
