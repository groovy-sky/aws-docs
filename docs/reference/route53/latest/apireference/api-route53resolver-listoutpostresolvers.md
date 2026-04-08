# ListOutpostResolvers

Lists all the Resolvers on Outposts that were created using the current AWS account.

## Request Syntax

```nohighlight

{
   "MaxResults": number,
   "NextToken": "string",
   "OutpostArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[MaxResults](#API_route53resolver_ListOutpostResolvers_RequestSyntax)**

The maximum number of Resolvers on the Outpost that you want to return in the response to a
`ListOutpostResolver` request. If you don't specify a value for
`MaxResults`, the request returns up to 100 Resolvers.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**[NextToken](#API_route53resolver_ListOutpostResolvers_RequestSyntax)**

For the first `ListOutpostResolver` request, omit this value.

Type: String

Required: No

**[OutpostArn](#API_route53resolver_ListOutpostResolvers_RequestSyntax)**

The Amazon Resource Name (ARN) of the Outpost.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^arn:aws([a-z-]+)?:outposts:[a-z\d-]+:\d{12}:outpost/op-[a-f0-9]{17}$`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "OutpostResolvers": [
      {
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
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_route53resolver_ListOutpostResolvers_ResponseSyntax)**

If more than `MaxResults` Resolvers match the specified criteria, you can submit another
`ListOutpostResolver` request to get the next group of results. In the next request, specify the value of `NextToken` from the previous response.

Type: String

**[OutpostResolvers](#API_route53resolver_ListOutpostResolvers_ResponseSyntax)**

The Resolvers on Outposts that were created by using the current AWS account,
and that match the specified filters, if any.

Type: Array of [OutpostResolver](api-route53resolver-outpostresolver.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53resolver-2018-04-01/listoutpostresolvers.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53resolver-2018-04-01/listoutpostresolvers.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListFirewallRules

ListResolverConfigs
