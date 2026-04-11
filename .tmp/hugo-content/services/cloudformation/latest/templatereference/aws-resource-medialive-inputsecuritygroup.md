This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::InputSecurityGroup

The AWS::MediaLive::InputSecurityGroup is a MediaLive resource type that creates an input security group.

A MediaLive input security group is associated with a MediaLive input. The input security group is an "allow
list" of IP addresses that controls whether an external IP address can push content to the associated MediaLive
input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::InputSecurityGroup",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "WhitelistRules" : [ InputWhitelistRuleCidr, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::InputSecurityGroup
Properties:
  Tags:
    - Tag
  WhitelistRules:
    - InputWhitelistRuleCidr

```

## Properties

`Tags`

A collection of tags for this input security group. Each tag is a key-value pair.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhitelistRules`

The list of IPv4 CIDR addresses to include in the input security group as "allowed" addresses.

_Required_: No

_Type_: Array of [InputWhitelistRuleCidr](aws-properties-medialive-inputsecuritygroup-inputwhitelistrulecidr.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the input security group.

For example: `{ "Ref": "myInputSecurityGroup" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the MediaLive input security group. For example:
arn:aws:medialive:us-west-1:111122223333:medialive:inputSecurityGroup:1234567

`Id`

The Id of the Input Security Group

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtSettingsRequest

InputWhitelistRuleCidr

All content copied from https://docs.aws.amazon.com/.
