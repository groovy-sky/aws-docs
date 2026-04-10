This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot PromptAttemptSpecification

Specifies the settings on a prompt attempt.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedInputTypes" : AllowedInputTypes,
  "AllowInterrupt" : Boolean,
  "AudioAndDTMFInputSpecification" : AudioAndDTMFInputSpecification,
  "TextInputSpecification" : TextInputSpecification
}

```

### YAML

```yaml

  AllowedInputTypes:
    AllowedInputTypes
  AllowInterrupt: Boolean
  AudioAndDTMFInputSpecification:
    AudioAndDTMFInputSpecification
  TextInputSpecification:
    TextInputSpecification

```

## Properties

`AllowedInputTypes`

Indicates the allowed input types of the prompt attempt.

_Required_: Yes

_Type_: [AllowedInputTypes](aws-properties-lex-bot-allowedinputtypes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowInterrupt`

Indicates whether the user can interrupt a speech prompt attempt from the bot.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioAndDTMFInputSpecification`

Specifies the settings on audio and DTMF input.

_Required_: No

_Type_: [AudioAndDTMFInputSpecification](aws-properties-lex-bot-audioanddtmfinputspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextInputSpecification`

Specifies the settings on text input.

_Required_: No

_Type_: [TextInputSpecification](aws-properties-lex-bot-textinputspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PostFulfillmentStatusSpecification

PromptSpecification

All content copied from https://docs.aws.amazon.com/.
