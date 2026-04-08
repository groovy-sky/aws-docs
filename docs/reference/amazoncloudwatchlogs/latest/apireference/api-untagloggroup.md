# UntagLogGroup

###### Important

The UntagLogGroup operation is on the path to deprecation. We recommend that you use
[UntagResource](api-untagresource.md) instead.

Removes the specified tags from the specified log group.

To list the tags for a log group, use [ListTagsForResource](api-listtagsforresource.md). To add tags, use [TagResource](api-tagresource.md).

When using IAM policies to control tag management for CloudWatch Logs log groups, the
condition keys `aws:Resource/key-name` and `aws:TagKeys` cannot be used
to restrict which tags users can assign.

## Request Syntax

```nohighlight

{
   "logGroupName": "string",
   "tags": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupName](#API_UntagLogGroup_RequestSyntax)**

The name of the log group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[\.\-_/#A-Za-z0-9]+`

Required: Yes

**[tags](#API_UntagLogGroup_RequestSyntax)**

The tag keys. The corresponding tags are removed from the log group.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

## Examples

### To remove tags from a log group

The following example removes the specified tags for the specified log
group.

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
X-Amz-Target: Logs_20140328.UntagLogGroup
{
  "logGroupName": "my-log-group",
  "tags": {"Project", "Environment"}
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/untagloggroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/untagloggroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TestTransformer

UntagResource
