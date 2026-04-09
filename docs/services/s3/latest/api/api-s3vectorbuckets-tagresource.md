# TagResource

Applies one or more user-defined tags to an Amazon S3 Vectors resource or updates existing
tags. Each tag is a label consisting of a key and value pair. Tags can help you organize,
track costs for, and control access to your resources. You can add up to 50 tags for each
resource.

###### Note

For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources](../userguide/tagging.md#manage-tags).

Permissions

For vector buckets and vector indexes, you must have the `s3vectors:TagResource` permission to use this operation.

## Request Syntax

```nohighlight

POST /tags/resourceArn HTTP/1.1
Content-type: application/json

{
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_S3VectorBuckets_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the Amazon S3 Vectors resource that you're applying tags to. The tagged resource can be a vector bucket or a vector index.

Length Constraints: Minimum length of 0. Maximum length of 1011.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[tags](#API_S3VectorBuckets_TagResource_RequestSyntax)**

The user-defined tag that you want to add to the specified S3 Vectors resource. For more information, see [Tagging for cost allocation or attribute-based access control (ABAC)](../userguide/tagging.md).

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

**AccessDeniedException**

Access denied.

HTTP Status Code: 403

**ConflictException**

The request failed because a vector bucket name or a vector index name already exists.
Vector bucket names must be unique within your AWS account for each AWS Region. Vector
index names must be unique within your vector bucket. Choose a different vector bucket name
or vector index name, and try again.

HTTP Status Code: 409

**InternalServerException**

The request failed due to an internal server error.

HTTP Status Code: 500

**NotFoundException**

The request was rejected because the specified resource can't be found.

HTTP Status Code: 404

**RequestTimeoutException**

The request timed out. Retry your request.

HTTP Status Code: 408

**ServiceUnavailableException**

The service is unavailable. Wait briefly and retry your request. If it continues to
fail, increase your waiting time between retries.

HTTP Status Code: 503

**TooManyRequestsException**

The request was denied due to request throttling.

HTTP Status Code: 429

**ValidationException**

The requested action isn't valid.

**fieldList**

A list of specific validation failures that are encountered during input processing. Each entry
in the list contains a path to the field that failed validation and a detailed message that
explains why the validation failed. To satisfy multiple constraints, a field can appear multiple times in this list if it
failed. You can use the information to identify and fix
the specific validation issues in your request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for Python](../../../goto/boto3/s3vectors-2025-07-15/tagresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryVectors

UntagResource

All content copied from https://docs.aws.amazon.com/.
