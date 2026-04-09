# TagResource

Applies one or more user-defined tags to an Amazon S3 Tables resource or updates existing tags. Each tag is a label consisting of a key and value pair. Tags can help you organize, track costs for, and control access to your resources. You can add up to 50 tags for each S3 resource.

###### Note

For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources](../userguide/tagging.md#manage-tags).

Permissions

For tables and table buckets, you must have the `s3tables:TagResource` permission to use this operation.

## Request Syntax

```nohighlight

POST /tag/resourceArn HTTP/1.1
Content-type: application/json

{
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_s3Buckets_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the Amazon S3 Tables resource that you're applying tags to. The tagged resource can be a table bucket or a table. For a list of all S3 resources that support tagging, see [Managing tags for Amazon S3 resources](../userguide/tagging.md#manage-tags).

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/.+`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[tags](#API_s3Buckets_TagResource_RequestSyntax)**

The user-defined tag that you want to add to the specified S3 Tables resource. For more information, see [Tagging for cost allocation or attribute-based access control (ABAC)](../userguide/tagging.md).

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

**BadRequestException**

The request is invalid or malformed.

HTTP Status Code: 400

**ConflictException**

The request failed because there is a conflict with a previous write. You can retry the
request.

HTTP Status Code: 409

**ForbiddenException**

The caller isn't authorized to make the request.

HTTP Status Code: 403

**InternalServerErrorException**

The request failed due to an internal server error.

HTTP Status Code: 500

**NotFoundException**

The request was rejected because the specified resource could not be found.

HTTP Status Code: 404

**TooManyRequestsException**

The limit on the number of requests per second was exceeded.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/tagresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RenameTable

UntagResource

All content copied from https://docs.aws.amazon.com/.
