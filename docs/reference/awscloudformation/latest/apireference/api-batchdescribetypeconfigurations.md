---
title: "BatchDescribeTypeConfigurations"
---

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cloudformation-2010-05-15/BatchDescribeTypeConfigurations)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivateType

CancelUpdateStack

All content copied from https://docs.aws.amazon.com/.
