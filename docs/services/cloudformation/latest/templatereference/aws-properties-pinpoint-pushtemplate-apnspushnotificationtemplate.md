This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::PushTemplate APNSPushNotificationTemplate

Specifies channel-specific content and settings for a message template that
can be used in push notifications that are sent through the APNs (Apple Push
Notification service) channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Body" : String,
  "MediaUrl" : String,
  "Sound" : String,
  "Title" : String,
  "Url" : String
}

```

### YAML

```yaml

  Action: String
  Body: String
  MediaUrl: String
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
interface in the app. This setting uses the deep-linking features of the
iOS platform.

- `URL` – The default mobile browser on the recipient's device
opens and loads the web page at a URL that you specify.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

The message body to use in push notifications that are based on the message
template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaUrl`

The URL of an image or video to display in push notifications that are based
on the message template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sound`

The key for the sound to play when the recipient receives a push notification
that's based on the message template. The value for this key is the name of a
sound file in your app's main bundle or the `Library/Sounds` folder
in your app's data container. If the sound file can't be found or you specify
`default` for the value, the system plays the default alert
sound.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title to use in push notifications that are based on the message template.
This title appears above the notification message on a recipient's
device.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL to open in the recipient's default mobile browser, if a recipient taps
a push notification that's based on the message template and the value of the
`Action` property is `URL`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AndroidPushNotificationTemplate

DefaultPushNotificationTemplate
