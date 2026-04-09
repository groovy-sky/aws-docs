# ListLegalHolds

This action returns metadata about active and previous legal holds.

## Request Syntax

```nohighlight

GET /legal-holds/?maxResults=MaxResults&nextToken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_ListLegalHolds_RequestSyntax)**

The maximum number of resource list items to be returned.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[NextToken](#API_ListLegalHolds_RequestSyntax)**

The next item following a partial list of returned resources. For example, if a request
is made to return `MaxResults` number of resources, `NextToken`
allows you to return more items in your list starting at the location pointed to by the
next token.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "LegalHolds": [
      {
         "CancellationDate": number,
         "CreationDate": number,
         "Description": "string",
         "LegalHoldArn": "string",
         "LegalHoldId": "string",
         "Status": "string",
         "Title": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[LegalHolds](#API_ListLegalHolds_ResponseSyntax)**

This is an array of returned legal holds, both active and previous.

Type: Array of [LegalHold](api-legalhold.md) objects

**[NextToken](#API_ListLegalHolds_ResponseSyntax)**

The next item following a partial list of returned resources. For example, if a request
is made to return `MaxResults` number of resources, `NextToken`
allows you to return more items in your list starting at the location pointed to by the
next token.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValueException**

Indicates that something is wrong with a parameter's value. For example, the value is
out of range.

**Context**

**Type**

HTTP Status Code: 400

**ServiceUnavailableException**

The request failed due to a temporary failure of the server.

**Context**

**Type**

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/listlegalholds.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/listlegalholds.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListIndexedRecoveryPoints

ListProtectedResources

All content copied from https://docs.aws.amazon.com/.
