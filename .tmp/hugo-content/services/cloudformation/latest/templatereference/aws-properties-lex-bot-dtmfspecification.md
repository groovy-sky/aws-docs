This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DTMFSpecification

Specifies the DTMF input specifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeletionCharacter" : String,
  "EndCharacter" : String,
  "EndTimeoutMs" : Integer,
  "MaxLength" : Integer
}

```

### YAML

```yaml

  DeletionCharacter: String
  EndCharacter: String
  EndTimeoutMs: Integer
  MaxLength: Integer

```

## Properties

`DeletionCharacter`

The DTMF character that clears the accumulated DTMF digits and immediately ends
the input.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-D0-9#*]{1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndCharacter`

The DTMF character that immediately ends input. If the user does not press this
character, the input ends after the end timeout.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-D0-9#*]{1}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTimeoutMs`

How long the bot should wait after the last DTMF character input before assuming
that the input has concluded.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxLength`

The maximum number of DTMF digits allowed in an utterance.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DialogState

ElicitationCodeHookInvocationSetting

All content copied from https://docs.aws.amazon.com/.
