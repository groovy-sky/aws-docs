---
title: "DescribeTableReplicaAutoScaling"
---

# DescribeTableReplicaAutoScaling

Describes auto scaling settings across replicas of the global table at once.

## Request Syntax

```nohighlight

{
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[TableName](#API_DescribeTableReplicaAutoScaling_RequestSyntax)**

The name of the table. You can also provide the Amazon Resource Name (ARN) of the table in this
parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Response Syntax

```nohighlight

{
   "TableAutoScalingDescription": {
      "Replicas": [
         {
            "GlobalSecondaryIndexes": [
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
                  }
               }
            ],
            "RegionName": "string",
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
            "ReplicaStatus": "string"
         }
      ],
      "TableName": "string",
      "TableStatus": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[TableAutoScalingDescription](#API_DescribeTableReplicaAutoScaling_ResponseSyntax)**

Represents the auto scaling properties of the table.

Type: [TableAutoScalingDescription](api-tableautoscalingdescription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/DescribeTableReplicaAutoScaling)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeTable

DescribeTimeToLive

All content copied from https://docs.aws.amazon.com/.
