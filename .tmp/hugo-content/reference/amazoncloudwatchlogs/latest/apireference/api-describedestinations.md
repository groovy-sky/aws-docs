# DescribeDestinations

Lists all your destinations. The results are ASCII-sorted by destination
name.

## Request Syntax

```nohighlight

{
   "DestinationNamePrefix": "string",
   "limit": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DestinationNamePrefix](#API_DescribeDestinations_RequestSyntax)**

The prefix to match. If you don't specify a value, no prefix filter is
applied.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: No

**[limit](#API_DescribeDestinations_RequestSyntax)**

The maximum number of items returned. If you don't specify a value, the default maximum
value of 50 items is used.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_DescribeDestinations_RequestSyntax)**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1.

Required: No

## Response Syntax

```nohighlight

{
   "destinations": [
      {
         "accessPolicy": "string",
         "arn": "string",
         "creationTime": number,
         "destinationName": "string",
         "roleArn": "string",
         "targetArn": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[destinations](#API_DescribeDestinations_ResponseSyntax)**

The destinations.

Type: Array of [Destination](api-destination.md) objects

**[nextToken](#API_DescribeDestinations_ResponseSyntax)**

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

## Examples

### To list all destinations

The following example lists all the destinations for the account.

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
X-Amz-Target: Logs_20140328.DescribeDestinations
{
  "destinationNamePrefix": "my-prefix"
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describedestinations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describedestinations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describedestinations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describedestinations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describedestinations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describedestinations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describedestinations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describedestinations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describedestinations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describedestinations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDeliverySources

DescribeExportTasks

All content copied from https://docs.aws.amazon.com/.
