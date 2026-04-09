# PutDestinationPolicy

Creates or updates an access policy associated with an existing destination. An access
policy is an [IAM\
policy document](../../../../services/iam/latest/userguide/policies-overview.md) that is used to authorize claims to register a subscription filter
against a given destination.

## Request Syntax

```nohighlight

{
   "accessPolicy": "string",
   "destinationName": "string",
   "forceUpdate": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[accessPolicy](#API_PutDestinationPolicy_RequestSyntax)**

An IAM policy document that authorizes cross-account users to deliver their log events
to the associated destination. This can be up to 5120 bytes.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[destinationName](#API_PutDestinationPolicy_RequestSyntax)**

A name for an existing destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 512.

Pattern: `[^:*]*`

Required: Yes

**[forceUpdate](#API_PutDestinationPolicy_RequestSyntax)**

Specify true if you are updating an existing destination policy to grant permission to an
organization ID instead of granting permission to individual AWS accounts.
Before you update a destination policy this way, you must first update the subscription
filters in the accounts that send logs to this destination. If you do not, the subscription
filters might stop working. By specifying `true` for `forceUpdate`, you
are affirming that you have already updated the subscription filters. For more information,
see [Updating an\
existing cross-account subscription](../../../../services/amazoncloudwatch/latest/logs/cross-account-log-subscription-update.md)

If you omit this parameter, the default of `false` is used.

Type: Boolean

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

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To create or update an access policy of a destination

The following example updates the access policy of the specified
destination.

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
X-Amz-Target: Logs_20140328.PutDestinationPolicy
{
  "destinationName": "my-destination",
  "accessPolicy": "{ \"Version\": \"2012-10-17\", \"Statement\": [{ \"Effect\": \"Allow\", \"Principal\": { \"AWS\": \"111111111111\"}, \"Action\": \"logs:PutSubscriptionFilter\",\"Resource\": \"arn:aws:logs:us-east-1:123456789012:destination:my-destination\"}]}"
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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putdestinationpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putdestinationpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutDestination

PutIndexPolicy

All content copied from https://docs.aws.amazon.com/.
