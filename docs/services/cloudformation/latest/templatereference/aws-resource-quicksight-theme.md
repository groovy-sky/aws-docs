This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Theme

Creates a theme.

A _theme_ is set of configuration options for color and layout. Themes apply to analyses and
dashboards. For more information, see [Using Themes in Amazon Quick](https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html) in the
_Amazon Quick User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Theme",
  "Properties" : {
      "AwsAccountId" : String,
      "BaseThemeId" : String,
      "Configuration" : ThemeConfiguration,
      "Name" : String,
      "Permissions" : [ ResourcePermission, ... ],
      "Tags" : [ Tag, ... ],
      "ThemeId" : String,
      "VersionDescription" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Theme
Properties:
  AwsAccountId: String
  BaseThemeId: String
  Configuration:
    ThemeConfiguration
  Name: String
  Permissions:
    - ResourcePermission
  Tags:
    - Tag
  ThemeId: String
  VersionDescription: String

```

## Properties

`AwsAccountId`

The ID of the AWS account where you want to store the new theme.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BaseThemeId`

The ID of the theme that a custom theme will inherit from. All themes inherit from one of
the starting themes defined by Amazon Quick Sight. For a list of the starting themes, use
`ListThemes` or choose **Themes** from
within an analysis.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

The theme configuration, which contains the theme display properties.

_Required_: Yes

_Type_: [ThemeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-themeconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A display name for the theme.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A valid grouping of resource permissions to apply to the new theme.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map of the key-value pairs for the resource tag or tags that you want to add to the
resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-theme-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThemeId`

An ID for the theme that you want to create. The theme ID is unique per AWS Region in
each AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VersionDescription`

A description of the first version of the theme that you're creating. Every time
`UpdateTheme` is called, a new version is created. Each version of the
theme has a description of the version in the `VersionDescription`
field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the theme.

`CreatedTime`

The time the theme was created.

`LastUpdatedTime`

The time the theme was last updated.

`Type`

Theme type.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

YAxisOptions

BorderStyle
