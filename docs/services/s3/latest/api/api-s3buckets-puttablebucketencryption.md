# PutTableBucketEncryption

Sets the encryption configuration for a table bucket.

Permissions

You must have the `s3tables:PutTableBucketEncryption` permission to use this operation.

###### Note

If you choose SSE-KMS encryption you must grant the S3 Tables maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption](../userguide/s3-tables-kms-permissions.md) in the _Amazon Simple Storage Service User Guide_.

## Request Syntax

```nohighlight

PUT /buckets/tableBucketARN/encryption HTTP/1.1
Content-type: application/json

{
   "encryptionConfiguration": {
      "kmsKeyArn": "string",
      "sseAlgorithm": "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[tableBucketARN](#API_s3Buckets_PutTableBucketEncryption_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[encryptionConfiguration](#API_s3Buckets_PutTableBucketEncryption_RequestSyntax)**

The encryption configuration to apply to the table bucket.

Type: [EncryptionConfiguration](api-s3buckets-encryptionconfiguration.md) object

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/puttablebucketencryption.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/puttablebucketencryption.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

PutTableBucketMaintenanceConfiguration

All content copied from https://docs.aws.amazon.com/.
