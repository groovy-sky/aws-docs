This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis

Creates an analysis in Amazon QuickSight.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::Analysis",
  "Properties" : {
      "AnalysisId" : String,
      "AwsAccountId" : String,
      "Definition" : AnalysisDefinition,
      "Errors" : [ AnalysisError, ... ],
      "FolderArns" : [ String, ... ],
      "Name" : String,
      "Parameters" : Parameters,
      "Permissions" : [ ResourcePermission, ... ],
      "Sheets" : [ Sheet, ... ],
      "SourceEntity" : AnalysisSourceEntity,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "ThemeArn" : String,
      "ValidationStrategy" : ValidationStrategy
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::Analysis
Properties:
  AnalysisId: String
  AwsAccountId: String
  Definition:
    AnalysisDefinition
  Errors:
    - AnalysisError
  FolderArns:
    - String
  Name: String
  Parameters:
    Parameters
  Permissions:
    - ResourcePermission
  Sheets:
    - Sheet
  SourceEntity:
    AnalysisSourceEntity
  Status: String
  Tags:
    - Tag
  ThemeArn: String
  ValidationStrategy:
    ValidationStrategy

```

## Properties

`AnalysisId`

The ID for the analysis that you're creating. This ID displays in the URL of the
analysis.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AwsAccountId`

The ID of the AWS account where you are creating an analysis.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Definition`

Property description not available.

_Required_: No

_Type_: [AnalysisDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-analysisdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Errors`

Errors associated with the analysis.

_Required_: No

_Type_: Array of [AnalysisError](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-analysiserror.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderArns`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A descriptive name for the analysis that you're creating. This name displays for the
analysis in the Amazon Quick Sight console.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameter names and override values that you want to use. An analysis can have
any parameter type, and some parameters might accept multiple values.

_Required_: No

_Type_: [Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-parameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A structure that describes the principals and the resource-level permissions on an
analysis. You can use the `Permissions` structure to grant permissions by
providing a list of AWS Identity and Access Management (IAM) action information for each
principal listed by Amazon Resource Name (ARN).

To specify no permissions, omit `Permissions`.

_Required_: No

_Type_: Array of [ResourcePermission](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-resourcepermission.html)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sheets`

A list of the associated sheets with the unique identifier and name of each sheet.

_Required_: No

_Type_: Array of [Sheet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-sheet.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceEntity`

A source entity to use for the analysis that you're creating. This metadata structure
contains details that describe a source template and one or more datasets.

Either a `SourceEntity` or a `Definition` must be provided in
order for the request to be valid.

_Required_: No

_Type_: [AnalysisSourceEntity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-analysissourceentity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Status associated with the analysis.

_Required_: No

_Type_: String

_Allowed values_: `CREATION_IN_PROGRESS | CREATION_SUCCESSFUL | CREATION_FAILED | UPDATE_IN_PROGRESS | UPDATE_SUCCESSFUL | UPDATE_FAILED | PENDING_UPDATE | DELETED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Contains a map of the key-value pairs for the resource tag or tags assigned to the
analysis.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-tag.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThemeArn`

The ARN for the theme to apply to the analysis that you're creating. To see the theme
in the Amazon Quick Sight console, make sure that you have access to it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationStrategy`

The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects. When you set this value to `LENIENT`, validation is skipped for specific errors.

_Required_: No

_Type_: [ValidationStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-analysis-validationstrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the analysis.

`CreatedTime`

The time that the analysis was created.

`DataSetArns`

The ARNs of the datasets of the analysis.

`LastUpdatedTime`

The time that the analysis was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AggregationFunction
