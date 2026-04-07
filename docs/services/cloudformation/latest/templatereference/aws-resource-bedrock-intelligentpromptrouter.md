This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::IntelligentPromptRouter

Specifies an intelligent prompt router resource for Amazon Bedrock.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::IntelligentPromptRouter",
  "Properties" : {
      "Description" : String,
      "FallbackModel" : PromptRouterTargetModel,
      "Models" : [ PromptRouterTargetModel, ... ],
      "PromptRouterName" : String,
      "RoutingCriteria" : RoutingCriteria,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::IntelligentPromptRouter
Properties:
  Description: String
  FallbackModel:
    PromptRouterTargetModel
  Models:
    - PromptRouterTargetModel
  PromptRouterName: String
  RoutingCriteria:
    RoutingCriteria
  Tags:
    - Tag

```

## Properties

`Description`

An optional description of the prompt router to help identify its purpose.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-zA-Z:.][ _-]?)+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FallbackModel`

The default model to use when the routing criteria is not met.

_Required_: Yes

_Type_: [PromptRouterTargetModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-intelligentpromptrouter-promptroutertargetmodel.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Models`

A list of foundation models that the prompt router can route requests to. At least one model must be specified.

_Required_: Yes

_Type_: Array of [PromptRouterTargetModel](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-intelligentpromptrouter-promptroutertargetmodel.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PromptRouterName`

The name of the prompt router. The name must be unique within your AWS account in the current region.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][ _-]?)+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoutingCriteria`

Routing criteria for a prompt router.

_Required_: Yes

_Type_: [RoutingCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-intelligentpromptrouter-routingcriteria.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource as tags. You can use tags to categorize and manage your AWS resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-intelligentpromptrouter-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the intelligent prompt router information.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

Indicates the time that the prompt router was created.

`PromptRouterArn`

The Amazon Resource Name (ARN) of the prompt router.

`Status`

The router's status.

`Type`

The router's type.

`UpdatedAt`

When the router was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Bedrock::GuardrailVersion

PromptRouterTargetModel
