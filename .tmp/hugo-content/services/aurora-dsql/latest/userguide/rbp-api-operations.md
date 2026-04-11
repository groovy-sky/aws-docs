# Aurora DSQL API Operations and Resource-Based Policies

Resource-based policies in Aurora DSQL control access to specific API operations. The following sections list all Aurora DSQL API operations organized by category, with an indication of which ones support resource-based policies.

The _Supports RBP_ column indicates whether the API operation is subject to resource-based policy evaluation when a policy is attached to the cluster.

## Tag APIs

API OperationDescriptionSupports RBP[ListTagsForResource](../../../../reference/aurora-dsql/latest/apireference/api-listtagsforresource.md)Lists the tags for a Aurora DSQL resourceYes[TagResource](../../../../reference/aurora-dsql/latest/apireference/api-tagresource.md)Adds tags to a Aurora DSQL resourceYes[UntagResource](../../../../reference/aurora-dsql/latest/apireference/api-untagresource.md)Removes tags from a Aurora DSQL resourceYes

## Cluster management APIs

API OperationDescriptionSupports RBP[CreateCluster](../../../../reference/aurora-dsql/latest/apireference/api-createcluster.md)Creates a new clusterNo[DeleteCluster](../../../../reference/aurora-dsql/latest/apireference/api-deletecluster.md)Deletes a clusterYes[GetCluster](../../../../reference/aurora-dsql/latest/apireference/api-getcluster.md)Retrieves information about a clusterYes[GetVpcEndpointServiceName](../../../../reference/aurora-dsql/latest/apireference/api-getvpcendpointservicename.md)Retrieves the VPC endpoint service name for a clusterYes[ListClusters](../../../../reference/aurora-dsql/latest/apireference/api-listclusters.md)Lists clusters in your accountNo[UpdateCluster](../../../../reference/aurora-dsql/latest/apireference/api-updatecluster.md)Updates the configuration of a clusterYes

## Multi-Region property APIs

API OperationDescriptionSupports RBP[AddPeerCluster](../../../../reference/aurora-dsql/latest/apireference/api-addpeercluster.md)Adds a peer cluster to a multi-region configurationYes[PutMultiRegionProperties](../../../../reference/aurora-dsql/latest/apireference/api-putmultiregionproperties.md)Sets multi-region properties for a clusterYes[PutWitnessRegion](../../../../reference/aurora-dsql/latest/apireference/api-putwitnessregion.md)Sets the witness region for a multi-region clusterYes

## Resource-based policy APIs

API OperationDescriptionSupports RBP[DeleteClusterPolicy](../../../../reference/aurora-dsql/latest/apireference/api-deleteclusterpolicy.md)Deletes the resource-based policy from a clusterYes[GetClusterPolicy](../../../../reference/aurora-dsql/latest/apireference/api-getclusterpolicy.md)Retrieves the resource-based policy for a clusterYes[PutClusterPolicy](../../../../reference/aurora-dsql/latest/apireference/api-putclusterpolicy.md)Creates or updates the resource-based policy for a clusterYes

## AWS Fault Injection Service APIs

API OperationDescriptionSupports RBP[InjectError](../../../../reference/aurora-dsql/latest/apireference/api-injecterror.md)Injects errors for fault injection testingNo

## Backup and restore APIs

API OperationDescriptionSupports RBP[GetBackupJob](../../../../reference/aurora-dsql/latest/apireference/api-getbackupjob.md)Retrieves information about a backup jobNo[GetRestoreJob](../../../../reference/aurora-dsql/latest/apireference/api-getrestorejob.md)Retrieves information about a restore jobNo[StartBackupJob](../../../../reference/aurora-dsql/latest/apireference/api-startbackupjob.md)Starts a backup job for a clusterYes[StartRestoreJob](../../../../reference/aurora-dsql/latest/apireference/api-startrestorejob.md)Starts a restore job from a backupNo[StopBackupJob](../../../../reference/aurora-dsql/latest/apireference/api-stopbackupjob.md)Stops a running backup jobNo[StopRestoreJob](../../../../reference/aurora-dsql/latest/apireference/api-stoprestorejob.md)Stops a running restore jobNo

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Block public access

Using a service-linked role

All content copied from https://docs.aws.amazon.com/.
