This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::OriginAccessControl

Creates a new origin access control in CloudFront. After you create an origin access
control, you can add it to an origin in a CloudFront distribution so that CloudFront sends
authenticated (signed) requests to the origin.

This makes it possible to block public access to the origin, allowing viewers (users) to
access the origin's content only through CloudFront.

For more information about using a CloudFront origin access control, see [Restricting access to an AWS origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html) in the
_Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::OriginAccessControl",
  "Properties" : {
      "OriginAccessControlConfig" : OriginAccessControlConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::OriginAccessControl
Properties:
  OriginAccessControlConfig:
    OriginAccessControlConfig

```

## Properties

`OriginAccessControlConfig`

The origin access control.

_Required_: Yes

_Type_: [OriginAccessControlConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-originaccesscontrol-originaccesscontrolconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique identifier of the origin access control.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RealtimeMetricsSubscriptionConfig

OriginAccessControlConfig
