This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment

Updates the configuration, dimension, and other settings for an existing
segment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::Segment",
  "Properties" : {
      "ApplicationId" : String,
      "Dimensions" : SegmentDimensions,
      "Name" : String,
      "SegmentGroups" : SegmentGroups,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::Segment
Properties:
  ApplicationId: String
  Dimensions:
    SegmentDimensions
  Name: String
  SegmentGroups:
    SegmentGroups
  Tags:
    - Tag

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that the segment is
associated with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Dimensions`

An array that defines the dimensions for the segment.

_Required_: No

_Type_: [SegmentDimensions](aws-properties-pinpoint-segment-segmentdimensions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the segment.

###### Note

A segment must have a name otherwise it will not appear in the Amazon Pinpoint console.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentGroups`

The segment group to use and the dimensions to apply to the group's base
segments in order to build the segment. A segment group can consist of zero or
more base segments. Your request can include only one segment group.

_Required_: No

_Type_: [SegmentGroups](aws-properties-pinpoint-segment-segmentgroups.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that the segment is associated with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the segment.

`SegmentId`

The unique identifier for the segment.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultPushNotificationTemplate

Behavior

All content copied from https://docs.aws.amazon.com/.
