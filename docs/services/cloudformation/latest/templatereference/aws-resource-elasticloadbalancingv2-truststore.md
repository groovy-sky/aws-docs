This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::TrustStore

Creates a trust store. You must specify `CaCertificatesBundleS3Bucket` and
`CaCertificatesBundleS3Key`. When you create a trust store, you must specify
`Name`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::TrustStore",
  "Properties" : {
      "CaCertificatesBundleS3Bucket" : String,
      "CaCertificatesBundleS3Key" : String,
      "CaCertificatesBundleS3ObjectVersion" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::TrustStore
Properties:
  CaCertificatesBundleS3Bucket: String
  CaCertificatesBundleS3Key: String
  CaCertificatesBundleS3ObjectVersion: String
  Name: String
  Tags:
    - Tag

```

## Properties

`CaCertificatesBundleS3Bucket`

The Amazon S3 bucket for the ca certificates bundle.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaCertificatesBundleS3Key`

The Amazon S3 path for the ca certificates bundle.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaCertificatesBundleS3ObjectVersion`

The Amazon S3 object version for the ca certificates bundle. If undefined the current version is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the trust store.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9]+-)*[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to assign to the trust store.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-truststore-tag.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the trust store.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`NumberOfCaCertificates`

The number of ca certificates in the trust store.

`Status`

The status of the trust store. The possible values are `CREATING` and `ACTIVE`.

`TrustStoreArn`

The Amazon Resource Name (ARN) of the trust store.

## Examples

The following example creates a trust store.

#### YAML

```yaml

myTrustStore:
  Type: 'AWS::ElasticLoadBalancingV2::TrustStore'
  Properties:
    Name: my-trust-store
    CaCertificatesBundleS3Bucket: amzn-s3-demo-bucket
    CaCertificatesBundleS3Key: certificates/ca-bundle.pem
```

#### JSON

```json

{
  "myTrustStore": {
    "Type": "AWS::ElasticLoadBalancingV2::TrustStore",
    "Properties": {
      "Name": "my-trust-store",
      "CaCertificatesBundleS3Bucket": "amzn-s3-demo-bucket",
      "CaCertificatesBundleS3Key": "certificates/ca-bundle.pem"
    }
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetGroupAttribute

Tag
