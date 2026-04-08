# DescribeChangeSetHooks

Returns Hook-related information for the change set and a list of changes that
CloudFormation makes when you run the change set.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ChangeSetName**

The name or Amazon Resource Name (ARN) of the change set that you want to describe.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z][-a-zA-Z0-9]*|arn:[-a-zA-Z0-9:/]*`

Required: Yes

**LogicalResourceId**

If specified, lists only the Hooks related to the specified
`LogicalResourceId`.

Type: String

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**StackName**

If you specified the name of a change set, specify the stack name or stack ID (ARN) of the
change set you want to describe.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `([a-zA-Z][-a-zA-Z0-9]*)|(arn:\b(aws|aws-us-gov|aws-cn)\b:[-a-zA-Z0-9:/._+]*)`

Required: No

## Response Elements

The following elements are returned by the service.

**ChangeSetId**

The change set identifier (stack ID).

Type: String

Length Constraints: Minimum length of 1.

Pattern: `arn:[-a-zA-Z0-9:/]*`

**ChangeSetName**

The change set name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z][-a-zA-Z0-9]*`

**Hooks.member.N**

List of Hook objects.

Type: Array of [ChangeSetHook](api-changesethook.md) objects

**NextToken**

Pagination token, `null` or empty if no more results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**StackId**

The stack identifier (stack ID).

Type: String

**StackName**

The stack name.

Type: String

**Status**

Provides the status of the change set Hook.

Type: String

Valid Values: `PLANNING | PLANNED | UNAVAILABLE`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ChangeSetNotFound**

The specified change set name or ID doesn't exit. To view valid change sets for a stack, use the
`ListChangeSets` operation.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describechangesethooks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describechangesethooks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeChangeSet

DescribeEvents
