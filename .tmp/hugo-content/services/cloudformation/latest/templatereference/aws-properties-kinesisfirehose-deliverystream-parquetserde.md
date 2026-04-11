This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream ParquetSerDe

A serializer to use for converting data to the Parquet format before storing it in
Amazon S3. For more information, see [Apache Parquet](https://parquet.apache.org/docs).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockSizeBytes" : Integer,
  "Compression" : String,
  "EnableDictionaryCompression" : Boolean,
  "MaxPaddingBytes" : Integer,
  "PageSizeBytes" : Integer,
  "WriterVersion" : String
}

```

### YAML

```yaml

  BlockSizeBytes: Integer
  Compression: String
  EnableDictionaryCompression: Boolean
  MaxPaddingBytes: Integer
  PageSizeBytes: Integer
  WriterVersion: String

```

## Properties

`BlockSizeBytes`

The Hadoop Distributed File System (HDFS) block size. This is useful if you intend to
copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the
minimum is 64 MiB. Firehose uses this value for padding calculations.

_Required_: No

_Type_: Integer

_Minimum_: `67108864`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compression`

The compression code to use over data blocks. The possible values are
`UNCOMPRESSED`, `SNAPPY`, and `GZIP`, with the default
being `SNAPPY`. Use `SNAPPY` for higher decompression speed. Use
`GZIP` if the compression ratio is more important than speed.

_Required_: No

_Type_: String

_Allowed values_: `UNCOMPRESSED | GZIP | SNAPPY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableDictionaryCompression`

Indicates whether to enable dictionary compression.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxPaddingBytes`

The maximum amount of padding to apply. This is useful if you intend to copy the data
from Amazon S3 to HDFS before querying. The default is 0.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PageSizeBytes`

The Parquet page size. Column chunks are divided into pages. A page is conceptually an
indivisible unit (in terms of compression and encoding). The minimum value is 64 KiB and
the default is 1 MiB.

_Required_: No

_Type_: Integer

_Minimum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriterVersion`

Indicates the version of row format to output. The possible values are `V1`
and `V2`. The default is `V1`.

_Required_: No

_Type_: String

_Allowed values_: `V1 | V2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputFormatConfiguration

PartitionField

All content copied from https://docs.aws.amazon.com/.
