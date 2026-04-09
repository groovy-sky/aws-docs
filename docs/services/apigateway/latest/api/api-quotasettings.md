# QuotaSettings

Quotas configured for a usage plan.

## Contents

**limit**

The target maximum number of requests that can be made in a given time period.

Type: Integer

Required: No

**offset**

The number of requests subtracted from the given limit in the initial time period.

Type: Integer

Required: No

**period**

The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".

Type: String

Valid Values: `DAY | WEEK | MONTH`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/quotasettings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/quotasettings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/quotasettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PatchOperation

RequestValidator

All content copied from https://docs.aws.amazon.com/.
