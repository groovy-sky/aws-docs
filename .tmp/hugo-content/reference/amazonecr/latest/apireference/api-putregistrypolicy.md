# PutRegistryPolicy

Creates or updates the permissions policy for your registry.

A registry policy is used to specify permissions for another AWS account and is used
when configuring cross-account replication. For more information, see [Registry permissions](../../../../services/amazonecr/latest/userguide/registry-permissions.md) in the _Amazon Elastic Container Registry User Guide_.

## Request Syntax

```nohighlight

{
   "policyText": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[policyText](#API_PutRegistryPolicy_RequestSyntax)**

The JSON policy text to apply to your registry. The policy text follows the same
format as IAM policy text. For more information, see [Registry\
permissions](../../../../services/amazonecr/latest/userguide/registry-permissions.md) in the _Amazon Elastic Container Registry User Guide_.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10240.

Required: Yes

## Response Syntax

```nohighlight

{
   "policyText": "string",
   "registryId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[policyText](#API_PutRegistryPolicy_ResponseSyntax)**

The JSON policy text for your registry.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 10240.

**[registryId](#API_PutRegistryPolicy_ResponseSyntax)**

The registry ID associated with the request.

Type: String

Pattern: `[0-9]{12}`

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ecr-2015-09-21/putregistrypolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/putregistrypolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutLifecyclePolicy

PutRegistryScanningConfiguration

All content copied from https://docs.aws.amazon.com/.
