# InputLogEvent

Represents a log event, which is a record of activity that was recorded by the
application or resource being monitored.

## Contents

**message**

The raw event message. Each log event can be no larger than 1 MB.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**timestamp**

The time the event occurred, expressed as the number of milliseconds after `Jan 1,
        1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/inputlogevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/inputlogevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/inputlogevent.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IndexPolicy

IntegrationDetails
