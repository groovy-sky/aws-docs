This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DAX::Cluster

Creates a DAX cluster. All nodes in the cluster run the same DAX caching software.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DAX::Cluster",
  "Properties" : {
      "AvailabilityZones" : [ String, ... ],
      "ClusterEndpointEncryptionType" : String,
      "ClusterName" : String,
      "Description" : String,
      "IAMRoleARN" : String,
      "NetworkType" : String,
      "NodeType" : String,
      "NotificationTopicARN" : String,
      "ParameterGroupName" : String,
      "PreferredMaintenanceWindow" : String,
      "ReplicationFactor" : Integer,
      "SecurityGroupIds" : [ String, ... ],
      "SSESpecification" : SSESpecification,
      "SubnetGroupName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DAX::Cluster
Properties:
  AvailabilityZones:
    - String
  ClusterEndpointEncryptionType: String
  ClusterName: String
  Description: String
  IAMRoleARN: String
  NetworkType: String
  NodeType: String
  NotificationTopicARN: String
  ParameterGroupName: String
  PreferredMaintenanceWindow: String
  ReplicationFactor: Integer
  SecurityGroupIds:
    - String
  SSESpecification:
    SSESpecification
  SubnetGroupName: String
  Tags:
    - Tag

```

## Properties

`AvailabilityZones`

The Availability Zones (AZs) in which the cluster nodes will reside after the
cluster has been created or updated. If provided, the length of this list must equal the
`ReplicationFactor` parameter. If you omit this parameter, DAX will spread the nodes across Availability Zones for the highest
availability.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterEndpointEncryptionType`

The encryption type of the cluster's endpoint. Available values are:

- `NONE` \- The cluster's endpoint will be unencrypted.

- `TLS` \- The cluster's endpoint will be encrypted with Transport
Layer Security, and will provide an x509 certificate for
authentication.

The default value is `NONE`.

_Required_: No

_Type_: String

_Allowed values_: `NONE | TLS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

The name of the DAX cluster.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`Description`

The description of the cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMRoleARN`

A valid Amazon Resource Name (ARN) that identifies an IAM role. At runtime, DAX
will assume this role and use the role's permissions to access DynamoDB on your
behalf.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`NetworkType`

The IP address type of the cluster. Values are:

- `ipv4` \- IPv4 addresses only

- `ipv6` \- IPv6 addresses only

- `dual_stack` \- Both IPv4 and IPv6 addresses

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | ipv6 | dual_stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeType`

The node type for the nodes in the cluster. (All nodes in a DAX cluster are of
the same type.)

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationTopicARN`

The Amazon Resource Name (ARN) of the Amazon SNS topic to which
notifications will be sent.

###### Note

The Amazon SNS topic owner must be same as the DAX
cluster owner.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterGroupName`

The parameter group to be associated with the DAX cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

A range of time when maintenance of DAX cluster software will be performed. For
example: `sun:01:00-sun:09:00`. Cluster maintenance normally takes less than
30 minutes, and is performed automatically within the maintenance window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationFactor`

The number of nodes in the DAX cluster. A replication factor of 1
will create a single-node cluster, without any read replicas. For additional fault
tolerance, you can create a multiple node cluster with one or more read replicas. To do
this, set `ReplicationFactor` to a number between 3 (one primary and two read
replicas) and 10 (one primary and nine read replicas). `If the
                AvailabilityZones` parameter is provided, its length must equal the
`ReplicationFactor`.

###### Note

AWS recommends that you have at least two read replicas per
cluster.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A list of security group IDs to be assigned to each node in the DAX
cluster. (Each of the security group ID is system-generated.)

If this parameter is not specified, DAX assigns the default VPC
security group to each node.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSESpecification`

Represents the settings used to enable server-side encryption on the
cluster.

_Required_: No

_Type_: [SSESpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dax-cluster-ssespecification.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetGroupName`

The name of the subnet group to be used for the replication group.

###### Important

DAX clusters can only run in an Amazon VPC environment. All of the subnets
that you specify in a subnet group must exist in the same VPC.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`Tags`

A set of tags to associate with the DAX cluster.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the created DAX cluster. For example:

```nohighlight

{ "Ref": "MyResource" }
```

Returns a value similar to the following:

```

MyDAXCluster
```

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`Arn`

Returns the ARN of the DAX cluster. For example:

```nohighlight

{ "Fn::GetAtt": ["MyDAXCluster", "Arn"] }
```

Returns a value similar to the following:

```nohighlight

arn:aws:dax:us-east-1:111122223333:cache/MyDAXCluster
```

`ClusterDiscoveryEndpoint`

Returns the endpoint of the DAX cluster. For example:

```nohighlight

{ "Fn::GetAtt": ["MyDAXCluster", "ClusterDiscoveryEndpoint"] }
```

Returns a value similar to the following:

```nohighlight

mydaxcluster.0h3d6x.clustercfg.dax.use1.cache.amazonaws.com:8111
```

`ClusterDiscoveryEndpointURL`

Returns the endpoint URL of the DAX cluster.

## Examples

### Create Cluster

The following example creates a DAX cluster.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Create a DAX cluster",
  "Resources": {
    "daxCluster": {
      "Type": "AWS::DAX::Cluster",
      "Properties": {
        "ClusterName": "MyDAXCluster",
        "NodeType": "dax.r3.large",
        "ReplicationFactor": 1,
        "IAMRoleARN": "arn:aws:iam::111122223333:role/DaxAccess",
        "Description": "DAX cluster created with CloudFormation",
        "SubnetGroupName": {"Ref":"subnetGroupClu"}
      }
    },
    "subnetGroupClu": {
      "Type": "AWS::DAX::SubnetGroup",
      "Properties": {
        "SubnetGroupName": "MySubnetGroup",
        "Description": "Subnet group for DAX cluster",
        "SubnetIds": [
          {"Ref":"subnet1"},
          {"Ref":"subnet2"}
        ]
      }
    },
    "subnet1": {
      "Type": "AWS::EC2::Subnet",
      "Properties": {
        "VpcId": {"Ref":"daxVpc"},
        "CidrBlock": "172.13.17.0/24",
        "AvailabilityZone": {
          "Fn::Select": [
            0,
            {
              "Fn::GetAZs": ""
            }
          ]
        }
      }
    },
    "subnet2": {
      "Type": "AWS::EC2::Subnet",
      "Properties": {
        "VpcId": {"Ref":"daxVpc"},
        "CidrBlock": "172.13.18.0/24",
        "AvailabilityZone": {
          "Fn::Select": [
            1,
            {
              "Fn::GetAZs": ""
            }
          ]
        }
      }
    },
    "daxVpc": {
      "Type": "AWS::EC2::VPC",
      "Properties": {
        "CidrBlock": "172.13.0.0/16"
      }
    }
  },
  "Outputs": {
    "Cluster": {
      "Value": {"Ref":"daxCluster"}
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Description: "Create a DAX cluster"
Resources:
  daxCluster:
    Type: AWS::DAX::Cluster
    Properties:
      ClusterName: "MyDAXCluster"
      NodeType: "dax.r3.large"
      ReplicationFactor: 1
      IAMRoleARN: "arn:aws:iam::111122223333:role/DaxAccess"
      Description: "DAX cluster created with CloudFormation"
      SubnetGroupName: !Ref subnetGroupClu
  subnetGroupClu:
    Type: AWS::DAX::SubnetGroup
    Properties:
      SubnetGroupName: "CFNClusterSubnetGrp"
      Description: "Subnet group for DAX cluster"
      SubnetIds:
        - !Ref subnet1
        - !Ref subnet2
  subnet1:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId:
        !Ref daxVpc
      CidrBlock: 172.13.17.0/24
      AvailabilityZone:
        Fn::Select:
          - 0
          - Fn::GetAZs: ""
  subnet2:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId:
        !Ref daxVpc
      CidrBlock: 172.13.18.0/24
      AvailabilityZone:
        Fn::Select:
          - 1
          - Fn::GetAZs: ""
  daxVpc:
     Type: AWS::EC2::VPC
     Properties:
       CidrBlock: 172.13.0.0/16
Outputs:
  Cluster:
    Value: !Ref daxCluster
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DynamoDB Accelerator

SSESpecification
