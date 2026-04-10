This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::PushTemplate AndroidPushNotificationTemplate

Specifies channel-specific content and settings for a message template that
can be used in push notifications that are sent through the ADM (Amazon Device
Messaging), Baidu (Baidu Cloud Push), or GCM (Firebase Cloud Messaging, formerly
Google Cloud Messaging) channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Body" : String,
  "ImageIconUrl" : String,
  "ImageUrl" : String,
  "SmallImageIconUrl" : String,
  "Sound" : String,
  "Title" : String,
  "Url" : String
}

```

### YAML

```yaml

  Action: String
  Body: String
  ImageIconUrl: String
  ImageUrl: String
  SmallImageIconUrl: String
  Sound: String
  Title: String
  Url: String

```

## Properties

`Action`

The action to occur if a recipient taps a push notification that's based on
the message template. Valid values are:

- `OPEN_APP` – Your app opens or it becomes the foreground app
if it was sent to the background. This is the default action.

- `DEEP_LINK` – Your app opens and displays a designated user
interface in the app. This action uses the deep-linking features of the
Android platform.

- `URL` – The default mobile browser on the recipient's device
opens and loads the web page at a URL that you specify.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

The message body to use in a push notification that's based on the message
template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageIconUrl`

The URL of the large icon image to display in the content view of a push
notification that's based on the message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUrl`

The URL of an image to display in a push notification that's based on the
message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmallImageIconUrl`

The URL of the small icon image to display in the status bar and the content
view of a push notification that's based on the message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sound`

The sound to play when a recipient receives a push notification that's based
on the message template. You can use the default stream or specify the file name
of a sound resource that's bundled in your app. On an Android platform, the
sound file must reside in `/res/raw/`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title to use in a push notification that's based on the message template.
This title appears above the notification message on a recipient's
device.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL to open in a recipient's default mobile browser, if a recipient taps a
push notification that's based on the message template and the value of the
`Action` property is `URL`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Pinpoint::PushTemplate

APNSPushNotificationTemplate

All content copied from https://docs.aws.amazon.com/.
