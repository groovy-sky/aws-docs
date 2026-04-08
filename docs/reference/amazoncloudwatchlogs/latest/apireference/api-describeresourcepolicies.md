# DescribeResourcePolicies

Lists the resource policies in this account.

## Request Syntax

```nohighlight

{
   "limit": number,
   "nextToken": "string",
   "policyScope": "string",
   "resourceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[limit](#API_DescribeResourcePolicies_RequestSyntax)**

The maximum number of resource policies to be displayed with one call of this
API.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 50.

Required: No

**[nextToken](#API_DescribeResourcePolicies_RequestSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**[policyScope](#API_DescribeResourcePolicies_RequestSyntax)**

Specifies the scope of the resource policy. Valid values are `ACCOUNT` or
`RESOURCE`. When not specified, defaults to `ACCOUNT`.

Type: String

Valid Values: `ACCOUNT | RESOURCE`

Required: No

**[resourceArn](#API_DescribeResourcePolicies_RequestSyntax)**

The ARN of the CloudWatch Logs resource for which to query the resource policy.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "resourcePolicies": [
      {
         "lastUpdatedTime": number,
         "policyDocument": "string",
         "policyName": "string",
         "policyScope": "string",
         "resourceArn": "string",
         "revisionId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_DescribeResourcePolicies_ResponseSyntax)**

The token for the next set of items to return. The token expires after 24
hours.

Type: String

Length Constraints: Minimum length of 1.

**[resourcePolicies](#API_DescribeResourcePolicies_ResponseSyntax)**

The resource policies that exist in this account.

Type: Array of [ResourcePolicy](api-resourcepolicy.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/describeresourcepolicies.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/describeresourcepolicies.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeQueryDefinitions

DescribeSubscriptionFilters
