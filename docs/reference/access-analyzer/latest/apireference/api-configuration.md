# Configuration

Access control configuration structures for your resource. You specify the configuration
as a type-value pair. You can specify only one type of access control configuration.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**dynamodbStream**

The access control configuration is for a DynamoDB stream.

Type: [DynamodbStreamConfiguration](api-dynamodbstreamconfiguration.md) object

Required: No

**dynamodbTable**

The access control configuration is for a DynamoDB table or index.

Type: [DynamodbTableConfiguration](api-dynamodbtableconfiguration.md) object

Required: No

**ebsSnapshot**

The access control configuration is for an Amazon EBS volume snapshot.

Type: [EbsSnapshotConfiguration](api-ebssnapshotconfiguration.md) object

Required: No

**ecrRepository**

The access control configuration is for an Amazon ECR repository.

Type: [EcrRepositoryConfiguration](api-ecrrepositoryconfiguration.md) object

Required: No

**efsFileSystem**

The access control configuration is for an Amazon EFS file system.

Type: [EfsFileSystemConfiguration](api-efsfilesystemconfiguration.md) object

Required: No

**iamRole**

The access control configuration is for an IAM role.

Type: [IamRoleConfiguration](api-iamroleconfiguration.md) object

Required: No

**kmsKey**

The access control configuration is for a KMS key.

Type: [KmsKeyConfiguration](api-kmskeyconfiguration.md) object

Required: No

**rdsDbClusterSnapshot**

The access control configuration is for an Amazon RDS DB cluster snapshot.

Type: [RdsDbClusterSnapshotConfiguration](api-rdsdbclustersnapshotconfiguration.md) object

Required: No

**rdsDbSnapshot**

The access control configuration is for an Amazon RDS DB snapshot.

Type: [RdsDbSnapshotConfiguration](api-rdsdbsnapshotconfiguration.md) object

Required: No

**s3Bucket**

The access control configuration is for an Amazon S3 bucket.

Type: [S3BucketConfiguration](api-s3bucketconfiguration.md) object

Required: No

**s3ExpressDirectoryBucket**

The access control configuration is for an Amazon S3 directory bucket.

Type: [S3ExpressDirectoryBucketConfiguration](api-s3expressdirectorybucketconfiguration.md) object

Required: No

**secretsManagerSecret**

The access control configuration is for a Secrets Manager secret.

Type: [SecretsManagerSecretConfiguration](api-secretsmanagersecretconfiguration.md) object

Required: No

**snsTopic**

The access control configuration is for an Amazon SNS topic

Type: [SnsTopicConfiguration](api-snstopicconfiguration.md) object

Required: No

**sqsQueue**

The access control configuration is for an Amazon SQS queue.

Type: [SqsQueueConfiguration](api-sqsqueueconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/Configuration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/Configuration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/Configuration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudTrailProperties

Criterion
