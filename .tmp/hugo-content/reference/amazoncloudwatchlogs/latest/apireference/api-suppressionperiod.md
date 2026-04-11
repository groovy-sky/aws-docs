# SuppressionPeriod

If you are suppressing an anomaly temporariliy, this structure defines how long the
suppression period is to be.

## Contents

**suppressionUnit**

Specifies whether the value of `value` is in seconds, minutes, or hours.

Type: String

Valid Values: `SECONDS | MINUTES | HOURS`

Required: No

**value**

Specifies the number of seconds, minutes or hours to suppress this anomaly. There is no
maximum.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/suppressionperiod.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/suppressionperiod.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/suppressionperiod.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubstituteStringEntry

TransformedLogRecord

All content copied from https://docs.aws.amazon.com/.
