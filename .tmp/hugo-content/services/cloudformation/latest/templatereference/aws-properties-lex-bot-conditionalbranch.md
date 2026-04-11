This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ConditionalBranch

A set of actions that Amazon Lex should run if the condition
is matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : Condition,
  "Name" : String,
  "NextStep" : DialogState,
  "Response" : ResponseSpecification
}

```

### YAML

```yaml

  Condition:
    Condition
  Name: String
  NextStep:
    DialogState
  Response:
    ResponseSpecification

```

## Properties

`Condition`

Contains the expression to evaluate. If the condition is true, the
branch's actions are taken.

_Required_: Yes

_Type_: [Condition](aws-properties-lex-bot-condition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the branch.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?)+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NextStep`

The next step in the conversation.

_Required_: Yes

_Type_: [DialogState](aws-properties-lex-bot-dialogstate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Response`

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

_Required_: No

_Type_: [ResponseSpecification](aws-properties-lex-bot-responsespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

ConditionalSpecification

All content copied from https://docs.aws.amazon.com/.
