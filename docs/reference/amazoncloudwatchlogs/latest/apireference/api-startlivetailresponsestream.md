# StartLiveTailResponseStream

This object includes the stream returned by your [StartLiveTail](api-startlivetail.md)
request.

## Contents

**sessionStart**

This object contains information about this Live Tail session, including the log groups
included and the log stream filters, if any.

Type: [LiveTailSessionStart](api-livetailsessionstart.md) object

Required: No

**SessionStreamingException**

This exception is returned if an unknown error occurs.

Type: Exception

HTTP Status Code:

Required: No

**SessionTimeoutException**

This exception is returned in the stream when the Live Tail session times out. Live Tail
sessions time out after three hours.

Type: Exception

HTTP Status Code:

Required: No

**sessionUpdate**

This object contains the log events and session metadata.

Type: [LiveTailSessionUpdate](api-livetailsessionupdate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/startlivetailresponsestream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/startlivetailresponsestream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/startlivetailresponsestream.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SplitStringEntry

SubscriptionFilter
