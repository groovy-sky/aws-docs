# DeleteLifecyclePolicy

Deletes the lifecycle policy associated with the specified repository.

## Request Syntax

```nohighlight

{
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[registryId](#API_DeleteLifecyclePolicy_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository.
If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_DeleteLifecyclePolicy_RequestSyntax)**

The name of the repository.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "lastEvaluatedAt": number,
   "lifecyclePolicyText": "string",
   "registryId": "string",
   "repositoryName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[lastEvaluatedAt](#API_DeleteLifecyclePolicy_ResponseSyntax)**

The time stamp of the last time that the lifecycle policy was run.

Type: Timestamp

**[lifecyclePolicyText](#API_DeleteLifecyclePolicy_ResponseSyntax)**

The JSON lifecycle policy text.

Type: String

Length Constraints: Minimum length of 100. Maximum length of 30720.

**[registryId](#API_DeleteLifecyclePolicy_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryName](#API_DeleteLifecyclePolicy_ResponseSyntax)**

The repository name associated with the request.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**LifecyclePolicyNotFoundException**

The lifecycle policy could not be found, and no policy is set to the
repository.

HTTP Status Code: 400

**RepositoryNotFoundException**

The specified repository could not be found. Check the spelling of the specified
repository and ensure that you are performing operations on the correct registry.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature Version 4
signature. For more information about creating these signatures, see [Signature\
Version 4 Signing Process](../../../../general/latest/gr/signature-version-4.md) in the _AWS General_
_Reference_.

You only need to learn how to sign HTTP requests if you intend to manually
create them. When you use the [AWS Command Line Interface (AWS CLI)](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make requests to AWS, these tools automatically sign the
requests for you with the access key that you specify when you configure the tools.
When you use these tools, you don't need to learn how to sign requests
yourself.

### Example

This example deletes a lifecycle policy for a repository called
`project-a/amazon-ecs-sample` in the default registry for an
account.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.DeleteLifecyclePolicy
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/1.11.144 Python/3.6.1 Darwin/16.6.0 botocore/1.7.2
X-Amz-Date: 20170901T223937Z
Authorization: AUTHPARAMS
Content-Length: 48

{
   "repositoryName": "project-a/amazon-ecs-sample",
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Fri, 01 Sep 2017 19:42:18 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 340
Connection: keep-alive
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f

{
   "lastEvaluatedAt":1.504295007E9,
   "lifecyclePolicyText":"{\"rules\":[{\"rulePriority\":1,\"description\":\"Expire images older than 14 days\",\"selection\":{\"tagStatus\":\"untagged\",\"countType\":\"sinceImagePushed\",\"countUnit\":\"days\",\"countNumber\":14},\"action\":{\"type\":\"expire\"}}]}",
   "registryId":"012345678910",
   "repositoryName":"project-a/amazon-ecs-sample"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/deletelifecyclepolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/deletelifecyclepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateRepositoryCreationTemplate

DeletePullThroughCacheRule

All content copied from https://docs.aws.amazon.com/.
