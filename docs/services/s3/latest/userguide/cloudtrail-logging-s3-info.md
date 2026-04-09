# Amazon S3 CloudTrail events

###### Important

Amazon S3 now applies server-side encryption with Amazon S3 managed keys (SSE-S3) as the base level of encryption for every bucket in Amazon S3. Starting January 5, 2023, all new object uploads to Amazon S3 are automatically encrypted at no additional cost and with no impact on performance. The automatic encryption status for S3 bucket default encryption configuration and for new object uploads is available in CloudTrail logs, S3 Inventory, S3 Storage Lens, the Amazon S3 console, and as an additional Amazon S3 API response header in the AWS CLI and AWS SDKs. For more information, see [Default encryption FAQ](default-encryption-faq.md).

This section provides information about the events that S3 logs to CloudTrail.

## Amazon S3 data events in CloudTrail

[Data events](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events) provide information about the resource operations
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
data events, see [Logging data events with the AWS Management Console](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#logging-data-events-console) and [Logging data events with the AWS Command Line Interface](../../../awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.md#creating-data-event-selectors-with-the-AWS-CLI) in the
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

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CopyObject](../api/api-copyobject.md)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjectTagging](../api/api-deleteobjecttagging.md)

- [DeleteObjects](../api/api-deleteobjects.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectAcl](../api/api-getobjectacl.md)

- [GetObjectAttributes](../api/api-getobjectattributes.md)

- [GetObjectLegalHold](../api/api-getobjectlegalhold.md)

- [GetObjectRetention](../api/api-getobjectretention.md)

- [GetObjectTagging](../api/api-getobjecttagging.md)

- [GetObjectTorrent](../api/api-getobjecttorrent.md)

- [HeadObject](../api/api-headobject.md)

- [HeadBucket](../api/api-headbucket.md)

- [ListObjectVersions](../api/api-listobjectversions.md)

- [ListObjects](../api/api-listobjects.md)

- [ListParts](../api/api-listparts.md)

- [PutObject](../api/api-putobject.md)

- [PutObjectAcl](../api/api-putobjectacl.md)

- [PutObjectLegalHold](../api/api-putobjectlegalhold.md)

- [PutObjectRetention](../api/api-putobjectretention.md)

- [PutObjectTagging](../api/api-putobjecttagging.md)

- [RestoreObject](../api/api-restoreobject.md)

- [SelectObjectContent](../api/api-selectobjectcontent.md)

- [UpdateObjectEncryption](../api/api-updateobjectencryption.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)

**S3 Express One Zone**`AWS::S3Express::Object`

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CreateSession](../api/api-createsession.md)

- [CopyObject](../api/api-copyobject.md)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjects](../api/api-deleteobjects.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectAttributes](../api/api-getobjectattributes.md)

- [HeadBucket](../api/api-headbucket.md)

- [HeadObject](../api/api-headobject.md)

- [ListObjectsV2](../api/api-listobjectsv2.md)

- [ListParts](../api/api-listparts.md)

- [PutObject](../api/api-putobject.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)

**S3 Access Point**`AWS::S3::Access Point`

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjectTagging](../api/api-deleteobjecttagging.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectAcl](../api/api-getobjectacl.md)

- [GetObjectAttributes](../api/api-getobjectattributes.md)

- [GetObjectLegalHold](../api/api-getobjectlegalhold.md)

- [GetObjectRetention](../api/api-getobjectretention.md)

- [GetObjectTagging](../api/api-getobjecttagging.md)

- [HeadBucket](../api/api-headbucket.md)

- [HeadObject](../api/api-headobject.md)

- [ListObjects](../api/api-listobjects.md)

- [ListObjectsV2](../api/api-listobjectsv2.md)

- [ListObjectVersions](../api/api-listobjectversions.md)

- [ListParts](../api/api-listparts.md)

- [Presign](../api/sigv4-query-string-auth.md)

- [PutObject](../api/api-putobject.md)

- [PutObjectLegalHold](../api/api-putobjectlegalhold.md)

- [PutObjectRetention](../api/api-putobjectretention.md)

- [PutObjectAcl](../api/api-putobjectacl.md)

- [PutObjectTagging](../api/api-putobjecttagging.md)

- [RestoreObject](../api/api-restoreobject.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)
(same-region copies only)

**S3 Object Lambda**`AWS::S3ObjectLambda::AccessPoint`

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjectTagging](../api/api-deleteobjecttagging.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectAcl](../api/api-getobjectacl.md)

- [GetObjectLegalHold](../api/api-getobjectlegalhold.md)

- [GetObjectRetention](../api/api-getobjectretention.md)

- [GetObjectTagging](../api/api-getobjecttagging.md)

- [HeadObject](../api/api-headobject.md)

- [ListObjects](../api/api-listobjects.md)

- [ListObjectVersions](../api/api-listobjectversions.md)

- [ListParts](../api/api-listparts.md)

- [PutObject](../api/api-putobject.md)

- [PutObjectLegalHold](../api/api-putobjectlegalhold.md)

- [PutObjectRetention](../api/api-putobjectretention.md)

- [PutObjectAcl](../api/api-putobjectacl.md)

- [PutObjectTagging](../api/api-putobjecttagging.md)

- [RestoreObject](../api/api-restoreobject.md)

- [UploadPart](../api/api-uploadpart.md)

- [WriteGetObjectResponse](../api/api-writegetobjectresponse.md)

**S3 Outposts**`AWS::S3Outposts::Object`

- [AbortMultipartUpload](../api/api-abortmultipartupload.md)

- [CompleteMultipartUpload](../api/api-completemultipartupload.md)

- [CopyObject](../api/api-copyobject.md) (same-region
copies only)

- [CreateMultipartUpload](../api/api-createmultipartupload.md)

- [DeleteObject](../api/api-deleteobject.md)

- [DeleteObjectTagging](../api/api-deleteobjecttagging.md)

- [GetObject](../api/api-getobject.md)

- [GetObjectTagging](../api/api-getobjecttagging.md)

- [HeadObject](../api/api-headobject.md)

- [ListObjects](../api/api-listobjects.md)

- [ListObjectsV2](../api/api-listobjectsv2.md)

- [ListParts](../api/api-listparts.md)

- [PutObject](../api/api-putobject.md)

- [PutObjectTagging](../api/api-putobjecttagging.md)

- [UploadPart](../api/api-uploadpart.md)

- [UploadPartCopy](../api/api-uploadpartcopy.md)

You can configure advanced event selectors to filter on the
`eventName`, `readOnly`, and
`resources.ARN` fields to log only those events that are
important to you. For more information about these fields, see [AdvancedFieldSelector](../../../../reference/awscloudtrail/latest/apireference/api-advancedfieldselector.md) in the
_AWS CloudTrail API Reference_.

## Amazon S3 management events in CloudTrail

Amazon S3 logs all control plane operations as management events. For more
information about S3 API operations, see the [Amazon S3 API\
Reference](../api/api-operations.md).

## How CloudTrail captures requests made to Amazon S3

By default, CloudTrail logs S3 bucket-level API calls that were made in the last 90
days, but not log requests made to objects. Bucket-level calls include events
such as `CreateBucket`, `DeleteBucket`,
`PutBucketLifecycle`, `PutBucketPolicy`, and so on.
You can see bucket-level events on the CloudTrail console. However, you can't view
data events (Amazon S3 object-level calls) there—you must parse or query CloudTrail
logs for them.

If you are logging data activity with AWS CloudTrail, the event record for an Amazon S3 `DeleteObjects` data event includes both the `DeleteObjects` event and a `DeleteObject` event for each object deleted as part of that operation. You can exclude the additional visibility about deleted objects from the event record. For more information, see [AWS CLI examples for filtering data events](../../../awscloudtrail/latest/userguide/filtering-data-events.md#filtering-data-events-deleteobjects) in the _AWS CloudTrail User Guide._

## Amazon S3 account-level actions tracked by CloudTrail logging

CloudTrail logs account-level actions. Amazon S3 records are written together with other
AWS service records in a log file. CloudTrail determines when to create and write to
a new file based on a time period and file size.

The tables in this section list the Amazon S3 account-level actions that are
supported for logging by CloudTrail.

Amazon S3 account-level API actions tracked by CloudTrail logging appear as the following
event names. The CloudTrail event names differ from the API action name. For example,
DeletePublicAccessBlock is DeleteAccountPublicAccessBlock.

- [DeleteAccountPublicAccessBlock](../api/api-control-deletepublicaccessblock.md)

- [GetAccountPublicAccessBlock](../api/api-control-getpublicaccessblock.md)

- [PutAccountPublicAccessBlock](../api/api-control-putpublicaccessblock.md)

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

- [CreateBucket](../api/api-createbucket.md)

- [CreateBucketMetadataConfiguration](../api/api-createbucketmetadataconfiguration.md) (V2 API operation)

- [CreateBucketMetadataTableConfiguration](../api/api-createbucketmetadatatableconfiguration.md) (V1 API operation)

- [DeleteBucket](../api/api-deletebucket.md)

- [DeleteBucketAnalyticsConfiguration](../api/api-deletebucketanalyticsconfiguration.md)

- [DeleteBucketCors](../api/api-deletebucketcors.md)

- [DeleteBucketEncryption](../api/api-deletebucketencryption.md)

- [DeleteBucketIntelligentTieringConfiguration](../api/api-deletebucketintelligenttieringconfiguration.md)

- [DeleteBucketInventoryConfiguration](../api/api-deletebucketinventoryconfiguration.md)

- [DeleteBucketLifecycle](../api/api-deletebucketlifecycle.md)

- [DeleteBucketMetadataConfiguration](../api/api-deletebucketmetadataconfiguration.md) (V2 API operation)

- [DeleteBucketMetadataTableConfiguration](../api/api-deletebucketmetadatatableconfiguration.md) (V1 API operation)

- [DeleteBucketMetricsConfiguration](../api/api-deletebucketmetricsconfiguration.md)

- [DeleteBucketOwnershipControls](../api/api-deletebucketownershipcontrols.md)

- [DeleteBucketPolicy](../api/api-deletebucketpolicy.md)

- [DeleteBucketPublicAccessBlock](../api/api-deletepublicaccessblock.md)

- [DeleteBucketReplication](../api/api-deletebucketreplication.md)

- [DeleteBucketTagging](../api/api-deletebuckettagging.md)

- [GetAccelerateConfiguration](../api/api-getbucketaccelerateconfiguration.md)

- [GetBucketAcl](../api/api-getbucketacl.md)

- [GetBucketAnalyticsConfiguration](../api/api-getbucketanalyticsconfiguration.md)

- [GetBucketCors](../api/api-getbucketcors.md)

- [GetBucketEncryption](../api/api-getbucketencryption.md)

- [GetBucketIntelligentTieringConfiguration](../api/api-getbucketintelligenttieringconfiguration.md)

- [GetBucketInventoryConfiguration](../api/api-getbucketinventoryconfiguration.md)

- [GetBucketLifecycle](../api/api-getbucketlifecycle.md)

- [GetBucketLocation](../api/api-getbucketlocation.md)

- [GetBucketLogging](../api/api-getbucketlogging.md)

- [GetBucketMetadataConfiguration](../api/api-getbucketmetadataconfiguration.md) (V2 API operation)

- [GetBucketMetadataTableConfiguration](../api/api-getbucketmetadatatableconfiguration.md) (V1 API operation)

- [GetBucketMetricsConfiguration](../api/api-getbucketmetricsconfiguration.md)

- [GetBucketNotification](../api/api-getbucketnotification.md)

- [GetBucketObjectLockConfiguration](../api/api-getobjectlockconfiguration.md)

- [GetBucketOwnershipControls](../api/api-getbucketownershipcontrols.md)

- [GetBucketPolicy](../api/api-getbucketpolicy.md)

- [GetBucketPolicyStatus](../api/api-getbucketpolicystatus.md)

- [GetBucketPublicAccessBlock](../api/api-getpublicaccessblock.md)

- [GetBucketReplication](../api/api-getbucketreplication.md)

- [GetBucketRequestPayment](../api/api-getbucketrequestpayment.md)

- [GetBucketTagging](../api/api-getbuckettagging.md)

- [GetBucketVersioning](../api/api-getbucketversioning.md)

- [GetBucketWebsite](../api/api-getbucketwebsite.md)

- [HeadBucket](../api/api-headbucket.md)

- [ListBuckets](../api/api-listbuckets.md)

- [PutAccelerateConfiguration](../api/api-putbucketaccelerateconfiguration.md)

- [PutBucketAcl](../api/api-putbucketacl.md)

- [PutBucketAnalyticsConfiguration](../api/api-putbucketanalyticsconfiguration.md)

- [PutBucketCors](../api/api-putbucketcors.md)

- [PutBucketEncryption](../api/api-putbucketencryption.md)

- [PutBucketIntelligentTieringConfiguration](../api/api-putbucketintelligenttieringconfiguration.md)

- [PutBucketInventoryConfiguration](../api/api-putbucketinventoryconfiguration.md)

- [PutBucketLifecycle](../api/api-putbucketlifecycle.md)

- [PutBucketLogging](../api/api-putbucketlogging.md)

- [PutBucketMetricsConfiguration](../api/api-putbucketmetricsconfiguration.md)

- [PutBucketNotification](../api/api-putbucketnotification.md)

- [PutBucketObjectLockConfiguration](../api/api-putobjectlockconfiguration.md)

- [PutBucketOwnershipControls](../api/api-putbucketownershipcontrols.md)

- [PutBucketPolicy](../api/api-putbucketpolicy.md)

- [PutBucketPublicAccessBlock](../api/api-putpublicaccessblock.md)

- [PutBucketReplication](../api/api-putbucketreplication.md)

- [PutBucketRequestPayment](../api/api-putbucketrequestpayment.md)

- [PutBucketTagging](../api/api-putbuckettagging.md)

- [PutBucketVersioning](../api/api-putbucketversioning.md)

- [PutBucketWebsite](../api/api-putbucketwebsite.md)

- [UpdateBucketMetadataJournalTableConfiguration](../api/api-updatebucketmetadatajournaltableconfiguratione.md)

- [UpdateBucketMetadataInventoryTableConfiguration](../api/api-updatebucketmetadatainventorytableconfiguration.md)

In addition to these API operations, you can also use the [OPTIONS\
object](../api/restoptionsobject.md) object-level action. This action is treated like a
bucket-level action in CloudTrail logging because the action checks the CORS
configuration of a bucket.

###### Note

The HeadBucket API is supported as an Amazon S3 data event in CloudTrail.

## Amazon S3 Express One Zone bucket-level (Regional API endpoint) actions tracked by CloudTrail logging

By default, CloudTrail logs bucket-level actions for directory buckets as management
events. The `eventsource` for CloudTrail management events for S3 Express One Zone
is `s3express.amazonaws.com`.

These following Regional endpoint API operations are logged to CloudTrail.

- [CreateBucket](../api/api-createbucket.md)

- [DeleteBucket](../api/api-deletebucket.md)

- [DeleteBucketPolicy](../api/api-deletebucketpolicy.md)

- [GetBucketPolicy](../api/api-getbucketpolicy.md)

- [PutBucketPolicy](../api/api-putbucketpolicy.md)

- [ListDirectoryBuckets](../api/api-listdirectorybuckets.md)

- [PutBucketEncryption](../api/api-putbucketencryption.md)

- [GetBucketEncryption](../api/api-getbucketencryption.md)

- [DeleteBucketEncryption](../api/api-deletebucketencryption.md)

For more information, see [Logging with AWS CloudTrail for S3 Express One Zone](s3-express-one-zone-logging.md)

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging with CloudTrail

Example log files

All content copied from https://docs.aws.amazon.com/.
