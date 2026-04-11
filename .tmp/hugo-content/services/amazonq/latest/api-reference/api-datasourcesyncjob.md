# DataSourceSyncJob

Provides information about an Amazon Q Business data source connector synchronization
job.

## Contents

**dataSourceErrorCode**

If the reason that the synchronization failed is due to an error with the underlying
data source, this field contains a code that identifies the error.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**endTime**

The Unix timestamp when the synchronization job completed.

Type: Timestamp

Required: No

**error**

If the `Status` field is set to `FAILED`, the
`ErrorCode` field indicates the reason the synchronization failed.

Type: [ErrorDetail](api-errordetail.md) object

Required: No

**executionId**

The identifier of a data source synchronization job.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**metrics**

Maps a batch delete document request to a specific data source sync job. This is
optional and should only be supplied when documents are deleted by a data source
connector.

Type: [DataSourceSyncJobMetrics](api-datasourcesyncjobmetrics.md) object

Required: No

**startTime**

The Unix time stamp when the data source synchronization job started.

Type: Timestamp

Required: No

**status**

The status of the synchronization job. When the `Status` field is set to
`SUCCEEDED`, the synchronization job is done. If the status code is
`FAILED`, the `ErrorCode` and `ErrorMessage` fields
give you the reason for the failure.

Type: String

Valid Values: `FAILED | SUCCEEDED | SYNCING | INCOMPLETE | STOPPING | ABORTED | SYNCING_INDEXING`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/datasourcesyncjob.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/datasourcesyncjob.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/datasourcesyncjob.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSource

DataSourceSyncJobMetrics

All content copied from https://docs.aws.amazon.com/.
