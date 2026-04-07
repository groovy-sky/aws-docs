This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::Cluster

Specifies a cluster. A _cluster_ is a fully managed data warehouse
that consists of a set of compute nodes.

To create a cluster in Virtual Private Cloud (VPC), you must provide a cluster subnet
group name. The cluster subnet group identifies the subnets of your VPC that Amazon
Redshift uses when creating the cluster. For more information about managing clusters,
go to [Amazon Redshift Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in the _Amazon Redshift Cluster_
_Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::Cluster",
  "Properties" : {
      "AllowVersionUpgrade" : Boolean,
      "AquaConfigurationStatus" : String,
      "AutomatedSnapshotRetentionPeriod" : Integer,
      "AvailabilityZone" : String,
      "AvailabilityZoneRelocation" : Boolean,
      "AvailabilityZoneRelocationStatus" : String,
      "Classic" : Boolean,
      "ClusterIdentifier" : String,
      "ClusterParameterGroupName" : String,
      "ClusterSecurityGroups" : [ String, ... ],
      "ClusterSubnetGroupName" : String,
      "ClusterType" : String,
      "ClusterVersion" : String,
      "DBName" : String,
      "DeferMaintenance" : Boolean,
      "DeferMaintenanceDuration" : Integer,
      "DeferMaintenanceEndTime" : String,
      "DeferMaintenanceStartTime" : String,
      "DestinationRegion" : String,
      "ElasticIp" : String,
      "Encrypted" : Boolean,
      "Endpoint" : Endpoint,
      "EnhancedVpcRouting" : Boolean,
      "HsmClientCertificateIdentifier" : String,
      "HsmConfigurationIdentifier" : String,
      "IamRoles" : [ String, ... ],
      "KmsKeyId" : String,
      "LoggingProperties" : LoggingProperties,
      "MaintenanceTrackName" : String,
      "ManageMasterPassword" : Boolean,
      "ManualSnapshotRetentionPeriod" : Integer,
      "MasterPasswordSecretKmsKeyId" : String,
      "MasterUsername" : String,
      "MasterUserPassword" : String,
      "MultiAZ" : Boolean,
      "NamespaceResourcePolicy" : Json,
      "NodeType" : String,
      "NumberOfNodes" : Integer,
      "OwnerAccount" : String,
      "Port" : Integer,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "ResourceAction" : String,
      "RevisionTarget" : String,
      "RotateEncryptionKey" : Boolean,
      "SnapshotClusterIdentifier" : String,
      "SnapshotCopyGrantName" : String,
      "SnapshotCopyManual" : Boolean,
      "SnapshotCopyRetentionPeriod" : Integer,
      "SnapshotIdentifier" : String,
      "Tags" : [ Tag, ... ],
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::Cluster
Properties:
  AllowVersionUpgrade: Boolean
  AquaConfigurationStatus: String
  AutomatedSnapshotRetentionPeriod: Integer
  AvailabilityZone: String
  AvailabilityZoneRelocation: Boolean
  AvailabilityZoneRelocationStatus: String
  Classic: Boolean
  ClusterIdentifier: String
  ClusterParameterGroupName: String
  ClusterSecurityGroups:
    - String
  ClusterSubnetGroupName: String
  ClusterType: String
  ClusterVersion: String
  DBName: String
  DeferMaintenance: Boolean
  DeferMaintenanceDuration: Integer
  DeferMaintenanceEndTime: String
  DeferMaintenanceStartTime: String
  DestinationRegion: String
  ElasticIp: String
  Encrypted: Boolean
  Endpoint:
    Endpoint
  EnhancedVpcRouting: Boolean
  HsmClientCertificateIdentifier: String
  HsmConfigurationIdentifier: String
  IamRoles:
    - String
  KmsKeyId: String
  LoggingProperties:
    LoggingProperties
  MaintenanceTrackName: String
  ManageMasterPassword: Boolean
  ManualSnapshotRetentionPeriod: Integer
  MasterPasswordSecretKmsKeyId: String
  MasterUsername: String
  MasterUserPassword: String
  MultiAZ: Boolean
  NamespaceResourcePolicy: Json
  NodeType: String
  NumberOfNodes: Integer
  OwnerAccount: String
  Port: Integer
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  ResourceAction: String
  RevisionTarget: String
  RotateEncryptionKey: Boolean
  SnapshotClusterIdentifier: String
  SnapshotCopyGrantName: String
  SnapshotCopyManual: Boolean
  SnapshotCopyRetentionPeriod: Integer
  SnapshotIdentifier: String
  Tags:
    - Tag
  VpcSecurityGroupIds:
    - String

```

## Properties

`AllowVersionUpgrade`

If `true`, major version upgrades can be applied during the maintenance
window to the Amazon Redshift engine that is running on the cluster.

When a new major version of the Amazon Redshift engine is released, you can request that
the service automatically apply upgrades during the maintenance window to the Amazon Redshift
engine that is running on your cluster.

Default: `true`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AquaConfigurationStatus`

This parameter is retired. It does not set the AQUA configuration status. Amazon Redshift automatically determines whether to use AQUA (Advanced Query Accelerator).

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled | auto`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomatedSnapshotRetentionPeriod`

The number of days that automated snapshots are retained. If the value is 0, automated
snapshots are disabled. Even if automated snapshots are disabled, you can still create
manual snapshots when you want with [CreateClusterSnapshot](https://docs.aws.amazon.com/redshift/latest/APIReference/API_CreateClusterSnapshot.html) in the _Amazon Redshift API_
_Reference_.

Default: `1`

Constraints: Must be a value from 0 to 35.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the
cluster. For example, if you have several EC2 instances running in a specific
Availability Zone, then you might want the cluster to be provisioned in the same zone in
order to decrease network latency.

Default: A random, system-chosen Availability Zone in the region that is specified
by the endpoint.

Example: `us-east-2d`

Constraint: The specified Availability Zone must be in the same region as the
current endpoint.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZoneRelocation`

The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is created.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZoneRelocationStatus`

Describes the status of the Availability Zone relocation operation.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Classic`

A boolean value indicating whether the resize operation is using the classic resize
process. If you don't provide this parameter or set the value to
`false`, the resize type is elastic.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterIdentifier`

A unique identifier for the cluster. You use this identifier to refer to the
cluster for any subsequent cluster operations such as deleting or modifying. The
identifier also appears in the Amazon Redshift console.

Constraints:

- Must contain from 1 to 63 alphanumeric characters or hyphens.

- Alphabetic characters must be lowercase.

- First character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

- Must be unique for all clusters within an AWS account.

Example: `myexamplecluster`

_Required_: No

_Type_: String

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterParameterGroupName`

The name of the parameter group to be associated with this cluster.

Default: The default Amazon Redshift cluster parameter group. For information about the
default parameter group, go to [Working with Amazon\
Redshift Parameter Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)

Constraints:

- Must be 1 to 255 alphanumeric characters or hyphens.

- First character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterSecurityGroups`

A list of security groups to be associated with this cluster.

Default: The default cluster security group for Amazon Redshift.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterSubnetGroupName`

The name of a cluster subnet group to be associated with this cluster.

If this parameter is not provided the resulting cluster will be deployed outside
virtual private cloud (VPC).

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterType`

The type of the cluster. When cluster type is specified as

- `single-node`, the **NumberOfNodes**
parameter is not required.

- `multi-node`, the **NumberOfNodes**
parameter is required.

Valid Values: `multi-node` \| `single-node`

Default: `multi-node`

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterVersion`

The version of the Amazon Redshift engine software that you want to deploy on the
cluster.

The version selected runs on all the nodes in the cluster.

Constraints: Only version 1.0 is currently available.

Example: `1.0`

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBName`

The name of the first database to be created when the cluster is created.

To create additional databases after the cluster is created, connect to the cluster
with a SQL client and use SQL commands to create a database. For more information, go to
[Create\
a Database](https://docs.aws.amazon.com/redshift/latest/dg/t_creating_database.html) in the Amazon Redshift Database Developer Guide.

Default: `dev`

Constraints:

- Must contain 1 to 64 alphanumeric characters.

- Must contain only lowercase letters.

- Cannot be a word that is reserved by the service. A list of reserved words
can be found in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the
Amazon Redshift Database Developer Guide.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeferMaintenance`

A Boolean indicating whether to enable the deferred maintenance window.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeferMaintenanceDuration`

An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time.
The duration must be 60 days or less.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeferMaintenanceEndTime`

A timestamp for the end of the time period when we defer maintenance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeferMaintenanceStartTime`

A timestamp indicating the start time for the deferred maintenance window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationRegion`

The destination region that snapshots are automatically copied to when cross-region
snapshot copy is enabled.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElasticIp`

The Elastic IP (EIP) address for the cluster.

Constraints: The cluster must be provisioned in EC2-VPC and publicly-accessible
through an Internet gateway. Don't specify the Elastic IP address for a publicly accessible
cluster with availability zone relocation turned on. For more information about provisioning clusters in
EC2-VPC, go to [Supported\
Platforms to Launch Your Cluster](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms) in the Amazon Redshift Cluster Management Guide.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Encrypted`

If `true`, the data in the cluster is encrypted at rest.
If you set the value on this parameter to `false`, the request will fail.

Default: true

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The connection endpoint.

_Required_: No

_Type_: [Endpoint](aws-properties-redshift-cluster-endpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnhancedVpcRouting`

An option that specifies whether to create the cluster with enhanced VPC routing
enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a
VPC. For more information, see [Enhanced VPC Routing](https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in
the Amazon Redshift Cluster Management Guide.

If this option is `true`, enhanced VPC routing is enabled.

Default: false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HsmClientCertificateIdentifier`

Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to
retrieve the data encryption keys stored in an HSM.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HsmConfigurationIdentifier`

Specifies the name of the HSM configuration that contains the information the
Amazon Redshift cluster can use to retrieve and store keys in an HSM.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoles`

A list of AWS Identity and Access Management (IAM) roles that can be used by the
cluster to access other AWS services. You must supply the IAM roles in their Amazon
Resource Name (ARN) format.

The maximum number of IAM roles that you can associate is subject to a quota.
For more information, go to [Quotas and limits](https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html)
in the _Amazon Redshift Cluster Management Guide_.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The AWS Key Management Service (KMS) key ID of the encryption key that you want to
use to encrypt data in the cluster.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingProperties`

Specifies logging information, such as queries and connection attempts, for the
specified Amazon Redshift cluster.

_Required_: No

_Type_: [LoggingProperties](aws-properties-redshift-cluster-loggingproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceTrackName`

An optional parameter for the name of the maintenance track for the cluster. If you
don't provide a maintenance track name, the cluster is assigned to the
`current` track.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManageMasterPassword`

If `true`, Amazon Redshift uses AWS Secrets Manager to manage this cluster's admin credentials.
You can't use `MasterUserPassword` if `ManageMasterPassword` is true.
If `ManageMasterPassword` is false or not set, Amazon Redshift uses
`MasterUserPassword` for the admin user account's password.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManualSnapshotRetentionPeriod`

The default number of days to retain a manual snapshot. If the value is -1, the
snapshot is retained indefinitely. This setting doesn't change the retention period
of existing snapshots.

The value must be either -1 or an integer between 1 and 3,653.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterPasswordSecretKmsKeyId`

The ID of the AWS Key Management Service (KMS) key used to encrypt and store the cluster's admin credentials secret.
You can only use this parameter if `ManageMasterPassword` is true.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MasterUsername`

The user name associated with the admin user account for the cluster that is being
created.

Constraints:

- Must be 1 - 128 alphanumeric characters or hyphens. The user name can't be
`PUBLIC`.

- Must contain only lowercase letters, numbers, underscore, plus sign, period (dot), at symbol (@), or hyphen.

- The first character must be a letter.

- Must not contain a colon (:) or a slash (/).

- Cannot be a reserved word. A list of reserved words can be found in [Reserved\
Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MasterUserPassword`

The password associated with the admin user account for the cluster that is being
created.

You can't use `MasterUserPassword` if `ManageMasterPassword` is `true`.

Constraints:

- Must be between 8 and 64 characters in length.

- Must contain at least one uppercase letter.

- Must contain at least one lowercase letter.

- Must contain one number.

- Can be any printable ASCII character (ASCII code 33-126) except `'`
(single quote), `"` (double quote), `\`, `/`, or `@`.

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiAZ`

A boolean indicating whether Amazon Redshift should deploy the cluster in two Availability Zones. The default is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceResourcePolicy`

The policy that is attached to a resource.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeType`

The node type to be provisioned for the cluster. For information about node types,
go to [Working with\
Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the _Amazon Redshift Cluster Management Guide_.

Valid Values:
`dc2.large` \| `dc2.8xlarge` \|
`ra3.large` \| `ra3.xlplus` \| `ra3.4xlarge` \| `ra3.16xlarge`

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfNodes`

The number of compute nodes in the cluster. This parameter is required when the
**ClusterType** parameter is specified as
`multi-node`.

For information about determining how many nodes you need, go to [Working with\
Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the _Amazon Redshift Cluster Management Guide_.

If you don't specify this parameter, you get a single-node cluster. When requesting
a multi-node cluster, you must specify the number of nodes that you want in the
cluster.

Default: `1`

Constraints: Value must be at least 1 and no more than 100.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerAccount`

The AWS account used to create or copy the snapshot. Required if you are
restoring a snapshot you do not own, optional if you own the snapshot.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Port`

The port number on which the cluster accepts incoming connections.

The cluster is accessible only via the JDBC and ODBC connection strings. Part of
the connection string requires the port on which the cluster will listen for incoming
connections.

Default: `5439`

Valid Values:

- For clusters with ra3 nodes - Select a port within the ranges `5431-5455` or `8191-8215`. (If you have an existing cluster
with ra3 nodes, it isn't required that you change the port to these ranges.)

- For clusters with dc2 nodes - Select a port within the range `1150-65535`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range (in UTC) during which automated cluster maintenance can
occur.

Format: `ddd:hh24:mi-ddd:hh24:mi`

Default: A 30-minute window selected at random from an 8-hour block of time per
region, occurring on a random day of the week. For more information about the time
blocks for each region, see [Maintenance Windows](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows) in Amazon Redshift Cluster Management Guide.

Valid Days: Mon \| Tue \| Wed \| Thu \| Fri \| Sat \| Sun

Constraints: Minimum 30-minute window.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

If `true`, the cluster can be accessed from a public network.

Default: false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAction`

The Amazon Redshift operation to be performed. Supported operations are `pause-cluster`,
`resume-cluster`, and `failover-primary-compute`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevisionTarget`

Describes a `RevisionTarget` object.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotateEncryptionKey`

Rotates the encryption keys for a cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotClusterIdentifier`

The name of the cluster the source snapshot was created from. This parameter is
required if your user or role has a policy containing a snapshot resource element that
specifies anything other than \* for the cluster name.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotCopyGrantName`

The name of the snapshot copy grant.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotCopyManual`

Indicates whether to apply the snapshot retention period to newly copied manual
snapshots instead of automated snapshots.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotCopyRetentionPeriod`

The number of days to retain automated snapshots in the destination AWS Region
after they are copied from the source AWS Region.

By default, this only changes the retention period of copied automated snapshots.

If you decrease the retention period for automated snapshots that are copied to a
destination AWS Region, Amazon Redshift deletes any existing automated snapshots that were
copied to the destination AWS Region and that fall outside of the new retention
period.

Constraints: Must be at least 1 and no more than 35 for automated snapshots.

If you specify the `manual` option, only newly copied manual snapshots will
have the new retention period.

If you specify the value of -1 newly copied manual snapshots are retained
indefinitely.

Constraints: The number of days must be either -1 or an integer between 1 and 3,653
for manual snapshots.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotIdentifier`

The name of the snapshot from which to create the new cluster. This parameter isn't
case sensitive. You must specify this parameter or `snapshotArn`, but not both.

Example: `my-snapshot-id`

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tag instances.

_Required_: No

_Type_: Array of [Tag](aws-properties-redshift-cluster-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

A list of Virtual Private Cloud (VPC) security groups to be associated with the
cluster.

Default: The default VPC security group is associated with the cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myCluster" }`

For the Amazon Redshift cluster `myCluster`, `Ref` returns the
name of the cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterNamespaceArn`

The namespace Amazon Resource Name (ARN) of the cluster.

`DeferMaintenanceIdentifier`

A unique identifier for the maintenance window.

`Endpoint.Address`

The connection endpoint for the Amazon Redshift cluster. For example:
`examplecluster.cg034hpkmmjt.us-east-1.redshift.amazonaws.com`.

`Endpoint.Port`

The port number on which the Amazon Redshift cluster accepts connections. For
example: `5439`.

`MasterPasswordSecretArn`

The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.

## Examples

### Single-Node Cluster

The following example describes a single-node Redshift cluster. The master
user password is referenced from an input parameter that's in the same
template.

#### JSON

```json

{
"myCluster": {
  "Type": "AWS::Redshift::Cluster",
  "Properties": {
    "DBName": "mydb",
    "MasterUsername": "master",
    "MasterUserPassword": { "Ref" : "MasterUserPassword" },
    "NodeType": "ds2.xlarge",
    "ClusterType": "single-node",
    "Tags": [{
        "Key": "foo",
        "Value": "bar"
     }
    ]
  }
}
```

#### YAML

```yaml

myCluster:
  Type: "AWS::Redshift::Cluster"
  Properties:
    DBName: "mydb"
    MasterUsername: "master"
    MasterUserPassword:
      Ref: "MasterUserPassword"
    NodeType: "ds2.xlarge"
    ClusterType: "single-node"
    Tags:
      - Key: foo
        Value: bar
```

## See also

- For a complete example template, see [Amazon\
Redshift Template Snippets](../userguide/quickref-redshift.md) .

- [CreateCluster](https://docs.aws.amazon.com/redshift/latest/APIReference/API_CreateCluster.html) in the _Redshift API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Redshift

Endpoint
