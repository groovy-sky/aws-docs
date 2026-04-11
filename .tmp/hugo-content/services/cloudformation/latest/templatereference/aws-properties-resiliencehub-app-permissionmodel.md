This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::App PermissionModel

Defines the roles and credentials that AWS Resilience Hub would use while creating the
application, importing its resources, and running an assessment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CrossAccountRoleArns" : [ String, ... ],
  "InvokerRoleName" : String,
  "Type" : String
}

```

### YAML

```yaml

  CrossAccountRoleArns:
    - String
  InvokerRoleName: String
  Type: String

```

## Properties

`CrossAccountRoleArns`

Defines a list of role Amazon Resource Names (ARNs) to be used in other accounts. These
ARNs are used for querying purposes while importing resources and assessing your
application.

###### Note

- These ARNs are required only when your resources are in other accounts and you have
different role name in these accounts. Else, the invoker role name will be used in the
other accounts.

- These roles must have a trust policy with `iam:AssumeRole` permission to
the invoker role in the primary account.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvokerRoleName`

Existing AWS
IAM role name in the primary AWS account that will be assumed by
AWS Resilience Hub Service Principle to obtain a read-only access to your application
resources while running an assessment.

If your IAM role includes a path, you must include the path in the `invokerRoleName` parameter.
For example, if your IAM role's ARN is `arn:aws:iam:123456789012:role/my-path/role-name`, you should pass `my-path/role-name`.

###### Note

- You must have `iam:passRole` permission for this role while creating or
updating the application.

- Currently, `invokerRoleName` accepts only `[A-Za-z0-9_+=,.@-]`
characters.

_Required_: No

_Type_: String

_Pattern_: `((\u002F[\u0021-\u007E]+\u002F){1,511})?[A-Za-z0-9+=,.@_/-]{1,64}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Defines how AWS Resilience Hub scans your resources. It can scan for the resources by
using a pre-existing role in your AWS account, or by using the credentials of
the current IAM user.

_Required_: Yes

_Type_: String

_Allowed values_: `LegacyIAMUser | RoleBased`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventSubscription

PhysicalResourceId

All content copied from https://docs.aws.amazon.com/.
