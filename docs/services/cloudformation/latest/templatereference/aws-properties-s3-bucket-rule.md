This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket Rule

Specifies lifecycle rules for an Amazon S3 bucket. For more information, see [Put Bucket\
Lifecycle Configuration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) in the _Amazon S3 API Reference_.

You must specify at least one of the following properties:
`AbortIncompleteMultipartUpload`, `ExpirationDate`,
`ExpirationInDays`, `NoncurrentVersionExpirationInDays`,
`NoncurrentVersionTransition`, `NoncurrentVersionTransitions`,
`Transition`, or `Transitions`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AbortIncompleteMultipartUpload" : AbortIncompleteMultipartUpload,
  "ExpirationDate" : String,
  "ExpirationInDays" : Integer,
  "ExpiredObjectDeleteMarker" : Boolean,
  "Id" : String,
  "NoncurrentVersionExpiration" : NoncurrentVersionExpiration,
  "NoncurrentVersionExpirationInDays" : Integer,
  "NoncurrentVersionTransition" : NoncurrentVersionTransition,
  "NoncurrentVersionTransitions" : [ NoncurrentVersionTransition, ... ],
  "ObjectSizeGreaterThan" : String,
  "ObjectSizeLessThan" : String,
  "Prefix" : String,
  "Status" : String,
  "TagFilters" : [ TagFilter, ... ],
  "Transition" : Transition,
  "Transitions" : [ Transition, ... ]
}

```

### YAML

```yaml

  AbortIncompleteMultipartUpload:
    AbortIncompleteMultipartUpload
  ExpirationDate: String
  ExpirationInDays: Integer
  ExpiredObjectDeleteMarker: Boolean
  Id: String
  NoncurrentVersionExpiration:
    NoncurrentVersionExpiration
  NoncurrentVersionExpirationInDays: Integer
  NoncurrentVersionTransition:
    NoncurrentVersionTransition
  NoncurrentVersionTransitions:
    - NoncurrentVersionTransition
  ObjectSizeGreaterThan: String
  ObjectSizeLessThan: String
  Prefix: String
  Status: String
  TagFilters:
    - TagFilter
  Transition:
    Transition
  Transitions:
    - Transition

```

## Properties

`AbortIncompleteMultipartUpload`

Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3
bucket.

_Required_: Conditional

_Type_: [AbortIncompleteMultipartUpload](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-abortincompletemultipartupload.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationDate`

Indicates when objects are deleted from Amazon S3 and Amazon S3 Glacier. The date value
must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and
transition time, you must use the same time unit for both properties (either in days or by
date). The expiration time must also be later than the transition time.

_Required_: Conditional

_Type_: String

_Pattern_: `^(\d{4})-(0[0-9]|1[0-2])-([0-2]\d|3[01])T([01]\d|2[0-4]):([0-5]\d):([0-6]\d)((\.\d{3})?)Z$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpirationInDays`

Indicates the number of days after creation when objects are deleted from Amazon S3 and
Amazon S3 Glacier. If you specify an expiration and transition time, you must use the same
time unit for both properties (either in days or by date). The expiration time must also be
later than the transition time.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpiredObjectDeleteMarker`

Indicates whether Amazon S3 will remove a delete marker without any noncurrent versions.
If set to true, the delete marker will be removed if there are no noncurrent versions. This
cannot be specified with `ExpirationInDays`, `ExpirationDate`, or
`TagFilters`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

Unique identifier for the rule. The value can't be longer than 255 characters.

_Required_: No

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoncurrentVersionExpiration`

Specifies when noncurrent object versions expire. Upon expiration, Amazon S3
permanently deletes the noncurrent object versions. You set this lifecycle configuration
action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's
lifetime.

_Required_: No

_Type_: [NoncurrentVersionExpiration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-noncurrentversionexpiration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoncurrentVersionExpirationInDays`

(Deprecated.) For buckets with versioning enabled (or suspended), specifies the time, in
days, between when a new version of the object is uploaded to the bucket and when old versions
of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you
specify a transition and expiration time, the expiration time must be later than the
transition time.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoncurrentVersionTransition`

(Deprecated.) For buckets with versioning enabled (or suspended), specifies when
non-current objects transition to a specified storage class. If you specify a transition and
expiration time, the expiration time must be later than the transition time. If you specify
this property, don't specify the `NoncurrentVersionTransitions` property.

_Required_: Conditional

_Type_: [NoncurrentVersionTransition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-noncurrentversiontransition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoncurrentVersionTransitions`

For buckets with versioning enabled (or suspended), one or more transition rules that
specify when non-current objects transition to a specified storage class. If you specify a
transition and expiration time, the expiration time must be later than the transition time. If
you specify this property, don't specify the `NoncurrentVersionTransition`
property.

_Required_: Conditional

_Type_: Array of [NoncurrentVersionTransition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-noncurrentversiontransition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectSizeGreaterThan`

Specifies the minimum object size in bytes for this rule to apply to. Objects must be
larger than this value in bytes. For more information about size based rules, see [Lifecycle configuration using size-based rules](../../../s3/latest/userguide/lifecycle-configuration-examples.md#lc-size-rules) in the _Amazon S3 User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectSizeLessThan`

Specifies the maximum object size in bytes for this rule to apply to. Objects must be
smaller than this value in bytes. For more information about sized based rules, see [Lifecycle configuration using size-based rules](../../../s3/latest/userguide/lifecycle-configuration-examples.md#lc-size-rules) in the _Amazon S3 User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `[0-9]+`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Prefix`

Object key prefix that identifies one or more objects to which this rule applies.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../../../s3/latest/userguide/object-keys.md#object-key-xml-related-constraints).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

If `Enabled`, the rule is currently being applied. If `Disabled`, the rule is
not currently being applied.

_Required_: Yes

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagFilters`

Tags to use to identify a subset of objects to which the lifecycle rule applies.

_Required_: No

_Type_: Array of [TagFilter](aws-properties-s3-bucket-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transition`

(Deprecated.) Specifies when an object transitions to a specified storage class. If you
specify an expiration and transition time, you must use the same time unit for both properties
(either in days or by date). The expiration time must also be later than the transition time.
If you specify this property, don't specify the `Transitions` property.

_Required_: Conditional

_Type_: [Transition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-transition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transitions`

One or more transition rules that specify when an object transitions to a specified
storage class. If you specify an expiration and transition time, you must use the same time
unit for both properties (either in days or by date). The expiration time must also be later
than the transition time. If you specify this property, don't specify the
`Transition` property.

_Required_: Conditional

_Type_: Array of [Transition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3-bucket-transition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Manage the lifecycle for S3 objects

The following example template shows an S3 bucket with a lifecycle configuration rule.
The rule applies to all objects with the `glacier` key prefix. The objects are
transitioned to Glacier after one day, and deleted after one year.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "Private",
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "Id": "GlacierRule",
                            "Prefix": "glacier",
                            "Status": "Enabled",
                            "ExpirationInDays": 365,
                            "Transitions": [
                                {
                                    "TransitionInDays": 1,
                                    "StorageClass": "GLACIER"
                                }
                            ]
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a lifecycle configuration."
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      AccessControl: Private
      LifecycleConfiguration:
        Rules:
          - Id: GlacierRule
            Prefix: glacier
            Status: Enabled
            ExpirationInDays: 365
            Transitions:
              - TransitionInDays: 1
                StorageClass: GLACIER
Outputs:
  BucketName:
    Value: !Ref S3Bucket
    Description: Name of the sample Amazon S3 bucket with a lifecycle configuration.
```

## See also

- AWS::S3::Bucket [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RoutingRuleCondition

S3KeyFilter
