This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage TransformInput

Describes the input source of a transform job and the way the transform job consumes
it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompressionType" : String,
  "ContentType" : String,
  "DataSource" : DataSource,
  "SplitType" : String
}

```

### YAML

```yaml

  CompressionType: String
  ContentType: String
  DataSource:
    DataSource
  SplitType: String

```

## Properties

`CompressionType`

If your transform data
is
compressed, specify the compression type. Amazon SageMaker automatically
decompresses the data for the transform job accordingly. The default value is
`None`.

_Required_: No

_Type_: String

_Allowed values_: `None | Gzip`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContentType`

The multipurpose internet mail extension
(MIME)
type of the data. Amazon SageMaker uses the MIME type with each http call to
transfer data to the transform job.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSource`

Describes the location of
the
channel data, which is, the S3 location of the input data that the
model can consume.

_Required_: Yes

_Type_: [DataSource](aws-properties-sagemaker-modelpackage-datasource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SplitType`

The method to use to split the transform job's data files into smaller batches.
Splitting is necessary when the total size of each object is too large to fit in a
single request. You can also use data splitting to improve performance by processing
multiple concurrent mini-batches. The default value for `SplitType` is
`None`, which indicates that input data files are not split, and request
payloads contain the entire contents of an input object. Set the value of this parameter
to `Line` to split records on a newline character boundary.
`SplitType` also supports a number of record-oriented binary data
formats. Currently, the supported record formats are:

- RecordIO

- TFRecord

When splitting is enabled, the size of a mini-batch depends on the values of the
`BatchStrategy` and `MaxPayloadInMB` parameters. When the
value of `BatchStrategy` is `MultiRecord`, Amazon SageMaker sends the maximum
number of records in each request, up to the `MaxPayloadInMB` limit. If the
value of `BatchStrategy` is `SingleRecord`, Amazon SageMaker sends individual
records in each request.

###### Note

Some data formats represent a record as a binary payload wrapped with extra
padding bytes. When splitting is applied to a binary data format, padding is removed
if the value of `BatchStrategy` is set to `SingleRecord`.
Padding is not removed if the value of `BatchStrategy` is set to
`MultiRecord`.

For more information about `RecordIO`, see [Create a Dataset Using\
RecordIO](https://mxnet.apache.org/api/faq/recordio) in the MXNet documentation. For more information about
`TFRecord`, see [Consuming TFRecord data](https://www.tensorflow.org/guide/data) in the TensorFlow documentation.

_Required_: No

_Type_: String

_Allowed values_: `None | TFRecord | Line | RecordIO`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TransformJobDefinition

All content copied from https://docs.aws.amazon.com/.
