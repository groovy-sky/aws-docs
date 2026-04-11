This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot ResponseSpecification

Specifies a list of message groups that Amazon Lex uses to respond the
user input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowInterrupt" : Boolean,
  "MessageGroupsList" : [ MessageGroup, ... ]
}

```

### YAML

```yaml

  AllowInterrupt: Boolean
  MessageGroupsList:
    - MessageGroup

```

## Properties

`AllowInterrupt`

Indicates whether the user can interrupt a speech response from
Amazon Lex.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageGroupsList`

A collection of responses that Amazon Lex can send to the user. Amazon Lex
chooses the actual response to send at runtime.

_Required_: Yes

_Type_: Array of [MessageGroup](aws-properties-lex-bot-messagegroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Replication

RuntimeSettings

All content copied from https://docs.aws.amazon.com/.
