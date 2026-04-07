This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ProductSubscription

The `AWS::SecurityHub::ProductSubscription` resource creates a subscription to a third-party product
that generates findings that you want to receive in AWS Security Hub CSPM. For a list of integrations to third-party
products, see [Available third-party partner product integrations](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-partner-providers.html)
in the _AWS Security Hub CSPM User Guide_.

To change a product subscription, remove the current product subscription resource, and then create a new one.

Tags aren't supported for this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::ProductSubscription",
  "Properties" : {
      "ProductArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::ProductSubscription
Properties:
  ProductArn: String

```

## Properties

`ProductArn`

The ARN of the product to enable the integration for.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws\S*:securityhub:\S*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of your subscription to the product to enable integrations
for. For example, `arn:aws:securityhub:us-east-1:123456789012:product-subscription/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProductSubscriptionArn`

The ARN of your subscription to the product to enable integrations for.

## Examples

### Creating a Security Hub CSPM product subscription

The following example creates a Security Hub CSPM product subscription to the specified third-party product.

#### JSON

```json

{
	"Description": "Example template to create a Security Hub product subscription",
	"Resources": {
		"SecurityHubProductSubscription": {
			"Type": "AWS::SecurityHub::ProductSubscription",
			"Properties": {
				"ProductArn": {
					"Fn::Sub": "arn:${AWS::Partition}:securityhub:${AWS::Region}::product/aws/example"
				}
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to create a Security Hub product subscription
Resources:
  SecurityHubProductSubscription:
    Type: 'AWS::SecurityHub::ProductSubscription'
    Properties:
      ProductArn: !Sub 'arn:${AWS::Partition}:securityhub:${AWS::Region}::product/aws/example'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::PolicyAssociation

AWS::SecurityHub::SecurityControl
