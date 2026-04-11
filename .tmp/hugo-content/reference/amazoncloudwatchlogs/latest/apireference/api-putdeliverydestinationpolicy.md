# PutDeliveryDestinationPolicy

Creates and assigns an IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account. To
configure the delivery of logs from an AWS service in another account to a logs
delivery destination in the current account, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](api-putdeliverysource.md).

- Create a _delivery destination_, which is a logical object that
represents the actual delivery destination. For more information, see [PutDeliveryDestination](api-putdeliverydestination.md).

- Use this operation in the destination account to assign an IAM policy
to the destination. This policy allows delivery to that destination.

- Create a _delivery_ by pairing exactly one delivery source and one
delivery destination. For more information, see [CreateDelivery](api-createdelivery.md).

Only some AWS services support being configured as a delivery source. These
services are listed as **Supported \[V2 Permissions\]** in the
table at [Enabling logging from\
AWS services.](../../../../services/amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

The contents of the policy must include two statements. One statement enables general logs
delivery, and the other allows delivery to the chosen destination. See the examples for the
needed policies.

## Request Syntax

```nohighlight

{
   "deliveryDestinationName": "string",
   "deliveryDestinationPolicy": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[deliveryDestinationName](#API_PutDeliveryDestinationPolicy_RequestSyntax)**

The name of the delivery destination to assign this policy to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 60.

Pattern: `[\w-]*`

Required: Yes

**[deliveryDestinationPolicy](#API_PutDeliveryDestinationPolicy_RequestSyntax)**

The contents of the policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 51200.

Required: Yes

## Response Syntax

```nohighlight

{
   "policy": {
      "deliveryDestinationPolicy": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[policy](#API_PutDeliveryDestinationPolicy_ResponseSyntax)**

The contents of the policy that you just created.

Type: [Policy](api-policy.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

**ValidationException**

One of the parameters for the request is not valid.

HTTP Status Code: 400

## Examples

### Policy to use with PutDeliveryDestination

The following example creates a policy that grants permission to CloudWatch Logs
to deliver logs cross-account to a destination in the current account.

#### Sample Request

```nohighlight

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.PutDeliveryDestinationPolicy
{
   "deliveryDestinationName": "DeliveryDestinationName",
   "deliveryDestinationPolicy": "{
      "Version": "2012-10-17",
      "Statement": [
        {
          "Sid": "AllowLogDeliveryActions",
          "Effect": "Allow",
          "Principal": {
            "AWS": "arn:aws:iam::AccountID:root"
          },
          "Action": "logs:CreateDelivery",
          "Resource": [
            "arn:aws:logs:us-east-1:AccountID:delivery-source:*",
            "arn:aws:logs:us-east-1:AccountID:delivery:*",
            "arn:aws:logs:us-east-1:AccountID:delivery-destination:*"
          ]
        }
      ]
    }"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putdeliverydestinationpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putdeliverydestinationpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutDeliveryDestination

PutDeliverySource

All content copied from https://docs.aws.amazon.com/.
