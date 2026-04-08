# DescribeEvents

Returns CloudFormation events based on flexible query criteria. Groups events by operation ID,
enabling you to focus on individual stack operations during deployment.

An operation is any action performed on a stack, including stack lifecycle actions
(Create, Update, Delete, Rollback), change set creation, nested stack creation, and automatic
rollbacks triggered by failures. Each operation has a unique identifier (Operation ID) and
represents a discrete change attempt on the stack.

Returns different types of events including:

- **Progress events** \- Status updates during stack operation
execution.

- **Validation errors** \- Failures from CloudFormation Early
Validations.

- **Provisioning errors** \- Resource creation and update
failures.

- **Hook invocation errors** \- Failures from CloudFormation
Hook during stack operations.

###### Note

One of `ChangeSetName`, `OperationId` or `StackName`
must be specified as input.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ChangeSetName**

The name or Amazon Resource Name (ARN) of the change set for which you want to retrieve
events.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Pattern: `[a-zA-Z][-a-zA-Z0-9]*|arn:[-a-zA-Z0-9:/]*`

Required: No

**Filters**

Filters to apply when retrieving events.

Type: [EventFilter](api-eventfilter.md) object

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**OperationId**

The unique identifier of the operation for which you want to retrieve events.

Type: String

Required: No

**StackName**

The name or unique stack ID for which you want to retrieve events. If you specified the
name of a change set, specify the stack name or ID (ARN) of the change set you want to
describe.

Type: String

Length Constraints: Minimum length of 1.

Pattern: `([a-zA-Z][-a-zA-Z0-9]*)|(arn:\b(aws|aws-us-gov|aws-cn)\b:[-a-zA-Z0-9:/._+]*)`

Required: No

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call `DescribeEvents` again and
assign that token to the request object's `NextToken` parameter. If the request
returns all results, `NextToken` is set to `null`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**OperationEvents.member.N**

A list of operation events that match the specified criteria.

Type: Array of [OperationEvent](api-operationevent.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describeevents.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describeevents.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeChangeSetHooks

DescribeGeneratedTemplate
