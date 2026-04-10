This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet

Creates a dataset. This operation doesn't support datasets that include uploaded files
as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::DataSet",
  "Properties" : {
      "AwsAccountId" : String,
      "ColumnGroups" : [ ColumnGroup, ... ],
      "ColumnLevelPermissionRules" : [ ColumnLevelPermissionRule, ... ],
      "DataPrepConfiguration" : DataPrepConfiguration,
      "DataSetId" : String,
      "DatasetParameters" : [ DatasetParameter, ... ],
      "DataSetRefreshProperties" : DataSetRefreshProperties,
      "DataSetUsageConfiguration" : DataSetUsageConfiguration,
      "FieldFolders" : {Key: Value, ...},
      "FolderArns" : [ String, ... ],
      "ImportMode" : String,
      "IngestionWaitPolicy" : IngestionWaitPolicy,
      "Name" : String,
      "PerformanceConfiguration" : PerformanceConfiguration,
      "Permissions" : [ ResourcePermission, ... ],
      "PhysicalTableMap" : {Key: Value, ...},
      "SemanticModelConfiguration" : SemanticModelConfiguration,
      "Tags" : [ Tag, ... ],
      "UseAs" : String
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::DataSet
Properties:
  AwsAccountId: String
  ColumnGroups:
    - ColumnGroup
  ColumnLevelPermissionRules:
    - ColumnLevelPermissionRule
  DataPrepConfiguration:
    DataPrepConfiguration
  DataSetId: String
  DatasetParameters:
    - DatasetParameter
  DataSetRefreshProperties:
    DataSetRefreshProperties
  DataSetUsageConfiguration:
    DataSetUsageConfiguration
  FieldFolders:
    Key: Value
  FolderArns:
    - String
  ImportMode: String
  IngestionWaitPolicy:
    IngestionWaitPolicy
  Name: String
  PerformanceConfiguration:
    PerformanceConfiguration
  Permissions:
    - ResourcePermission
  PhysicalTableMap:
    Key: Value
  SemanticModelConfiguration:
    SemanticModelConfiguration
  Tags:
    - Tag
  UseAs: String

```

## Properties

`AwsAccountId`

The AWS account ID.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnGroups`

Groupings of columns that work together in certain Amazon Quick Sight features.
Currently, only geospatial hierarchy is supported.

_Required_: No

_Type_: Array of [ColumnGroup](aws-properties-quicksight-dataset-columngroup.md)

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnLevelPermissionRules`

A set of one or more definitions of a ` ColumnLevelPermissionRule `.

_Required_: No

_Type_: Array of [ColumnLevelPermissionRule](aws-properties-quicksight-dataset-columnlevelpermissionrule.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPrepConfiguration`

The data preparation configuration associated with this dataset.

_Required_: No

_Type_: [DataPrepConfiguration](aws-properties-quicksight-dataset-dataprepconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetId`

An ID for the dataset that you want to create. This ID is unique per AWS Region for each AWS account.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatasetParameters`

The parameters that are declared in a dataset.

_Required_: No

_Type_: Array of [DatasetParameter](aws-properties-quicksight-dataset-datasetparameter.md)

_Minimum_: `0`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetRefreshProperties`

The refresh properties of a dataset.

_Required_: No

_Type_: [DataSetRefreshProperties](aws-properties-quicksight-dataset-datasetrefreshproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetUsageConfiguration`

The usage configuration to apply to child datasets that reference this dataset as a source.

_Required_: No

_Type_: [DataSetUsageConfiguration](aws-properties-quicksight-dataset-datasetusageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldFolders`

The folder that contains fields and nested subfolders for your dataset.

_Required_: No

_Type_: Object of [FieldFolder](aws-properties-quicksight-dataset-fieldfolder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderArns`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImportMode`

Indicates whether you want to import the data into SPICE.

_Required_: No

_Type_: String

_Allowed values_: `SPICE | DIRECT_QUERY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IngestionWaitPolicy`

The wait policy to use when creating or updating a Dataset. The default is to wait for SPICE ingestion to finish
with timeout of 36 hours.

_Required_: No

_Type_: [IngestionWaitPolicy](aws-properties-quicksight-dataset-ingestionwaitpolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The display name for the dataset.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceConfiguration`

The performance optimization configuration of a dataset.

_Required_: No

_Type_: [PerformanceConfiguration](aws-properties-quicksight-dataset-performanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A list of resource permissions on the dataset.

_Required_: No

_Type_: Array of [ResourcePermission](aws-properties-quicksight-dataset-resourcepermission.md)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhysicalTableMap`

Declares the physical tables that are available in the underlying data sources.

_Required_: No

_Type_: Object of [PhysicalTable](aws-properties-quicksight-dataset-physicaltable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticModelConfiguration`

The semantic model configuration associated with this dataset.

_Required_: No

_Type_: [SemanticModelConfiguration](aws-properties-quicksight-dataset-semanticmodelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Contains a map of the key-value pairs for the resource tag or tags assigned to the
dataset.

_Required_: No

_Type_: Array of [Tag](aws-properties-quicksight-dataset-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAs`

The usage of the dataset.

_Required_: No

_Type_: String

_Allowed values_: `RLS_RULES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the dataset.

`ConsumedSpiceCapacityInBytes`

`CreatedTime`

The time this dataset version was created.

`LastUpdatedTime`

The time this dataset version was last updated.

`OutputColumns`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

YAxisOptions

AggregateOperation

All content copied from https://docs.aws.amazon.com/.
