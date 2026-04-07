This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings

This resource specifies user settings that can be associated with a web portal. Once
associated with a web portal, user settings control how users can transfer data between a
streaming session and the their local devices.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WorkSpacesWeb::UserSettings",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "BrandingConfiguration" : BrandingConfiguration,
      "CookieSynchronizationConfiguration" : CookieSynchronizationConfiguration,
      "CopyAllowed" : String,
      "CustomerManagedKey" : String,
      "DeepLinkAllowed" : String,
      "DisconnectTimeoutInMinutes" : Number,
      "DownloadAllowed" : String,
      "IdleDisconnectTimeoutInMinutes" : Number,
      "PasteAllowed" : String,
      "PrintAllowed" : String,
      "Tags" : [ Tag, ... ],
      "ToolbarConfiguration" : ToolbarConfiguration,
      "UploadAllowed" : String,
      "WebAuthnAllowed" : String
    }
}

```

### YAML

```yaml

Type: AWS::WorkSpacesWeb::UserSettings
Properties:
  AdditionalEncryptionContext:
    Key: Value
  BrandingConfiguration:
    BrandingConfiguration
  CookieSynchronizationConfiguration:
    CookieSynchronizationConfiguration
  CopyAllowed: String
  CustomerManagedKey: String
  DeepLinkAllowed: String
  DisconnectTimeoutInMinutes: Number
  DownloadAllowed: String
  IdleDisconnectTimeoutInMinutes: Number
  PasteAllowed: String
  PrintAllowed: String
  Tags:
    - Tag
  ToolbarConfiguration:
    ToolbarConfiguration
  UploadAllowed: String
  WebAuthnAllowed: String

```

## Properties

`AdditionalEncryptionContext`

The additional encryption context of the user settings.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrandingConfiguration`

The branding configuration that customizes the appearance of the web portal for end users. This includes a custom logo, favicon, localized strings, color theme, and optionally a wallpaper and terms of service.

_Required_: No

_Type_: [BrandingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-brandingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CookieSynchronizationConfiguration`

The configuration that specifies which cookies should be synchronized from the end
user's local browser to the remote browser.

_Required_: No

_Type_: [CookieSynchronizationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-cookiesynchronizationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CopyAllowed`

Specifies whether the user can copy text from the streaming session to the local
device.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerManagedKey`

The customer managed key used to encrypt sensitive information in the user
settings.

_Required_: No

_Type_: String

_Pattern_: `^arn:[\w+=\/,.@-]+:kms:[a-zA-Z0-9\-]*:[a-zA-Z0-9]{1,12}:key\/[a-zA-Z0-9-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeepLinkAllowed`

Specifies whether the user can use deep links that open automatically when connecting to
a session.

_Required_: No

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisconnectTimeoutInMinutes`

The amount of time that a streaming session remains active after users
disconnect.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DownloadAllowed`

Specifies whether the user can download files from the streaming session to the local
device.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdleDisconnectTimeoutInMinutes`

The amount of time that users can be idle (inactive) before they are disconnected from
their streaming session and the disconnect timeout interval begins.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PasteAllowed`

Specifies whether the user can paste text from the local device to the streaming
session.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrintAllowed`

Specifies whether the user can print to the local device.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the user settings resource. A tag is a key-value pair.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ToolbarConfiguration`

The configuration of the toolbar. This allows administrators to select the toolbar type and visual mode, set maximum display resolution for sessions, and choose which items are visible to end users during their sessions. If administrators do not modify these settings, end users retain control over their toolbar preferences.

_Required_: No

_Type_: [ToolbarConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-usersettings-toolbarconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UploadAllowed`

Specifies whether the user can upload files from the local device to the streaming
session.

_Required_: Yes

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAuthnAllowed`

Specifies whether the user can use WebAuthn redirection for passwordless login to websites within the streaming session.

_Required_: No

_Type_: String

_Allowed values_: `Disabled | Enabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function,
`Ref` returns the resource's Amazon Resource Name (ARN).

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`AssociatedPortalArns`

A list of web portal ARNs that this user settings resource is associated with.

`UserSettingsArn`

The ARN of the user settings.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

BrandingConfiguration
