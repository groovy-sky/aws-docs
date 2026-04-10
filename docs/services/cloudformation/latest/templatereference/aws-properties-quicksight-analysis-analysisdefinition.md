This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AnalysisDefinition

The definition of an analysis.

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
  "QueryExecutionOptions" : QueryExecutionOptions,
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
  QueryExecutionOptions:
    QueryExecutionOptions
  Sheets:
    - SheetDefinition
  StaticFiles:
    - StaticFile

```

## Properties

`AnalysisDefaults`

Property description not available.

_Required_: No

_Type_: [AnalysisDefaults](aws-properties-quicksight-analysis-analysisdefaults.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFields`

An array of calculated field definitions for the analysis.

_Required_: No

_Type_: Array of [CalculatedField](aws-properties-quicksight-analysis-calculatedfield.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnConfigurations`

An array of analysis-level column configurations. Column configurations can be used to set default
formatting for a column to be used throughout an analysis.

_Required_: No

_Type_: Array of [ColumnConfiguration](aws-properties-quicksight-analysis-columnconfiguration.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetIdentifierDeclarations`

An array of dataset identifier declarations. This mapping allows the usage of dataset identifiers instead
of dataset ARNs throughout analysis sub-structures.

_Required_: Yes

_Type_: Array of [DataSetIdentifierDeclaration](aws-properties-quicksight-analysis-datasetidentifierdeclaration.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroups`

Filter definitions for an analysis.

For more information, see [Filtering Data in Amazon Quick Sight](../../../quicksight/latest/user/adding-a-filter.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [FilterGroup](aws-properties-quicksight-analysis-filtergroup.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

An array of option definitions for an analysis.

_Required_: No

_Type_: [AssetOptions](aws-properties-quicksight-analysis-assetoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterDeclarations`

An array of parameter declarations for an analysis.

Parameters are named variables that can transfer a value for use by an action or an object.

For more information, see [Parameters in Amazon Quick Sight](../../../quicksight/latest/user/parameters-in-quicksight.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [ParameterDeclaration](aws-properties-quicksight-analysis-parameterdeclaration.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryExecutionOptions`

Property description not available.

_Required_: No

_Type_: [QueryExecutionOptions](aws-properties-quicksight-analysis-queryexecutionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheets`

An array of sheet definitions for an analysis. Each `SheetDefinition` provides detailed information about
a sheet within this analysis.

_Required_: No

_Type_: Array of [SheetDefinition](aws-properties-quicksight-analysis-sheetdefinition.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticFiles`

The static files for the definition.

_Required_: No

_Type_: Array of [StaticFile](aws-properties-quicksight-analysis-staticfile.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisDefaults

AnalysisError

All content copied from https://docs.aws.amazon.com/.
