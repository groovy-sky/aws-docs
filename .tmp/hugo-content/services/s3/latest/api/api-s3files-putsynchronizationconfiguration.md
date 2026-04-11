# PutSynchronizationConfiguration

Creates or updates the synchronization configuration for the specified S3 File System,
including import data rules and expiration data rules.

## Request Syntax

```nohighlight

PUT /file-systems/fileSystemId/synchronization-configuration HTTP/1.1
Content-type: application/json

{
   "expirationDataRules": [
      {
         "daysAfterLastAccess": number
      }
   ],
   "importDataRules": [
      {
         "prefix": "string",
         "sizeLessThan": number,
         "trigger": "string"
      }
   ],
   "latestVersionNumber": number
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[fileSystemId](#API_S3Files_PutSynchronizationConfiguration_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the S3 File System to configure synchronization for.

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[expirationDataRules](#API_S3Files_PutSynchronizationConfiguration_RequestSyntax)**

An array of expiration data rules that control when cached data expires from the file
system.

Type: Array of [ExpirationDataRule](api-s3files-expirationdatarule.md) objects

Array Members: Fixed number of 1 item.

Required: Yes

**[importDataRules](#API_S3Files_PutSynchronizationConfiguration_RequestSyntax)**

An array of import data rules that control how data is imported from S3 into the file
system.

Type: Array of [ImportDataRule](api-s3files-importdatarule.md) objects

Array Members: Minimum number of 1 item. Maximum number of 10 items.

Required: Yes

**[latestVersionNumber](#API_S3Files_PutSynchronizationConfiguration_RequestSyntax)**

The version number of the current synchronization configuration. Omit this value
when creating a synchronization configuration for the first time. For subsequent
updates, provide this value for optimistic concurrency control. If the version number
does not match the current configuration, the request fails with a
`ConflictException`.

Type: Integer

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

**ConflictException**

The request conflicts with the current state of the resource. This can occur when trying to create a resource that already exists or delete a resource that is in use.

**errorCode**

The error code associated with the exception.

**resourceId**

The identifier of the resource that caused the conflict.

**resourceType**

The type of the resource that caused the conflict.

HTTP Status Code: 409

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/putsynchronizationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/putsynchronizationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutFileSystemPolicy

TagResource

All content copied from https://docs.aws.amazon.com/.
