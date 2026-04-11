This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model HubAccessConfig

The configuration for a private hub model reference that points to a public SageMaker JumpStart model.

For more information about private hubs, see [Private curated hubs for foundation model access control in JumpStart](../../../sagemaker/latest/dg/jumpstart-curated-hubs.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HubContentArn" : String
}

```

### YAML

```yaml

  HubContentArn: String

```

## Properties

`HubContentArn`

The ARN of your private model hub content. This should be a `ModelReference`
resource type that points to a SageMaker JumpStart public hub model.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContainerDefinition

ImageConfig

All content copied from https://docs.aws.amazon.com/.
