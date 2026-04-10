This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MonitoringSchedule S3Output

Information about where and how you want to store the results of a monitoring job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LocalPath" : String,
  "S3UploadMode" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  LocalPath: String
  S3UploadMode: String
  S3Uri: String

```

## Properties

`LocalPath`

The local path to the S3 storage location where SageMaker saves the results of a monitoring job. LocalPath is an
absolute path for the output data.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3UploadMode`

Whether to upload the results of the monitoring job continuously or after the job completes.

_Required_: No

_Type_: String

_Allowed values_: `Continuous | EndOfJob`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Uri`

A URI that identifies the S3 storage location where SageMaker saves the results of a monitoring job.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfig

ScheduleConfig

All content copied from https://docs.aws.amazon.com/.
