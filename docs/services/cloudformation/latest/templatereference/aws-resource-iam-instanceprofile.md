This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::InstanceProfile

Creates a new instance profile. For information about instance profiles, see [Using\
instance profiles](../../../iam/latest/userguide/id-roles-use-switch-role-ec2-instance-profiles.md).

For information about the number of instance profiles you can create, see [IAM object quotas](../../../iam/latest/userguide/reference-iam-quotas.md) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::InstanceProfile",
  "Properties" : {
      "InstanceProfileName" : String,
      "Path" : String,
      "Roles" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::InstanceProfile
Properties:
  InstanceProfileName: String
  Path: String
  Roles:
    - String

```

## Properties

`InstanceProfileName`

The name of the instance profile to create.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Path`

The path to the instance profile. For more information about paths, see [IAM\
Identifiers](../../../iam/latest/userguide/using-identifiers.md) in the _IAM User Guide_.

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

`Roles`

The name of the role to associate with the instance profile. Only one role can be
assigned to an EC2 instance at a time, and all applications on the instance share the same
role and permissions.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "MyProfile" }`

For the `AWS::IAM::InstanceProfile` resource with the logical ID
`MyProfile`, Ref returns the name of the instance profile.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the instance profile. For example:

`{"Fn::GetAtt" : ["MyProfile", "Arn"] }`

This returns a value such as
`arn:aws:iam::1234567890:instance-profile/MyProfile-ASDNSDLKJ`.

## Examples

### Instance Profile

In this example, the InstanceProfile resource refers to the role by specifying its
name, "MyRole".

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Resources": {
      "MyInstanceProfile": {
         "Type": "AWS::IAM::InstanceProfile",
         "Properties": {
            "Path": "/",
            "Roles": [ {
               "Ref": "MyRole"
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
  MyInstanceProfile:
    Type: "AWS::IAM::InstanceProfile"
    Properties:
      Path: "/"
      Roles:
        -
          Ref: "MyRole"

```

## See also

- [CreateInstanceProfile](../../../../reference/iam/latest/apireference/api-createinstanceprofile.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::GroupPolicy

AWS::IAM::ManagedPolicy

All content copied from https://docs.aws.amazon.com/.
