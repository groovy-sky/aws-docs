# QueryRuntimeStatistics

The query execution timeline, statistics on input and output rows and bytes, and the
different query stages that form the query execution plan.

## Contents

**OutputStage**

Stage statistics such as input and output rows and bytes, execution time, and stage
state. This information also includes substages and the query stage plan.

Type: [QueryStage](api-querystage.md) object

Required: No

**Rows**

Statistics such as input rows and bytes read by the query, rows and bytes output by
the query, and the number of rows written by the query.

Type: [QueryRuntimeStatisticsRows](api-queryruntimestatisticsrows.md) object

Required: No

**Timeline**

Timeline statistics such as query queue time, planning time, execution time, service
processing time, and total execution time.

Type: [QueryRuntimeStatisticsTimeline](api-queryruntimestatisticstimeline.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/queryruntimestatistics.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/queryruntimestatistics.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/queryruntimestatistics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryResultsS3AccessGrantsConfiguration

QueryRuntimeStatisticsRows

All content copied from https://docs.aws.amazon.com/.
