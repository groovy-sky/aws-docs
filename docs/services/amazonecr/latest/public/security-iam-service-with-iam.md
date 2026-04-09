# How Amazon ECR Public works with IAM

Before you use IAM to manage access to Amazon ECR, you should understand what IAM
features are available to use with Amazon ECR. To get a high-level view of how Amazon ECR and
other AWS services work with IAM, see [AWS services\
that work with IAM](../../../iam/latest/userguide/reference-aws-services-that-work-with-iam.md) in the _IAM User Guide_.

###### Topics

- [Amazon ECR identity-based policies](#security_iam_service-with-iam-id-based-policies)

- [Amazon ECR resource-based policies](#security_iam_service-with-iam-resource-based-policies)

- [Amazon ECR IAM roles](#security_iam_service-with-iam-roles)

## Amazon ECR identity-based policies

With IAM identity-based policies, you can specify allowed or denied actions and
resources as well as the conditions under which actions are allowed or denied. Amazon ECR
supports specific actions, resources, and condition keys when managing a public
registry and the resources in a public registry. The image pull permissions can't be
changed because Amazon ECR Public repositories are public accessible. To learn about all
of the elements that you use in a JSON policy, see [IAM JSON policy elements\
reference](../../../iam/latest/userguide/reference-policies-elements.md) in the _IAM User Guide_.

### Actions

Administrators can use AWS JSON policies to specify who has access to what. That is, which **principal** can perform
**actions** on what **resources**, and under what **conditions**.

The `Action` element of a JSON policy describes the
actions that you can use to allow or deny access in a policy. Include actions in a policy to grant permissions to perform the associated operation.

Policy actions in Amazon ECR Public use the following prefix before the action:
`ecr-public:`. For example, to grant someone permission to create
a public Amazon ECR repository with the Amazon ECR Public `CreateRepository`
API operation, you include the `ecr-public:CreateRepository` action
in their policy. Policy statements must include either an `Action` or
`NotAction` element. Amazon ECR defines its own set of actions that
describe tasks that you can perform with this service.

To specify multiple actions in a single statement, separate them with commas
as follows:

```JSON

"Action": [
      "ecr-public:action1",
      "ecr-public:action2"
```

You can specify multiple actions using wildcards (\*). For example, to specify
all actions that begin with the word `Describe`, include the
following action:

```JSON

"Action": "ecr-public:Describe*"
```

To see a list of Amazon ECR actions, see [Actions, Resources, and Condition Keys for Amazon ECR Public](../../../iam/latest/userguide/list-amazonelasticcontainerregistrypublic.md) in the
_IAM User Guide_.

### Resources

Administrators can use AWS JSON policies to specify who has access to what. That is, which **principal** can perform
**actions** on what **resources**, and under what **conditions**.

The `Resource` JSON policy element specifies the object or objects to which the action applies. As a best practice, specify a resource using its [Amazon Resource Name (ARN)](../../../iam/latest/userguide/reference-arns.md). For actions that don't support resource-level permissions, use a wildcard (\*) to indicate that the statement applies to all resources.

```

"Resource": "*"
```

An Amazon ECR public repository resource has the following ARN. You don't use a
Region name in the ARN.

```JSON

arn:${Partition}:ecr-public::${Account}:repository/${Repository-name}
```

For more information about the format of ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](../../../../general/latest/gr/aws-arns-and-namespaces.md).

For example, to specify the `my-repo` public repository in your
statement, use the following ARN:

```JSON

"Resource": "arn:aws:ecr-public::123456789012:repository/my-repo"
```

To specify all public repositories that belong to a specific account, use the
wildcard (\*):

```JSON

"Resource": "arn:aws:ecr-public::123456789012:repository/*"
```

To specify multiple resources in a single statement, separate the ARNs with
commas.

```JSON

"Resource": [
      "resource1",
      "resource2"
```

To see a list of Amazon ECR Public resource types and their ARNs, see [Resources defined by Amazon ECR Public](../../../iam/latest/userguide/list-amazonelasticcontainerregistrypublic.md#amazonelasticcontainerregistrypublic-resources-for-iam-policies) in the
_IAM User Guide_. To learn with which actions you can
specify the ARN of each resource, see [Actions defined by Amazon ECR](../../../iam/latest/userguide/list-amazonelasticcontainerregistrypublic.md#amazonelasticcontainerregistrypublic-actions-as-permissions).

### Condition Keys

Administrators can use AWS JSON policies to specify who has access to what. That is, which **principal** can perform
**actions** on what **resources**, and under what **conditions**.

The `Condition` element specifies when statements execute based on defined criteria. You can create conditional expressions that use [condition\
operators](../../../iam/latest/userguide/reference-policies-elements-condition-operators.md), such as equals or less than, to match the condition in the
policy with values in the request. To see all AWS global
condition keys, see [AWS global condition context keys](../../../iam/latest/userguide/reference-policies-condition-keys.md) in the
_IAM User Guide_.

Amazon ECR supports using some global condition keys. To
see all AWS global condition keys, see [AWS global\
condition context keys](../../../iam/latest/userguide/reference-policies-condition-keys.md) in the
_IAM User Guide_.

To see a list of Amazon ECR Public condition keys, see [Condition keys defined by Amazon ECR Public](../../../service-authorization/latest/reference/list-amazonelasticcontainerregistrypublic.md#amazonelasticcontainerregistrypublic-policy-keys) in the
_IAM User Guide_. To learn with which actions and
resources you can use a condition key, see [Actions defined by Amazon ECR](../../../iam/latest/userguide/list-amazonelasticcontainerregistrypublic.md#amazonelasticcontainerregistrypublic-actions-as-permissions).

### Examples

To view examples of Amazon ECR identity-based policies, see [Amazon ECR public identity-based policy examples](security-iam-id-based-policy-examples.md).

## Amazon ECR resource-based policies

Resource-based policies are JSON policy documents that specify what actions a
specified principal can perform on an Amazon ECR Public resource and under what
conditions. Amazon ECR supports resource-based permissions policies for Amazon ECR public
repositories. Resource-based policies let you grant usage permission to other
accounts on a per-resource basis. You can also use a resource-based policy to allow
an AWS service to access your Amazon ECR public repositories.

To enable cross-account access, you can specify an entire account or IAM
entities in another account as the [principal in a\
resource-based policy](../../../iam/latest/userguide/reference-policies-elements-principal.md). Adding a cross-account principal to a
resource-based policy is only half of establishing the trust relationship. When the
principal and the resource are in different AWS accounts, you must also grant the
principal entity permission to access the resource. Grant permission by attaching an
identity-based policy to the entity. However, if a resource-based policy grants
access to a principal in the same account, no additional identity-based policy is
required. For more information, see [How IAM roles\
differ from resource-based policies](../../../iam/latest/userguide/access-policies-identity-vs-resource.md) in the
_IAM User Guide_.

The Amazon ECR Public service supports only one type of resource-based policy called a
_repository policy_, which is attached to a
_repository_. This policy defines which principal entities
(accounts, users, roles, and federated users) can perform actions on the
repository.

To learn how to attach a resource-based policy to a repository, see [Public repository policies in Amazon ECR Public](public-repository-policies.md).

### Examples

To view examples of Amazon ECR resource-based policies, see [Public repository policy examples in Amazon ECR Public](public-repository-policy-examples.md).

## Amazon ECR IAM roles

An [IAM role](../../../iam/latest/userguide/id-roles.md) is an entity within
your AWS account that has specific permissions.

### Using temporary credentials with Amazon ECR

You can use temporary credentials to sign in with federation, assume an IAM
role, or to assume a cross-account role. You obtain temporary security
credentials by calling AWS STS API operations such as [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) or [GetFederationToken](../../../../reference/sts/latest/apireference/api-getfederationtoken.md).

Amazon ECR Public supports using temporary credentials.

### Service-linked roles

[Service-linked roles](../../../iam/latest/userguide/id-roles-terms-and-concepts.md#iam-term-service-linked-role) allow AWS services to access resources in
other services to complete an action on your behalf. Service-linked roles appear
in your IAM account and are owned by the service. An IAM administrator can
view but not edit the permissions for service-linked roles.

Amazon ECR Public does not support service-linked
roles.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Identity and Access Management

AWS managed policies for Amazon ECR Public

All content copied from https://docs.aws.amazon.com/.
