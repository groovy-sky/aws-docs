This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow S3OutputFormatConfig

The configuration that determines how Amazon AppFlow should format the flow output
data when Amazon S3 is used as the destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregationConfig" : AggregationConfig,
  "FileType" : String,
  "PrefixConfig" : PrefixConfig,
  "PreserveSourceDataTyping" : Boolean
}

```

### YAML

```yaml

  AggregationConfig:
    AggregationConfig
  FileType: String
  PrefixConfig:
    PrefixConfig
  PreserveSourceDataTyping: Boolean

```

## Properties

`AggregationConfig`

The aggregation settings that you can use to customize the output format of your flow
data.

_Required_: No

_Type_: [AggregationConfig](aws-properties-appflow-flow-aggregationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileType`

Indicates the file type that Amazon AppFlow places in the Amazon S3 bucket.

_Required_: No

_Type_: String

_Allowed values_: `CSV | JSON | PARQUET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixConfig`

Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket. You can name folders according to the flow frequency and date.

_Required_: No

_Type_: [PrefixConfig](aws-properties-appflow-flow-prefixconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveSourceDataTyping`

If your file output format is Parquet, use this parameter to set whether Amazon AppFlow preserves the data types in your source data when it writes the output to Amazon S3.

- `true`: Amazon AppFlow preserves the data types when it writes to
Amazon S3. For example, an integer or `1` in your source data is
still an integer in your output.

- `false`: Amazon AppFlow converts all of the source data into strings
when it writes to Amazon S3. For example, an integer of `1` in your
source data becomes the string `"1"` in the output.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3OutputFormatConfig](../../../../reference/appflow/1-0/apireference/api-s3outputformatconfig.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3InputFormatConfig

S3SourceProperties

All content copied from https://docs.aws.amazon.com/.
