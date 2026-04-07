This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::InAppTemplate

Creates a message template that you can use to send in-app messages. A message
template is a set of content and settings that you can define, save, and reuse in
messages for any of your Amazon Pinpoint applications. The In-App channel is unavailable in AWS GovCloud (US).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::InAppTemplate",
  "Properties" : {
      "Content" : [ InAppMessageContent, ... ],
      "CustomConfig" : Json,
      "Layout" : String,
      "Tags" : [ Tag, ... ],
      "TemplateDescription" : String,
      "TemplateName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::InAppTemplate
Properties:
  Content:
    - InAppMessageContent
  CustomConfig: Json
  Layout: String
  Tags:
    - Tag
  TemplateDescription: String
  TemplateName: String

```

## Properties

`Content`

An object that contains information about the content of an in-app message,
including its title and body text, text colors, background colors, images,
buttons, and behaviors.

_Required_: No

_Type_: Array of [InAppMessageContent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-inapptemplate-inappmessagecontent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConfig`

Custom data, in the form of key-value pairs, that is included in an in-app messaging
payload.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Layout`

A string that determines the appearance of the in-app message. You can specify
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

_Allowed values_: `BOTTOM_BANNER | TOP_BANNER | OVERLAYS | MOBILE_FEED | MIDDLE_BANNER | CAROUSEL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateDescription`

An optional description of the in-app template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateName`

The name of the in-app message template.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the message template.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::GCMChannel

BodyConfig
