This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration TargetContainerRepository

The container repository where the output container image is stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepositoryName" : String,
  "Service" : String
}

```

### YAML

```yaml

  RepositoryName: String
  Service: String

```

## Properties

`RepositoryName`

The name of the container repository where the output container image is stored.
This name is prefixed by the repository location. For example,
`<repository location url>/repository_name`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Service`

Specifies the service in which this image was registered.

_Required_: No

_Type_: String

_Allowed values_: `ECR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SsmParameterConfiguration

AWS::ImageBuilder::Image

All content copied from https://docs.aws.amazon.com/.
