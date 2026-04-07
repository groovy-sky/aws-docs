This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::TrustStore

A trust store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::TrustStore",
  "Properties" : {
      "CaCertificatesBundleSource" : CaCertificatesBundleSource,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::TrustStore
Properties:
  CaCertificatesBundleSource:
    CaCertificatesBundleSource
  Name: String
  Tags:
    - Tag

```

## Properties

`CaCertificatesBundleSource`

A CA certificates bundle source.

_Required_: No

_Type_: [CaCertificatesBundleSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-truststore-cacertificatesbundlesource.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The trust store's name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-truststore-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The trust store's Amazon Resource Name (ARN).

`ETag`

The version identifier for the current version of the trust store.

`Id`

The trust store's ID.

`LastModifiedTime`

The trust store's last modified time.

`NumberOfCaCertificates`

The trust store's number of CA certificates.

`Status`

The trust store's status.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrustedSigners

CaCertificatesBundleS3Location
