# ListTagsForResource

This operation returns the tags for a resource type.

## Request Syntax

```nohighlight

GET /tags/ResourceArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[ResourceArn](#API_BKS_ListTagsForResource_RequestSyntax)**

The Amazon Resource Name (ARN) that uniquely identifies
the resource.>

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Tags": {
      "string" : "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Tags](#API_BKS_ListTagsForResource_ResponseSyntax)**

List of tags returned by the operation.

Type: String to string map

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

**message**

User does not have sufficient access to perform this action.

HTTP Status Code: 403

**InternalServerException**

An internal server error occurred. Retry your request.

**message**

Unexpected error during processing of request.

**retryAfterSeconds**

Retry the call after number of seconds.

HTTP Status Code: 500

**ResourceNotFoundException**

The resource was not found for this request.

Confirm the resource information, such as the ARN or type is correct
and exists, then retry the request.

**message**

Request references a resource which does not exist.

**resourceId**

Hypothetical identifier of the resource affected.

**resourceType**

Hypothetical type of the resource affected.

HTTP Status Code: 404

**ThrottlingException**

The request was denied due to request throttling.

**message**

Request was unsuccessful due to request throttling.

**quotaCode**

This is the code unique to the originating service with the quota.

**retryAfterSeconds**

Retry the call after number of seconds.

**serviceCode**

This is the code unique to the originating service.

HTTP Status Code: 429

**ValidationException**

The input fails to satisfy the constraints specified by a service.

**message**

The input fails to satisfy the constraints specified by an Amazon service.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for Python](../../../goto/boto3/backupsearch-2018-05-10/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backupsearch-2018-05-10/listtagsforresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListSearchResultExportJobs

StartSearchJob

All content copied from https://docs.aws.amazon.com/.
