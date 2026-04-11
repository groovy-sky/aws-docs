This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob S3Input

Configuration for downloading input data from Amazon S3 into the processing
container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LocalPath" : String,
  "S3CompressionType" : String,
  "S3DataDistributionType" : String,
  "S3DataType" : String,
  "S3InputMode" : String,
  "S3Uri" : String
}

```

### YAML

```yaml

  LocalPath: String
  S3CompressionType: String
  S3DataDistributionType: String
  S3DataType: String
  S3InputMode: String
  S3Uri: String

```

## Properties

`LocalPath`

The local path in your container where you want Amazon SageMaker to write input data to.
`LocalPath` is an absolute path to the input data and must begin with
`/opt/ml/processing/`. `LocalPath` is a required parameter
when `AppManaged` is `False` (default).

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3CompressionType`

Whether to GZIP-decompress the data in Amazon S3 as it is streamed into the processing
container. `Gzip` can only be used when `Pipe` mode is specified
as the `S3InputMode`. In `Pipe` mode, Amazon SageMaker streams input data from
the source directly to your container without using the EBS volume.

_Required_: No

_Type_: String

_Allowed values_: `None | Gzip`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3DataDistributionType`

Whether to distribute the data from Amazon S3 to all processing instances with
`FullyReplicated`, or whether the data from Amazon S3 is sharded by Amazon S3 key,
downloading one shard of data to each processing instance.

_Required_: No

_Type_: String

_Allowed values_: `FullyReplicated | ShardedByS3Key`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3DataType`

Whether you use an `S3Prefix` or a `ManifestFile` for the data
type. If you choose `S3Prefix`, `S3Uri` identifies a key name
prefix. Amazon SageMaker uses all objects with the specified key name prefix for the processing job.
If you choose `ManifestFile`, `S3Uri` identifies an object that is
a manifest file containing a list of object keys that you want Amazon SageMaker to use for the
processing job.

_Required_: Yes

_Type_: String

_Allowed values_: `ManifestFile | S3Prefix`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3InputMode`

Whether to use `File` or `Pipe` input mode. In File mode, Amazon SageMaker
copies the data from the input source onto the local ML storage volume before starting
your processing container. This is the most commonly used input mode. In
`Pipe` mode, Amazon SageMaker streams input data from the source directly to your
processing container into named pipes without using the ML storage volume.

_Required_: No

_Type_: String

_Allowed values_: `File | Pipe`

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

RedshiftDatasetDefinition

S3Output

All content copied from https://docs.aws.amazon.com/.
