This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationS3

The `AWS::DataSync::LocationS3` resource specifies an endpoint for an
Amazon S3 bucket.

For more information, see the [_AWS DataSync User Guide_](../../../datasync/latest/userguide/create-s3-location.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationS3",
  "Properties" : {
      "S3BucketArn" : String,
      "S3Config" : S3Config,
      "S3StorageClass" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationS3
Properties:
  S3BucketArn: String
  S3Config:
    S3Config
  S3StorageClass: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`S3BucketArn`

The ARN of the Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):s3:[a-z\-0-9]*:[0-9]*:.*$`

_Maximum_: `156`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3Config`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that is used
to access an Amazon S3 bucket.

For detailed information about using such a role, see [Creating a Location for Amazon S3](../../../datasync/latest/userguide/working-with-locations.md#create-s3-location) in the _AWS DataSync_
_User Guide_.

_Required_: Yes

_Type_: [S3Config](aws-properties-datasync-locations3-s3config.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3StorageClass`

The Amazon S3 storage class that you want to store your files in when this location is
used as a task destination. For buckets in AWS Regions, the storage
class defaults to S3 Standard.

For more information about S3 storage classes, see [Amazon S3 Storage Classes](https://aws.amazon.com/s3/storage-classes). Some storage classes
have behaviors that can affect your S3 storage costs. For detailed information, see
[Considerations When Working with Amazon S3 Storage Classes in\
DataSync](../../../datasync/latest/userguide/create-s3-location.md#using-storage-classes).

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | STANDARD_IA | ONEZONE_IA | INTELLIGENT_TIERING | GLACIER | GLACIER_INSTANT_RETRIEVAL | DEEP_ARCHIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies a prefix in the S3 bucket that DataSync reads from or writes to
(depending on whether the bucket is a source or destination location).

###### Note

DataSync can't transfer objects with a prefix that begins with a slash ( `/`)
or includes `//`, `/./`, or `/../` patterns. For
example:

- `/photos`

- `photos//2006/January`

- `photos/./2006/February`

- `photos/../2006/March`

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{M}\p{Z}\p{S}\p{N}\p{P}\p{C}]*$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS resources. We recommend creating at least a name tag for your transfer location.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locations3-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource Amazon Resource Name (ARN). For
example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`LocationArn`

The Amazon Resource Name (ARN) of the specified Amazon S3 location.

`LocationUri`

The URI of the specified Amazon S3 location.

## Examples

### Create an Amazon S3 location for DataSync

The following example specifies an S3 location for DataSync. In this example,
the S3 location uses the bucket specified by the Amazon Resource Name (ARN)
`arn:aws:s3:::amzn-s3-demo-bucket`. The access
role is specified by the ARN `arn:aws:iam::111222333444:role/amzn-s3-demo-bucket-access-role`. The example uses the S3
Standard storage class and refers to the subdirectory `/MyFolder`.

AWS Security Token Service (AWS STS) must be activated in your
account and Region for DataSync to assume the IAM role. For more information
about temporary security credentials, see [Temporary security\
credentials in IAM](../../../iam/latest/userguide/id-credentials-temp.md) in the _IAM User Guide_.

#### JSON

```json

{
"AWSTemplateFormatVersion": "2010-09-09",
"Description": "Specifies an S3 location for DataSync",
"Resources":
  {
    "LocationS3": {
      "Type": "AWS::DataSync::LocationS3",
      "Properties": {
        "S3BucketArn": "arn:aws:s3:::amzn-s3-demo-bucket",
        "S3Config": {
          "BucketAccessRoleArn": "arn:aws:iam::111222333444:role/amzn-s3-demo-bucket-access-role"
        },
      "S3StorageClass": "STANDARD",
      "Subdirectory": "/MyFolder"
      }
    }
  }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an S3 location for DataSync
Resources:
  LocationS3:
    Type: AWS::DataSync::LocationS3
    Properties:
      S3BucketArn: arn:aws:s3:::amzn-s3-demo-bucket
      S3Config:
        BucketAccessRoleArn: arn:aws:iam::111222333444:role/amzn-s3-demo-bucket-access-role
      S3StorageClass: STANDARD
      Subdirectory: /MyFolder

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

S3Config

All content copied from https://docs.aws.amazon.com/.
