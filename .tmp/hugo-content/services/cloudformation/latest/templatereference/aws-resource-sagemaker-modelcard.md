This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard

Creates an Amazon SageMaker Model Card.

For information about how to use model cards, see [Amazon SageMaker Model Card](../../../sagemaker/latest/dg/model-cards.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::ModelCard",
  "Properties" : {
      "Content" : Content,
      "CreatedBy" : UserContext,
      "LastModifiedBy" : UserContext,
      "ModelCardName" : String,
      "ModelCardStatus" : String,
      "SecurityConfig" : SecurityConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::ModelCard
Properties:
  Content:
    Content
  CreatedBy:
    UserContext
  LastModifiedBy:
    UserContext
  ModelCardName: String
  ModelCardStatus: String
  SecurityConfig:
    SecurityConfig
  Tags:
    - Tag

```

## Properties

`Content`

The content of the model card. Content uses the [model card JSON schema](../../../sagemaker/latest/dg/model-cards.md#model-cards-json-schema).

_Required_: Yes

_Type_: [Content](aws-properties-sagemaker-modelcard-content.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedBy`

Information about the user who created or modified one or more of the following:

- Experiment

- Trial

- Trial component

- Lineage group

- Project

- Model Card

_Required_: No

_Type_: [UserContext](aws-properties-sagemaker-modelcard-usercontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastModifiedBy`

Information about the user who last modified the model card.

_Required_: No

_Type_: [UserContext](aws-properties-sagemaker-modelcard-usercontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelCardName`

The unique name of the model card.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}$`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelCardStatus`

The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.

- `Draft`: The model card is a work in progress.

- `PendingReview`: The model card is pending review.

- `Approved`: The model card is approved.

- `Archived`: The model card is archived. No more updates should be made to the model
card, but it can still be exported.

_Required_: Yes

_Type_: String

_Allowed values_: `Draft | PendingReview | Approved | Archived`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityConfig`

The security configuration used to protect model card data.

_Required_: No

_Type_: [SecurityConfig](aws-properties-sagemaker-modelcard-securityconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs used to manage metadata for the model card.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-modelcard-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the model card name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedBy.DomainId`

Property description not available.

`LastModifiedBy.UserProfileArn`

Property description not available.

`ModelCardArn`

The Amazon Resource Number (ARN) of the model card. For example,
`arn:aws:sagemaker:us-west-2:012345678901:modelcard/examplemodelcard`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

AdditionalInformation

All content copied from https://docs.aws.amazon.com/.
