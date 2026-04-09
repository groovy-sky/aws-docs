# QueryInfo

Information about one CloudWatch Logs Insights query that matches the request in a
`DescribeQueries` operation.

## Contents

**bytesScanned**

The total number of bytes scanned by the query. This indicates the cost associated with the query.

Type: Double

Required: No

**createTime**

The date and time that this query was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupName**

The name of the log group scanned by this query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**queryDuration**

The duration in milliseconds that the query took to execute.

Type: Long

Required: No

**queryId**

The unique ID number of this query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**queryLanguage**

The query language used for this query. For more information about the query languages
that CloudWatch Logs supports, see [Supported query\
languages](../../../../services/amazoncloudwatch/latest/logs/cwl-analyzelogdata-languages.md).

Type: String

Valid Values: `CWLI | SQL | PPL`

Required: No

**queryString**

The query string used in this query.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10000.

Required: No

**status**

The status of this query. Possible values are `Cancelled`,
`Complete`, `Failed`, `Running`, `Scheduled`,
and `Unknown`.

Type: String

Valid Values: `Scheduled | Running | Complete | Failed | Cancelled | Timeout | Unknown`

Required: No

**userIdentity**

The ARN of the user who ran the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/queryinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/queryinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/queryinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryDefinition

QueryParameter

All content copied from https://docs.aws.amazon.com/.
