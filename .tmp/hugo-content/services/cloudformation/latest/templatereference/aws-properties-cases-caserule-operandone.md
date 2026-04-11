This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule OperandOne

Represents the left hand operand in the condition. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String
}

```

### YAML

```yaml

  FieldId: String

```

## Properties

`FieldId`

The field ID that this operand should take the value of.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HiddenCaseRule

OperandTwo

All content copied from https://docs.aws.amazon.com/.
