This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::EvaluationForm

Creates an evaluation form for the specified Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::EvaluationForm",
  "Properties" : {
      "AutoEvaluationConfiguration" : AutoEvaluationConfiguration,
      "Description" : String,
      "InstanceArn" : String,
      "Items" : [ EvaluationFormBaseItem, ... ],
      "LanguageConfiguration" : EvaluationFormLanguageConfiguration,
      "ReviewConfiguration" : EvaluationReviewConfiguration,
      "ScoringStrategy" : ScoringStrategy,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "TargetConfiguration" : EvaluationFormTargetConfiguration,
      "Title" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::EvaluationForm
Properties:
  AutoEvaluationConfiguration:
    AutoEvaluationConfiguration
  Description: String
  InstanceArn: String
  Items:
    - EvaluationFormBaseItem
  LanguageConfiguration:
    EvaluationFormLanguageConfiguration
  ReviewConfiguration:
    EvaluationReviewConfiguration
  ScoringStrategy:
    ScoringStrategy
  Status: String
  Tags:
    - Tag
  TargetConfiguration:
    EvaluationFormTargetConfiguration
  Title: String

```

## Properties

`AutoEvaluationConfiguration`

The automatic evaluation configuration of an evaluation form.

_Required_: No

_Type_: [AutoEvaluationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-autoevaluationconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the evaluation form.

_Length Constraints_: Minimum length of 0. Maximum length of
1024.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The identifier of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Items`

Items that are part of the evaluation form. The total number of sections and questions
must not exceed 100 each. Questions must be contained in a section.

_Minimum size_: 1

_Maximum size_: 100

_Required_: Yes

_Type_: Array of [EvaluationFormBaseItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-evaluationformbaseitem.html)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LanguageConfiguration`

Configuration for language settings of this evaluation form.

_Required_: No

_Type_: [EvaluationFormLanguageConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-evaluationformlanguageconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReviewConfiguration`

Configuration for evaluation review settings of this evaluation form.

_Required_: No

_Type_: [EvaluationReviewConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-evaluationreviewconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScoringStrategy`

A scoring strategy of the evaluation form.

_Required_: No

_Type_: [ScoringStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-scoringstrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the evaluation form.

_Allowed values_: `DRAFT` \| `ACTIVE`

_Required_: Yes

_Type_: String

_Allowed values_: `DRAFT | ACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetConfiguration`

Configuration that specifies the target for this evaluation form.

_Required_: No

_Type_: [EvaluationFormTargetConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

A title of the evaluation form.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the evaluation form name. For example:

`{ "Ref": "myEvaluationFormName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EvaluationFormArn`

The Amazon Resource Name (ARN) of the evaluation form.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AutoEvaluationConfiguration
