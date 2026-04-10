This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage S3DataSource

Describes the S3 data source.

Your input bucket must be in the same AWS region as your training
job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3DataType" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  S3DataType: String
  S3Uri: String

```

## Properties

`S3DataType`

If you choose `S3Prefix`, `S3Uri` identifies a key name prefix.
SageMaker uses all objects that match the specified key name prefix for model training.

If you choose `ManifestFile`, `S3Uri` identifies an object that
is a manifest file containing a list of object keys that you want SageMaker to use for model
training.

If you choose `AugmentedManifestFile`, `S3Uri` identifies an
object that is an augmented manifest file in JSON lines format. This file contains the
data you want to use for model training. `AugmentedManifestFile` can only be
used if the Channel's input mode is `Pipe`.

If you choose `Converse`, `S3Uri` identifies an Amazon S3 location
that contains data formatted according to Converse format. This format structures conversational
messages with specific roles and content types used for training and fine-tuning foundational models.

_Required_: Yes

_Type_: String

_Allowed values_: `ManifestFile | S3Prefix | AugmentedManifestFile`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Uri`

Depending on the value specified for the `S3DataType`, identifies either
a key name prefix or a manifest. For example:

- A key name prefix might look like this:
`s3://bucketname/exampleprefix/`

- A manifest might look like this:
`s3://bucketname/example.manifest`

A manifest is an S3 object which is a JSON file consisting of an array of
elements. The first element is a prefix which is followed by one or more
suffixes. SageMaker appends the suffix elements to the prefix to get a full set of
`S3Uri`. Note that the prefix must be a valid non-empty
`S3Uri` that precludes users from specifying a manifest whose
individual `S3Uri` is sourced from different S3 buckets.

The following code example shows a valid manifest format:

`[ {"prefix": "s3://customer_bucket/some/prefix/"},`

` "relative/path/to/custdata-1",`

` "relative/path/custdata-2",`

` ...`

` "relative/path/custdata-N"`

`]`

This JSON is equivalent to the following `S3Uri`
list:

`s3://customer_bucket/some/prefix/relative/path/to/custdata-1`

`s3://customer_bucket/some/prefix/relative/path/custdata-2`

`...`

`s3://customer_bucket/some/prefix/relative/path/custdata-N`

The complete set of `S3Uri` in this manifest is the input data
for the channel for this data source. The object that each `S3Uri`
points to must be readable by the IAM role that SageMaker uses to perform tasks on
your behalf.

Your input bucket must be located in same AWS region as your
training job.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|s3)://([^/]+)/?(.*)$`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelQuality

S3ModelDataSource

All content copied from https://docs.aws.amazon.com/.
