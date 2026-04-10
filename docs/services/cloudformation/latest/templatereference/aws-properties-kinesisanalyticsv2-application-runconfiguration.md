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

_Type_: [ApplicationRestoreConfiguration](aws-properties-kinesisanalyticsv2-application-applicationrestoreconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlinkRunConfiguration`

Describes the starting parameters for a Managed Service for Apache Flink application.

_Required_: No

_Type_: [FlinkRunConfiguration](aws-properties-kinesisanalyticsv2-application-flinkrunconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordFormat

S3ContentBaseLocation

All content copied from https://docs.aws.amazon.com/.
