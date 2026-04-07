This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationReviewConfiguration

Configuration settings for evaluation reviews.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EligibilityDays" : Integer,
  "ReviewNotificationRecipients" : [ EvaluationReviewNotificationRecipient, ... ]
}

```

### YAML

```yaml

  EligibilityDays: Integer
  ReviewNotificationRecipients:
    - EvaluationReviewNotificationRecipient

```

## Properties

`EligibilityDays`

Number of days during which a request for review can be submitted for evaluations created from this form.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `90`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReviewNotificationRecipients`

List of recipients who should be notified when a review is requested.

_Required_: Yes

_Type_: Array of [EvaluationReviewNotificationRecipient](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-evaluationreviewnotificationrecipient.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EvaluationFormTextQuestionProperties

EvaluationReviewNotificationRecipient
