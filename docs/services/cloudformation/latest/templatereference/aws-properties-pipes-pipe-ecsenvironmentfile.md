This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsEnvironmentFile

A list of files containing the environment variables to pass to a container. You can
specify up to ten environment files. The file must have a `.env` file extension.
Each line in an environment file should contain an environment variable in
`VARIABLE=VALUE` format. Lines beginning with `#` are treated as
comments and are ignored. For more information about the environment variable file syntax,
see [Declare default environment\
variables in file](https://docs.docker.com/compose/env-file).

If there are environment variables specified using the `environment`
parameter in a container definition, they take precedence over the variables contained
within an environment file. If multiple environment files are specified that contain the
same variable, they're processed from the top down. We recommend that you use unique
variable names. For more information, see [Specifying environment\
variables](../../../amazonecs/latest/developerguide/taskdef-envfiles.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

This parameter is only supported for tasks hosted on Fargate using the
following platform versions:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

The file type to use. The only supported value is `s3`.

_Required_: Yes

_Type_: String

_Allowed values_: `s3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The Amazon Resource Name (ARN) of the Amazon S3 object containing the
environment variable file.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsContainerOverride

EcsEnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
