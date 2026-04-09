# ListAppBundles

Returns a list of app bundles.

## Request Syntax

```nohighlight

GET /appbundles?maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[maxResults](#API_ListAppBundles_RequestSyntax)**

The maximum number of results that are returned per call. You can use
`nextToken` to obtain further pages of results.

This is only an upper limit. The actual number of results returned per call might be
fewer than the specified maximum.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_ListAppBundles_RequestSyntax)**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "appBundleSummaryList": [
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

**[appBundleSummaryList](#API_ListAppBundles_ResponseSyntax)**

Contains a list of app bundle summaries.

Type: Array of [AppBundleSummary](api-appbundlesummary.md) objects

**[nextToken](#API_ListAppBundles_ResponseSyntax)**

If `nextToken` is returned, there are more results available. The value of
`nextToken` is a unique pagination token for each page. Make the call again
using the returned token to retrieve the next page. Keep all other arguments unchanged.
Each pagination token expires after 24 hours. Using an expired pagination token will return
an _HTTP 400 InvalidToken error_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/listappbundles.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/listappbundles.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAppAuthorizations

ListIngestionDestinations

All content copied from https://docs.aws.amazon.com/.
