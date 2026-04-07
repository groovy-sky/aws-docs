This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelPackage ModelCard

An Amazon SageMaker Model Card.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ModelCardContent" : String,
  "ModelCardStatus" : String
}

```

### YAML

```yaml

  ModelCardContent: String
  ModelCardStatus: String

```

## Properties

`ModelCardContent`

The content of the model card associated with the model package.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `100000`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

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

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelAccessConfig

ModelDataQuality
