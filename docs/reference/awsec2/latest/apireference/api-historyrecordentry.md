# HistoryRecordEntry

Describes an event in the history of an EC2 Fleet.

## Contents

**eventInformation**

Information about the event.

Type: [EventInformation](api-eventinformation.md) object

Required: No

**eventType**

The event type.

Type: String

Valid Values: `instance-change | fleet-change | service-error`

Required: No

**timestamp**

The date and time of the event, in UTC format (for example,
_YYYY_- _MM_- _DD_ T _HH_: _MM_: _SS_ Z).

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/historyrecordentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/historyrecordentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/historyrecordentry.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HistoryRecord

Host
