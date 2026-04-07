# LiveTailSessionUpdate

This object contains the log events and metadata for a Live Tail session.

## Contents

**sessionMetadata**

This object contains the session metadata for a Live Tail session.

Type: [LiveTailSessionMetadata](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionMetadata.html) object

Required: No

**sessionResults**

An array, where each member of the array includes the information for one log event in the
Live Tail session.

A `sessionResults` array can include as many as 500 log events. If the number
of log events matching the request exceeds 500 per second, the log events are sampled down to
500 log events to be included in each `sessionUpdate` structure.

Type: Array of [LiveTailSessionLogEvent](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LiveTailSessionLogEvent.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/LiveTailSessionUpdate)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/LiveTailSessionUpdate)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/LiveTailSessionUpdate)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LiveTailSessionStart

LogEvent
