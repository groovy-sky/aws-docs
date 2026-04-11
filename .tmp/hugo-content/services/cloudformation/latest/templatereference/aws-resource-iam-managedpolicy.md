This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::ManagedPolicy

Creates a new managed policy for your AWS account.

This operation creates a policy version with a version identifier of `v1`
and sets v1 as the policy's default version. For more information about policy versions,
see [Versioning for managed policies](../../../iam/latest/userguide/policies-managed-versions.md) in the
_IAM User Guide_.

As a best practice, you can validate your IAM policies.
To learn more, see [Validating IAM policies](../../../iam/latest/userguide/access-policies-policy-validator.md)
in the _IAM User Guide_.

For more information about managed policies in general, see [Managed\
policies and inline policies](../../../iam/latest/userguide/policies-managed-vs-inline.md) in the
_IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::ManagedPolicy",
  "Properties" : {
      "Description" : String,
      "Groups" : [ String, ... ],
      "ManagedPolicyName" : String,
      "Path" : String,
      "PolicyDocument" : Json,
      "Roles" : [ String, ... ],
      "Users" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::ManagedPolicy
Properties:
  Description: String
  Groups:
    - String
  ManagedPolicyName: String
  Path: String
  PolicyDocument: Json
  Roles:
    - String
  Users:
    - String

```

## Properties

`Description`

A friendly description of the policy.

Typically used to store information about the permissions defined in the policy. For
example, "Grants access to production DynamoDB tables."

The policy description is immutable. After a value is assigned, it cannot be
changed.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Groups`

The name (friendly name, not ARN) of the group to attach the policy to.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: Array of String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedPolicyName`

The friendly name of the policy.

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to
acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in CloudFormation\
Templates](../userguide/using-iam-template.md#using-iam-capabilities).

###### Important

Naming an IAM resource can cause an unrecoverable error if you reuse
the same template in multiple Regions. To prevent this, we recommend using
`Fn::Join` and `AWS::Region` to create a Region-specific name,
as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref":
               "MyResourceName"}]]}`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The path for the policy.

For more information about paths, see [IAM identifiers](../../../iam/latest/userguide/using-identifiers.md) in the
_IAM User Guide_.

This parameter is optional. If it is not included, it defaults to a slash (/).

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting
of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
In addition, it can contain any ASCII character from the ! ( `\u0021`) through the DEL character ( `\u007F`), including
most punctuation characters, digits, and upper and lowercased letters.

###### Note

You cannot use an asterisk (\*) in the path name.

_Required_: No

_Type_: String

_Pattern_: `((/[A-Za-z0-9\.,\+@=_-]+)*)/`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDocument`

The JSON policy document that you want to use as the content for the new
policy.

You must provide policies in JSON format in IAM. However, for CloudFormation
templates formatted in YAML, you can provide the policy in JSON or YAML format. CloudFormation always converts a YAML policy to JSON format before submitting it to
IAM.

The maximum length of the policy document that you can pass in this operation,
including whitespace, is listed below. To view the maximum character counts of a managed policy with no whitespaces, see [IAM and AWS STS character quotas](../../../iam/latest/userguide/reference-iam-quotas.md#reference_iam-quotas-entity-length).

To learn more about JSON policy grammar, see [Grammar of the IAM JSON\
policy language](../../../iam/latest/userguide/reference-policies-grammar.md) in the _IAM User Guide_.

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

_Pattern_: `[\u0009\u000A\u000D\u0020-\u00FF]+`

_Minimum_: `1`

_Maximum_: `131072`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Roles`

The name (friendly name, not ARN) of the role to attach the policy to.

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

The name (friendly name, not ARN) of the IAM user to attach the policy to.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: Array of String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN.

In the following sample, the `Ref` function returns the ARN of the
`CreateTestDBPolicy` managed policy, such as
`arn:aws:iam::123456789012:policy/teststack-CreateTestDBPolicy-16M23YE3CS700`.

`{ "Ref": "CreateTestDBPolicy" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttachmentCount`

The number of principal entities (users, groups, and roles) that the policy is attached
to.

`CreateDate`

The date and time, in [ISO 8601 date-time\
format](http://www.iso.org/iso/iso8601), when the policy was created.

`DefaultVersionId`

The identifier for the version of the policy that is set as the default (operative)
version.

For more information about policy versions, see [Versioning for managed\
policies](../../../iam/latest/userguide/policies-managed-versions.md) in the _IAM User Guide_.

`IsAttachable`

Specifies whether the policy can be attached to an IAM user, group, or role.

`PermissionsBoundaryUsageCount`

The number of entities (users and roles) for which the policy is used as the permissions
boundary.

For more information about permissions boundaries, see [Permissions boundaries for IAM\
identities](../../../iam/latest/userguide/access-policies-boundaries.md) in the _IAM User Guide_.

`PolicyArn`

The Amazon Resource Name (ARN) of the managed policy that you want information
about.

For more information about ARNs, see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the _AWS General Reference_.

`PolicyId`

The stable and unique string identifying the policy.

For more information about IDs, see [IAM identifiers](../../../iam/latest/userguide/using-identifiers.md) in the
_IAM User Guide_.

`UpdateDate`

The date and time, in [ISO 8601 date-time\
format](http://www.iso.org/iso/iso8601), when the policy was last updated.

When a policy has only one version, this field contains the date and time when the
policy was created. When a policy has more than one version, this field contains the date
and time when the most recent policy version was created.

## Examples

### Create managed policy

The following example creates a managed policy and associates it with the
`TestDBGroup` group. The managed policy grants users permission to
create t2.micro database instances. The database must use the MySQL database engine
and the instance name must include the prefix `test`.

#### JSON

```json

{
    "CreateTestDBPolicy": {
        "Type": "AWS::IAM::ManagedPolicy",
        "Properties": {
            "Description": "Policy for creating a test database",
            "Path": "/",
            "PolicyDocument": {
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Action": "rds:CreateDBInstance",
                        "Resource": {
                            "Fn::Join": [
                                "",
                                [
                                    "arn:aws:rds:",
                                    {
                                        "Ref": "AWS::Region"
                                    },
                                    ":",
                                    {
                                        "Ref": "AWS::AccountId"
                                    },
                                    ":db:test*"
                                ]
                            ]
                        },
                        "Condition": {
                            "StringEquals": {
                                "rds:DatabaseEngine": "mysql"
                            }
                        }
                    },
                    {
                        "Effect": "Allow",
                        "Action": "rds:CreateDBInstance",
                        "Resource": {
                            "Fn::Join": [
                                "",
                                [
                                    "arn:aws:rds:",
                                    {
                                        "Ref": "AWS::Region"
                                    },
                                    ":",
                                    {
                                        "Ref": "AWS::AccountId"
                                    },
                                    ":db:test*"
                                ]
                            ]
                        },
                        "Condition": {
                            "StringEquals": {
                                "rds:DatabaseClass": "db.t2.micro"
                            }
                        }
                    }
                ]
            },
            "Groups": [
                "TestDBGroup"
            ]
        }
    }
}
```

#### YAML

```yaml

CreateTestDBPolicy:
  Type: 'AWS::IAM::ManagedPolicy'
  Properties:
    Description: Policy for creating a test database
    Path: /
    PolicyDocument:
      Version: "2012-10-17"
      Statement:
        - Effect: Allow
          Action: 'rds:CreateDBInstance'
          Resource: !Join
            - ''
            - - 'arn:aws:rds:'
              - !Ref 'AWS::Region'
              - ':'
              - !Ref 'AWS::AccountId'
              - ':db:test*'
          Condition:
            StringEquals:
              'rds:DatabaseEngine': mysql
        - Effect: Allow
          Action: 'rds:CreateDBInstance'
          Resource: !Join
            - ''
            - - 'arn:aws:rds:'
              - !Ref 'AWS::Region'
              - ':'
              - !Ref 'AWS::AccountId'
              - ':db:test*'
          Condition:
            StringEquals:
              'rds:DatabaseClass': db.t2.micro
    Groups:
      - TestDBGroup
```

## See also

- [CreatePolicy](../../../../reference/iam/latest/apireference/api-createpolicy.md) in the _AWS Identity and Access Management API Reference_

- [CreatePolicyVersion](../../../../reference/iam/latest/apireference/api-createpolicyversion.md) in the _AWS Identity and Access Management API Reference_

- [AttachGroupPolicy](../../../../reference/iam/latest/apireference/api-attachgrouppolicy.md) in the _AWS Identity and Access Management API Reference_

- [AttachRolePolicy](../../../../reference/iam/latest/apireference/api-attachrolepolicy.md) in the _AWS Identity and Access Management API Reference_

- [AttachUserPolicy](../../../../reference/iam/latest/apireference/api-attachuserpolicy.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::InstanceProfile

AWS::IAM::OIDCProvider

All content copied from https://docs.aws.amazon.com/.
