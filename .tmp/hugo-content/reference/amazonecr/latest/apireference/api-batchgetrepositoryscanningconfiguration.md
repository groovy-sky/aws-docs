# BatchGetRepositoryScanningConfiguration

Gets the scanning configuration for one or more repositories.

## Request Syntax

```nohighlight

{
   "repositoryNames": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[repositoryNames](#API_BatchGetRepositoryScanningConfiguration_RequestSyntax)**

One or more repository names to get the scanning configuration for.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 25 items.

Length Constraints: Minimum length of 2. Maximum length of 256.

Pattern: `[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*(\/[a-z0-9]+((\.|_|__|-+)[a-z0-9]+)*)*`

Required: Yes

## Response Syntax

```nohighlight

{
   "failures": [
      {
         "failureCode": "string",
         "failureReason": "string",
         "repositoryName": "string"
      }
   ],
   "scanningConfigurations": [
      {
         "appliedScanFilters": [
            {
               "filter": "string",
               "filterType": "string"
            }
         ],
         "repositoryArn": "string",
         "repositoryName": "string",
         "scanFrequency": "string",
         "scanOnPush": boolean
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[failures](#API_BatchGetRepositoryScanningConfiguration_ResponseSyntax)**

Any failures associated with the call.

Type: Array of [RepositoryScanningConfigurationFailure](api-repositoryscanningconfigurationfailure.md) objects

**[scanningConfigurations](#API_BatchGetRepositoryScanningConfiguration_ResponseSyntax)**

The scanning configuration for the requested repositories.

Type: Array of [RepositoryScanningConfiguration](api-repositoryscanningconfiguration.md) objects

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/batchgetrepositoryscanningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchGetImage

CompleteLayerUpload

All content copied from https://docs.aws.amazon.com/.
