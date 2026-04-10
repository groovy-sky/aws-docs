This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow IdMappingTechniques

An object which defines the ID mapping technique and any additional
configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdMappingType" : String,
  "NormalizationVersion" : String,
  "ProviderProperties" : ProviderProperties,
  "RuleBasedProperties" : IdMappingRuleBasedProperties
}

```

### YAML

```yaml

  IdMappingType: String
  NormalizationVersion: String
  ProviderProperties:
    ProviderProperties
  RuleBasedProperties:
    IdMappingRuleBasedProperties

```

## Properties

`IdMappingType`

The type of ID mapping.

_Required_: No

_Type_: String

_Allowed values_: `PROVIDER | RULE_BASED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NormalizationVersion`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderProperties`

An object which defines any additional configurations required by the provider
service.

_Required_: No

_Type_: [ProviderProperties](aws-properties-entityresolution-idmappingworkflow-providerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleBasedProperties`

An object which defines any additional configurations required by rule-based
matching.

_Required_: No

_Type_: [IdMappingRuleBasedProperties](aws-properties-entityresolution-idmappingworkflow-idmappingrulebasedproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdMappingRuleBasedProperties

IdMappingWorkflowInputSource

All content copied from https://docs.aws.amazon.com/.
