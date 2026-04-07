This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Group

Creates or updates a group which you can use to associate canaries with
each other, including cross-Region canaries. Using groups can help you with managing
and automating your canaries, and you can also view aggregated run results and statistics for all canaries in a group.

Groups are global resources. When you create a group, it is replicated across all AWS Regions, and you
can add canaries from any Region to it, and view it in any Region. Although the group ARN format
reflects the Region name where it was created, a group is not constrained to any Region. This
means that you can put canaries from multiple Regions into the same group, and then use that
group to view and manage all of those canaries in a single view.

Each group can contain as many as 10 canaries. You can have as many as 20 groups in your account.
Any single canary can be a member of up to 10 groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Synthetics::Group",
  "Properties" : {
      "Name" : String,
      "ResourceArns" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Synthetics::Group
Properties:
  Name: String
  ResourceArns:
    - String
  Tags:
    - Tag

```

## Properties

`Name`

A name for the group. It can include any Unicode characters.

The names for all groups in your account, across all Regions, must be unique.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z_\-]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArns`

The ARNs of the canaries that
you want to associate with this group.

_Required_: No

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value pairs that are associated with the
group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-synthetics-group-tag.html)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the group.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The Id of the group.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VPCConfig

Tag
