# AnalyzedResource

Contains details about the analyzed resource.

## Contents

**analyzedAt**

The time at which the resource was analyzed.

Type: Timestamp

Required: Yes

**createdAt**

The time at which the finding was created.

Type: Timestamp

Required: Yes

**isPublic**

Indicates whether the policy that generated the finding grants public access to the
resource.

Type: Boolean

Required: Yes

**resourceArn**

The ARN of the resource that was analyzed.

Type: String

Pattern: `arn:[^:]*:[^:]*:[^:]*:[^:]*:.*`

Required: Yes

**resourceOwnerAccount**

The AWS account ID that owns the resource.

Type: String

Required: Yes

**resourceType**

The type of the resource that was analyzed.

Type: String

Valid Values: `AWS::S3::Bucket | AWS::IAM::Role | AWS::SQS::Queue | AWS::Lambda::Function | AWS::Lambda::LayerVersion | AWS::KMS::Key | AWS::SecretsManager::Secret | AWS::EFS::FileSystem | AWS::EC2::Snapshot | AWS::ECR::Repository | AWS::RDS::DBSnapshot | AWS::RDS::DBClusterSnapshot | AWS::SNS::Topic | AWS::S3Express::DirectoryBucket | AWS::DynamoDB::Table | AWS::DynamoDB::Stream | AWS::IAM::User`

Required: Yes

**updatedAt**

The time at which the finding was updated.

Type: Timestamp

Required: Yes

**actions**

The actions that an external principal is granted permission to use by the policy that
generated the finding.

Type: Array of strings

Required: No

**error**

An error message.

Type: String

Required: No

**sharedVia**

Indicates how the access that generated the finding is granted. This is populated for
Amazon S3 bucket findings.

Type: Array of strings

Required: No

**status**

The current status of the finding generated from the analyzed resource.

Type: String

Valid Values: `ACTIVE | ARCHIVED | RESOLVED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/AnalyzedResource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/AnalyzedResource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/AnalyzedResource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalysisRuleCriteria

AnalyzedResourceSummary
