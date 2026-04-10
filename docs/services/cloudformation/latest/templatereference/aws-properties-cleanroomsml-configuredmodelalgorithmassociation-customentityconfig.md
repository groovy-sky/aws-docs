This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation CustomEntityConfig

The configuration for defining custom patterns to be redacted from logs and error messages. This is for the CUSTOM
config under entitiesToRedact. Both CustomEntityConfig and entitiesToRedact need to be present or not present.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomDataIdentifiers" : [ String, ... ]
}

```

### YAML

```yaml

  CustomDataIdentifiers:
    - String

```

## Properties

`CustomDataIdentifiers`

Defines data identifiers for the custom entity configuration. Provide this only if CUSTOM redaction is configured.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `200 | 10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation

LogRedactionConfiguration

All content copied from https://docs.aws.amazon.com/.
