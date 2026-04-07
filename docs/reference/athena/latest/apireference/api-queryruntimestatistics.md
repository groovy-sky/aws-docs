# QueryRuntimeStatistics

The query execution timeline, statistics on input and output rows and bytes, and the
different query stages that form the query execution plan.

## Contents

**OutputStage**

Stage statistics such as input and output rows and bytes, execution time, and stage
state. This information also includes substages and the query stage plan.

Type: [QueryStage](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryStage.html) object

Required: No

**Rows**

Statistics such as input rows and bytes read by the query, rows and bytes output by
the query, and the number of rows written by the query.

Type: [QueryRuntimeStatisticsRows](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryRuntimeStatisticsRows.html) object

Required: No

**Timeline**

Timeline statistics such as query queue time, planning time, execution time, service
processing time, and total execution time.

Type: [QueryRuntimeStatisticsTimeline](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryRuntimeStatisticsTimeline.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryRuntimeStatistics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryRuntimeStatistics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryRuntimeStatistics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryResultsS3AccessGrantsConfiguration

QueryRuntimeStatisticsRows
