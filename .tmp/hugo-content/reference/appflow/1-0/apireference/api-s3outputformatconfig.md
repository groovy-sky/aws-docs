# S3OutputFormatConfig

The configuration that determines how Amazon AppFlow should format the flow output
data when Amazon S3 is used as the destination.

## Contents

**aggregationConfig**

The aggregation settings that you can use to customize the output format of your flow
data.

Type: [AggregationConfig](api-aggregationconfig.md) object

Required: No

**fileType**

Indicates the file type that Amazon AppFlow places in the Amazon S3 bucket.

Type: String

Valid Values: `CSV | JSON | PARQUET`

Required: No

**prefixConfig**

Determines the prefix that Amazon AppFlow applies to the folder name in the Amazon S3 bucket. You can name folders according to the flow frequency and date.

Type: [PrefixConfig](api-prefixconfig.md) object

Required: No

**preserveSourceDataTyping**

If your file output format is Parquet, use this parameter to set whether Amazon AppFlow preserves the data types in your source data when it writes the output to Amazon S3.

- `true`: Amazon AppFlow preserves the data types when it writes to
Amazon S3. For example, an integer or `1` in your source data is
still an integer in your output.

- `false`: Amazon AppFlow converts all of the source data into strings
when it writes to Amazon S3. For example, an integer of `1` in your
source data becomes the string `"1"` in the output.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/s3outputformatconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/s3outputformatconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/s3outputformatconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Metadata

S3SourceProperties

All content copied from https://docs.aws.amazon.com/.
