# PutDestination

Creates or updates a destination. This operation is used only to create destinations
for cross-account subscriptions.

A destination encapsulates a physical resource (such as an Amazon Kinesis stream). With
a destination, you can subscribe to a real-time stream of log events for a different account,
ingested using [PutLogEvents](api-putlogevents.md).

Through an access policy, a destination controls what is written to it. By default,
`PutDestination` does not set any access policy with the destination, which means
a cross-account user cannot call [PutSubscriptionFilter](api-putsubscriptionfilter.md) against this destination. To enable this, the destination
owner must call [PutDestinationPolicy](api-putdestinationpolicy.md) after `PutDestination`.

To perform a `PutDestination` operation, you must also have the
`iam:PassRole` permission.

## Request Syntax

```nohighlight

{
   "destinationName": "string",
   "roleArn": "string",
   "tags": {
      "string" : "string"
   },
   "targetArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[destinationName](#API_PutDestination_RequestSyntax)**

A name for the destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: Yes

**[roleArn](#API_PutDestination_RequestSyntax)**

The ARN of an IAM role that grants CloudWatch Logs permissions to call the Amazon
Kinesis `PutRecord` operation on the destination stream.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[tags](#API_PutDestination_RequestSyntax)**

An optional list of key-value pairs to associate with the resource.

For more information about tagging, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md)

Type: String to string map

Map Entries: Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]+)$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

**[targetArn](#API_PutDestination_RequestSyntax)**

The ARN of an Amazon Kinesis stream to which to deliver matching log events.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

## Response Syntax

```nohighlight

{
   "destination": {
      "accessPolicy": "string",
      "arn": "string",
      "creationTime": number,
      "destinationName": "string",
      "roleArn": "string",
      "targetArn": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[destination](#API_PutDestination_ResponseSyntax)**

The destination.

Type: [Destination](api-destination.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To create or update a destination

The following example creates the specified destination.

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
X-Amz-Target: Logs_20140328.PutDestination
{
  "destinationName": "my-destination",
  "targetArn": "arn:aws:kinesis:us-east-1:123456789012:stream/my-kinesis-stream",
  "roleArn": "arn:aws:iam::123456789012:role/my-subscription-role"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
  "destination": [
    {
      "destinationName": "my-destination",
      "targetArn": "arn:aws:kinesis:us-east-1:123456789012:stream/my-kinesis-stream",
      "roleArn": "arn:aws:iam::123456789012:role/my-subscription-role",
      "arn": "arn:aws:logs:us-east-1:123456789012:destination:my-destination",
      "creationTime": 1437584472382
    }
  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putdestination.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putdestination.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putdestination.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putdestination.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putdestination.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putdestination.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putdestination.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putdestination.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putdestination.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putdestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutDeliverySource

PutDestinationPolicy

All content copied from https://docs.aws.amazon.com/.
