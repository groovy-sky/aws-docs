This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Profiles::Profile

A complex type that includes settings for a Route 53 Profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Profiles::Profile",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Profiles::Profile
Properties:
  Name: String
  Tags:
    - Tag

```

## Properties

`Name`

Name of the Profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of the tag keys and values that you want to associate with the profile.

_Required_: No

_Type_: Array of [Tag](aws-properties-route53profiles-profile-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `Profile` ID for the profile, such as `rp-9e6daa461example`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the Profile.

`ClientToken`

The `ClientToken` value that was assigned when the Profile was created.

`Id`

ID of the Profile.

`ShareStatus`

Sharing status for the Profile.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Route 53 Profiles

Tag

All content copied from https://docs.aws.amazon.com/.
