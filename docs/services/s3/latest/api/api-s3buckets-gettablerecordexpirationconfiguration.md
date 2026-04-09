# GetTableRecordExpirationConfiguration

Retrieves the expiration configuration settings for records in a table, and the status of the configuration. If the status of the configuration is `enabled`, records expire and are automatically removed from the table after the specified number of days.

Permissions

You must have the `s3tables:GetTableRecordExpirationConfiguration` permission to use this operation.

## Request Syntax

```nohighlight

GET /table-record-expiration?tableArn=tableArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[tableArn](#API_s3Buckets_GetTableRecordExpirationConfiguration_RequestSyntax)**

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
      "settings": {
         "days": number
      },
      "status": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[configuration](#API_s3Buckets_GetTableRecordExpirationConfiguration_ResponseSyntax)**

The record expiration configuration for the table, including the status and retention settings.

Type: [TableRecordExpirationConfigurationValue](api-s3buckets-tablerecordexpirationconfigurationvalue.md) object

## Errors

**BadRequestException**

The request is invalid or malformed.

HTTP Status Code: 400

**ForbiddenException**

The caller isn't authorized to make the request.

HTTP Status Code: 403

**InternalServerErrorException**

The request failed due to an internal server error.

HTTP Status Code: 500

**MethodNotAllowedException**

The requested operation is not allowed on this resource. This may occur when attempting to modify a resource that is managed by a service or has restrictions that prevent the operation.

HTTP Status Code: 405

**NotFoundException**

The request was rejected because the specified resource could not be found.

HTTP Status Code: 404

**TooManyRequestsException**

The limit on the number of requests per second was exceeded.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/gettablerecordexpirationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetTablePolicy

GetTableRecordExpirationJobStatus

All content copied from https://docs.aws.amazon.com/.
