# InternalAccessAnalysisRuleCriteria

The criteria for an analysis rule for an internal access analyzer.

## Contents

**accountIds**

A list of AWS account IDs to apply to the internal access analysis rule criteria.
Account IDs can only be applied to the analysis rule criteria for organization-level
analyzers.

Type: Array of strings

Required: No

**resourceArns**

A list of resource ARNs to apply to the internal access analysis rule criteria. The
analyzer will only generate findings for resources that match these ARNs.

Type: Array of strings

Required: No

**resourceTypes**

A list of resource types to apply to the internal access analysis rule criteria. The
analyzer will only generate findings for resources of these types. These resource types are
currently supported for internal access analyzers:

- `AWS::S3::Bucket`

- `AWS::RDS::DBSnapshot`

- `AWS::RDS::DBClusterSnapshot`

- `AWS::S3Express::DirectoryBucket`

- `AWS::DynamoDB::Table`

- `AWS::DynamoDB::Stream`

Type: Array of strings

Valid Values: `AWS::S3::Bucket | AWS::IAM::Role | AWS::SQS::Queue | AWS::Lambda::Function | AWS::Lambda::LayerVersion | AWS::KMS::Key | AWS::SecretsManager::Secret | AWS::EFS::FileSystem | AWS::EC2::Snapshot | AWS::ECR::Repository | AWS::RDS::DBSnapshot | AWS::RDS::DBClusterSnapshot | AWS::SNS::Topic | AWS::S3Express::DirectoryBucket | AWS::DynamoDB::Table | AWS::DynamoDB::Stream | AWS::IAM::User`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/internalaccessanalysisrulecriteria.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/internalaccessanalysisrulecriteria.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/internalaccessanalysisrulecriteria.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InternalAccessAnalysisRule

InternalAccessConfiguration
