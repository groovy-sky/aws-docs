# AggregationConfig

The aggregation settings that you can use to customize the output format of your flow
data.

## Contents

**aggregationType**

Specifies whether Amazon AppFlow aggregates the flow records into a single file, or
leave them unaggregated.

Type: String

Valid Values: `None | SingleFile`

Required: No

**targetFileSize**

The desired file size, in MB, for each output file that Amazon AppFlow writes to the
flow destination. For each file, Amazon AppFlow attempts to achieve the size that you
specify. The actual file sizes might differ from this target based on the number and size of
the records that each file contains.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/AggregationConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/AggregationConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/AggregationConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Data Types

AmplitudeConnectorProfileCredentials
