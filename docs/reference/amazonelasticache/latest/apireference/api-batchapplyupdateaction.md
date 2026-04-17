---
title: "BatchApplyUpdateAction"
---

# BatchApplyUpdateAction

Apply the service update. For more information on service updates and applying them,
see [Applying Service\
Updates](../../../../services/amazonelasticache/latest/dg/applying-updates.md).

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ServiceUpdateName**

The unique ID of the service update

Type: String

Required: Yes

**CacheClusterIds.member.N**

The cache cluster IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

**ReplicationGroupIds.member.N**

The replication group IDs

Type: Array of strings

Array Members: Maximum number of 20 items.

Required: No

## Response Elements

The following elements are returned by the service.

**ProcessedUpdateActions.ProcessedUpdateAction.N**

Update actions that have been processed successfully

Type: Array of [ProcessedUpdateAction](api-processedupdateaction.md) objects

**UnprocessedUpdateActions.UnprocessedUpdateAction.N**

Update actions that haven't been processed successfully

Type: Array of [UnprocessedUpdateAction](api-unprocessedupdateaction.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value for a parameter is invalid.

**message**

A parameter value is invalid.

HTTP Status Code: 400

**ServiceUpdateNotFoundFault**

The service update doesn't exist

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/elasticache-2015-02-02/BatchApplyUpdateAction)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/BatchApplyUpdateAction)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizeCacheSecurityGroupIngress

BatchStopUpdateAction

All content copied from https://docs.aws.amazon.com/.
