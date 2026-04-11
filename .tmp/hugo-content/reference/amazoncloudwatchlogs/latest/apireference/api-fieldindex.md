# FieldIndex

This structure describes one log event field that is used as an index in at least one
index policy in this account.

## Contents

**fieldIndexName**

The string that this field index matches.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**firstEventTime**

The time and date of the earliest log event that matches this field index, after the index
policy that contains it was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastEventTime**

The time and date of the most recent log event that matches this field index.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastScanTime**

The most recent time that CloudWatch Logs scanned ingested log events to search for
this field index to improve the speed of future CloudWatch Logs Insights queries that
search for this field index.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupIdentifier**

If this field index appears in an index policy that applies only to a single log group,
the ARN of that log group is displayed here.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**type**

The type of index. Specify `FACET` for facet-based indexing or
`FIELD_INDEX` for field-based indexing. This determines how the field is indexed
and can be queried.

Type: String

Valid Values: `FACET | FIELD_INDEX`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/fieldindex.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/fieldindex.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/fieldindex.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportTaskStatus

FieldsData

All content copied from https://docs.aws.amazon.com/.
