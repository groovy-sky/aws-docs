# Required permissions for Amazon S3 API operations

###### Note

This page is about Amazon S3 policy actions for general purpose buckets. To learn more about Amazon S3 policy actions for directory buckets, see
[Actions for directory buckets](s3-express-security-iam.md#s3-express-security-iam-actions).

To perform an S3 API operation, you must have the right permissions. This page maps S3 API operations to the required permissions. To grant permissions to
perform an S3 API operation, you must compose a valid policy (such as an S3 bucket
policy or IAM identity-based policy), and specify corresponding actions in the
`Action` element of the policy. These actions are called policy actions.
Not every S3 API operation is represented by a single permission (a single policy action), and some permissions (some policy actions) are required for many different API operations.

When you compose policies, you must specify the `Resource` element based on the correct resource type required by the corresponding Amazon S3 policy actions. This page categorizes permissions to S3 API operations by the resource types.
For more information about the resource types, see
[Resource types defined by Amazon S3](../../../service-authorization/latest/reference/list-amazons3.md#amazons3-resources-for-iam-policies) in the _Service Authorization_
_Reference_. For a full list of Amazon S3 policy actions, resources, and condition keys for use in policies, see [Actions, resources, and condition keys for Amazon S3](../../../service-authorization/latest/reference/list-amazons3.md) in the _Service Authorization_
_Reference_. For a complete list of Amazon S3 API operations, see [Amazon S3 API Actions](../api/api-operations.md) in the _Amazon Simple Storage Service API Reference_.

For more information on how to address the HTTP `403 Forbidden` errors in S3, see [Troubleshoot access denied (403 Forbidden) errors in Amazon S3](troubleshoot-403-errors.md). For
more information on the IAM features to use with S3, see [How Amazon S3 works with IAM](security-iam-service-with-iam.md). For more information on S3 security best practices,
see [Security best practices for Amazon S3](security-best-practices.md).

###### Topics

- [Bucket operations and permissions](#using-with-s3-policy-actions-related-to-buckets)

- [Object operations and permissions](#using-with-s3-policy-actions-related-to-objects)

- [Access point for general purpose buckets operations and permissions](#using-with-s3-policy-actions-related-to-accesspoint)

- [Object Lambda Access Point operations and permissions](#using-with-s3-policy-actions-related-to-olap)

- [Multi-Region Access Point operations and permissions](#using-with-s3-policy-actions-related-to-mrap)

- [Batch job operations and permissions](#using-with-s3-policy-actions-related-to-batchops)

- [S3 Storage Lens configuration operations and permissions](#using-with-s3-policy-actions-related-to-lens)

- [S3 Storage Lens groups operations and permissions](#using-with-s3-policy-actions-related-to-lens-groups)

- [S3 Access Grants instance operations and permissions](#using-with-s3-policy-actions-related-to-s3ag-instances)

- [S3 Access Grants location operations and permissions](#using-with-s3-policy-actions-related-to-s3ag-locations)

- [S3 Access Grants grant operations and permissions](#using-with-s3-policy-actions-related-to-s3ag-grants)

- [Account operations and permissions](#using-with-s3-policy-actions-related-to-accounts)

## Bucket operations and permissions

Bucket operations are S3 API operations that operate on the bucket resource type. You must specify S3 policy actions for bucket operations
in bucket policies or IAM
identity-based policies.

In the policies, the `Resource` element must be the bucket Amazon Resource Name (ARN). For more information about the `Resource` element format and
example policies, see [Bucket operations](security-iam-service-with-iam.md#using-with-s3-actions-related-to-buckets).

###### Note

To grant permissions to bucket operations in access point policies, note the following:

- Permissions granted for bucket operations in an access point policy are effective only if the underlying bucket allows
the same permissions. When you use an access point, you must delegate access
control from the bucket to the access point or add the same permissions in the access point
policy to the underlying bucket's policy.

- In access point policies that grant permissions to bucket operations, the `Resource` element must be the `accesspoint` ARN. For more information about the `Resource` element format and
example policies, see [Bucket operations in policies for access points for general purpose buckets](security-iam-service-with-iam.md#bucket-operations-ap). For more information about access point policies, see [Configuring IAM policies for using access points](access-points-policies.md).

- Not all bucket operations are supported by access points. For more information, see [Access points compatibility with S3 operations](access-points-service-api-support.md#access-points-operations-support).

The following is the mapping of bucket operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateBucket](../api/api-createbucket.md)

(Required) `s3:CreateBucket`

Required to create a new s3 bucket.

(Conditionally required) `s3:PutBucketAcl`

Required if you want to use access control list (ACL) to specify permissions on a bucket when you make a `CreateBucket` request.

(Conditionally required) `s3:PutBucketObjectLockConfiguration`, `s3:PutBucketVersioning`

Required if you want to enable Object Lock when you create a bucket.

(Conditionally required) `s3:PutBucketOwnershipControls`

Required if you want to specify S3 Object Ownership when you create a bucket.

[CreateBucketMetadataConfiguration](../api/api-createbucketmetadataconfiguration.md) (V2 API operation. The IAM
policy action name is the same for the V1 and V2 API operations.)

(Required) `s3:CreateBucketMetadataTableConfiguration`,
`s3tables:CreateTableBucket`, `s3tables:CreateNamespace`,
`s3tables:CreateTable`, `s3tables:GetTable`,
`s3tables:PutTablePolicy`, `s3tables:PutTableEncryption`,
`kms:DescribeKey`

Required to create a metadata table configuration on a general purpose bucket.

To create your AWS managed table bucket and the metadata tables that are specified in
your metadata table configuration, you must have the specified `s3tables`
permissions.

If you want to encrypt your metadata tables with server-side encryption with AWS Key Management Service
(AWS KMS) keys (SSE-KMS), you need additional permissions in your KMS key policy. For more
information, see [Setting up permissions for configuring metadata tables](metadata-tables-permissions.md).

If you also want to integrate your AWS managed table bucket with AWS analytics
services so that you can query your metadata table, you need additional permissions. For more
information, see [Integrating Amazon S3 Tables with\
AWS analytics services](s3-tables-integrating-aws.md).

[CreateBucketMetadataTableConfiguration](../api/api-createbucketmetadatatableconfiguration.md) (V1 API operation)

(Required) `s3:CreateBucketMetadataTableConfiguration`,
`s3tables:CreateNamespace`, `s3tables:CreateTable`,
`s3tables:GetTable`, `s3tables:PutTablePolicy`

Required to create a metadata table configuration on a general purpose bucket.

To create the metadata table in the table bucket that's specified in your metadata table
configuration, you must have the specified `s3tables` permissions.

If you want to encrypt your metadata tables with server-side encryption with AWS Key Management Service
(AWS KMS) keys (SSE-KMS), you need additional permissions. For more information, see [Setting up permissions for configuring metadata tables](metadata-tables-permissions.md).

If you also want to integrate your table bucket with AWS analytics services so that you
can query your metadata table, you need additional permissions. For more information, see
[Integrating Amazon S3 Tables with AWS analytics services](s3-tables-integrating-aws.md).

[DeleteBucket](../api/api-deletebucket.md)

(Required) `s3:DeleteBucket`

Required to delete an S3 bucket.

[DeleteBucketAnalyticsConfiguration](../api/api-deletebucketanalyticsconfiguration.md)

(Required) `s3:PutAnalyticsConfiguration`

Required to delete an S3 analytics configuration from an S3 bucket.

[DeleteBucketCors](../api/api-deletebucketcors.md)

(Required) `s3:PutBucketCORS`

Required to delete the cross-origin resource sharing (CORS) configuration for an bucket.

[DeleteBucketEncryption](../api/api-deletebucketencryption.md)

(Required) `s3:PutEncryptionConfiguration`

Required to reset the default encryption configuration for an S3 bucket as server-side encryption with Amazon S3 managed keys (SSE-S3).

[DeleteBucketIntelligentTieringConfiguration](../api/api-deletebucketintelligenttieringconfiguration.md)

(Required) `s3:PutIntelligentTieringConfiguration`

Required to delete the existing S3 Intelligent-Tiering configuration from an S3 bucket.

[DeleteBucketInventoryConfiguration](../api/api-deletebucketinventoryconfiguration.md)

(Required) `s3:PutInventoryConfiguration`

Required to delete an S3 Inventory configuration from an S3 bucket.

[DeleteBucketLifecycle](../api/api-deletebucketlifecycle.md)

(Required) `s3:PutLifecycleConfiguration`

Required to delete the S3 Lifecycle configuration for an S3 bucket.

[DeleteBucketMetadataConfiguration](../api/api-deletebucketmetadatatableconfiguration.md) (V2 API operation. The IAM
policy action name is the same for the V1 and V2 API operations.)

(Required) `s3:DeleteBucketMetadataTableConfiguration`

Required to delete a metadata table configuration from a general purpose bucket.

[DeleteBucketMetadataTableConfiguration](../api/api-deletebucketmetadatatableconfiguration.md) (V1 API operation)

(Required) `s3:DeleteBucketMetadataTableConfiguration`

Required to delete a metadata table configuration from a general purpose bucket.

[DeleteBucketMetricsConfiguration](../api/api-deletebucketmetricsconfiguration.md)

(Required) `s3:PutMetricsConfiguration`

Required to delete a metrics configuration for the Amazon CloudWatch request metrics from an S3 bucket.

[DeleteBucketOwnershipControls](../api/api-deletebucketownershipcontrols.md)

(Required) `s3:PutBucketOwnershipControls`

Required to remove the Object Ownership setting for an S3 bucket. After removal, the Object Ownership setting becomes `Object writer`.

[DeleteBucketPolicy](../api/api-deletebucketpolicy.md)

(Required) `s3:DeleteBucketPolicy`

Required to delete the policy of an S3 bucket.

[DeleteBucketReplication](../api/api-deletebucketreplication.md)

(Required) `s3:PutReplicationConfiguration`

Required to delete the replication configuration of an S3 bucket.

[DeleteBucketTagging](../api/api-deletebuckettagging.md)

(Required) `s3:PutBucketTagging`

Required to delete tags from an S3 bucket.

[DeleteBucketWebsite](../api/api-deletebucketwebsite.md)

(Required) `s3:DeleteBucketWebsite`

Required to remove the website configuration for an S3 bucket.

[DeletePublicAccessBlock](../api/api-deletepublicaccessblock.md) (Bucket-level)

(Required) `s3:PutBucketPublicAccessBlock`

Required to remove the block public access configuration for an S3 bucket.

[GetBucketAccelerateConfiguration](../api/api-getbucketaccelerateconfiguration.md)

(Required) `s3:GetAccelerateConfiguration`

Required to use the accelerate subresource to return the Amazon S3 Transfer Acceleration state of a bucket, which is either Enabled or Suspended.

[GetBucketAcl](../api/api-getbucketacl.md)

(Required) `s3:GetBucketAcl`

Required to return the access control list (ACL) of an S3 bucket.

[GetBucketAnalyticsConfiguration](../api/api-getbucketanalyticsconfiguration.md)

(Required) `s3:GetAnalyticsConfiguration`

Required to return an analytics configuration that's identified by the analytics configuration ID from an S3 bucket.

[GetBucketCors](../api/api-getbucketcors.md)

(Required) `s3:GetBucketCORS`

Required to return the cross-origin resource sharing (CORS) configuration for an S3 bucket.

[GetBucketEncryption](../api/api-getbucketencryption.md)

(Required) `s3:GetEncryptionConfiguration`

Required to return the default encryption configuration for an S3 bucket.

[GetBucketIntelligentTieringConfiguration](../api/api-getbucketintelligenttieringconfiguration.md)

(Required) `s3:GetIntelligentTieringConfiguration`

Required to get the S3 Intelligent-Tiering configuration of an S3 bucket.

[GetBucketInventoryConfiguration](../api/api-getbucketinventoryconfiguration.md)

(Required) `s3:GetInventoryConfiguration`

Required to return an inventory configuration that's identified by the inventory configuration ID from the bucket.

[GetBucketLifecycle](../api/api-getbucketlifecycle.md)

(Required) `s3:GetLifecycleConfiguration`

Required to return the S3 Lifecycle configuration of the bucket.

[GetBucketLocation](../api/api-getbucketlocation.md)

(Required) `s3:GetBucketLocation`

Required to return the AWS Region that an S3 bucket resides in.

[GetBucketLogging](../api/api-getbucketlogging.md)

(Required) `s3:GetBucketLogging`

Required to return the logging status of an S3 bucket and the permissions that users have to view and modify that status.

[GetBucketMetadataConfiguration](../api/api-getbucketmetadatatableconfiguration.md) (V2 API operation. The IAM policy
action name is the same for the V1 and V2 API operations.)

(Required) `s3:GetBucketMetadataTableConfiguration`

Required to retrieve a metadata table configuration for a general purpose
bucket.

[GetBucketMetadataTableConfiguration](../api/api-getbucketmetadatatableconfiguration.md) (V1 API operation)

(Required) `s3:GetBucketMetadataTableConfiguration`

Required to retrieve a metadata table configuration for a general purpose bucket.

[GetBucketMetricsConfiguration](../api/api-getbucketmetricsconfiguration.md)

(Required) `s3:GetMetricsConfiguration`

Required to get a metrics configuration that's specified by the metrics configuration ID from the bucket.

[GetBucketNotificationConfiguration](../api/api-getbucketnotificationconfiguration.md)

(Required) `s3:GetBucketNotification`

Required to return the notification configuration of an S3 bucket.

[GetBucketOwnershipControls](../api/api-getbucketownershipcontrols.md)

(Required) `s3:GetBucketOwnershipControls`

Required to retrieve the Object Ownership setting for an S3 bucket.

[GetBucketPolicy](../api/api-getbucketpolicy.md)

(Required) `s3:GetBucketPolicy`

Required to return the policy of an S3 bucket.

[GetBucketPolicyStatus](../api/api-getbucketpolicystatus.md)

(Required) `s3:GetBucketPolicyStatus`

Required to retrieve the policy status for an S3 bucket, indicating whether the bucket is public.

[GetBucketReplication](../api/api-getbucketreplication.md)

(Required) `s3:GetReplicationConfiguration`

Required to return the replication configuration of an S3 bucket.

[GetBucketRequestPayment](../api/api-getbucketrequestpayment.md)

(Required) `s3:GetBucketRequestPayment`

Required to return the request payment configuration for an S3 bucket.

[GetBucketVersioning](../api/api-getbucketversioning.md)

(Required) `s3:GetBucketVersioning`

Required to return the versioning state of an S3 bucket.

[GetBucketTagging](../api/api-getbuckettagging.md)

(Required) `s3:GetBucketTagging`

Required to return the tag set that's associated with an S3 bucket.

[GetBucketWebsite](../api/api-getbucketwebsite.md)

(Required) `s3:GetBucketWebsite`

Required to return the website configuration for an S3 bucket.

[GetObjectLockConfiguration](../api/api-getobjectlockconfiguration.md)

(Required) `s3:GetBucketObjectLockConfiguration`

Required to get the Object Lock configuration for an S3 bucket.

[GetPublicAccessBlock](../api/api-getpublicaccessblock.md) (Bucket-level)

(Required) `s3:GetBucketPublicAccessBlock`

Required to retrieve the block public access configuration for an S3 bucket.

[HeadBucket](../api/api-headbucket.md)

(Required) `s3:ListBucket`

Required to determine if a bucket exists and if you have permission to access it.

[ListBucketAnalyticsConfigurations](../api/api-listbucketanalyticsconfigurations.md)

(Required) `s3:GetAnalyticsConfiguration`

Required to list the analytics configurations for an S3 bucket.

[ListBucketIntelligentTieringConfigurations](../api/api-listbucketintelligenttieringconfigurations.md)

(Required) `s3:GetIntelligentTieringConfiguration`

Required to list the S3 Intelligent-Tiering configurations of an S3 bucket.

[ListBucketInventoryConfigurations](../api/api-listbucketinventoryconfigurations.md)

(Required) `s3:GetInventoryConfiguration`

Required to return a list of inventory configurations for an S3 bucket.

[ListBucketMetricsConfigurations](../api/api-listbucketmetricsconfigurations.md)

(Required) `s3:GetMetricsConfiguration`

Required to list the metrics configurations for an S3 bucket.

[ListObjects](../api/api-listobjects.md)

(Required) `s3:ListBucket`

Required to list some or all (up to 1,000) of the objects in an S3 bucket.

(Conditionally required) `s3:GetObjectAcl`

Required if you want to display object owner information.

[ListObjectsV2](../api/api-listobjectsv2.md)

(Required) `s3:ListBucket`

Required to list some or all (up to 1,000) of the objects in an S3 bucket.

(Conditionally required) `s3:GetObjectAcl`

Required if you want to display object owner information.

[ListObjectVersions](../api/api-listobjectversions.md)

(Required) `s3:ListBucketVersions`

Required to get metadata about all the versions of objects in an S3 bucket.

[PutBucketAccelerateConfiguration](../api/api-putbucketaccelerateconfiguration.md)

(Required) `s3:PutAccelerateConfiguration`

Required to set the accelerate configuration of an existing bucket.

[PutBucketAcl](../api/api-putbucketacl.md)

(Required) `s3:PutBucketAcl`

Required to use access control lists (ACLs) to set the permissions on an existing bucket.

[PutBucketAnalyticsConfiguration](../api/api-putbucketanalyticsconfiguration.md)

(Required) `s3:PutAnalyticsConfiguration`

Required to set an analytics configuration for an S3 bucket.

[PutBucketCors](../api/api-putbucketcors.md)

(Required) `s3:PutBucketCORS`

Required to set the cross-origin resource sharing (CORS) configuration for an S3 bucket.

[PutBucketEncryption](../api/api-putbucketencryption.md)

(Required) `s3:PutEncryptionConfiguration`

Required to configure the default encryption for an S3 bucket.

[PutBucketIntelligentTieringConfiguration](../api/api-putbucketintelligenttieringconfiguration.md)

(Required) `s3:PutIntelligentTieringConfiguration`

Required to put the S3 Intelligent-Tiering configuration to an S3 bucket.

[PutBucketInventoryConfiguration](../api/api-putbucketinventoryconfiguration.md)

(Required) `s3:PutInventoryConfiguration`

Required to add an inventory configuration to an S3 bucket.

[PutBucketLifecycle](../api/api-putbucketlifecycle.md)

(Required) `s3:PutLifecycleConfiguration`

Required to create a new S3 Lifecycle configuration or replace an existing lifecycle configuration for an S3 bucket.

[PutBucketLogging](../api/api-putbucketlogging.md)

(Required) `s3:PutBucketLogging`

Required to set the logging parameters for an S3 bucket and specify permissions for who can view and modify the logging parameters.

[PutBucketMetricsConfiguration](../api/api-putbucketmetricsconfiguration.md)

(Required) `s3:PutMetricsConfiguration`

Required to set or update a metrics configuration for the Amazon CloudWatch request metrics of an S3 bucket.

[PutBucketNotificationConfiguration](../api/api-putbucketnotificationconfiguration.md)

(Required) `s3:PutBucketNotification`

Required to enable notifications of specified events for an S3 bucket.

[PutBucketOwnershipControls](../api/api-putbucketownershipcontrols.md)

(Required) `s3:PutBucketOwnershipControls`

Required to create or modify the Object Ownership setting for an S3 bucket.

[PutBucketPolicy](../api/api-putbucketpolicy.md)

(Required) `s3:PutBucketPolicy`

Required to apply an S3 bucket policy to a bucket.

[PutBucketReplication](../api/api-putbucketreplication.md)

(Required) `s3:PutReplicationConfiguration`

Required to create a new replication configuration or replace an existing one for an S3 bucket.

[PutBucketRequestPayment](../api/api-putbucketrequestpayment.md)

(Required) `s3:PutBucketRequestPayment`

Required to set the request payment configuration for a bucket.

[PutBucketTagging](../api/api-putbuckettagging.md)

(Required) `s3:PutBucketTagging`

Required to add a set of tags to an S3 bucket.

[PutBucketVersioning](../api/api-putbucketversioning.md)

(Required) `s3:PutBucketVersioning`

Required to set the versioning state of an S3 bucket.

[PutBucketWebsite](../api/api-putbucketwebsite.md)

(Required) `s3:PutBucketWebsite`

Required to configure a bucket as a website and set the configuration of the website.

[PutObjectLockConfiguration](../api/api-putobjectlockconfiguration.md)

(Required) `s3:PutBucketObjectLockConfiguration`

Required to put Object Lock configuration on an S3 bucket.

[PutPublicAccessBlock](../api/api-putpublicaccessblock.md) (Bucket-level)

(Required) `s3:PutBucketPublicAccessBlock`

Required to create or modify the block public access configuration for an S3 bucket.

[UpdateBucketMetadataInventoryTableConfiguration](../api/api-updatebucketmetadatainventorytableconfiguration.md)

(Required) `s3:UpdateBucketMetadataInventoryTableConfiguration`,
`s3tables:CreateTableBucket`, `s3tables:CreateNamespace`,
`s3tables:CreateTable`, `s3tables:GetTable`,
`s3tables:PutTablePolicy`, `s3tables:PutTableEncryption`,
`kms:DescribeKey`

Required to enable or disable an inventory table for a metadata table configuration on a
general purpose bucket.

If you want to encrypt your inventory table with server-side encryption with AWS Key Management Service
(AWS KMS) keys (SSE-KMS), you need additional permissions in your KMS key policy. For more
information, see [Setting up permissions for configuring metadata tables](metadata-tables-permissions.md).

[UpdateBucketMetadataJournalTableConfiguration](../api/api-updatebucketmetadatajournaltableconfiguration.md)

(Required) `s3:UpdateBucketMetadataJournalTableConfiguration`

Required to enable or disable journal table record expiration for a metadata table
configuration on a general purpose bucket.

## Object operations and permissions

Object operations are S3 API operations that operate on the object resource type. You must specify S3 policy actions for object operations in resource-based policies (such as bucket policies, access point policies, Multi-Region Access Point policies, VPC endpoint policies) or IAM
identity-based policies.

In the policies, the `Resource` element must be the object ARN. For more information about the `Resource` element format and
example policies, see [Object operations](security-iam-service-with-iam.md#using-with-s3-actions-related-to-objects).

###### Note

- AWS KMS policy actions ( `kms:GenerateDataKey` and `kms:Decrypt`) are only applicable for the AWS KMS resource type and must be specified in IAM identity-based policies and AWS KMS resource-based policies (AWS KMS key policies).
You can't specify AWS KMS policy actions in S3 resource-based policies, such as S3 bucket policies.

- When you use access points to control access to object operations, you can use
access point policies. To grant permissions to object operations in access point policies, note the following:

- In access point policies that grant permissions to object operations, the `Resource` element must be the ARNs for objects accessed through an access point. For more information about the `Resource` element format and example
policies, see [Object operations in access point policies](security-iam-service-with-iam.md#object-operations-ap).

- Not all object operations are supported by access points. For more information, see [Access points compatibility with S3 operations](access-points-service-api-support.md#access-points-operations-support).

- Not all object operations are supported by Multi-Region Access Points. For more information, see [Multi-Region Access Point compatibility with S3 operations](mrapoperations.md#mrap-operations-support).

The following is the mapping of object operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[AbortMultipartUpload](../api/api-abortmultipartupload.md)

(Required) `s3:AbortMultipartUpload`

Required to abort a multipart upload.

[CompleteMultipartUpload](../api/api-completemultipartupload.md)

(Required) `s3:PutObject`

Required to complete a multipart upload.

(Conditionally required) `kms:Decrypt`

Required if you want to complete a multipart upload for an AWS KMS customer managed key encrypted object.

[CopyObject](../api/api-copyobject.md)

For source object:

For source object:

(Required) Either `s3:GetObject` or `s3:GetObjectVersion`

- `s3:GetObject` – Required if you want to copy an object from the source bucket without specifying `versionId` in the request.

- `s3:GetObjectVersion` – Required if you want to copy a specific version of an object from the source bucket by specifying `versionId` in the request.

(Conditionally required) `kms:Decrypt`

Required if you want to copy an AWS KMS customer managed key encrypted object from the source bucket.

For destination object:

For destination object:

(Required) `s3:PutObject`

Required to put the copied object in the destination bucket.

(Conditionally required) `s3:PutObjectAcl`

Required if you want to put the copied object with the object access control list (ACL) to the destination bucket when you make a `CopyObject` request.

(Conditionally required) `s3:PutObjectTagging`

Required if you want to put the copied object with object tagging to the destination bucket when you make a `CopyObject` request.

(Conditionally required) `kms:GenerateDataKey`

Required if you want to encrypt the copied object with an AWS KMS customer managed key and put it to the destination bucket.

(Conditionally required) `s3:PutObjectRetention`

Required if you want to set an Object Lock retention configuration for the new object.

(Conditionally required) `s3:PutObjectLegalHold`

Required if you want to place an Object Lock legal hold on the new object.

[CreateMultipartUpload](../api/api-createmultipartupload.md)

(Required) `s3:PutObject`

Required to create multipart upload.

(Conditionally required) `s3:PutObjectAcl`

Required if you want to set the object access control list (ACL) permissions for the uploaded object.

(Conditionally required) `s3:PutObjectTagging`

Required if you want to add object tagging(s) to the uploaded object.

(Conditionally required) `kms:GenerateDataKey`

Required if you want to use an AWS KMS customer managed key to encrypt an object when you initiate a multipart upload.

(Conditionally required) `s3:PutObjectRetention`

Required if you want to set an Object Lock retention configuration for the uploaded object.

(Conditionally required) `s3:PutObjectLegalHold`

Required if you want to apply an Object Lock legal hold to the uploaded object.

[DeleteObject](../api/api-deleteobject.md)

(Required) Either `s3:DeleteObject` or `s3:DeleteObjectVersion`

- `s3:DeleteObject` – Required if you want to remove an object without specifying `versionId` in the request.

- `s3:DeleteObjectVersion` – Required if you want to remove a specific version of an object by specifying `versionId` in the request.

(Conditionally required) `s3:BypassGovernanceRetention`

Required if you want to delete an object that's protected by governance mode for Object Lock retention.

[DeleteObjects](../api/api-deleteobjects.md)

(Required) Either `s3:DeleteObject` or `s3:DeleteObjectVersion`

- `s3:DeleteObject` – Required if you want to remove an object without specifying `versionId` in the request.

- `s3:DeleteObjectVersion` – Required if you want to remove a specific version of an object by specifying `versionId` in the request.

(Conditionally required) `s3:BypassGovernanceRetention`

Required if you want to delete objects that are protected by governance mode for Object Lock retention.

[DeleteObjectTagging](../api/api-deleteobjecttagging.md)

(Required) Either `s3:DeleteObjectTagging` or `s3:DeleteObjectVersionTagging`

- `s3:DeleteObjectTagging` – Required if you want to remove the entire tag set of an object without specifying `versionId` in the request.

- `s3:DeleteObjectVersionTagging` – Required if you want to delete tags of a specific object version by specifying `versionId` in the request.

[GetObject](../api/api-getobject.md)

(Required) Either `s3:GetObject` or `s3:GetObjectVersion`

- `s3:GetObject` – Required if you want to get an object without specifying `versionId` in the request.

- `s3:GetObjectVersion` – Required if you want to get a specific version of an object by specifying `versionId` in the request.

(Conditionally required) `kms:Decrypt`

Required if you want to get and decrypt an AWS KMS customer managed key encrypted object.

(Conditionally required) `s3:GetObjectTagging`

Required if you want to get the tag-set of an object when you make a `GetObject` request.

(Conditionally required) `s3:GetObjectLegalHold`

Required if you want to get an object's current Object Lock legal hold status.

(Conditionally required) `s3:GetObjectRetention`

Required if you want to retrieve the Object Lock retention settings for an object.

[GetObjectAcl](../api/api-getobjectacl.md)

(Required) Either `s3:GetObjectAcl` or `s3:GetObjectVersionAcl`

- `s3:GetObjectAcl` – Required if you want to get the access control list (ACL) of an object without specifying `versionId` in the request.

- `s3:GetObjectVersionAcl` – Required if you want to get the access control list (ACL) of an object by specifying `versionId` in the request.

[GetObjectAttributes](../api/api-getobjectattributes.md)

(Required) Either `s3:GetObject` or `s3:GetObjectVersion`

- `s3:GetObject` – Required if you want to retrieve attributes related to an object without specifying `versionId` in the request.

- `s3:GetObjectVersion` – Required if you want to retrieve attributes related to a specific object version by specifying `versionId` in the request.

(Conditionally required) `kms:Decrypt`

Required if you want to retrieve attributes related to an AWS KMS customer managed key encrypted object.

[GetObjectLegalHold](../api/api-getobjectlegalhold.md)

(Required) `s3:GetObjectLegalHold`

Required to get an object's current Object Lock legal hold status.

[GetObjectRetention](../api/api-getobjectretention.md)

(Required) `s3:GetObjectRetention`

Required to retrieve the Object Lock retention settings for an object.

[GetObjectTagging](../api/api-getobjecttagging.md)

(Required) Either `s3:GetObjectTagging` or `s3:GetObjectVersionTagging`

- `s3:GetObjectTagging` – Required if you want to get the tag set of an object without specifying `versionId` in the request.

- `s3:GetObjectVersionTagging` – Required if you want to get the tags of a specific object version by specifying `versionId` in the request.

[GetObjectTorrent](../api/api-getobjecttorrent.md)

(Required) `s3:GetObject`

Required to return torrent files of an object.

[HeadObject](../api/api-headobject.md)

(Required) `s3:GetObject`

Required to retrieve metadata from an object without returning the object itself.

(Conditionally required) `s3:GetObjectLegalHold`

Required if you want to get an object's current Object Lock legal hold status.

(Conditionally required) `s3:GetObjectRetention`

Required if you want to retrieve the Object Lock retention settings for an object.

[ListMultipartUploads](../api/api-listmultipartuploads.md)

(Required) `s3:ListBucketMultipartUploads`

Required to list in-progress multipart uploads in a bucket.

[ListParts](../api/api-listparts.md)

(Required) `s3:ListMultipartUploadParts`

Required to list the parts that have been uploaded for a specific multipart upload.

(Conditionally required) `kms:Decrypt`

Required if you want to list parts of an AWS KMS customer managed key encrypted multipart upload.

[PutObject](../api/api-putobject.md)

(Required) `s3:PutObject`

Required to put an object.

(Conditionally required) `s3:PutObjectAcl`

Required if you want to put the object access control list (ACL) when you make a `PutObject` request.

(Conditionally required) `s3:PutObjectTagging`

Required if you want to put object tagging when you make a `PutObject` request.

(Conditionally required) `kms:GenerateDataKey`

Required if you want to encrypt an object with an AWS KMS customer managed key.

(Conditionally required) `s3:PutObjectRetention`

Required if you want to set an Object Lock retention configuration on an object.

(Conditionally required) `s3:PutObjectLegalHold`

Required if you want to apply an Object Lock legal hold configuration to a specified object.

[PutObjectAcl](../api/api-putobjectacl.md)

(Required) Either `s3:PutObjectAcl` or `s3:PutObjectVersionAcl`

- `s3:PutObjectAcl` – Required if you want to set the access control list (ACL) permissions for a new or existing object without specifying `versionId` in the request.

- `s3:PutObjectVersionAcl` – Required if you want to set the access control list (ACL) permissions for a new or existing object by specifying `versionId` in the request.

[PutObjectLegalHold](../api/api-putobjectlegalhold.md)

(Required) `s3:PutObjectLegalHold`

Required to apply an Object Lock legal hold configuration to an object.

[PutObjectRetention](../api/api-putobjectretention.md)

(Required) `s3:PutObjectRetention`

Required to apply an Object Lock retention configuration to an object.

(Conditionally required) `s3:BypassGovernanceRetention`

Required if you want to bypass the governance mode of an Object Lock retention configuration.

[PutObjectTagging](../api/api-putobjecttagging.md)

(Required) Either `s3:PutObjectTagging` or `s3:PutObjectVersionTagging`

- `s3:PutObjectTagging` – Required if you want to set the supplied tag set to an object that already exists in a bucket without specifying `versionId` in the request.

- `s3:PutObjectVersionTagging` – Required if you want to set the supplied tag set to an object that already exists in a bucket by specifying `versionId` in the request.

[RestoreObject](../api/api-restoreobject.md)

(Required) `s3:RestoreObject`

Required to restore a copy of an archived object.

[SelectObjectContent](../api/api-selectobjectcontent.md)

(Required) `s3:GetObject`

Required to filter the contents of an S3 object based on a simple structured query language (SQL) statement.

(Conditionally required) `kms:Decrypt`

Required if you want to filter the contents of an S3 object that's encrypted with an AWS KMS customer managed key.

[UpdateObjectEncryption](../api/api-updateobjectencryption.md)

(Required) `s3:UpdateObjectEncryption`, `s3:PutObject`,
`kms:Encrypt`, `kms:Decrypt`, `kms:GenerateDataKey`,
`kms:ReEncrypt*`

Required if you want to change encrypted objects between server-side encryption with Amazon S3
managed encryption (SSE-S3) and server-side encryption with AWS Key Management Service (AWS KMS) encryption keys
(SSE-KMS). You can also use the `UpdateObjectEncryption` operation to apply S3 Bucket Keys to
reduce AWS KMS request costs or change the customer-managed KMS key that's used to encrypt your
data so that you can comply with custom key-rotation standards.

(Conditionally required) `organizations:DescribeAccount`

If you're using AWS Organizations, to use the `UpdateObjectEncryption` operation with customer-managed
KMS keys from other AWS accounts within your organization, you must have the
`organizations:DescribeAccount` permission.

###### Note

You must also request the ability to use AWS KMS keys owned by other member accounts
within your organization by contacting AWS Support.

[UploadPart](../api/api-uploadpart.md)

(Required) `s3:PutObject`

Required to upload a part in a multipart upload.

(Conditionally required) `kms:GenerateDataKey`

Required if you want to put an upload part and encrypt it with an AWS KMS customer managed key.

[UploadPartCopy](../api/api-uploadpartcopy.md)

For source object:

For source object:

(Required) Either `s3:GetObject` or `s3:GetObjectVersion`

- `s3:GetObject` – Required if you want to copy an object from the source bucket without specifying `versionId` in the request.

- `s3:GetObjectVersion` – Required if you want to copy a specific version of an object from the source bucket by specifying `versionId` in the request.

(Conditionally required) `kms:Decrypt`

Required if you want to copy an AWS KMS customer managed key encrypted object from the source bucket.

For destination part:

For destination part:

(Required) `s3:PutObject`

Required to upload a multipart upload part to the destination bucket.

(Conditionally required) `kms:GenerateDataKey`

Required if you want to encrypt a part with an AWS KMS customer managed key when you upload the part to the destination bucket.

## Access point for general purpose buckets operations and permissions

Access point operations are S3 API operations that operate on the `accesspoint`
resource type. You must specify S3
policy actions for access point operations in IAM identity-based
policies, not in bucket policies or access point policies.

In the policies, the `Resource` element must be the `accesspoint` ARN. For more information about the `Resource` element format and
example policies, see [Access point for general purpose bucket operations](security-iam-service-with-iam.md#using-with-s3-actions-related-to-accesspoint).

###### Note

If you want to use access points to control access to bucket or object operations, note the following:

- For using access points to control access to bucket operations, see [Bucket operations in policies for access points for general purpose buckets](security-iam-service-with-iam.md#bucket-operations-ap).

- For using access points to control access to object operations, see [Object operations in access point policies](security-iam-service-with-iam.md#object-operations-ap).

- For more information about how to configure access point policies, see [Configuring IAM policies for using access points](access-points-policies.md).

The following is the mapping of access point operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateAccessPoint](../api/api-control-createaccesspoint.md)

(Required) `s3:CreateAccessPoint`

Required to create an access point that's associated with an S3 bucket.

[DeleteAccessPoint](../api/api-control-deleteaccesspoint.md)

(Required) `s3:DeleteAccessPoint`

Required to delete an access point.

[DeleteAccessPointPolicy](../api/api-control-deleteaccesspointpolicy.md)

(Required) `s3:DeleteAccessPointPolicy`

Required to delete an access point policy.

[GetAccessPointPolicy](../api/api-control-getaccesspointpolicy.md)

(Required) `s3:GetAccessPointPolicy`

Required to retrieve an access point policy.

[GetAccessPointPolicyStatus](../api/api-control-getaccesspointpolicystatus.md)

(Required) `s3:GetAccessPointPolicyStatus`

Required to retrieve the information on whether the specified access point currently has a policy that allows public access.

[PutAccessPointPolicy](../api/api-control-putaccesspointpolicy.md)

(Required) `s3:PutAccessPointPolicy`

Required to put an access point policy.

## Object Lambda Access Point operations and permissions

Object Lambda Access Point operations are S3 API operations that operate on the `objectlambdaaccesspoint` resource type. For more information about how to configure policies for Object Lambda Access Point operations, see [Configuring IAM policies for Object Lambda Access Points](olap-policies.md).

The following is the mapping of Object Lambda Access Point operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateAccessPointForObjectLambda](../api/api-control-createaccesspointforobjectlambda.md)

(Required) `s3:CreateAccessPointForObjectLambda`

Required to create an Object Lambda Access Point.

[DeleteAccessPointForObjectLambda](../api/api-control-deleteaccesspointforobjectlambda.md)

(Required) `s3:DeleteAccessPointForObjectLambda`

Required to delete a specified Object Lambda Access Point.

[DeleteAccessPointPolicyForObjectLambda](../api/api-control-deleteaccesspointpolicyforobjectlambda.md)

(Required) `s3:DeleteAccessPointPolicyForObjectLambda`

Required to delete the policy on a specified Object Lambda Access Point.

[GetAccessPointConfigurationForObjectLambda](../api/api-control-getaccesspointconfigurationforobjectlambda.md)

(Required) `s3:GetAccessPointConfigurationForObjectLambda`

Required to retrieve the configuration of the Object Lambda Access Point.

[GetAccessPointForObjectLambda](../api/api-getaccesspointforobjectlambda.md)

(Required) `s3:GetAccessPointForObjectLambda`

Required to retrieve information about the Object Lambda Access Point.

[GetAccessPointPolicyForObjectLambda](../api/api-getaccesspointpolicyforobjectlambda.md)

(Required) `s3:GetAccessPointPolicyForObjectLambda`

Required to return the access point policy that's associated with the specified Object Lambda Access Point.

[GetAccessPointPolicyStatusForObjectLambda](../api/api-getaccesspointpolicystatusforobjectlambda.md)

(Required) `s3:GetAccessPointPolicyStatusForObjectLambda`

Required to return the policy status for a specific Object Lambda Access Point policy.

[PutAccessPointConfigurationForObjectLambda](../api/api-putaccesspointconfigurationforobjectlambda.md)

(Required) `s3:PutAccessPointConfigurationForObjectLambda`

Required to set the configuration of the Object Lambda Access Point.

[PutAccessPointPolicyForObjectLambda](../api/api-putaccesspointpolicyforobjectlambda.md)

(Required) `s3:PutAccessPointPolicyForObjectLambda`

Required to associate an access policy with a specified Object Lambda Access Point.

## Multi-Region Access Point operations and permissions

Multi-Region Access Point operations are S3 API operations that operate on the `multiregionaccesspoint` resource type.
For more information about how to configure policies for Multi-Region Access Point operations, see [Multi-Region Access Point policy examples](multiregionaccesspointpermissions.md#MultiRegionAccessPointPolicyExamples).

The following is the mapping of Multi-Region Access Point operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateMultiRegionAccessPoint](../api/api-control-createmultiregionaccesspoint.md)

(Required) `s3:CreateMultiRegionAccessPoint`

Required to create a Multi-Region Access Point and associate it with S3 buckets.

[DeleteMultiRegionAccessPoint](../api/api-control-deletemultiregionaccesspoint.md)

(Required) `s3:DeleteMultiRegionAccessPoint`

Required to delete a Multi-Region Access Point.

[DescribeMultiRegionAccessPointOperation](../api/api-control-describemultiregionaccesspointoperation.md)

(Required) `s3:DescribeMultiRegionAccessPointOperation`

Required to retrieve the status of an asynchronous request to manage a Multi-Region Access Point.

[GetMultiRegionAccessPoint](../api/api-control-getmultiregionaccesspoint.md)

(Required) `s3:GetMultiRegionAccessPoint`

Required to return configuration information about the specified Multi-Region Access Point.

[GetMultiRegionAccessPointPolicy](../api/api-control-getmultiregionaccesspointpolicy.md)

(Required) `s3:GetMultiRegionAccessPointPolicy`

Required to return the access control policy of the specified Multi-Region Access Point.

[GetMultiRegionAccessPointPolicyStatus](../api/api-control-getmultiregionaccesspointpolicystatus.md)

(Required) `s3:GetMultiRegionAccessPointPolicyStatus`

Required to return the policy status for a specific Multi-Region Access Point about whether the specified Multi-Region Access Point has an access control policy that allows public access.

[GetMultiRegionAccessPointRoutes](../api/api-control-getmultiregionaccesspointroutes.md)

(Required) `s3:GetMultiRegionAccessPointRoutes`

Required to return the routing configuration for a Multi-Region Access Point.

[PutMultiRegionAccessPointPolicy](../api/api-control-putmultiregionaccesspointpolicy.md)

(Required) `s3:PutMultiRegionAccessPointPolicy`

Required to update the access control policy of the specified Multi-Region Access Point.

[SubmitMultiRegionAccessPointRoutes](../api/api-control-submitmultiregionaccesspointroutes.md)

(Required) `s3:SubmitMultiRegionAccessPointRoutes`

Required to submit an updated route configuration for a Multi-Region Access Point.

## Batch job operations and permissions

(Batch Operations) job operations are S3 API operations that operate on the `job` resource type.
You must specify S3 policy
actions for job operations in IAM identity-based policies,
not in bucket policies.

In the policies, the `Resource`
element must be the `job` ARN. For more information about the `Resource` element format and
example policies, see [Batch job operations](security-iam-service-with-iam.md#using-with-s3-actions-related-to-batchops).

The following is the mapping of batch job operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[DeleteJobTagging](../api/api-control-deletejobtagging.md)

(Required) `s3:DeleteJobTagging`

Required to remove tags from an existing S3 Batch Operations job.

[DescribeJob](../api/api-control-describejob.md)

(Required) `s3:DescribeJob`

Required to retrieve the configuration parameters and status for a Batch Operations job.

[GetJobTagging](../api/api-control-getjobtagging.md)

(Required) `s3:GetJobTagging`

Required to return the tag set of an existing S3 Batch Operations job.

[PutJobTagging](../api/api-control-putjobtagging.md)

(Required) `s3:PutJobTagging`

Required to put or replace tags on an existing S3 Batch Operations job.

[UpdateJobPriority](../api/api-control-updatejobpriority.md)

(Required) `s3:UpdateJobPriority`

Required to update the priority of an existing job.

[UpdateJobStatus](../api/api-control-updatejobstatus.md)

(Required) `s3:UpdateJobStatus`

Required to update the status for the specified job.

## S3 Storage Lens configuration operations and permissions

S3 Storage Lens configuration operations are S3 API operations that operate on the `storagelensconfiguration` resource type.
For more information about how to configure S3 Storage Lens configuration operations, see [Setting Amazon S3 Storage Lens permissions](storage-lens-iam-permissions.md).

The following is the mapping of S3 Storage Lens configuration operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[DeleteStorageLensConfiguration](../api/api-control-deletestoragelensconfiguration.md)

(Required) `s3:DeleteStorageLensConfiguration`

Required to delete the S3 Storage Lens configuration.

[DeleteStorageLensConfigurationTagging](../api/api-control-deletestoragelensconfigurationtagging.md)

(Required) `s3:DeleteStorageLensConfigurationTagging`

Required to delete the S3 Storage Lens configuration tags.

[GetStorageLensConfiguration](../api/api-control-getstoragelensconfiguration.md)

(Required) `s3:GetStorageLensConfiguration`

Required to get the S3 Storage Lens configuration.

[GetStorageLensConfigurationTagging](../api/api-control-getstoragelensconfigurationtagging.md)

(Required) `s3:GetStorageLensConfigurationTagging`

Required to get the tags of S3 Storage Lens configuration.

[PutStorageLensConfigurationTagging](../api/api-control-putstoragelensconfigurationtagging.md)

(Required) `s3:PutStorageLensConfigurationTagging`

Required to put or replace tags on an existing S3 Storage Lens configuration.

## S3 Storage Lens groups operations and permissions

S3 Storage Lens groups operations are S3 API operations that operate on the `storagelensgroup` resource type.
For more information about how to configure S3 Storage Lens groups permissions, see [Storage Lens groups permissions](storage-lens-groups.md#storage-lens-group-permissions).

The following is the mapping of S3 Storage Lens groups operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[DeleteStorageLensGroup](../api/api-control-deletestoragelensgroup.md)

(Required) `s3:DeleteStorageLensGroup`

Required to delete an existing S3 Storage Lens group.

[GetStorageLensGroup](../api/api-control-getstoragelensgroup.md)

(Required) `s3:GetStorageLensGroup`

Required to retrieve the S3 Storage Lens group configuration details.

[UpdateStorageLensGroup](../api/api-control-updatestoragelensgroup.md)

(Required) `s3:UpdateStorageLensGroup`

Required to update the existing S3 Storage Lens group.

[CreateStorageLensGroup](../api/api-control-createstoragelensgroup.md)

(Required) `s3:CreateStorageLensGroup`

Required to create a new Storage Lens group.

[CreateStorageLensGroup](../api/api-control-createstoragelensgroup.md), [TagResource](../api/api-control-tagresource.md)

(Required) `s3:CreateStorageLensGroup`,
`s3:TagResource`

Required to create a new Storage Lens group with tags.

[ListStorageLensGroups](../api/api-control-liststoragelensgroups.md)

(Required) `s3:ListStorageLensGroups`

Required to list all Storage Lens groups in your home Region.

[ListTagsForResource](../api/api-control-listtagsforresource.md)

(Required) `s3:ListTagsForResource`

Required to list the tags that were added to your Storage Lens group.

[TagResource](../api/api-control-tagresource.md)

(Required) `s3:TagResource`

Required to add or update a Storage Lens group tag for an existing Storage Lens group.

[UntagResource](../api/api-control-untagresource.md)

(Required) `s3:UntagResource`

Required to delete a tag from a Storage Lens group.

## S3 Access Grants instance operations and permissions

S3 Access Grants instance operations are S3 API operations that operate on the `accessgrantsinstance` resource type. An S3 Access Grants instance is a logical container for your access grants. For more information on working with S3 Access Grants instances, see [Working with S3 Access Grants instances](access-grants-instance.md).

The following is the mapping of the S3 Access Grants instance configuration operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[AssociateAccessGrantsIdentityCenter](../api/api-control-associateaccessgrantsidentitycenter.md)

(Required) `s3:AssociateAccessGrantsIdentityCenter`

Required to associate an AWS IAM Identity Center instance with your S3 Access Grants instance, thus enabling you to create access grants for users and groups in your corporate identity directory. You must also have the following permissions:

`sso:CreateApplication`, `sso:PutApplicationGrant`, and `sso:PutApplicationAuthenticationMethod`.

[CreateAccessGrantsInstance](../api/api-control-createaccessgrantsinstance.md)

(Required) `s3:CreateAccessGrantsInstance`

Required to create an S3 Access Grants instance ( `accessgrantsinstance` resource) which is a container for your individual access grants.

To associate an AWS IAM Identity Center instance with your S3 Access Grants instance, you must also have the `sso:DescribeInstance`, `sso:CreateApplication`, `sso:PutApplicationGrant`, and `sso:PutApplicationAuthenticationMethod` permissions.

[DeleteAccessGrantsInstance](../api/api-control-deleteaccessgrantsinstance.md)

(Required) `s3:DeleteAccessGrantsInstance`

Required to delete an S3 Access Grants instance ( `accessgrantsinstance` resource) from an AWS Region in your account.

[DeleteAccessGrantsInstanceResourcePolicy](../api/api-control-deleteaccessgrantsinstanceresourcepolicy.md)

(Required) `s3:DeleteAccessGrantsInstanceResourcePolicy`

Required to delete a resource policy for your S3 Access Grants instance.

[DissociateAccessGrantsIdentityCenter](../api/api-control-dissociateaccessgrantsidentitycenter.md)

(Required) `s3:DissociateAccessGrantsIdentityCenter`

Required to disassociate an AWS IAM Identity Center instance from your S3 Access Grants instance. You must also have the following permissions:

`sso:DeleteApplication`

[GetAccessGrantsInstance](../api/api-control-getaccessgrantsinstance.md)

(Required) `s3:GetAccessGrantsInstance`

Required to retrieve the S3 Access Grants instance for an AWS Region in your account.

[GetAccessGrantsInstanceForPrefix](../api/api-control-getaccessgrantsinstanceforprefix.md)

(Required) `s3:GetAccessGrantsInstanceForPrefix`

Required to retrieve the S3 Access Grants instance that contains a particular prefix.

[GetAccessGrantsInstanceResourcePolicy](../api/api-control-getaccessgrantsinstanceresourcepolicy.md)

(Required) `s3:GetAccessGrantsInstanceResourcePolicy`

Required to return the resource policy of your S3 Access Grants instance.

[ListAccessGrantsInstances](../api/api-control-listaccessgrantsinstances.md)

(Required) `s3:ListAccessGrantsInstances`

Required to return a list of the S3 Access Grants instances in your account.

[PutAccessGrantsInstanceResourcePolicy](../api/api-control-putaccessgrantsinstanceresourcepolicy.md)

(Required) `s3:PutAccessGrantsInstanceResourcePolicy`

Required to update the resource policy of the S3 Access Grants instance.

## S3 Access Grants location operations and permissions

S3 Access Grants location operations are S3 API operations that operate on the `accessgrantslocation` resource type. For more information on working with S3 Access Grants locations, see [Working with S3 Access Grants locations](access-grants-location.md).

The following is the mapping of the S3 Access Grants location configuration operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateAccessGrantsLocation](../api/api-control-createaccessgrantslocation.md)

(Required) `s3:CreateAccessGrantsLocation`

Required to register a location in your S3 Access Grants instance (create an `accessgrantslocation` resource). You must also have the following permission for the specified IAM role:

`iam:PassRole`

[DeleteAccessGrantsLocation](../api/api-control-deleteaccessgrantslocation.md)

(Required) `s3:DeleteAccessGrantsLocation`

Required to remove a registered location from your S3 Access Grants instance.

[GetAccessGrantsLocation](../api/api-control-getaccessgrantslocation.md)

(Required) `s3:GetAccessGrantsLocation`

Required to retrieve the details of a particular location registered in your S3 Access Grants instance.

[ListAccessGrantsLocations](../api/api-control-listaccessgrantslocations.md)

(Required) `s3:ListAccessGrantsLocations`

Required to return a list of the locations registered in your S3 Access Grants instance.

[UpdateAccessGrantsLocation](../api/api-control-updateaccessgrantslocation.md)

(Required) `s3:UpdateAccessGrantsLocation`

Required to update the IAM role of a registered location in your S3 Access Grants instance.

## S3 Access Grants grant operations and permissions

S3 Access Grants grant operations are S3 API operations that operate on the `accessgrant` resource type. For more information on working with individual grants using S3 Access Grants, see [Working with grants in S3 Access Grants](access-grants-grant.md).

The following is the mapping of the S3 Access Grants grant configuration operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateAccessGrant](../api/api-control-createaccessgrant.md)

(Required) `s3:CreateAccessGrant`

Required to create an individual grant ( `accessgrant` resource) for a user or group in your S3 Access Grants instance. You must also have the following permissions:

For any directory identity — `sso:DescribeInstance` and `sso:DescribeApplication`

For directory users — `identitystore:DescribeUser`

[DeleteAccessGrant](../api/api-control-deleteaccessgrant.md)

(Required) `s3:DeleteAccessGrant`

Required to delete an individual access grant ( `accessgrant` resource) from your S3 Access Grants instance.

[GetAccessGrant](../api/api-control-getaccessgrant.md)

(Required) `s3:GetAccessGrant`

Required to get the details about an individual access grant in your S3 Access Grants instance.

[ListAccessGrants](../api/api-control-listaccessgrants.md)

(Required) `s3:ListAccessGrants`

Required to return a list of individual access grants in your S3 Access Grants instance.

[ListCallerAccessGrants](../api/api-control-listcalleraccessgrants.md)

(Required) `s3:ListCallerAccessGrants`

Required to list the access grants that grant the caller access to Amazon S3 data through S3 Access Grants.

## Account operations and permissions

Account operations are S3 API operations that operate on the account level. Account isn't a resource type defined by Amazon S3.
You must specify S3 policy actions for account operations in IAM identity-based policies, not in bucket policies.

In the policies, the `Resource` element must be `"*"`.
For more information about example policies, see [Account operations](security-iam-service-with-iam.md#using-with-s3-actions-related-to-accounts).

The following is the mapping of account operations and required policy actions.

API operationsPolicy actionsDescription of policy actions

[CreateJob](../api/api-control-createjob.md)

(Required) `s3:CreateJob`

Required to create a new S3 Batch Operations job.

[CreateStorageLensGroup](../api/api-control-createstoragelensgroup.md)

(Required) `s3:CreateStorageLensGroup`

Required to create a new S3 Storage Lens group and associate it with the specified AWS account ID.

(Conditionally required) `s3:TagResource`

Required if you want to create an S3 Storage Lens group with AWS resource tags.

[DeletePublicAccessBlock](../api/api-control-deletepublicaccessblock.md) (Account-level)

(Required) `s3:PutAccountPublicAccessBlock`

Required to remove the block public access configuration from an AWS account.

[GetAccessPoint](../api/api-control-getaccesspoint.md)

(Required) `s3:GetAccessPoint`

Required to retrieve configuration information about the specified access point.

[GetAccessPointPolicy](../api/api-control-getaccesspointpolicy.md) (Account-level)

(Required) `s3:GetAccountPublicAccessBlock`

Required to retrieve the block public access configuration for an AWS account.

[ListAccessPoints](../api/api-control-listaccesspoints.md)

(Required) `s3:ListAccessPoints`

Required to list access points of an S3 bucket that are owned by an AWS account.

[ListAccessPointsForObjectLambda](../api/api-control-listaccesspointsforobjectlambda.md)

(Required) `s3:ListAccessPointsForObjectLambda`

Required to list the Object Lambda Access Points.

[ListBuckets](../api/api-listbuckets.md)

(Required) `s3:ListAllMyBuckets`

Required to return a list of all buckets that are owned by the authenticated sender of the request.

[ListJobs](../api/api-control-listjobs.md)

(Required) `s3:ListJobs`

Required to list current jobs and jobs that have ended recently.

[ListMultiRegionAccessPoints](../api/api-control-listmultiregionaccesspoints.md)

(Required) `s3:ListMultiRegionAccessPoints`

Required to return a list of the Multi-Region Access Points that are currently associated with the specified AWS account.

[ListStorageLensConfigurations](../api/api-control-liststoragelensconfigurations.md)

(Required) `s3:ListStorageLensConfigurations`

Required to get a list of S3 Storage Lens configurations for an AWS account.

[ListStorageLensGroups](../api/api-control-liststoragelensgroups.md)

(Required) `s3:ListStorageLensGroups`

Required to list all the S3 Storage Lens groups in the specified home AWS Region.

[PutPublicAccessBlock](../api/api-putpublicaccessblock.md) (Account-level)

(Required) `s3:PutAccountPublicAccessBlock`

Required to create or modify the block public access configuration for an AWS account.

[PutStorageLensConfiguration](../api/api-control-putstoragelensconfiguration.md)

(Required) `s3:PutStorageLensConfiguration`

Required to put an S3 Storage Lens configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

For an object operation

Policies and permissions

All content copied from https://docs.aws.amazon.com/.
