This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::RolePolicy

Adds or updates an inline policy document that is embedded in the specified IAM role.

When you embed an inline policy in a role, the inline policy is used as part of the
role's access (permissions) policy. The role's trust policy is created at the same time as
the role, using [`CreateRole`](../../../../reference/iam/latest/apireference/api-createrole.md). You can update a role's trust policy using [`UpdateAssumeRolePolicy`](../../../../reference/iam/latest/apireference/api-updateassumerolepolicy.md). For information about roles, see [IAM roles](../../../iam/latest/userguide/roles-toplevel.md) in the _IAM User Guide_.

A role can also have a managed policy attached to it. To attach a managed policy to a
role, use [`AWS::IAM::Role`](../userguide/aws-resource-iam-role.md). To create a new managed policy, use [`AWS::IAM::ManagedPolicy`](../userguide/aws-resource-iam-managedpolicy.md). For information about policies, see [Managed policies and inline\
policies](../../../iam/latest/userguide/policies-managed-vs-inline.md) in the _IAM User Guide_.

For information about the maximum number of inline policies that you can embed with a
role, see [IAM and AWS STS quotas](../../../iam/latest/userguide/reference-iam-quotas.md) in the _IAM User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::RolePolicy",
  "Properties" : {
      "PolicyDocument" : Json,
      "PolicyName" : String,
      "RoleName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::RolePolicy
Properties:
  PolicyDocument: Json
  PolicyName: String
  RoleName: String

```

## Properties

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

_Required_: No

_Type_: Json

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the policy document.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleName`

The name of the role to associate the policy with.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Role embedded inline policy document](#aws-resource-iam-rolepolicy--examples--Role_embedded_inline_policy_document)

- [Role embedded inline policy document with Ref function](#aws-resource-iam-rolepolicy--examples--Role_embedded_inline_policy_document_with_Ref_function)

### Role embedded inline policy document

This example shows an inline policy document in `AWS::IAM::RolePolicy`.
The policy will be embedded in the specified IAM role,
`RootRole`.

#### JSON

```json

{
    "Type": "AWS::IAM::RolePolicy",
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
        "RoleName": "RootRole"
    }
}
```

#### YAML

```yaml

Type: 'AWS::IAM::RolePolicy'
Properties:
  PolicyName: root
  PolicyDocument:
    Version: "2012-10-17"
    Statement:
      - Effect: Allow
        Action: '*'
        Resource: '*'
  RoleName: RootRole
```

### Role embedded inline policy document with Ref function

This example shows an inline policy document in `AWS::IAM::RolePolicy`.
The example uses the `Ref` function in the `RoleName` property
to specify the logical ID of a `AWS::IAM::Role` resource. For the logical
ID of the `AWS::IAM::Role` resource, `Ref` returns the role
name.

#### JSON

```json

{
    "Type": "AWS::IAM::RolePolicy",
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
        "RoleName": {
            "Ref": "RootRole"
        }
    }
}
```

#### YAML

```yaml

Type: 'AWS::IAM::RolePolicy'
Properties:
  PolicyName: root
  PolicyDocument:
    Version: "2012-10-17"
    Statement:
      - Effect: Allow
        Action: '*'
        Resource: '*'
  RoleName: !Ref RootRole
```

## See also

- [PutRolePolicy](../../../../reference/iam/latest/apireference/api-putrolepolicy.md) in the _AWS Identity and Access Management API Reference_

- [AWS::IAM::Role](../userguide/aws-resource-iam-role.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IAM::SAMLProvider

All content copied from https://docs.aws.amazon.com/.
