This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Chatbot::CustomAction CustomActionAttachment

###### Note

AWS Chatbot is now Amazon Q Developer. [Learn more](https://docs.aws.amazon.com/chatbot/latest/adminguide/service-rename.html)

`Type` attribute values remain unchanged.

Defines when a custom action button should be attached to a notification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ButtonText" : String,
  "Criteria" : [ CustomActionAttachmentCriteria, ... ],
  "NotificationType" : String,
  "Variables" : {Key: Value, ...}
}

```

### YAML

```yaml

  ButtonText: String
  Criteria:
    - CustomActionAttachmentCriteria
  NotificationType: String
  Variables:
    Key: Value

```

## Properties

`ButtonText`

The text of the button that appears on the notification.

_Required_: No

_Type_: String

_Pattern_: `^[\S\s]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criteria`

The criteria for when a button should be shown based on values in the notification.

_Required_: No

_Type_: Array of [CustomActionAttachmentCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-chatbot-customaction-customactionattachmentcriteria.html)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationType`

The type of notification that the custom action should be attached to.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The variables to extract from the notification.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Chatbot::CustomAction

CustomActionAttachmentCriteria
