This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Host

Determine whether your data volume persists on the host container instance and where it's
stored. If this parameter is empty, then the Docker daemon assigns a host path for your data
volume. However, the data isn't guaranteed to persist after the containers that are associated
with it stop running.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourcePath" : String
}

```

### YAML

```yaml

  SourcePath: String

```

## Properties

`SourcePath`

The path on the host container instance that's presented to the container. If this parameter
is empty, then the Docker daemon has assigned a host path for you. If this parameter contains a
file location, then the data volume persists at the specified location on the host container
instance until you delete it manually. If the source path location doesn't exist on the host
container instance, the Docker daemon creates it. If the location does exist, the contents of the
source path folder are exported.

###### Note

This parameter isn't applicable to jobs that run on Fargate resources. Don't provide this
for these jobs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirelensConfiguration

ImagePullSecret

All content copied from https://docs.aws.amazon.com/.
