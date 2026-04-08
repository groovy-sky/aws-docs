# DeleteGeneratedTemplate

Deleted a generated template.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**GeneratedTemplateName**

The name or Amazon Resource Name (ARN) of a generated template.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConcurrentResourcesLimitExceeded**

No more than 5 generated templates can be in an `InProgress` or `Pending` status at one
time. This error is also returned if a generated template that is in an `InProgress` or
`Pending` status is attempted to be updated or deleted.

HTTP Status Code: 429

**GeneratedTemplateNotFound**

The generated template was not found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/deletegeneratedtemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/deletegeneratedtemplate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteChangeSet

DeleteStack
