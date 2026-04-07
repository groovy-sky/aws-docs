This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::AccessPoint

Access points simplify managing data access at scale for shared datasets in Amazon S3.
Access points are unique hostnames you create to enforce distinct permissions and network controls
for all requests made through an access point. You can create hundreds of access points per bucket,
each with a distinct name and permissions customized for each application. Each access point works
in conjunction with the bucket policy that is attached to the underlying bucket. For more information, see [Managing access to shared datasets in directory buckets with access points](../../../s3/latest/userguide/access-points-directory-buckets.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Express::AccessPoint",
  "Properties" : {
      "Bucket" : String,
      "BucketAccountId" : String,
      "Name" : String,
      "Policy" : Json,
      "PublicAccessBlockConfiguration" : PublicAccessBlockConfiguration,
      "Scope" : Scope,
      "Tags" : [ Tag, ... ],
      "VpcConfiguration" : VpcConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::S3Express::AccessPoint
Properties:
  Bucket: String
  BucketAccountId: String
  Name: String
  Policy: Json
  PublicAccessBlockConfiguration:
    PublicAccessBlockConfiguration
  Scope:
    Scope
  Tags:
    - Tag
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`Bucket`

The name of the bucket that you want to associate the access point with.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BucketAccountId`

The AWS account ID that owns the bucket associated with this access point.

_Required_: No

_Type_: String

_Pattern_: `^\d{12}$`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

An access point name consists of a base name you provide, followed by the zoneID (AWS Local Zone) followed by the prefix `--xa-s3`. For example, accesspointname--zoneID--xa-s3.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The access point policy associated with the specified access point.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicAccessBlockConfiguration`

Public access is blocked by default to access points for directory buckets.

_Required_: No

_Type_: [PublicAccessBlockConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3express-accesspoint-publicaccessblockconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

You can use the access point scope to restrict access to specific prefixes, API operations, or a combination of both.

For more information, see [Manage the scope of your access points for directory buckets.](../../../s3/latest/userguide/access-points-directory-buckets-manage-scope.md)

_Required_: No

_Type_: [Scope](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3express-accesspoint-scope.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of tags that you can apply to access points. Tags are key-value pairs of metadata used to categorize your access points and control access. For more information, see [Using tags for attribute-based access control (ABAC)](../../../s3/latest/userguide/tagging.md#using-tags-for-abac).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3express-accesspoint-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

If you include this field, Amazon S3 restricts access to this access point to requests from the specified virtual private cloud (VPC).

_Required_: No

_Type_: [VpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3express-accesspoint-vpcconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the access point.

`NetworkOrigin`

The network configuration of the access point.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon S3 Express

PublicAccessBlockConfiguration
