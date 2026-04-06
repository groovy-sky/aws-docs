# Finding

Contains information about a finding.

## Contents

**analyzedAt**

The time at which the resource was analyzed.

Type: Timestamp

Required: Yes

**condition**

The condition in the analyzed policy statement that resulted in a finding.

Type: String to string map

Required: Yes

**createdAt**

The time at which the finding was generated.

Type: Timestamp

Required: Yes

**id**

The ID of the finding.

Type: String

Required: Yes

**resourceOwnerAccount**

The AWS account ID that owns the resource.

Type: String

Required: Yes

**resourceType**

The type of the resource identified in the finding.

Type: String

Valid Values: `AWS::S3::Bucket | AWS::IAM::Role | AWS::SQS::Queue | AWS::Lambda::Function | AWS::Lambda::LayerVersion | AWS::KMS::Key | AWS::SecretsManager::Secret | AWS::EFS::FileSystem | AWS::EC2::Snapshot | AWS::ECR::Repository | AWS::RDS::DBSnapshot | AWS::RDS::DBClusterSnapshot | AWS::SNS::Topic | AWS::S3Express::DirectoryBucket | AWS::DynamoDB::Table | AWS::DynamoDB::Stream | AWS::IAM::User`

Required: Yes

**status**

The current status of the finding.

Type: String

Valid Values: `ACTIVE | ARCHIVED | RESOLVED`

Required: Yes

**updatedAt**

The time at which the finding was updated.

Type: Timestamp

Required: Yes

**action**

The action in the analyzed policy statement that an external principal has permission to
use.

Type: Array of strings

Required: No

**error**

An error.

Type: String

Required: No

**isPublic**

Indicates whether the policy that generated the finding allows public access to the
resource.

Type: Boolean

Required: No

**principal**

The external principal that has access to a resource within the zone of trust.

Type: String to string map

Required: No

**resource**

The resource that an external principal has access to.

Type: String

Required: No

**resourceControlPolicyRestriction**

The type of restriction applied to the finding by the resource owner with an Organizations
resource control policy (RCP).

Type: String

Valid Values: `APPLICABLE | FAILED_TO_EVALUATE_RCP | NOT_APPLICABLE | APPLIED`

Required: No

**sources**

The sources of the finding. This indicates how the access that generated the finding is
granted. It is populated for Amazon S3 bucket findings.

Type: Array of [FindingSource](api-findingsource.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/finding.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/finding.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/finding.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExternalAccessFindingsStatistics

FindingAggregationAccountDetails
