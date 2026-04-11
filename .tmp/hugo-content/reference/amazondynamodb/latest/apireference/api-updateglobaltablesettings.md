# UpdateGlobalTableSettings

Updates settings for a global table.

###### Important

This documentation is for version 2017.11.29 (Legacy) of global tables, which should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)](../../../../services/dynamodb/latest/developerguide/globaltables.md) when possible, because it provides greater flexibility, higher efficiency, and consumes less write capacity than 2017.11.29 (Legacy).

To determine which version you're using, see [Determining the global table version you are using](../../../../services/dynamodb/latest/developerguide/globaltables-determineversion.md). To update existing global tables from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables](../../../../services/dynamodb/latest/developerguide/v2globaltables-upgrade.md).

## Request Syntax

```nohighlight

{
   "GlobalTableBillingMode": "string",
   "GlobalTableGlobalSecondaryIndexSettingsUpdate": [
      {
         "IndexName": "string",
         "ProvisionedWriteCapacityAutoScalingSettingsUpdate": {
            "AutoScalingDisabled": boolean,
            "AutoScalingRoleArn": "string",
            "MaximumUnits": number,
            "MinimumUnits": number,
            "ScalingPolicyUpdate": {
               "PolicyName": "string",
               "TargetTrackingScalingPolicyConfiguration": {
                  "DisableScaleIn": boolean,
                  "ScaleInCooldown": number,
                  "ScaleOutCooldown": number,
                  "TargetValue": number
               }
            }
         },
         "ProvisionedWriteCapacityUnits": number
      }
   ],
   "GlobalTableName": "string",
   "GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate": {
      "AutoScalingDisabled": boolean,
      "AutoScalingRoleArn": "string",
      "MaximumUnits": number,
      "MinimumUnits": number,
      "ScalingPolicyUpdate": {
         "PolicyName": "string",
         "TargetTrackingScalingPolicyConfiguration": {
            "DisableScaleIn": boolean,
            "ScaleInCooldown": number,
            "ScaleOutCooldown": number,
            "TargetValue": number
         }
      }
   },
   "GlobalTableProvisionedWriteCapacityUnits": number,
   "ReplicaSettingsUpdate": [
      {
         "RegionName": "string",
         "ReplicaGlobalSecondaryIndexSettingsUpdate": [
            {
               "IndexName": "string",
               "ProvisionedReadCapacityAutoScalingSettingsUpdate": {
                  "AutoScalingDisabled": boolean,
                  "AutoScalingRoleArn": "string",
                  "MaximumUnits": number,
                  "MinimumUnits": number,
                  "ScalingPolicyUpdate": {
                     "PolicyName": "string",
                     "TargetTrackingScalingPolicyConfiguration": {
                        "DisableScaleIn": boolean,
                        "ScaleInCooldown": number,
                        "ScaleOutCooldown": number,
                        "TargetValue": number
                     }
                  }
               },
               "ProvisionedReadCapacityUnits": number
            }
         ],
         "ReplicaProvisionedReadCapacityAutoScalingSettingsUpdate": {
            "AutoScalingDisabled": boolean,
            "AutoScalingRoleArn": "string",
            "MaximumUnits": number,
            "MinimumUnits": number,
            "ScalingPolicyUpdate": {
               "PolicyName": "string",
               "TargetTrackingScalingPolicyConfiguration": {
                  "DisableScaleIn": boolean,
                  "ScaleInCooldown": number,
                  "ScaleOutCooldown": number,
                  "TargetValue": number
               }
            }
         },
         "ReplicaProvisionedReadCapacityUnits": number,
         "ReplicaTableClass": "string"
      }
   ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[GlobalTableName](#API_UpdateGlobalTableSettings_RequestSyntax)**

The name of the global table

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

**[GlobalTableBillingMode](#API_UpdateGlobalTableSettings_RequestSyntax)**

The billing mode of the global table. If `GlobalTableBillingMode` is not
specified, the global table defaults to `PROVISIONED` capacity billing
mode.

- `PROVISIONED` \- We recommend using `PROVISIONED` for
predictable workloads. `PROVISIONED` sets the billing mode to [Provisioned capacity mode](../../../../services/dynamodb/latest/developerguide/provisioned-capacity-mode.md).

- `PAY_PER_REQUEST` \- We recommend using `PAY_PER_REQUEST`
for unpredictable workloads. `PAY_PER_REQUEST` sets the billing mode
to [On-demand capacity mode](../../../../services/dynamodb/latest/developerguide/on-demand-capacity-mode.md).

Type: String

Valid Values: `PROVISIONED | PAY_PER_REQUEST`

Required: No

**[GlobalTableGlobalSecondaryIndexSettingsUpdate](#API_UpdateGlobalTableSettings_RequestSyntax)**

Represents the settings of a global secondary index for a global table that will be
modified.

Type: Array of [GlobalTableGlobalSecondaryIndexSettingsUpdate](api-globaltableglobalsecondaryindexsettingsupdate.md) objects

Array Members: Minimum number of 1 item. Maximum number of 20 items.

Required: No

**[GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate](#API_UpdateGlobalTableSettings_RequestSyntax)**

Auto scaling settings for managing provisioned write capacity for the global
table.

Type: [AutoScalingSettingsUpdate](api-autoscalingsettingsupdate.md) object

Required: No

**[GlobalTableProvisionedWriteCapacityUnits](#API_UpdateGlobalTableSettings_RequestSyntax)**

The maximum number of writes consumed per second before DynamoDB returns a
`ThrottlingException.`

Type: Long

Valid Range: Minimum value of 1.

Required: No

**[ReplicaSettingsUpdate](#API_UpdateGlobalTableSettings_RequestSyntax)**

Represents the settings for a global table in a Region that will be modified.

Type: Array of [ReplicaSettingsUpdate](api-replicasettingsupdate.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Required: No

## Response Syntax

```nohighlight

{
   "GlobalTableName": "string",
   "ReplicaSettings": [
      {
         "RegionName": "string",
         "ReplicaBillingModeSummary": {
            "BillingMode": "string",
            "LastUpdateToPayPerRequestDateTime": number
         },
         "ReplicaGlobalSecondaryIndexSettings": [
            {
               "IndexName": "string",
               "IndexStatus": "string",
               "ProvisionedReadCapacityAutoScalingSettings": {
                  "AutoScalingDisabled": boolean,
                  "AutoScalingRoleArn": "string",
                  "MaximumUnits": number,
                  "MinimumUnits": number,
                  "ScalingPolicies": [
                     {
                        "PolicyName": "string",
                        "TargetTrackingScalingPolicyConfiguration": {
                           "DisableScaleIn": boolean,
                           "ScaleInCooldown": number,
                           "ScaleOutCooldown": number,
                           "TargetValue": number
                        }
                     }
                  ]
               },
               "ProvisionedReadCapacityUnits": number,
               "ProvisionedWriteCapacityAutoScalingSettings": {
                  "AutoScalingDisabled": boolean,
                  "AutoScalingRoleArn": "string",
                  "MaximumUnits": number,
                  "MinimumUnits": number,
                  "ScalingPolicies": [
                     {
                        "PolicyName": "string",
                        "TargetTrackingScalingPolicyConfiguration": {
                           "DisableScaleIn": boolean,
                           "ScaleInCooldown": number,
                           "ScaleOutCooldown": number,
                           "TargetValue": number
                        }
                     }
                  ]
               },
               "ProvisionedWriteCapacityUnits": number
            }
         ],
         "ReplicaProvisionedReadCapacityAutoScalingSettings": {
            "AutoScalingDisabled": boolean,
            "AutoScalingRoleArn": "string",
            "MaximumUnits": number,
            "MinimumUnits": number,
            "ScalingPolicies": [
               {
                  "PolicyName": "string",
                  "TargetTrackingScalingPolicyConfiguration": {
                     "DisableScaleIn": boolean,
                     "ScaleInCooldown": number,
                     "ScaleOutCooldown": number,
                     "TargetValue": number
                  }
               }
            ]
         },
         "ReplicaProvisionedReadCapacityUnits": number,
         "ReplicaProvisionedWriteCapacityAutoScalingSettings": {
            "AutoScalingDisabled": boolean,
            "AutoScalingRoleArn": "string",
            "MaximumUnits": number,
            "MinimumUnits": number,
            "ScalingPolicies": [
               {
                  "PolicyName": "string",
                  "TargetTrackingScalingPolicyConfiguration": {
                     "DisableScaleIn": boolean,
                     "ScaleInCooldown": number,
                     "ScaleOutCooldown": number,
                     "TargetValue": number
                  }
               }
            ]
         },
         "ReplicaProvisionedWriteCapacityUnits": number,
         "ReplicaStatus": "string",
         "ReplicaTableClassSummary": {
            "LastUpdateDateTime": number,
            "TableClass": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[GlobalTableName](#API_UpdateGlobalTableSettings_ResponseSyntax)**

The name of the global table.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

**[ReplicaSettings](#API_UpdateGlobalTableSettings_ResponseSyntax)**

The Region-specific settings for the global table.

Type: Array of [ReplicaSettingsDescription](api-replicasettingsdescription.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**GlobalTableNotFoundException**

The specified global table does not exist.

HTTP Status Code: 400

**IndexNotFoundException**

The operation tried to access a nonexistent index.

HTTP Status Code: 400

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**LimitExceededException**

There is no limit to the number of daily on-demand backups that can be taken.

For most purposes, up to 500 simultaneous table operations are allowed per account.
These operations include `CreateTable`, `UpdateTable`,
`DeleteTable`, `UpdateTimeToLive`,
`RestoreTableFromBackup`, and `RestoreTableToPointInTime`.

When you are creating a table with one or more secondary indexes, you can have up
to 250 such requests running at a time. However, if the table or index specifications
are complex, then DynamoDB might temporarily reduce the number of concurrent
operations.

When importing into DynamoDB, up to 50 simultaneous import table operations are
allowed per account.

There is a soft account quota of 2,500 tables.

GetRecords was called with a value of more than 1000 for the limit request
parameter.

More than 2 processes are reading from the same streams shard at the same time.
Exceeding this limit may result in request throttling.

**message**

Too many operations for a given subscriber.

HTTP Status Code: 400

**ReplicaNotFoundException**

The specified replica is no longer part of the global table.

HTTP Status Code: 400

**ResourceInUseException**

The operation conflicts with the resource's availability. For example:

- You attempted to recreate an existing table.

- You tried to delete a table currently in the `CREATING`
state.

- You tried to update a resource that was already being updated.

When appropriate, wait for the ongoing update to complete and attempt the request
again.

**message**

The resource which is being attempted to be changed is in use.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/updateglobaltablesettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/updateglobaltablesettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateGlobalTable

UpdateItem

All content copied from https://docs.aws.amazon.com/.
