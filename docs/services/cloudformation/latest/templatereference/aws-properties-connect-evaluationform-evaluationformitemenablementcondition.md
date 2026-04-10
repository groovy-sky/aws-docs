This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormItemEnablementCondition

A condition for item enablement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operands" : [ EvaluationFormItemEnablementConditionOperand, ... ],
  "Operator" : String
}

```

### YAML

```yaml

  Operands:
    - EvaluationFormItemEnablementConditionOperand
  Operator: String

```

## Properties

`Operands`

Operands of the enablement condition.

_Required_: Yes

_Type_: Array of [EvaluationFormItemEnablementConditionOperand](aws-properties-connect-evaluationform-evaluationformitemenablementconditionoperand.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

The operator to be used to be applied to operands if more than one provided.

_Required_: No

_Type_: String

_Allowed values_: `OR | AND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormItem

EvaluationFormItemEnablementConditionOperand

All content copied from https://docs.aws.amazon.com/.
