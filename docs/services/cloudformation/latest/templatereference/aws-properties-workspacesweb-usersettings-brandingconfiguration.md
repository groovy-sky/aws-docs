This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings BrandingConfiguration

The branding configuration that customizes the appearance of the web portal for end users. This includes a custom logo, favicon, localized strings, color theme, and optionally a wallpaper and terms of service. If you use an empty string as the wallpaper input during an update, it will remove your previously uploaded wallpaper and use the default wallpaper instead.

###### Note

The `LogoMetadata`, `FaviconMetadata`, and `WallpaperMetadata` properties are read-only and cannot be specified in your template. They are automatically populated by the service after you upload images and can be retrieved using the `Fn::GetAtt` intrinsic function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColorTheme" : String,
  "Favicon" : String,
  "FaviconMetadata" : ImageMetadata,
  "LocalizedStrings" : {Key: Value, ...},
  "Logo" : String,
  "LogoMetadata" : ImageMetadata,
  "TermsOfService" : String,
  "Wallpaper" : String,
  "WallpaperMetadata" : ImageMetadata
}

```

### YAML

```yaml

  ColorTheme: String
  Favicon: String
  FaviconMetadata:
    ImageMetadata
  LocalizedStrings:
    Key: Value
  Logo: String
  LogoMetadata:
    ImageMetadata
  TermsOfService: String
  Wallpaper: String
  WallpaperMetadata:
    ImageMetadata

```

## Properties

`ColorTheme`

The color theme for components on the web portal. Choose `Light` if you upload a dark wallpaper, or `Dark` for a light wallpaper.

_Required_: No

_Type_: String

_Allowed values_: `Light | Dark`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Favicon`

The favicon image for the portal. Provide either a binary image file or an S3 URI pointing to the image file. Maximum 100 KB in JPEG, PNG, or ICO format.

_Required_: No

_Type_: String

_Pattern_: `(^s3://[a-z0-9][a-z0-9\.\-]{1,61}[a-z0-9]/.+$)|(^(?=(.{4})*$)[A-Za-z0-9+/]*={0,2}$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FaviconMetadata`

Read-only. Metadata for the favicon image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.

_Required_: No

_Type_: [ImageMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-imagemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalizedStrings`

A map of localized text strings for different languages, allowing the portal to display content in the user's preferred language.

_Required_: No

_Type_: Object of [LocalizedBrandingStrings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-localizedbrandingstrings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logo`

The logo image for the portal. Provide either a binary image file or an S3 URI pointing to the image file. Maximum 100 KB in JPEG, PNG, or ICO format.

_Required_: No

_Type_: String

_Pattern_: `(^s3://[a-z0-9][a-z0-9\.\-]{1,61}[a-z0-9]/.+$)|(^(?=(.{4})*$)[A-Za-z0-9+/]*={0,2}$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogoMetadata`

Read-only. Metadata for the logo image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.

_Required_: No

_Type_: [ImageMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-imagemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TermsOfService`

The terms of service text in Markdown format that users must accept before accessing the portal.

_Required_: No

_Type_: String

_Maximum_: `153600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Wallpaper`

The wallpaper image for the portal. Provide either a binary image file or an S3 URI pointing to the image file. Maximum 5 MB in JPEG or PNG format. If not provided, a default wallpaper will be used as the background image.

_Required_: No

_Type_: String

_Pattern_: `(^s3://[a-z0-9][a-z0-9\.\-]{1,61}[a-z0-9]/.+$)|(^(?=(.{4})*$)[A-Za-z0-9+/]*={0,2}$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WallpaperMetadata`

Read-only. Metadata for the wallpaper image file, including the MIME type, file extension, and upload timestamp. This property is automatically populated by the service and cannot be specified in your template. It can be retrieved using the `Fn::GetAtt` intrinsic function.

_Required_: No

_Type_: [ImageMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-imagemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::WorkSpacesWeb::UserSettings

CookieSpecification
