This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Profiles::ProfileResourceAssociation

The association between a Route 53 Profile and resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Profiles::ProfileResourceAssociation",
  "Properties" : {
      "Name" : String,
      "ProfileId" : String,
      "ResourceArn" : String,
      "ResourceProperties" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Profiles::ProfileResourceAssociation
Properties:
  Name: String
  ProfileId: String
  ResourceArn: String
  ResourceProperties: String

```

## Properties

`Name`

Name of the Profile resource association.

_Required_: Yes

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileId`

Profile ID of the Profile that the resources are associated with.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArn`

The Amazon Resource Name (ARN) of the resource association.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceProperties`

If the DNS resource is a DNS Firewall rule group, this indicates the priority.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `ProfileResourceAssociation` ID for the profile resource association, such as `rpr-001ba2563f9example6`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

ID of the Profile resource association.

`ResourceType`

Resource type, such as a private hosted zone, interface VPC endpoint, or DNS Firewall rule group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
