# DescribeWorkflowExecution

Returns information about the specified workflow execution including its type and some
statistics.

###### Note

This operation is eventually consistent. The results are best effort and may not
exactly reflect recent updates and changes.

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
   "domain": "string",
   "execution": {
      "runId": "string",
      "workflowId": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[domain](#API_DescribeWorkflowExecution_RequestSyntax)**

The name of the domain containing the workflow execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[execution](#API_DescribeWorkflowExecution_RequestSyntax)**

The workflow execution to describe.

Type: [WorkflowExecution](api-workflowexecution.md) object

Required: Yes

## Response Syntax

```nohighlight

{
   "executionConfiguration": {
      "childPolicy": "string",
      "executionStartToCloseTimeout": "string",
      "lambdaRole": "string",
      "taskList": {
         "name": "string"
      },
      "taskPriority": "string",
      "taskStartToCloseTimeout": "string"
   },
   "executionInfo": {
      "cancelRequested": boolean,
      "closeStatus": "string",
      "closeTimestamp": number,
      "execution": {
         "runId": "string",
         "workflowId": "string"
      },
      "executionStatus": "string",
      "parent": {
         "runId": "string",
         "workflowId": "string"
      },
      "startTimestamp": number,
      "tagList": [ "string" ],
      "workflowType": {
         "name": "string",
         "version": "string"
      }
   },
   "latestActivityTaskTimestamp": number,
   "latestExecutionContext": "string",
   "openCounts": {
      "openActivityTasks": number,
      "openChildWorkflowExecutions": number,
      "openDecisionTasks": number,
      "openLambdaFunctions": number,
      "openTimers": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[executionConfiguration](#API_DescribeWorkflowExecution_ResponseSyntax)**

The configuration settings for this workflow execution including timeout values, tasklist etc.

Type: [WorkflowExecutionConfiguration](api-workflowexecutionconfiguration.md) object

**[executionInfo](#API_DescribeWorkflowExecution_ResponseSyntax)**

Information about the workflow execution.

Type: [WorkflowExecutionInfo](api-workflowexecutioninfo.md) object

**[latestActivityTaskTimestamp](#API_DescribeWorkflowExecution_ResponseSyntax)**

The time when the last activity task was scheduled for this workflow execution. You can use this information to determine if the workflow has not made progress for an unusually long period of time and might require a corrective action.

Type: Timestamp

**[latestExecutionContext](#API_DescribeWorkflowExecution_ResponseSyntax)**

The latest executionContext provided by the decider for this workflow execution. A decider can provide an
executionContext (a free-form string) when closing a decision task using [RespondDecisionTaskCompleted](api-responddecisiontaskcompleted.md).

Type: String

Length Constraints: Maximum length of 32768.

**[openCounts](#API_DescribeWorkflowExecution_ResponseSyntax)**

The number of tasks for this workflow execution. This includes open and closed tasks of all types.

Type: [WorkflowExecutionOpenCounts](api-workflowexecutionopencounts.md) object

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

### DescribeWorkflowExecution Example

This example illustrates one usage of DescribeWorkflowExecution.

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
X-Amz-Date: Sun, 15 Jan 2012 02:05:18 GMT
X-Amz-Target: SimpleWorkflowService.DescribeWorkflowExecution
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=ufQVcSkfUyGPLiS8xbkEBqEc2PmEEE/3Lb9Kr8yozs8=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 127
Pragma: no-cache
Cache-Control: no-cache

{
  "domain": "867530901",
  "execution": {
  "workflowId": "20110927-T-1",
  "runId": "06b8f87a-24b3-40b6-9ceb-9676f28e9493"
  }
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 577
Content-Type: application/json
x-amzn-RequestId: 5f85ef79-3f1d-11e1-9e8f-57bb03e21482

{
  "executionConfiguration": {
  "executionStartToCloseTimeout": "3600",
  "childPolicy": "TERMINATE",
  "taskPriority": "100",
  "taskStartToCloseTimeout": "600",
  "taskList": {
    "name": "specialTaskList"
  }
  },
  "openCounts": {
  "openTimers": 0,
  "openDecisionTasks": 1,
  "openActivityTasks": 0,
  "openChildWorkflowExecutions": 0
  },
  "executionInfo": {
  "execution": {
    "workflowId": "20110927-T-1",
    "runId": "06b8f87a-24b3-40b6-9ceb-9676f28e9493"
  },
  "startTimestamp": 1326592619.474,
  "executionStatus": "OPEN",
  "workflowType": {
    "version": "1.0",
    "name": "customerOrderWorkflow"
  },
  "cancelRequested": false,
  "tagList": [
    "music purchase",
    "digital",
    "ricoh-the-dog"
  ]
  }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/describeworkflowexecution.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/describeworkflowexecution.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeDomain

DescribeWorkflowType

All content copied from https://docs.aws.amazon.com/.
