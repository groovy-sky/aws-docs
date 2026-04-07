This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service ServiceObservabilityConfiguration

Describes the observability configuration of an AWS App Runner service. These are additional observability features, like tracing, that you choose to
enable. They're configured in a separate resource that you associate with your service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObservabilityConfigurationArn" : String,
  "ObservabilityEnabled" : Boolean
}

```

### YAML

```yaml

  ObservabilityConfigurationArn: String
  ObservabilityEnabled: Boolean

```

## Properties

`ObservabilityConfigurationArn`

The Amazon Resource Name (ARN) of the observability configuration that is associated with the service. Specified only when
`ObservabilityEnabled` is `true`.

Specify an ARN with a name and a revision number to associate that revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing/3`

Specify just the name to associate the latest revision. For example:
`arn:aws:apprunner:us-east-1:123456789012:observabilityconfiguration/xray-tracing`

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[0-9]{12}:(\w|/|-){1,1011}`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObservabilityEnabled`

When `true`, an observability configuration resource is associated with the service, and an `ObservabilityConfigurationArn` is
specified.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NetworkConfiguration

SourceCodeVersion
