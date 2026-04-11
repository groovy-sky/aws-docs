This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign Message

Specifies the content and settings for a push notification that's sent to
recipients of a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Body" : String,
  "ImageIconUrl" : String,
  "ImageSmallIconUrl" : String,
  "ImageUrl" : String,
  "JsonBody" : String,
  "MediaUrl" : String,
  "RawContent" : String,
  "SilentPush" : Boolean,
  "TimeToLive" : Integer,
  "Title" : String,
  "Url" : String
}

```

### YAML

```yaml

  Action: String
  Body: String
  ImageIconUrl: String
  ImageSmallIconUrl: String
  ImageUrl: String
  JsonBody: String
  MediaUrl: String
  RawContent: String
  SilentPush: Boolean
  TimeToLive: Integer
  Title: String
  Url: String

```

## Properties

`Action`

The action to occur if a recipient taps the push notification. Valid values
are:

- `OPEN_APP` – Your app opens or it becomes the foreground app
if it was sent to the background. This is the default action.

- `DEEP_LINK` – Your app opens and displays a designated user
interface in the app. This setting uses the deep-linking features of iOS
and Android.

- `URL` – The default mobile browser on the recipient's device
opens and loads the web page at a URL that you specify.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Body`

The body of the notification message. The maximum number of characters is
200.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageIconUrl`

The URL of the image to display as the push notification icon, such as the
icon for the app.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageSmallIconUrl`

The URL of the image to display as the small, push notification icon, such as
a small version of the icon for the app.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUrl`

The URL of an image to display in the push notification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JsonBody`

The JSON payload to use for a silent push notification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaUrl`

The URL of the image or video to display in the push notification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RawContent`

The raw, JSON-formatted string to use as the payload for the notification
message. If specified, this value overrides all other content for the
message.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SilentPush`

Specifies whether the notification is a silent push notification, which is a
push notification that doesn't display on a recipient's device. Silent push
notifications can be used for cases such as updating an app's configuration,
displaying messages in an in-app message center, or supporting phone home
functionality.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeToLive`

The number of seconds that the push notification service should keep the
message, if the service is unable to deliver the notification the first time.
This value is converted to an expiration value when it's sent to a push
notification service. If this value is `0`, the service treats the
notification as if it expires immediately and the service doesn't store or try
to deliver the notification again.

This value doesn't apply to messages that are sent through the Amazon Device
Messaging (ADM) service.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title to display above the notification message on a recipient's
device.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL to open in a recipient's default mobile browser, if a recipient taps
the push notification and the value of the `Action` property is
`URL`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Limits

MessageConfiguration

All content copied from https://docs.aws.amazon.com/.
