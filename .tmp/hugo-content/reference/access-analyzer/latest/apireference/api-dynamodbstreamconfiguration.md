# DynamodbStreamConfiguration

The proposed access control configuration for a DynamoDB stream. You can propose a
configuration for a new DynamoDB stream or an existing DynamoDB stream that you own by specifying
the policy for the DynamoDB stream. For more information, see [PutResourcePolicy](../../../amazondynamodb/latest/apireference/api-putresourcepolicy.md).

- If the configuration is for an existing DynamoDB stream and you do not specify the
DynamoDB policy, then the access preview uses the existing DynamoDB policy for the
stream.

- If the access preview is for a new resource and you do not specify the policy,
then the access preview assumes a DynamoDB stream without a policy.

- To propose deletion of an existing DynamoDB stream policy, you can specify an empty
string for the DynamoDB policy.

## Contents

**streamPolicy**

The proposed resource policy defining who can access or manage the DynamoDB stream.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/dynamodbstreamconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/dynamodbstreamconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/dynamodbstreamconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Criterion

DynamodbTableConfiguration
