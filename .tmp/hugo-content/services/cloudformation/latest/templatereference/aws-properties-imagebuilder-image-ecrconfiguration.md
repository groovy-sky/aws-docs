This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image EcrConfiguration

Settings that Image Builder uses to configure the ECR repository and the output container
images that Amazon Inspector scans.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerTags" : [ String, ... ],
  "RepositoryName" : String
}

```

### YAML

```yaml

  ContainerTags:
    - String
  RepositoryName: String

```

## Properties

`ContainerTags`

Tags for Image Builder to apply to the output container image that Amazon Inspector scans. Tags can
help you identify and manage your scanned images.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryName`

The name of the container repository that Amazon Inspector scans to identify findings for your
container images. The name includes the path for the repository location. If you don’t
provide this information, Image Builder creates a repository in your account named
`image-builder-image-scanning-repository` for vulnerability scans of your
output container images.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeletionSettings

ImageLoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
