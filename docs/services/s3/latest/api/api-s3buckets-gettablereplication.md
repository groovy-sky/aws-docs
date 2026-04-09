# GetTableReplication

Retrieves the replication configuration for a specific table.

Permissions

You must have the `s3tables:GetTableReplication` permission to use this operation.

## Request Syntax

```nohighlight

GET /table-replication?tableArn=tableArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[tableArn](#API_s3Buckets_GetTableReplication_RequestSyntax)**

The Amazon Resource Name (ARN) of the table.

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "configuration": {
      "role": "string",
      "rules": [
         {
            "destinations": [
               {
                  "destinationTableBucketARN": "string"
               }
            ]
         }
      ]
   },
   "versionToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configuration](#API_s3Buckets_GetTableReplication_ResponseSyntax)**

The replication configuration for the table, including the IAM role and replication rules.

Type: [TableReplicationConfiguration](api-s3buckets-tablereplicationconfiguration.md) object

**[versionToken](#API_s3Buckets_GetTableReplication_ResponseSyntax)**

A version token that represents the current state of the table's replication configuration. Use this token when updating the configuration to ensure consistency.

Type: String

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/gettablereplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/gettablereplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetTableRecordExpirationJobStatus

GetTableReplicationStatus

All content copied from https://docs.aws.amazon.com/.
