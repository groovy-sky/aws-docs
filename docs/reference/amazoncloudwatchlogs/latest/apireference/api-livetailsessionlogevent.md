# LiveTailSessionLogEvent

This object contains the information for one log event returned in a Live Tail
stream.

## Contents

**ingestionTime**

The timestamp specifying when this log event was ingested into the log group.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupIdentifier**

The name or ARN of the log group that ingested this log event.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**logStreamName**

The name of the log stream that ingested this log event.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**message**

The log event message text.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**timestamp**

The timestamp specifying when this log event was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/livetailsessionlogevent.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/livetailsessionlogevent.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/livetailsessionlogevent.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListToMap

LiveTailSessionMetadata

All content copied from https://docs.aws.amazon.com/.
