This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign InAppMessageContent

Specifies the configuration and contents of an in-app message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "BodyConfig" : InAppMessageBodyConfig,
  "HeaderConfig" : InAppMessageHeaderConfig,
  "ImageUrl" : String,
  "PrimaryBtn" : InAppMessageButton,
  "SecondaryBtn" : InAppMessageButton
}

```

### YAML

```yaml

  BackgroundColor: String
  BodyConfig:
    InAppMessageBodyConfig
  HeaderConfig:
    InAppMessageHeaderConfig
  ImageUrl: String
  PrimaryBtn:
    InAppMessageButton
  SecondaryBtn:
    InAppMessageButton

```

## Properties

`BackgroundColor`

The background color for an in-app message banner, expressed as a hex color code (such
as #000000 for black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BodyConfig`

Specifies the configuration of main body text in an in-app message template.

_Required_: No

_Type_: [InAppMessageBodyConfig](aws-properties-pinpoint-campaign-inappmessagebodyconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderConfig`

Specifies the configuration and content of the header or title text of the in-app
message.

_Required_: No

_Type_: [InAppMessageHeaderConfig](aws-properties-pinpoint-campaign-inappmessageheaderconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageUrl`

The URL of the image that appears on an in-app message banner.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryBtn`

An object that contains configuration information about the primary button in
an in-app message.

_Required_: No

_Type_: [InAppMessageButton](aws-properties-pinpoint-campaign-inappmessagebutton.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryBtn`

An object that contains configuration information about the secondary button
in an in-app message.

_Required_: No

_Type_: [InAppMessageButton](aws-properties-pinpoint-campaign-inappmessagebutton.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InAppMessageButton

InAppMessageHeaderConfig

All content copied from https://docs.aws.amazon.com/.
