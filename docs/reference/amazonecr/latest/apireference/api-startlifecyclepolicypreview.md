# StartLifecyclePolicyPreview

Starts a preview of a lifecycle policy for the specified repository. This allows you
to see the results before associating the lifecycle policy with the repository.

## Request Syntax

```nohighlight

{
   "lifecyclePolicyText": "string",
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[lifecyclePolicyText](#API_StartLifecyclePolicyPreview_RequestSyntax)**

The policy to be evaluated against. If you do not specify a policy, the current policy
for the repository is used.

Type: String

Length Constraints: Minimum length of 100. Maximum length of 30720.

Required: No

**[registryId](#API_StartLifecyclePolicyPreview_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository.
If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_StartLifecyclePolicyPreview_RequestSyntax)**

The name of the repository to be evaluated.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "lifecyclePolicyText": "string",
   "registryId": "string",
   "repositoryName": "string",
   "status": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[lifecyclePolicyText](#API_StartLifecyclePolicyPreview_ResponseSyntax)**

The JSON repository policy text.

Type: String

Length Constraints: Minimum length of 100. Maximum length of 30720.

**[registryId](#API_StartLifecyclePolicyPreview_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryName](#API_StartLifecyclePolicyPreview_ResponseSyntax)**

The repository name associated with the request.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

**[status](#API_StartLifecyclePolicyPreview_ResponseSyntax)**

The status of the lifecycle policy preview request.

Type: String

Valid Values: `IN_PROGRESS | COMPLETE | EXPIRED | FAILED`

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

**LifecyclePolicyPreviewInProgressException**

The previous lifecycle policy preview request has not completed. Wait and try
again.

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

This example creates a lifecycle policy preview to expire images older than 14
days for a repository called `project-a/amazon-ecs-sample` in the
default registry for an account.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.StartLifecyclePolicyPreview
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/1.11.144 Python/3.6.1 Darwin/16.6.0 botocore/1.7.2
X-Amz-Date: 20170901T221604Z
Authorization: AUTHPARAMS
Content-Length: 535

{
   "repositoryName": "project-a/amazon-ecs-sample",
   "lifecyclePolicyText": "{\n    \"rules\": [\n        {\n            \"rulePriority\": 1,\n            \"description\": \"Expire images older than 14 days\",\n            \"selection\": {\n                \"tagStatus\": \"untagged\",\n                \"countType\": \"sinceImagePushed\",\n                \"countUnit\": \"days\",\n                \"countNumber\": 14\n            },\n            \"action\": {\n                \"type\": \"expire\"\n            }\n        }\n    ]\n}\n"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Fri, 01 Sep 2017 22:16:05 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 583
Connection: keep-alive
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f

{
      "lifecyclePolicyText":"{\n    \"rules\": [\n        {\n            \"rulePriority\": 1,\n            \"description\": \"Expire images older than 14 days\",\n            \"selection\": {\n                \"tagStatus\": \"untagged\",\n                \"countType\": \"sinceImagePushed\",\n                \"countUnit\": \"days\",\n                \"countNumber\": 14\n            },\n            \"action\": {\n                \"type\": \"expire\"\n            }\n        }\n    ]\n}\n",
      "registryId":"012345678910",
      "repositoryName":"project-a/amazon-ecs-sample",
      "status":"IN_PROGRESS"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/startlifecyclepolicypreview.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/startlifecyclepolicypreview.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StartImageScan

TagResource

All content copied from https://docs.aws.amazon.com/.
