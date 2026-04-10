This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Multiplexprogram MultiplexProgramServiceDescriptor

Transport stream service descriptor configuration for the Multiplex program.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProviderName" : String,
  "ServiceName" : String
}

```

### YAML

```yaml

  ProviderName: String
  ServiceName: String

```

## Properties

`ProviderName`

Name of the provider.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

Name of the service.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexProgramPipelineDetail

MultiplexProgramSettings

All content copied from https://docs.aws.amazon.com/.
