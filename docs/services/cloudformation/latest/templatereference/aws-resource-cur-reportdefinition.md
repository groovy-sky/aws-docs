This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CUR::ReportDefinition

The definition of AWS Cost and Usage Report. You can specify the report name,
time unit, report format, compression format, S3 bucket, additional artifacts, and schema
elements in the definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CUR::ReportDefinition",
  "Properties" : {
      "AdditionalArtifacts" : [ String, ... ],
      "AdditionalSchemaElements" : [ String, ... ],
      "BillingViewArn" : String,
      "Compression" : String,
      "Format" : String,
      "RefreshClosedReports" : Boolean,
      "ReportName" : String,
      "ReportVersioning" : String,
      "S3Bucket" : String,
      "S3Prefix" : String,
      "S3Region" : String,
      "Tags" : [ Tag, ... ],
      "TimeUnit" : String
    }
}

```

### YAML

```yaml

Type: AWS::CUR::ReportDefinition
Properties:
  AdditionalArtifacts:
    - String
  AdditionalSchemaElements:
    - String
  BillingViewArn: String
  Compression: String
  Format: String
  RefreshClosedReports: Boolean
  ReportName: String
  ReportVersioning: String
  S3Bucket: String
  S3Prefix: String
  S3Region: String
  Tags:
    - Tag
  TimeUnit: String

```

## Properties

`AdditionalArtifacts`

A list of manifests that you want AWS to create for this report.

_Required_: No

_Type_: Array of String

_Allowed values_: `REDSHIFT | QUICKSIGHT | ATHENA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalSchemaElements`

A list of strings that indicate additional content that AWS includes in the report, such as individual resource IDs.

_Required_: No

_Type_: Array of String

_Allowed values_: `RESOURCES | SPLIT_COST_ALLOCATION_DATA | MANUAL_DISCOUNT_COMPATIBILITY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BillingViewArn`

The Amazon Resource Name (ARN) of the billing view. You can get this value by using the billing view service public APIs.

_Required_: No

_Type_: String

_Pattern_: `(arn:aws(-cn)?:billing::[0-9]{12}:billingview/)?[a-zA-Z0-9_\+=\.\-@].{1,30}`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Compression`

The compression format that Amazon Web Services uses for the report.

_Required_: Yes

_Type_: String

_Allowed values_: `ZIP | GZIP | Parquet`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The format that Amazon Web Services saves the report in.

_Required_: Yes

_Type_: String

_Allowed values_: `textORcsv | Parquet`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshClosedReports`

Whether you want AWS to update your reports after they have been finalized if AWS detects charges related to
previous months. These charges can include refunds, credits, or support fees.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportName`

The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z!\-_.*\'()]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReportVersioning`

Whether you want AWS to overwrite the previous version of each report or
to deliver the report in addition to the previous versions.

_Required_: Yes

_Type_: String

_Allowed values_: `CREATE_NEW_REPORT | OVERWRITE_REPORT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Bucket`

The S3 bucket where Amazon Web Services delivers the report.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9_\.\-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Prefix`

The prefix that Amazon Web Services adds to the report name when Amazon Web Services delivers the report. Your prefix can't include spaces.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9A-Za-z!\-_.*\'()/]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Region`

The Region of the S3 bucket that Amazon Web Services delivers the report into.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the report definition resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cur-reportdefinition-tag.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeUnit`

The granularity of the line items in the report.

_Required_: Yes

_Type_: String

_Allowed values_: `HOURLY | DAILY | MONTHLY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns

`{ "Ref": "ReportName" }`

The name of the report that you want to create. The name must be unique, is case sensitive, and can't include spaces.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Cost and Usage Report

Tag
