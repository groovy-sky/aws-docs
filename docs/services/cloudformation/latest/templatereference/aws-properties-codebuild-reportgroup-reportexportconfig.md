This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::ReportGroup ReportExportConfig

Information about the location where the run of a report is exported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExportConfigType" : String,
  "S3Destination" : S3ReportExportConfig
}

```

### YAML

```yaml

  ExportConfigType: String
  S3Destination:
    S3ReportExportConfig

```

## Properties

`ExportConfigType`

The export configuration type. Valid values are:

- `S3`: The report results are exported to an S3 bucket.

- `NO_EXPORT`: The report results are not exported.

_Required_: Yes

_Type_: String

_Allowed values_: `S3 | NO_EXPORT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Destination`

A `S3ReportExportConfig` object that contains information about the S3
bucket where the run of a report is exported.

_Required_: No

_Type_: [S3ReportExportConfig](aws-properties-codebuild-reportgroup-s3reportexportconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CodeBuild::ReportGroup

S3ReportExportConfig

All content copied from https://docs.aws.amazon.com/.
