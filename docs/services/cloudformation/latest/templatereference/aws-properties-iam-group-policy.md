This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::Group Policy

Contains information about an attached policy.

An attached policy is a managed policy that has been attached to a user, group, or
role.

For more information about managed policies, see [Managed Policies and Inline\
Policies](../../../iam/latest/userguide/policies-managed-vs-inline.md) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyDocument" : Json,
  "PolicyName" : String
}

```

### YAML

```yaml

  PolicyDocument: Json
  PolicyName: String

```

## Properties

`PolicyDocument`

The policy document.

_Required_: Yes

_Type_: Json

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The friendly name (not ARN) identifying the policy.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PolicyDetail](../../../../reference/iam/latest/apireference/api-policydetail.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::Group

AWS::IAM::GroupPolicy

All content copied from https://docs.aws.amazon.com/.
