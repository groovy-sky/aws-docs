This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp WebAppCustomization

A structure that contains the customization fields for the web app. You can provide a
title, logo, and icon to customize the appearance of your web app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FaviconFile" : String,
  "LogoFile" : String,
  "Title" : String
}

```

### YAML

```yaml

  FaviconFile: String
  LogoFile: String
  Title: String

```

## Properties

`FaviconFile`

Returns an icon file data string (in base64 encoding).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `20960`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogoFile`

Returns a logo file data string (in base64 encoding).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

Returns the page title that you defined for your web app.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Vpc

WebAppUnits
