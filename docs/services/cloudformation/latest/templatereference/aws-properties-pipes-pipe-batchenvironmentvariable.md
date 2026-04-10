This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe BatchEnvironmentVariable

The environment variables to send to the container. You can add new environment
variables, which are added to the container at launch, or you can override the existing
environment variables from the Docker image or the task definition.

###### Note

Environment variables cannot start with " `
                        AWS Batch
                    `". This
naming convention is reserved for variables that AWS Batch sets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the key-value pair. For environment variables, this is the name of the
environment variable.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the key-value pair. For environment variables, this is the value of the
environment variable.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchContainerOverrides

BatchJobDependency

All content copied from https://docs.aws.amazon.com/.
