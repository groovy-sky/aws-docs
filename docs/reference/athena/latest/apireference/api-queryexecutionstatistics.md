# QueryExecutionStatistics

The amount of data scanned during the query execution and the amount of time that it
took to execute, and the type of statement that was run.

## Contents

**DataManifestLocation**

The location and file name of a data manifest file. The manifest file is saved to the
Athena query results location in Amazon S3. The manifest file
tracks files that the query wrote to Amazon S3. If the query fails, the manifest
file also tracks files that the query intended to write. The manifest is useful for
identifying orphaned files resulting from a failed query. For more information, see
[Working with Query\
Results, Output Files, and Query History](../../../../services/athena/latest/ug/querying.md) in the _Amazon Athena User Guide_.

Type: String

Required: No

**DataScannedInBytes**

The number of bytes in the data that was queried.

Type: Long

Required: No

**DpuCount**

The number of Data Processing Units (DPUs) that Athena used to run the query.

Type: Double

Required: No

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

**ResultReuseInformation**

Contains information about whether previous query results were reused for the
query.

Type: [ResultReuseInformation](https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultReuseInformation.html) object

Required: No

**ServicePreProcessingTimeInMillis**

The number of milliseconds that Athena took to preprocess the query before
submitting the query to the query engine.

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryExecutionStatistics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryExecutionStatistics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryExecutionStatistics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryExecutionContext

QueryExecutionStatus
