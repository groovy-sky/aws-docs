# RegisterDomain

Registers a new domain.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as
follows:

- You cannot use an IAM policy to control domain access for this action. The name of
the domain being registered is available as the resource of this action.

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
   "description": "string",
   "name": "string",
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ],
   "workflowExecutionRetentionPeriodInDays": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[description](#API_RegisterDomain_RequestSyntax)**

A text description of the domain.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**[name](#API_RegisterDomain_RequestSyntax)**

Name of the domain to register. The name must be unique in the region that the domain
is registered in.

The specified string must not start or end with whitespace. It must not contain a
`:` (colon), `/` (slash), `|` (vertical bar), or any
control characters ( `\u0000-\u001f` \| `\u007f-\u009f`). Also, it must
_not_ be the literal string `arn`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**[tags](#API_RegisterDomain_RequestSyntax)**

Tags to be added when registering a domain.

Tags may only contain unicode letters, digits, whitespace, or these symbols: `_ . : / = + - @`.

Type: Array of [ResourceTag](api-resourcetag.md) objects

Required: No

**[workflowExecutionRetentionPeriodInDays](#API_RegisterDomain_RequestSyntax)**

The duration (in days) that records and histories of workflow executions on the domain
should be kept by the service. After the retention period, the workflow execution isn't
available in the results of visibility calls.

If you pass the value `NONE` or `0` (zero), then the workflow
execution history isn't retained. As soon as the workflow execution completes, the execution
record and its history are deleted.

The maximum workflow execution retention period is 90 days. For more information about
Amazon SWF service limits, see: [Amazon SWF Service Limits](../../../../services/amazonswf/latest/developerguide/swf-dg-limits.md) in the
_Amazon SWF Developer Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DomainAlreadyExistsFault**

Returned if the domain already exists. You may get this fault if you are registering a domain that is either already registered or deprecated, or if you undeprecate a domain that is currently registered.

**message**

A description that may help with diagnosing the cause of the fault.

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

**TooManyTagsFault**

You've exceeded the number of tags allowed for a domain.

HTTP Status Code: 400

## Examples

### RegisterDomain Example

This example illustrates one usage of RegisterDomain.

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
X-Amz-Date: Fri, 13 Jan 2012 18:42:12 GMT
X-Amz-Target: SimpleWorkflowService.RegisterDomain
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=tzjkF55lxAxPhzp/BRGFYQRQRq6CqrM254dTDE/EncI=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 91
Pragma: no-cache
Cache-Control: no-cache

{"name": "867530902",
 "description": "music",
 "workflowExecutionRetentionPeriodInDays": "60"}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 0
Content-Type: application/json
x-amzn-RequestId: 4ec4ac3f-3e16-11e1-9b11-7182192d0b57
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/registerdomain.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/registerdomain.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/registerdomain.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/registerdomain.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/registerdomain.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/registerdomain.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/registerdomain.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/registerdomain.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/registerdomain.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/registerdomain.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RegisterActivityType

RegisterWorkflowType

All content copied from https://docs.aws.amazon.com/.
