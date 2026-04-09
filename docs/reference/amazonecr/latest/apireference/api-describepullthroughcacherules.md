# DescribePullThroughCacheRules

Returns the pull through cache rules for a registry.

## Request Syntax

```nohighlight

{
   "ecrRepositoryPrefixes": [ "string" ],
   "maxResults": number,
   "nextToken": "string",
   "registryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ecrRepositoryPrefixes](#API_DescribePullThroughCacheRules_RequestSyntax)**

The Amazon ECR repository prefixes associated with the pull through cache rules to return.
If no repository prefix value is specified, all pull through cache rules are
returned.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: No

**[maxResults](#API_DescribePullThroughCacheRules_RequestSyntax)**

The maximum number of pull through cache rules returned by
`DescribePullThroughCacheRulesRequest` in paginated output. When this
parameter is used, `DescribePullThroughCacheRulesRequest` only returns
`maxResults` results in a single page along with a `nextToken`
response element. The remaining results of the initial request can be seen by sending
another `DescribePullThroughCacheRulesRequest` request with the returned
`nextToken` value. This value can be between 1 and 1000. If this
parameter is not used, then `DescribePullThroughCacheRulesRequest` returns up
to 100 results and a `nextToken` value, if applicable.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[nextToken](#API_DescribePullThroughCacheRules_RequestSyntax)**

The `nextToken` value returned from a previous paginated
`DescribePullThroughCacheRulesRequest` request where
`maxResults` was used and the results exceeded the value of that
parameter. Pagination continues from the end of the previous results that returned the
`nextToken` value. This value is null when there are no more results to
return.

Type: String

Required: No

**[registryId](#API_DescribePullThroughCacheRules_RequestSyntax)**

The AWS account ID associated with the registry to return the pull through cache
rules for. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

## Response Syntax

```nohighlight

{
   "nextToken": "string",
   "pullThroughCacheRules": [
      {
         "createdAt": number,
         "credentialArn": "string",
         "customRoleArn": "string",
         "ecrRepositoryPrefix": "string",
         "registryId": "string",
         "updatedAt": number,
         "upstreamRegistry": "string",
         "upstreamRegistryUrl": "string",
         "upstreamRepositoryPrefix": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_DescribePullThroughCacheRules_ResponseSyntax)**

The `nextToken` value to include in a future
`DescribePullThroughCacheRulesRequest` request. When the results of a
`DescribePullThroughCacheRulesRequest` request exceed
`maxResults`, this value can be used to retrieve the next page of
results. This value is null when there are no more results to return.

Type: String

**[pullThroughCacheRules](#API_DescribePullThroughCacheRules_ResponseSyntax)**

The details of the pull through cache rules.

Type: Array of [PullThroughCacheRule](api-pullthroughcacherule.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**PullThroughCacheRuleNotFoundException**

The pull through cache rule was not found. Specify a valid pull through cache rule and
try again.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/describepullthroughcacherules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/describepullthroughcacherules.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeImageSigningStatus

DescribeRegistry

All content copied from https://docs.aws.amazon.com/.
