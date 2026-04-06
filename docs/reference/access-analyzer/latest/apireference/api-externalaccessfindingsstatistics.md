# ExternalAccessFindingsStatistics

Provides aggregate statistics about the findings for the specified external access
analyzer.

## Contents

**resourceTypeStatistics**

The total number of active cross-account and public findings for each resource type of
the specified external access analyzer.

Type: String to [ResourceTypeDetails](api-resourcetypedetails.md) object map

Valid Keys: `AWS::S3::Bucket | AWS::IAM::Role | AWS::SQS::Queue | AWS::Lambda::Function | AWS::Lambda::LayerVersion | AWS::KMS::Key | AWS::SecretsManager::Secret | AWS::EFS::FileSystem | AWS::EC2::Snapshot | AWS::ECR::Repository | AWS::RDS::DBSnapshot | AWS::RDS::DBClusterSnapshot | AWS::SNS::Topic | AWS::S3Express::DirectoryBucket | AWS::DynamoDB::Table | AWS::DynamoDB::Stream | AWS::IAM::User`

Required: No

**totalActiveFindings**

The number of active findings for the specified external access analyzer.

Type: Integer

Required: No

**totalArchivedFindings**

The number of archived findings for the specified external access analyzer.

Type: Integer

Required: No

**totalResolvedFindings**

The number of resolved findings for the specified external access analyzer.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/ExternalAccessFindingsStatistics)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/ExternalAccessFindingsStatistics)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/ExternalAccessFindingsStatistics)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExternalAccessDetails

Finding
