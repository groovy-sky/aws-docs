This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation MetricsConfigurationPolicy

Provides the configuration policy for metrics generation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NoiseLevel" : String
}

```

### YAML

```yaml

  NoiseLevel: String

```

## Properties

`NoiseLevel`

The noise level for the generated metrics.

_Required_: Yes

_Type_: String

_Allowed values_: `HIGH | MEDIUM | LOW | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogsConfigurationPolicy

PrivacyConfiguration

All content copied from https://docs.aws.amazon.com/.
