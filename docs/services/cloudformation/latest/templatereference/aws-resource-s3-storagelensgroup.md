This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLensGroup

The `AWS::S3::StorageLensGroup` resource creates an S3 Storage Lens group. A
Storage Lens group is a custom grouping of objects that include filters for prefixes,
suffixes, object tags, object size, or object age. You can create an S3 Storage Lens group
that includes a single filter or multiple filter conditions. To specify multiple filter
conditions, you use `AND` or `OR` logical operators. For more
information about S3 Storage Lens groups, see [Working with S3 Storage Lens groups](../../../s3/latest/userguide/storage-lens-groups-overview.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::StorageLensGroup",
  "Properties" : {
      "Filter" : Filter,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3::StorageLensGroup
Properties:
  Filter:
    Filter
  Name: String
  Tags:
    - Tag

```

## Properties

`Filter`

This property contains the criteria for the Storage Lens group data that is
displayed

_Required_: Yes

_Type_: [Filter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelensgroup-filter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

This property contains the Storage Lens group name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

This property contains the AWS resource tags that you're adding to your Storage Lens
group. This parameter is optional.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-storagelensgroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the S3 Storage Lens group name. For more information about using the Ref function, see
[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. For more
information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`StorageLensGroupArn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

And
