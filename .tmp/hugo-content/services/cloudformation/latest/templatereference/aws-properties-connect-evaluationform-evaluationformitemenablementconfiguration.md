This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormItemEnablementConfiguration

An item enablement configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Condition" : EvaluationFormItemEnablementCondition,
  "DefaultAction" : String
}

```

### YAML

```yaml

  Action: String
  Condition:
    EvaluationFormItemEnablementCondition
  DefaultAction: String

```

## Properties

`Action`

An enablement action that if condition is satisfied.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLE | ENABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

A condition for item enablement configuration.

_Required_: Yes

_Type_: [EvaluationFormItemEnablementCondition](aws-properties-connect-evaluationform-evaluationformitemenablementcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAction`

An enablement action that if condition is not satisfied.

_Required_: No

_Type_: String

_Allowed values_: `DISABLE | ENABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormItemEnablementConditionOperand

EvaluationFormItemEnablementExpression

All content copied from https://docs.aws.amazon.com/.
