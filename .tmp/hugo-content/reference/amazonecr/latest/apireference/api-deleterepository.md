# DeleteRepository

Deletes a repository. If the repository isn't empty, you must either delete the
contents of the repository or use the `force` option to delete the repository
and have Amazon ECR delete all of its contents on your behalf.

## Request Syntax

```nohighlight

{
   "force": boolean,
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[force](#API_DeleteRepository_RequestSyntax)**

If true, deleting the repository force deletes the contents of the repository. If
false, the repository must be empty before attempting to delete it.

Type: Boolean

Required: No

**[registryId](#API_DeleteRepository_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository to
delete. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_DeleteRepository_RequestSyntax)**

The name of the repository to delete.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "repository": {
      "createdAt": number,
      "encryptionConfiguration": {
         "encryptionType": "string",
         "kmsKey": "string"
      },
      "imageScanningConfiguration": {
         "scanOnPush": boolean
      },
      "imageTagMutability": "string",
      "imageTagMutabilityExclusionFilters": [
         {
            "filter": "string",
            "filterType": "string"
         }
      ],
      "registryId": "string",
      "repositoryArn": "string",
      "repositoryName": "string",
      "repositoryUri": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[repository](#API_DeleteRepository_ResponseSyntax)**

The repository that was deleted.

Type: [Repository](api-repository.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**KmsException**

The operation failed due to a KMS exception.

**kmsError**

The error code returned by AWS KMS.

HTTP Status Code: 400

**RepositoryNotEmptyException**

The specified repository contains images. To delete a repository that contains images,
you must force the deletion with the `force` parameter.

**message**

The error message associated with the exception.

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

This example deletes a repository named `ubuntu` in the default
registry for an account.

#### Sample Request

```

POST / HTTP/1.1
Host: api.ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.DeleteRepository
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/1.16.190 Python/3.6.1 Darwin/16.7.0 botocore/1.12.180
X-Amz-Date: 20190715T205933Z
Authorization: AUTHPARAMS
Content-Length: 33

{
   "repositoryName":"sample-repo"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 252
Connection: keep-alive

{
   "repository":{
      "createdAt":1.563223656E9,
      "registryId":"012345678910",
      "repositoryArn":"arn:aws:ecr:us-west-2:012345678910:repository/sample-repo",
      "repositoryName":"sample-repo",
      "repositoryUri":"012345678910.dkr.ecr.us-west-2.amazonaws.com/sample-repo"
   }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/deleterepository.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/deleterepository.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteRegistryPolicy

DeleteRepositoryCreationTemplate

All content copied from https://docs.aws.amazon.com/.
