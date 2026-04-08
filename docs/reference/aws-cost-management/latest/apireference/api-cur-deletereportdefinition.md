# DeleteReportDefinition

Deletes the specified report. Any tags associated with the report are also
deleted.

## Request Syntax

```nohighlight

{
   "ReportName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ReportName](#API_cur_DeleteReportDefinition_RequestSyntax)**

The name of the report that you want to delete. The name must be unique, is case sensitive, and can't include spaces.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[0-9A-Za-z!\-_.*\'()]+`

Required: Yes

## Response Syntax

```nohighlight

{
   "ResponseMessage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ResponseMessage](#API_cur_DeleteReportDefinition_ResponseSyntax)**

Whether the deletion was successful or not.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 500

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

## Examples

### The following is a sample request of the DeleteReportDefinition operation.

This example illustrates one usage of DeleteReportDefinition.

#### Sample Request

```

POST / HTTP/1.1
Host: api.cur.<region>.<domain>
x-amz-Date: <Date>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=contenttype;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid,Signature=<Signature>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: AWSOrigamiServiceGateway.DeleteReportDefinition
{
        "ReportName": "ExampleReport"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cur-2017-01-06/deletereportdefinition.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cur-2017-01-06/deletereportdefinition.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AWS Cost and Usage Report

DescribeReportDefinitions
