This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service ImageConfiguration

Describes the configuration that AWS App Runner uses to run an App Runner service using an image pulled from a source image repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Port" : String,
  "RuntimeEnvironmentSecrets" : [ KeyValuePair, ... ],
  "RuntimeEnvironmentVariables" : [ KeyValuePair, ... ],
  "StartCommand" : String
}

```

### YAML

```yaml

  Port: String
  RuntimeEnvironmentSecrets:
    - KeyValuePair
  RuntimeEnvironmentVariables:
    - KeyValuePair
  StartCommand: String

```

## Properties

`Port`

The port that your application listens to in the container.

Default: `8080`

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeEnvironmentSecrets`

An array of key-value pairs representing the secrets and parameters that get referenced to your service as an environment variable. The supported
values are either the full Amazon Resource Name (ARN) of the AWS Secrets Manager secret or the full ARN of the parameter in the AWS Systems Manager
Parameter Store.

###### Note

- If the AWS Systems Manager Parameter Store parameter exists in the same AWS Region as the service that you're launching, you can use
either the full ARN or name of the secret. If the parameter exists in a different Region, then the full ARN must be specified.

- Currently, cross account referencing of AWS Systems Manager Parameter Store parameter is not supported.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-apprunner-service-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeEnvironmentVariables`

Environment variables that are available to your running App Runner service. An array of key-value pairs.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-apprunner-service-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartCommand`

An optional command that App Runner runs to start the application in the source image. If specified, this command overrides the Docker image’s default start
command.

_Required_: No

_Type_: String

_Pattern_: `[^\x0a\x0d]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HealthCheckConfiguration

ImageRepository

All content copied from https://docs.aws.amazon.com/.
