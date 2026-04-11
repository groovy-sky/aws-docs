# QueryStatistics

Contains the number of log events scanned by the query, the number of log events that
matched the query criteria, and the total number of bytes in the log events that were
scanned.

If the query involved log groups that have field index policies, the estimated number of
skipped log events and the total bytes of those skipped log events are included. Using field
indexes to skip log events in queries reduces scan volume and improves performance. For more
information, see [Create field indexes\
to improve query performance and reduce scan volume](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-field-indexing.md).

## Contents

**bytesScanned**

The total number of bytes in the log events scanned during the query.

Type: Double

Required: No

**estimatedBytesSkipped**

An estimate of the number of bytes in the log events that were skipped when processing
this query, because the query contained an indexed field. Skipping these entries lowers query
costs and improves the query performance time. For more information about field indexes, see
[PutIndexPolicy](api-putindexpolicy.md).

Type: Double

Required: No

**estimatedRecordsSkipped**

An estimate of the number of log events that were skipped when processing this query,
because the query contained an indexed field. Skipping these entries lowers query costs and
improves the query performance time. For more information about field indexes, see [PutIndexPolicy](api-putindexpolicy.md).

Type: Double

Required: No

**logGroupsScanned**

The number of log groups that were scanned by this query.

Type: Double

Required: No

**recordsMatched**

The number of log events that matched the query string.

Type: Double

Required: No

**recordsScanned**

The total number of log events scanned during the query.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/querystatistics.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/querystatistics.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/querystatistics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryParameter

RecordField

All content copied from https://docs.aws.amazon.com/.
