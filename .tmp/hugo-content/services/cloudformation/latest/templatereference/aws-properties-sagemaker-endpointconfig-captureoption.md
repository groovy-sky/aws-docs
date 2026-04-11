This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig CaptureOption

Specifies whether the endpoint captures input data or output data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaptureMode" : String
}

```

### YAML

```yaml

  CaptureMode: String

```

## Properties

`CaptureMode`

Specifies whether the endpoint captures input data or output data.

_Required_: Yes

_Type_: String

_Allowed values_: `Input | Output | InputAndOutput`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CaptureContentTypeHeader

ClarifyExplainerConfig

All content copied from https://docs.aws.amazon.com/.
