# DeleteFirewallDomainList

Deletes the specified domain list.

## Request Syntax

```nohighlight

{
   "FirewallDomainListId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallDomainListId](#API_route53resolver_DeleteFirewallDomainList_RequestSyntax)**

The ID of the domain list that you want to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "FirewallDomainList": {
      "Arn": "string",
      "CreationTime": "string",
      "CreatorRequestId": "string",
      "DomainCount": number,
      "Id": "string",
      "ManagedOwnerName": "string",
      "ModificationTime": "string",
      "Name": "string",
      "Status": "string",
      "StatusMessage": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallDomainList](#API_route53resolver_DeleteFirewallDomainList_ResponseSyntax)**

The domain list that you just deleted.

Type: [FirewallDomainList](api-route53resolver-firewalldomainlist.md) object

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/deletefirewalldomainlist.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/deletefirewalldomainlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateResolverRule

DeleteFirewallRule
