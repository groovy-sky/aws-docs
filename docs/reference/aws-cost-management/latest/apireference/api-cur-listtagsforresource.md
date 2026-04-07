# ListTagsForResource

Lists the tags associated with the specified report definition.

## Request Syntax

```nohighlight

{
   "ReportName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ReportName](#API_cur_ListTagsForResource_RequestSyntax)**

The report name of the report definition that tags are to be returned for.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `[0-9A-Za-z!\-_.*\'()]+`

Required: Yes

## Response Syntax

```nohighlight

{
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Tags](#API_cur_ListTagsForResource_ResponseSyntax)**

The tags assigned to the report definition resource.

Type: Array of [Tag](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_cur_Tag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalErrorException**

An error on the server occurred during the processing of your request. Try again later.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified report ( `ReportName`) in the request doesn't exist.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

**ValidationException**

The input fails to satisfy the constraints specified by an AWS service.

**Message**

A message to show the detail of the exception.

HTTP Status Code: 400

## Examples

### The following is a sample request of the ListTagsForResource operation.

This example illustrates one usage of ListTagsForResource.

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
X-Amz-Target: AWSOrigamiServiceGateway.ListTagsForResource
{
  "ReportName": "ExampleReport",
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/cur-2017-01-06/ListTagsForResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cur-2017-01-06/ListTagsForResource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeReportDefinitions

ModifyReportDefinition
