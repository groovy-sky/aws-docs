This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::ReportGroup

Represents a report group. A report group contains a collection of reports.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeBuild::ReportGroup",
  "Properties" : {
      "DeleteReports" : Boolean,
      "ExportConfig" : ReportExportConfig,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodeBuild::ReportGroup
Properties:
  DeleteReports: Boolean
  ExportConfig:
    ReportExportConfig
  Name: String
  Tags:
    - Tag
  Type: String

```

## Properties

`DeleteReports`

When deleting a report group, specifies if reports within the report group should be
deleted.

true

Deletes any reports that belong to the report group before deleting the
report group.

false

You must delete any reports in the report group. This is the default
value. If you delete a report group that contains one or more reports, an
exception is thrown.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExportConfig`

Information about the destination where the raw data of this `ReportGroup`
is exported.

_Required_: Yes

_Type_: [ReportExportConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-reportgroup-reportexportconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the `ReportGroup`.

_Required_: No

_Type_: String

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tag key and value pairs associated with this report group.

These tags are available for use by AWS services that support AWS CodeBuild report group
tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codebuild-reportgroup-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the `ReportGroup`. This can be one of the following
values:

CODE\_COVERAGE

The report group contains code coverage reports.

TEST

The report group contains test reports.

_Required_: Yes

_Type_: String

_Allowed values_: `TEST | CODE_COVERAGE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the ARN of the AWS CodeBuild report group, such as `arn:aws:codebuild:region:123456789012:report-group/myReportGroupName`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available
attributes and sample return values. For more information about using `Fn::GetAtt`,
see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

The ARN of the AWS CodeBuild report group, such as
`arn:aws:codebuild:region:123456789012:report-group/myReportGroupName`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WebhookFilter

ReportExportConfig
