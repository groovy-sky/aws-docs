This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Addon

Creates an Amazon EKS add-on.

Amazon EKS add-ons help to automate the provisioning and lifecycle management of common
operational software for Amazon EKS clusters. For more information, see [Amazon EKS\
add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::Addon",
  "Properties" : {
      "AddonName" : String,
      "AddonVersion" : String,
      "ClusterName" : String,
      "ConfigurationValues" : String,
      "NamespaceConfig" : NamespaceConfig,
      "PodIdentityAssociations" : [ PodIdentityAssociation, ... ],
      "PreserveOnDelete" : Boolean,
      "ResolveConflicts" : String,
      "ServiceAccountRoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EKS::Addon
Properties:
  AddonName: String
  AddonVersion: String
  ClusterName: String
  ConfigurationValues: String
  NamespaceConfig:
    NamespaceConfig
  PodIdentityAssociations:
    - PodIdentityAssociation
  PreserveOnDelete: Boolean
  ResolveConflicts: String
  ServiceAccountRoleArn: String
  Tags:
    - Tag

```

## Properties

`AddonName`

The name of the add-on.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AddonVersion`

The version of the add-on.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of your cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationValues`

The configuration values that you provided.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceConfig`

The namespace configuration for the addon. This specifies the Kubernetes namespace where the addon is installed.

_Required_: No

_Type_: [NamespaceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-addon-namespaceconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PodIdentityAssociations`

An array of EKS Pod Identity associations owned by the add-on. Each association maps a role to a service
account in a namespace in the cluster.

For more information, see [Attach an IAM Role to an Amazon EKS add-on\
using EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/add-ons-iam.html) in the _Amazon EKS User Guide_.

_Required_: No

_Type_: Array of [PodIdentityAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-addon-podidentityassociation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveOnDelete`

Specifying this option preserves the add-on software on your cluster but Amazon EKS stops
managing any settings for the add-on. If an IAM account is associated with the add-on,
it isn't removed.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolveConflicts`

How to resolve field value conflicts for an Amazon EKS add-on. Conflicts are handled based
on the value you choose:

- **None** – If the self-managed version of
the add-on is installed on your cluster, Amazon EKS doesn't change the value.
Creation of the add-on might fail.

- **Overwrite** – If the self-managed
version of the add-on is installed on your cluster and the Amazon EKS default value
is different than the existing value, Amazon EKS changes the value to the Amazon EKS
default value.

- **Preserve** – This is similar to the NONE
option. If the self-managed version of the add-on is installed on your cluster
Amazon EKS doesn't change the add-on resource properties. Creation of the add-on
might fail if conflicts are detected. This option works differently during the
update operation. For more information, see [`UpdateAddon`](https://docs.aws.amazon.com/eks/latest/APIReference/API_UpdateAddon.html).

If you don't currently have the self-managed version of the add-on installed on your
cluster, the Amazon EKS add-on is installed. Amazon EKS sets all values to default values,
regardless of the option that you specify.

_Required_: No

_Type_: String

_Allowed values_: `NONE | OVERWRITE | PRESERVE`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccountRoleArn`

The Amazon Resource Name (ARN) of an existing IAM role to bind to the add-on's service account. The role must be assigned the IAM permissions required by the add-on. If you don't specify an existing IAM role, then the add-on uses the
permissions assigned to the node IAM role. For more information, see [Amazon EKS node IAM role](https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html) in the _Amazon EKS User Guide_.

###### Note

To specify an existing IAM role, you must have an IAM OpenID Connect (OIDC) provider created for
your cluster. For more information, see [Enabling\
IAM roles for service accounts on your cluster](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html) in the
_Amazon EKS User Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The metadata that you apply to the add-on to assist with categorization and
organization. Each tag consists of a key and an optional value, both of which you
define. Add-on tags do not propagate to any other resources associated with the
cluster.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-addon-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "vpc-cni" }`

For the add-on `vpc-cni`, `Ref` returns the name of the add-on.
For example, `cluster-name|vpc-cni`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the add-on, such as
`arn:aws:eks:us-west-2:111122223333:addon/1-19/vpc-cni/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`.

## See also

- [Amazon EKS add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the _Amazon EKS User Guide_.

- [`CreateAddon`](https://docs.aws.amazon.com/eks/latest/APIReference/API_CreateAddon.html) in the _Amazon EKS API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

NamespaceConfig
