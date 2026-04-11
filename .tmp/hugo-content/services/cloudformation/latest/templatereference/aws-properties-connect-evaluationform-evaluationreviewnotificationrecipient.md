This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm EvaluationReviewNotificationRecipient

Information about a recipient who should be notified when an evaluation review is requested.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : EvaluationReviewNotificationRecipientValue
}

```

### YAML

```yaml

  Type: String
  Value:
    EvaluationReviewNotificationRecipientValue

```

## Properties

`Type`

The type of notification recipient.

_Required_: Yes

_Type_: String

_Allowed values_: `USER_ID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value associated with the notification recipient type.

_Required_: Yes

_Type_: [EvaluationReviewNotificationRecipientValue](aws-properties-connect-evaluationform-evaluationreviewnotificationrecipientvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationReviewConfiguration

EvaluationReviewNotificationRecipientValue

All content copied from https://docs.aws.amazon.com/.
