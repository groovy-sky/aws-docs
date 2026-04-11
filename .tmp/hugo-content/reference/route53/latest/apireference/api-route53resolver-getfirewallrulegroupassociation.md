# GetFirewallRuleGroupAssociation

Retrieves a firewall rule group association, which enables DNS filtering for a VPC with one rule group. A VPC can have more than one firewall rule group association, and a rule group can be associated with more than one VPC.

## Request Syntax

```nohighlight

{
   "FirewallRuleGroupAssociationId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallRuleGroupAssociationId](#API_route53resolver_GetFirewallRuleGroupAssociation_RequestSyntax)**

The identifier of the [FirewallRuleGroupAssociation](api-route53resolver-firewallrulegroupassociation.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "FirewallRuleGroupAssociation": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "FirewallRuleGroupId": "string",
      "Id": "string",
      "ManagedOwnerName": "string",
      "ModificationTime": "string",
      "MutationProtection": "string",
      "Name": "string",
      "Priority": number,
      "Status": "string",
      "StatusMessage": "string",
      "VpcId": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallRuleGroupAssociation](#API_route53resolver_GetFirewallRuleGroupAssociation_ResponseSyntax)**

The association that you requested.

Type: [FirewallRuleGroupAssociation](api-route53resolver-firewallrulegroupassociation.md) object

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

**ResourceNotFoundException**

The specified resource doesn't exist.

**ResourceType**

For a `ResourceNotFoundException` error, the type of resource that doesn't exist.

HTTP Status Code: 400

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/getfirewallrulegroupassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetFirewallRuleGroup

GetFirewallRuleGroupPolicy
