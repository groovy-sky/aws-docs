This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdMappingWorkflow IdMappingRuleBasedProperties

An object that defines the list of matching rules to run in an ID mapping
workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeMatchingModel" : String,
  "RecordMatchingModel" : String,
  "RuleDefinitionType" : String,
  "Rules" : [ Rule, ... ]
}

```

### YAML

```yaml

  AttributeMatchingModel: String
  RecordMatchingModel: String
  RuleDefinitionType: String
  Rules:
    - Rule

```

## Properties

`AttributeMatchingModel`

The comparison type. You can either choose `ONE_TO_ONE` or
`MANY_TO_MANY` as the `attributeMatchingModel`.

If you choose `ONE_TO_ONE`, the system can only match attributes if the
sub-types are an exact match. For example, for the `Email` attribute type, the
system will only consider it a match if the value of the `Email` field of
Profile A matches the value of the `Email` field of Profile B.

If you choose `MANY_TO_MANY`, the system can match attributes across the
sub-types of an attribute type. For example, if the value of the `Email` field
of Profile A matches the value of the `BusinessEmail` field of Profile B, the
two profiles are matched on the `Email` attribute type.

_Required_: Yes

_Type_: String

_Allowed values_: `ONE_TO_ONE | MANY_TO_MANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordMatchingModel`

The type of matching record that is allowed to be used in an ID mapping workflow.

If the value is set to `ONE_SOURCE_TO_ONE_TARGET`, only one record in the
source can be matched to the same record in the target.

If the value is set to `MANY_SOURCE_TO_ONE_TARGET`, multiple records in the
source can be matched to one record in the target.

_Required_: Yes

_Type_: String

_Allowed values_: `ONE_SOURCE_TO_ONE_TARGET | MANY_SOURCE_TO_ONE_TARGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleDefinitionType`

The set of rules you can use in an ID mapping workflow. The limitations specified for
the source or target to define the match rules must be compatible.

_Required_: No

_Type_: String

_Allowed values_: `SOURCE | TARGET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The rules that can be used for ID mapping.

_Required_: No

_Type_: Array of [Rule](aws-properties-entityresolution-idmappingworkflow-rule.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdMappingIncrementalRunConfig

IdMappingTechniques

All content copied from https://docs.aws.amazon.com/.
