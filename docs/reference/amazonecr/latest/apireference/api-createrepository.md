# CreateRepository

Creates a repository. For more information, see [Amazon ECR repositories](../../../../services/amazonecr/latest/userguide/repositories.md) in the
_Amazon Elastic Container Registry User Guide_.

## Request Syntax

```nohighlight

{
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
   "repositoryName": "string",
   "tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[encryptionConfiguration](#API_CreateRepository_RequestSyntax)**

The encryption configuration for the repository. This determines how the contents of
your repository are encrypted at rest.

Type: [EncryptionConfiguration](api-encryptionconfiguration.md) object

Required: No

**[imageScanningConfiguration](#API_CreateRepository_RequestSyntax)**

###### Important

The `imageScanningConfiguration` parameter is being deprecated, in
favor of specifying the image scanning configuration at the registry level. For more
information, see `PutRegistryScanningConfiguration`.

The image scanning configuration for the repository. This determines whether images
are scanned for known vulnerabilities after being pushed to the repository.

Type: [ImageScanningConfiguration](api-imagescanningconfiguration.md) object

Required: No

**[imageTagMutability](#API_CreateRepository_RequestSyntax)**

The tag mutability setting for the repository. If this parameter is omitted, the
default setting of `MUTABLE` will be used which will allow image tags to be
overwritten. If `IMMUTABLE` is specified, all image tags within the
repository will be immutable which will prevent them from being overwritten.

Type: String

Valid Values: `MUTABLE | IMMUTABLE | IMMUTABLE_WITH_EXCLUSION | MUTABLE_WITH_EXCLUSION`

Required: No

**[imageTagMutabilityExclusionFilters](#API_CreateRepository_RequestSyntax)**

A list of filters that specify which image tags should be excluded from the
repository's image tag mutability setting.

Type: Array of [ImageTagMutabilityExclusionFilter](api-imagetagmutabilityexclusionfilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Required: No

**[registryId](#API_CreateRepository_RequestSyntax)**

The AWS account ID associated with the registry to create the repository.
If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_CreateRepository_RequestSyntax)**

The name to use for the repository. The repository name may be specified on its own
(such as `nginx-web-app`) or it can be prepended with a namespace to group
the repository into a category (such as `project-a/nginx-web-app`).

The repository name must start with a letter and can only contain lowercase letters,
numbers, hyphens, underscores, and forward slashes.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

**[tags](#API_CreateRepository_RequestSyntax)**

The metadata that you apply to the repository to help you categorize and organize
them. Each tag consists of a key and an optional value, both of which you define.
Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

Type: Array of [Tag](api-tag.md) objects

Required: No

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

**[repository](#API_CreateRepository_ResponseSyntax)**

The repository that was created.

Type: [Repository](api-repository.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**InvalidTagParameterException**

An invalid parameter has been specified. Tag keys can have a maximum character length of 128 characters, and tag values can have
a maximum length of 256 characters.

HTTP Status Code: 400

**KmsException**

The operation failed due to a KMS exception.

**kmsError**

The error code returned by AWS KMS.

HTTP Status Code: 400

**LimitExceededException**

The operation did not succeed because it would have exceeded a service limit for your
account. For more information, see [Amazon ECR service quotas](../../../../services/amazonecr/latest/userguide/service-quotas.md) in
the Amazon Elastic Container Registry User Guide.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**RepositoryAlreadyExistsException**

The specified repository already exists in the specified registry.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**TooManyTagsException**

The list of tags on the repository is over the limit. The maximum number of tags that
can be applied to a repository is 50.

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

This example creates a repository called `sample-repo` in the
default registry for an account.

#### Sample Request

```

POST / HTTP/1.1
Host: api.ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.CreateRepository
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/1.16.190 Python/3.6.1 Darwin/16.7.0 botocore/1.12.180
X-Amz-Date: 20190715T204735Z
Authorization: AUTHPARAMS
Content-Length: 33

{
   "repositoryName": "sample-repo"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 339
Connection: keep-alive

{
   "repository":{
      "repositoryArn":"arn:aws:ecr:us-west-2:012345678910:repository/sample-repo",
      "registryId":"012345678910",
      "repositoryName":"sample-repo",
      "repositoryUri":"012345678910.dkr.ecr.us-west-2.amazonaws.com/sample-repo",
      "createdAt":1.563223656E9,
      "imageTagMutability":"MUTABLE",
      "imageScanningConfiguration": {
            "scanOnPush": false
      }
   }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/createrepository.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/createrepository.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/createrepository.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/createrepository.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/createrepository.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/createrepository.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/createrepository.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/createrepository.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/createrepository.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/createrepository.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreatePullThroughCacheRule

CreateRepositoryCreationTemplate

All content copied from https://docs.aws.amazon.com/.
