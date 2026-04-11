This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::CaseRule BooleanCondition

Boolean condition for a rule. In the Amazon Connect admin website, case rules are known as _case field conditions_.
For more
information about case field conditions, see [Add case field conditions to a\
case template](../../../connect/latest/adminguide/case-field-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EqualTo" : BooleanOperands,
  "NotEqualTo" : BooleanOperands
}

```

### YAML

```yaml

  EqualTo:
    BooleanOperands
  NotEqualTo:
    BooleanOperands

```

## Properties

`EqualTo`

Tests that operandOne is equal to operandTwo.

_Required_: No

_Type_: [BooleanOperands](aws-properties-cases-caserule-booleanoperands.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEqualTo`

Tests that operandOne is not equal to operandTwo.

_Required_: No

_Type_: [BooleanOperands](aws-properties-cases-caserule-booleanoperands.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cases::CaseRule

BooleanOperands

All content copied from https://docs.aws.amazon.com/.
