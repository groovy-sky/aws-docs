# DataSourceSyncJobMetrics

Maps a batch delete document request to a specific Amazon Q Business data source connector
sync job.

## Contents

**documentsAdded**

The current count of documents added from the data source during the data source
sync.

Type: String

Pattern: `(([1-9][0-9]*)|0)`

Required: No

**documentsDeleted**

The current count of documents deleted from the data source during the data source
sync.

Type: String

Pattern: `(([1-9][0-9]*)|0)`

Required: No

**documentsFailed**

The current count of documents that failed to sync from the data source during the
data source sync.

Type: String

Pattern: `(([1-9][0-9]*)|0)`

Required: No

**documentsModified**

The current count of documents modified in the data source during the data source
sync.

Type: String

Pattern: `(([1-9][0-9]*)|0)`

Required: No

**documentsScanned**

The current count of documents crawled by the ongoing sync job in the data
source.

Type: String

Pattern: `(([1-9][0-9]*)|0)`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/DataSourceSyncJobMetrics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/DataSourceSyncJobMetrics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/DataSourceSyncJobMetrics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataSourceSyncJob

DataSourceVpcConfiguration
