# LiveTailSessionMetadata

This object contains the metadata for one `LiveTailSessionUpdate` structure. It
indicates whether that update includes only a sample of 500 log events out of a larger number
of ingested log events, or if it contains all of the matching log events ingested during that
second of time.

## Contents

**sampled**

If this is `true`, then more than 500 log events matched the request for this
update, and the `sessionResults` includes a sample of 500 of those events.

If this is `false`, then 500 or fewer log events matched the request for this
update, so no sampling was necessary. In this case, the `sessionResults` array
includes all log events that matched your request during this time.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/livetailsessionmetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/livetailsessionmetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/livetailsessionmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LiveTailSessionLogEvent

LiveTailSessionStart

All content copied from https://docs.aws.amazon.com/.
