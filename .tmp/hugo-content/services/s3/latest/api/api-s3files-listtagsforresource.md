# ListTagsForResource

Lists all tags for S3 Files resources.

## Request Syntax

```nohighlight

GET /resource-tags/resourceId?MaxResults=maxResults&NextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_S3Files_ListTagsForResource_RequestSyntax)**

The maximum number of tags to return in a single response.

Valid Range: Minimum value of 1. Maximum value of 50.

**[nextToken](#API_S3Files_ListTagsForResource_RequestSyntax)**

A pagination token returned from a previous call to continue listing tags.

**[resourceId](#API_S3Files_ListTagsForResource_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the resource to list tags for.

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}(/access-point/fsap-[0-9a-f]{17,40})?|fs(ap)?-[0-9a-f]{17,40})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "nextToken": "string",
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[nextToken](#API_S3Files_ListTagsForResource_ResponseSyntax)**

A pagination token to use in a subsequent request if more results are
available.

Type: String

**[tags](#API_S3Files_ListTagsForResource_ResponseSyntax)**

An array of tags associated with the resource.

Type: Array of [Tag](api-s3files-tag.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

## Errors

**InternalServerException**

An internal server error occurred. Retry your request.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource was not found. Verify that the resource exists and that you have permission to access it.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 404

**ThrottlingException**

The request was throttled. Retry your request using exponential backoff.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 429

**ValidationException**

The input parameters are not valid. Check the parameter values and try again.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listtagsforresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListMountTargets

PutFileSystemPolicy

All content copied from https://docs.aws.amazon.com/.
