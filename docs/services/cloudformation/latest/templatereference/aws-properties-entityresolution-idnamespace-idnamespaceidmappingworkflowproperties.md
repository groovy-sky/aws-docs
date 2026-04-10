This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdNamespace IdNamespaceIdMappingWorkflowProperties

An object containing `idMappingType`, `providerProperties`, and
`ruleBasedProperties`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdMappingType" : String,
  "ProviderProperties" : NamespaceProviderProperties,
  "RuleBasedProperties" : NamespaceRuleBasedProperties
}

```

### YAML

```yaml

  IdMappingType: String
  ProviderProperties:
    NamespaceProviderProperties
  RuleBasedProperties:
    NamespaceRuleBasedProperties

```

## Properties

`IdMappingType`

The type of ID mapping.

_Required_: Yes

_Type_: String

_Allowed values_: `PROVIDER | RULE_BASED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderProperties`

An object which defines any additional configurations required by the provider
service.

_Required_: No

_Type_: [NamespaceProviderProperties](aws-properties-entityresolution-idnamespace-namespaceproviderproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleBasedProperties`

An object which defines any additional configurations required by rule-based
matching.

_Required_: No

_Type_: [NamespaceRuleBasedProperties](aws-properties-entityresolution-idnamespace-namespacerulebasedproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EntityResolution::IdNamespace

IdNamespaceInputSource

All content copied from https://docs.aws.amazon.com/.
