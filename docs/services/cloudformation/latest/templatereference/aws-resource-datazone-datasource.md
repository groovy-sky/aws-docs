This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource

The `AWS::DataZone::DataSource` resource specifies an Amazon DataZone data source that is used to
import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::DataSource",
  "Properties" : {
      "AssetFormsInput" : [ FormInput, ... ],
      "Configuration" : DataSourceConfigurationInput,
      "ConnectionIdentifier" : String,
      "Description" : String,
      "DomainIdentifier" : String,
      "EnableSetting" : String,
      "EnvironmentIdentifier" : String,
      "Name" : String,
      "ProjectIdentifier" : String,
      "PublishOnImport" : Boolean,
      "Recommendation" : RecommendationConfiguration,
      "Schedule" : ScheduleConfiguration,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::DataSource
Properties:
  AssetFormsInput:
    - FormInput
  Configuration:
    DataSourceConfigurationInput
  ConnectionIdentifier: String
  Description: String
  DomainIdentifier: String
  EnableSetting: String
  EnvironmentIdentifier: String
  Name: String
  ProjectIdentifier: String
  PublishOnImport: Boolean
  Recommendation:
    RecommendationConfiguration
  Schedule:
    ScheduleConfiguration
  Type: String

```

## Properties

`AssetFormsInput`

The metadata forms attached to the assets that the data source works with.

_Required_: No

_Type_: Array of [FormInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-forminput.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

The configuration of the data source.

_Required_: No

_Type_: [DataSourceConfigurationInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-datasourceconfigurationinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionIdentifier`

The ID of the connection.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the data source.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The ID of the Amazon DataZone domain where the data source is created.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableSetting`

Specifies whether the data source is enabled.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentIdentifier`

The unique identifier of the Amazon DataZone environment to which the data source publishes
assets.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectIdentifier`

The identifier of the Amazon DataZone project in which you want to add this data
source.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublishOnImport`

Specifies whether the assets that this data source creates in the inventory are to be
also automatically published to the catalog.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recommendation`

Specifies whether the business name generation is to be enabled for this data
source.

_Required_: No

_Type_: [RecommendationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-recommendationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule of the data source runs.

_Required_: No

_Type_: [ScheduleConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-datasource-scheduleconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the data source. In Amazon DataZone, you can use data sources to import
technical metadata of assets (data) from the source databases or data warehouses into
Amazon DataZone. In the current release of Amazon DataZone, you can create and run data
sources for AWS Glue and Amazon Redshift.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and the
`DataSourceId`, which uniquely identifies the data source. For example: `{ "Ref": "MyDataSource"
    }` for the resource with the logical ID `MyDataSource`, `Ref` returns
`DomainId|DataSourceId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectionId`

The connection ID that's part of the data source summary.

`CreatedAt`

The timestamp of when the data source was created.

`DomainId`

The ID of the Amazon DataZone domain in which the data source exists.

`EnvironmentId`

The ID of the environment in which the data source exists.

`Id`

The identifier of the data source run.

`LastRunAssetCount`

The count of the assets created during the last data source run.

`LastRunAt`

The timestamp of when the data source run was last performed.

`LastRunStatus`

The status of the last data source run.

`ProjectId`

The project ID included in the data source run activity.

`Status`

The status of the data source.

`UpdatedAt`

The timestamp of when the data source was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkflowsMwaaPropertiesInput

DataSourceConfigurationInput
