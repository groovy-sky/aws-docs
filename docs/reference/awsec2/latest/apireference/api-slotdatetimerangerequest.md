# SlotDateTimeRangeRequest

Describes the time period for a Scheduled Instance to start its first schedule. The time period must span less than one day.

## Contents

**EarliestTime**

The earliest date and time, in UTC, for the Scheduled Instance to start.

Type: Timestamp

Required: Yes

**LatestTime**

The latest date and time, in UTC, for the Scheduled Instance to start. This value must be later than or equal to the earliest date and at most three months in the future.

Type: Timestamp

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/slotdatetimerangerequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/slotdatetimerangerequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/slotdatetimerangerequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceTypeDetail

SlotStartTimeRangeRequest
