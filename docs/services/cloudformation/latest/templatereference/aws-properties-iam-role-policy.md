This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::Role Policy

Contains information about an attached policy.

An attached policy is a managed policy that has been attached to a user, group, or
role.

For more information about managed policies, refer to [Managed Policies and Inline\
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

The entire contents of the policy that defines permissions. For more information, see
[Overview of JSON\
policies](../../../iam/latest/userguide/access-policies.md#access_policies-json).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The friendly name (not ARN) identifying the policy.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Role Policy

This example shows how the policy document is declared.

#### JSON

```json

{
    "PolicyName": "root",
    "PolicyDocument": {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "IamListAccess",
                "Effect": "Allow",
                "Action": [
                    "iam:ListRoles",
                    "iam:ListUsers"
                ],
                "Resource": "*"
            }
        ]
    }
}
```

#### YAML

```yaml

PolicyName: root
PolicyDocument:
   Version: 2012-10-17
   Statement:
      - Sid: IamListAccess
        Effect: Allow
        Action:
          - 'iam:ListRoles'
          - 'iam:ListUsers'
        Resource: '*'
```

## See also

- [PolicyDetail](../../../../reference/iam/latest/apireference/api-policydetail.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::Role

Tag

All content copied from https://docs.aws.amazon.com/.
