This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelQualityJobDefinition S3Output

The Amazon S3 storage location where the results of a monitoring job are saved.

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

The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
LocalPath is an absolute path for the output data.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3UploadMode`

Whether to upload the results of the monitoring job continuously or after the job completes.

_Required_: No

_Type_: String

_Allowed values_: `Continuous | EndOfJob`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring
job.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfig

StoppingCondition

All content copied from https://docs.aws.amazon.com/.
