# DynamodbTableConfiguration

The proposed access control configuration for a DynamoDB table or index. You can propose a
configuration for a new DynamoDB table or index or an existing DynamoDB table or index that you
own by specifying the policy for the DynamoDB table or index. For more information, see [PutResourcePolicy](../../../amazondynamodb/latest/apireference/api-putresourcepolicy.md).

- If the configuration is for an existing DynamoDB table or index and you do not
specify the DynamoDB policy, then the access preview uses the existing DynamoDB policy for
the table or index.

- If the access preview is for a new resource and you do not specify the policy,
then the access preview assumes a DynamoDB table without a policy.

- To propose deletion of an existing DynamoDB table or index policy, you can specify an
empty string for the DynamoDB policy.

## Contents

**tablePolicy**

The proposed resource policy defining who can access or manage the DynamoDB table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/dynamodbtableconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/dynamodbtableconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/dynamodbtableconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DynamodbStreamConfiguration

EbsSnapshotConfiguration
