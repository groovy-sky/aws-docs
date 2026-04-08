# ListLogGroups

Returns a list of log groups in the Region in your account. If you are performing this
action in a monitoring account, you can choose to also return log groups from source accounts
that are linked to the monitoring account. For more information about using cross-account
observability to set up monitoring accounts and source accounts, see [CloudWatch cross-account observability](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-unified-cross-account.md).

You can optionally filter the list by log group class, by using regular expressions in
your request to match strings in the log group names, by using the fieldIndexes parameter to
filter log groups based on which field indexes are configured, by using the dataSources
parameter to filter log groups by data source types, and by using the fieldIndexNames
parameter to filter by specific field index names.

This operation is paginated. By default, your first use of this operation returns 50
results, and includes a token to use in a subsequent operation to return more results.

## Request Syntax

```nohighlight

{
   "accountIdentifiers": [ "string" ],
   "dataSources": [
      {
         "name": "string",
         "type": "string"
      }
   ],
   "fieldIndexNames": [ "string" ],
   "includeLinkedAccounts": boolean,
   "limit": number,
   "logGroupClass": "string",
   "logGroupNamePattern": "string",
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[accountIdentifiers](#API_ListLogGroups_RequestSyntax)**

When `includeLinkedAccounts` is set to `true`, use this parameter to
specify the list of accounts to search. You can specify as many as 20 account IDs in the
array.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 20 items.

Length Constraints: Fixed length of 12.

Pattern: `^\d{12}$`

Required: No

**[dataSources](#API_ListLogGroups_RequestSyntax)**

An array of data source filters to filter log groups by their associated data sources. You
can filter by data source name, type, or both. Multiple filters within the same dimension are
combined with OR logic, while filters across different dimensions are combined with AND
logic.

Type: Array of [DataSourceFilter](api-datasourcefilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: No

**[fieldIndexNames](#API_ListLogGroups_RequestSyntax)**

An array of field index names to filter log groups that have specific field indexes. Only
log groups containing all specified field indexes are returned. You can specify 1 to 20 field
index names, each with 1 to 512 characters.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[includeLinkedAccounts](#API_ListLogGroups_RequestSyntax)**

If you are using a monitoring account, set this to `true` to have the operation
return log groups in the accounts listed in `accountIdentifiers`.

If this parameter is set to `true` and `accountIdentifiers` contains
a null value, the operation returns all log groups in the monitoring account and all log
groups in all source accounts that are linked to the monitoring account.

The default for this parameter is `false`.

Type: Boolean

Required: No

**[limit](#API_ListLogGroups_RequestSyntax)**

The maximum number of log groups to return. If you omit this parameter, the default is
up to 50 log groups.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[logGroupClass](#API_ListLogGroups_RequestSyntax)**

Use this parameter to limit the results to only those log groups in the specified log
group class. If you omit this parameter, log groups of all classes can be returned.

Type: String

Valid Values: `STANDARD | INFREQUENT_ACCESS | DELIVERY`

Required: No

**[logGroupNamePattern](#API_ListLogGroups_RequestSyntax)**

Use this parameter to limit the returned log groups to only those with names that match
the pattern that you specify. This parameter is a regular expression that can match prefixes
and substrings, and supports wildcard matching and matching multiple patterns, as in the
following examples.

- Use `^` to match log group names by prefix.

- For a substring match, specify the string to match. All matches are case
sensitive

- To match multiple patterns, separate them with a `|` as in the example
`^/aws/lambda|discovery`

You can specify as many as five different regular expression patterns in this field, each
of which must be between 3 and 24 characters. You can include the `^` symbol as
many as five times, and include the `|` symbol as many as four times.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 129.

Pattern: `(\^?[\.\-_\/#A-Za-z0-9]{3,24})(\|\^?[\.\-_\/#A-Za-z0-9]{3,24}){0,4}`

Required: No

**[nextToken](#API_ListLogGroups_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "logGroups": [
      {
         "logGroupArn": "string",
         "logGroupClass": "string",
         "logGroupName": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[logGroups](#API_ListLogGroups_ResponseSyntax)**

An array of structures, where each structure contains the information about one log
group.

Type: Array of [LogGroupSummary](api-loggroupsummary.md) objects

**[nextToken](#API_ListLogGroups_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/listloggroups.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/listloggroups.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/listloggroups.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/listloggroups.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/listloggroups.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/listloggroups.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/listloggroups.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/listloggroups.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/listloggroups.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/listloggroups.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListLogAnomalyDetectors

ListLogGroupsForQuery
