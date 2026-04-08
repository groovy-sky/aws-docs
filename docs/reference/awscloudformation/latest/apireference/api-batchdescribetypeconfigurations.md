# BatchDescribeTypeConfigurations

Returns configuration data for the specified CloudFormation extensions, from the CloudFormation
registry in your current account and Region.

For more information, see [Edit configuration\
data for extensions in your account](../../../../services/cloudformation/latest/userguide/registry-set-configuration.md) in the
_AWS CloudFormation User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**TypeConfigurationIdentifiers.member.N**

The list of identifiers for the desired extension configurations.

Type: Array of [TypeConfigurationIdentifier](api-typeconfigurationidentifier.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

## Response Elements

The following elements are returned by the service.

**Errors.member.N**

A list of information concerning any errors generated during the setting of the specified
configurations.

Type: Array of [BatchDescribeTypeConfigurationsError](api-batchdescribetypeconfigurationserror.md) objects

**TypeConfigurations.member.N**

A list of any of the specified extension configurations from the CloudFormation
registry.

Type: Array of [TypeConfigurationDetails](api-typeconfigurationdetails.md) objects

**UnprocessedTypeConfigurations.member.N**

A list of any of the specified extension configurations that CloudFormation could not process
for any reason.

Type: Array of [TypeConfigurationIdentifier](api-typeconfigurationidentifier.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CFNRegistry**

An error occurred during a CloudFormation registry operation.

**Message**

A message with details about the error that occurred.

HTTP Status Code: 400

**TypeConfigurationNotFound**

The specified extension configuration can't be found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/batchdescribetypeconfigurations.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ActivateType

CancelUpdateStack
