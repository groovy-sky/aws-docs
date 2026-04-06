# UpsolverS3OutputFormatConfig

The configuration that determines how Amazon AppFlow formats the flow output data
when Upsolver is used as the destination.

## Contents

**prefixConfig**

Specifies elements that Amazon AppFlow includes in the file and folder names in the flow
destination.

Type: [PrefixConfig](api-prefixconfig.md) object

Required: Yes

**aggregationConfig**

The aggregation settings that you can use to customize the output format of your flow
data.

Type: [AggregationConfig](api-aggregationconfig.md) object

Required: No

**fileType**

Indicates the file type that Amazon AppFlow places in the Upsolver Amazon S3
bucket.

Type: String

Valid Values: `CSV | JSON | PARQUET`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/UpsolverS3OutputFormatConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/UpsolverS3OutputFormatConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/UpsolverS3OutputFormatConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpsolverMetadata

VeevaConnectorProfileCredentials
