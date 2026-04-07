This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDBElastic::Cluster

Creates a new Amazon DocumentDB elastic cluster and returns its cluster structure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDBElastic::Cluster",
  "Properties" : {
      "AdminUserName" : String,
      "AdminUserPassword" : String,
      "AuthType" : String,
      "BackupRetentionPeriod" : Integer,
      "ClusterName" : String,
      "KmsKeyId" : String,
      "PreferredBackupWindow" : String,
      "PreferredMaintenanceWindow" : String,
      "ShardCapacity" : Integer,
      "ShardCount" : Integer,
      "ShardInstanceCount" : Integer,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcSecurityGroupIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DocDBElastic::Cluster
Properties:
  AdminUserName: String
  AdminUserPassword: String
  AuthType: String
  BackupRetentionPeriod: Integer
  ClusterName: String
  KmsKeyId: String
  PreferredBackupWindow: String
  PreferredMaintenanceWindow: String
  ShardCapacity: Integer
  ShardCount: Integer
  ShardInstanceCount: Integer
  SubnetIds:
    - String
  Tags:
    - Tag
  VpcSecurityGroupIds:
    - String

```

## Properties

`AdminUserName`

The name of the Amazon DocumentDB elastic clusters administrator.

_Constraints_:

- Must be from 1 to 63 letters or numbers.

- The first character must be a letter.

- Cannot be a reserved word.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AdminUserPassword`

The password for the Elastic DocumentDB cluster administrator and can
contain any printable ASCII characters.

_Constraints_:

- Must contain from 8 to 100 characters.

- Cannot contain a forward slash (/), double quote ("), or the "at" symbol (@).

- A valid `AdminUserName` entry is also required.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthType`

The authentication type used to determine where to fetch the password used for accessing the elastic cluster.
Valid types are `PLAIN_TEXT` or `SECRET_ARN`.

_Required_: Yes

_Type_: String

_Allowed values_: `PLAIN_TEXT | SECRET_ARN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupRetentionPeriod`

The number of days for which automatic snapshots are retained.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of the new elastic cluster. This parameter is stored as
a lowercase string.

_Constraints_:

- Must contain from 1 to 63 letters, numbers, or hyphens.

- The first character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

_Example_: `my-cluster`

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-z][a-zA-Z0-9]*(-[a-zA-Z0-9]+)*`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The KMS key identifier to use to encrypt the new elastic cluster.

The KMS key identifier is the Amazon Resource Name (ARN) for the KMS
encryption key. If you are creating a cluster using the same Amazon account
that owns this KMS encryption key, you can use the KMS key alias instead
of the ARN as the KMS encryption key.

If an encryption key is not specified, Amazon DocumentDB uses the
default encryption key that KMS creates for your account. Your account
has a different default encryption key for each Amazon Region.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredBackupWindow`

The daily time range during which automated backups are created if automated backups are enabled, as determined by `backupRetentionPeriod`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The weekly time range during which system maintenance can occur,
in Universal Coordinated Time (UTC).

_Format_: `ddd:hh24:mi-ddd:hh24:mi`

_Default_: a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week.

_Valid days_: Mon, Tue, Wed, Thu, Fri, Sat, Sun

_Constraints_: Minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShardCapacity`

The number of vCPUs assigned to each elastic cluster shard. Maximum is 64. Allowed values are 2, 4, 8, 16, 32, 64.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShardCount`

The number of shards assigned to the elastic cluster. Maximum is 32.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShardInstanceCount`

The number of replica instances applying to all shards in the cluster.
A `shardInstanceCount` value of 1 means there is one writer instance, and any additional instances are replicas that can be used for reads and to improve availability.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The Amazon EC2 subnet IDs for the new elastic cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the new elastic cluster.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-docdbelastic-cluster-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSecurityGroupIds`

A list of EC2 VPC security groups to associate with the new
elastic cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ClusterArn`

Property description not available.

`ClusterEndpoint`

The URL used to connect to the elastic cluster.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon DocumentDB (with MongoDB compatibility) elastic

Tag
