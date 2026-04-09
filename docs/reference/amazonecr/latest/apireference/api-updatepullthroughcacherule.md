# UpdatePullThroughCacheRule

Updates an existing pull through cache rule.

## Request Syntax

```nohighlight

{
   "credentialArn": "string",
   "customRoleArn": "string",
   "ecrRepositoryPrefix": "string",
   "registryId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[credentialArn](#API_UpdatePullThroughCacheRule_RequestSyntax)**

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that identifies the credentials to authenticate
to the upstream registry.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 612.

Pattern: `^arn:aws(-\w+)*:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\-pullthroughcache\/[a-zA-Z0-9\/_+=.@-]+$`

Required: No

**[customRoleArn](#API_UpdatePullThroughCacheRule_RequestSyntax)**

Amazon Resource Name (ARN) of the IAM role to be assumed by Amazon ECR to authenticate to
the ECR upstream registry. This role must be in the same account as the registry that
you are configuring.

Type: String

Length Constraints: Maximum length of 2048.

Required: No

**[ecrRepositoryPrefix](#API_UpdatePullThroughCacheRule_RequestSyntax)**

The repository name prefix to use when caching images from the source registry.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

Required: Yes

**[registryId](#API_UpdatePullThroughCacheRule_RequestSyntax)**

The AWS account ID associated with the registry associated with the pull through
cache rule. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

## Response Syntax

```nohighlight

{
   "credentialArn": "string",
   "customRoleArn": "string",
   "ecrRepositoryPrefix": "string",
   "registryId": "string",
   "updatedAt": number,
   "upstreamRepositoryPrefix": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[credentialArn](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret associated with the pull through cache
rule.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 612.

Pattern: `^arn:aws(-\w+)*:secretsmanager:[a-zA-Z0-9-:]+:secret:ecr\-pullthroughcache\/[a-zA-Z0-9\/_+=.@-]+$`

**[customRoleArn](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

The ARN of the IAM role associated with the pull through cache rule.

Type: String

Length Constraints: Maximum length of 2048.

**[ecrRepositoryPrefix](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

The Amazon ECR repository prefix associated with the pull through cache rule.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 30.

Pattern: `^([a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*\/?|ROOT)$`

**[registryId](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[updatedAt](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

The date and time, in JavaScript date format, when the pull through cache rule was
updated.

Type: Timestamp

**[upstreamRepositoryPrefix](#API_UpdatePullThroughCacheRule_ResponseSyntax)**

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

**SecretNotFoundException**

The ARN of the secret specified in the pull through cache rule was not found. Update
the pull through cache rule with a valid secret ARN and try again.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**UnableToAccessSecretException**

The secret is unable to be accessed. Verify the resource permissions for the secret
and try again.

HTTP Status Code: 400

**UnableToDecryptSecretValueException**

The secret is accessible but is unable to be decrypted. Verify the resource
permisisons and try again.

HTTP Status Code: 400

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/updatepullthroughcacherule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/updatepullthroughcacherule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateImageStorageClass

UpdateRepositoryCreationTemplate

All content copied from https://docs.aws.amazon.com/.
