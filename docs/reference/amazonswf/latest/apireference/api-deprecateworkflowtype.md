# DeprecateWorkflowType

Deprecates the specified _workflow type_. After a workflow type has
been deprecated, you cannot create new executions of that type. Executions that were started
before the type was deprecated continues to run. A deprecated workflow type may still be used
when calling visibility actions.

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

- Constrain the following parameters by using a `Condition` element with
the appropriate keys.

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
   "domain": "string",
   "workflowType": {
      "name": "string",
      "version": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[domain](#API_DeprecateWorkflowType_RequestSyntax)**

The name of the domain in which the workflow type is registered.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[workflowType](#API_DeprecateWorkflowType_RequestSyntax)**

The workflow type to deprecate.

Type: [WorkflowType](api-workflowtype.md) object

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

## Examples

### DeprecateWorkflowType Example

This example illustrates one usage of DeprecateWorkflowType.

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
X-Amz-Date: Mon, 16 Jan 2012 05:04:47 GMT
X-Amz-Target: SimpleWorkflowService.DeprecateWorkflowType
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=BGrr1djQvp+YLq3ci2ffpK8KWhZm/PakBL2fFhc3zds=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 102
Pragma: no-cache
Cache-Control: no-cache

{"domain": "867530901",
"workflowType":
{"name": "customerOrderWorkflow",
"version": "1.0"}
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 0
Content-Type: application/json
x-amzn-RequestId: 9c8d6d3b-3fff-11e1-9e8f-57bb03e21482
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/deprecateworkflowtype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/deprecateworkflowtype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeprecateDomain

DescribeActivityType

All content copied from https://docs.aws.amazon.com/.
