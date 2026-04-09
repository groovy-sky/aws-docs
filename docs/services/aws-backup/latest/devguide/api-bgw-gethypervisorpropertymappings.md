# GetHypervisorPropertyMappings

This action retrieves the property mappings for the specified hypervisor.
A hypervisor property mapping displays the relationship of entity properties
available from the hypervisor to the properties available in AWS.

## Request Syntax

```nohighlight

{
   "HypervisorArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[HypervisorArn](#API_BGW_GetHypervisorPropertyMappings_RequestSyntax)**

The Amazon Resource Name (ARN) of the hypervisor.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 500.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

Required: Yes

## Response Syntax

```nohighlight

{
   "HypervisorArn": "string",
   "IamRoleArn": "string",
   "VmwareToAwsTagMappings": [
      {
         "AwsTagKey": "string",
         "AwsTagValue": "string",
         "VmwareCategory": "string",
         "VmwareTagName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[HypervisorArn](#API_BGW_GetHypervisorPropertyMappings_ResponseSyntax)**

The Amazon Resource Name (ARN) of the hypervisor.

Type: String

Length Constraints: Minimum length of 50. Maximum length of 500.

Pattern: `arn:(aws|aws-cn|aws-us-gov):backup-gateway(:[a-zA-Z-0-9]+){3}\/[a-zA-Z-0-9]+`

**[IamRoleArn](#API_BGW_GetHypervisorPropertyMappings_ResponseSyntax)**

The Amazon Resource Name (ARN) of the IAM role.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws|aws-cn|aws-us-gov):iam::([0-9]+):role/(\S+)`

**[VmwareToAwsTagMappings](#API_BGW_GetHypervisorPropertyMappings_ResponseSyntax)**

This is a display of the mappings of VMware tags to the
AWS tags.

Type: Array of [VmwareToAwsTagMapping](api-bgw-vmwaretoawstagmapping.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

The operation did not succeed because an internal error occurred. Try again later.

**ErrorCode**

A description of which internal error occured.

HTTP Status Code: 500

**ResourceNotFoundException**

A resource that is required for the action wasn't found.

**ErrorCode**

A description of which resource wasn't found.

HTTP Status Code: 400

**ThrottlingException**

TPS has been limited to protect against intentional or unintentional
high request volumes.

**ErrorCode**

Error: TPS has been limited to protect against intentional or unintentional
high request volumes.

HTTP Status Code: 400

**ValidationException**

The operation did not succeed because a validation error occurred.

**ErrorCode**

A description of what caused the validation error.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for Python](../../../goto/boto3/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-gateway-2021-01-01/gethypervisorpropertymappings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetHypervisor

GetVirtualMachine

All content copied from https://docs.aws.amazon.com/.
