# ExportInfo

Information regarding the export status of discovered data. The value is an array of
objects.

## Contents

**exportId**

A unique identifier used to query an export.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: Yes

**exportRequestTime**

The time that the data export was initiated.

Type: Timestamp

Required: Yes

**exportStatus**

The status of the data export job.

Type: String

Valid Values: `FAILED | SUCCEEDED | IN_PROGRESS`

Required: Yes

**statusMessage**

A status message provided for API callers.

Type: String

Required: Yes

**configurationsDownloadUrl**

A URL for an Amazon S3 bucket where you can review the exported data. The URL is
displayed only if the export succeeded.

Type: String

Required: No

**isTruncated**

If true, the export of agent information exceeded the size limit for a single export
and the exported data is incomplete for the requested time range. To address this, select a
smaller time range for the export by using `startDate` and
`endDate`.

Type: Boolean

Required: No

**requestedEndTime**

The `endTime` used in the `StartExportTask` request. If no
`endTime` was requested, this result does not appear in
`ExportInfo`.

Type: Timestamp

Required: No

**requestedStartTime**

The value of `startTime` parameter in the `StartExportTask`
request. If no `startTime` was requested, this result does not appear in
`ExportInfo`.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/exportinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/exportinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/exportinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportFilter

ExportPreferences

All content copied from https://docs.aws.amazon.com/.
