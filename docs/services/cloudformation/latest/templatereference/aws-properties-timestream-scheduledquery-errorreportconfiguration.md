This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery ErrorReportConfiguration

Configuration required for error reporting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Configuration" : S3Configuration
}

```

### YAML

```yaml

  S3Configuration:
    S3Configuration

```

## Properties

`S3Configuration`

The S3 configuration for the error reports.

_Required_: Yes

_Type_: [S3Configuration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-timestream-scheduledquery-s3configuration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DimensionMapping

MixedMeasureMapping
