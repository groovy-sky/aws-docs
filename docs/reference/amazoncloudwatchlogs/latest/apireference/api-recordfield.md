---
title: "RecordField"
---

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/RecordField)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/RecordField)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/RecordField)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryStatistics

RejectedEntityInfo

All content copied from https://docs.aws.amazon.com/.
