# DescribeExportTasks

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Retrieve status of one or more export tasks. You can retrieve the status of up to 100
export tasks.

## Request Syntax

```nohighlight

{
   "exportIds": [ "string" ],
   "filters": [
      {
         "condition": "string",
         "name": "string",
         "values": [ "string" ]
      }
   ],
   "maxResults": number,
   "nextToken": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[exportIds](#API_DescribeExportTasks_RequestSyntax)**

One or more unique identifiers used to query the status of an export request.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: No

**[filters](#API_DescribeExportTasks_RequestSyntax)**

One or more filters.

- `AgentId` \- ID of the agent whose collected data will be
exported

Type: Array of [ExportFilter](api-exportfilter.md) objects

Required: No

**[maxResults](#API_DescribeExportTasks_RequestSyntax)**

The maximum number of volume results returned by `DescribeExportTasks` in
paginated output. When this parameter is used, `DescribeExportTasks` only returns
`maxResults` results in a single page along with a `nextToken`
response element.

Type: Integer

Required: No

**[nextToken](#API_DescribeExportTasks_RequestSyntax)**

The `nextToken` value returned from a previous paginated
`DescribeExportTasks` request where `maxResults` was used and the
results exceeded the value of that parameter. Pagination continues from the end of the
previous results that returned the `nextToken` value. This value is null when there
are no more results to return.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "exportsInfo": [
      {
         "configurationsDownloadUrl": "string",
         "exportId": "string",
         "exportRequestTime": number,
         "exportStatus": "string",
         "isTruncated": boolean,
         "requestedEndTime": number,
         "requestedStartTime": number,
         "statusMessage": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[exportsInfo](#API_DescribeExportTasks_ResponseSyntax)**

Contains one or more sets of export request details. When the status of a request is
`SUCCEEDED`, the response includes a URL for an Amazon S3 bucket where you can
view the data in a CSV file.

Type: Array of [ExportInfo](api-exportinfo.md) objects

**[nextToken](#API_DescribeExportTasks_ResponseSyntax)**

The `nextToken` value to include in a future
`DescribeExportTasks` request. When the results of a
`DescribeExportTasks` request exceed `maxResults`, this value can be
used to retrieve the next page of results. This value is null when there are no more results
to return.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AuthorizationErrorException**

The user does not have permission to perform the action. Check the IAM policy
associated with this user.

HTTP Status Code: 400

**HomeRegionNotSetException**

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

The home Region is not set. Set the home Region to continue.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters are not valid. Verify the parameters and try again.

HTTP Status Code: 400

**InvalidParameterValueException**

The value of one or more parameters are either invalid or out of range. Verify the
parameter values and try again.

HTTP Status Code: 400

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

In the following example or examples, the Authorization header contents
( `AUTHPARAMS`) must be replaced with an AWS Signature Version 4 signature.
For more information about creating these signatures, see [Signature Version 4\
Signing Process](../../../../general/general/latest/gr/signature-version-4.md) in the _AWS General Reference_.

You only need to learn how to sign HTTP requests if you intend to manually create
them. When you use the [AWS Command Line Interface (AWS\
CLI)](http://aws.amazon.com/cli) or one of the [AWS SDKs](http://aws.amazon.com/tools) to make
requests to AWS, these tools automatically sign the requests for you with the access key
that you specify when you configure the tools. When you use these tools, you don't need
to learn how to sign requests yourself.

### Retrieve status of two specified export tasks

This example illustrates one usage of DescribeExportTasks.

#### Sample Request

```

POST / HTTP/1.1
Host: discovery.us-west-2.amazonaws.com
Accept-Encoding: identity
Content-Length: 109
X-Amz-Target: AWSPoseidonService_V2015_11_01.DescribeExportTasks
X-Amz-Date: 20170308T232123Z
Content-Type: application/x-amz-json-1.1
Authorization: AUTHPARAMS

{
   "exportIds":[
      "export-123a4b56-7c89-01d2-3ef4-example5678f",
      "export-654a3b21-7c89-01d2-3ef4-example8765f"
   ]
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 123a4b56-7c89-01d2-3ef4-example5678f
Content-Type: application/x-amz-json-1.1
Content-Length: 1140
Date: Wed, 08 Mar 2017 23:21:25 GMT

{
   "exportsInfo":[
      {
         "configurationsDownloadUrl":"[URL]",
         "exportId":"export-123a4b56-7c89-01d2-3ef4-example5678f",
         "exportRequestTime":1.489001254713E9,
         "exportStatus":"SUCCEEDED",
         "statusMessage":"Data export ran successfully and is accessible from the download URL. The URL will expire in 24 hours. The export data expires in 10 days."
      },
      {
         "configurationsDownloadUrl":"[URL]",
         "exportId":"export-654a3b21-7c89-01d2-3ef4-example8765f",
         "exportRequestTime":1.488920016713E9,
         "exportStatus":"SUCCEEDED",
         "statusMessage":"Data export ran successfully and is accessible from the download URL. The URL will expire in 24 hours. The export data expires in 10 days."
      }
   ],
   "nextToken":""
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/discovery-2015-11-01/describeexporttasks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/describeexporttasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeExportConfigurations

DescribeImportTasks
