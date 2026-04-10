# UpdateTerminationProtection

Updates termination protection for the specified stack. If a user attempts to delete a
stack with termination protection enabled, the operation fails and the stack remains
unchanged. For more information, see [Protect a CloudFormation\
stack from being deleted](../../../../services/cloudformation/latest/userguide/using-cfn-protect-stacks.md) in the _AWS CloudFormation User Guide_.

For [nested stacks](../../../../services/cloudformation/latest/userguide/using-cfn-nested-stacks.md),
termination protection is set on the root stack and can't be changed directly on the nested
stack.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**EnableTerminationProtection**

Whether to enable termination protection on the specified stack.

Type: Boolean

Required: Yes

**StackName**

The name or unique ID of the stack for which you want to set termination
protection.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `([a-zA-Z][-a-zA-Z0-9]*)|(arn:\b(aws|aws-us-gov|aws-cn)\b:[-a-zA-Z0-9:/._+]*)`

Required: Yes

## Response Elements

The following element is returned by the service.

**StackId**

The unique ID of the stack.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/updateterminationprotection.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/updateterminationprotection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateStackSet

ValidateTemplate

All content copied from https://docs.aws.amazon.com/.
