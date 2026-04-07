This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::ThreatEntitySet

The `AWS::GuardDuty::ThreatEntitySet` resource helps you create a list of
known malicious IP addresses and domain names in your AWS environment. Once
you activate this list, GuardDuty will use the entries in this list as an
additional source of threat detection and generate findings when there is an activity
associated with these known malicious IP addresses and domain names. GuardDuty
continues to monitor independently of this custom threat entity set.

Only the users of the GuardDuty administrator account can manage this list.
These settings automatically apply to the member accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::ThreatEntitySet",
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

Type: AWS::GuardDuty::ThreatEntitySet
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
custom threat detection. For GuardDuty to consider the entries in this list and
generate findings based on associated activity, this list must be active.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The unique regional detector ID of the GuardDuty account for which you want to
create a threat entity set.

To find the `detectorId` in the current Region, see the Settings page in the
GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html)
API.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExpectedBucketOwner`

The AWS account ID that owns the Amazon S3 bucket specified in the
_Location_ field.

Whether or not you provide the account ID for this optional field, GuardDuty
validates that the account ID associated with the `DetectorId` owns the S3 bucket in the `Location`
field. If GuardDuty finds that this S3 bucket doesn't belong to the specified account ID, you will
get an error at the time of activating this list.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The format of the file that contains the threat entity set. For information about
supported formats, see [List formats](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#prepare_list)
in the _Amazon GuardDuty User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

The URI of the file that contains the threat entity set.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The user-friendly name to identify the threat entity set. Valid characters are
alphanumeric, whitespace, dash (-), and underscores (\_).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be added to a new threat entity set resource. Each tag consists of a key and
an optional value, both of which you define.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [TagItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-threatentityset-tagitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID associated with the newly created threat entity
set.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the threat entity set was created.

`ErrorDetails`

The details associated with the **Error** status of your
threat entity list.

`Id`

Returns the unique ID associated with the newly created threat entity set.

`Status`

The status of your `ThreatEntitySet`. For information about valid status
values, see [Understanding list statuses](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#guardduty-entity-list-statuses) in the _Amazon GuardDuty User Guide_.

`UpdatedAt`

The timestamp when the threat entity set was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagItem

TagItem
