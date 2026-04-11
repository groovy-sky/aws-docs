# ListTieringConfigurations

Returns a list of tiering configurations.

## Request Syntax

```nohighlight

GET /tiering-configurations/?maxResults=MaxResults&nextToken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_ListTieringConfigurations_RequestSyntax)**

The maximum number of items to be returned.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[NextToken](#API_ListTieringConfigurations_RequestSyntax)**

The next item following a partial list of returned items. For example, if a request
is made to return `MaxResults` number of items, `NextToken`
allows you to return more items in your list starting at the location pointed to by the
next token.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "NextToken": "string",
   "TieringConfigurations": [
      {
         "BackupVaultName": "string",
         "CreationTime": number,
         "LastUpdatedTime": number,
         "TieringConfigurationArn": "string",
         "TieringConfigurationName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListTieringConfigurations_ResponseSyntax)**

The next item following a partial list of returned items. For example, if a request
is made to return `MaxResults` number of items, `NextToken`
allows you to return more items in your list starting at the location pointed to by the
next token.

Type: String

**[TieringConfigurations](#API_ListTieringConfigurations_ResponseSyntax)**

An array of tiering configurations returned by the `ListTieringConfigurations` call.

Type: Array of [TieringConfigurationsListMember](api-tieringconfigurationslistmember.md) objects

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/listtieringconfigurations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/listtieringconfigurations.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTags

PutBackupVaultAccessPolicy

All content copied from https://docs.aws.amazon.com/.
