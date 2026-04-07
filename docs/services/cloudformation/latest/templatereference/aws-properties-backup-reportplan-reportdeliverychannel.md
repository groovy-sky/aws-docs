This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::ReportPlan ReportDeliveryChannel

Contains information from your report plan about where to deliver your reports,
specifically your Amazon S3 bucket name, S3 key prefix, and the formats of your
reports.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Formats" : [ String, ... ],
  "S3BucketName" : String,
  "S3KeyPrefix" : String
}

```

### YAML

```yaml

  Formats:
    - String
  S3BucketName: String
  S3KeyPrefix: String

```

## Properties

`Formats`

The format of your reports: `CSV`, `JSON`, or both. If
not specified, the default format is `CSV`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

The unique name of the S3 bucket that receives your reports.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3KeyPrefix`

The prefix for where AWS Backup Audit Manager delivers your reports to Amazon S3. The prefix is this part of the following path:
s3://your-bucket-name/ `prefix`/Backup/us-west-2/year/month/day/report-name.
If not specified, there is no prefix.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Backup::ReportPlan

ReportSetting
