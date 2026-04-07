This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::TrustedEntitySet

Creates a new trusted entity set. In the trusted entity set, you can provide IP addresses
and domains that you believe are secure for communication in your AWS environment. GuardDuty
will not generate findings for the entries that are specified in a trusted entity set. At any
given time, you can have only one trusted entity set.

Only users of the administrator account can manage the entity sets, which automatically
apply to member accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::TrustedEntitySet",
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

Type: AWS::GuardDuty::TrustedEntitySet
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

The unique regional detector ID of the GuardDuty account for which you want to
create a trusted entity set.

To find the `detectorId` in the current Region, see the Settings page in the
GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html)
API.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExpectedBucketOwner`

The AWS account ID that owns the Amazon S3 bucket specified in the _Location_ field.

Whether or not you provide the account ID for this optional field, GuardDuty
validates that the account ID associated with the `DetectorId` value owns the S3 bucket in the `Location`
field. If GuardDuty finds that this S3 bucket doesn't belong to the specified account ID, you will
get an error at the time of activating this list.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Format`

The format of the file that contains the trusted entity set. For information about
supported formats, see [List formats](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#prepare_list)
in the _Amazon GuardDuty User Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Location`

The URI of the file that contains the trusted entity set.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A user-friendly name to identify the trusted entity set. Valid characters include
lowercase letters, uppercase letters, numbers, dash(-), and underscore (\_).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be added to a new trusted entity set resource. Each tag consists of a key
and an optional value, both of which you define.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [TagItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-trustedentityset-tagitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID of the `TrustedEntitySet`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the trusted entity set was created.

`ErrorDetails`

Specifies the error details when the status of the trusted entity set shows as **Error**.

`Status`

The status of your `TrustedEntitySet`. For information about valid status
values, see [Understanding list statuses](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_upload-lists.html#guardduty-entity-list-statuses) in the _Amazon GuardDuty User Guide_.

`UpdatedAt`

The timestamp when the trusted entity set was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagItem

TagItem
