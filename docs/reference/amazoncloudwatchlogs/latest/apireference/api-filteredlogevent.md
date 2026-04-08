# FilteredLogEvent

Represents a matched event.

## Contents

**eventId**

The ID of the event.

Type: String

Required: No

**ingestionTime**

The time the event was ingested, expressed as the number of milliseconds after
`Jan 1, 1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logStreamName**

The name of the log stream to which this event belongs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**message**

The data contained in the log event.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**timestamp**

The time the event occurred, expressed as the number of milliseconds after `Jan 1,
        1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/filteredlogevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/filteredlogevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/filteredlogevent.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FieldsData

GetLogObjectResponseStream
