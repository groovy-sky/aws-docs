This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::PermissionSet PermissionsBoundary

Specifies the configuration of the AWS managed or customer managed policy that you
want to set as a permissions boundary. Specify either
`CustomerManagedPolicyReference` to use the name and path of a customer
managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed
policy. A permissions boundary represents the maximum permissions that any policy can
grant your role. For more information, see [Permissions boundaries\
for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the _IAM User Guide_.

###### Important

Policies used as permissions boundaries don't provide permissions. You must also
attach an IAM policy to the role. To learn how the effective permissions for a
role are evaluated, see [IAM JSON\
policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the _IAM User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManagedPolicyReference" : CustomerManagedPolicyReference,
  "ManagedPolicyArn" : String
}

```

### YAML

```yaml

  CustomerManagedPolicyReference:
    CustomerManagedPolicyReference
  ManagedPolicyArn: String

```

## Properties

`CustomerManagedPolicyReference`

Specifies the name and path of a customer managed policy. You must have an IAM policy that matches the name and path in each AWS account where you want to deploy your permission set.

_Required_: No

_Type_: [CustomerManagedPolicyReference](aws-properties-sso-permissionset-customermanagedpolicyreference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedPolicyArn`

The AWS managed policy ARN that you want to attach to a permission set as a
permissions boundary.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomerManagedPolicyReference

Tag
