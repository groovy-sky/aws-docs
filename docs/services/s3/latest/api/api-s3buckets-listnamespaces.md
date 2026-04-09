# ListNamespaces

Lists the namespaces within a table bucket. For more information, see [Table namespaces](../userguide/s3-tables-namespace.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:ListNamespaces` permission to use this operation.

## Request Syntax

```nohighlight

GET /namespaces/tableBucketARN?continuationToken=continuationToken&maxNamespaces=maxNamespaces&prefix=prefix HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[continuationToken](#API_s3Buckets_ListNamespaces_RequestSyntax)**

`ContinuationToken` indicates to Amazon S3 that the list is being continued on
this bucket with a token. `ContinuationToken` is obfuscated and is not a real key.
You can use this `ContinuationToken` for pagination of the list results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[maxNamespaces](#API_s3Buckets_ListNamespaces_RequestSyntax)**

The maximum number of namespaces to return in the list.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[prefix](#API_s3Buckets_ListNamespaces_RequestSyntax)**

The prefix of the namespaces.

Length Constraints: Minimum length of 1. Maximum length of 255.

**[tableBucketARN](#API_s3Buckets_ListNamespaces_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "continuationToken": "string",
   "namespaces": [
      {
         "createdAt": "string",
         "createdBy": "string",
         "namespace": [ "string" ],
         "namespaceId": "string",
         "ownerAccountId": "string",
         "tableBucketId": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[continuationToken](#API_s3Buckets_ListNamespaces_ResponseSyntax)**

The `ContinuationToken` for pagination of the list results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[namespaces](#API_s3Buckets_ListNamespaces_ResponseSyntax)**

A list of namespaces.

Type: Array of [NamespaceSummary](api-s3buckets-namespacesummary.md) objects

## Errors

**AccessDeniedException**

The action cannot be performed because you do not have the required permission.

HTTP Status Code: 403

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/listnamespaces.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/listnamespaces.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetTableStorageClass

ListTableBuckets

All content copied from https://docs.aws.amazon.com/.
