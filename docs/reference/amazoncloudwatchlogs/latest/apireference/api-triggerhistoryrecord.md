---
title: "TriggerHistoryRecord"
---

# TriggerHistoryRecord

A record of a scheduled query execution, including execution status, timestamp, and
destination processing results.

## Contents

**destinations**

Information about destination processing for this query execution.

Type: Array of [ScheduledQueryDestination](api-scheduledquerydestination.md) objects

Required: No

**errorMessage**

Error message if the query execution failed.

Type: String

Required: No

**executionStatus**

The execution status of the scheduled query run.

Type: String

Valid Values: `Running | InvalidQuery | Complete | Failed | Timeout`

Required: No

**queryId**

The unique identifier for this query execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**triggeredTimestamp**

The timestamp when the scheduled query execution was triggered.

Type: Long

Valid Range: Minimum value of 0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/TriggerHistoryRecord)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/TriggerHistoryRecord)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/TriggerHistoryRecord)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransformedLogRecord

TrimString

All content copied from https://docs.aws.amazon.com/.
