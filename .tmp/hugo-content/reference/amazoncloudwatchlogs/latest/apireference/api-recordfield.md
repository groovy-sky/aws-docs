# RecordField

A structure that represents a valid record field header and whether it is
mandatory.

## Contents

**mandatory**

If this is `true`, the record field must be present in the
`recordFields` parameter provided to a [CreateDelivery](api-createdelivery.md) or [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md) operation.

Type: Boolean

Required: No

**name**

The name to use when specifying this record field in a [CreateDelivery](api-createdelivery.md) or [UpdateDeliveryConfiguration](api-updatedeliveryconfiguration.md) operation.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/recordfield.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/recordfield.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/recordfield.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryStatistics

RejectedEntityInfo

All content copied from https://docs.aws.amazon.com/.
