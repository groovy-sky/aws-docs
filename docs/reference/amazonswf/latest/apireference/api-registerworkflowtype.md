# RegisterWorkflowType

Registers a new _workflow type_ and its configuration settings in
the specified domain.

The retention period for the workflow history is set by the [RegisterDomain](api-registerdomain.md) action.

###### Important

If the type already exists, then a `TypeAlreadyExists` fault is returned.
You cannot change the configuration settings of a workflow type once it is registered and it
must be registered as a new version.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as
follows:

- Use a `Resource` element with the domain name to limit the action to
only specified domains.

- Use an `Action` element to allow or deny permission to call this
action.

- Constrain the following parameters by using a `Condition` element with
the appropriate keys.

- `defaultTaskList.name`: String constraint. The key is
`swf:defaultTaskList.name`.

- `name`: String constraint. The key is `swf:name`.

- `version`: String constraint. The key is
`swf:version`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated
event attribute's `cause` parameter is set to `OPERATION_NOT_PERMITTED`.
For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF\
Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Request Syntax

```nohighlight

{
   "defaultChildPolicy": "string",
   "defaultExecutionStartToCloseTimeout": "string",
   "defaultLambdaRole": "string",
   "defaultTaskList": {
      "name": "string"
   },
   "defaultTaskPriority": "string",
   "defaultTaskStartToCloseTimeout": "string",
   "description": "string",
   "domain": "string",
   "name": "string",
   "version": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[defaultChildPolicy](#API_RegisterWorkflowType_RequestSyntax)**

If set, specifies the default policy to use for the child workflow executions when a
workflow execution of this type is terminated, by calling the [TerminateWorkflowExecution](api-terminateworkflowexecution.md) action explicitly or due to an expired timeout. This
default can be overridden when starting a workflow execution using the [StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

The supported child policies are:

- `TERMINATE` – The child executions are terminated.

- `REQUEST_CANCEL` – A request to cancel is attempted for each child
execution by recording a `WorkflowExecutionCancelRequested` event in its
history. It is up to the decider to take appropriate actions when it receives an execution
history with this event.

- `ABANDON` – No action is taken. The child executions continue to
run.

Type: String

Valid Values: `TERMINATE | REQUEST_CANCEL | ABANDON`

Required: No

**[defaultExecutionStartToCloseTimeout](#API_RegisterWorkflowType_RequestSyntax)**

If set, specifies the default maximum duration for executions of this workflow type.
You can override this default when starting an execution through the [StartWorkflowExecution](api-startworkflowexecution.md) Action or `StartChildWorkflowExecution` [Decision](api-decision.md).

The duration is specified in seconds; an integer greater than or equal to 0. Unlike
some of the other timeout parameters in Amazon SWF, you cannot specify a value of "NONE" for
`defaultExecutionStartToCloseTimeout`; there is a one-year max limit on the time
that a workflow execution can run. Exceeding this limit always causes the workflow execution
to time out.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**[defaultLambdaRole](#API_RegisterWorkflowType_RequestSyntax)**

The default IAM role attached to this workflow type.

###### Note

Executions of this workflow type need IAM roles to invoke Lambda functions. If you
don't specify an IAM role when you start this workflow type, the default Lambda role is
attached to the execution. For more information, see [https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html](../../../../services/amazonswf/latest/developerguide/lambda-task.md) in the
_Amazon SWF Developer Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**[defaultTaskList](#API_RegisterWorkflowType_RequestSyntax)**

If set, specifies the default task list to use for scheduling decision tasks for
executions of this workflow type. This default is used only if a task list isn't provided when
starting the execution through the [StartWorkflowExecution](api-startworkflowexecution.md) Action or
`StartChildWorkflowExecution` [Decision](api-decision.md).

Type: [TaskList](api-tasklist.md) object

Required: No

**[defaultTaskPriority](#API_RegisterWorkflowType_RequestSyntax)**

The default task priority to assign to the workflow type. If not assigned, then
`0` is used. Valid values are integers that range from Java's
`Integer.MIN_VALUE` (-2147483648) to `Integer.MAX_VALUE` (2147483647).
Higher numbers indicate higher priority.

For more information about setting task priority, see [Setting Task\
Priority](../../../../services/amazonswf/latest/developerguide/programming-priority.md) in the _Amazon SWF Developer Guide_.

Type: String

Required: No

**[defaultTaskStartToCloseTimeout](#API_RegisterWorkflowType_RequestSyntax)**

If set, specifies the default maximum duration of decision tasks for this workflow
type. This default can be overridden when starting a workflow execution using the [StartWorkflowExecution](api-startworkflowexecution.md) action or the `StartChildWorkflowExecution` [Decision](api-decision.md).

The duration is specified in seconds, an integer greater than or equal to
`0`. You can use `NONE` to specify unlimited duration.

Type: String

Length Constraints: Maximum length of 8.

Required: No

**[description](#API_RegisterWorkflowType_RequestSyntax)**

Textual description of the workflow type.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**[domain](#API_RegisterWorkflowType_RequestSyntax)**

The name of the domain in which to register the workflow type.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[name](#API_RegisterWorkflowType_RequestSyntax)**

The name of the workflow type.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[version](#API_RegisterWorkflowType_RequestSyntax)**

The version of the workflow type.

###### Note

The workflow type consists of the name and version, the combination of which must be
unique within the domain. To get a list of all currently registered workflow types, use the
[ListWorkflowTypes](api-listworkflowtypes.md) action.

The specified string must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**TypeAlreadyExistsFault**

Returned if the type already exists in the specified domain. You may get this fault if you are registering a type that is either already registered or deprecated, or if you undeprecate a type that is currently registered.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**UnknownResourceFault**

Returned when the named resource cannot be found with in the scope of this operation (region or domain). This could happen if the named resource was never created or is no longer available for this operation.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

## Examples

### RegisterWorkflowType Example

This example illustrates one usage of RegisterWorkflowType.

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
X-Amz-Date: Fri, 13 Jan 2012 18:59:33 GMT
X-Amz-Target: SimpleWorkflowService.RegisterWorkflowType
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=p5FUOoV3QXAafb7aK5z79Ztu5v0w9NeEqLu0ei+P9FA=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 300
Pragma: no-cache
Cache-Control: no-cache

{
  "defaultExecutionStartToCloseTimeout": "3600",
  "domain": "867530901",
  "name": "customerOrderWorkflow",
  "defaultChildPolicy": "TERMINATE",
  "defaultTaskPriority": "10",
  "defaultTaskStartToCloseTimeout": "600",
  "version": "1.0",
  "defaultTaskList": {
  "name": "mainTaskList"
  },
  "description": "Handle customer orders"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 0
Content-Type: application/json
x-amzn-RequestId: bb469e67-3e18-11e1-9914-a356b6ea8bdf
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/registerworkflowtype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/registerworkflowtype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegisterDomain

RequestCancelWorkflowExecution

All content copied from https://docs.aws.amazon.com/.
