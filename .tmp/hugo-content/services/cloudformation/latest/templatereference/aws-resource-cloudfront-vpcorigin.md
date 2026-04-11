This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::VpcOrigin

An Amazon CloudFront VPC origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::VpcOrigin",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "VpcOriginEndpointConfig" : VpcOriginEndpointConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::VpcOrigin
Properties:
  Tags:
    - Tag
  VpcOriginEndpointConfig:
    VpcOriginEndpointConfig

```

## Properties

`Tags`

A complex type that contains zero or more `Tag` elements.

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudfront-vpcorigin-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcOriginEndpointConfig`

The VPC origin endpoint configuration.

_Required_: Yes

_Type_: [VpcOriginEndpointConfig](aws-properties-cloudfront-vpcorigin-vpcoriginendpointconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AccountId`

The account ID of the AWS account that owns the VPC origin.

`Arn`

The VPC origin ARN.

`CreatedTime`

The VPC origin created time.

`Id`

The VPC origin ID.

`LastModifiedTime`

The VPC origin last modified time.

`Status`

The VPC origin status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
