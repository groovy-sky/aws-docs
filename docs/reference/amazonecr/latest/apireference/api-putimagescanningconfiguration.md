# PutImageScanningConfiguration

###### Important

The `PutImageScanningConfiguration` API is being deprecated, in favor
of specifying the image scanning configuration at the registry level. For more
information, see [PutRegistryScanningConfiguration](api-putregistryscanningconfiguration.md).

Updates the image scanning configuration for the specified repository.

## Request Syntax

```nohighlight

{
   "imageScanningConfiguration": {
      "scanOnPush": boolean
   },
   "registryId": "string",
   "repositoryName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[imageScanningConfiguration](#API_PutImageScanningConfiguration_RequestSyntax)**

The image scanning configuration for the repository. This setting determines whether
images are scanned for known vulnerabilities after being pushed to the
repository.

Type: [ImageScanningConfiguration](api-imagescanningconfiguration.md) object

Required: Yes

**[registryId](#API_PutImageScanningConfiguration_RequestSyntax)**

The AWS account ID associated with the registry that contains the repository in
which to update the image scanning configuration setting.
If you do not specify a registry, the default registry is assumed.

Type: String

Pattern: `[0-9]{12}`

Required: No

**[repositoryName](#API_PutImageScanningConfiguration_RequestSyntax)**

The name of the repository in which to update the image scanning configuration
setting.

Type: String

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "imageScanningConfiguration": {
      "scanOnPush": boolean
   },
   "registryId": "string",
   "repositoryName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[imageScanningConfiguration](#API_PutImageScanningConfiguration_ResponseSyntax)**

The image scanning configuration setting for the repository.

Type: [ImageScanningConfiguration](api-imagescanningconfiguration.md) object

**[registryId](#API_PutImageScanningConfiguration_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[repositoryName](#API_PutImageScanningConfiguration_ResponseSyntax)**

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/putimagescanningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/putimagescanningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutImage

PutImageTagMutability

All content copied from https://docs.aws.amazon.com/.
