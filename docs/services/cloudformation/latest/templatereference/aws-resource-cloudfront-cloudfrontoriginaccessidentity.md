This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::CloudFrontOriginAccessIdentity

The request to create a new origin access identity (OAI). An origin access identity is
a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all
or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-s3.md) in
the _Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::CloudFrontOriginAccessIdentity",
  "Properties" : {
      "CloudFrontOriginAccessIdentityConfig" : CloudFrontOriginAccessIdentityConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::CloudFrontOriginAccessIdentity
Properties:
  CloudFrontOriginAccessIdentityConfig:
    CloudFrontOriginAccessIdentityConfig

```

## Properties

`CloudFrontOriginAccessIdentityConfig`

The current configuration information for the identity.

_Required_: Yes

_Type_: [CloudFrontOriginAccessIdentityConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the origin access identity, such as
`E15MNIMTCFKK4C`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID for the origin access identity, for example, `E74FTE3AJFJ256A`.

`S3CanonicalUserId`

The Amazon S3 canonical user ID for the origin access identity, used when
giving the origin access identity read permission to an object in Amazon S3. For
example:
`b970b42360b81c8ddbd79d2f5df0069ba9033c8a79655752abe380cd6d63ba8bcf23384d568fcf89fc49700b5e11a0fd`.

## Examples

### Specify the comment for an origin access identity

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "cloudfrontoriginaccessidentity": {
            "Type": "AWS::CloudFront::CloudFrontOriginAccessIdentity",
            "Properties": {
                "CloudFrontOriginAccessIdentityConfig": {
                    "Comment": "string-value"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  cloudfrontoriginaccessidentity:
    Type: AWS::CloudFront::CloudFrontOriginAccessIdentity
    Properties:
      CloudFrontOriginAccessIdentityConfig:
        Comment: string-value
```

## See also

- [OriginAccessIdentity](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_S3OriginConfig.html#cloudfront-Type-S3OriginConfig-OriginAccessIdentity) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryStringsConfig

CloudFrontOriginAccessIdentityConfig
