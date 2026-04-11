This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule SendNotificationAction

Information about the send notification action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "ContentType" : String,
  "DeliveryMethod" : String,
  "Recipient" : NotificationRecipientType,
  "Subject" : String
}

```

### YAML

```yaml

  Content: String
  ContentType: String
  DeliveryMethod: String
  Recipient:
    NotificationRecipientType
  Subject: String

```

## Properties

`Content`

Notification content. Supports variable injection. For more information, see [JSONPath\
reference](../../../connect/latest/adminguide/contact-lens-variable-injection.md) in the _Amazon Connect Administrators_
_Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentType`

Content type format.

_Allowed value_: `PLAIN_TEXT`

_Required_: Yes

_Type_: String

_Allowed values_: `PLAIN_TEXT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryMethod`

Notification delivery method.

_Allowed value_: `EMAIL`

_Required_: Yes

_Type_: String

_Allowed values_: `EMAIL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recipient`

Notification recipient.

_Required_: Yes

_Type_: [NotificationRecipientType](aws-properties-connect-rule-notificationrecipienttype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subject`

The subject of the email if the delivery method is `EMAIL`. Supports
variable injection. For more information, see [JSONPath\
reference](../../../connect/latest/adminguide/contact-lens-variable-injection.md) in the _Amazon Connect Administrators_
_Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleTriggerEventSource

SubmitAutoEvaluationAction

All content copied from https://docs.aws.amazon.com/.
