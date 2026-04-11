# DescribeGlobalTableSettings

Describes Region-specific settings for a global table.

###### Important

This documentation is for version 2017.11.29 (Legacy) of global tables, which should be avoided for new global tables. Customers should use [Global Tables version 2019.11.21 (Current)](../../../../services/dynamodb/latest/developerguide/globaltables.md) when possible, because it provides greater flexibility, higher efficiency, and consumes less write capacity than 2017.11.29 (Legacy).

To determine which version you're using, see [Determining the global table version you are using](../../../../services/dynamodb/latest/developerguide/globaltables-determineversion.md). To update existing global tables from version 2017.11.29 (Legacy) to version 2019.11.21 (Current), see [Upgrading global tables](../../../../services/dynamodb/latest/developerguide/v2globaltables-upgrade.md).

## Request Syntax

```nohighlight

{
   "GlobalTableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[GlobalTableName](#API_DescribeGlobalTableSettings_RequestSyntax)**

The name of the global table to describe.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: Yes

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

**[GlobalTableName](#API_DescribeGlobalTableSettings_ResponseSyntax)**

The name of the global table.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

**[ReplicaSettings](#API_DescribeGlobalTableSettings_ResponseSyntax)**

The Region-specific settings for the global table.

Type: Array of [ReplicaSettingsDescription](api-replicasettingsdescription.md) objects

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/describeglobaltablesettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/describeglobaltablesettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeGlobalTable

DescribeImport

All content copied from https://docs.aws.amazon.com/.
