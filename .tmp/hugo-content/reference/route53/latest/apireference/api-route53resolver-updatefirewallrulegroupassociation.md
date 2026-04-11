# UpdateFirewallRuleGroupAssociation

Changes the association of a [FirewallRuleGroup](api-route53resolver-firewallrulegroup.md) with a VPC. The association enables DNS filtering for the VPC.

## Request Syntax

```nohighlight

{
   "FirewallRuleGroupAssociationId": "string",
   "MutationProtection": "string",
   "Name": "string",
   "Priority": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallRuleGroupAssociationId](#API_route53resolver_UpdateFirewallRuleGroupAssociation_RequestSyntax)**

The identifier of the [FirewallRuleGroupAssociation](api-route53resolver-firewallrulegroupassociation.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

**[MutationProtection](#API_route53resolver_UpdateFirewallRuleGroupAssociation_RequestSyntax)**

If enabled, this setting disallows modification or removal of the association, to help prevent against accidentally altering DNS firewall protections.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**[Name](#API_route53resolver_UpdateFirewallRuleGroupAssociation_RequestSyntax)**

The name of the rule group association.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**[Priority](#API_route53resolver_UpdateFirewallRuleGroupAssociation_RequestSyntax)**

The setting that determines the processing order of the rule group among the rule
groups that you associate with the specified VPC. DNS Firewall filters VPC traffic
starting from the rule group with the lowest numeric priority setting.

You must specify a unique priority for each rule group that you associate with a single VPC.
To make it easier to insert rule groups later, leave space between the numbers, for example, use 100, 200, and so on. You
can change the priority setting for a rule group association after you create it.

Type: Integer

Required: No

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

**[FirewallRuleGroupAssociation](#API_route53resolver_UpdateFirewallRuleGroupAssociation_ResponseSyntax)**

The association that you just updated.

Type: [FirewallRuleGroupAssociation](api-route53resolver-firewallrulegroupassociation.md) object

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

**ThrottlingException**

The request was throttled. Try again in a few minutes.

HTTP Status Code: 400

**ValidationException**

You have provided an invalid command. If you ran the `UpdateFirewallDomains` request. supported values are `ADD`,
`REMOVE`, or `REPLACE` a domain.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/updatefirewallrulegroupassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateFirewallRule

UpdateOutpostResolver
