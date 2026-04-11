This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::BillingGroup

Creates a new billing group.

Requires permission to access the [CreateBillingGroup](../../../service-authorization/latest/reference/list-awsiot.md#awsiot-actions-as-permissions) action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::BillingGroup",
  "Properties" : {
      "BillingGroupName" : String,
      "BillingGroupProperties" : BillingGroupProperties,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::BillingGroup
Properties:
  BillingGroupName: String
  BillingGroupProperties:
    BillingGroupProperties
  Tags:
    - Tag

```

## Properties

`BillingGroupName`

The name of the billing group.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BillingGroupProperties`

The properties of the billing group.

_Required_: No

_Type_: [BillingGroupProperties](aws-properties-iot-billinggroup-billinggroupproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the billing group.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-billinggroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the billing group id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the billing group.

`Id`

The ID of the billing group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

BillingGroupProperties

All content copied from https://docs.aws.amazon.com/.
