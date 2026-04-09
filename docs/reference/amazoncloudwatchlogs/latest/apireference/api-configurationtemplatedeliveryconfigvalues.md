# ConfigurationTemplateDeliveryConfigValues

This structure contains the default values that are used for each configuration parameter
when you use [CreateDelivery](api-createdelivery.md) to create a deliver under the current service type, resource type,
and log type.

## Contents

**fieldDelimiter**

The default field delimiter that is used in a [CreateDelivery](api-createdelivery.md) operation when the field delimiter is not specified in that
operation. The field delimiter is used only when the final output delivery is in
`Plain`, `W3C`, or `Raw` format.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 5.

Required: No

**recordFields**

The default record fields that will be delivered when a list of record fields is not
provided in a [CreateDelivery](api-createdelivery.md) operation.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 128 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**s3DeliveryConfiguration**

The delivery parameters that are used when you create a delivery to a delivery destination
that is an S3 Bucket.

Type: [S3DeliveryConfiguration](api-s3deliveryconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/configurationtemplatedeliveryconfigvalues.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/configurationtemplatedeliveryconfigvalues.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/configurationtemplatedeliveryconfigvalues.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationTemplate

CopyValue

All content copied from https://docs.aws.amazon.com/.
