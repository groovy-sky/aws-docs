This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm NumericQuestionPropertyValueAutomation

Information about the property value used in automation of a numeric questions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Label" : String
}

```

### YAML

```yaml

  Label: String

```

## Properties

`Label`

The property label of the automation.

_Required_: Yes

_Type_: String

_Allowed values_: `OVERALL_CUSTOMER_SENTIMENT_SCORE | OVERALL_AGENT_SENTIMENT_SCORE | NON_TALK_TIME | NON_TALK_TIME_PERCENTAGE | NUMBER_OF_INTERRUPTIONS | CONTACT_DURATION | AGENT_INTERACTION_DURATION | CUSTOMER_HOLD_TIME | LONGEST_HOLD_DURATION | NUMBER_OF_HOLDS | AGENT_INTERACTION_AND_HOLD_DURATION | CUSTOMER_SENTIMENT_SCORE_WITHOUT_AGENT | CUSTOMER_SENTIMENT_SCORE_WITH_AGENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiSelectQuestionRuleCategoryAutomation

ScoringStrategy

All content copied from https://docs.aws.amazon.com/.
