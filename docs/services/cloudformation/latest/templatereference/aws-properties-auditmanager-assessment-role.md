This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment Role

The `Role` property type specifies the wrapper that contains AWS Audit Manager role information, such as the role type and IAM Amazon
Resource Name (ARN).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "RoleType" : String
}

```

### YAML

```yaml

  RoleArn: String
  RoleType: String

```

## Properties

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*:iam:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleType`

The type of customer persona.

###### Note

In `CreateAssessment`, `roleType` can only be
`PROCESS_OWNER`.

In `UpdateSettings`, `roleType` can only be
`PROCESS_OWNER`.

In `BatchCreateDelegationByAssessment`, `roleType` can only be
`RESOURCE_OWNER`.

_Required_: No

_Type_: String

_Allowed values_: `PROCESS_OWNER | RESOURCE_OWNER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Role](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_Role.html) in the
_AWS Audit Manager API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Delegation

Scope
