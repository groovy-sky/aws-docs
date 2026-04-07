# UpdateGlobalTable

Adds or removes replicas in the specified global table. The global table must already
exist to be able to use this operation. Any replica to be added must be empty, have the
same name as the global table, have the same key schema, have DynamoDB Streams enabled,
and have the same provisioned and maximum write capacity units.

###### Important

This documentation is for version 2017.11.29 (Legacy) of global tables, which should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)](../../../../services/dynamodb/latest/developerguide/globaltables.md) when possible, because it provides greater flexibility, higher efficiency, and consumes less write capacity than 2017.11.29 (Legacy).

To determine which version you're using, see [Determining the global table version you are using](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.DetermineVersion.html). To update existing global tables from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_upgrade.html).

###### Note

If you are using global tables [Version\
2019.11.21](../../../../services/dynamodb/latest/developerguide/globaltables.md) (Current) you can use [UpdateTable](api-updatetable.md) instead.

Although you can use `UpdateGlobalTable` to add replicas and remove
replicas in a single request, for simplicity we recommend that you issue separate
requests for adding or removing replicas.

If global secondary indexes are specified, then the following conditions must also be
met:

- The global secondary indexes must have the same name.

- The global secondary indexes must have the same hash key and sort key (if
present).

- The global secondary indexes must have the same provisioned and maximum write
capacity units.

## Request Syntax

```nohighlight

{
   "GlobalTableName": "string",
   "ReplicaUpdates": [
      {
         "Create": {
            "RegionName": "string"
         },
         "Delete": {
            "RegionName": "string"
         }
      }
   ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[GlobalTableName](#API_UpdateGlobalTable_RequestSyntax)**

The global table name.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**[ReplicaUpdates](#API_UpdateGlobalTable_RequestSyntax)**

A list of Regions that should be added or removed from the global table.

Type: Array of [ReplicaUpdate](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaUpdate.html) objects

Required: Yes

## Response Syntax

```nohighlight

{
   "GlobalTableDescription": {
      "CreationDateTime": number,
      "GlobalTableArn": "string",
      "GlobalTableName": "string",
      "GlobalTableStatus": "string",
      "ReplicationGroup": [
         {
            "GlobalSecondaryIndexes": [
               {
                  "IndexName": "string",
                  "OnDemandThroughputOverride": {
                     "MaxReadRequestUnits": number
                  },
                  "ProvisionedThroughputOverride": {
                     "ReadCapacityUnits": number
                  },
                  "WarmThroughput": {
                     "ReadUnitsPerSecond": number,
                     "Status": "string",
                     "WriteUnitsPerSecond": number
                  }
               }
            ],
            "GlobalTableSettingsReplicationMode": "string",
            "KMSMasterKeyId": "string",
            "OnDemandThroughputOverride": {
               "MaxReadRequestUnits": number
            },
            "ProvisionedThroughputOverride": {
               "ReadCapacityUnits": number
            },
            "RegionName": "string",
            "ReplicaArn": "string",
            "ReplicaInaccessibleDateTime": number,
            "ReplicaStatus": "string",
            "ReplicaStatusDescription": "string",
            "ReplicaStatusPercentProgress": "string",
            "ReplicaTableClassSummary": {
               "LastUpdateDateTime": number,
               "TableClass": "string"
            },
            "WarmThroughput": {
               "ReadUnitsPerSecond": number,
               "Status": "string",
               "WriteUnitsPerSecond": number
            }
         }
      ]
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GlobalTableDescription](#API_UpdateGlobalTable_ResponseSyntax)**

Contains the details of the global table.

Type: [GlobalTableDescription](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**GlobalTableNotFoundException**

The specified global table does not exist.

HTTP Status Code: 400

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ReplicaAlreadyExistsException**

The specified replica is already part of the global table.

HTTP Status Code: 400

**ReplicaNotFoundException**

The specified replica is no longer part of the global table.

HTTP Status Code: 400

**TableNotFoundException**

A source table with the name `TableName` does not currently exist within
the subscriber's account or the subscriber is operating in the wrong AWS
Region.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/UpdateGlobalTable)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/UpdateGlobalTable)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateContributorInsights

UpdateGlobalTableSettings
