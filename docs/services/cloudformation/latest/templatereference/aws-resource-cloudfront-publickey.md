This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::PublicKey

A public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](../../../amazoncloudfront/latest/developerguide/field-level-encryption.md).

CloudFront supports signed URLs and signed cookies with RSA 2048 or ECDSA 256 key signatures. Field-level encryption is only compatible with RSA 2048 key signatures.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::PublicKey",
  "Properties" : {
      "PublicKeyConfig" : PublicKeyConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::PublicKey
Properties:
  PublicKeyConfig:
    PublicKeyConfig

```

## Properties

`PublicKeyConfig`

Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](../../../amazoncloudfront/latest/developerguide/field-level-encryption.md).

_Required_: Yes

_Type_: [PublicKeyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-publickey-publickeyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the public key. For example:
`K36X4X2EO997HM`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedTime`

The date and time when the public key was uploaded.

`Id`

The identifier of the public key.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryStringsConfig

PublicKeyConfig
