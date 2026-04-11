This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob S3Output

Configuration for uploading output data to Amazon S3 from the processing container.

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

The local path of a directory where you want Amazon SageMaker to upload its contents to Amazon S3.
`LocalPath` is an absolute path to a directory containing output files.
This directory will be created by the platform and exist when your container's
entrypoint is invoked.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3UploadMode`

Whether to upload the results of the processing job continuously or after the job
completes.

_Required_: Yes

_Type_: String

_Allowed values_: `Continuous | EndOfJob`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

The URI of the Amazon S3 prefix Amazon SageMaker downloads data required to run a processing
job.

_Required_: Yes

_Type_: String

_Pattern_: `(https|s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Input

StoppingCondition

All content copied from https://docs.aws.amazon.com/.
