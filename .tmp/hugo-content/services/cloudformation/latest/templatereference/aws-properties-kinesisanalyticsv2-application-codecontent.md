This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application CodeContent

Specifies either the application code, or the location of the application code, for a
Managed Service for Apache Flink application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3ContentLocation" : S3ContentLocation,
  "TextContent" : String,
  "ZipFileContent" : String
}

```

### YAML

```yaml

  S3ContentLocation:
    S3ContentLocation
  TextContent: String
  ZipFileContent: String

```

## Properties

`S3ContentLocation`

Information about the Amazon S3 bucket that contains the application code.

_Required_: No

_Type_: [S3ContentLocation](aws-properties-kinesisanalyticsv2-application-s3contentlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TextContent`

The text-format code for a Managed Service for Apache Flink application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `102400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZipFileContent`

The zip-format code for a Managed Service for Apache Flink application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CodeContent](../../../managed-flink/latest/apiv2/api-codecontent.md) in the _Amazon Kinesis Data Analytics API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CheckpointConfiguration

CSVMappingParameters

All content copied from https://docs.aws.amazon.com/.
