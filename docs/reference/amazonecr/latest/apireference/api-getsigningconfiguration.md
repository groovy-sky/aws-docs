---
title: "GetSigningConfiguration"
---

# GetSigningConfiguration

Retrieves the registry's signing configuration, which defines
rules for automatically signing images using AWS Signer.

For more information, see [Managed signing](../../../../services/amazonecr/latest/userguide/managed-signing.md) in the
_Amazon Elastic Container Registry User Guide_.

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

**[registryId](#API_GetSigningConfiguration_ResponseSyntax)**

The AWS account ID associated with the registry.

Type: String

Pattern: `[0-9]{12}`

**[signingConfiguration](#API_GetSigningConfiguration_ResponseSyntax)**

The registry's signing configuration.

Type: [SigningConfiguration](api-signingconfiguration.md) object

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ecr-2015-09-21/GetSigningConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecr-2015-09-21/GetSigningConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetRepositoryPolicy

InitiateLayerUpload

All content copied from https://docs.aws.amazon.com/.
