This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::Role

Creates a new role for your AWS account.

For more information about roles, see [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html) in the
_IAM User Guide_. For information about quotas for role names
and the number of roles you can create, see [IAM and AWS STS quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the
_IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::Role",
  "Properties" : {
      "AssumeRolePolicyDocument" : Json,
      "Description" : String,
      "ManagedPolicyArns" : [ String, ... ],
      "MaxSessionDuration" : Integer,
      "Path" : String,
      "PermissionsBoundary" : String,
      "Policies" : [ Policy, ... ],
      "RoleName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::Role
Properties:
  AssumeRolePolicyDocument: Json
  Description: String
  ManagedPolicyArns:
    - String
  MaxSessionDuration: Integer
  Path: String
  PermissionsBoundary: String
  Policies:
    - Policy
  RoleName: String
  Tags:
    - Tag

```

## Properties

`AssumeRolePolicyDocument`

The trust policy that is associated with this role. Trust policies define which entities
can assume the role. You can associate only one trust policy with a role. For an example of
a policy that can be used to assume a role, see [Template Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#aws-resource-iam-role--examples). For more information about the elements that you can use in
an IAM policy, see [IAM Policy Elements Reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html) in the _IAM User Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the role that you provide.

_Required_: No

_Type_: String

_Pattern_: `[\u0009\u000A\u000D\u0020-\u007E\u00A1-\u00FF]*`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedPolicyArns`

A list of Amazon Resource Names (ARNs) of the IAM managed policies that
you want to attach to the role.

For more information about ARNs, see [Amazon Resource Names (ARNs) and\
AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSessionDuration`

The maximum session duration (in seconds) that you want to set for the specified role.
If you do not specify a value for this setting, the default value of one hour is
applied. This setting can have a value from 1 hour to 12 hours.

Anyone who assumes the role from the AWS CLI or API can use the
`DurationSeconds` API parameter or the `duration-seconds` AWS CLI parameter to request a longer session. The `MaxSessionDuration` setting
determines the maximum duration that can be requested using the
`DurationSeconds` parameter. If users don't specify a value for the
`DurationSeconds` parameter, their security credentials are valid for one
hour by default. This applies when you use the `AssumeRole*` API operations
or the `assume-role*` AWS CLI operations but does not apply when you use those
operations to create a console URL. For more information, see [Using IAM\
roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the _IAM User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `3600`

_Maximum_: `43200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to the role. For more information about paths, see [IAM\
Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the _IAM User Guide_.

This parameter is optional. If it is not included, it defaults to a slash (/).

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting
of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
In addition, it can contain any ASCII character from the ! ( `\u0021`) through the DEL character ( `\u007F`), including
most punctuation characters, digits, and upper and lowercased letters.

_Required_: No

_Type_: String

_Pattern_: `(\u002F)|(\u002F[\u0021-\u007E]+\u002F)`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionsBoundary`

The ARN of the policy used to set the permissions boundary for the role.

For more information about permissions boundaries, see [Permissions boundaries for IAM\
identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the _IAM User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

Adds or updates an inline policy document that is embedded in the specified IAM role.

When you embed an inline policy in a role, the inline policy is used as part of the
role's access (permissions) policy. The role's trust policy is created at the same time as
the role. You can update a role's trust policy later. For more information about IAM roles, go to [Using Roles to Delegate Permissions and\
Federate Identities](https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).

A role can also have an attached managed policy. For information about policies, see
[Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the _IAM User Guide_.

For information about limits on the number of inline policies that you can embed with a
role, see [Limitations on IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the _IAM User Guide_.

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

_Type_: Array of [Policy](aws-properties-iam-role-policy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleName`

A name for the IAM role, up to 64 characters in length. For valid values,
see the `RoleName` parameter for the [`CreateRole`](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) action in the _IAM User Guide_.

This parameter allows (per its [regex\
pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-.
The role name must be unique within the account. Role names are not distinguished by case.
For example, you cannot create roles named both "Role1" and "role1".

If you don't specify a name, CloudFormation generates a unique physical ID and
uses that ID for the role name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to
acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in CloudFormation\
Templates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities).

###### Important

Naming an IAM resource can cause an unrecoverable error if you reuse
the same template in multiple Regions. To prevent this, we recommend using
`Fn::Join` and `AWS::Region` to create a Region-specific name,
as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref":
               "MyResourceName"}]]}`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags that are attached to the role. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the
_IAM User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iam-role-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For example:

`{ "Ref": "RootRole" }`

For the `AWS::IAM::Role` resource with the logical ID `RootRole`,
`Ref` will return the role name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the role. For example:

`{"Fn::GetAtt" : ["MyRole", "Arn"] }`

This will return a value such as
`arn:aws:iam::1234567890:role/MyRole-AJJHDSKSDF`.

`RoleId`

Returns the stable and unique string identifying the role. For example,
`AIDAJQABLZS4A3QDU576Q`.

For more information about IDs, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html) in the _IAM User Guide_.

## Examples

- [Role with Embedded Policy and Instance Profiles](#aws-resource-iam-role--examples--Role_with_Embedded_Policy_and_Instance_Profiles)

- [Role with External Policy and Instance Profiles](#aws-resource-iam-role--examples--Role_with_External_Policy_and_Instance_Profiles)

### Role with Embedded Policy and Instance Profiles

This example shows an embedded policy in the `AWS::IAM::Role`. The
policy is specified inline in the `Policies` property of the
`AWS::IAM::Role`.

```language_sc3_fgs_qjb

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "RootRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "ec2.amazonaws.com"
                                ]
                            },
                            "Action": [
                                "sts:AssumeRole"
                            ]
                        }
                    ]
                },
                "Path": "/",
                "Policies": [
                    {
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
                        }
                    }
                ]
            }
        },
        "RootInstanceProfile": {
            "Type": "AWS::IAM::InstanceProfile",
            "Properties": {
                "Path": "/",
                "Roles": [
                    {
                        "Ref": "RootRole"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Resources:
  RootRole:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: Allow
            Principal:
              Service:
                - ec2.amazonaws.com
            Action:
              - 'sts:AssumeRole'
      Path: /
      Policies:
        - PolicyName: root
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              - Effect: Allow
                Action: '*'
                Resource: '*'
  RootInstanceProfile:
    Type: 'AWS::IAM::InstanceProfile'
    Properties:
      Path: /
      Roles:
        - !Ref RootRole
```

### Role with External Policy and Instance Profiles

In this example, the Policy and InstanceProfile resources are specified externally
to the IAM Role. They refer to the role by specifying its name,
"RootRole", in their respective `Roles` properties.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Resources": {
      "RootRole": {
         "Type": "AWS::IAM::Role",
         "Properties": {
            "AssumeRolePolicyDocument": {
               "Version" : "2012-10-17",
               "Statement": [ {
                  "Effect": "Allow",
                  "Principal": {
                     "Service": [ "ec2.amazonaws.com" ]
                  },
                  "Action": [ "sts:AssumeRole" ]
               } ]
            },
            "Path": "/"
         }
      },
      "RolePolicies": {
         "Type": "AWS::IAM::Policy",
         "Properties": {
            "PolicyName": "root",
            "PolicyDocument": {
               "Version" : "2012-10-17",
               "Statement": [ {
                  "Effect": "Allow",
                  "Action": "*",
                  "Resource": "*"
               } ]
            },
            "Roles": [ {
               "Ref": "RootRole"
            } ]
         }
      },
      "RootInstanceProfile": {
         "Type": "AWS::IAM::InstanceProfile",
         "Properties": {
            "Path": "/",
            "Roles": [ {
               "Ref": "RootRole"
            } ]
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Resources:
  RootRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Principal:
              Service:
                - "ec2.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
  RolePolicies:
    Type: "AWS::IAM::Policy"
    Properties:
      PolicyName: "root"
      PolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Action: "*"
            Resource: "*"
      Roles:
        - Ref: "RootRole"
  RootInstanceProfile:
    Type: "AWS::IAM::InstanceProfile"
    Properties:
      Path: "/"
      Roles:
        - Ref: "RootRole"

```

## See also

- To view `AWS::IAM::User` template example snippets, see [IAM role template examples](../userguide/quickref-iam.md#scenarios-iamroles).

- [AWS Identity and Access Management Template Snippets](../userguide/quickref-iam.md)

- [CreateRole](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) in the _AWS Identity and Access Management API Reference_

- [AWS::IAM::InstanceProfile](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::Policy

Policy
