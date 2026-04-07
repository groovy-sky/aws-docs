This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::MeteredProduct

Adds a metered product.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Deadline::MeteredProduct",
  "Properties" : {
      "LicenseEndpointId" : String,
      "ProductId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Deadline::MeteredProduct
Properties:
  LicenseEndpointId: String
  ProductId: String

```

## Properties

`LicenseEndpointId`

The Amazon EC2 identifier of the license endpoint.

_Required_: No

_Type_: String

_Pattern_: `^le-[0-9a-f]{32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProductId`

The product ID.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z]{1,32}-[.0-9a-z]{1,32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the metered product.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the metered product.

`Family`

The family to which the metered product belongs.

`Port`

The port on which the metered product should run.

`Vendor`

The vendor.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Deadline::Limit

AWS::Deadline::Monitor
