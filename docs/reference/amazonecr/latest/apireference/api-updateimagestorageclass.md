# UpdateImageStorageClass

Transitions an image between storage classes. You can transition images from Amazon ECR standard storage class to Amazon ECR archival storage class for long-term storage, or restore archived images back to Amazon ECR standard.

## Request Syntax

```nohighlight

{
   "imageId": {
      "imageDigest": "string",
      "imageTag": "string"
   },
   "registryId": "string",
   "repositoryName": "string",
   "targetStorageClass": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[imageId](#API_UpdateImageStorageClass_RequestSyntax)**

An object with identifying information for an image in an Amazon ECR repository.

Type: [ImageIdentifier](api-imageidentifier.md) object

Required: Yes

**[registryId](#API_UpdateImageStorageClass_RequestSyntax)**

The AWS account ID associated with the registry that contains the image to transition. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_UpdateImageStorageClass_RequestSyntax)**

The name of the repository that contains the image to transition.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

**[targetStorageClass](#API_UpdateImageStorageClass_RequestSyntax)**

The target storage class for the image.

Type: String

Valid Values: `STANDARD | ARCHIVE`

Required: Yes

## Response Syntax

```nohighlight

{
   "imageId": {
      "imageDigest": "string",
      "imageTag": "string"
   },
   "imageStatus": "string",
   "registryId": "string",
   "repositoryName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[imageId](#API_UpdateImageStorageClass_ResponseSyntax)**

An object with identifying information for an image in an Amazon ECR repository.

Type: [ImageIdentifier](api-imageidentifier.md) object

**[imageStatus](#API_UpdateImageStorageClass_ResponseSyntax)**

The current status of the image after the call to UpdateImageStorageClass is complete. Valid values are `ACTIVE`, `ARCHIVED`, and `ACTIVATING`.

Type: String

Valid Values: `ACTIVE | ARCHIVED | ACTIVATING`

**[registryId](#API_UpdateImageStorageClass_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryName](#API_UpdateImageStorageClass_ResponseSyntax)**

The repository name associated with the request.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ImageNotFoundException**

The image requested does not exist in the specified repository.

HTTP Status Code: 400

**ImageStorageClassUpdateNotSupportedException**

The requested image storage class update is not supported.

HTTP Status Code: 400

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

### To transition an image to Amazon ECR archive

This example transitions an image with a specific digest in the `hello-repository` repository to Amazon ECR archive storage for long-term archival.

#### Sample Request

```

POST / HTTP/1.1
Host: api.ecr.us-east-1.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.UpdateImageStorageClass
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/2.0 Python/3.8.0 Darwin/20.0.0 botocore/2.0.0
X-Amz-Date: 20251117T220800Z
Authorization: AUTHPARAMS
Content-Length: 225

{
   "registryId": "724772093679",
   "repositoryName": "hello-repository",
   "imageId": {
      "imageDigest": "sha256:0b1a4e0c81c434fa7928e5c4a2651a521ebabc4ff200c65f7e25b99373efca3b"
   },
   "targetStorageClass": "ARCHIVE"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 215
Connection: keep-alive

{
   "registryId": "724772093679",
   "repositoryName": "hello-repository",
   "imageId": {
      "imageDigest": "sha256:0b1a4e0c81c434fa7928e5c4a2651a521ebabc4ff200c65f7e25b99373efca3b"
   },
   "imageStatus": "ARCHIVED"
}
```

### To restore an archived image to Amazon ECR standard

This example restores an archived image with a specific digest back to Amazon ECR standard storage.

#### Sample Request

```

POST / HTTP/1.1
Host: api.ecr.us-east-1.amazonaws.com
Accept-Encoding: identity
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.UpdateImageStorageClass
Content-Type: application/x-amz-json-1.1
User-Agent: aws-cli/2.0 Python/3.8.0 Darwin/20.0.0 botocore/2.0.0
X-Amz-Date: 20251117T220900Z
Authorization: AUTHPARAMS
Content-Length: 226

{
   "registryId": "724772093679",
   "repositoryName": "hello-repository",
   "imageId": {
      "imageDigest": "sha256:0b1a4e0c81c434fa7928e5c4a2651a521ebabc4ff200c65f7e25b99373efca3b"
   },
   "targetStorageClass": "STANDARD"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 217
Connection: keep-alive

{
   "registryId": "724772093679",
   "repositoryName": "hello-repository",
   "imageId": {
      "imageDigest": "sha256:0b1a4e0c81c434fa7928e5c4a2651a521ebabc4ff200c65f7e25b99373efca3b"
   },
   "imageStatus": "ACTIVATING"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/updateimagestorageclass.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/updateimagestorageclass.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

UpdatePullThroughCacheRule

All content copied from https://docs.aws.amazon.com/.
