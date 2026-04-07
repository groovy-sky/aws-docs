This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::PushTemplate DefaultPushNotificationTemplate

Specifies the default settings and content for a message template that can be
used in messages that are sent through a push notification channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Body" : String,
  "Sound" : String,
  "Title" : String,
  "Url" : String
}

```

### YAML

```yaml

  Action: String
  Body: String
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
iOS and Android platforms.

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

`Sound`

The sound to play when a recipient receives a push notification that's based
on the message template. You can use the default stream or specify the file name
of a sound resource that's bundled in your app. On an Android platform, the
sound file must reside in `/res/raw/`.

For an iOS platform, this value is the key for the name of a sound file in
your app's main bundle or the `Library/Sounds` folder in your app's
data container. If the sound file can't be found or you specify
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

The URL to open in a recipient's default mobile browser, if a recipient taps a
push notification that's based on the message template and the value of the
`Action` property is `URL`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

APNSPushNotificationTemplate

AWS::Pinpoint::Segment
