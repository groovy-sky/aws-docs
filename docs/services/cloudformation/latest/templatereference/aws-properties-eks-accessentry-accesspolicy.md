This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::AccessEntry AccessPolicy

An access policy includes permissions that allow Amazon EKS to authorize an
IAM principal to work with Kubernetes objects on your cluster. The policies are
managed by Amazon EKS, but they're not IAM policies. You can't view the permissions in the
policies using the API. The permissions for many of the policies are similar to the
Kubernetes `cluster-admin`, `admin`, `edit`, and
`view` cluster roles. For more information about these cluster roles, see
[User-facing roles](https://kubernetes.io/docs/reference/access-authn-authz/rbac) in the Kubernetes documentation. To view the contents of the
policies, see [Access\
policy permissions](https://docs.aws.amazon.com/eks/latest/userguide/access-policies.html#access-policy-permissions) in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessScope" : AccessScope,
  "PolicyArn" : String
}

```

### YAML

```yaml

  AccessScope:
    AccessScope
  PolicyArn: String

```

## Properties

`AccessScope`

The scope of an `AccessPolicy` that's associated to an
`AccessEntry`.

_Required_: Yes

_Type_: [AccessScope](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-accessentry-accessscope.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyArn`

The ARN of the access policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EKS::AccessEntry

AccessScope
