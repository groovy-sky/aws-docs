This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard

Creates a dashboard from a template. To first create a template, see the `CreateTemplate` API
operation.

A dashboard is an entity in Quick that identifies Quick reports, created from
analyses. You can share Quick dashboards. With the right permissions, you can create scheduled email
reports from them. If you have the correct permissions, you can create a dashboard from a template that exists in a
different AWS account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Dashboard",
  "Properties" : {
      "AwsAccountId" : String,
      "DashboardId" : String,
      "DashboardPublishOptions" : DashboardPublishOptions,
      "Definition" : DashboardVersionDefinition,
      "FolderArns" : [ String, ... ],
      "LinkEntities" : [ String, ... ],
      "LinkSharingConfiguration" : LinkSharingConfiguration,
      "Name" : String,
      "Parameters" : Parameters,
      "Permissions" : [ ResourcePermission, ... ],
      "SourceEntity" : DashboardSourceEntity,
      "Tags" : [ Tag, ... ],
      "ThemeArn" : String,
      "ValidationStrategy" : ValidationStrategy,
      "VersionDescription" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Dashboard
Properties:
  AwsAccountId: String
  DashboardId: String
  DashboardPublishOptions:
    DashboardPublishOptions
  Definition:
    DashboardVersionDefinition
  FolderArns:
    - String
  LinkEntities:
    - String
  LinkSharingConfiguration:
    LinkSharingConfiguration
  Name: String
  Parameters:
    Parameters
  Permissions:
    - ResourcePermission
  SourceEntity:
    DashboardSourceEntity
  Tags:
    - Tag
  ThemeArn: String
  ValidationStrategy:
    ValidationStrategy
  VersionDescription: String

```

## Properties

`AwsAccountId`

The ID of the AWS account where you want to create the
dashboard.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DashboardId`

The ID for the dashboard, also added to the IAM policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DashboardPublishOptions`

Options for publishing the dashboard when you create it:

- `AvailabilityStatus` for `AdHocFilteringOption` \- This
status can be either `ENABLED` or `DISABLED`. When this is
set to `DISABLED`, Amazon Quick Sight disables the left filter pane on
the published dashboard, which can be used for ad hoc (one-time) filtering. This
option is `ENABLED` by default.

- `AvailabilityStatus` for `ExportToCSVOption` \- This
status can be either `ENABLED` or `DISABLED`. The visual
option to export data to .CSV format isn't enabled when this is set to
`DISABLED`. This option is `ENABLED` by default.

- `VisibilityState` for `SheetControlsOption` \- This
visibility state can be either `COLLAPSED` or `EXPANDED`.
This option is `COLLAPSED` by default.

- `AvailabilityStatus` for `QuickSuiteActionsOption` -
This status can be either `ENABLED` or `DISABLED`.
Features related to Actions in Amazon Quick Suite on dashboards are disabled
when this is set to `DISABLED`. This option is `DISABLED`
by default.

- `AvailabilityStatus` for `ExecutiveSummaryOption` \- This
status can be either `ENABLED` or `DISABLED`. The option
to build an executive summary is disabled when this is set to
`DISABLED`. This option is `ENABLED` by
default.

- `AvailabilityStatus` for `DataStoriesSharingOption` -
This status can be either `ENABLED` or `DISABLED`. The
option to share a data story is disabled when this is set to
`DISABLED`. This option is `ENABLED` by
default.

_Required_: No

_Type_: [DashboardPublishOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-dashboardpublishoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Definition`

Property description not available.

_Required_: No

_Type_: [DashboardVersionDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-dashboardversiondefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderArns`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkEntities`

A list of analysis Amazon Resource Names (ARNs) to be linked to the dashboard.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `1024 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkSharingConfiguration`

A structure that contains the link sharing configurations that you want to apply
overrides to.

_Required_: No

_Type_: [LinkSharingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-linksharingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The display name of the dashboard.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the creation of the dashboard, which you want to use to override
the default settings. A dashboard can have any type of parameters, and some parameters
might accept multiple values.

_Required_: No

_Type_: [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-parameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A structure that contains the permissions of the dashboard. You can use this structure
for granting permissions by providing a list of IAM action information
for each principal ARN.

To specify no permissions, omit the permissions list.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceEntity`

The entity that you are using as a source when you create the dashboard. In `SourceEntity`, you
specify the type of object that you want to use. You can only create a dashboard from a template, so you use a
`SourceTemplate` entity. If you need to create a dashboard from an analysis, first convert the analysis to
a template by using the `CreateTemplate` API operation. For `SourceTemplate`, specify the
Amazon Resource Name (ARN) of the source template. The `SourceTemplate` ARN can contain any AWS account; and any QuickSight-supported AWS Region.

Use the `DataSetReferences` entity within `SourceTemplate` to list the replacement
datasets for the placeholders listed in the original. The schema in each dataset must match its placeholder.

_Required_: No

_Type_: [DashboardSourceEntity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-dashboardsourceentity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Contains a map of the key-value pairs for the resource tag or tags assigned to the
dashboard.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThemeArn`

The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If
you add a value for this field, it overrides the value that is used in the source
entity. The theme ARN must exist in the same AWS account where you create
the dashboard.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationStrategy`

The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to `LENIENT`, validation is skipped for specific errors.

_Required_: No

_Type_: [ValidationStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-validationstrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionDescription`

A description for the first version of the dashboard being created.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the dashboard.

`CreatedTime`

The time this dashboard version was created.

`LastPublishedTime`

The time that the dashboard was last published.

`LastUpdatedTime`

The time that the dashboard was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AdHocFilteringOption
