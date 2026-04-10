This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::PrefixList

Specifies a managed prefix list. You can add one or more entries to the prefix list.
Each entry consists of a CIDR block and an optional description.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::PrefixList",
  "Properties" : {
      "AddressFamily" : String,
      "Entries" : [ Entry, ... ],
      "MaxEntries" : Integer,
      "PrefixListName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::PrefixList
Properties:
  AddressFamily: String
  Entries:
    - Entry
  MaxEntries: Integer
  PrefixListName: String
  Tags:
    - Tag

```

## Properties

`AddressFamily`

The IP address type.

Valid Values: `IPv4` \| `IPv6`

_Required_: Yes

_Type_: String

_Allowed values_: `IPv4 | IPv6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Entries`

The entries for the prefix list.

_Required_: No

_Type_: Array of [Entry](aws-properties-ec2-prefixlist-entry.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxEntries`

The maximum number of entries for the prefix list. You can't modify the entries and the size
of a prefix list at the same time.

This property is required when you create a prefix list.

_Required_: Conditional

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrefixListName`

A name for the prefix list.

Constraints: Up to 255 characters in length. The name cannot start with `com.amazonaws`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the prefix list.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-prefixlist-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the prefix list.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the prefix list. For example,
`arn:aws:ec2:us-east-1:123456789012:prefix-list/pl-0123123123123abcd`.

`OwnerId`

The ID of the owner of the prefix list. For example, `123456789012`.

`PrefixListId`

The ID of the prefix list. For example, `pl-0123123123123abcd`.

`Version`

The version of the prefix list. For example, `1`.

## Examples

### Create a prefix list

The following example creates an IPv4 prefix list with a maximum of 10 entries,
and creates 2 entries in the prefix list.

#### JSON

```json

{
    "Resources": {
        "NewPrefixList": {
            "Type": "AWS::EC2::PrefixList",
            "Properties": {
                "PrefixListName": "vpc-1-servers",
                "AddressFamily": "IPv4",
                "MaxEntries": 10,
                "Entries": [
                    {
                        "Cidr": "10.0.0.5/32",
                        "Description": "Server 1"
                    },
                    {
                        "Cidr": "10.0.0.10/32",
                        "Description": "Server 2"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": "VPC-1-Servers"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  NewPrefixList:
    Type: AWS::EC2::PrefixList
    Properties:
      PrefixListName: "vpc-1-servers"
      AddressFamily: "IPv4"
      MaxEntries: 10
      Entries:
        - Cidr: "10.0.0.5/32"
          Description: "Server 1"
        - Cidr: "10.0.0.10/32"
          Description: "Server 2"
      Tags:
        - Key: "Name"
          Value: "VPC-1-Servers"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Entry

All content copied from https://docs.aws.amazon.com/.
