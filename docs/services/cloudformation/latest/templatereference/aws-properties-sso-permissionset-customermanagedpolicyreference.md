This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::PermissionSet CustomerManagedPolicyReference

Specifies the name and path of a customer managed policy. You must have an IAM policy that matches the name and path in each AWS account where you want to deploy your permission set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Path" : String
}

```

### YAML

```yaml

  Name: String
  Path: String

```

## Properties

`Name`

The name of the IAM policy that you have configured in each account where you want
to deploy your permission set.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to the IAM policy that you have configured in each account where you want
to deploy your permission set. The default is `/`. For more information, see
[Friendly names and paths](../../../iam/latest/userguide/reference-identifiers.md#identifiers-friendly-names) in the _IAM User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `((/[A-Za-z0-9\.,\+@=_-]+)*)/`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSO::PermissionSet

PermissionsBoundary

All content copied from https://docs.aws.amazon.com/.
