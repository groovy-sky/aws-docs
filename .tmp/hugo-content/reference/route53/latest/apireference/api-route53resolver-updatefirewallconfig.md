# UpdateFirewallConfig

Updates the configuration of the firewall behavior provided by DNS Firewall for a single
VPC from Amazon Virtual Private Cloud (Amazon VPC).

## Request Syntax

```nohighlight

{
   "FirewallFailOpen": "string",
   "ResourceId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[FirewallFailOpen](#API_route53resolver_UpdateFirewallConfig_RequestSyntax)**

Determines how Route 53 Resolver handles queries during failures, for example when all traffic that is sent to DNS Firewall fails to receive a reply.

- By default, fail open is disabled, which means the failure mode is closed. This approach favors security over availability.
DNS Firewall blocks queries that it is unable to evaluate properly.

- If you enable this option, the failure mode is open. This approach favors availability over security. DNS Firewall allows queries to proceed if it
is unable to properly evaluate them.

This behavior is only enforced for VPCs that have at least one DNS Firewall rule group association.

Type: String

Valid Values: `ENABLED | DISABLED | USE_LOCAL_RESOURCE_SETTING`

Required: Yes

**[ResourceId](#API_route53resolver_UpdateFirewallConfig_RequestSyntax)**

The ID of the VPC that the configuration is for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Syntax

```nohighlight

{
   "FirewallConfig": {
      "FirewallFailOpen": "string",
      "Id": "string",
      "OwnerId": "string",
      "ResourceId": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[FirewallConfig](#API_route53resolver_UpdateFirewallConfig_ResponseSyntax)**

Configuration of the firewall behavior provided by DNS Firewall for a single VPC.

Type: [FirewallConfig](api-route53resolver-firewallconfig.md) object

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

**ValidationException**

You have provided an invalid command. If you ran the `UpdateFirewallDomains` request. supported values are `ADD`,
`REMOVE`, or `REPLACE` a domain.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/updatefirewallconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/updatefirewallconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UntagResource

UpdateFirewallDomains
