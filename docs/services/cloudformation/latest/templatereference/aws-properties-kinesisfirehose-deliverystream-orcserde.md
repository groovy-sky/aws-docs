This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream OrcSerDe

A serializer to use for converting data to the ORC format before storing it in Amazon
S3. For more information, see [Apache\
ORC](https://orc.apache.org/docs).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockSizeBytes" : Integer,
  "BloomFilterColumns" : [ String, ... ],
  "BloomFilterFalsePositiveProbability" : Number,
  "Compression" : String,
  "DictionaryKeyThreshold" : Number,
  "EnablePadding" : Boolean,
  "FormatVersion" : String,
  "PaddingTolerance" : Number,
  "RowIndexStride" : Integer,
  "StripeSizeBytes" : Integer
}

```

### YAML

```yaml

  BlockSizeBytes: Integer
  BloomFilterColumns:
    - String
  BloomFilterFalsePositiveProbability: Number
  Compression: String
  DictionaryKeyThreshold: Number
  EnablePadding: Boolean
  FormatVersion: String
  PaddingTolerance: Number
  RowIndexStride: Integer
  StripeSizeBytes: Integer

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

`BloomFilterColumns`

The column names for which you want Firehose to create bloom filters. The
default is `null`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BloomFilterFalsePositiveProbability`

The Bloom filter false positive probability (FPP). The lower the FPP, the bigger the
Bloom filter. The default value is 0.05, the minimum is 0, and the maximum is 1.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compression`

The compression code to use over data blocks. The default is `SNAPPY`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | ZLIB | SNAPPY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DictionaryKeyThreshold`

Represents the fraction of the total number of non-null rows. To turn off dictionary
encoding, set this fraction to a number that is less than the number of distinct keys in a
dictionary. To always use dictionary encoding, set this threshold to 1.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePadding`

Set this to `true` to indicate that you want stripes to be padded to the HDFS
block boundaries. This is useful if you intend to copy the data from Amazon S3 to HDFS
before querying. The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormatVersion`

The version of the file to write. The possible values are `V0_11` and
`V0_12`. The default is `V0_12`.

_Required_: No

_Type_: String

_Allowed values_: `V0_11 | V0_12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PaddingTolerance`

A number between 0 and 1 that defines the tolerance for block padding as a decimal
fraction of stripe size. The default value is 0.05, which means 5 percent of stripe
size.

For the default values of 64 MiB ORC stripes and 256 MiB HDFS blocks, the default block
padding tolerance of 5 percent reserves a maximum of 3.2 MiB for padding within the 256 MiB
block. In such a case, if the available size within the block is more than 3.2 MiB, a new,
smaller stripe is inserted to fit within that space. This ensures that no stripe crosses
block boundaries and causes remote reads within a node-local task.

Kinesis Data Firehose ignores this parameter when `EnablePadding` is
`false`.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RowIndexStride`

The number of rows between index entries. The default is 10,000 and the minimum is
1,000.

_Required_: No

_Type_: Integer

_Minimum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StripeSizeBytes`

The number of bytes in each stripe. The default is 64 MiB and the minimum is 8
MiB.

_Required_: No

_Type_: Integer

_Minimum_: `8388608`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenXJsonSerDe

OutputFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
