This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage TransformOutput

Describes the results of a transform job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Accept" : String,
  "AssembleWith" : String,
  "KmsKeyId" : String,
  "S3OutputPath" : String
}

```

### YAML

```yaml

  Accept: String
  AssembleWith: String
  KmsKeyId: String
  S3OutputPath: String

```

## Properties

`Accept`

The MIME type used to specify the output data. Amazon SageMaker uses the MIME type with each http
call to transfer data from the transform job.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssembleWith`

Defines how to assemble the results of the transform job as a single S3 object. Choose
a format that is most convenient to you. To concatenate the results in binary format,
specify `None`. To add a newline character at the end of every transformed
record, specify
`Line`.

_Required_: No

_Type_: String

_Allowed values_: `None | Line`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using
Amazon S3 server-side encryption. The `KmsKeyId` can be any of the following
formats:

- Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`

- Key ARN:
`arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`

- Alias name: `alias/ExampleAlias`

- Alias name ARN:
`arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias`

If you don't provide a KMS key ID, Amazon SageMaker uses the default KMS key for Amazon S3 for your
role's account. For more information, see [KMS-Managed Encryption Keys](../../../s3/latest/dev/usingkmsencryption.md) in the
_Amazon Simple Storage Service_
_Developer Guide._

The KMS key policy must grant permission to the IAM role that you specify in your
[CreateModel](../../../../reference/sagemaker/latest/apireference/api-createmodel.md)
request. For more information, see [Using\
Key Policies in AWS KMS](../../../kms/latest/developerguide/key-policies.md) in the _AWS Key Management Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3OutputPath`

The Amazon S3 path where you want Amazon SageMaker to store the results of the transform job. For
example, `s3://bucket-name/key-name-prefix`.

For every S3 object used as input for the transform job, batch transform stores the
transformed data with an . `out` suffix in a corresponding subfolder in the
location in the output prefix. For example, for the input data stored at
`s3://bucket-name/input-name-prefix/dataset01/data.csv`, batch transform
stores the transformed data at
`s3://bucket-name/output-name-prefix/input-name-prefix/data.csv.out`.
Batch transform doesn't upload partially processed objects. For an input S3 object that
contains multiple records, it creates an . `out` file only if the transform
job succeeds on the entire file. When the input contains multiple S3 objects, the batch
transform job processes the listed S3 objects and uploads only the output for
successfully processed objects. If any object fails in the transform job batch transform
marks the job as failed to prompt investigation.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformJobDefinition

TransformResources

All content copied from https://docs.aws.amazon.com/.
