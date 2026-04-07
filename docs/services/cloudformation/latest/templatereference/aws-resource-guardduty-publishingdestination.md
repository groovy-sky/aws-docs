This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::PublishingDestination

Creates a publishing destination where you can export your GuardDuty findings. Before
you start exporting the findings, the destination resource must exist.

For more information about considerations and permissions, see [Exporting\
GuardDuty findings to Amazon S3 buckets](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_exportfindings.html) in the _Amazon GuardDuty User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::PublishingDestination",
  "Properties" : {
      "DestinationProperties" : CFNDestinationProperties,
      "DestinationType" : String,
      "DetectorId" : String,
      "Tags" : [ TagItem, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::PublishingDestination
Properties:
  DestinationProperties:
    CFNDestinationProperties
  DestinationType: String
  DetectorId: String
  Tags:
    - TagItem

```

## Properties

`DestinationProperties`

Contains the Amazon Resource Name (ARN) of the resource to publish to, such as an S3
bucket, and the ARN of the KMS key to use to encrypt published findings.

_Required_: Yes

_Type_: [CFNDestinationProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-publishingdestination-cfndestinationproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationType`

The type of publishing destination. GuardDuty supports Amazon S3
buckets as a publishing destination.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The ID of the GuardDuty detector where the publishing destination exists.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Describes a tag.

_Required_: No

_Type_: Array of [TagItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-publishingdestination-tagitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource publishing destination ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the publishing destination.

`PublishingFailureStartTimestamp`

The time, in epoch millisecond format, at which GuardDuty was first unable to publish
findings to the destination.

`Status`

The status of the publishing destination.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GuardDuty::Member

CFNDestinationProperties
