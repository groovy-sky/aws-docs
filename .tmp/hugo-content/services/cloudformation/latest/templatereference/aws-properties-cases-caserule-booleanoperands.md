This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule BooleanOperands

Boolean operands for a condition. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OperandOne" : OperandOne,
  "OperandTwo" : OperandTwo,
  "Result" : Boolean
}

```

### YAML

```yaml

  OperandOne:
    OperandOne
  OperandTwo:
    OperandTwo
  Result: Boolean

```

## Properties

`OperandOne`

Represents the left hand operand in the condition.

_Required_: Yes

_Type_: [OperandOne](aws-properties-cases-caserule-operandone.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperandTwo`

Represents the right hand operand in the condition.

_Required_: Yes

_Type_: [OperandTwo](aws-properties-cases-caserule-operandtwo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Result`

The value of the outer rule if the condition evaluates to true.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BooleanCondition

CaseRuleDetails

All content copied from https://docs.aws.amazon.com/.
