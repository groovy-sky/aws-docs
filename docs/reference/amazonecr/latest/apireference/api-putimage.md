# PutImage

Creates or updates the image manifest and tags associated with an image.

When an image is pushed and all new image layers have been uploaded, the PutImage API
is called once to create or update the image manifest and the tags associated with the
image.

###### Note

This operation is used by the Amazon ECR proxy and is not generally used by
customers for pulling and pushing images. In most cases, you should use the `docker` CLI to pull, tag, and push images.

## Request Syntax

```nohighlight

{
   "imageDigest": "string",
   "imageManifest": "string",
   "imageManifestMediaType": "string",
   "imageTag": "string",
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[imageDigest](#API_PutImage_RequestSyntax)**

The image digest of the image manifest corresponding to the image.

Type: String

Required: No

**[imageManifest](#API_PutImage_RequestSyntax)**

The image manifest corresponding to the image to be uploaded.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4194304.

Required: Yes

**[imageManifestMediaType](#API_PutImage_RequestSyntax)**

The media type of the image manifest. If you push an image manifest that does not
contain the `mediaType` field, you must specify the
`imageManifestMediaType` in the request.

Type: String

Required: No

**[imageTag](#API_PutImage_RequestSyntax)**

The tag to associate with the image. This parameter is optional.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Required: No

**[registryId](#API_PutImage_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository in
which to put the image. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_PutImage_RequestSyntax)**

The name of the repository in which to put the image.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "image": {
      "imageId": {
         "imageDigest": "string",
         "imageTag": "string"
      },
      "imageManifest": "string",
      "imageManifestMediaType": "string",
      "registryId": "string",
      "repositoryName": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[image](#API_PutImage_ResponseSyntax)**

Details of the image uploaded.

Type: [Image](api-image.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ImageAlreadyExistsException**

The specified image has already been pushed, and there were no changes to the manifest
or image tag after the last push.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ImageDigestDoesNotMatchException**

The specified image digest does not match the digest that Amazon ECR calculated for the
image.

HTTP Status Code: 400

**ImageTagAlreadyExistsException**

The specified image is tagged with a tag that already exists. The repository is
configured for tag immutability.

HTTP Status Code: 400

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

**LayersNotFoundException**

The specified layers could not be found, or the specified layer is not valid for this
repository.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**LimitExceededException**

The operation did not succeed because it would have exceeded a service limit for your
account. For more information, see [Amazon ECR service quotas](../../../../services/amazonecr/latest/userguide/service-quotas.md) in
the Amazon Elastic Container Registry User Guide.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ReferencedImagesNotFoundException**

The manifest list is referencing an image that does not exist.

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

This example puts an image to the `amazonlinux` repository with the
tag `2016.09`.

#### Sample Request

```

POST / HTTP/1.1
Host: ecr.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 653
X-Amz-Target: AmazonEC2ContainerRegistry_V20150921.PutImage
X-Amz-Date: 20161216T201255Z
User-Agent: aws-cli/1.11.22 Python/2.7.12 Darwin/16.3.0 botocore/1.4.79
Content-Type: application/x-amz-json-1.1
Authorization: AUTHPARAMS

{
  "imageManifest": "{\n   \"schemaVersion\": 2,\n   \"mediaType\": \"application/vnd.docker.distribution.manifest.v2+json\",\n   \"config\": {\n      \"mediaType\": \"application/vnd.docker.container.image.v1+json\",\n      \"size\": 1486,\n      \"digest\": \"sha256:5b52b314511a611975c2c65e695d920acdf8ae8848fe0ef00b7d018d1f118b64\"\n   },\n   \"layers\": [\n      {\n         \"mediaType\": \"application/vnd.docker.image.rootfs.diff.tar.gzip\",\n         \"size\": 91768077,\n         \"digest\": \"sha256:8e3fa21c4cc40232e835a6761332d225c7af3235c5755f44ada2ed9d0e4ab7e8\"\n      }\n   ]\n}\n",
  "repositoryName": "amazonlinux",
  "imageTag": "2016.09"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Server: Server
Date: Fri, 16 Dec 2016 20:12:56 GMT
Content-Type: application/x-amz-json-1.1
Content-Length: 786
Connection: keep-alive
x-amzn-RequestId: 084038f1-c3cc-11e6-8d10-9da51cf53fd3

{
  "image": {
    "imageId": {
      "imageDigest": "sha256:f1d4ae3f7261a72e98c6ebefe9985cf10a0ea5bd762585a43e0700ed99863807",
      "imageTag": "2016.09"
    },
    "imageManifest": "{\n   \"schemaVersion\": 2,\n   \"mediaType\": \"application/vnd.docker.distribution.manifest.v2+json\",\n   \"config\": {\n      \"mediaType\": \"application/vnd.docker.container.image.v1+json\",\n      \"size\": 1486,\n      \"digest\": \"sha256:5b52b314511a611975c2c65e695d920acdf8ae8848fe0ef00b7d018d1f118b64\"\n   },\n   \"layers\": [\n      {\n         \"mediaType\": \"application/vnd.docker.image.rootfs.diff.tar.gzip\",\n         \"size\": 91768077,\n         \"digest\": \"sha256:8e3fa21c4cc40232e835a6761332d225c7af3235c5755f44ada2ed9d0e4ab7e8\"\n      }\n   ]\n}\n",
    "registryId": "012345678910",
    "repositoryName": "amazonlinux"
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/putimage.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/putimage.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/putimage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/putimage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/putimage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/putimage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/putimage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/putimage.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/putimage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/putimage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutAccountSetting

PutImageScanningConfiguration

All content copied from https://docs.aws.amazon.com/.
