This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::WebExperience

Creates an Amazon Q Business web experience.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::WebExperience",
  "Properties" : {
      "ApplicationId" : String,
      "BrowserExtensionConfiguration" : BrowserExtensionConfiguration,
      "CustomizationConfiguration" : CustomizationConfiguration,
      "IdentityProviderConfiguration" : IdentityProviderConfiguration,
      "Origins" : [ String, ... ],
      "RoleArn" : String,
      "SamplePromptsControlMode" : String,
      "Subtitle" : String,
      "Tags" : [ Tag, ... ],
      "Title" : String,
      "WelcomeMessage" : String
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::WebExperience
Properties:
  ApplicationId: String
  BrowserExtensionConfiguration:
    BrowserExtensionConfiguration
  CustomizationConfiguration:
    CustomizationConfiguration
  IdentityProviderConfiguration:
    IdentityProviderConfiguration
  Origins:
    - String
  RoleArn: String
  SamplePromptsControlMode: String
  Subtitle: String
  Tags:
    - Tag
  Title: String
  WelcomeMessage: String

```

## Properties

`ApplicationId`

The identifier of the Amazon Q Business web experience.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BrowserExtensionConfiguration`

The container for browser extension configuration for an Amazon Q Business web experience.

_Required_: No

_Type_: [BrowserExtensionConfiguration](aws-properties-qbusiness-webexperience-browserextensionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomizationConfiguration`

Contains the configuration information to customize the logo, font, and color of an Amazon Q Business web experience with individual files for each property or a CSS file for them all.

_Required_: No

_Type_: [CustomizationConfiguration](aws-properties-qbusiness-webexperience-customizationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderConfiguration`

Provides information about the identity provider (IdP) used to authenticate end users
of an Amazon Q Business web experience.

_Required_: No

_Type_: [IdentityProviderConfiguration](aws-properties-qbusiness-webexperience-identityproviderconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Origins`

Sets the website domain origins that are allowed to embed the Amazon Q Business web
experience. The _domain origin_ refers to the base URL for accessing
a website including the protocol ( `http/https`), the domain name, and the
port number (if specified).

###### Note

You must only submit a _base URL_ and not a full path. For
example, `https://docs.aws.amazon.com`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the service role attached to your web
experience.

###### Note

The `roleArn` parameter is required when your Amazon Q Business application
is created with IAM Identity Center. It is not required for SAML-based
applications.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SamplePromptsControlMode`

Determines whether sample prompts are enabled in the web experience for an end
user.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subtitle`

A subtitle to personalize your Amazon Q Business web experience.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize your Amazon Q Business web
experience. You can also use tags to help control access to the web experience. Tag keys
and values can consist of Unicode letters, digits, white space, and any of the following
symbols: \_ . : / = + - @.

_Required_: No

_Type_: Array of [Tag](aws-properties-qbusiness-webexperience-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title for your Amazon Q Business web experience.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WelcomeMessage`

A message in an Amazon Q Business web experience.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application and web experience ID. For
example:

`{"Ref": "ApplicationId|WebExperienceId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`CreatedAt`

The Unix timestamp when the Amazon Q Business application was last updated.

`DefaultEndpoint`

The endpoint URLs for your Amazon Q Business web experience. The URLs are unique and fully
hosted by AWS.

`Status`

The status of your Amazon Q Business web experience.

`UpdatedAt`

The Unix timestamp when your Amazon Q Business web experience was updated.

`WebExperienceArn`

The Amazon Resource Name (ARN) of an Amazon Q Business web experience.

`WebExperienceId`

The identifier of your Amazon Q Business web experience.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BrowserExtensionConfiguration

All content copied from https://docs.aws.amazon.com/.
