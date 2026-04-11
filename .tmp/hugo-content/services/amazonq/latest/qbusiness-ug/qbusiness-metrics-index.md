# Amazon Q Business index metrics

The following table shows the [Index](concepts-terms.md#index) metrics that
Amazon Q Business sends to CloudWatch in real time.

Metric nameUnitDescription

`DocumentCount`

Count

The number of documents. This metric is published every 15 minutes.

Valid dimensions: `ApplicationId`, `IndexId`

`DocumentsIndexed`

Count

The number of documents that were indexed.

Valid dimensions: `ApplicationId`, `IndexId`,
`DataSourceId`

`DocumentsFailedToIndex`

Count

The number of documents that failed to index.

Valid dimensions: `ApplicationId`, `IndexId`,
`DataSourceId`

`DocumentsFailedToIndexDueToCDE`

Count

The number of documents that failed to index because of custom document
enrichment.

Valid dimensions: `ApplicationId`, `IndexId`,
`DataSourceId`

`ExtractedTextSize`

MB

Size of the extracted text

Valid dimensions: `ApplicationId`, `IndexId`

`MonthlyDataSyncDuration`Count

Duration of data synchronization operations over a monthly period.

Valid dimensions: `ApplicationId`, `IndexId`,
`DataSourceId`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Q Business API operation metrics

Amazon Q Apps metrics

All content copied from https://docs.aws.amazon.com/.
