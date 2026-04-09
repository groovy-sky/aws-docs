# StartWorkflowExecution

Starts an execution of the workflow type in the specified domain using the provided
`workflowId` and input data.

This action returns the newly started workflow execution.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as
follows:

- Use a `Resource` element with the domain name to limit the action to
only specified domains.

- Use an `Action` element to allow or deny permission to call this
action.

- Constrain the following parameters by using a `Condition` element with
the appropriate keys.

- `tagList.member.0`: The key is `swf:tagList.member.0`.

- `tagList.member.1`: The key is `swf:tagList.member.1`.

- `tagList.member.2`: The key is `swf:tagList.member.2`.

- `tagList.member.3`: The key is `swf:tagList.member.3`.

- `tagList.member.4`: The key is `swf:tagList.member.4`.

- `taskList`: String constraint. The key is
`swf:taskList.name`.

- `workflowType.name`: String constraint. The key is
`swf:workflowType.name`.

- `workflowType.version`: String constraint. The key is
`swf:workflowType.version`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated
event attribute's `cause` parameter is set to `OPERATION_NOT_PERMITTED`.
For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF\
Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Request Syntax

```nohighlight

{
   "childPolicy": "string",
   "domain": "string",
   "executionStartToCloseTimeout": "string",
   "input": "string",
   "lambdaRole": "string",
   "tagList": [ "string" ],
   "taskList": {
      "name": "string"
   },
   "taskPriority": "string",
   "taskStartToCloseTimeout": "string",
   "workflowId": "string",
   "workflowType": {
      "name": "string",
      "version": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[childPolicy](#API_StartWorkflowExecution_RequestSyntax)**

If set, specifies the policy to use for the child workflow executions of this workflow
execution if it is terminated, by calling the [TerminateWorkflowExecution](api-terminateworkflowexecution.md)
action explicitly or due to an expired timeout. This policy overrides the default child policy
specified when registering the workflow type using [RegisterWorkflowType](api-registerworkflowtype.md).

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its
history. It is up to the decider to take appropriate actions when it receives an execution
history with this event.

- `ABANDON` – No action is taken. The child executions continue to
run.

###### Note

A child policy for this workflow execution must be specified either as a default for
the workflow type or through this parameter. If neither this parameter is set nor a default
child policy was specified at registration time then a fault is returned.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: No

**[domain](#API_StartWorkflowExecution_RequestSyntax)**

The name of the domain in which the workflow execution is created.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[executionStartToCloseTimeout](#API_StartWorkflowExecution_RequestSyntax)**

The total duration for this workflow execution. This overrides the
defaultExecutionStartToCloseTimeout specified when registering the workflow type.

The duration is specified in seconds; an integer greater than or equal to
`0`. Exceeding this limit causes the workflow execution to time out. Unlike some
of the other timeout parameters in Amazon SWF, you cannot specify a value of "NONE" for this
timeout; there is a one-year max limit on the time that a workflow execution can
run.

###### Note

An execution start-to-close timeout must be specified either through this parameter
or as a default when the workflow type is registered. If neither this parameter nor a
default execution start-to-close timeout is specified, a fault is returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**[input](#API_StartWorkflowExecution_RequestSyntax)**

The input for the workflow execution. This is a free form string which should be
meaningful to the workflow you are starting. This `input` is made available to the
new workflow execution in the `WorkflowExecutionStarted` history event.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**[lambdaRole](#API_StartWorkflowExecution_RequestSyntax)**

The IAM role to attach to this workflow execution.

###### Note

Executions of this workflow type need IAM roles to invoke Lambda functions. If you
don't attach an IAM role, any attempt to schedule a Lambda task fails. This results in a
`ScheduleLambdaFunctionFailed` history event. For more information, see [https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html](../../../../services/amazonswf/latest/developerguide/lambda-task.md) in the
_Amazon SWF Developer Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**[tagList](#API_StartWorkflowExecution_RequestSyntax)**

The list of tags to associate with the workflow execution. You can specify a maximum of
5 tags. You can list workflow executions with a specific tag by calling [ListOpenWorkflowExecutions](api-listopenworkflowexecutions.md) or [ListClosedWorkflowExecutions](api-listclosedworkflowexecutions.md) and
specifying a [TagFilter](api-tagfilter.md).

Type: Array of strings

Array Members: Maximum number of 5 items.

Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

**[taskList](#API_StartWorkflowExecution_RequestSyntax)**

The task list to use for the decision tasks generated for this workflow execution. This
overrides the `defaultTaskList` specified when registering the workflow
type.

###### Note

A task list for this workflow execution must be specified either as a default for the
workflow type or through this parameter. If neither this parameter is set nor a default task
list was specified at registration time then a fault is returned.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: [TaskList](api-tasklist.md) object

Required: No

**[taskPriority](#API_StartWorkflowExecution_RequestSyntax)**

The task priority to use for this workflow execution. This overrides any default
priority that was assigned when the workflow type was registered. If not set, then the default
task priority for the workflow type is used. Valid values are integers that range from Java's
`Integer.MIN_VALUE` (-2147483648) to `Integer.MAX_VALUE` (2147483647).
Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task\
Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**[taskStartToCloseTimeout](#API_StartWorkflowExecution_RequestSyntax)**

Specifies the maximum duration of decision tasks for this workflow execution. This
parameter overrides the `defaultTaskStartToCloseTimout` specified when registering
the workflow type using [RegisterWorkflowType](api-registerworkflowtype.md).

The duration is specified in seconds, an integer greater than or equal to
`0`. You can use `NONE` to specify unlimited duration.

###### Note

A task start-to-close timeout for this workflow execution must be specified either as
a default for the workflow type or through this parameter. If neither this parameter is set
nor a default task start-to-close timeout was specified at registration time then a fault is
returned.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**[workflowId](#API_StartWorkflowExecution_RequestSyntax)**

The user defined identifier associated with the workflow execution. You can use this to
associate a custom identifier with the workflow execution. You may specify the same identifier
if a workflow execution is logically a _restart_ of a previous execution.
You cannot have two open workflow executions with the same `workflowId` at the same
time within the same domain.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[workflowType](#API_StartWorkflowExecution_RequestSyntax)**

The type of the workflow to start.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

## Response Syntax

```nohighlight

{
   "runId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[runId](#API_StartWorkflowExecution_ResponseSyntax)**

The `runId` of a workflow execution. This ID is generated by the service and
can be used to uniquely identify the workflow execution within a domain.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DefaultUndefinedFault**

The `StartWorkflowExecution` API action was called without the required
parameters set.

Some workflow execution parameters, such as the decision `taskList`, must be
set to start the execution. However, these parameters might have been set as defaults when the
workflow type was registered. In this case, you can omit these parameters from the
`StartWorkflowExecution` call and Amazon SWF uses the values defined in the workflow
type.

###### Note

If these parameters aren't set and no default parameters were defined in the workflow
type, this error is displayed.

HTTP Status Code: 400

**LimitExceededFault**

Returned by any operation if a system imposed limitation has been reached. To address this fault you should either clean up unused resources or increase the limit by contacting AWS.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**OperationNotPermittedFault**

Returned when the caller doesn't have sufficient permissions to invoke the action.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**TypeDeprecatedFault**

Returned when the specified activity or workflow type was already deprecated.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**UnknownResourceFault**

Returned when the named resource cannot be found with in the scope of this operation (region or domain). This could happen if the named resource was never created or is no longer available for this operation.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**WorkflowExecutionAlreadyStartedFault**

Returned by [StartWorkflowExecution](api-startworkflowexecution.md) when an open execution with the same workflowId is already running in
the specified domain.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

## Examples

### StartWorkflowExecution Example

This example illustrates one usage of StartWorkflowExecution.

#### Sample Request

```

POST / HTTP/1.1
Host: swf.us-east-1.amazonaws.com
User-Agent: Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.2.25) Gecko/20111212 Firefox/3.6.25 ( .NET CLR 3.5.30729; .NET4.0E)
Accept: application/json, text/javascript, */*
Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Keep-Alive: 115
Connection: keep-alive
Content-Type: application/x-amz-json-1.0
X-Requested-With: XMLHttpRequest
X-Amz-Date: Sat, 14 Jan 2012 22:45:13 GMT
X-Amz-Target: SimpleWorkflowService.StartWorkflowExecution
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=aYxuqLX+TO91kPVg+jh+aA8PWxQazQRN2+SZUGdOgU0=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 417
Pragma: no-cache
Cache-Control: no-cache

{
  "domain": "867530901",
  "taskList": {
  "name": "specialTaskList"
  },
  "taskPriority": "100",
  "taskStartToCloseTimeout": "600",
  "workflowId": "20110927-T-1",
  "childPolicy": "TERMINATE",
  "executionStartToCloseTimeout": "1800",
  "input": "arbitrary-string-that-is-meaningful-to-the-workflow",
  "workflowType": {
  "version": "1.0",
  "name": "customerOrderWorkflow"
  },
  "tagList": [
  "music purchase",
  "digital",
  "ricoh-the-dog"
  ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 48
Content-Type: application/json
x-amzn-RequestId: 6c25f6e6-3f01-11e1-9a27-0760db01a4a8

{
  "runId": "1e536162-f1ea-48b0-85f3-aade88eef2f7"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/startworkflowexecution.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/startworkflowexecution.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignalWorkflowExecution

TagResource

All content copied from https://docs.aws.amazon.com/.
