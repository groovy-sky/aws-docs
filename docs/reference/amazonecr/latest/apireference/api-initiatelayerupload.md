# InitiateLayerUpload

Notifies Amazon ECR that you intend to upload an image layer.

When an image is pushed, the InitiateLayerUpload API is called once per image layer
that has not already been uploaded. Whether or not an image layer has been uploaded is
determined by the BatchCheckLayerAvailability API action.

###### Note

This operation is used by the Amazon ECR proxy and is not generally used by
customers for pulling and pushing images. In most cases, you should use the `docker` CLI to pull, tag, and push images.

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

**[registryId](#API_InitiateLayerUpload_RequestSyntax)**

The AWS account ID associated with the registry to which you intend to upload
layers. If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_InitiateLayerUpload_RequestSyntax)**

The name of the repository to which you intend to upload layers.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "partSize": number,
   "uploadId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[partSize](#API_InitiateLayerUpload_ResponseSyntax)**

The size, in bytes, that Amazon ECR expects future layer part uploads to be.

Type: Long

Valid Range: Minimum value of 0.

**[uploadId](#API_InitiateLayerUpload_ResponseSyntax)**

The upload ID for the layer upload. This parameter is passed to further [UploadLayerPart](api-uploadlayerpart.md) and [CompleteLayerUpload](api-completelayerupload.md)
operations.

Type: String

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/initiatelayerupload.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/initiatelayerupload.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetSigningConfiguration

ListImageReferrers

All content copied from https://docs.aws.amazon.com/.
