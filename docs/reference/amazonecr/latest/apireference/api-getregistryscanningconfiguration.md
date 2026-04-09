# GetRegistryScanningConfiguration

Retrieves the scanning configuration for a registry.

## Response Syntax

```nohighlight

{
   "registryId": "string",
   "scanningConfiguration": {
      "rules": [
         {
            "repositoryFilters": [
               {
                  "filter": "string",
                  "filterType": "string"
               }
            ],
            "scanFrequency": "string"
         }
      ],
      "scanType": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[registryId](#API_GetRegistryScanningConfiguration_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

**[scanningConfiguration](#API_GetRegistryScanningConfiguration_ResponseSyntax)**

The scanning configuration for the registry.

Type: [RegistryScanningConfiguration](api-registryscanningconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

The specified parameter is invalid. Review the available parameters for the API
request.

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/getregistryscanningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/getregistryscanningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetRegistryPolicy

GetRepositoryPolicy

All content copied from https://docs.aws.amazon.com/.
