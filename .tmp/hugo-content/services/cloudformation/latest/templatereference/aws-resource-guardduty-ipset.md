This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::IPSet

The `AWS::GuardDuty::IPSet` resource helps you create a list of trusted IP
addresses that you can use for secure communication with AWS
infrastructure and applications. Once you activate this list, GuardDuty will not
generate findings when there is an activity associated with these safe IP addresses.

Only the users of the GuardDuty administrator account can manage this list.
These settings are also applied to the member accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::IPSet",
  "Properties" : {
      "Activate" : Boolean,
      "DetectorId" : String,
      "ExpectedBucketOwner" : String,
      "Format" : String,
      "Location" : String,
      "Name" : String,
      "Tags" : [ TagItem, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::IPSet
Properties:
  Activate: Boolean
  DetectorId: String
  ExpectedBucketOwner: String
  Format: String
  Location: String
  Name: String
  Tags:
    - TagItem

```

## Properties

`Activate`

A boolean value that determines if GuardDuty can start using this list for
custom threat detection. For GuardDuty to prevent generating findings based on an
activity associated with these entries, this list must be active.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The unique ID of the detector of the GuardDuty account for which you want to create an IPSet.

To find the `detectorId` in the current Region, see the
Settings page in the GuardDuty console, or run the [ListDetectors](../../../../reference/guardduty/latest/apireference/api-listdetectors.md) API.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExpectedBucketOwner`

The AWS account ID that owns the Amazon S3 bucket specified in
the _Location_ field.

When you provide this account ID, GuardDuty will validate that the S3 bucket
belongs to this account. If you don't specify an account ID owner, GuardDuty
doesn't perform any validation.

_Required_: No

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The format of the file that contains the IPSet. For information about supported formats,
see [List formats](../../../guardduty/latest/ug/guardduty-upload-lists.md#prepare_list) in the _Amazon GuardDuty User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `TXT | STIX | OTX_CSV | ALIEN_VAULT | PROOF_POINT | FIRE_EYE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

The URI of the file that contains the IPSet.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The user-friendly name to identify the IPSet.

The name of your list must be unique within an AWS account and
Region. Valid characters are alphanumeric, whitespace, dash (-), and underscores (\_).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be added to a new threat entity set resource. Each tag consists of a key and
an optional value, both of which you define.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [TagItem](aws-properties-guardduty-ipset-tagitem.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID of the `IPSet`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

### Declare an IPSet Resource

The following example shows how to declare a GuardDuty `IPSet` resource:

#### JSON

```json

"myipset": {
    "Type" : "AWS::GuardDuty::IPSet",
    "Properties" : {
        "Activate" : True,
        "DetectorId" : "12abc34d567e8f4912ab3d45e67891f2",
        "ExpectedBucketOwner" : "111122223333",
        "Format" : "TXT",
        "Location" : "https://s3-us-west-2.amazonaws.com/amzn-s3-demo-bucket/myipset.txt",
        "Name" : "MyIPSet"
    }
}
```

#### YAML

```yaml

myipset:
    Type: AWS::GuardDuty::IPSet
    Properties:
        Activate: True
        DetectorId: "12abc34d567e8f4912ab3d45e67891f2"
        ExpectedBucketOwner : "111122223333"
        Format: "TXT"
        Location: "https://s3-us-west-2.amazonaws.com/amzn-s3-demo-bucket/myipset.txt"
        Name: "MyIPSet"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagItem

TagItem

All content copied from https://docs.aws.amazon.com/.
