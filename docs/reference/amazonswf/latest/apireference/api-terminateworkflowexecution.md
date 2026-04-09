# TerminateWorkflowExecution

Records a `WorkflowExecutionTerminated` event and forces closure of the
workflow execution identified by the given domain, runId, and workflowId. The child policy,
registered with the workflow type or specified when starting this execution, is applied to any
open child workflow executions of this workflow execution.

###### Important

If the identified workflow execution was in progress, it is terminated
immediately.

###### Note

If a runId isn't specified, then the `WorkflowExecutionTerminated` event
is recorded in the history of the current open workflow with the matching workflowId in the
domain.

###### Note

You should consider using [RequestCancelWorkflowExecution](api-requestcancelworkflowexecution.md) action
instead because it allows the workflow to gracefully close while TerminateWorkflowExecution doesn't.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as
follows:

- Use a `Resource` element with the domain name to limit the action to
only specified domains.

- Use an `Action` element to allow or deny permission to call this
action.

- You cannot use an IAM policy to constrain this action's parameters.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated
event attribute's `cause` parameter is set to `OPERATION_NOT_PERMITTED`.
For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF\
Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Request Syntax

```nohighlight

{
   "childPolicy": "string",
   "details": "string",
   "domain": "string",
   "reason": "string",
   "runId": "string",
   "workflowId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[childPolicy](#API_TerminateWorkflowExecution_RequestSyntax)**

If set, specifies the policy to use for the child workflow executions of the workflow
execution being terminated. This policy overrides the child policy specified for the workflow
execution at registration time or when starting the execution.

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

**[details](#API_TerminateWorkflowExecution_RequestSyntax)**

Details for terminating the workflow execution.

Type: String

Length Constraints: Maximum length of 32768.

Required: No

**[domain](#API_TerminateWorkflowExecution_RequestSyntax)**

The domain of the workflow execution to terminate.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[reason](#API_TerminateWorkflowExecution_RequestSyntax)**

A descriptive reason for terminating the workflow execution.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**[runId](#API_TerminateWorkflowExecution_RequestSyntax)**

The runId of the workflow execution to terminate.

Type: String

Length Constraints: Maximum length of 64.

Required: No

**[workflowId](#API_TerminateWorkflowExecution_RequestSyntax)**

The workflowId of the workflow execution to terminate.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**OperationNotPermittedFault**

Returned when the caller doesn't have sufficient permissions to invoke the action.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**UnknownResourceFault**

Returned when the named resource cannot be found with in the scope of this operation (region or domain). This could happen if the named resource was never created or is no longer available for this operation.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

## Examples

### TerminateWorkflowExecution Example

This example illustrates one usage of TerminateWorkflowExecution.

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
X-Amz-Date: Mon, 16 Jan 2012 04:56:34 GMT
X-Amz-Target: SimpleWorkflowService.TerminateWorkflowExecution
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=JHMRAjN6JGPawEuhiANHfiCil9KOGfDF/cuXYmuu9S4=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 218
Pragma: no-cache
Cache-Control: no-cache

{"domain": "867530901",
 "workflowId": "20110927-T-1",
 "runId": "94861fda-a714-4126-95d7-55ba847da8ab",
 "reason": "transaction canceled",
 "details": "customer credit card declined",
 "childPolicy": "TERMINATE"}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 0
Content-Type: application/json
x-amzn-RequestId: 76d68a47-3ffe-11e1-b118-3bfa5e8e7fc3
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/terminateworkflowexecution.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/terminateworkflowexecution.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

UndeprecateActivityType

All content copied from https://docs.aws.amazon.com/.
