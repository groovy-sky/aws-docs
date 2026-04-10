# CreateGeneratedTemplate

Creates a template from existing resources that are not already managed with CloudFormation.
You can check the status of the template generation using the
`DescribeGeneratedTemplate` API action.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GeneratedTemplateName**

The name assigned to the generated template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Resources.member.N**

An optional list of resources to be included in the generated template.

If no resources are specified,the template will be created without any resources.
Resources can be added to the template using the `UpdateGeneratedTemplate` API
action.

Type: Array of [ResourceDefinition](api-resourcedefinition.md) objects

Array Members: Minimum number of 1 item. Maximum number of 500 items.

Required: No

**StackName**

An optional name or ARN of a stack to use as the base stack for the generated
template.

Type: String

Required: No

**TemplateConfiguration**

The configuration details of the generated template, including the
`DeletionPolicy` and `UpdateReplacePolicy`.

Type: [TemplateConfiguration](api-templateconfiguration.md) object

Required: No

## Response Elements

The following element is returned by the service.

**GeneratedTemplateId**

The ID of the generated template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AlreadyExists**

The resource with the name requested already exists.

HTTP Status Code: 400

**ConcurrentResourcesLimitExceeded**

No more than 5 generated templates can be in an `InProgress` or `Pending` status at one
time. This error is also returned if a generated template that is in an `InProgress` or
`Pending` status is attempted to be updated or deleted.

HTTP Status Code: 429

**LimitExceeded**

The quota for the resource has already been reached.

For information about resource and stack limitations, see [CloudFormation quotas](../../../../services/cloudformation/latest/userguide/cloudformation-limits.md) in the
_AWS CloudFormation User Guide_.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/creategeneratedtemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/creategeneratedtemplate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateChangeSet

CreateStack

All content copied from https://docs.aws.amazon.com/.
