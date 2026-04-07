This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model InferenceExecutionConfig

Specifies details about how containers in a multi-container endpoint are run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

How containers in a multi-container are run. The following values are valid.

- `Serial` \- Containers run as a serial pipeline.

- `Direct` \- Only the individual container that you specify is run.

_Required_: Yes

_Type_: String

_Allowed values_: `Serial | Direct`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImageConfig

ModelAccessConfig
