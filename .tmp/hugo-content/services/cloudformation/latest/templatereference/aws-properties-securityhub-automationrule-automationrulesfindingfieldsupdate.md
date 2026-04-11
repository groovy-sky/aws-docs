This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule AutomationRulesFindingFieldsUpdate

Identifies the finding fields that the automation rule action updates when a finding matches the defined criteria.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Confidence" : Integer,
  "Criticality" : Integer,
  "Note" : NoteUpdate,
  "RelatedFindings" : [ RelatedFinding, ... ],
  "Severity" : SeverityUpdate,
  "Types" : [ String, ... ],
  "UserDefinedFields" : {Key: Value, ...},
  "VerificationState" : String,
  "Workflow" : WorkflowUpdate
}

```

### YAML

```yaml

  Confidence: Integer
  Criticality: Integer
  Note:
    NoteUpdate
  RelatedFindings:
    - RelatedFinding
  Severity:
    SeverityUpdate
  Types:
    - String
  UserDefinedFields:
    Key: Value
  VerificationState: String
  Workflow:
    WorkflowUpdate

```

## Properties

`Confidence`

The rule action updates the `Confidence` field of a finding.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criticality`

The rule action updates the `Criticality` field of a finding.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Note`

The rule action will update the `Note` field of a finding.

_Required_: No

_Type_: [NoteUpdate](aws-properties-securityhub-automationrule-noteupdate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedFindings`

The rule action will update the `RelatedFindings` field of a finding.

_Required_: No

_Type_: Array of [RelatedFinding](aws-properties-securityhub-automationrule-relatedfinding.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Severity`

The rule action will update the `Severity` field of a finding.

_Required_: No

_Type_: [SeverityUpdate](aws-properties-securityhub-automationrule-severityupdate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Types`

The rule action updates the `Types` field of a finding.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserDefinedFields`

The rule action updates the `UserDefinedFields` field of a finding.

_Required_: No

_Type_: Object of String

_Pattern_: `^[-_+=.:/@\w\s]{1,128}$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerificationState`

The rule action updates the `VerificationState` field of a finding.

_Required_: No

_Type_: String

_Allowed values_: `UNKNOWN | TRUE_POSITIVE | FALSE_POSITIVE | BENIGN_POSITIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workflow`

The rule action will update the `Workflow` field of a finding.

_Required_: No

_Type_: [WorkflowUpdate](aws-properties-securityhub-automationrule-workflowupdate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutomationRulesAction

AutomationRulesFindingFilters

All content copied from https://docs.aws.amazon.com/.
