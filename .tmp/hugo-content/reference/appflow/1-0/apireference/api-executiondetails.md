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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/executiondetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/executiondetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/executiondetails.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventBridgeMetadata

ExecutionRecord

All content copied from https://docs.aws.amazon.com/.
