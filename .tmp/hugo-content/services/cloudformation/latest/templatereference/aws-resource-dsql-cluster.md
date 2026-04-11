This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DSQL::Cluster

The `AWS::DSQL::Cluster` resource specifies an cluster.
You can use this resource to create, modify, and manage clusters.

This resource supports both single-Region clusters and multi-Region clusters through
the `MultiRegionProperties` parameter.

###### Note

Creating multi-Region clusters requires additional IAM permissions beyond those needed for single-Region clusters.

###### Important

- The witness Region specified in `multiRegionProperties.witnessRegion` cannot be the same as the cluster's Region.

**Required permissions**

dsql:CreateCluster

Required to create a cluster.

Resources: `arn:aws:dsql:region:account-id:cluster/*`

dsql:TagResource

Permission to add tags to a resource.

Resources: `arn:aws:dsql:region:account-id:cluster/*`

dsql:PutMultiRegionProperties

Permission to configure multi-Region properties for a cluster.

Resources: `arn:aws:dsql:region:account-id:cluster/*`

dsql:AddPeerCluster

When specifying `multiRegionProperties.clusters`, permission to add peer clusters.

Resources:

- Local cluster: `arn:aws:dsql:region:account-id:cluster/*`

- Each peer cluster: exact ARN of each specified peer cluster

dsql:PutWitnessRegion

When specifying `multiRegionProperties.witnessRegion`, permission to set a witness Region. This permission is checked both in the cluster Region and in the witness Region.

Resources: `arn:aws:dsql:region:account-id:cluster/*`

Condition Keys: `dsql:WitnessRegion` (matching the specified witness region)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DSQL::Cluster",
  "Properties" : {
      "DeletionProtectionEnabled" : Boolean,
      "KmsEncryptionKey" : String,
      "MultiRegionProperties" : MultiRegionProperties,
      "PolicyDocument" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DSQL::Cluster
Properties:
  DeletionProtectionEnabled: Boolean
  KmsEncryptionKey: String
  MultiRegionProperties:
    MultiRegionProperties
  PolicyDocument: String
  Tags:
    - Tag

```

## Properties

`DeletionProtectionEnabled`

Whether deletion protection is enabled on this cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsEncryptionKey`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiRegionProperties`

Defines the structure for multi-Region cluster configurations, containing the witness
Region and peered cluster settings.

_Required_: No

_Type_: [MultiRegionProperties](aws-properties-dsql-cluster-multiregionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDocument`

A resource-based policy document in JSON format. Length constraints: Minimum length of 1. Maximum length of 20480 characters (approximately 20KB).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A map of key and value pairs this cluster is tagged with.

_Required_: No

_Type_: Array of [Tag](aws-properties-dsql-cluster-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the cluster identifier.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The timestamp when the cluster was created, in ISO 8601 format.

`Endpoint`

The connection endpoint for the created cluster.

`Identifier`

The unique identifier assigned to the cluster upon creation.

`PolicyVersion`

Property description not available.

`ResourceArn`

The Amazon Resource Name (ARN) of the cluster. Used for IAM permissions and resource identification.

`Status`

The current status of the cluster. Possible values include: CREATING, ACTIVE, DELETING, FAILED.

The cluster can have two additional status values when working with multi-Region
clusters:

`PENDING_SETUP`—Indicates the cluster is being configured

`PENDING_DELETE`—Indicates the cluster is being deleted

**Note:** These status values only appear for multi-Region cluster operations.

`VpcEndpoint`

Property description not available.

`VpcEndpointServiceName`

The VPC Endpoint Service name for the cluster. This can be used to create a VPC endpoint to connect to the cluster from within a VPC.

## See also

See [API reference](../../../aurora-dsql/latest/userguide/chap-api-reference.md) for a full list of API operations to manage your resources in Aurora DSQL.

See [MultiRegionProperties](../../../../reference/aurora-dsql/latest/apireference/api-multiregionproperties.md) for the data structure used for multi-Region
clusters.

See [Configuring\
multi-Region clusters using CloudFormation](../../../aurora-dsql/latest/userguide/mr-cluster-setup.md) to create multi-Region
clusters using this CloudFormation resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Aurora DSQL

EncryptionDetails

All content copied from https://docs.aws.amazon.com/.
