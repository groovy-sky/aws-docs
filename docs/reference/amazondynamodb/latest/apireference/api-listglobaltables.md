# ListGlobalTables

Lists all global tables that have a replica in the specified Region.

###### Important

This documentation is for version 2017.11.29 (Legacy) of global tables, which should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)](../../../../services/dynamodb/latest/developerguide/globaltables.md) when possible, because it provides greater flexibility, higher efficiency, and consumes less write capacity than 2017.11.29 (Legacy).

To determine which version you're using, see [Determining the global table version you are using](../../../../services/dynamodb/latest/developerguide/globaltables-determineversion.md). To update existing global tables from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables](../../../../services/dynamodb/latest/developerguide/v2globaltables-upgrade.md).

## Request Syntax

```nohighlight

{
   "ExclusiveStartGlobalTableName": "string",
   "Limit": number,
   "RegionName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ExclusiveStartGlobalTableName](#API_ListGlobalTables_RequestSyntax)**

The first global table name that this operation will evaluate.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**[Limit](#API_ListGlobalTables_RequestSyntax)**

The maximum number of table names to return, if the parameter is not specified
DynamoDB defaults to 100.

If the number of global tables DynamoDB finds reaches this limit, it stops the
operation and returns the table names collected up to that point, with a table name in
the `LastEvaluatedGlobalTableName` to apply in a subsequent operation to the
`ExclusiveStartGlobalTableName` parameter.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[RegionName](#API_ListGlobalTables_RequestSyntax)**

Lists the global tables in a specific Region.

Type: String

Required: No

## Response Syntax

```nohighlight

{
   "GlobalTables": [
      {
         "GlobalTableName": "string",
         "ReplicationGroup": [
            {
               "RegionName": "string"
            }
         ]
      }
   ],
   "LastEvaluatedGlobalTableName": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GlobalTables](#API_ListGlobalTables_ResponseSyntax)**

List of global table names.

Type: Array of [GlobalTable](api-globaltable.md) objects

**[LastEvaluatedGlobalTableName](#API_ListGlobalTables_ResponseSyntax)**

Last evaluated global table name.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/listglobaltables.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/listglobaltables.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListExports

ListImports
