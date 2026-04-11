This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard Container

Information about the container used for the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Image" : String,
  "ModelDataUrl" : String,
  "NearestModelName" : String
}

```

### YAML

```yaml

  Image: String
  ModelDataUrl: String
  NearestModelName: String

```

## Properties

`Image`

The Amazon EC2 Container Registry (Amazon ECR) path where the model container image is stored.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelDataUrl`

The Amazon S3 path where the model artifacts are stored.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NearestModelName`

The name of a pre-trained model in the registry that is similar to this model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BusinessDetails

Content

All content copied from https://docs.aws.amazon.com/.
