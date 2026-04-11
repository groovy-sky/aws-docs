This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Budgets::BudgetsAction IamActionDefinition

The AWS Identity and Access Management (IAM) action definition details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Groups" : [ String, ... ],
  "PolicyArn" : String,
  "Roles" : [ String, ... ],
  "Users" : [ String, ... ]
}

```

### YAML

```yaml

  Groups:
    - String
  PolicyArn: String
  Roles:
    - String
  Users:
    - String

```

## Properties

`Groups`

A list of groups to be attached. There must be at least one group.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyArn`

The Amazon Resource Name (ARN) of the policy to be attached.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-eusc|-cn|-us-gov|-iso|-iso-[a-z]{1})?:iam::(\d{12}|aws):policy(\u002F[\u0021-\u007F]+\u002F|\u002F)[\w+=,.@-]+$`

_Minimum_: `25`

_Maximum_: `684`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Roles`

A list of roles to be attached. There must be at least one role.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Users`

A list of users to be attached. There must be at least one user.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Definition

ResourceTag

All content copied from https://docs.aws.amazon.com/.
