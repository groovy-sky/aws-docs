This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::IntegrationResourceProperty

The `AWS::Glue::IntegrationResourceProperty` resource type can be used to setup `ResourceProperty` of the AWS Glue connection (for the SaaS source), DynamoDB Database (for DynamoDB source), or AWS Glue database ARN (for the target). ResourceProperty is used to define the properties requires to setup the integration, including the role to access the connection or database, KMS keys, event bus for event notifications and VPC connection. To set both source and target properties the same API needs to be invoked twice, once with the AWS Glue connection ARN as ResourceArn with SourceProcessingProperties and next, with the AWS Glue database ARN as ResourceArn with TargetProcessingProperties respectively.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::IntegrationResourceProperty",
  "Properties" : {
      "ResourceArn" : String,
      "SourceProcessingProperties" : SourceProcessingProperties,
      "Tags" : [ Tag, ... ],
      "TargetProcessingProperties" : TargetProcessingProperties
    }
}

```

### YAML

```yaml

Type: AWS::Glue::IntegrationResourceProperty
Properties:
  ResourceArn: String
  SourceProcessingProperties:
    SourceProcessingProperties
  Tags:
    - Tag
  TargetProcessingProperties:
    TargetProcessingProperties

```

## Properties

`ResourceArn`

The connection ARN of the source, or the database ARN of the target.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:.*:.*:[0-9]+:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceProcessingProperties`

The resource properties associated with the integration source.

_Required_: No

_Type_: [SourceProcessingProperties](aws-properties-glue-integrationresourceproperty-sourceprocessingproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-integrationresourceproperty-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetProcessingProperties`

The structure used to define the resource properties associated with the integration target.

_Required_: No

_Type_: [TargetProcessingProperties](aws-properties-glue-integrationresourceproperty-targetprocessingproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns The resource property arn.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourcePropertyArn`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

SourceProcessingProperties

All content copied from https://docs.aws.amazon.com/.
