# CreateTable

Creates a new table associated with the given namespace in a table bucket. For more information, see [Creating an Amazon S3 table](../userguide/s3-tables-create.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

- You must have the `s3tables:CreateTable` permission to use this operation.

- If you use this operation with the optional `metadata` request parameter you must have the `s3tables:PutTableData` permission.

- If you use this operation with the optional `encryptionConfiguration` request parameter you must have the `s3tables:PutTableEncryption` permission.

- If you use this operation with the `storageClassConfiguration` request parameter, you must have the `s3tables:PutTableStorageClass` permission.

- To create a table with tags, you must have the `s3tables:TagResource` permission in addition to `s3tables:CreateTable` permission.

###### Note

Additionally, If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](../userguide/s3-tables-kms-permissions.md).

## Request Syntax

```nohighlight

PUT /tables/tableBucketARN/namespace HTTP/1.1
Content-type: application/json

{
   "encryptionConfiguration": {
      "kmsKeyArn": "string",
      "sseAlgorithm": "string"
   },
   "format": "string",
   "metadata": { ... },
   "name": "string",
   "storageClassConfiguration": {
      "storageClass": "string"
   },
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[namespace](#API_s3Buckets_CreateTable_RequestSyntax)**

The namespace to associated with the table.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**[tableBucketARN](#API_s3Buckets_CreateTable_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket to create the table in.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[encryptionConfiguration](#API_s3Buckets_CreateTable_RequestSyntax)**

The encryption configuration to use for the table. This configuration specifies the encryption algorithm and, if using SSE-KMS, the KMS key to use for encrypting the table.

###### Note

If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](../userguide/s3-tables-kms-permissions.md).

Type: [EncryptionConfiguration](api-s3buckets-encryptionconfiguration.md) object

Required: No

**[format](#API_s3Buckets_CreateTable_RequestSyntax)**

The format for the table.

Type: String

Valid Values: `ICEBERG`

Required: Yes

**[metadata](#API_s3Buckets_CreateTable_RequestSyntax)**

The metadata for the table.

Type: [TableMetadata](api-s3buckets-tablemetadata.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[name](#API_s3Buckets_CreateTable_RequestSyntax)**

The name for the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**[storageClassConfiguration](#API_s3Buckets_CreateTable_RequestSyntax)**

The storage class configuration for the table. If not specified, the table inherits the storage class configuration from its table bucket. Specify this parameter to override the bucket's default storage class for this table.

Type: [StorageClassConfiguration](api-s3buckets-storageclassconfiguration.md) object

Required: No

**[tags](#API_s3Buckets_CreateTable_RequestSyntax)**

A map of user-defined tags that you would like to apply to the table that you are creating. A tag is a key-value pair that you apply to your resources. Tags can help you organize, track costs for, and control access to resources. For more information, see [Tagging for cost allocation or attribute-based access control (ABAC)](../userguide/tagging.md).

###### Note

You must have the `s3tables:TagResource` permission in addition to `s3tables:CreateTable` permission to create a table with tags.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "tableARN": "string",
   "versionToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[tableARN](#API_s3Buckets_CreateTable_ResponseSyntax)**

The Amazon Resource Name (ARN) of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

**[versionToken](#API_s3Buckets_CreateTable_ResponseSyntax)**

The version token of the table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/createtable.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/createtable.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/createtable.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/createtable.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/createtable.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/createtable.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/createtable.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/createtable.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/createtable.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/createtable.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateNamespace

CreateTableBucket

All content copied from https://docs.aws.amazon.com/.
