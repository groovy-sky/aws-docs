This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application RunConfiguration

Describes the starting parameters for an Managed Service for Apache Flink application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationRestoreConfiguration" : ApplicationRestoreConfiguration,
  "FlinkRunConfiguration" : FlinkRunConfiguration
}

```

### YAML

```yaml

  ApplicationRestoreConfiguration:
    ApplicationRestoreConfiguration
  FlinkRunConfiguration:
    FlinkRunConfiguration

```

## Properties

`ApplicationRestoreConfiguration`

Describes the restore behavior of a restarting application.

_Required_: No

_Type_: [ApplicationRestoreConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-applicationrestoreconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlinkRunConfiguration`

Describes the starting parameters for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [FlinkRunConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisanalyticsv2-application-flinkrunconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecordFormat

S3ContentBaseLocation
