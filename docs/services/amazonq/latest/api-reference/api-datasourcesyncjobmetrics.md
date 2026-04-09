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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/datasourcesyncjobmetrics.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/datasourcesyncjobmetrics.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/datasourcesyncjobmetrics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceSyncJob

DataSourceVpcConfiguration

All content copied from https://docs.aws.amazon.com/.
