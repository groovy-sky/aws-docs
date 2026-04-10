This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow ResolutionTechniques

An object which defines the `resolutionType` and the
`ruleBasedProperties`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProviderProperties" : ProviderProperties,
  "ResolutionType" : String,
  "RuleBasedProperties" : RuleBasedProperties,
  "RuleConditionProperties" : RuleConditionProperties
}

```

### YAML

```yaml

  ProviderProperties:
    ProviderProperties
  ResolutionType: String
  RuleBasedProperties:
    RuleBasedProperties
  RuleConditionProperties:
    RuleConditionProperties

```

## Properties

`ProviderProperties`

The properties of the provider service.

_Required_: No

_Type_: [ProviderProperties](aws-properties-entityresolution-matchingworkflow-providerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolutionType`

The type of matching workflow to create. Specify one of the following types:

- `RULE_MATCHING`: Match records using configurable rule-based criteria

- `ML_MATCHING`: Match records using machine learning models

- `PROVIDER`: Match records using a third-party matching provider

_Required_: No

_Type_: String

_Allowed values_: `RULE_MATCHING | ML_MATCHING | PROVIDER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleBasedProperties`

An object which defines the list of matching rules to run and has a field
`rules`, which is a list of rule objects.

_Required_: No

_Type_: [RuleBasedProperties](aws-properties-entityresolution-matchingworkflow-rulebasedproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleConditionProperties`

An object containing the `rules` for a matching workflow.

_Required_: No

_Type_: [RuleConditionProperties](aws-properties-entityresolution-matchingworkflow-ruleconditionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProviderProperties

Rule

All content copied from https://docs.aws.amazon.com/.
