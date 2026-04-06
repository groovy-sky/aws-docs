# ExecutionDetails

Describes the details of the flow run, including the timestamp, status, and message.

## Contents

**mostRecentExecutionMessage**

Describes the details of the most recent flow run.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `[\w!@#\-.?,\s]*`

Required: No

**mostRecentExecutionStatus**

Specifies the status of the most recent flow run.

Type: String

Valid Values: `InProgress | Successful | Error | CancelStarted | Canceled`

Required: No

**mostRecentExecutionTime**

Specifies the time of the most recent flow run.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ExecutionDetails)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ExecutionDetails)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ExecutionDetails)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridgeMetadata

ExecutionRecord
