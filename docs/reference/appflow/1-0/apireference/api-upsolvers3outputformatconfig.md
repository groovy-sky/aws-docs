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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/upsolvers3outputformatconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/upsolvers3outputformatconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/upsolvers3outputformatconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpsolverMetadata

VeevaConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
