This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::AccessEntry

Creates an access entry.

An access entry allows an IAM principal to access your cluster. Access
entries can replace the need to maintain entries in the `aws-auth` `ConfigMap` for authentication. You have the following options for
authorizing an IAM principal to access Kubernetes objects on your cluster: Kubernetes
role-based access control (RBAC), Amazon EKS, or both. Kubernetes RBAC authorization requires you
to create and manage Kubernetes `Role`, `ClusterRole`,
`RoleBinding`, and `ClusterRoleBinding` objects, in addition
to managing access entries. If you use Amazon EKS authorization exclusively, you don't need
to create and manage Kubernetes `Role`, `ClusterRole`,
`RoleBinding`, and `ClusterRoleBinding` objects.

For more information about access entries, see [Access entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html) in the
_Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::AccessEntry",
  "Properties" : {
      "AccessPolicies" : [ AccessPolicy, ... ],
      "ClusterName" : String,
      "KubernetesGroups" : [ String, ... ],
      "PrincipalArn" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "Username" : String
    }
}

```

### YAML

```yaml

Type: AWS::EKS::AccessEntry
Properties:
  AccessPolicies:
    - AccessPolicy
  ClusterName: String
  KubernetesGroups:
    - String
  PrincipalArn: String
  Tags:
    - Tag
  Type: String
  Username: String

```

## Properties

`AccessPolicies`

The access policies to associate to the access entry.

_Required_: No

_Type_: Array of [AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-accessentry-accesspolicy.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of your cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KubernetesGroups`

The value for `name` that you've specified for `kind: Group` as
a `subject` in a Kubernetes `RoleBinding` or
`ClusterRoleBinding` object. Amazon EKS doesn't confirm that the value for
`name` exists in any bindings on your cluster. You can specify one or
more names.

Kubernetes authorizes the `principalArn` of the access entry to access any
cluster objects that you've specified in a Kubernetes `Role` or
`ClusterRole` object that is also specified in a binding's
`roleRef`. For more information about creating Kubernetes
`RoleBinding`, `ClusterRoleBinding`, `Role`, or
`ClusterRole` objects, see [Using RBAC\
Authorization in the Kubernetes documentation](https://kubernetes.io/docs/reference/access-authn-authz/rbac).

If you want Amazon EKS to authorize the `principalArn` (instead of, or in
addition to Kubernetes authorizing the `principalArn`), you can associate one or
more access policies to the access entry using `AssociateAccessPolicy`. If
you associate any access policies, the `principalARN` has all permissions
assigned in the associated access policies and all permissions in any Kubernetes
`Role` or `ClusterRole` objects that the group names are bound
to.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrincipalArn`

The ARN of the IAM principal for the `AccessEntry`. You can specify one ARN for each access entry. You can't specify the
same ARN in more than one access entry. This value can't be changed after access entry
creation.

The valid principals differ depending on the type of the access entry in the
`type` field. For `STANDARD` access entries, you can use every
IAM principal type. For nodes ( `EC2` (for EKS Auto Mode),
`EC2_LINUX`, `EC2_WINDOWS`, `FARGATE_LINUX`, and
`HYBRID_LINUX`), the only valid ARN is IAM roles.

You can't use the STS session principal type with access entries because this is a
temporary principal for each session and not a permanent identity that can be assigned
permissions.

[IAM best\
practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#bp-users-federation-idp) recommend using IAM roles with temporary credentials, rather
than IAM users with long-term credentials.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value. You define both. Tags don't
propagate to any other cluster or AWS resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-accessentry-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the new access entry. Valid values are `STANDARD`,
`FARGATE_LINUX`, `EC2_LINUX`, `EC2_WINDOWS`,
`EC2` (for EKS Auto Mode), `HYBRID_LINUX`, and
`HYPERPOD_LINUX`.

If the `principalArn` is for an IAM role that's used for self-managed
Amazon EC2 nodes, specify `EC2_LINUX` or `EC2_WINDOWS`. Amazon EKS grants
the necessary permissions to the node for you. If the `principalArn` is for
any other purpose, specify `STANDARD`. If you don't specify a value, Amazon EKS
sets the value to `STANDARD`. If you have the access mode of the cluster set
to `API_AND_CONFIG_MAP`, it's unnecessary to create access entries for IAM
roles used with Fargate profiles or managed Amazon EC2 nodes, because Amazon EKS creates entries
in the `aws-auth` `ConfigMap` for the roles. You can't change this value once you've created
the access entry.

If you set the value to `EC2_LINUX` or `EC2_WINDOWS`, you can't
specify values for `kubernetesGroups`, or associate an
`AccessPolicy` to the access entry.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Username`

The username to authenticate to Kubernetes with. We recommend not specifying a username and
letting Amazon EKS specify it for you. For more information about the value Amazon EKS specifies
for you, or constraints before specifying your own username, see [Creating\
access entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html#creating-access-entries) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "arn:aws:iam::012345678910:role/my-role" }`

For the access entry `arn:aws:iam::012345678910:role/my-role`,
`Ref` returns the ARN of the access entry.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccessEntryArn`

The ARN of the access entry.

## Examples

### Create an access entry

The following example creates an access entry of type `STANDARD`
for an IAM role named `my-role`. The role has
administrator permissions for all resources in the namespace named
`my-namespace` and edit permissions for all resources on the
cluster.

#### JSON

```json

{
    "myAccessEntry": {
        "Type" : "AWS::EKS::AccessEntry",
        "Properties" : {
            "ClusterName": "my-cluster",
            "PrincipalArn": "arn:aws:iam::012345678910:role/my-role",
            "Username": "admin",
            "KubernetesGroups": ["my-group1", "my-group2"],
            "AccessPolicies": [
                {
                    "PolicyArn": "arn:aws:eks::aws:cluster-access-policy/AmazonEKSAdminPolicy",
                    "AccessScope": {
                        "Type": "namespace",
                        "Namespaces": ["my-namespace"]
                    }
                },
                {
                    "PolicyArn": "arn:aws:eks::aws:cluster-access-policy/AmazonEKSEditPolicy",
                    "AccessScope": {
                        "Type": "cluster"
                    }
                }
            ],
            "Type": "STANDARD",
            "Tags" : [
                {
                    "Key": "my-tagkey",
                    "Value": "my-tagvalue"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

MyAccessEntry:
    Type: 'AWS::EKS::AccessEntry'
    Properties:
        ClusterName: 'my-cluster'
        PrincipalArn: 'arn:aws:iam::012345678910:role/my-role'
        Username: 'admin'
        KubernetesGroups:
            - 'my-group1'
            - 'my-group2'
        AccessPolicies:
            - PolicyArn: 'arn:aws:eks::aws:cluster-access-policy/AmazonEKSAdminPolicy'
              AccessScope:
                Type: 'namespace'
                Namespaces:
                    - 'my-namespace'
            - PolicyArn: 'arn:aws:eks::aws:cluster-access-policy/AmazonEKSEditPolicy'
              AccessScope:
                Type: 'cluster'
        Type: 'STANDARD'
        Tags:
            - Key: 'my-tagkey'
              Value: 'my-tagvalue'
```

## See also

- [Access entries](https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html) in the _Amazon EKS User Guide_.

- [`CreateAccessEntry`](https://docs.aws.amazon.com/eks/latest/APIReference/API_CreateAccessEntry.html) in the _Amazon EKS API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Elastic Kubernetes Service

AccessPolicy
