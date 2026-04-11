This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot OutputContext

Describes a session context that is activated when an intent is
fulfilled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "TimeToLiveInSeconds" : Integer,
  "TurnsToLive" : Integer
}

```

### YAML

```yaml

  Name: String
  TimeToLiveInSeconds: Integer
  TurnsToLive: Integer

```

## Properties

`Name`

The name of the output context.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeToLiveInSeconds`

The amount of time, in seconds, that the output context should
remain active. The time is figured from the first time the context is
sent to the user.

_Required_: Yes

_Type_: Integer

_Minimum_: `5`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TurnsToLive`

The number of conversation turns that the output context should
remain active. The number of turns is counted from the first time that
the context is sent to the user.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpensearchConfiguration

PlainTextMessage

All content copied from https://docs.aws.amazon.com/.
