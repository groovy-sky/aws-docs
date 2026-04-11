This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInsightsAccessScopeAnalysis

Describes a Network Access Scope analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInsightsAccessScopeAnalysis",
  "Properties" : {
      "NetworkInsightsAccessScopeId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInsightsAccessScopeAnalysis
Properties:
  NetworkInsightsAccessScopeId: String
  Tags:
    - Tag

```

## Properties

`NetworkInsightsAccessScopeId`

The ID of the Network Access Scope.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-networkinsightsaccessscopeanalysis-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network insights analysis.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AnalyzedEniCount`

The number of network interfaces analyzed.

`EndDate`

The end date of the analysis.

`FindingsFound`

Indicates whether there are findings (true \| false \| unknown).

`NetworkInsightsAccessScopeAnalysisArn`

The ARN of the Network Access Scope analysis.

`NetworkInsightsAccessScopeAnalysisId`

The ID of the Network Access Scope analysis.

`StartDate`

The start date of the analysis.

`Status`

The status of the analysis (running \| succeeded \| failed).

`StatusMessage`

The status message.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThroughResourcesStatementRequest

Tag

All content copied from https://docs.aws.amazon.com/.
