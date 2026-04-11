# CreateIngestionDestination

Creates an ingestion destination, which specifies how an application's ingested data is
processed by AWS AppFabric and where it's delivered.

## Request Syntax

```nohighlight

POST /appbundles/appBundleIdentifier/ingestions/ingestionIdentifier/ingestiondestinations HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
   "destinationConfiguration": { ... },
   "processingConfiguration": { ... },
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[appBundleIdentifier](#API_CreateIngestionDestination_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle
to use for the request.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[ingestionIdentifier](#API_CreateIngestionDestination_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to
use for the request.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_CreateIngestionDestination_RequestSyntax)**

Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency
of the request. This lets you safely retry the request without accidentally performing the
same operation a second time. Passing the same value to a later call to an operation
requires that you also pass the same value for all other parameters. We recommend that you
use a [UUID type of\
value](https://wikipedia.org/wiki/Universally_unique_identifier).

If you don't provide this value, then AWS generates a random one for
you.

If you retry the operation with the same `ClientToken`, but with different
parameters, the retry fails with an `IdempotentParameterMismatch` error.

Type: String

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**[destinationConfiguration](#API_CreateIngestionDestination_RequestSyntax)**

Contains information about the destination of ingested data.

Type: [DestinationConfiguration](api-destinationconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**[processingConfiguration](#API_CreateIngestionDestination_RequestSyntax)**

Contains information about how ingested data is processed.

Type: [ProcessingConfiguration](api-processingconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

**[tags](#API_CreateIngestionDestination_RequestSyntax)**

A map of the key-value pairs of the tag or tags to assign to the resource.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 201
Content-type: application/json

{
   "ingestionDestination": {
      "arn": "string",
      "createdAt": "string",
      "destinationConfiguration": { ... },
      "ingestionArn": "string",
      "processingConfiguration": { ... },
      "status": "string",
      "statusReason": "string",
      "updatedAt": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 201 response.

The following data is returned in JSON format by the service.

**[ingestionDestination](#API_CreateIngestionDestination_ResponseSyntax)**

Contains information about an ingestion destination.

Type: [IngestionDestination](api-ingestiondestination.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You are not authorized to perform this operation.

HTTP Status Code: 403

**ConflictException**

The request has created a conflict. Check the request parameters and try again.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 409

**InternalServerException**

The request processing has failed because of an unknown error, exception, or failure
with an internal server.

**retryAfterSeconds**

The period of time after which you should retry your request.

HTTP Status Code: 500

**ServiceQuotaExceededException**

The request exceeds a service quota.

**quotaCode**

The code for the quota exceeded.

**resourceId**

The resource ID.

**resourceType**

The resource type.

**serviceCode**

The code of the service.

HTTP Status Code: 402

**ThrottlingException**

The request rate exceeds the limit.

**quotaCode**

The code for the quota exceeded.

**retryAfterSeconds**

The period of time after which you should retry your request.

**serviceCode**

The code of the service.

HTTP Status Code: 429

**ValidationException**

The request has invalid or missing parameters.

**fieldList**

The field list.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/createingestiondestination.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/createingestiondestination.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateIngestion

DeleteAppAuthorization

All content copied from https://docs.aws.amazon.com/.
