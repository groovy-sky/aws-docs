This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ClusterSecurityGroup

Specifies a new Amazon Redshift security group. You use security groups to control
access to non-VPC clusters.

For information about managing security groups, go to [Amazon Redshift\
Cluster Security Groups](../../../redshift/latest/mgmt/working-with-security-groups.md) in the _Amazon Redshift Cluster_
_Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::ClusterSecurityGroup",
  "Properties" : {
      "Description" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::ClusterSecurityGroup
Properties:
  Description: String
  Tags:
    - Tag

```

## Properties

`Description`

A description for the security group.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Specifies an arbitrary set of tags (key–value pairs) to associate with this security
group. Use tags to manage your resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-redshift-clustersecuritygroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myClusterSecurityGroup" }`

For the Amazon Redshift cluster security group `myClusterSecurityGroup`,
Ref returns the name of the cluster security group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

### Specify a Cluster Security Group

The following example describes an Amazon Redshift cluster security group
that you can associate cluster security group ingress rules with.

#### JSON

```json

"myClusterSecurityGroup": {
  "Type": "AWS::Redshift::ClusterSecurityGroup",
  "Properties": {
    "Description": "Security group to determine where connections to the Amazon Redshift cluster can come from",
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

myClusterSecurityGroup:
  Type: "AWS::Redshift::ClusterSecurityGroup"
  Properties:
    Description: "Security group to determine where connections to the Amazon Redshift cluster can come from"
    Tags:
      - Key: foo
        Value: bar
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
