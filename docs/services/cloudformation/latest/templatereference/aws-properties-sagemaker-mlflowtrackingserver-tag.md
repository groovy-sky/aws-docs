This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MlflowTrackingServer Tag

A tag object that consists of a key and an optional value, used to manage metadata
for SageMaker AWS resources.

You can add tags to notebook instances, training jobs, hyperparameter tuning jobs,
batch transform jobs, models, labeling jobs, work teams, endpoint configurations, and
endpoints. For more information on adding tags to SageMaker resources, see [AddTags](../../../../reference/sagemaker/latest/apireference/api-addtags.md).

For more information on adding metadata to your AWS resources with
tagging, see [Tagging AWS resources](../../../../general/latest/gr/aws-tagging.md). For advice on best practices for
managing AWS resources with tagging, see [Tagging\
Best Practices: Implement an Effective AWS Resource Tagging\
Strategy](https://d1.awsstatic.com/whitepapers/aws-tagging-best-practices.pdf).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The tag key. Tag keys must be unique per resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::MlflowTrackingServer

AWS::SageMaker::Model

All content copied from https://docs.aws.amazon.com/.
