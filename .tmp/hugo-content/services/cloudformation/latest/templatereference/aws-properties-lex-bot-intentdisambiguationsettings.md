This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot IntentDisambiguationSettings

Configures the Intent Disambiguation feature that helps resolve ambiguous user inputs when multiple intents could match. When enabled, the system presents clarifying questions to users, helping them specify their exact intent for improved conversation accuracy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomDisambiguationMessage" : String,
  "Enabled" : Boolean,
  "MaxDisambiguationIntents" : Integer
}

```

### YAML

```yaml

  CustomDisambiguationMessage: String
  Enabled: Boolean
  MaxDisambiguationIntents: Integer

```

## Properties

`CustomDisambiguationMessage`

Provides a custom message that will be displayed before presenting the disambiguation options to users. This message helps set the context for users and can be customized to match your bot's tone and brand. If not specified, a default message will be used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Determines whether the Intent Disambiguation feature is enabled. When set to `true`, Amazon Lex will present disambiguation options to users when multiple intents could match their input, with the default being `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxDisambiguationIntents`

Specifies the maximum number of intent options (2-5) to present to users when disambiguation is needed. This setting determines how many intent options will be shown to users when the system detects ambiguous input. The default value is 3.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntentConfirmationSetting

IntentOverride

All content copied from https://docs.aws.amazon.com/.
