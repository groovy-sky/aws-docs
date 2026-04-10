This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::InAppTemplate InAppMessageContent

Specifies the configuration of an in-app message, including its header, body, buttons,
colors, and images.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BackgroundColor" : String,
  "BodyConfig" : BodyConfig,
  "HeaderConfig" : HeaderConfig,
  "ImageUrl" : String,
  "PrimaryBtn" : ButtonConfig,
  "SecondaryBtn" : ButtonConfig
}

```

### YAML

```yaml

  BackgroundColor: String
  BodyConfig:
    BodyConfig
  HeaderConfig:
    HeaderConfig
  ImageUrl: String
  PrimaryBtn:
    ButtonConfig
  SecondaryBtn:
    ButtonConfig

```

## Properties

`BackgroundColor`

The background color for an in-app message banner, expressed as a hex color code (such
as #000000 for black).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BodyConfig`

An object that contains configuration information about the header or title
text of the in-app message.

_Required_: No

_Type_: [BodyConfig](aws-properties-pinpoint-inapptemplate-bodyconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderConfig`

An object that contains configuration information about the header or title
text of the in-app message.

_Required_: No

_Type_: [HeaderConfig](aws-properties-pinpoint-inapptemplate-headerconfig.md)

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

_Type_: [ButtonConfig](aws-properties-pinpoint-inapptemplate-buttonconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryBtn`

An object that contains configuration information about the secondary button
in an in-app message.

_Required_: No

_Type_: [ButtonConfig](aws-properties-pinpoint-inapptemplate-buttonconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeaderConfig

OverrideButtonConfiguration

All content copied from https://docs.aws.amazon.com/.
