This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Bucket

The `AWS::Lightsail::Bucket` resource specifies a bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Bucket",
  "Properties" : {
      "AccessRules" : AccessRules,
      "BucketName" : String,
      "BundleId" : String,
      "ObjectVersioning" : Boolean,
      "ReadOnlyAccessAccounts" : [ String, ... ],
      "ResourcesReceivingAccess" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Bucket
Properties:
  AccessRules:
    AccessRules
  BucketName: String
  BundleId: String
  ObjectVersioning: Boolean
  ReadOnlyAccessAccounts:
    - String
  ResourcesReceivingAccess:
    - String
  Tags:
    - Tag

```

## Properties

`AccessRules`

An object that describes the access rules for the bucket.

_Required_: No

_Type_: [AccessRules](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-bucket-accessrules.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

The name of the bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][a-z0-9-]{1,52}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `54`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BundleId`

The bundle ID for the bucket (for example, `small_1_0`).

A bucket bundle specifies the monthly cost, storage space, and data transfer quota for a
bucket.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectVersioning`

Indicates whether object versioning is enabled for the bucket.

The following options can be configured:

- `Enabled` \- Object versioning is enabled.

- `Suspended` \- Object versioning was previously enabled but is currently
suspended. Existing object versions are retained.

- `NeverEnabled` \- Object versioning has never been enabled.

_Required_: No

_Type_: Boolean

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnlyAccessAccounts`

An array of AWS account IDs that have read-only access to the
bucket.

_Required_: No

_Type_: Array of String

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcesReceivingAccess`

An array of Lightsail instances that have access to the bucket.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-bucket-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AbleToUpdateBundle`

A Boolean value indicating whether the bundle that is currently applied to your
distribution can be changed to another bundle.

`BucketArn`

The Amazon Resource Name (ARN) of the bucket.

`Url`

The URL of the bucket.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Lightsail::Alarm

AccessRules
