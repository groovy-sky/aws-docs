This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::PodIdentityAssociation

Amazon EKS Pod Identity associations provide the ability to manage credentials for your applications, similar to the way that Amazon EC2 instance profiles provide credentials to Amazon EC2 instances.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::PodIdentityAssociation",
  "Properties" : {
      "ClusterName" : String,
      "DisableSessionTags" : Boolean,
      "Namespace" : String,
      "Policy" : String,
      "RoleArn" : String,
      "ServiceAccount" : String,
      "Tags" : [ Tag, ... ],
      "TargetRoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::EKS::PodIdentityAssociation
Properties:
  ClusterName: String
  DisableSessionTags: Boolean
  Namespace: String
  Policy: String
  RoleArn: String
  ServiceAccount: String
  Tags:
    - Tag
  TargetRoleArn: String

```

## Properties

`ClusterName`

The name of the cluster that the association is in.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DisableSessionTags`

The state of the automatic sessions tags. The value of _true_
disables these tags.

EKS Pod Identity adds a pre-defined set of session tags when it assumes the role. You
can use these tags to author a single role that can work across resources by allowing
access to AWS resources based on matching tags. By default, EKS Pod Identity attaches
six tags, including tags for cluster name, namespace, and service account name. For the
list of tags added by EKS Pod Identity, see [List of session tags\
added by EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/pod-id-abac.html#pod-id-abac-tags) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: Boolean

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespace`

The name of the Kubernetes namespace inside the cluster to create the association in. The
service account and the Pods that use the service account must be in this
namespace.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

An optional IAM policy in JSON format (as an escaped string) that applies additional
restrictions to this pod identity association beyond the IAM policies attached to the
IAM role. This policy is applied as the intersection of the role's policies and this
policy, allowing you to reduce the permissions that applications in the pods can use.
Use this policy to enforce least privilege access while still leveraging a shared IAM
role across multiple applications.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to associate with the service account. The EKS Pod Identity
agent manages credentials to assume this role for applications in the containers in the
Pods that use this service account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccount`

The name of the Kubernetes service account inside the cluster to associate the IAM
credentials with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value. You define both. Tags don't
propagate to any other cluster or AWS resources.

The following basic restrictions apply to tags:

- Maximum number of tags per resource – 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length – 128 Unicode characters in UTF-8

- Maximum value length – 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with
this prefix do not count against your tags per resource limit.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-podidentityassociation-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRoleArn`

The Amazon Resource Name (ARN) of the target IAM role to associate with the service account. This
role is assumed by using the EKS Pod Identity association role, then the credentials for this
role are injected into the Pod.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myAssociation" }`

For the Amazon EKS Pod Identity association `myAssociation`, Ref returns the
physical resource ID of the association. For example,
`a-abcdefghijklmnop1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationArn`

The Amazon Resource Name (ARN) of the association.

`AssociationId`

The ID of the association.

`ExternalId`

The unique identifier for this EKS Pod Identity association for a target IAM role. You put this value
in the trust policy of the target role, in a `Condition` to match the
`sts.ExternalId`. This ensures that the target role can only be assumed
by this association. This prevents the _confused deputy problem_. For
more information about the confused deputy problem, see [The confused deputy\
problem](https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html) in the _IAM User Guide_.

If you want to use the same target role with multiple associations or
other roles, use independent statements in the trust policy to allow
`sts:AssumeRole` access from each role.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WarmPoolConfig

Tag
