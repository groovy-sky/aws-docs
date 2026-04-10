This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ConsumableResource

The `AWS::Batch::ConsumableResource` resource specifies the parameters for an
AWS Batch consumable resource. For more information, see [Resource-aware scheduling](../../../batch/latest/userguide/resource-aware-scheduling.md) in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Batch::ConsumableResource",
  "Properties" : {
      "ConsumableResourceName" : String,
      "ResourceType" : String,
      "Tags" : {Key: Value, ...},
      "TotalQuantity" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Batch::ConsumableResource
Properties:
  ConsumableResourceName: String
  ResourceType: String
  Tags:
    Key: Value
  TotalQuantity: Integer

```

## Properties

`ConsumableResourceName`

The name of the consumable resource.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

Indicates whether the resource is available to be re-used after a job completes. Can be
one of:

- `REPLENISHABLE`

- `NON_REPLENISHABLE`

_Required_: Yes

_Type_: String

_Allowed values_: `REPLENISHABLE | NON_REPLENISHABLE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags that you apply to the consumable resource to help you categorize and organize your
resources. Each tag consists of a key and an optional value. For more information, see [Tagging your AWS Batch resources](../../../batch/latest/userguide/using-tags.md).

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TotalQuantity`

The total amount of the consumable resource that is available.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the consumable resource ARN, such as `arn:aws:batch:us-east-1:555555555555:consumable-resource/Resource`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AvailableQuantity`

The amount of the consumable resource that is currently available to use.

`ConsumableResourceArn`

The Amazon Resource Name (ARN) of the consumable resource.

`CreatedAt`

The Unix timestamp (in milliseconds) for when the consumable resource was created.

`InUseQuantity`

The amount of the consumable resource that is currently in use.

## See also

- [Resource-aware scheduling](../../../batch/latest/userguide/resource-aware-scheduling.md) in the _AWS Batch User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdatePolicy

AWS::Batch::JobDefinition

All content copied from https://docs.aws.amazon.com/.
