# ValidatePullThroughCacheRule

Validates an existing pull through cache rule for an upstream registry that requires
authentication. This will retrieve the contents of the AWS Secrets Manager secret, verify the
syntax, and then validate that authentication to the upstream registry is
successful.

## Request Syntax

```nohighlight

{
   "ecrRepositoryPrefix": "string",
   "registryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ecrRepositoryPrefix](#API_ValidatePullThroughCacheRule_RequestSyntax)**

The repository name prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: Yes

**[registryId](#API_ValidatePullThroughCacheRule_RequestSyntax)**

The registry ID associated with the pull through cache rule.
If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

## Response Syntax

```nohighlight

{
   "credentialArn": "string",
   "customRoleArn": "string",
   "ecrRepositoryPrefix": "string",
   "failure": "string",
   "isValid": boolean,
   "registryId": "string",
   "upstreamRegistryUrl": "string",
   "upstreamRepositoryPrefix": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[credentialArn](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret associated with the pull through cache
rule.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 612.

Pattern: `^arn:aws(-\w+)*:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\-pullthroughcache\/[a-zA-Z0-9\/_+=.@-]+$`

**[customRoleArn](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The ARN of the IAM role associated with the pull through cache rule.

Type: String

Length Constraints: Maximum length of 2048.

**[ecrRepositoryPrefix](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The Amazon ECR repository prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

**[failure](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The reason the validation failed. For more details about possible causes and how to
address them, see [Using pull through cache\
rules](../../../../services/amazonecr/latest/userguide/pull-through-cache.md) in the _Amazon Elastic Container Registry User Guide_.

Type: String

**[isValid](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

Whether or not the pull through cache rule was validated. If `true`, Amazon ECR
was able to reach the upstream registry and authentication was successful. If
`false`, there was an issue and validation failed. The
`failure` reason indicates the cause.

Type: Boolean

**[registryId](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[upstreamRegistryUrl](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The upstream registry URL associated with the pull through cache rule.

Type: String

**[upstreamRepositoryPrefix](#API_ValidatePullThroughCacheRule_ResponseSyntax)**

The upstream repository prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/validatepullthroughcacherule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/validatepullthroughcacherule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UploadLayerPart

Data Types

All content copied from https://docs.aws.amazon.com/.
