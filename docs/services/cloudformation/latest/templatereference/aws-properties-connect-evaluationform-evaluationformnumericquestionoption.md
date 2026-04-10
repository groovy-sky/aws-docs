This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationFormNumericQuestionOption

Information about the option range used for scoring in numeric questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutomaticFail" : Boolean,
  "AutomaticFailConfiguration" : AutomaticFailConfiguration,
  "MaxValue" : Integer,
  "MinValue" : Integer,
  "Score" : Integer
}

```

### YAML

```yaml

  AutomaticFail: Boolean
  AutomaticFailConfiguration:
    AutomaticFailConfiguration
  MaxValue: Integer
  MinValue: Integer
  Score: Integer

```

## Properties

`AutomaticFail`

The flag to mark the option as automatic fail. If an automatic fail answer is
provided, the overall evaluation gets a score of 0.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticFailConfiguration`

A configuration for automatic fail.

_Required_: No

_Type_: [AutomaticFailConfiguration](aws-properties-connect-evaluationform-automaticfailconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxValue`

The maximum answer value of the range option.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValue`

The minimum answer value of the range option.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Score`

The score assigned to answer values within the range option.

_Minimum_: 0

_Maximum_: 10

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationFormNumericQuestionAutomation

EvaluationFormNumericQuestionProperties

All content copied from https://docs.aws.amazon.com/.
