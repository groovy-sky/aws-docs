This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TemplateVersionDefinition

The detailed definition of a template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnalysisDefaults" : AnalysisDefaults,
  "CalculatedFields" : [ CalculatedField, ... ],
  "ColumnConfigurations" : [ ColumnConfiguration, ... ],
  "DataSetConfigurations" : [ DataSetConfiguration, ... ],
  "FilterGroups" : [ FilterGroup, ... ],
  "Options" : AssetOptions,
  "ParameterDeclarations" : [ ParameterDeclaration, ... ],
  "QueryExecutionOptions" : QueryExecutionOptions,
  "Sheets" : [ SheetDefinition, ... ]
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
  DataSetConfigurations:
    - DataSetConfiguration
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

```

## Properties

`AnalysisDefaults`

Property description not available.

_Required_: No

_Type_: [AnalysisDefaults](aws-properties-quicksight-template-analysisdefaults.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFields`

An array of calculated field definitions for the template.

_Required_: No

_Type_: Array of [CalculatedField](aws-properties-quicksight-template-calculatedfield.md)

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnConfigurations`

An array of template-level column
configurations. Column configurations are used to set default formatting for a column that's used throughout a template.

_Required_: No

_Type_: Array of [ColumnConfiguration](aws-properties-quicksight-template-columnconfiguration.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetConfigurations`

An array of dataset configurations. These configurations define the required columns for each dataset used within a template.

_Required_: Yes

_Type_: Array of [DataSetConfiguration](aws-properties-quicksight-template-datasetconfiguration.md)

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterGroups`

Filter definitions for a template.

For more information, see [Filtering Data](../../../quicksight/latest/user/filtering-visual-data.md) in the _Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [FilterGroup](aws-properties-quicksight-template-filtergroup.md)

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Options`

An array of option definitions for a template.

_Required_: No

_Type_: [AssetOptions](aws-properties-quicksight-template-assetoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterDeclarations`

An array of parameter declarations for a template.

_Parameters_ are named variables that can transfer a value for use by an action or an object.

For more information, see [Parameters in Amazon Quick Sight](../../../quicksight/latest/user/parameters-in-quicksight.md) in the
_Amazon Quick Suite User Guide_.

_Required_: No

_Type_: Array of [ParameterDeclaration](aws-properties-quicksight-template-parameterdeclaration.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryExecutionOptions`

Property description not available.

_Required_: No

_Type_: [QueryExecutionOptions](aws-properties-quicksight-template-queryexecutionoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheets`

An array of sheet definitions for a template.

_Required_: No

_Type_: Array of [SheetDefinition](aws-properties-quicksight-template-sheetdefinition.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateVersion

TextAreaControlDisplayOptions

All content copied from https://docs.aws.amazon.com/.
