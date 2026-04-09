# ConfigurationTemplate

A structure containing information about the deafult settings and available settings that
you can use to configure a [delivery](api-delivery.md) or a
[delivery\
destination](api-deliverydestination.md).

## Contents

**allowedActionForAllowVendedLogsDeliveryForResource**

The action permissions that a caller needs to have to be able to successfully create a
delivery source on the desired resource type when calling [PutDeliverySource](api-putdeliverysource.md).

Type: String

Required: No

**allowedFieldDelimiters**

The valid values that a caller can use as field delimiters when calling [CreateDelivery](api-createdelivery.md) or [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md) on a delivery that delivers in `Plain`,
`W3C`, or `Raw` format.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 5.

Required: No

**allowedFields**

The allowed fields that a caller can use in the `recordFields` parameter of a
[CreateDelivery](api-createdelivery.md) or [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md) operation.

Type: Array of [RecordField](api-recordfield.md) objects

Required: No

**allowedOutputFormats**

The list of delivery destination output formats that are supported by this log
source.

Type: Array of strings

Valid Values: `json | plain | w3c | raw | parquet`

Required: No

**allowedSuffixPathFields**

The list of variable fields that can be used in the suffix path of a delivery that
delivers to an S3 bucket.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 128 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**defaultDeliveryConfigValues**

A mapping that displays the default value of each property within a delivery's
configuration, if it is not specified in the request.

Type: [ConfigurationTemplateDeliveryConfigValues](api-configurationtemplatedeliveryconfigvalues.md) object

Required: No

**deliveryDestinationType**

A string specifying which destination type this configuration template applies to.

Type: String

Valid Values: `S3 | CWL | FH | XRAY`

Required: No

**logType**

A string specifying which log type this configuration template applies to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w]*`

Required: No

**resourceType**

A string specifying which resource type this configuration template applies to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w-_]*`

Required: No

**service**

A string specifying which service this configuration template applies to. For more
information about supported services see [Enable logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\w_-]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/configurationtemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/configurationtemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/configurationtemplate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnomalyDetector

ConfigurationTemplateDeliveryConfigValues

All content copied from https://docs.aws.amazon.com/.
