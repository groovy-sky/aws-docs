This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::MultiRegionAccessPoint

The `AWS::S3::MultiRegionAccessPoint` resource creates an Amazon S3
Multi-Region Access Point. To learn more about Multi-Region Access Points, see [Multi-Region Access Points in Amazon S3](../../../s3/latest/userguide/multiregionaccesspoints.md) in the in the _Amazon S3 User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::MultiRegionAccessPoint",
  "Properties" : {
      "Name" : String,
      "PublicAccessBlockConfiguration" : PublicAccessBlockConfiguration,
      "Regions" : [ Region, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::S3::MultiRegionAccessPoint
Properties:
  Name: String
  PublicAccessBlockConfiguration:
    PublicAccessBlockConfiguration
  Regions:
    - Region

```

## Properties

`Name`

The name of the Multi-Region Access Point.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicAccessBlockConfiguration`

The PublicAccessBlock configuration that you want to apply to this Multi-Region Access
Point. You can enable the configuration options in any combination. For more information about
when Amazon S3 considers an object public, see [The Meaning of "Public"](../../../s3/latest/dev/access-control-block-public-access.md#access-control-block-public-access-policy-status) in the _Amazon S3 User Guide_.

_Required_: No

_Type_: [PublicAccessBlockConfiguration](aws-properties-s3-multiregionaccesspoint-publicaccessblockconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Regions`

A collection of the Regions and buckets associated with the Multi-Region Access
Point.

_Required_: Yes

_Type_: Array of [Region](aws-properties-s3-multiregionaccesspoint-region.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the Multi-Region Access Point.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Alias`

The alias for the Multi-Region Access Point. For more information about the distinction
between the name and the alias of an Multi-Region Access Point, see [Managing Multi-Region Access Points](../../../s3/latest/userguide/creatingmultiregionaccesspoints.md#multi-region-access-point-naming) in the _Amazon S3 User_
_Guide_.

`CreatedAt`

The timestamp of when the Multi-Region Access Point is created.

## Examples

You can use AWSCloudFormation to create a Multi-Region Access
Point. When you create the Multi-Region Access Point, you must provide all the S3 buckets that
it supports. Be aware that you can't add any S3 buckets to the Multi-Region Access Point after
it's been created.

### Multi-Region Access Point with two Regions

The following template can be used to create a Multi-Region Access Point (with two Regions) through AWS CloudFormation.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "DOC-EXAMPLE-MULTI-REGION-ACCESS-POINT": {
          "Type" : "AWS::S3::MultiRegionAccessPoint",
          "Properties" : {
            "PublicAccessBlockConfiguration" : {
              "BlockPublicAcls" : "True",
              "BlockPublicPolicy" : "True",
              "IgnorePublicAcls" : "True",
              "RestrictPublicBuckets" : "True"
            },
            "Regions" : [ {"Bucket":"DOC-EXAMPLE-BUCKET1"}, {"Bucket": "DOC-EXAMPLE-BUCKET2"} ]
        }
      }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Resources:
  DOC-EXAMPLE-MULTI-REGION-ACCESS-POINT:
    Type: AWS::S3::MultiRegionAccessPoint
    Properties:
      PublicAccessBlockConfiguration:
        BlockPublicAcls: "True"
        BlockPublicPolicy: "True"
        IgnorePublicAcls: "True"
        RestrictPublicBuckets: "True"
      Regions:
        - Bucket: DOC-EXAMPLE-BUCKET1
        - Bucket: DOC-EXAMPLE-BUCKET2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::S3::BucketPolicy

PublicAccessBlockConfiguration

All content copied from https://docs.aws.amazon.com/.
