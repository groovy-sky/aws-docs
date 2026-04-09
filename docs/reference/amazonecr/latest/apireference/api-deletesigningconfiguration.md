# DeleteSigningConfiguration

Deletes the registry's signing configuration. Images pushed after deletion of the signing
configuration will no longer be automatically signed.

For more information, see [Managed signing](../../../../services/amazonecr/latest/userguide/managed-signing.md) in the
_Amazon Elastic Container Registry User Guide_.

###### Note

Deleting the signing configuration does not affect existing image signatures.

## Response Syntax

```nohighlight

{
   "registryId": "string",
   "signingConfiguration": {
      "rules": [
         {
            "repositoryFilters": [
               {
                  "filter": "string",
                  "filterType": "string"
               }
            ],
            "signingProfileArn": "string"
         }
      ]
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[registryId](#API_DeleteSigningConfiguration_ResponseSyntax)**

The AWS account ID associated with the registry.

Type: String

Pattern: `[0-9]{12}`

**[signingConfiguration](#API_DeleteSigningConfiguration_ResponseSyntax)**

The registry's deleted signing configuration.

Type: [SigningConfiguration](api-signingconfiguration.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ServerException**

These errors are usually caused by a server-side issue.

**message**

The error message associated with the exception.

HTTP Status Code: 500

**SigningConfigurationNotFoundException**

The specified signing configuration was not found. This occurs when
attempting to retrieve or delete a signing configuration that does not exist.

**message**

The error message associated with the exception.

HTTP Status Code: 400

**ValidationException**

There was an exception validating this request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/deletesigningconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/deletesigningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteRepositoryPolicy

DeregisterPullTimeUpdateExclusion

All content copied from https://docs.aws.amazon.com/.
