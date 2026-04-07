This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ClusterSubnetGroup

Specifies an Amazon Redshift subnet group. You must provide a list of one or more
subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when creating
Amazon Redshift subnet group.

For information about subnet groups, go to [Amazon Redshift\
Cluster Subnet Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html) in the _Amazon Redshift Cluster Management_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::ClusterSubnetGroup",
  "Properties" : {
      "Description" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::ClusterSubnetGroup
Properties:
  Description: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`Description`

A description for the subnet group.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
request.

_Required_: Yes

_Type_: Array of String

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies an arbitrary set of tags (key–value pairs) to associate with this subnet
group. Use tags to manage your resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-redshift-clustersubnetgroup-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myClusterSubnetGroup" }`

For the Amazon Redshift subnet group `myClusterSubnetGroup`, Ref returns
the name of the cluster subnet group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClusterSubnetGroupName`

The name of the cluster subnet group.

## Examples

### Specify a Subnet

The following example specifies one subnet for an Amazon Redshift cluster
subnet group.

#### JSON

```json

"myClusterSubnetGroup": {
  "Type": "AWS::Redshift::ClusterSubnetGroup",
  "Properties": {
    "Description": "My ClusterSubnetGroup",
    "SubnetIds": [
      "subnet-7fbc2813"
    ],
    "Tags": [
      {
        "Key": "foo",
        "Value": "bar"
      }
    ]
  }
}
```

#### YAML

```yaml

myClusterSubnetGroup:
  Type: 'AWS::Redshift::ClusterSubnetGroup'
  Properties:
    Description: My ClusterSubnetGroup
    SubnetIds:
      - subnet-7fbc2813
    Tags:
      - Key: foo
        Value: bar
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Redshift::ClusterSecurityGroupIngress

Tag
