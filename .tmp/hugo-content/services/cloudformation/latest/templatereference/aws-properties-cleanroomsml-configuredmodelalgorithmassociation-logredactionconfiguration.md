This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation LogRedactionConfiguration

The configuration for log redaction.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomEntityConfig" : CustomEntityConfig,
  "EntitiesToRedact" : [ String, ... ]
}

```

### YAML

```yaml

  CustomEntityConfig:
    CustomEntityConfig
  EntitiesToRedact:
    - String

```

## Properties

`CustomEntityConfig`

Specifies the configuration for custom entities in the context of log redaction.

_Required_: No

_Type_: [CustomEntityConfig](aws-properties-cleanroomsml-configuredmodelalgorithmassociation-customentityconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntitiesToRedact`

Specifies the entities to be redacted from logs. Entities to redact are "ALL\_PERSONALLY\_IDENTIFIABLE\_INFORMATION",
"NUMBERS","CUSTOM". If CUSTOM is supplied or configured, custom patterns (customDataIdentifiers) should be provided,
and the patterns will be redacted in logs or error messages.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomEntityConfig

LogsConfigurationPolicy

All content copied from https://docs.aws.amazon.com/.
