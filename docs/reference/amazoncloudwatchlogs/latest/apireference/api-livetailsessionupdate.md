# LiveTailSessionUpdate

This object contains the log events and metadata for a Live Tail session.

## Contents

**sessionMetadata**

This object contains the session metadata for a Live Tail session.

Type: [LiveTailSessionMetadata](api-livetailsessionmetadata.md) object

Required: No

**sessionResults**

An array, where each member of the array includes the information for one log event in the
Live Tail session.

A `sessionResults` array can include as many as 500 log events. If the number
of log events matching the request exceeds 500 per second, the log events are sampled down to
500 log events to be included in each `sessionUpdate` structure.

Type: Array of [LiveTailSessionLogEvent](api-livetailsessionlogevent.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/livetailsessionupdate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/livetailsessionupdate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/livetailsessionupdate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LiveTailSessionStart

LogEvent
