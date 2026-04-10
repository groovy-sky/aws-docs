This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::IdNamespace NamespaceRuleBasedProperties

The rule-based properties of an ID namespace. These properties define how the ID
namespace can be used in an ID mapping workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeMatchingModel" : String,
  "RecordMatchingModels" : [ String, ... ],
  "RuleDefinitionTypes" : [ String, ... ],
  "Rules" : [ Rule, ... ]
}

```

### YAML

```yaml

  AttributeMatchingModel: String
  RecordMatchingModels:
    - String
  RuleDefinitionTypes:
    - String
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
of Profile A matches the value of `BusinessEmail` field of Profile B, the two
profiles are matched on the `Email` attribute type.

_Required_: No

_Type_: String

_Allowed values_: `ONE_TO_ONE | MANY_TO_MANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordMatchingModels`

The type of matching record that is allowed to be used in an ID mapping workflow.

If the value is set to `ONE_SOURCE_TO_ONE_TARGET`, only one record in the
source is matched to one record in the target.

If the value is set to `MANY_SOURCE_TO_ONE_TARGET`, all matching records in
the source are matched to one record in the target.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleDefinitionTypes`

The sets of rules you can use in an ID mapping workflow. The limitations specified for
the source and target must be compatible.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The rules for the ID namespace.

_Required_: No

_Type_: Array of [Rule](aws-properties-entityresolution-idnamespace-rule.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NamespaceProviderProperties

Rule

All content copied from https://docs.aws.amazon.com/.
