# Actions, resources, and condition keys for Amazon S3

Amazon S3 (service prefix: `s3`) provides the following service-specific resources, actions, and condition context keys for use in IAM permission policies.

References:

- Learn how to [configure this service](../../../s3/latest/userguide/welcome.md).

- View a list of the [API operations available for this service](../../../s3/latest/api.md).

- Learn how to secure this service and its resources by [using IAM](../../../s3/latest/userguide/access-control-overview.md) permission policies.

###### Topics

- [Actions defined by Amazon S3](#amazons3-actions-as-permissions)

- [Resource types defined by Amazon S3](#amazons3-resources-for-iam-policies)

- [Condition keys for Amazon S3](#amazons3-policy-keys)

## Actions defined by Amazon S3

You can specify the following actions in the `Action` element of an IAM policy statement. Use policies to grant permissions to perform an operation in AWS. When you use an action in a policy, you usually allow or deny access to the API operation or CLI command with the same name. However, in some cases, a single action controls access to more than one operation. Alternatively, some operations require several different actions.

The **Access level** column of the Actions table describes how the action is classified (List, Read, Permissions management, or Tagging). This classification can help you understand the level of access that an action grants when you use it in a policy. For more information about access levels, see [Access levels in policy summaries](../../../iam/latest/userguide/access-policies-understand-policy-summary-access-level-summaries.md).

The **Resource types** column of the Actions table indicates whether each action supports resource-level permissions. If there is no value for this column, you must specify all resources ("\*") to which the policy applies in the `Resource` element of your policy statement. If the column includes a resource type, then you can specify an ARN of that type in a statement with that action. If the action has one or more required resources, the caller must have permission to use the action with those resources. Required resources are indicated in the table with an asterisk (\*). If you limit resource access with the `Resource` element in an IAM policy, you must include an ARN or pattern for each required resource type. Some actions support multiple resource types. If the resource type is optional (not indicated as required), then you can choose to use one of the optional resource types.

The **Condition keys** column of the Actions table includes keys that you can specify in a policy statement's `Condition` element. For more information on the condition keys that are associated with resources for the service, see the **Condition keys** column of the Resource types table.

The **Dependent actions** column of the Actions table shows additional permissions that may be required to successfully call an action. These permissions may be needed in addition to the permission for the action itself. When an action specifies dependent actions, those dependencies may apply to additional resources defined for that action, not only the first resource listed in the table.

###### Note

Resource condition keys are listed in the [Resource types](#amazons3-resources-for-iam-policies) table. You can find a link to the resource type that applies to an action in the **Resource types (\*required)** column of the Actions table. The resource type in the Resource types table includes the **Condition keys** column, which are the resource condition keys that apply to an action in the Actions table.

For details about the columns in the following table, see [Actions table](reference-policies-actions-resources-contextkeys.md#actions_table).

ActionsDescriptionAccess levelResource types (\*required)Condition keysDependent actions[AbortMultipartUpload](../../../s3/latest/api/api-abortmultipartupload.md)Grants permission to abort a multipart uploadWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[AssociateAccessGrantsIdentityCenter](../../../s3/latest/api/api-control-associateaccessgrantsidentitycenter.md)Grants permission to associate Access Grants identity centerPermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[BypassGovernanceRetention](../../../s3/latest/userguide/object-lock-managing.md#object-lock-managing-bypass)Grants permission to allow circumvention of governance-mode object retention settingsPermissions management

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:RequestObjectTag/<key>](#amazons3-s3_RequestObjectTag_key)

[s3:RequestObjectTagKeys](#amazons3-s3_RequestObjectTagKeys)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-copy-source](#amazons3-s3_x-amz-copy-source)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[s3:x-amz-metadata-directive](#amazons3-s3_x-amz-metadata-directive)

[s3:x-amz-server-side-encryption](#amazons3-s3_x-amz-server-side-encryption)

[s3:x-amz-server-side-encryption-aws-kms-key-id](#amazons3-s3_x-amz-server-side-encryption-aws-kms-key-id)

[s3:x-amz-server-side-encryption-customer-algorithm](#amazons3-s3_x-amz-server-side-encryption-customer-algorithm)

[s3:x-amz-storage-class](#amazons3-s3_x-amz-storage-class)

[s3:x-amz-website-redirect-location](#amazons3-s3_x-amz-website-redirect-location)

[CreateAccessGrant](../../../s3/latest/api/api-control-createaccessgrant.md)Grants permission to create Access GrantPermissions management

[accessgrantslocation\*](#amazons3-accessgrantslocation)

[s3:AccessGrantScope](#amazons3-s3_AccessGrantScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[CreateAccessGrantsInstance](../../../s3/latest/api/api-control-createaccessgrantsinstance.md)Grants permission to Create Access Grants InstancePermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[CreateAccessGrantsLocation](../../../s3/latest/api/api-control-createaccessgrantslocation.md)Grants permission to create Access Grants locationPermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:AccessGrantsLocationScope](#amazons3-s3_AccessGrantsLocationScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[CreateAccessPoint](../../../s3/latest/api/api-control-createaccesspoint.md)Grants permission to create a new access pointWrite

[accesspoint\*](#amazons3-accesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:locationconstraint](#amazons3-s3_locationconstraint)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[CreateAccessPointForObjectLambda](../../../s3/latest/api/api-control-createaccesspointforobjectlambda.md)Grants permission to create an object lambda enabled accesspointWrite

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[CreateBucket](../../../s3/latest/api/api-createbucket.md)Grants permission to create a new bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:locationconstraint](#amazons3-s3_locationconstraint)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[s3:x-amz-bucket-namespace](#amazons3-s3_x-amz-bucket-namespace)

[s3:x-amz-object-ownership](#amazons3-s3_x-amz-object-ownership)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[CreateBucketMetadataTableConfiguration](../../../s3/latest/api/api-createbucketmetadataconfiguration.md)Grants permission to create a new S3 Metadata configuration for a specified general purpose bucketWrite

[bucket\*](#amazons3-bucket)

kms:DescribeKey

s3tables:CreateNamespace

s3tables:CreateTable

s3tables:CreateTableBucket

s3tables:GetTable

s3tables:PutTableBucketPolicy

s3tables:PutTableEncryption

s3tables:PutTablePolicy

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[CreateJob](../../../s3/latest/api/api-control-createjob.md)Grants permission to create a new Amazon S3 Batch Operations jobWrite

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:RequestJobPriority](#amazons3-s3_RequestJobPriority)

[s3:RequestJobOperation](#amazons3-s3_RequestJobOperation)

[aws:TagKeys](#amazons3-aws_TagKeys)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

iam:PassRole

[CreateMultiRegionAccessPoint](../../../s3/latest/api/api-control-createmultiregionaccesspoint.md)Grants permission to create a new Multi-Region Access PointWrite

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[CreateStorageLensGroup](../../../s3/latest/api/api-control-createstoragelensgroup.md)Grants permission to create an Amazon S3 Storage Lens groupWrite

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[DeleteAccessGrant](../../../s3/latest/api/api-control-deleteaccessgrant.md)Grants permission to delete Access GrantPermissions management

[accessgrant\*](#amazons3-accessgrant)

[s3:AccessGrantScope](#amazons3-s3_AccessGrantScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessGrantsInstance](../../../s3/latest/api/api-control-deleteaccessgrantsinstance.md)Grants permission to Delete Access Grants InstancePermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessGrantsInstanceResourcePolicy](../../../s3/latest/api/api-control-deleteaccessgrantsinstanceresourcepolicy.md)Grants permission to read Access grants instance resource policyPermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessGrantsLocation](../../../s3/latest/api/api-control-deleteaccessgrantslocation.md)Grants permission to delete Access Grants locationPermissions management

[accessgrantslocation\*](#amazons3-accessgrantslocation)

[s3:AccessGrantsLocationScope](#amazons3-s3_AccessGrantsLocationScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessPoint](../../../s3/latest/api/api-control-deleteaccesspoint.md)Grants permission to delete the access point named in the URIWrite

[accesspoint\*](#amazons3-accesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessPointForObjectLambda](../../../s3/latest/api/api-control-deleteaccesspointforobjectlambda.md)Grants permission to delete the object lambda enabled access point named in the URIWrite

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteAccessPointPolicy](../../../s3/latest/api/api-control-deleteaccesspointpolicy.md)Grants permission to delete the policy on a specified access pointPermissions management

[accesspoint\*](#amazons3-accesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[DeleteAccessPointPolicyForObjectLambda](../../../s3/latest/api/api-control-deleteaccesspointpolicyforobjectlambda.md)Grants permission to delete the policy on a specified object lambda enabled access pointPermissions management

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteBucket](../../../s3/latest/api/api-deletebucket.md)Grants permission to delete the bucket named in the URIWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteBucketMetadataTableConfiguration](../../../s3/latest/api/api-deletebucketmetadataconfiguration.md)Grants permission to delete the S3 Metadata configuration for a specified general purpose bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteBucketPolicy](../../../s3/latest/api/api-deletebucketpolicy.md)Grants permission to delete the policy on a specified bucketPermissions management

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteBucketWebsite](../../../s3/latest/api/api-deletebucketwebsite.md)Grants permission to remove the website configuration for a bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteJobTagging](../../../s3/latest/api/api-control-deletejobtagging.md)Grants permission to remove tags from an existing Amazon S3 Batch Operations jobTagging

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:ExistingJobPriority](#amazons3-s3_ExistingJobPriority)

[s3:ExistingJobOperation](#amazons3-s3_ExistingJobOperation)

[DeleteMultiRegionAccessPoint](../../../s3/latest/api/api-control-deletemultiregionaccesspoint.md)Grants permission to delete the Multi-Region Access Point named in the URIWrite

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[DeleteObject](../../../s3/latest/api/api-deleteobject.md)Grants permission to remove the null version of an object and insert a delete marker, which becomes the current version of the objectWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:if-match](#amazons3-s3_if-match)

[DeleteObjectTagging](../../../s3/latest/api/api-deleteobjecttagging.md)Grants permission to use the tagging subresource to remove the entire tag set from the specified objectTagging

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteObjectVersion](../../../s3/latest/api/api-deleteobject.md)Grants permission to remove a specific version of an objectWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteObjectVersionTagging](../../../s3/latest/api/api-deleteobjecttagging.md)Grants permission to remove the entire tag set for a specific version of the objectTagging

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteStorageLensConfiguration](../../../s3/latest/api/api-control-deletestoragelensconfiguration.md)Grants permission to delete an existing Amazon S3 Storage Lens configurationWrite

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteStorageLensConfigurationTagging](../../../s3/latest/api/api-control-deletestoragelensconfigurationtagging.md)Grants permission to remove tags from an existing Amazon S3 Storage Lens configurationTagging

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DeleteStorageLensGroup](../../../s3/latest/api/api-control-deletestoragelensgroup.md)Grants permission to delete an existing S3 Storage Lens groupWrite

[storagelensgroup\*](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DescribeJob](../../../s3/latest/api/api-control-describejob.md)Grants permission to retrieve the configuration parameters and status for a batch operations jobRead

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[DescribeMultiRegionAccessPointOperation](../../../s3/latest/api/api-control-describemultiregionaccesspointoperation.md)Grants permission to retrieve the configurations for a Multi-Region Access PointRead

[multiregionaccesspointrequestarn\*](#amazons3-multiregionaccesspointrequestarn)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[DissociateAccessGrantsIdentityCenter](../../../s3/latest/api/api-control-dissociateaccessgrantsidentitycenter.md)Grants permission to disassociate Access Grants identity centerPermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccelerateConfiguration](../../../s3/latest/api/api-getbucketaccelerateconfiguration.md)Grants permission to uses the accelerate subresource to return the Transfer Acceleration state of a bucket, which is either Enabled or SuspendedRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAccessGrant](../../../s3/latest/api/api-control-getaccessgrant.md)Grants permission to read Access GrantRead

[accessgrant\*](#amazons3-accessgrant)

[s3:AccessGrantScope](#amazons3-s3_AccessGrantScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessGrantsInstance](../../../s3/latest/api/api-control-getaccessgrantsinstance.md)Grants permission to Read Access Grants InstanceRead

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessGrantsInstanceForPrefix](../../../s3/latest/api/api-control-getaccessgrantsinstanceforprefix.md)Grants permission to Read Access Grants Instance by prefixRead

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessGrantsInstanceResourcePolicy](../../../s3/latest/api/api-control-getaccessgrantsinstanceresourcepolicy.md)Grants permission to read Access grants instance resource policyRead

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessGrantsLocation](../../../s3/latest/api/api-control-getaccessgrantslocation.md)Grants permission to read Access Grants locationRead

[accessgrantslocation\*](#amazons3-accessgrantslocation)

[s3:AccessGrantsLocationScope](#amazons3-s3_AccessGrantsLocationScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessPoint](../../../s3/latest/api/api-control-getaccesspoint.md)Grants permission to return configuration information about the specified access pointRead

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessPointConfigurationForObjectLambda](../../../s3/latest/api/api-control-getaccesspointconfigurationforobjectlambda.md)Grants permission to retrieve the configuration of the object lambda enabled access pointRead

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAccessPointForObjectLambda](../../../s3/latest/api/api-control-getaccesspointforobjectlambda.md)Grants permission to create an object lambda enabled accesspointRead

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAccessPointPolicy](../../../s3/latest/api/api-control-getaccesspointpolicy.md)Grants permission to return the access point policy associated with the specified access pointRead

[accesspoint\*](#amazons3-accesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessPointPolicyForObjectLambda](../../../s3/latest/api/api-control-getaccesspointpolicyforobjectlambda.md)Grants permission to return the access point policy associated with the specified object lambda enabled access pointRead

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAccessPointPolicyStatus](../../../s3/latest/api/api-control-getaccesspointpolicystatus.md)Grants permission to return the policy status for a specific access point policyRead

[accesspoint\*](#amazons3-accesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetAccessPointPolicyStatusForObjectLambda](../../../s3/latest/api/api-control-getaccesspointpolicystatusforobjectlambda.md)Grants permission to return the policy status for a specific object lambda access point policyRead

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAccountPublicAccessBlock](../../../s3/latest/api/api-control-getpublicaccessblock.md)Grants permission to retrieve the PublicAccessBlock configuration for an AWS accountRead

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetAnalyticsConfiguration](../../../s3/latest/api/api-getbucketanalyticsconfiguration.md)Grants permission to get an analytics configuration from an Amazon S3 bucket, identified by the analytics configuration IDRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketAbac](../../../s3/latest/api/api-getbucketabac.md)Grants permission to retrieve ABAC configuration for a general purpose bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketAcl](../../../s3/latest/api/api-getbucketacl.md)Grants permission to use the acl subresource to return the access control list (ACL) of an Amazon S3 bucketRead

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketCORS](../../../s3/latest/api/api-getbucketcors.md)Grants permission to return the CORS configuration information set for an Amazon S3 bucketRead

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketLocation](../../../s3/latest/api/api-getbucketlocation.md)Grants permission to return the Region that an Amazon S3 bucket resides inRead

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketLogging](../../../s3/latest/api/api-getbucketlogging.md)Grants permission to return the logging status of an Amazon S3 bucket and the permissions users have to view or modify that statusRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketMetadataTableConfiguration](../../../s3/latest/api/api-getbucketmetadataconfiguration.md)Grants permission to return the S3 Metadata configuration for a specified general purpose bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketNotification](../../../s3/latest/api/api-getbucketnotification.md)Grants permission to get the notification configuration of an Amazon S3 bucketRead

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketObjectLockConfiguration](../../../s3/latest/api/api-getobjectlockconfiguration.md)Grants permission to get the Object Lock configuration of an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:signatureversion](#amazons3-s3_signatureversion)

[GetBucketOwnershipControls](../../../s3/latest/api/api-getbucketownershipcontrols.md)Grants permission to retrieve ownership controls on a bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketPolicy](../../../s3/latest/api/api-getbucketpolicy.md)Grants permission to return the policy of the specified bucketRead

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketPolicyStatus](../../../s3/latest/api/api-getbucketpolicystatus.md)Grants permission to retrieve the policy status for a specific Amazon S3 bucket, which indicates whether the bucket is publicRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketPublicAccessBlock](../../../s3/latest/api/api-getpublicaccessblock.md)Grants permission to retrieve the PublicAccessBlock configuration for an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketRequestPayment](../../../s3/latest/api/api-getbucketrequestpayment.md)Grants permission to return the request payment configuration for an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketTagging](../../../s3/latest/api/api-getbuckettagging.md)Grants permission to return the tag set associated with an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketVersioning](../../../s3/latest/api/api-getbucketversioning.md)Grants permission to return the versioning state of an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetBucketWebsite](../../../s3/latest/api/api-getbucketwebsite.md)Grants permission to return the website configuration for an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetDataAccess](../../../s3/latest/api/api-control-getdataaccess.md)Grants permission to get AccessRead

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[GetEncryptionConfiguration](../../../s3/latest/api/api-getbucketencryption.md)Grants permission to return the default encryption configuration an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetIntelligentTieringConfiguration](../../../s3/latest/api/api-getbucketintelligenttieringconfiguration.md)Grants permission to get an or list all Amazon S3 Intelligent Tiering configuration in a S3 BucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetInventoryConfiguration](../../../s3/latest/api/api-getbucketinventoryconfiguration.md)Grants permission to return an inventory configuration from an Amazon S3 bucket, identified by the inventory configuration IDRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetJobTagging](../../../s3/latest/api/api-control-getjobtagging.md)Grants permission to return the tag set of an existing Amazon S3 Batch Operations jobRead

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetLifecycleConfiguration](../../../s3/latest/api/api-getbucketlifecycleconfiguration.md)Grants permission to return the lifecycle configuration information set on an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetMetricsConfiguration](../../../s3/latest/api/api-getbucketmetricsconfiguration.md)Grants permission to get a metrics configuration from an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetMultiRegionAccessPoint](../../../s3/latest/api/api-control-getmultiregionaccesspoint.md)Grants permission to return configuration information about the specified Multi-Region Access PointRead

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[GetMultiRegionAccessPointPolicy](../../../s3/latest/api/api-control-getmultiregionaccesspointpolicy.md)Grants permission to return the access point policy associated with the specified Multi-Region Access PointRead

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[GetMultiRegionAccessPointPolicyStatus](../../../s3/latest/api/api-control-getmultiregionaccesspointpolicystatus.md)Grants permission to return the policy status for a specific Multi-Region Access Point policyRead

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[GetMultiRegionAccessPointRoutes](../../../s3/latest/api/api-control-getmultiregionaccesspointroutes.md)Grants permission to return the route configuration for a Multi-Region Access PointRead

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[GetObject](../../../s3/latest/api/api-getobject.md)Grants permission to retrieve objects from Amazon S3Read

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectAcl](../../../s3/latest/api/api-getobjectacl.md)Grants permission to return the access control list (ACL) of an objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectAttributes](../../../s3/latest/api/api-getobjectattributes.md)Grants permission to retrieve attributes related to a specific objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectLegalHold](../../../s3/latest/api/api-getobjectlegalhold.md)Grants permission to get an object's current Legal Hold statusRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectRetention](../../../s3/latest/api/api-getobjectretention.md)Grants permission to retrieve the retention settings for an objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectTagging](../../../s3/latest/api/api-getobjecttagging.md)Grants permission to return the tag set of an objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectTorrent](../../../s3/latest/api/api-getobjecttorrent.md)Grants permission to return torrent files from an Amazon S3 bucketRead

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersion](../../../s3/latest/api/api-getobject.md)Grants permission to retrieve a specific version of an objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersionAcl](../../../s3/latest/api/api-getobjectacl.md)Grants permission to return the access control list (ACL) of a specific object versionRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersionAttributes](../../../s3/latest/api/api-getobjectattributes.md)Grants permission to retrieve attributes related to a specific version of an objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersionForReplication](../../../s3/latest/userguide/replication-config-for-kms-objects.md)Grants permission to replicate both unencrypted objects and objects encrypted with SSE-S3 or SSE-KMSRead

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersionTagging](../../../s3/latest/userguide/setting-repl-config-perm-overview.md)Grants permission to return the tag set for a specific version of the objectRead

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetObjectVersionTorrent](../../../s3/latest/api/api-getobjecttorrent.md)Grants permission to get Torrent files about a different version using the versionId subresourceRead

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetReplicationConfiguration](../../../s3/latest/api/api-getbucketreplication.md)Grants permission to get the replication configuration information set on an Amazon S3 bucketRead

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetStorageLensConfiguration](../../../s3/latest/api/api-control-getstoragelensconfiguration.md)Grants permission to get an Amazon S3 Storage Lens configurationRead

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetStorageLensConfigurationTagging](../../../s3/latest/api/api-control-getstoragelensconfigurationtagging.md)Grants permission to get the tag set of an existing Amazon S3 Storage Lens configurationRead

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetStorageLensDashboard](../../../s3/latest/userguide/storage-lens-dashboard.md)Grants permission to get an Amazon S3 Storage Lens dashboardRead

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[GetStorageLensGroup](../../../s3/latest/api/api-control-getstoragelensgroup.md)Grants permission to get an Amazon S3 Storage Lens groupRead

[storagelensgroup\*](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[InitiateReplication](../../../s3/latest/userguide/setting-repl-config-perm-overview.md) \[permission only\]Grants permission to initiate the replication process by setting replication status of an object to pendingWrite

[object\*](#amazons3-object)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[ListAccessGrants](../../../s3/latest/api/api-control-listaccessgrants.md)Grants permission to list Access GrantList

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[ListAccessGrantsInstances](../../../s3/latest/api/api-control-listaccessgrantsinstances.md)Grants permission to List Access Grants InstancesList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListAccessGrantsLocations](../../../s3/latest/api/api-control-listaccessgrantslocations.md)Grants permission to list Access Grants locationsList

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[ListAccessPoints](../../../s3/latest/api/api-control-listaccesspoints.md)Grants permission to list access pointsList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListAccessPointsForObjectLambda](../../../s3/latest/api/api-control-listaccesspointsforobjectlambda.md)Grants permission to list object lambda enabled accesspointsList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListAllMyBuckets](../../../s3/latest/api/api-listbuckets.md)Grants permission to list all buckets owned by the authenticated sender of the requestList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListBucket](../../../s3/latest/api/api-listobjectsv2.md)Grants permission to list some or all of the objects in an Amazon S3 bucket (up to 1000)List

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:delimiter](#amazons3-s3_delimiter)

[s3:max-keys](#amazons3-s3_max-keys)

[s3:prefix](#amazons3-s3_prefix)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListBucketMultipartUploads](../../../s3/latest/api/api-listmultipartuploads.md)Grants permission to list in-progress multipart uploadsList

[bucket\*](#amazons3-bucket)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[ListBucketVersions](../../../s3/latest/api/api-listobjectversions.md)Grants permission to list metadata about all the versions of objects in an Amazon S3 bucketList

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:delimiter](#amazons3-s3_delimiter)

[s3:max-keys](#amazons3-s3_max-keys)

[s3:prefix](#amazons3-s3_prefix)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListCallerAccessGrants](../../../s3/latest/api/api-control-listcalleraccessgrants.md)Grants permission to list caller's Access GrantList

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[ListJobs](../../../s3/latest/api/api-control-listjobs.md)Grants permission to list current jobs and jobs that have ended recentlyList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListMultiRegionAccessPoints](../../../s3/latest/api/api-control-listmultiregionaccesspoints.md)Grants permission to list Multi-Region Access PointsList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[ListMultipartUploadParts](../../../s3/latest/api/api-listparts.md)Grants permission to list the parts that have been uploaded for a specific multipart uploadList

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListStorageLensConfigurations](../../../s3/latest/api/api-control-liststoragelensconfigurations.md)Grants permission to list Amazon S3 Storage Lens configurationsList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListStorageLensGroups](../../../s3/latest/api/api-control-liststoragelensgroups.md)Grants permission to list S3 Storage Lens groupsList

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ListTagsForResource](../../../s3/latest/api/api-control-listtagsforresource.md)Grants permission to list the tags attached to the specified resourceList

[accessgrant](#amazons3-accessgrant)

[accessgrantsinstance](#amazons3-accessgrantsinstance)

[accessgrantslocation](#amazons3-accessgrantslocation)

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[storagelensgroup](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ObjectOwnerOverrideToBucketOwner](../../../s3/latest/userguide/replication-change-owner.md#repl-ownership-add-role-permission)Grants permission to change replica ownershipPermissions management

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PauseReplication](../../../fis/latest/userguide/fis-actions-reference.md#bucket-pause-replication) \[permission only\]Grants permission to pause S3 Replication from target source buckets to destination bucketsWrite

[bucket\*](#amazons3-bucket)

s3:GetReplicationConfiguration

s3:PutReplicationConfiguration

[s3:destinationRegion](#amazons3-s3_destinationRegion)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAccelerateConfiguration](../../../s3/latest/api/api-putbucketaccelerateconfiguration.md)Grants permission to use the accelerate subresource to set the Transfer Acceleration state of an existing S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAccessGrantsInstanceResourcePolicy](../../../s3/latest/api/api-control-putaccessgrantsinstanceresourcepolicy.md)Grants permission to put Access grants instance resource policyPermissions management

[accessgrantsinstance\*](#amazons3-accessgrantsinstance)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[PutAccessPointConfigurationForObjectLambda](../../../s3/latest/api/api-control-putaccesspointconfigurationforobjectlambda.md)Grants permission to set the configuration of the object lambda enabled access pointWrite

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAccessPointPolicy](../../../s3/latest/api/api-control-putaccesspointpolicy.md)Grants permission to associate an access policy with a specified access pointPermissions management

[accesspoint\*](#amazons3-accesspoint)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAccessPointPolicyForObjectLambda](../../../s3/latest/api/api-control-putaccesspointpolicyforobjectlambda.md)Grants permission to associate an access policy with a specified object lambda enabled access pointPermissions management

[objectlambdaaccesspoint\*](#amazons3-objectlambdaaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAccessPointPublicAccessBlock](../../../s3/latest/userguide/access-control-block-public-access.md#access-control-block-public-access-examples-access-point)Grants permission to associate public access block configurations with a specified access point, while creating a access pointPermissions management[PutAccountPublicAccessBlock](../../../s3/latest/api/api-control-putpublicaccessblock.md)Grants permission to create or modify the PublicAccessBlock configuration for an AWS accountPermissions management

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutAnalyticsConfiguration](../../../s3/latest/api/api-putbucketanalyticsconfiguration.md)Grants permission to set an analytics configuration for the bucket, specified by the analytics configuration IDWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketAbac](../../../s3/latest/api/api-putbucketabac.md)Grants permission to set ABAC configuration for a general purpose bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketAcl](../../../s3/latest/api/api-putbucketacl.md)Grants permission to set the permissions on an existing bucket using access control lists (ACLs)Permissions management

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[PutBucketCORS](../../../s3/latest/api/api-putbucketcors.md)Grants permission to set the CORS configuration for an Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketLogging](../../../s3/latest/api/api-putbucketlogging.md)Grants permission to set the logging parameters for an Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketNotification](../../../s3/latest/api/api-putbucketnotification.md)Grants permission to receive notifications when certain events happen in an Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketObjectLockConfiguration](../../../s3/latest/api/api-putobjectlockconfiguration.md)Grants permission to put Object Lock configuration on a specific bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:signatureversion](#amazons3-s3_signatureversion)

[PutBucketOwnershipControls](../../../s3/latest/api/api-putbucketownershipcontrols.md)Grants permission to add, replace or delete ownership controls on a bucketPermissions management

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketPolicy](../../../s3/latest/api/api-putbucketpolicy.md)Grants permission to add or replace a bucket policy on a bucketPermissions management

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketPublicAccessBlock](../../../s3/latest/api/api-putpublicaccessblock.md)Grants permission to create or modify the PublicAccessBlock configuration for a specific Amazon S3 bucketPermissions management

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketRequestPayment](../../../s3/latest/api/api-putbucketrequestpayment.md)Grants permission to set the request payment configuration of a bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketTagging](../../../s3/latest/api/api-putbuckettagging.md)Grants permission to add a set of tags to an existing Amazon S3 bucketTagging

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketVersioning](../../../s3/latest/api/api-putbucketversioning.md)Grants permission to set the versioning state of an existing Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutBucketWebsite](../../../s3/latest/api/api-putbucketwebsite.md)Grants permission to set the configuration of the website that is specified in the website subresourceWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutEncryptionConfiguration](../../../s3/latest/api/api-putbucketencryption.md)Grants permission to set the encryption configuration for an Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutIntelligentTieringConfiguration](../../../s3/latest/api/api-putbucketintelligenttieringconfiguration.md)Grants permission to create new or update or delete an existing Amazon S3 Intelligent Tiering configurationWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutInventoryConfiguration](../../../s3/latest/api/api-putbucketinventoryconfiguration.md)Grants permission to add an inventory configuration to the bucket, identified by the inventory IDWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:InventoryAccessibleOptionalFields](#amazons3-s3_InventoryAccessibleOptionalFields)

[PutJobTagging](../../../s3/latest/api/api-control-putjobtagging.md)Grants permission to replace tags on an existing Amazon S3 Batch Operations jobTagging

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:ExistingJobPriority](#amazons3-s3_ExistingJobPriority)

[s3:ExistingJobOperation](#amazons3-s3_ExistingJobOperation)

[aws:TagKeys](#amazons3-aws_TagKeys)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[PutLifecycleConfiguration](../../../s3/latest/api/api-putbucketlifecycleconfiguration.md)Grants permission to create a new lifecycle configuration for the bucket or replace an existing lifecycle configurationWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutMetricsConfiguration](../../../s3/latest/api/api-putbucketmetricsconfiguration.md)Grants permission to set or update a metrics configuration for the CloudWatch request metrics from an Amazon S3 bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutMultiRegionAccessPointPolicy](../../../s3/latest/api/api-control-putmultiregionaccesspointpolicy.md)Grants permission to associate an access policy with a specified Multi-Region Access PointPermissions management

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[PutObject](../../../s3/latest/api/api-putobject.md)Grants permission to add an object to a bucketWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:RequestObjectTag/<key>](#amazons3-s3_RequestObjectTag_key)

[s3:RequestObjectTagKeys](#amazons3-s3_RequestObjectTagKeys)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-copy-source](#amazons3-s3_x-amz-copy-source)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[s3:x-amz-metadata-directive](#amazons3-s3_x-amz-metadata-directive)

[s3:x-amz-server-side-encryption](#amazons3-s3_x-amz-server-side-encryption)

[s3:x-amz-server-side-encryption-aws-kms-key-id](#amazons3-s3_x-amz-server-side-encryption-aws-kms-key-id)

[s3:x-amz-server-side-encryption-customer-algorithm](#amazons3-s3_x-amz-server-side-encryption-customer-algorithm)

[s3:x-amz-storage-class](#amazons3-s3_x-amz-storage-class)

[s3:x-amz-website-redirect-location](#amazons3-s3_x-amz-website-redirect-location)

[s3:object-lock-mode](#amazons3-s3_object-lock-mode)

[s3:object-lock-retain-until-date](#amazons3-s3_object-lock-retain-until-date)

[s3:object-lock-remaining-retention-days](#amazons3-s3_object-lock-remaining-retention-days)

[s3:object-lock-legal-hold](#amazons3-s3_object-lock-legal-hold)

[s3:if-match](#amazons3-s3_if-match)

[s3:if-none-match](#amazons3-s3_if-none-match)

[s3:ObjectCreationOperation](#amazons3-s3_ObjectCreationOperation)

[PutObjectAcl](../../../s3/latest/api/api-putobjectacl.md)Grants permission to set the access control list (ACL) permissions for new or existing objects in an S3 bucketPermissions management

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[s3:x-amz-storage-class](#amazons3-s3_x-amz-storage-class)

[PutObjectLegalHold](../../../s3/latest/api/api-putobjectlegalhold.md)Grants permission to apply a Legal Hold configuration to the specified objectWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:object-lock-legal-hold](#amazons3-s3_object-lock-legal-hold)

[PutObjectRetention](../../../s3/latest/api/api-putobjectretention.md)Grants permission to place an Object Retention configuration on an objectWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:object-lock-mode](#amazons3-s3_object-lock-mode)

[s3:object-lock-retain-until-date](#amazons3-s3_object-lock-retain-until-date)

[s3:object-lock-remaining-retention-days](#amazons3-s3_object-lock-remaining-retention-days)

[PutObjectTagging](../../../s3/latest/api/api-putobjecttagging.md)Grants permission to set the supplied tag-set to an object that already exists in a bucketTagging

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:RequestObjectTag/<key>](#amazons3-s3_RequestObjectTag_key)

[s3:RequestObjectTagKeys](#amazons3-s3_RequestObjectTagKeys)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutObjectVersionAcl](../../../s3/latest/api/api-putobjectacl.md)Grants permission to use the acl subresource to set the access control list (ACL) permissions for an object that already exists in a bucketPermissions management

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:AccessGrantsInstanceArn](#amazons3-s3_AccessGrantsInstanceArn)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-acl](#amazons3-s3_x-amz-acl)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-grant-full-control](#amazons3-s3_x-amz-grant-full-control)

[s3:x-amz-grant-read](#amazons3-s3_x-amz-grant-read)

[s3:x-amz-grant-read-acp](#amazons3-s3_x-amz-grant-read-acp)

[s3:x-amz-grant-write](#amazons3-s3_x-amz-grant-write)

[s3:x-amz-grant-write-acp](#amazons3-s3_x-amz-grant-write-acp)

[s3:x-amz-storage-class](#amazons3-s3_x-amz-storage-class)

[PutObjectVersionTagging](../../../s3/latest/api/api-putobjecttagging.md)Grants permission to set the supplied tag-set for a specific version of an objectTagging

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:ExistingObjectTag/<key>](#amazons3-s3_ExistingObjectTag_key)

[s3:RequestObjectTag/<key>](#amazons3-s3_RequestObjectTag_key)

[s3:RequestObjectTagKeys](#amazons3-s3_RequestObjectTagKeys)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:versionid](#amazons3-s3_versionid)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[PutReplicationConfiguration](../../../s3/latest/api/api-putbucketreplication.md)Grants permission to create a new replication configuration or replace an existing oneWrite

[bucket\*](#amazons3-bucket)

iam:PassRole

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:isReplicationPauseRequest](#amazons3-s3_isReplicationPauseRequest)

[PutStorageLensConfiguration](../../../s3/latest/api/api-control-putstoragelensconfiguration.md)Grants permission to create or update an Amazon S3 Storage Lens configurationWrite

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:TagKeys](#amazons3-aws_TagKeys)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[PutStorageLensConfigurationTagging](../../../s3/latest/api/api-control-putstoragelensconfigurationtagging.md)Grants permission to put or replace tags on an existing Amazon S3 Storage Lens configurationTagging

[storagelensconfiguration\*](#amazons3-storagelensconfiguration)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:TagKeys](#amazons3-aws_TagKeys)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[ReplicateDelete](../../../s3/latest/userguide/setting-repl-config-perm-overview.md)Grants permission to replicate delete markers to the destination bucketWrite

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[ReplicateObject](../../../s3/latest/userguide/setting-repl-config-perm-overview.md)Grants permission to replicate objects and object tags to the destination bucketWrite

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-server-side-encryption](#amazons3-s3_x-amz-server-side-encryption)

[s3:x-amz-server-side-encryption-aws-kms-key-id](#amazons3-s3_x-amz-server-side-encryption-aws-kms-key-id)

[s3:x-amz-server-side-encryption-customer-algorithm](#amazons3-s3_x-amz-server-side-encryption-customer-algorithm)

[ReplicateTags](../../../s3/latest/userguide/setting-repl-config-perm-overview.md)Grants permission to replicate object tags to the destination bucketTagging

[object\*](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[RestoreObject](../../../s3/latest/api/api-restoreobject.md)Grants permission to restore an archived copy of an object back into Amazon S3Write

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[SubmitMultiRegionAccessPointRoutes](../../../s3/latest/api/api-control-submitmultiregionaccesspointroutes.md)Grants permission to submit a route configuration update for a Multi-Region Access PointWrite

[multiregionaccesspoint\*](#amazons3-multiregionaccesspoint)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[TagResource](../../../s3/latest/api/api-control-tagresource.md)Grants permission to add tags to the specified resourceTagging

[accessgrant](#amazons3-accessgrant)

[accessgrantsinstance](#amazons3-accessgrantsinstance)

[accessgrantslocation](#amazons3-accessgrantslocation)

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[storagelensgroup](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:TagKeys](#amazons3-aws_TagKeys)

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[UntagResource](../../../s3/latest/api/api-control-untagresource.md)Grants permission to remove tags from the specified resourceTagging

[accessgrant](#amazons3-accessgrant)

[accessgrantsinstance](#amazons3-accessgrantsinstance)

[accessgrantslocation](#amazons3-accessgrantslocation)

[accesspoint](#amazons3-accesspoint)

[bucket](#amazons3-bucket)

[storagelensgroup](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:TagKeys](#amazons3-aws_TagKeys)

[UpdateAccessGrantsLocation](../../../s3/latest/api/api-control-updateaccessgrantslocation.md)Grants permission to update Access Grants locationPermissions management

[accessgrantslocation\*](#amazons3-accessgrantslocation)

[s3:AccessGrantsLocationScope](#amazons3-s3_AccessGrantsLocationScope)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[UpdateBucketMetadataInventoryTableConfiguration](../../../s3/latest/api/api-updatebucketmetadatainventorytableconfiguration.md)Grants permission to update the inventory table configuration on an existing S3 Metadata configuration for a specified general purpose bucketWrite

[bucket\*](#amazons3-bucket)

kms:DescribeKey

s3tables:CreateNamespace

s3tables:CreateTable

s3tables:CreateTableBucket

s3tables:GetTable

s3tables:PutTableEncryption

s3tables:PutTablePolicy

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[UpdateBucketMetadataJournalTableConfiguration](../../../s3/latest/api/api-updatebucketmetadatajournaltableconfiguration.md)Grants permission to update the journal table configuration on an existing S3 Metadata configuration for a specified general purpose bucketWrite

[bucket\*](#amazons3-bucket)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[UpdateJobPriority](../../../s3/latest/api/api-control-updatejobpriority.md)Grants permission to update the priority of an existing jobWrite

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:RequestJobPriority](#amazons3-s3_RequestJobPriority)

[s3:ExistingJobPriority](#amazons3-s3_ExistingJobPriority)

[s3:ExistingJobOperation](#amazons3-s3_ExistingJobOperation)

[UpdateJobStatus](../../../s3/latest/api/api-control-updatejobstatus.md)Grants permission to update the status for the specified jobWrite

[job\*](#amazons3-job)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:ExistingJobPriority](#amazons3-s3_ExistingJobPriority)

[s3:ExistingJobOperation](#amazons3-s3_ExistingJobOperation)

[s3:JobSuspendedCause](#amazons3-s3_JobSuspendedCause)

[UpdateObjectEncryption](../../../s3/latest/api/api-updateobjectencryption.md)Grants permission to update the server-side encryption type of an existing object in a general purpose bucketWrite

[accesspointobject](#amazons3-accesspointobject)

[object](#amazons3-object)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

[s3:x-amz-server-side-encryption](#amazons3-s3_x-amz-server-side-encryption)

[s3:x-amz-server-side-encryption-aws-kms-key-id](#amazons3-s3_x-amz-server-side-encryption-aws-kms-key-id)

[UpdateStorageLensGroup](../../../s3/latest/api/api-control-updatestoragelensgroup.md)Grants permission to update an existing S3 Storage Lens groupWrite

[storagelensgroup\*](#amazons3-storagelensgroup)

[s3:authType](#amazons3-s3_authType)

[s3:ResourceAccount](#amazons3-s3_ResourceAccount)

[s3:signatureAge](#amazons3-s3_signatureAge)

[s3:signatureversion](#amazons3-s3_signatureversion)

[s3:TlsVersion](#amazons3-s3_TlsVersion)

[s3:x-amz-content-sha256](#amazons3-s3_x-amz-content-sha256)

## Resource types defined by Amazon S3

The following resource types are defined by this service and can be used in the `Resource` element of IAM permission policy statements. Each action in the [Actions table](#amazons3-actions-as-permissions) identifies the resource types that can be specified with that action. A resource type can also define which condition keys you can include in a policy. These keys are displayed in the last column of the table. For details about the columns in the following table, see [Resource types table](reference-policies-actions-resources-contextkeys.md#resources_table).

Resource typesARNCondition keys[accesspoint](../../../s3/latest/userguide/access-points.md)`arn:${Partition}:s3:${Region}:${Account}:accesspoint/${AccessPointName}`

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[accesspointobject](../../../s3/latest/userguide/access-points.md)`arn:${Partition}:s3:${Region}:${Account}:accesspoint/${AccessPointName}/object/${ObjectName}`

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[s3:AccessPointNetworkOrigin](#amazons3-s3_AccessPointNetworkOrigin)

[s3:AccessPointTag/${TagKey}](#amazons3-s3_AccessPointTag___TagKey_)

[s3:BucketTag/${TagKey}](#amazons3-s3_BucketTag___TagKey_)

[s3:DataAccessPointAccount](#amazons3-s3_DataAccessPointAccount)

[s3:DataAccessPointArn](#amazons3-s3_DataAccessPointArn)

[bucket](../../../s3/latest/userguide/usingbucket.md)`arn:${Partition}:s3:::${BucketName}`

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[s3:BucketTag/${TagKey}](#amazons3-s3_BucketTag___TagKey_)

[object](../../../s3/latest/userguide/usingobjects.md)`arn:${Partition}:s3:::${BucketName}/${ObjectName}`

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[s3:BucketTag/${TagKey}](#amazons3-s3_BucketTag___TagKey_)

[job](../../../s3/latest/userguide/batch-ops-managing-jobs.md)`arn:${Partition}:s3:${Region}:${Account}:job/${JobId}`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[storagelensconfiguration](../../../s3/latest/userguide/storage-lens.md)`arn:${Partition}:s3:${Region}:${Account}:storage-lens/${ConfigId}`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[storagelensgroup](../../../s3/latest/userguide/storage-lens-group.md)`arn:${Partition}:s3:${Region}:${Account}:storage-lens-group/${Name}`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[objectlambdaaccesspoint](../../../s3/latest/userguide/transforming-objects.md)`arn:${Partition}:s3-object-lambda:${Region}:${Account}:accesspoint/${AccessPointName}`[multiregionaccesspoint](../../../s3/latest/userguide/multiregionaccesspointrequests.md)`arn:${Partition}:s3::${Account}:accesspoint/${AccessPointAlias}`[multiregionaccesspointrequestarn](../../../s3/latest/userguide/multiregionaccesspointrequests.md)`arn:${Partition}:s3:us-west-2:${Account}:async-request/mrap/${Operation}/${Token}`[accessgrantsinstance](../../../s3/latest/userguide/access-grants-instance.md)`arn:${Partition}:s3:${Region}:${Account}:access-grants/default`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[accessgrantslocation](../../../s3/latest/userguide/access-grants-location.md)`arn:${Partition}:s3:${Region}:${Account}:access-grants/default/location/${Token}`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

[accessgrant](../../../s3/latest/userguide/access-grants-grant.md)`arn:${Partition}:s3:${Region}:${Account}:access-grants/default/grant/${Token}`

[aws:RequestTag/${TagKey}](#amazons3-aws_RequestTag___TagKey_)

[aws:ResourceTag/${TagKey}](#amazons3-aws_ResourceTag___TagKey_)

[aws:TagKeys](#amazons3-aws_TagKeys)

## Condition keys for Amazon S3

Amazon S3 defines the following condition keys that can be used in the `Condition` element of an IAM policy. You can use these keys to further refine the conditions under which the policy statement applies. For details about the columns in the following table, see [Condition keys table](reference-policies-actions-resources-contextkeys.md#context_keys_table).

To view the global condition keys that are available to all services, see [AWS global condition context keys](../../../iam/latest/userguide/reference-policies-condition-keys.md).

Condition keysDescriptionType[aws:RequestTag/${TagKey}](../../../iam/latest/userguide/reference-policies-condition-keys.md#condition-keys-requesttag)Filters access by the tags that are passed in the requestString[aws:ResourceTag/${TagKey}](../../../iam/latest/userguide/reference-policies-condition-keys.md#condition-keys-resourcetag)Filters access by the tags associated with the resourceString[aws:TagKeys](../../../iam/latest/userguide/reference-policies-condition-keys.md#condition-keys-tagkeys)Filters access by the tag keys that are passed in the requestArrayOfString[s3:AccessGrantScope](../../../s3/latest/userguide/access-grants-grant.md)Filters access by the grant scope of access grants grantString[s3:AccessGrantsInstanceArn](../../../s3/latest/userguide/access-grants-instance.md)Filters access by access grants instance ARNARN[s3:AccessGrantsLocationScope](../../../s3/latest/userguide/access-grants-location.md)Filters access by the location scope of access grants locationString[s3:AccessPointNetworkOrigin](../../../s3/latest/userguide/creating-access-points.md#access-points-policies)Filters access by the network origin (Internet or VPC)String[s3:AccessPointTag/${TagKey}](../../../s3/latest/userguide/access-points.md#tagging-and-policies)Filters access by existing access point tag key and valueString[s3:BucketTag/${TagKey}](../../../s3/latest/userguide/buckets-tagging.md)Filters access by the tags associated with the bucketString[s3:DataAccessPointAccount](../../../s3/latest/userguide/creating-access-points.md#access-points-policies)Filters access by the AWS Account ID that owns the access pointString[s3:DataAccessPointArn](../../../s3/latest/userguide/creating-access-points.md#access-points-policies)Filters access by an access point Amazon Resource Name (ARN)ARN[s3:ExistingJobOperation](../../../s3/latest/userguide/batch-ops-job-tags-examples.md)Filters access by operation to updating the job priorityString[s3:ExistingJobPriority](../../../s3/latest/userguide/batch-ops-job-tags-examples.md)Filters access by priority range to cancelling existing jobsNumeric[s3:ExistingObjectTag/<key>](../../../s3/latest/userguide/object-tagging.md#tagging-and-policies)Filters access by existing object tag key and valueString[s3:InventoryAccessibleOptionalFields](../../../s3/latest/userguide/example-bucket-policies.md#example-bucket-policies-s3-inventory-2)Filters access by restricting which optional metadata fields a user can add when configuring S3 Inventory reportsArrayOfString[s3:JobSuspendedCause](../../../s3/latest/userguide/batch-ops-job-tags-examples.md)Filters access by a specific job suspended cause (for example, AWAITING\_CONFIRMATION) to cancelling suspended jobsString[s3:ObjectCreationOperation](../../../s3/latest/userguide/conditional-writes-enforce.md)Filters access by whether or not the operation creates an objectBool[s3:RequestJobOperation](../../../s3/latest/userguide/batch-ops-job-tags-examples.md)Filters access by operation to creating jobsString[s3:RequestJobPriority](../../../s3/latest/userguide/batch-ops-job-tags-examples.md)Filters access by priority range to creating new jobsNumeric[s3:RequestObjectTag/<key>](../../../s3/latest/userguide/object-tagging.md#tagging-and-policies)Filters access by the tag keys and values to be added to objectsString[s3:RequestObjectTagKeys](../../../s3/latest/userguide/object-tagging.md#tagging-and-policies)Filters access by the tag keys to be added to objectsArrayOfString[s3:ResourceAccount](../../../s3/latest/userguide/amazon-s3-policy-keys.md#example-object-resource-account)Filters access by the resource owner AWS account IDString[s3:TlsVersion](../../../s3/latest/userguide/amazon-s3-policy-keys.md#example-object-tls-version)Filters access by the TLS version used by the clientNumeric[s3:authType](../../../s3/latest/api/bucket-policy-s3-sigv4-conditions.md)Filters access by authentication methodString[s3:delimiter](../../../s3/latest/userguide/walkthrough1.md)Filters access by delimiter parameterString[s3:destinationRegion](../../../s3/latest/userguide/replication.md)Filters access by a specific replication destination region for targeted buckets of the AWS FIS action aws:s3:bucket-pause-replicationString[s3:if-match](../../../s3/latest/userguide/conditional-writes-enforce.md)Filters access by the request's 'If-Match' conditional headerString[s3:if-none-match](../../../s3/latest/userguide/conditional-writes-enforce.md)Filters access by the request's 'If-None-Match' conditional headerString[s3:isReplicationPauseRequest](../../../fis/latest/userguide/security-iam-id-based-policy-examples.md#security-iam-policy-examples-s3)Filters access by request made via AWS FIS action aws:s3:bucket-pause-replicationBool[s3:locationconstraint](../../../s3/latest/userguide/amazon-s3-policy-keys.md#condition-key-bucket-ops-1)Filters access by a specific RegionString[s3:max-keys](../../../s3/latest/userguide/amazon-s3-policy-keys.md#example-numeric-condition-operators)Filters access by maximum number of keys returned in a ListBucket requestNumeric[s3:object-lock-legal-hold](../../../s3/latest/userguide/object-lock-overview.md#object-lock-legal-holds)Filters access by object legal hold statusString[s3:object-lock-mode](../../../s3/latest/userguide/object-lock-overview.md#object-lock-retention-modes)Filters access by object retention mode (COMPLIANCE or GOVERNANCE)String[s3:object-lock-remaining-retention-days](../../../s3/latest/userguide/object-lock-managing.md#object-lock-managing-retention-limits)Filters access by remaining object retention daysNumeric[s3:object-lock-retain-until-date](../../../s3/latest/userguide/object-lock-overview.md#object-lock-retention-periods)Filters access by object retain-until dateDate[s3:prefix](../../../s3/latest/userguide/amazon-s3-policy-keys.md#condition-key-bucket-ops-2)Filters access by key name prefixString[s3:signatureAge](../../../s3/latest/api/bucket-policy-s3-sigv4-conditions.md)Filters access by the age in milliseconds of the request signatureNumeric[s3:signatureversion](../../../s3/latest/api/bucket-policy-s3-sigv4-conditions.md)Filters access by the version of AWS Signature used on the requestString[s3:versionid](../../../s3/latest/userguide/amazon-s3-policy-keys.md#getobjectversion-limit-access-to-specific-version-3)Filters access by a specific object versionString[s3:x-amz-acl](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by canned ACL in the request's x-amz-acl headerString[s3:x-amz-bucket-namespace](../../../s3/latest/userguide/gpbucketnamespaces.md)Filters access by general purpose bucket namespace typeString[s3:x-amz-content-sha256](../../../s3/latest/api/bucket-policy-s3-sigv4-conditions.md)Filters access by unsigned content in your bucketString[s3:x-amz-copy-source](../../../s3/latest/userguide/amazon-s3-policy-keys.md#putobject-limit-copy-source-3)Filters access by copy source bucket, prefix, or object in the copy object requestsString[s3:x-amz-grant-full-control](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by x-amz-grant-full-control (full control) headerString[s3:x-amz-grant-read](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by x-amz-grant-read (read access) headerString[s3:x-amz-grant-read-acp](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by the x-amz-grant-read-acp (read permissions for the ACL) headerString[s3:x-amz-grant-write](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by the x-amz-grant-write (write access) headerString[s3:x-amz-grant-write-acp](../../../s3/latest/userguide/acl-overview.md#permissions)Filters access by the x-amz-grant-write-acp (write permissions for the ACL) headerString[s3:x-amz-metadata-directive](../../../s3/latest/api/api-copyobject.md)Filters access by object metadata behavior (COPY or REPLACE) when objects are copiedString[s3:x-amz-object-ownership](../../../s3/latest/userguide/ensure-object-ownership.md#object-ownership-requiring-bucket-owner-enforced)Filters access by Object OwnershipString[s3:x-amz-server-side-encryption](../../../s3/latest/userguide/usingserversideencryption.md)Filters access by server-side encryptionString[s3:x-amz-server-side-encryption-aws-kms-key-id](../../../s3/latest/userguide/usingkmsencryption.md#require-sse-kms)Filters access by AWS KMS customer managed CMK for server-side encryptionARN[s3:x-amz-server-side-encryption-customer-algorithm](../../../s3/latest/userguide/serversideencryptioncustomerkeys.md)Filters access by customer specified algorithm for server-side encryptionString[s3:x-amz-storage-class](../../../s3/latest/userguide/storage-class-intro.md#sc-howtoset)Filters access by storage classString[s3:x-amz-website-redirect-location](../../../s3/latest/userguide/how-to-page-redirect.md#page-redirect-using-rest-api)Filters access by a specific website redirect location for buckets that are configured as static websitesString

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Previous

Next

All content copied from https://docs.aws.amazon.com/.
