# ExecutionResult

Specifies the end result of the flow run.

## Contents

**bytesProcessed**

The total number of bytes processed by the flow run.

Type: Long

Required: No

**bytesWritten**

The total number of bytes written as a result of the flow run.

Type: Long

Required: No

**errorInfo**

Provides any error message information related to the flow run.

Type: [ErrorInfo](api-errorinfo.md) object

Required: No

**maxPageSize**

The maximum number of records that Amazon AppFlow receives in each page of the
response from your SAP application.

Type: Long

Required: No

**numParallelProcesses**

The number of processes that Amazon AppFlow ran at the same time when it retrieved
your data.

Type: Long

Required: No

**recordsProcessed**

The number of records processed in the flow run.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/executionresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/executionresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/executionresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionRecord

FieldTypeDetails

All content copied from https://docs.aws.amazon.com/.
