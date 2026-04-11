This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier DocumentClassifierOutputDataConfig

Provide the location for output data from a custom classifier job. This field is mandatory
if you are training a native document model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  KmsKeyId: String
  S3Uri: String

```

## Properties

`KmsKeyId`

ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to encrypt the
output results from an analysis job. The KmsKeyId can be one of the following formats:

- KMS Key ID: `"1234abcd-12ab-34cd-56ef-1234567890ab"`

- Amazon Resource Name (ARN) of a KMS Key:
`"arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"`

- KMS Key Alias: `"alias/ExampleAlias"`

- ARN of a KMS Key Alias:
`"arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias"`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

When you use the `OutputDataConfig` object while creating a custom
classifier, you specify the Amazon S3 location where you want to write the confusion matrix
and other output files.
The URI must be in the same Region as the API endpoint that you are calling. The location is
used as the prefix for the actual location of this output file.

When the custom classifier job is finished, the service creates the output file in a
directory specific to the job. The `S3Uri` field contains the location of the
output file, called `output.tar.gz`. It is a compressed archive that contains the
confusion matrix.

_Required_: No

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentClassifierInputDataConfig

DocumentReaderConfig

All content copied from https://docs.aws.amazon.com/.
