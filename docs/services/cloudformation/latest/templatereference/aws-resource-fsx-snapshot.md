This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Snapshot

A snapshot of an Amazon FSx for OpenZFS volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::Snapshot",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "VolumeId" : String
    }
}

```

### YAML

```yaml

Type: AWS::FSx::Snapshot
Properties:
  Name: String
  Tags:
    - Tag
  VolumeId: String

```

## Properties

`Name`

The name of the snapshot.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_:.-]{1,203}$`

_Minimum_: `1`

_Maximum_: `203`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of `Tag` values, with a maximum of 50 elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-snapshot-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeId`

The ID of the volume that the snapshot is of.

_Required_: Yes

_Type_: String

_Pattern_: `^(fsvol-[0-9a-f]{17,})$`

_Minimum_: `23`

_Maximum_: `23`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the snapshot. For example:

`{"Ref":"logical_snapshot_id"}`

Returns `fsvolsnap-0123456789abcedf5`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceARN`

Returns the snapshot's Amazon Resource Name (ARN).

Example: `arn:aws:fsx:us-east-2:111133334444:snapshot/fsvol-01234567890123456/fsvolsnap-0123456789abcedf5`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3AccessPointVpcConfiguration

Tag
