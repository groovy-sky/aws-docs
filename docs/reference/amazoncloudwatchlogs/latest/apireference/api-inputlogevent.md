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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/InputLogEvent)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/InputLogEvent)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/InputLogEvent)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IndexPolicy

IntegrationDetails
