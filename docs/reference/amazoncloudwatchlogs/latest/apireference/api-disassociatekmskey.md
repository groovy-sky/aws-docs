# DisassociateKmsKey

Disassociates the specified AWS KMS key from the specified log group or
from all CloudWatch Logs Insights query results in the account.

When you use `DisassociateKmsKey`, you specify either the
`logGroupName` parameter or the `resourceIdentifier` parameter. You
can't specify both of those parameters in the same operation.

- Specify the `logGroupName` parameter to stop using the AWS KMS key to encrypt future log events ingested and stored in the log group.
Instead, they will be encrypted with the default CloudWatch Logs method. The log events
that were ingested while the key was associated with the log group are still encrypted
with that key. Therefore, CloudWatch Logs will need permissions for the key whenever
that data is accessed.

- Specify the `resourceIdentifier` parameter with the
`query-result` resource to stop using the AWS KMS key to
encrypt the results of all future [StartQuery](api-startquery.md)
operations in the account. They will instead be encrypted with the default CloudWatch Logs method. The results from queries that ran while the key was associated with
the account are still encrypted with that key. Therefore, CloudWatch Logs will need
permissions for the key whenever that data is accessed.

It can take up to 5 minutes for this operation to take effect.

## Request Syntax

```nohighlight

{
   "logGroupName": "string",
   "resourceIdentifier": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupName](#API_DisassociateKmsKey_RequestSyntax)**

The name of the log group.

In your `DisassociateKmsKey` operation, you must specify either the
`resourceIdentifier` parameter or the `logGroup` parameter, but you
can't specify both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: No

**[resourceIdentifier](#API_DisassociateKmsKey_RequestSyntax)**

Specifies the target for this operation. You must specify one of the following:

- Specify the ARN of a log group to stop having CloudWatch Logs use the AWS KMS key to encrypt log events that are ingested and stored by that log
group. After you run this operation, CloudWatch Logs encrypts ingested log events with
the default CloudWatch Logs method. The log group ARN must be in the following format.
Replace _REGION_ and _ACCOUNT\_ID_ with your Region
and account ID.

`arn:aws:logs:REGION:ACCOUNT_ID:log-group:LOG_GROUP_NAME
                          `

- Specify the following ARN to stop using this key to encrypt the results of future
[StartQuery](api-startquery.md)
operations in this account. Replace _REGION_ and
_ACCOUNT\_ID_ with your Region and account ID.

`arn:aws:logs:REGION:ACCOUNT_ID:query-result:*`

In your `DisssociateKmsKey` operation, you must specify either the
`resourceIdentifier` parameter or the `logGroup` parameter, but you
can't specify both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w+=/:,.@\-\*]*`

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

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

### To disassociate an KMS key from a log group

The following example disassociates the associated KMS key from the
specified log group.

#### Sample Request

```

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.DisassociateKmsKey
{
  "logGroupName": "my-log-group",
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/disassociatekmskey.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/disassociatekmskey.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSubscriptionFilters

DisassociateSourceFromS3TableIntegration
