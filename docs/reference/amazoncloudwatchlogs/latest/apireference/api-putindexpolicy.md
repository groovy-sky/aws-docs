# PutIndexPolicy

Creates or updates a _field index policy_ for the specified log group.
Only log groups in the Standard log class support field index policies. For more information
about log classes, see [Log\
classes](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-log-classes.md).

You can use field index policies to create _field indexes_ on fields
found in log events in the log group. Creating field indexes speeds up and lowers the costs
for CloudWatch Logs Insights queries that reference those field indexes, because these
queries attempt to skip the processing of log events that are known to not match the indexed
field. Good fields to index are fields that you often need to query for and fields or values
that match only a small fraction of the total log events. Common examples of indexes include
request ID, session ID, userID, and instance IDs. For more information, see [Create field indexes to improve query performance and reduce costs](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-field-indexing.md).

You can configure indexed fields as _facets_ to enable interactive
exploration and filtering of your logs in the CloudWatch Logs Insights console. Facets
allow you to view value distributions and counts for indexed fields without running queries.
When you create a field index, you can optionally set it as a facet to enable this interactive
analysis capability. For more information, see [Use facets to group and\
explore logs](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-facets.md).

To find the fields that are in your log group events, use the [GetLogGroupFields](api-getloggroupfields.md) operation.

For example, suppose you have created a field index for `requestId`. Then, any
CloudWatch Logs Insights query on that log group that includes `requestId =
          value
      ` or `requestId IN [value,
          value, ...]` will process fewer log events to reduce costs, and
have improved performance.

CloudWatch Logs provides default field indexes for all log groups in the Standard log
class. Default field indexes are automatically available for the following fields:

- `@logStream`

- `@aws.region`

- `@aws.account`

- `@source.log`

- `traceId`

Default field indexes are in addition to any custom field indexes you define within your
policy. Default field indexes are not counted towards your field index quota.

Each index policy has the following quotas and restrictions:

- As many as 20 fields can be included in the policy.

- Each field name can include as many as 100 characters.

Matches of log events to the names of indexed fields are case-sensitive. For example, a
field index of `RequestId` won't match a log event containing
`requestId`.

Log group-level field index policies created with `PutIndexPolicy` override
account-level field index policies created with [PutAccountPolicy](api-putaccountpolicy.md) that apply to log groups. If you use `PutIndexPolicy`
to create a field index policy for a log group, that log group uses only that policy for log
group-level indexing, including any facet configurations. The log group ignores any
account-wide field index policy that applies to log groups, but data source-based account
policies may still apply.

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string",
   "policyDocument": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_PutIndexPolicy_RequestSyntax)**

Specify either the log group name or log group ARN to apply this field index policy to. If
you specify an ARN, use the format
arn:aws:logs: _region_: _account-id_:log-group: _log\_group\_name_
Don't include an \* at the end.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[policyDocument](#API_PutIndexPolicy_RequestSyntax)**

The index policy document, in JSON format. The following is an example of an index policy
document that creates indexes with different types.

`"policyDocument": "{"Fields": [ "TransactionId" ], "FieldsV2": {"RequestId":
        {"type": "FIELD_INDEX"}, "APIName": {"type": "FACET"}, "StatusCode": {"type":
        "FACET"}}}"`

You can use `FieldsV2` to specify the type for each field. Supported types are
`FIELD_INDEX` and `FACET`. Field names within `Fields` and
`FieldsV2` must be mutually exclusive.

The policy document must include at least one field index. For more information about the
fields that can be included and other restrictions, see [Field index\
syntax and quotas](../../../../services/amazoncloudwatch/latest/logs/cloudwatchlogs-field-indexing-syntax.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 51200.

Required: Yes

## Response Syntax

```nohighlight

{
   "indexPolicy": {
      "lastUpdateTime": number,
      "logGroupIdentifier": "string",
      "policyDocument": "string",
      "policyName": "string",
      "source": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[indexPolicy](#API_PutIndexPolicy_ResponseSyntax)**

The index policy that you just created or updated.

Type: [IndexPolicy](api-indexpolicy.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### Create index policy for a log group

The following example creates an index policy that indexes two fields,
`RequestId` and `TransactionId`.

#### Sample Request

```

{
    "logGroupIdentifier": "service-logs",
    "policyDocument": {
        "Fields": ["RequestId", "TransactionId"]
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putindexpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putindexpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutDestinationPolicy

PutIntegration
