This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot SlotDefaultValueSpecification

The default value to use when a user doesn't provide a value for a
slot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValueList" : [ SlotDefaultValue, ... ]
}

```

### YAML

```yaml

  DefaultValueList:
    - SlotDefaultValue

```

## Properties

`DefaultValueList`

A list of default values. Amazon Lex chooses the default value to use in
the order that they are presented in the list.

_Required_: Yes

_Type_: Array of [SlotDefaultValue](aws-properties-lex-bot-slotdefaultvalue.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlotDefaultValue

SlotPriority

All content copied from https://docs.aws.amazon.com/.
