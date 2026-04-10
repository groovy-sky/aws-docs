This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::FargateProfile

Creates an AWS Fargate profile for your Amazon EKS cluster. You must have at least one
Fargate profile in a cluster to be able to run pods on Fargate.

The Fargate profile allows an administrator to declare which pods run on Fargate
and specify which pods run on which Fargate profile. This declaration is done through
the profile's selectors. Each profile can have up to five selectors that contain a
namespace and labels. A namespace is required for every selector. The label field
consists of multiple optional key-value pairs. Pods that match the selectors are
scheduled on Fargate. If a to-be-scheduled pod matches any of the selectors in the
Fargate profile, then that pod is run on Fargate.

When you create a Fargate profile, you must specify a pod execution role to use with
the pods that are scheduled with the profile. This role is added to the cluster's Kubernetes
[Role\
Based Access Control](https://kubernetes.io/docs/reference/access-authn-authz/rbac) (RBAC) for authorization so that the
`kubelet` that is running on the Fargate infrastructure can register
with your Amazon EKS cluster so that it can appear in your cluster as a node. The pod
execution role also provides IAM permissions to the Fargate infrastructure to allow
read access to Amazon ECR image repositories. For more information, see [Pod\
Execution Role](../../../eks/latest/userguide/pod-execution-role.md) in the _Amazon EKS User Guide_.

Fargate profiles are immutable. However, you can create a new updated profile to
replace an existing profile and then delete the original after the updated profile has
finished creating.

If any Fargate profiles in a cluster are in the `DELETING` status, you
must wait for that Fargate profile to finish deleting before you can create any other
profiles in that cluster.

For more information, see [AWS Fargate profile](../../../eks/latest/userguide/fargate-profile.md) in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::FargateProfile",
  "Properties" : {
      "ClusterName" : String,
      "FargateProfileName" : String,
      "PodExecutionRoleArn" : String,
      "Selectors" : [ Selector, ... ],
      "Subnets" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EKS::FargateProfile
Properties:
  ClusterName: String
  FargateProfileName: String
  PodExecutionRoleArn: String
  Selectors:
    - Selector
  Subnets:
    - String
  Tags:
    - Tag

```

## Properties

`ClusterName`

The name of your cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FargateProfileName`

The name of the Fargate profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PodExecutionRoleArn`

The Amazon Resource Name (ARN) of the `Pod` execution role to use for a `Pod`
that matches the selectors in the Fargate profile. The `Pod` execution role
allows Fargate infrastructure to register with your cluster as a node, and it provides
read access to Amazon ECR image repositories. For more information, see [`Pod` execution role](../../../eks/latest/userguide/pod-execution-role.md) in the _Amazon EKS User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Selectors`

The selectors to match for a `Pod` to use this Fargate profile. Each
selector must have an associated Kubernetes `namespace`. Optionally, you can also
specify `labels` for a `namespace`. You may specify up to five
selectors in a Fargate profile.

_Required_: Yes

_Type_: Array of [Selector](aws-properties-eks-fargateprofile-selector.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Subnets`

The IDs of subnets to launch a `Pod` into. A `Pod` running on
Fargate isn't assigned a public IP address, so only private subnets (with no direct
route to an Internet Gateway) are accepted for this parameter.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value. You define both. Tags don't
propagate to any other cluster or AWS resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-eks-fargateprofile-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myFargateProfile" }`

For the Fargate profile `myFargateProfile`, Ref returns the
physical resource ID of the Fargate profile. For example,
`<cluster-name>/<Fargate_profile_name>`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the cluster, such as
`arn:aws:eks:us-west-2:666666666666:fargateprofile/myCluster/myFargateProfile/1cb1a11a-1dc1-1d11-cf11-1111f11fa111`.

## Remarks

_Creating a Fargate profile and identity provider config_
_resources in the same template._

If CloudFormation attempts to create both resources at the same time, resource
creation fails. If you want to create both resources in the same template, then add
the `DependsOn` property in your template, as shown in the
examples.

## Examples

### Create a Fargate profile

The following example creates a Fargate profile for pods
deployed to a namespace with the name `my-namespace` that have a
label with a key value pair assigned to them. If you're not creating an
`EKSIdpConfig` in the same template, remove the
`"DependsOn"` line in the following example. For more
information, see [`AWS::EKS::IdentityProviderConfig`](../userguide/aws-resource-eks-identityproviderconfig.md).

#### JSON

```json

{
   "Resources" : {
      "EKSFargateProfile" : {
         "DependsOn" : "EKSIdpConfig",
         "Type" : "AWS::EKS::FargateProfile",
         "Properties" : {
            "FargateProfileName" : "my-fargate-profile",
            "ClusterName" : "my-cluster",
            "PodExecutionRoleArn" : "arn:aws:iam::012345678910:role/AmazonEKSFargatePodExecutionRole",
            "Subnets" : [ "subnet-6782e71e", "subnet-e7e761ac" ],
            "Selectors" : [
               {
                  "Namespace" : "my-namespace",
                  "Labels" : [
                     {
                        "Key" : "my-key",
                        "Value" : "my-value"
                     }
                  ]
               }
            ]
         }
      }
   }
}
```

#### YAML

```yaml

Resources:
  EKSFargateProfile:
    DependsOn: EKSIdpConfig
    Type: 'AWS::EKS::FargateProfile'
    Properties:
      FargateProfileName: my-fargate-profile
      ClusterName: my-cluster
      PodExecutionRoleArn: 'arn:aws:iam::012345678910:role/AmazonEKSFargatePodExecutionRole'
      Subnets:
        - subnet-6782e71e
        - subnet-e7e761ac
      Selectors:
        - Namespace: my-namespace
          Labels:
            - Key: my-key
              Value: my-value
```

## See also

- [Fargate profile](../../../eks/latest/userguide/fargate-profile.md) in the _Amazon EKS User Guide_.

- [`CreateFargateProfile`](../../../../reference/eks/latest/apireference/api-createfargateprofile.md) in the _Amazon EKS API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZonalShiftConfig

Label

All content copied from https://docs.aws.amazon.com/.
