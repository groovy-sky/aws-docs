This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::Policy

Adds or updates an inline policy document that is embedded in the specified IAM group, user or role.

An IAM user can also have a managed policy attached to it. For
information about policies, see [Managed Policies and Inline\
Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the _IAM User Guide_.

The Groups, Roles, and Users properties are optional. However, you must specify at least
one of these properties.

For information about policy documents, see [Creating IAM\
policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html) in the _IAM User Guide_.

For information about limits on the number of inline policies that you can embed in an
identity, see [Limitations on IAM\
Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the _IAM User Guide_.

###### Important

This resource does not support [drift detection](../userguide/using-cfn-stack-drift.md). The following inline policy resource types support drift detection:

- [`AWS::IAM::GroupPolicy`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html)

- [`AWS::IAM::RolePolicy`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html)

- [`AWS::IAM::UserPolicy`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::Policy",
  "Properties" : {
      "Groups" : [ String, ... ],
      "PolicyDocument" : Json,
      "PolicyName" : String,
      "Roles" : [ String, ... ],
      "Users" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::Policy
Properties:
  Groups:
    - String
  PolicyDocument: Json
  PolicyName: String
  Roles:
    - String
  Users:
    - String

```

## Properties

`Groups`

The name of the group to associate the policy with.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-.

_Required_: No

_Type_: Array of String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDocument`

The policy document.

You must provide policies in JSON format in IAM. However, for CloudFormation
templates formatted in YAML, you can provide the policy in JSON or YAML format. CloudFormation always converts a YAML policy to JSON format before submitting it to
IAM.

The [regex pattern](http://wikipedia.org/wiki/regex)
used to validate this parameter is a string of characters consisting of the following:

- Any printable ASCII
character ranging from the space character ( `\u0020`) through the end of the ASCII character range

- The printable characters in the Basic Latin and Latin-1 Supplement character set
(through `\u00FF`)

- The special characters tab ( `\u0009`), line feed ( `\u000A`), and
carriage return ( `\u000D`)

_Required_: Yes

_Type_: Json

_Minimum_: `1`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the policy document.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Roles`

The name of the role to associate the policy with.

This parameter allows (per its [regex\
pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters:
\_+=,.@-

###### Note

If an external policy (such as `AWS::IAM::Policy` or
`AWS::IAM::ManagedPolicy`) has a `Ref` to a role and if a
resource (such as `AWS::ECS::Service`) also has a `Ref` to the
same role, add a `DependsOn` attribute to the resource to make the resource
depend on the external policy. This dependency ensures that the role's policy is
available throughout the resource's lifecycle. For example, when you delete a stack with
an `AWS::ECS::Service` resource, the `DependsOn` attribute ensures
that CloudFormation deletes the `AWS::ECS::Service` resource before
deleting its role's policy.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Users`

The name of the user to associate the policy with.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: Array of String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The stable and unique string identifying the policy.

For more information about IDs, see [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the
_IAM User Guide_.

## Examples

- [Policy with policy group](#aws-resource-iam-policy--examples--Policy_with_policy_group)

- [Policy with specified role](#aws-resource-iam-policy--examples--Policy_with_specified_role)

### Policy with policy group

#### JSON

```json

{
    "Type": "AWS::IAM::Policy",
    "Properties": {
        "PolicyName": "CFNUsers",
        "PolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Action": [
                        "cloudformation:Describe*",
                        "cloudformation:List*",
                        "cloudformation:Get*"
                    ],
                    "Resource": "*"
                }
            ]
        },
        "Groups": [
            {
                "Ref": "CFNUserGroup"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::IAM::Policy'
Properties:
  PolicyName: CFNUsers
  PolicyDocument:
    Version: "2012-10-17"
    Statement:
      - Effect: Allow
        Action:
          - 'cloudformation:Describe*'
          - 'cloudformation:List*'
          - 'cloudformation:Get*'
        Resource: '*'
  Groups:
    - !Ref CFNUserGroup
```

### Policy with specified role

#### JSON

```json

{
    "Type": "AWS::IAM::Policy",
    "Properties": {
        "PolicyName": "root",
        "PolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Action": "*",
                    "Resource": "*"
                }
            ]
        },
        "Roles": [
            {
                "Ref": "RootRole"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::IAM::Policy'
Properties:
  PolicyName: root
  PolicyDocument:
    Version: "2012-10-17"
    Statement:
      - Effect: Allow
        Action: '*'
        Resource: '*'
  Roles:
    - !Ref RootRole
```

## See also

- [AWS::IAM::GroupPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html)

- [AWS::IAM::RolePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html)

- [AWS::IAM::UserPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-userpolicy.html)

- [PutGroupPolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutGroupPolicy.html) in the _AWS Identity and Access Management API Reference_

- [PutRolePolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutRolePolicy.html) in the _AWS Identity and Access Management API Reference_

- [PutUserPolicy](https://docs.aws.amazon.com/IAM/latest/APIReference/API_PutUserPolicy.html) in the _AWS Identity and Access Management API Reference_

- [IAM\
JSON policy reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in the _AWS Identity and Access Management User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IAM::Role
