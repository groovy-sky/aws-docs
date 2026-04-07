This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAnalysis

Specifies a network insights analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInsightsAnalysis",
  "Properties" : {
      "AdditionalAccounts" : [ String, ... ],
      "FilterInArns" : [ String, ... ],
      "FilterOutArns" : [ String, ... ],
      "NetworkInsightsPathId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInsightsAnalysis
Properties:
  AdditionalAccounts:
    - String
  FilterInArns:
    - String
  FilterOutArns:
    - String
  NetworkInsightsPathId: String
  Tags:
    - Tag

```

## Properties

`AdditionalAccounts`

The member accounts that contain resources that the path can traverse.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterInArns`

The Amazon Resource Names (ARN) of the resources that the path must traverse.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilterOutArns`

The Amazon Resource Names (ARN) of the resources that the path must ignore.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInsightsPathId`

The ID of the path.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-networkinsightsanalysis-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network insights analysis.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AlternatePathHints`

Potential intermediate components.

`Explanations`

The explanations. For more information, see [Reachability Analyzer explanation\
codes](https://docs.aws.amazon.com/vpc/latest/reachability/explanation-codes.html).

`ForwardPathComponents`

The components in the path from source to destination.

`NetworkInsightsAnalysisArn`

The Amazon Resource Name (ARN) of the network insights analysis.

`NetworkInsightsAnalysisId`

The ID of the network insights analysis.

`NetworkPathFound`

Indicates whether the destination is reachable from the source.

`ReturnPathComponents`

The components in the path from destination to source.

`StartDate`

The time the analysis started.

`Status`

The status of the network insights analysis.

`StatusMessage`

The status message, if the status is `failed`.

`SuggestedAccounts`

The IDs of potential intermediate accounts.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AdditionalDetail
