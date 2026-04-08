# DeleteQueryDefinition

Deletes a saved CloudWatch Logs Insights query definition. A query definition contains
details about a saved CloudWatch Logs Insights query.

Each `DeleteQueryDefinition` operation can delete one query definition.

You must have the `logs:DeleteQueryDefinition` permission to be able to perform
this operation.

## Request Syntax

```nohighlight

{
   "queryDefinitionId": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[queryDefinitionId](#API_DeleteQueryDefinition_RequestSyntax)**

The ID of the query definition that you want to delete. You can use [DescribeQueryDefinitions](api-describequerydefinitions.md) to retrieve the IDs of your saved query
definitions.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## Response Syntax

```nohighlight

{
   "success": boolean
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[success](#API_DeleteQueryDefinition_ResponseSyntax)**

A value of TRUE indicates that the operation succeeded. FALSE indicates that the operation
failed.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### Example

This example deletes a query definition.

#### Sample Request

```

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.DeleteQueryDefinition
{
   "queryDefinitionId": "123456ab-12ab-123a-789e-1234567890ab"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
{
   "success": True
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/deletequerydefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/deletequerydefinition.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteMetricFilter

DeleteResourcePolicy
