# CreateStackRefactor

Creates a refactor across multiple stacks, with the list of stacks and resources that are
affected.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**Description**

A description to help you identify the stack refactor.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**EnableStackCreation**

Determines if a new stack is created with the refactor.

Type: Boolean

Required: No

**ResourceMappings.member.N**

The mappings for the stack resource `Source` and stack resource
`Destination`.

Type: Array of [ResourceMapping](api-resourcemapping.md) objects

Required: No

**StackDefinitions.member.N**

The stacks being refactored.

Type: Array of [StackDefinition](api-stackdefinition.md) objects

Required: Yes

## Response Elements

The following element is returned by the service.

**StackRefactorId**

The ID associated with the stack refactor created from the [CreateStackRefactor](api-createstackrefactor.md) action.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/createstackrefactor.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/createstackrefactor.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateStackInstances

CreateStackSet
