# GetTable

Gets details about a table. For more information, see [S3 Tables](../userguide/s3-tables-tables.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:GetTable` permission to use this operation.

## Request Syntax

```nohighlight

GET /get-table?name=name&namespace=namespace&tableArn=tableArn&tableBucketARN=tableBucketARN HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_s3Buckets_GetTable_RequestSyntax)**

The name of the table.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

**[namespace](#API_s3Buckets_GetTable_RequestSyntax)**

The name of the namespace the table is associated with.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

**[tableArn](#API_s3Buckets_GetTable_RequestSyntax)**

The Amazon Resource Name (ARN) of the table.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

**[tableBucketARN](#API_s3Buckets_GetTable_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket associated with the table.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "createdAt": "string",
   "createdBy": "string",
   "format": "string",
   "managedByService": "string",
   "managedTableInformation": {
      "replicationInformation": {
         "sourceTableARN": "string"
      }
   },
   "metadataLocation": "string",
   "modifiedAt": "string",
   "modifiedBy": "string",
   "name": "string",
   "namespace": [ "string" ],
   "namespaceId": "string",
   "ownerAccountId": "string",
   "tableARN": "string",
   "tableBucketId": "string",
   "type": "string",
   "versionToken": "string",
   "warehouseLocation": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[createdAt](#API_s3Buckets_GetTable_ResponseSyntax)**

The date and time the table bucket was created at.

Type: Timestamp

**[createdBy](#API_s3Buckets_GetTable_ResponseSyntax)**

The ID of the account that created the table.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9].*`

**[format](#API_s3Buckets_GetTable_ResponseSyntax)**

The format of the table.

Type: String

Valid Values: `ICEBERG`

**[managedByService](#API_s3Buckets_GetTable_ResponseSyntax)**

The service that manages the table.

Type: String

**[managedTableInformation](#API_s3Buckets_GetTable_ResponseSyntax)**

If this table is managed by S3 Tables, contains additional information such as replication details.

Type: [ManagedTableInformation](api-s3buckets-managedtableinformation.md) object

**[metadataLocation](#API_s3Buckets_GetTable_ResponseSyntax)**

The metadata location of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[modifiedAt](#API_s3Buckets_GetTable_ResponseSyntax)**

The date and time the table was last modified on.

Type: Timestamp

**[modifiedBy](#API_s3Buckets_GetTable_ResponseSyntax)**

The ID of the account that last modified the table.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9].*`

**[name](#API_s3Buckets_GetTable_ResponseSyntax)**

The name of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

**[namespace](#API_s3Buckets_GetTable_ResponseSyntax)**

The namespace associated with the table.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

**[namespaceId](#API_s3Buckets_GetTable_ResponseSyntax)**

The unique identifier of the namespace containing this table.

Type: String

**[ownerAccountId](#API_s3Buckets_GetTable_ResponseSyntax)**

The ID of the account that owns the table.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9].*`

**[tableARN](#API_s3Buckets_GetTable_ResponseSyntax)**

The Amazon Resource Name (ARN) of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

**[tableBucketId](#API_s3Buckets_GetTable_ResponseSyntax)**

The unique identifier of the table bucket containing this table.

Type: String

**[type](#API_s3Buckets_GetTable_ResponseSyntax)**

The type of the table.

Type: String

Valid Values: `customer | aws`

**[versionToken](#API_s3Buckets_GetTable_ResponseSyntax)**

The version token of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[warehouseLocation](#API_s3Buckets_GetTable_ResponseSyntax)**

The warehouse location of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/gettable.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/gettable.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/gettable.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/gettable.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/gettable.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/gettable.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/gettable.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/gettable.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/gettable.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/gettable.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetNamespace

GetTableBucket

All content copied from https://docs.aws.amazon.com/.
