# DeprecateActivityType

Deprecates the specified _activity type_. After an activity type has
been deprecated, you cannot create new tasks of that activity type. Tasks of this type that
were scheduled before the type was deprecated continue to run.

**Access Control**

You can use IAM policies to control this action's access to Amazon SWF resources as
follows:

- Use a `Resource` element with the domain name to limit the action to
only specified domains.

- Use an `Action` element to allow or deny permission to call this
action.

- Constrain the following parameters by using a `Condition` element with
the appropriate keys.

- `activityType.name`: String constraint. The key is
`swf:activityType.name`.

- `activityType.version`: String constraint. The key is
`swf:activityType.version`.

If the caller doesn't have sufficient permissions to invoke the action, or the
parameter values fall outside the specified constraints, the action fails. The associated
event attribute's `cause` parameter is set to `OPERATION_NOT_PERMITTED`.
For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF\
Workflows](../../../../services/amazonswf/latest/developerguide/swf-dev-iam.md) in the _Amazon SWF Developer Guide_.

## Request Syntax

```nohighlight

{
   "activityType": {
      "name": "string",
      "version": "string"
   },
   "domain": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[activityType](#API_DeprecateActivityType_RequestSyntax)**

The activity type to deprecate.

Type: [ActivityType](api-activitytype.md) object

Required: Yes

**[domain](#API_DeprecateActivityType_RequestSyntax)**

The name of the domain in which the activity type is registered.

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

### DeprecateActivityType Example

This example illustrates one usage of DeprecateActivityType.

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
X-Amz-Date: Mon, 16 Jan 2012 05:01:06 GMT
X-Amz-Target: SimpleWorkflowService.DeprecateActivityType
Content-Encoding: amz-1.0
X-Amzn-Authorization: AWS3 AWSAccessKeyId=AWS_ACCESS_KEY_ID_REDACTED,Algorithm=HmacSHA256,SignedHeaders=Host;X-Amz-Date;X-Amz-Target;Content-Encoding,Signature=iX/mNMtNH6IaSNwfZq9hHOhDlLnp7buuj9tO93kRIrQ=
Referer: http://swf.us-east-1.amazonaws.com/explorer/index.html
Content-Length: 95
Pragma: no-cache
Cache-Control: no-cache

{"domain": "867530901",
"activityType":
{"name": "activityVerify",
"version": "1.0"}
}
```

#### Sample Response

```

HTTP/1.1 200 OK
Content-Length: 0
Content-Type: application/json
x-amzn-RequestId: 191ee17e-3fff-11e1-a23a-99d60383ae71
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/deprecateactivitytype.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/deprecateactivitytype.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteWorkflowType

DeprecateDomain

All content copied from https://docs.aws.amazon.com/.
