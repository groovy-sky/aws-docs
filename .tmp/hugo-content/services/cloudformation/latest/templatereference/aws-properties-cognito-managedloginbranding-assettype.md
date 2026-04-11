This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::ManagedLoginBranding AssetType

An image file from a managed login branding style in a user pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bytes" : String,
  "Category" : String,
  "ColorMode" : String,
  "Extension" : String,
  "ResourceId" : String
}

```

### YAML

```yaml

  Bytes: String
  Category: String
  ColorMode: String
  Extension: String
  ResourceId: String

```

## Properties

`Bytes`

The image file, in Base64-encoded binary.

_Required_: No

_Type_: String

_Maximum_: `1000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Category`

The category that the image corresponds to in your managed login configuration.
Managed login has asset categories for different types of logos, backgrounds, and
icons.

_Required_: Yes

_Type_: String

_Allowed values_: `FAVICON_ICO | FAVICON_SVG | EMAIL_GRAPHIC | SMS_GRAPHIC | AUTH_APP_GRAPHIC | PASSWORD_GRAPHIC | PASSKEY_GRAPHIC | PAGE_HEADER_LOGO | PAGE_HEADER_BACKGROUND | PAGE_FOOTER_LOGO | PAGE_FOOTER_BACKGROUND | PAGE_BACKGROUND | FORM_BACKGROUND | FORM_LOGO | IDP_BUTTON_ICON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColorMode`

The display-mode target of the asset: light, dark, or browser-adaptive. For example,
Amazon Cognito displays a dark-mode image only when the browser or application is in dark mode,
but displays a browser-adaptive file in all contexts.

_Required_: Yes

_Type_: String

_Allowed values_: `LIGHT | DARK | DYNAMIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extension`

The file type of the image file.

_Required_: Yes

_Type_: String

_Allowed values_: `ICO | JPEG | PNG | SVG | WEBP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The ID of the asset.

_Required_: No

_Type_: String

_Pattern_: `^[\w\- ]+$`

_Minimum_: `1`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cognito::ManagedLoginBranding

AWS::Cognito::Terms

All content copied from https://docs.aws.amazon.com/.
