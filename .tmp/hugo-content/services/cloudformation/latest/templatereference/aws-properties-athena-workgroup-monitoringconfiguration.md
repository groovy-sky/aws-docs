This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup MonitoringConfiguration

Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets,
Amazon CloudWatch log groups etc.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLoggingConfiguration" : CloudWatchLoggingConfiguration,
  "ManagedLoggingConfiguration" : ManagedLoggingConfiguration,
  "S3LoggingConfiguration" : S3LoggingConfiguration
}

```

### YAML

```yaml

  CloudWatchLoggingConfiguration:
    CloudWatchLoggingConfiguration
  ManagedLoggingConfiguration:
    ManagedLoggingConfiguration
  S3LoggingConfiguration:
    S3LoggingConfiguration

```

## Properties

`CloudWatchLoggingConfiguration`

Configuration settings for delivering logs to Amazon CloudWatch log groups.

_Required_: No

_Type_: [CloudWatchLoggingConfiguration](aws-properties-athena-workgroup-cloudwatchloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedLoggingConfiguration`

Configuration settings for managed log persistence.

_Required_: No

_Type_: [ManagedLoggingConfiguration](aws-properties-athena-workgroup-managedloggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3LoggingConfiguration`

Configuration settings for delivering logs to Amazon S3 buckets.

_Required_: No

_Type_: [S3LoggingConfiguration](aws-properties-athena-workgroup-s3loggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedStorageEncryptionConfiguration

ResultConfiguration

All content copied from https://docs.aws.amazon.com/.
