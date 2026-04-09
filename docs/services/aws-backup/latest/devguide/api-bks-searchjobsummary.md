# SearchJobSummary

This is information pertaining to a search job.

## Contents

**CompletionTime**

This is the completion time of the search job.

Type: Timestamp

Required: No

**CreationTime**

This is the creation time of the search job.

Type: Timestamp

Required: No

**Name**

This is the name of the search job.

Type: String

Required: No

**SearchJobArn**

The unique string that identifies the Amazon Resource
Name (ARN) of the specified search job.

Type: String

Required: No

**SearchJobIdentifier**

The unique string that specifies the search job.

Type: String

Required: No

**SearchScopeSummary**

Returned summary of the specified search job scope,
including:

- TotalBackupsToScanCount, the number of
recovery points returned by the search.

- TotalItemsToScanCount, the number of
items returned by the search.

Type: [SearchScopeSummary](api-bks-searchscopesummary.md) object

Required: No

**Status**

This is the status of the search job.

Type: String

Valid Values: `RUNNING | COMPLETED | STOPPING | STOPPED | FAILED`

Required: No

**StatusMessage**

A status message will be returned for either a
earch job with a status of `ERRORED` or a status of
`COMPLETED` jobs with issues.

For example, a message may say that a search
contained recovery points unable to be scanned because
of a permissions issue.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/searchjobsummary.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/searchjobsummary.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/searchjobsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SearchJobBackupsResult

SearchScope

All content copied from https://docs.aws.amazon.com/.
