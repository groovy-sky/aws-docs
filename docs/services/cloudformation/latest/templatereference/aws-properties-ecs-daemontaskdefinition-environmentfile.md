This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition EnvironmentFile

A list of files containing the environment variables to pass to a container. You can
specify up to ten environment files. The file must have a `.env` file
extension. Each line in an environment file should contain an environment variable in
`VARIABLE=VALUE` format. Lines beginning with `#` are treated
as comments and are ignored.

If there are environment variables specified using the `environment`
parameter in a container definition, they take precedence over the variables contained
within an environment file. If multiple environment files are specified that contain the
same variable, they're processed from the top down. We recommend that you use unique
variable names. For more information, see [Use a file to pass\
environment variables to a container](../../../amazonecs/latest/developerguide/use-environment-file.md) in the _Amazon Elastic_
_Container Service Developer Guide_.

Environment variable files are objects in Amazon S3 and all Amazon S3 security
considerations apply.

You must use the following platforms for the Fargate launch type:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

Consider the following when using the Fargate launch type:

- The file is handled like a native Docker env-file.

- There is no support for shell escape handling.

- The container entry point interperts the `VARIABLE` values.

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

The file type to use. Environment files are objects in Amazon S3. The only supported
value is `s3`.

_Required_: No

_Type_: String

_Allowed values_: `s3`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment
variable file.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Device

FirelensConfiguration

All content copied from https://docs.aws.amazon.com/.
