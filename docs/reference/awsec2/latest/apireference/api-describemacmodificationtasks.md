# DescribeMacModificationTasks

Describes a System Integrity Protection (SIP) modification task or volume ownership delegation
task for an Amazon EC2 Mac instance. For more information, see [Configure \
SIP for Amazon EC2 instances](../../../../services/ec2/latest/userguide/mac-sip-settings.md#mac-sip-configure) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

Specifies one or more filters for the request:

- `instance-id` \- The ID of the instance for which the task was created.

- `task-state` \- The state of the task ( `successful` \| `failed` \|
`in-progress` \| `pending`).

- `mac-system-integrity-protection-configuration.sip-status` \- The overall SIP
state requested in the task ( `enabled` \| `disabled`).

- `start-time` \- The date and time the task was created.

- `task-type` \- The type of task ( `sip-modification` \| `volume-ownership-delegation`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MacModificationTaskId.N**

The ID of task.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining results can be seen by sending another request with the returned `nextToken` value. This value can be between 5 and 500. If `maxResults` is given a larger value than 500, you receive an error.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**NextToken**

The token to use to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**macModificationTaskSet**

Information about the tasks.

Type: Array of [MacModificationTask](api-macmodificationtask.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeMacModificationTasks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeMacModificationTasks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describemacmodificationtasks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeMacModificationTasks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describemacmodificationtasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeMacHosts

DescribeManagedPrefixLists
