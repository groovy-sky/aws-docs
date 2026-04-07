# QueryRuntimeStatisticsTimeline

Timeline statistics such as query queue time, planning time, execution time, service
processing time, and total execution time.

## Contents

**EngineExecutionTimeInMillis**

The number of milliseconds that the query took to execute.

Type: Long

Required: No

**QueryPlanningTimeInMillis**

The number of milliseconds that Athena took to plan the query processing
flow. This includes the time spent retrieving table partitions from the data source.
Note that because the query engine performs the query planning, query planning time is a
subset of engine processing time.

Type: Long

Required: No

**QueryQueueTimeInMillis**

The number of milliseconds that the query was in your query queue waiting for
resources. Note that if transient errors occur, Athena might automatically
add the query back to the queue.

Type: Long

Required: No

**ServicePreProcessingTimeInMillis**

The number of milliseconds that Athena spends on preprocessing before it
submits the query to the engine.

Type: Long

Required: No

**ServiceProcessingTimeInMillis**

The number of milliseconds that Athena took to finalize and publish the
query results after the query engine finished running the query.

Type: Long

Required: No

**TotalExecutionTimeInMillis**

The number of milliseconds that Athena took to run the query.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryRuntimeStatisticsTimeline)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryRuntimeStatisticsTimeline)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryRuntimeStatisticsTimeline)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryRuntimeStatisticsRows

QueryStage
