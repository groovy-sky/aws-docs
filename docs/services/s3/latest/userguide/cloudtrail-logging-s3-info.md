# Amazon S3 CloudTrail events

###### Important

Amazon S3 now applies server-side encryption with Amazon S3 managed keys (SSE-S3) as the base level of encryption for every bucket in Amazon S3. Starting January 5, 2023, all new object uploads to Amazon S3 are automatically encrypted at no additional cost and with no impact on performance. The automatic encryption status for S3 bucket default encryption configuration and for new object uploads is available in CloudTrail logs, S3 Inventory, S3 Storage Lens, the Amazon S3 console, and as an additional Amazon S3 API response header in the AWS CLI and AWS SDKs. For more information, see [Default encryption FAQ](https://docs.aws.amazon.com/AmazonS3/latest/userguide/default-encryption-faq.html).

This section provides information about the events that S3 logs to CloudTrail.

## Amazon S3 data events in CloudTrail

[Data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#logging-data-events) provide information about the resource operations
performed on or in a resource (for example, reading or writing to an Amazon S3
object). These are also
known as data plane operations. Data events are often high-volume activities. By
default, CloudTrail doesn’t log data events. The CloudTrail **Event**
**history** doesn't record data events.

Additional charges apply for data events. For more information about CloudTrail
pricing, see [AWS CloudTrail\
Pricing](https://aws.amazon.com/cloudtrail/pricing).

You can log data events for the Amazon S3 resource types by using the CloudTrail
console, AWS CLI, or CloudTrail API operations. For more information about how to log
data events, see [Logging data events with the AWS Management Console](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#logging-data-events-console) and [Logging data events with the AWS Command Line Interface](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-with-the-AWS-CLI) in the
_AWS CloudTrail User Guide_.

The following table lists the Amazon S3 resource types for which you can
log data events. The **Data event type (console)**
column shows the value to choose from the **Data event type**
list on the CloudTrail console. The **resources.type**
**value** column shows the `resources.type` value, which
you would specify when configuring advanced event selectors using the AWS CLI or
CloudTrail APIs. The **Data APIs logged to CloudTrail** column
shows the API calls logged to CloudTrail for the resource type.

Data event type (console)resources.type valueData APIs logged to CloudTrail**S3**`AWS::S3::Object`

- [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)

- [CompleteMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)

- [CopyObject](../api/api-copyobject.md)

- [CreateMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)

- [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)

- [DeleteObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html)

- [DeleteObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjects.html)

- [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)

- [GetObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)

- [GetObjectAttributes](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html)

- [GetObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html)

- [GetObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html)

- [GetObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html)

- [GetObjectTorrent](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTorrent.html)

- [HeadObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html)

- [HeadBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html)

- [ListObjectVersions](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectVersions.html)

- [ListObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)

- [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)

- [PutObject](../api/api-putobject.md)

- [PutObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectAcl.html)

- [PutObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLegalHold.html)

- [PutObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectRetention.html)

- [PutObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html)

- [RestoreObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)

- [SelectObjectContent](https://docs.aws.amazon.com/AmazonS3/latest/API/API_SelectObjectContent.html)

- [UpdateObjectEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateObjectEncryption.html)

- [UploadPart](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)

- [UploadPartCopy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html)

**S3 Express One Zone**`AWS::S3Express::Object`

- [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)

- [CompleteMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)

- [CreateSession](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html)

- [CopyObject](../api/api-copyobject.md)

- [CreateMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)

- [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)

- [DeleteObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjects.html)

- [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)

- [GetObjectAttributes](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html)

- [HeadBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html)

- [HeadObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html)

- [ListObjectsV2](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)

- [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)

- [PutObject](../api/api-putobject.md)

- [UploadPart](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)

- [UploadPartCopy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html)

**S3 Access Point**`AWS::S3::Access Point`

- [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)

- [CompleteMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)

- [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)

- [DeleteObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html)

- [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)

- [GetObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)

- [GetObjectAttributes](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html)

- [GetObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html)

- [GetObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html)

- [GetObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html)

- [HeadBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html)

- [HeadObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html)

- [ListObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)

- [ListObjectsV2](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)

- [ListObjectVersions](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectVersions.html)

- [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)

- [Presign](https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)

- [PutObject](../api/api-putobject.md)

- [PutObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLegalHold.html)

- [PutObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectRetention.html)

- [PutObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectAcl.html)

- [PutObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html)

- [RestoreObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)

- [UploadPart](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)

- [UploadPartCopy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html)
(same-region copies only)

**S3 Object Lambda**`AWS::S3ObjectLambda::AccessPoint`

- [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)

- [CompleteMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)

- [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)

- [DeleteObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html)

- [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)

- [GetObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)

- [GetObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html)

- [GetObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html)

- [GetObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html)

- [HeadObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html)

- [ListObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)

- [ListObjectVersions](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectVersions.html)

- [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)

- [PutObject](../api/api-putobject.md)

- [PutObjectLegalHold](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLegalHold.html)

- [PutObjectRetention](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectRetention.html)

- [PutObjectAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectAcl.html)

- [PutObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html)

- [RestoreObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html)

- [UploadPart](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)

- [WriteGetObjectResponse](https://docs.aws.amazon.com/AmazonS3/latest/API/API_WriteGetObjectResponse.html)

**S3 Outposts**`AWS::S3Outposts::Object`

- [AbortMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)

- [CompleteMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)

- [DeleteObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)

- [DeleteObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html)

- [GetObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)

- [GetObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html)

- [HeadObject](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html)

- [ListObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)

- [ListObjectsV2](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)

- [ListParts](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)

- [PutObject](../api/api-putobject.md)

- [PutObjectTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html)

- [UploadPart](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)

- [UploadPartCopy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html)

You can configure advanced event selectors to filter on the
`eventName`, `readOnly`, and
`resources.ARN` fields to log only those events that are
important to you. For more information about these fields, see [AdvancedFieldSelector](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html) in the
_AWS CloudTrail API Reference_.

## Amazon S3 management events in CloudTrail

Amazon S3 logs all control plane operations as management events. For more
information about S3 API operations, see the [Amazon S3 API\
Reference](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Operations.html).

## How CloudTrail captures requests made to Amazon S3

By default, CloudTrail logs S3 bucket-level API calls that were made in the last 90
days, but not log requests made to objects. Bucket-level calls include events
such as `CreateBucket`, `DeleteBucket`,
`PutBucketLifecycle`, `PutBucketPolicy`, and so on.
You can see bucket-level events on the CloudTrail console. However, you can't view
data events (Amazon S3 object-level calls) there—you must parse or query CloudTrail
logs for them.

If you are logging data activity with AWS CloudTrail, the event record for an Amazon S3 `DeleteObjects` data event includes both the `DeleteObjects` event and a `DeleteObject` event for each object deleted as part of that operation. You can exclude the additional visibility about deleted objects from the event record. For more information, see [AWS CLI examples for filtering data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/filtering-data-events.html#filtering-data-events-deleteobjects) in the _AWS CloudTrail User Guide._

## Amazon S3 account-level actions tracked by CloudTrail logging

CloudTrail logs account-level actions. Amazon S3 records are written together with other
AWS service records in a log file. CloudTrail determines when to create and write to
a new file based on a time period and file size.

The tables in this section list the Amazon S3 account-level actions that are
supported for logging by CloudTrail.

Amazon S3 account-level API actions tracked by CloudTrail logging appear as the following
event names. The CloudTrail event names differ from the API action name. For example,
DeletePublicAccessBlock is DeleteAccountPublicAccessBlock.

- [DeleteAccountPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html)

- [GetAccountPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html)

- [PutAccountPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html)

## Amazon S3 bucket-level actions that are tracked by CloudTrail logging

By default, CloudTrail logs bucket-level actions for general purpose buckets. Amazon S3
records are written together with other AWS service records in a log file.
CloudTrail determines when to create and write to a new file based on a time period
and file size.

This section lists the Amazon S3 bucket-level actions that are supported for
logging by CloudTrail.

Amazon S3 bucket-level API actions tracked by CloudTrail logging appear as the following
event names. In some cases, the CloudTrail event name differs from the API action
name. For example, `PutBucketLifecycleConfiguration` is
`PutBucketLifecycle`.

- [CreateBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)

- [CreateBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html) (V2 API operation)

- [CreateBucketMetadataTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataTableConfiguration.html) (V1 API operation)

- [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)

- [DeleteBucketAnalyticsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html)

- [DeleteBucketCors](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html)

- [DeleteBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)

- [DeleteBucketIntelligentTieringConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html)

- [DeleteBucketInventoryConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html)

- [DeleteBucketLifecycle](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)

- [DeleteBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html) (V2 API operation)

- [DeleteBucketMetadataTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html) (V1 API operation)

- [DeleteBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html)

- [DeleteBucketOwnershipControls](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketOwnershipControls.html)

- [DeleteBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html)

- [DeleteBucketPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html)

- [DeleteBucketReplication](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html)

- [DeleteBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)

- [GetAccelerateConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAccelerateConfiguration.html)

- [GetBucketAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAcl.html)

- [GetBucketAnalyticsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html)

- [GetBucketCors](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketCors.html)

- [GetBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)

- [GetBucketIntelligentTieringConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html)

- [GetBucketInventoryConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html)

- [GetBucketLifecycle](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html)

- [GetBucketLocation](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLocation.html)

- [GetBucketLogging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html)

- [GetBucketMetadataConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html) (V2 API operation)

- [GetBucketMetadataTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html) (V1 API operation)

- [GetBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html)

- [GetBucketNotification](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketNotification.html)

- [GetBucketObjectLockConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html)

- [GetBucketOwnershipControls](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketOwnershipControls.html)

- [GetBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html)

- [GetBucketPolicyStatus](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html)

- [GetBucketPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html)

- [GetBucketReplication](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html)

- [GetBucketRequestPayment](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketRequestPayment.html)

- [GetBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html)

- [GetBucketVersioning](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html)

- [GetBucketWebsite](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketWebsite.html)

- [HeadBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html)

- [ListBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html)

- [PutAccelerateConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAccelerateConfiguration.html)

- [PutBucketAcl](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAcl.html)

- [PutBucketAnalyticsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html)

- [PutBucketCors](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html)

- [PutBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html)

- [PutBucketIntelligentTieringConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html)

- [PutBucketInventoryConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html)

- [PutBucketLifecycle](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html)

- [PutBucketLogging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLogging.html)

- [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html)

- [PutBucketNotification](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketNotification.html)

- [PutBucketObjectLockConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectLockConfiguration.html)

- [PutBucketOwnershipControls](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketOwnershipControls.html)

- [PutBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html)

- [PutBucketPublicAccessBlock](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html)

- [PutBucketReplication](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html)

- [PutBucketRequestPayment](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketRequestPayment.html)

- [PutBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html)

- [PutBucketVersioning](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html)

- [PutBucketWebsite](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html)

- [UpdateBucketMetadataJournalTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguratione.html)

- [UpdateBucketMetadataInventoryTableConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html)

In addition to these API operations, you can also use the [OPTIONS\
object](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html) object-level action. This action is treated like a
bucket-level action in CloudTrail logging because the action checks the CORS
configuration of a bucket.

###### Note

The HeadBucket API is supported as an Amazon S3 data event in CloudTrail.

## Amazon S3 Express One Zone bucket-level (Regional API endpoint) actions tracked by CloudTrail logging

By default, CloudTrail logs bucket-level actions for directory buckets as management
events. The `eventsource` for CloudTrail management events for S3 Express One Zone
is `s3express.amazonaws.com`.

These following Regional endpoint API operations are logged to CloudTrail.

- [CreateBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)

- [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)

- [DeleteBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html)

- [GetBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html)

- [PutBucketPolicy](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html)

- [ListDirectoryBuckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListDirectoryBuckets.html)

- [PutBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html)

- [GetBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)

- [DeleteBucketEncryption](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)

For more information, see [Logging with AWS CloudTrail for S3 Express One Zone](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-one-zone-logging.html)

## Amazon S3 object-level actions in cross-account scenarios

The following are special use cases involving the object-level API calls in
cross-account scenarios and how CloudTrail logs are reported. CloudTrail delivers logs to
the requester (the account that made the API call), except in some access denied
cases where log entries are redacted or omitted. When setting up cross-account
access, consider the examples in this section.

###### Note

The examples assume that CloudTrail logs are appropriately configured.

### Example 1: CloudTrail delivers logs to the bucket owner

CloudTrail delivers logs to the bucket owner even if the bucket owner does not
have permissions for the same object API operation. Consider the following
cross-account scenario:

- Account A owns the bucket.

- Account B (the requester) tries to access an object in that
bucket.

- Account C owns the object. Account C might or might not be the
same account as Account A.

###### Note

CloudTrail always delivers object-level API logs to the requester (Account
B). In addition, CloudTrail also delivers the same logs to the bucket owner
(Account A) even when the bucket owner does not own the object (Account
C) or have permissions for those same API operations on that
object.

### Example 2: CloudTrail does not proliferate email addresses that are used in setting object ACLs

Consider the following cross-account scenario:

- Account A owns the bucket.

- Account B (the requester) sends a request to set an object ACL
grant by using an email address. For more information about ACLs,
see [Access control list (ACL) overview](acl-overview.md).

The requester gets the logs along with the email information. However, the
bucket owner—if they are eligible to receive logs, as in example
1—gets the CloudTrail log reporting the event. However, the bucket owner
doesn't get the ACL configuration information, specifically the grantee
email address and the grant. The only information that the log tells the
bucket owner is that an ACL API call was made by Account B.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logging with CloudTrail

Example log files
