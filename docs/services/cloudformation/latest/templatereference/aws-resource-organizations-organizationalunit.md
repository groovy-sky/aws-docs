This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Organizations::OrganizationalUnit

Creates an organizational unit (OU) within a root or parent OU. An OU is a container
for accounts that enables you to organize your accounts to apply policies according to
your business requirements. The number of levels deep that you can nest OUs is dependent
upon the policy types enabled for that root. For service control policies, the limit is
five.

For more information about OUs, see [Managing organizational units (OUs)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html) in the
_AWS Organizations User Guide_.

If the request includes tags, then the requester must have the
`organizations:TagResource` permission.

You can only call this operation from the management account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Organizations::OrganizationalUnit",
  "Properties" : {
      "Name" : String,
      "ParentId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Organizations::OrganizationalUnit
Properties:
  Name: String
  ParentId: String
  Tags:
    - Tag

```

## Properties

`Name`

The friendly name of this OU.

The [regex pattern](http://wikipedia.org/wiki/regex)
that is used to validate this parameter is a string of any of the characters in the ASCII
character range.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentId`

The unique identifier (ID) of the parent root or OU that you want to create the new OU
in.

###### Important

To update the `ParentId` parameter value, you must first remove all
accounts attached to the organizational unit (OU). OUs can't be moved within the
organization with accounts still attached.

The [regex pattern](http://wikipedia.org/wiki/regex) for a parent ID
string requires one of the following:

- **Root** \- A string that begins with "r-" followed
by from 4 to 32 lowercase letters or digits.

- **Organizational unit (OU)** \- A string that begins
with "ou-" followed by from 4 to 32 lowercase letters or digits (the ID of the
root that the OU is in). This string is followed by a second "-" dash and from 8
to 32 additional lowercase letters or digits.

_Required_: Yes

_Type_: String

_Pattern_: `^(r-[0-9a-z]{4,32})|(ou-[0-9a-z]{4,32}-[a-z0-9]{8,32})$`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags that you want to attach to the newly created OU. For each tag in the
list, you must specify both a tag key and a value. You can set the value to an empty
string, but you can't set it to `null`. For more information about tagging,
see [Tagging AWS Organizations\
resources](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) in the AWS Organizations User Guide.

###### Note

If any one of the tags is not valid or if you exceed the allowed number of tags
for an OU, then the entire request fails and the OU is not created.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-organizations-organizationalunit-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Id`. For example:
`ou-examplerootid111-exampleouid111`.

###### Note

When creating child OUs, we recommend that you use the `Ref` function
instead of `Fn::GetAtt`. For example, in the properties for the child OU,
use `ParentId: !Ref ParentOU`, instead of `ParentId: !GetAtt
                    'ParentOU.Id'`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of this OU. For example:
`arn:aws:organizations::111111111111:ou/o-exampleorgid/ou-examplerootid111-exampleouid111`.

`Id`

The unique identifier (ID) associated with this OU. For example:
`ou-examplerootid111-exampleouid111`.

## Examples

- [Specify an OU under the root](#aws-resource-organizations-organizationalunit--examples--Specify_an_OU_under_the_root)

- [Specify an OU under a parent OU](#aws-resource-organizations-organizationalunit--examples--Specify_an_OU_under_a_parent_OU)

### Specify an OU under the root

This example illustrates how to specify an OU using
`AWS::Organizations::OrganizationalUnit` directly under the
root.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation Organizations Template Example",
    "Parameters": {
        "OrganizationRootId": {
            "Type": "String"
        }
    },
    "Resources": {
        "TestTemplateOU": {
            "Type": "AWS::Organizations::OrganizationalUnit",
            "Properties": {
                "Name": "TestTemplateOU",
                "ParentId": {
                    "Ref": "OrganizationRootId"
                }
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS CloudFormation Organizations Template Example
Parameters:
  OrganizationRootId:
    Type: String
Resources:
  TestTemplateOU:
    Type: 'AWS::Organizations::OrganizationalUnit'
    Properties:
      Name: TestTemplateOU
      ParentId: !Ref OrganizationRootId

```

### Specify an OU under a parent OU

This example illustrates how to specify a nested OU using
`AWS::Organizations::OrganizationalUnit` by referencing another
OU.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation Nested OU Template Example",
    "Parameters": {
        "OrganizationRootId": {
            "Type": "String"
        }
    },
    "Resources": {
        "ParentOU": {
            "Type": "AWS::Organizations::OrganizationalUnit",
            "Properties": {
                "Name": "ParentOU",
                "ParentId": {
                    "Ref": "OrganizationRootId"
                }
            }
        },
        "ChildOU": {
            "Type": "AWS::Organizations::OrganizationalUnit",
            "Properties": {
                "Name": "ChildOU",
                "ParentId": {
                    "Ref": "ParentOU"
                }
            }
        }
    }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS CloudFormation Nested OU Template Example
Parameters:
  OrganizationRootId:
    Type: String
Resources:
  ParentOU:
    Type: 'AWS::Organizations::OrganizationalUnit'
    Properties:
      Name: ParentOU
      ParentId: !Ref OrganizationRootId
  ChildOU:
    Type: 'AWS::Organizations::OrganizationalUnit'
    Properties:
      Name: ChildOU
      ParentId: !Ref ParentOU

```

## See also

- [Creating an OU](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html#create_ou) in the _AWS Organizations User_
_Guide_.

- [CreateOrganizationalUnit](https://docs.aws.amazon.com/organizations/latest/APIReference/API_CreateOrganizationalUnit.html) in the _AWS Organizations_
_API Reference Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Organizations::Organization

Tag
