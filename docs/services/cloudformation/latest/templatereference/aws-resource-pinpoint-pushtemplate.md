This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::PushTemplate

Creates a message template that you can use in messages that are sent through a push
notification channel. A _message template_ is a set of content and
settings that you can define, save, and reuse in messages for any of your Amazon
Pinpoint applications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::PushTemplate",
  "Properties" : {
      "ADM" : AndroidPushNotificationTemplate,
      "APNS" : APNSPushNotificationTemplate,
      "Baidu" : AndroidPushNotificationTemplate,
      "Default" : DefaultPushNotificationTemplate,
      "DefaultSubstitutions" : String,
      "GCM" : AndroidPushNotificationTemplate,
      "Tags" : [ Tag, ... ],
      "TemplateDescription" : String,
      "TemplateName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::PushTemplate
Properties:
  ADM:
    AndroidPushNotificationTemplate
  APNS:
    APNSPushNotificationTemplate
  Baidu:
    AndroidPushNotificationTemplate
  Default:
    DefaultPushNotificationTemplate
  DefaultSubstitutions: String
  GCM:
    AndroidPushNotificationTemplate
  Tags:
    - Tag
  TemplateDescription: String
  TemplateName: String

```

## Properties

`ADM`

The message template to use for the ADM (Amazon Device Messaging) channel. This
message template overrides the default template for push notification channels
( `Default`).

_Required_: No

_Type_: [AndroidPushNotificationTemplate](aws-properties-pinpoint-pushtemplate-androidpushnotificationtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`APNS`

The message template to use for the APNs (Apple Push Notification service) channel.
This message template overrides the default template for push notification channels
( `Default`).

_Required_: No

_Type_: [APNSPushNotificationTemplate](aws-properties-pinpoint-pushtemplate-apnspushnotificationtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Baidu`

The message template to use for the Baidu (Baidu Cloud Push) channel. This message
template overrides the default template for push notification channels
( `Default`).

_Required_: No

_Type_: [AndroidPushNotificationTemplate](aws-properties-pinpoint-pushtemplate-androidpushnotificationtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Default`

The default message template to use for push notification channels.

_Required_: No

_Type_: [DefaultPushNotificationTemplate](aws-properties-pinpoint-pushtemplate-defaultpushnotificationtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSubstitutions`

A JSON object that specifies the default values to use for message variables
in the message template. This object is a set of key-value pairs. Each key
defines a message variable in the template. The corresponding value defines the
default value for that variable. When you create a message that's based on the
template, you can override these defaults with message-specific and
address-specific variables and values.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GCM`

The message template to use for the GCM channel, which is used to send notifications
through the Firebase Cloud Messaging (FCM), formerly Google Cloud Messaging (GCM),
service. This message template overrides the default template for push notification
channels ( `Default`).

_Required_: No

_Type_: [AndroidPushNotificationTemplate](aws-properties-pinpoint-pushtemplate-androidpushnotificationtemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateDescription`

A custom description of the message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the message template to use for the message. If specified, this
value must match the name of an existing message template.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the name of the message template
( `TemplateName`).

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the message template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OverrideButtonConfiguration

AndroidPushNotificationTemplate

All content copied from https://docs.aws.amazon.com/.
