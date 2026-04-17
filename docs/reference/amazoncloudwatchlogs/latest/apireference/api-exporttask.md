---
title: "ExportTask"
---

# ExportTask

Represents an export task.

## Contents

**destination**

The name of the S3 bucket to which the log data was exported.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

**destinationPrefix**

The prefix that was used as the start of Amazon S3 key for every object
exported.

Type: String

Required: No

**executionInfo**

Execution information about the export task.

Type: [ExportTaskExecutionInfo](api-exporttaskexecutioninfo.md) object

Required: No

**from**

The start time, expressed as the number of milliseconds after `Jan 1, 1970
        00:00:00 UTC`. Events with a timestamp before this time are not exported.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupName**

The name of the log group from which logs data was exported.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**status**

The status of the export task.

Type: [ExportTaskStatus](api-exporttaskstatus.md) object

Required: No

**taskId**

The ID of the export task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

**taskName**

The name of the export task.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Required: No

**to**

The end time, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00
        UTC`. Events with a timestamp later than this time are not exported.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/ExportTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/ExportTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/ExportTask)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Entity

ExportTaskExecutionInfo

All content copied from https://docs.aws.amazon.com/.
