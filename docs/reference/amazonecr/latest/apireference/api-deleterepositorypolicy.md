# DeleteRepositoryPolicy

Deletes the repository policy associated with the specified repository.

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

**[registryId](#API_DeleteRepositoryPolicy_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository policy
to delete. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_DeleteRepositoryPolicy_RequestSyntax)**

The name of the repository that is associated with the repository policy to
delete.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "policyText": "string",
   "registryId": "string",
   "repositoryName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[policyText](#API_DeleteRepositoryPolicy_ResponseSyntax)**

The JSON repository policy that was deleted from the repository.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10240.

**[registryId](#API_DeleteRepositoryPolicy_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryName](#API_DeleteRepositoryPolicy_ResponseSyntax)**

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

**RepositoryNotFoundException**

The specified repository could not be found. Check the spelling of the specified
repository and ensure that you are performing operations on the correct registry.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**RepositoryPolicyNotFoundException**

The specified repository and registry combination does not have an associated
repository policy.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

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

This example deletes the repository policy from the `ubuntu`
repository.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Length: 28
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.DeleteRepositoryPolicy
X-Amz-Date: 20151215T003722Z
User-Agent: aws-cli/1.9.10 Python/2.7.10 Darwin/14.5.0 botocore/1.3.10
Content-Type: application/x-amz-json-1.1
Authorization: AUTHPARAMS

{
  "repositoryName": "ubuntu"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Tue, 15 Dec 2015 00:37:22 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 301
Connection: keep-alive
x-amzn-RequestId: 01817918-a2c4-11e5-a19f-014c7a9aad99

{
  "policyText": "{\n  \"Version\" : \"2012-10-17\",\n  \"Statement\" : [ {\n    \"Sid\" : \"AllowPull\",\n    \"Effect\" : \"Allow\",\n    \"Principal\" : \"*\",\n    \"Action\" : [ \"ecr:BatchGetImage\", \"ecr:GetDownloadUrlForLayer\" ]\n  } ]\n}",
  "registryId": "012345678910",
  "repositoryName": "ubuntu"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/deleterepositorypolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/deleterepositorypolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteRepositoryCreationTemplate

DeleteSigningConfiguration

All content copied from https://docs.aws.amazon.com/.
