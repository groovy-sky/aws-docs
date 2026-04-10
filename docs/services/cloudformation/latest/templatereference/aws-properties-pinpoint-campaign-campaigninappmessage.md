This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignInAppMessage

Specifies the appearance of an in-app message, including the message type, the
title and body text, text and background colors, and the configurations of
buttons that appear in the message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : [ InAppMessageContent, ... ],
  "CustomConfig" : Json,
  "Layout" : String
}

```

### YAML

```yaml

  Content:
    - InAppMessageContent
  CustomConfig: Json
  Layout: String

```

## Properties

`Content`

An array that contains configurtion information about the in-app message for
the campaign, including title and body text, text colors, background colors,
image URLs, and button configurations.

_Required_: No

_Type_: Array of [InAppMessageContent](aws-properties-pinpoint-campaign-inappmessagecontent.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConfig`

Custom data, in the form of key-value pairs, that is included in an in-app messaging
payload.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Layout`

A string that describes how the in-app message will appear. You can specify
one of the following:

- `BOTTOM_BANNER` – a message that appears as a banner at the
bottom of the page.

- `TOP_BANNER` – a message that appears as a banner at the
top of the page.

- `OVERLAYS` – a message that covers entire screen.

- `MOBILE_FEED` – a message that appears in a window in front
of the page.

- `MIDDLE_BANNER` – a message that appears as a banner in the
middle of the page.

- `CAROUSEL` – a scrollable layout of up to five unique
messages.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CampaignHook

CampaignSmsMessage

All content copied from https://docs.aws.amazon.com/.
