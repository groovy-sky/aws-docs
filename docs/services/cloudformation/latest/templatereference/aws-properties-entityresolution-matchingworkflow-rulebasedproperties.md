This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow RuleBasedProperties

An object which defines the list of matching rules to run in a matching workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeMatchingModel" : String,
  "MatchPurpose" : String,
  "Rules" : [ Rule, ... ]
}

```

### YAML

```yaml

  AttributeMatchingModel: String
  MatchPurpose: String
  Rules:
    - Rule

```

## Properties

`AttributeMatchingModel`

The comparison type. You can choose `ONE_TO_ONE` or `MANY_TO_MANY` as the `attributeMatchingModel`.

If you choose `ONE_TO_ONE`, the system can only match attributes if the
sub-types are an exact match. For example, for the `Email` attribute type, the
system will only consider it a match if the value of the `Email` field of
Profile A matches the value of the `Email` field of Profile B.

If you choose `MANY_TO_MANY`, the system can match attributes across the
sub-types of an attribute type. For example, if the value of the `Email` field
of Profile A and the value of `BusinessEmail` field of Profile B matches, the
two profiles are matched on the `Email` attribute type.

_Required_: Yes

_Type_: String

_Allowed values_: `ONE_TO_ONE | MANY_TO_MANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchPurpose`

An indicator of whether to generate IDs and index the data or not.

If you choose `IDENTIFIER_GENERATION`, the process generates IDs and indexes
the data.

If you choose `INDEXING`, the process indexes the data without generating
IDs.

_Required_: No

_Type_: String

_Allowed values_: `IDENTIFIER_GENERATION | INDEXING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

A list of `Rule` objects, each of which have fields `RuleName` and
`MatchingKeys`.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-entityresolution-matchingworkflow-rule.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

RuleCondition

All content copied from https://docs.aws.amazon.com/.
