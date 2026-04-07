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

_Type_: [AnalysisDefaults](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-analysisdefaults.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFields`

An array of calculated field definitions for the analysis.

_Required_: No

_Type_: Array of [CalculatedField](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-calculatedfield.html)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnConfigurations`

An array of analysis-level column configurations. Column configurations can be used to set default
formatting for a column to be used throughout an analysis.

_Required_: No

_Type_: Array of [ColumnConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-columnconfiguration.html)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetIdentifierDeclarations`

An array of dataset identifier declarations. This mapping allows the usage of dataset identifiers instead
of dataset ARNs throughout analysis sub-structures.

_Required_: Yes

_Type_: Array of [DataSetIdentifierDeclaration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-datasetidentifierdeclaration.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroups`

Filter definitions for an analysis.

For more information, see [Filtering Data in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/adding-a-filter.html) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [FilterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-filtergroup.html)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

An array of option definitions for an analysis.

_Required_: No

_Type_: [AssetOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-assetoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterDeclarations`

An array of parameter declarations for an analysis.

Parameters are named variables that can transfer a value for use by an action or an object.

For more information, see [Parameters in Amazon Quick Sight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-in-quicksight.html) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [ParameterDeclaration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-parameterdeclaration.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryExecutionOptions`

Property description not available.

_Required_: No

_Type_: [QueryExecutionOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-queryexecutionoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheets`

An array of sheet definitions for an analysis. Each `SheetDefinition` provides detailed information about
a sheet within this analysis.

_Required_: No

_Type_: Array of [SheetDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-sheetdefinition.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticFiles`

The static files for the definition.

_Required_: No

_Type_: Array of [StaticFile](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-staticfile.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalysisDefaults

AnalysisError
