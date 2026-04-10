This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DashboardVersionDefinition

The contents of a dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnalysisDefaults" : AnalysisDefaults,
  "CalculatedFields" : [ CalculatedField, ... ],
  "ColumnConfigurations" : [ ColumnConfiguration, ... ],
  "DataSetIdentifierDeclarations" : [ DataSetIdentifierDeclaration, ... ],
  "FilterGroups" : [ FilterGroup, ... ],
  "Options" : AssetOptions,
  "ParameterDeclarations" : [ ParameterDeclaration, ... ],
  "Sheets" : [ SheetDefinition, ... ],
  "StaticFiles" : [ StaticFile, ... ]
}

```

### YAML

```yaml

  AnalysisDefaults:
    AnalysisDefaults
  CalculatedFields:
    - CalculatedField
  ColumnConfigurations:
    - ColumnConfiguration
  DataSetIdentifierDeclarations:
    - DataSetIdentifierDeclaration
  FilterGroups:
    - FilterGroup
  Options:
    AssetOptions
  ParameterDeclarations:
    - ParameterDeclaration
  Sheets:
    - SheetDefinition
  StaticFiles:
    - StaticFile

```

## Properties

`AnalysisDefaults`

Property description not available.

_Required_: No

_Type_: [AnalysisDefaults](aws-properties-quicksight-dashboard-analysisdefaults.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFields`

An array of calculated field definitions for the dashboard.

_Required_: No

_Type_: Array of [CalculatedField](aws-properties-quicksight-dashboard-calculatedfield.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnConfigurations`

An array of dashboard-level column configurations. Column configurations
are used to set the default formatting for a column that
is used throughout a dashboard.

_Required_: No

_Type_: Array of [ColumnConfiguration](aws-properties-quicksight-dashboard-columnconfiguration.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetIdentifierDeclarations`

An array of dataset identifier declarations. With
this mapping,you can use dataset identifiers instead of dataset Amazon Resource Names (ARNs) throughout the dashboard's sub-structures.

_Required_: Yes

_Type_: Array of [DataSetIdentifierDeclaration](aws-properties-quicksight-dashboard-datasetidentifierdeclaration.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroups`

The filter definitions for a dashboard.

For more information, see [Filtering Data in Amazon Quick Sight](../../../quicksight/latest/user/adding-a-filter.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [FilterGroup](aws-properties-quicksight-dashboard-filtergroup.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

An array of option definitions for a dashboard.

_Required_: No

_Type_: [AssetOptions](aws-properties-quicksight-dashboard-assetoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterDeclarations`

The parameter declarations for a dashboard. Parameters are named variables that can transfer a value for use by an action or an object.

For more information, see [Parameters in Amazon Quick Sight](../../../quicksight/latest/user/parameters-in-quicksight.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [ParameterDeclaration](aws-properties-quicksight-dashboard-parameterdeclaration.md)

_Minimum_: `0`

_Maximum_: `400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheets`

An array of sheet definitions for a dashboard.

_Required_: No

_Type_: Array of [SheetDefinition](aws-properties-quicksight-dashboard-sheetdefinition.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticFiles`

The static files for the definition.

_Required_: No

_Type_: Array of [StaticFile](aws-properties-quicksight-dashboard-staticfile.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashboardVersion

DashboardVisualPublishOptions

All content copied from https://docs.aws.amazon.com/.
