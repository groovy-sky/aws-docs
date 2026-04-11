This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset S3DestinationConfiguration

Configuration information for delivery of dataset contents to Amazon Simple Storage Service (Amazon S3).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "GlueConfiguration" : GlueConfiguration,
  "Key" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  Bucket: String
  GlueConfiguration:
    GlueConfiguration
  Key: String
  RoleArn: String

```

## Properties

`Bucket`

The name of the S3 bucket to which dataset contents are delivered.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9.\-_]*$`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueConfiguration`

Configuration information for coordination with AWS Glue, a fully managed extract, transform
and load (ETL) service.

_Required_: No

_Type_: [GlueConfiguration](aws-properties-iotanalytics-dataset-glueconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the dataset contents object in an S3 bucket. Each object has a key that is a
unique identifier. Each object has exactly one key.

You can create a unique key with the following options:

- Use `!{iotanalytics:scheduleTime}` to insert the time of a scheduled SQL
query run.

- Use `!{iotanalytics:versionId}` to insert a unique hash that identifies a
dataset content.

- Use `!{iotanalytics:creationTime}` to insert the creation time of a dataset
content.

The following example creates a unique key for a CSV file:
`dataset/mydataset/!{iotanalytics:scheduleTime}/!{iotanalytics:versionId}.csv`

###### Note

If you don't use `!{iotanalytics:versionId}` to specify the key, you might
get duplicate keys. For example, you might have two dataset contents with the same
`scheduleTime` but different `versionId` s. This means that one
dataset content overwrites the other.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9!_.*'()/{}:-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 and AWS Glue
resources.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetentionPeriod

Schedule

All content copied from https://docs.aws.amazon.com/.
