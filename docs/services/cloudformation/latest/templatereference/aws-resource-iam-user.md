This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::User

Creates a new IAM user for your AWS account.

For information about quotas for the number of IAM users you can create, see [IAM and AWS STS\
quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::User",
  "Properties" : {
      "Groups" : [ String, ... ],
      "LoginProfile" : LoginProfile,
      "ManagedPolicyArns" : [ String, ... ],
      "Path" : String,
      "PermissionsBoundary" : String,
      "Policies" : [ Policy, ... ],
      "Tags" : [ Tag, ... ],
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::User
Properties:
  Groups:
    - String
  LoginProfile:
    LoginProfile
  ManagedPolicyArns:
    - String
  Path: String
  PermissionsBoundary: String
  Policies:
    - Policy
  Tags:
    - Tag
  UserName: String

```

## Properties

`Groups`

A list of group names to which you want to add the user.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginProfile`

Creates a password for the specified IAM user. A password allows an
IAM user to access AWS services through the AWS Management Console.

You can use the AWS CLI, the AWS API, or the **Users** page in the IAM console to create a password
for any IAM user. Use [ChangePassword](https://docs.aws.amazon.com/IAM/latest/APIReference/API_ChangePassword.html) to update
your own existing password in the **My Security Credentials**
page in the AWS Management Console.

For more information about managing passwords, see [Managing passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the _IAM User Guide_.

_Required_: No

_Type_: [LoginProfile](aws-properties-iam-user-loginprofile.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedPolicyArns`

A list of Amazon Resource Names (ARNs) of the IAM managed policies that
you want to attach to the user.

For more information about ARNs, see [Amazon Resource Names (ARNs) and\
AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path for the user name. For more information about paths, see [IAM\
identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the _IAM User Guide_.

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

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionsBoundary`

The ARN of the managed policy that is used to set the permissions boundary for the
user.

A permissions boundary policy defines the maximum permissions that identity-based
policies can grant to an entity, but does not grant permissions. Permissions boundaries
do not define the maximum permissions that a resource-based policy can grant to an
entity. To learn more, see [Permissions boundaries\
for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the _IAM User Guide_.

For more information about policy types, see [Policy types](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policy-types) in the _IAM User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

Adds or updates an inline policy document that is embedded in the specified IAM user. To view AWS::IAM::User snippets, see [Declaring\
an IAM User Resource](../userguide/quickref-iam.md#scenario-iam-user).

###### Important

The name of each policy for a role, user, or group must be unique. If you don't
choose unique names, updates to the IAM identity will fail.

For information about limits on the number of inline policies that you can embed in a
user, see [Limitations on IAM Entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html) in the _IAM User Guide_.

_Required_: No

_Type_: Array of [Policy](aws-properties-iam-user-policy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags that you want to attach to the new user. Each tag consists of a key name and an associated value.
For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the
_IAM User Guide_.

###### Note

If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request
fails and the resource is not created.

_Required_: No

_Type_: Array of [Tag](aws-properties-iam-user-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

The name of the user to create. Do not include the path in this value.

This parameter allows (per its [regex\
pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-.
The user name must be unique within the account. User names are not distinguished by case.
For example, you cannot create users named both "John" and "john".

If you don't specify a name, CloudFormation generates a unique physical ID and
uses that ID for the user name.

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

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `UserName`. For example:
`mystack-myuser-1CCXAFG2H2U4D`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified `AWS::IAM::User`
resource. For example:
`arn:aws:iam::123456789012:user/mystack-myuser-1CCXAFG2H2U4D`.

## See also

- To view `AWS::IAM::User` template example snippets, see [Declaring an IAM User Resource](../userguide/quickref-iam.md#scenario-iam-user).

- [CreateUser](https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateUser.html) in the _AWS Identity and Access Management API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::ServiceLinkedRole

LoginProfile
