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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/StartLiveTailResponseStream)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/StartLiveTailResponseStream)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/StartLiveTailResponseStream)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SplitStringEntry

SubscriptionFilter
