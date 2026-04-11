This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow ProviderProperties

An object containing the `providerServiceARN`,
`intermediateSourceConfiguration`, and
`providerConfiguration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IntermediateSourceConfiguration" : IntermediateSourceConfiguration,
  "ProviderConfiguration" : {Key: Value, ...},
  "ProviderServiceArn" : String
}

```

### YAML

```yaml

  IntermediateSourceConfiguration:
    IntermediateSourceConfiguration
  ProviderConfiguration:
    Key: Value
  ProviderServiceArn: String

```

## Properties

`IntermediateSourceConfiguration`

The Amazon S3 location that temporarily stores your data while it processes.
Your information won't be saved permanently.

_Required_: No

_Type_: [IntermediateSourceConfiguration](aws-properties-entityresolution-idmappingworkflow-intermediatesourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderConfiguration`

The required configuration fields to use with the provider service.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderServiceArn`

The ARN of the provider service.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):(entityresolution):([a-z]{2}-[a-z]{1,10}-[0-9])::providerservice/([a-zA-Z0-9_-]{1,255})/([a-zA-Z0-9_-]{1,255})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntermediateSourceConfiguration

Rule

All content copied from https://docs.aws.amazon.com/.
