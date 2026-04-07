# UpdateOutpostResolver

You can use `UpdateOutpostResolver` to update the instance count, type, or name of a Resolver on an Outpost.

## Request Syntax

```nohighlight

{
   "Id": "string",
   "InstanceCount": number,
   "Name": "string",
   "PreferredInstanceType": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Id](#API_route53resolver_UpdateOutpostResolver_RequestSyntax)**

A unique string that identifies Resolver on an Outpost.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[InstanceCount](#API_route53resolver_UpdateOutpostResolver_RequestSyntax)**

The Amazon EC2 instance count for a Resolver on the Outpost.

Type: Integer

Required: No

**[Name](#API_route53resolver_UpdateOutpostResolver_RequestSyntax)**

Name of the Resolver on the Outpost.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**[PreferredInstanceType](#API_route53resolver_UpdateOutpostResolver_RequestSyntax)**

Amazon EC2 instance type.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## Response Syntax

```nohighlight

{
   "OutpostResolver": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "Id": "string",
      "InstanceCount": number,
      "ModificationTime": "string",
      "Name": "string",
      "OutpostArn": "string",
      "PreferredInstanceType": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[OutpostResolver](#API_route53resolver_UpdateOutpostResolver_ResponseSyntax)**

The response to an `UpdateOutpostResolver` request.

Type: [OutpostResolver](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_OutpostResolver.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

**ConflictException**

The requested state transition isn't valid. For example, you can't delete a firewall
domain list if it is in the process of being deleted, or you can't import domains into a
domain list that is in the process of being deleted.

HTTP Status Code: 400

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ServiceQuotaExceededException**

Fulfilling the request would cause one or more quotas to be exceeded.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

**ValidationException**

You have provided an invalid command. If you ran the `UpdateFirewallDomains` request. supported values are `ADD`,
`REMOVE`, or `REPLACE` a domain.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53resolver-2018-04-01/UpdateOutpostResolver)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/UpdateOutpostResolver)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateFirewallRuleGroupAssociation

UpdateResolverConfig
