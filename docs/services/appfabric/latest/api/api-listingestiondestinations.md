# ListIngestionDestinations

Returns a list of all ingestion destinations configured for an ingestion.

## Request Syntax

```nohighlight

GET /appbundles/appBundleIdentifier/ingestions/ingestionIdentifier/ingestiondestinations?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[appBundleIdentifier](#API_ListIngestionDestinations_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the app bundle
to use for the request.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[ingestionIdentifier](#API_ListIngestionDestinations_RequestSyntax)**

The Amazon Resource Name (ARN) or Universal Unique Identifier (UUID) of the ingestion to
use for the request.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

**[maxResults](#API_ListIngestionDestinations_RequestSyntax)**

The maximum number of results that are returned per call. You can use
`nextToken` to obtain further pages of results.

This is only an upper limit. The actual number of results returned per call might be
fewer than the specified maximum.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_ListIngestionDestinations_RequestSyntax)**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "ingestionDestinations": [
      {
         "arn": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ingestionDestinations](#API_ListIngestionDestinations_ResponseSyntax)**

Contains a list of ingestion destination summaries.

Type: Array of [IngestionDestinationSummary](api-ingestiondestinationsummary.md) objects

**[nextToken](#API_ListIngestionDestinations_ResponseSyntax)**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You are not authorized to perform this operation.

HTTP Status Code: 403

**InternalServerException**

The request processing has failed because of an unknown error, exception, or failure
with an internal server.

**retryAfterSeconds**

The period of time after which you should retry your request.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource does not exist.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/listingestiondestinations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/listingestiondestinations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAppBundles

ListIngestions

All content copied from https://docs.aws.amazon.com/.
