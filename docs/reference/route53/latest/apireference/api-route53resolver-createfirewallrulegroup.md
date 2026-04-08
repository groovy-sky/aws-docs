# CreateFirewallRuleGroup

Creates an empty DNS Firewall rule group for filtering DNS network traffic in a VPC. You can add rules to the new rule group
by calling [CreateFirewallRule](api-route53resolver-createfirewallrule.md).

## Request Syntax

```nohighlight

{
   "CreatorRequestId": "string",
   "Name": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CreatorRequestId](#API_route53resolver_CreateFirewallRuleGroup_RequestSyntax)**

A unique string defined by you to identify the request. This allows you to retry failed
requests without the risk of running the operation twice. This can be any unique string,
for example, a timestamp.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**[Name](#API_route53resolver_CreateFirewallRuleGroup_RequestSyntax)**

A name that lets you identify the rule group, to manage and use it.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: Yes

**[Tags](#API_route53resolver_CreateFirewallRuleGroup_RequestSyntax)**

A list of the tag keys and values that you want to associate with the rule group.

Type: Array of [Tag](api-route53resolver-tag.md) objects

Array Members: Maximum number of 200 items.

Required: No

## Response Syntax

```nohighlight

{
   "FirewallRuleGroup": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "Id": "string",
      "ModificationTime": "string",
      "Name": "string",
      "OwnerId": "string",
      "RuleCount": number,
      "ShareStatus": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRuleGroup](#API_route53resolver_CreateFirewallRuleGroup_ResponseSyntax)**

A collection of rules used to filter DNS network traffic.

Type: [FirewallRuleGroup](api-route53resolver-firewallrulegroup.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

The current account doesn't have the IAM permissions required to perform the specified Resolver operation.

This error can also be thrown when a customer has reached the 5120 character limit for a
resource policy for CloudWatch Logs.

HTTP Status Code: 400

**InternalServiceErrorException**

We encountered an unknown error. Try again in a few minutes.

HTTP Status Code: 400

**LimitExceededException**

The request caused one or more limits to be exceeded.

**ResourceType**

For a `LimitExceededException` error, the type of resource that exceeded the current limit.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/createfirewallrulegroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/createfirewallrulegroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateFirewallRule

CreateOutpostResolver
