This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule

The `AWS::SecurityHub::AutomationRule` resource specifies an automation rule based on input parameters. For more information, see
[Automation rules](../../../securityhub/latest/userguide/automation-rules.md) in the _AWS Security Hub CSPM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::AutomationRule",
  "Properties" : {
      "Actions" : [ AutomationRulesAction, ... ],
      "Criteria" : AutomationRulesFindingFilters,
      "Description" : String,
      "IsTerminal" : Boolean,
      "RuleName" : String,
      "RuleOrder" : Integer,
      "RuleStatus" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::AutomationRule
Properties:
  Actions:
    - AutomationRulesAction
  Criteria:
    AutomationRulesFindingFilters
  Description: String
  IsTerminal: Boolean
  RuleName: String
  RuleOrder: Integer
  RuleStatus: String
  Tags:
    Key: Value

```

## Properties

`Actions`

One or more actions to update finding fields if a finding matches the conditions
specified in `Criteria`.

_Required_: Yes

_Type_: Array of [AutomationRulesAction](aws-properties-securityhub-automationrule-automationrulesaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criteria`

A set of [AWS Security Finding Format (ASFF)](../../../securityhub/latest/userguide/securityhub-findings-format.md) finding field attributes and corresponding expected values that
Security Hub CSPM uses to filter findings. If a rule is enabled and a finding matches the criteria specified in
this parameter, Security Hub CSPM applies the rule action to the finding.

_Required_: Yes

_Type_: [AutomationRulesFindingFilters](aws-properties-securityhub-automationrule-automationrulesfindingfilters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsTerminal`

Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria. This is useful when a finding
matches the criteria for multiple rules, and each rule has different actions. If a rule is terminal, Security Hub CSPM applies the rule action to a finding that matches
the rule criteria and doesn't evaluate other rules for the finding. By default, a rule isn't terminal.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleName`

The name of the rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleOrder`

An integer ranging from 1 to 1000 that represents the order in which the rule action is
applied to findings. Security Hub CSPM applies rules with lower values for this parameter
first.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleStatus`

Whether the rule is active after it is created. If
this parameter is equal to `ENABLED`, Security Hub CSPM applies the rule to findings
and finding updates after the rule is created.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

User-defined tags associated with an automation rule.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9]{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `RuleArn`. For example, `arn:aws:securityhub:us-east-1:123456789012:automation-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

A timestamp that indicates when the rule was created.

Uses the `date-time` format specified in [RFC 3339 section 5.6, Internet\
Date/Time Format](https://tools.ietf.org/html/rfc3339). The value cannot contain spaces. For example,
`2020-03-22T13:22:13.933Z`.

`CreatedBy`

The principal that created the rule. For example, `arn:aws:sts::123456789012:assumed-role/Developer-Role/JaneDoe`.

`RuleArn`

The Amazon Resource Name (ARN) of the automation rule that you create. For example, `arn:aws:securityhub:us-east-1:123456789012:automation-rule/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

`UpdatedAt`

A timestamp that indicates when the rule was most recently updated.

Uses the `date-time` format specified in [RFC 3339 section 5.6, Internet\
Date/Time Format](https://tools.ietf.org/html/rfc3339). The value cannot contain spaces. For example,
`2020-03-22T13:22:13.933Z`.

## Examples

The following examples demonstrate how to declare an `AWS::SecurityHub::AutomationRule` resource.

### Creating an automation rule

This example creates a Security Hub CSPM automation rule. This example contains all available
fields for `Actions` and `Criteria` for demonstration purposes.

#### JSON

```json

{
  "Description": "Example template to create a Security Hub automation rule",
  "Resources": {
    "RuleWithCriteriaActionsTags": {
      "Type": "AWS::SecurityHub::AutomationRule",
      "Properties": {
        "RuleName": "Example rule name",
        "RuleOrder": 5,
        "Description": "Example rule description.",
        "IsTerminal": false,
        "RuleStatus": "ENABLED",
        "Criteria": {
          "ProductName": [
            {
              "Comparison": "EQUALS",
              "Value": "GuardDuty"
            },
            {
              "Comparison": "PREFIX",
              "Value": "SecurityHub"
            }
          ],
          "CompanyName": [
            {
              "Comparison": "EQUALS",
              "Value": "AWS"
            },
            {
              "Comparison": "PREFIX",
              "Value": "Private"
            }
          ],
          "ProductArn": [
            {
              "Comparison": "EQUALS",
              "Value": "arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default"
            },
            {
              "Comparison": "PREFIX",
              "Value": "arn:aws:securityhub:us-west-2:123456789012:product/aws"
            }
          ],
          "AwsAccountId": [
            {
              "Comparison": "EQUALS",
              "Value": 123456789012
            }
          ],
          "Id": [
            {
              "Comparison": "EQUALS",
              "Value": "example-finding-id"
            }
          ],
          "GeneratorId": [
            {
              "Comparison": "EQUALS",
              "Value": "example-generator-id"
            }
          ],
          "Type": [
            {
              "Comparison": "EQUALS",
              "Value": "type-1"
            },
            {
              "Comparison": "EQUALS",
              "Value": "type-2"
            }
          ],
          "Description": [
            {
              "Comparison": "EQUALS",
              "Value": "description1"
            },
            {
              "Comparison": "EQUALS",
              "Value": "description2"
            }
          ],
          "SourceUrl": [
            {
              "Comparison": "PREFIX",
              "Value": "https"
            },
            {
              "Comparison": "PREFIX",
              "Value": "ftp"
            }
          ],
          "Title": [
            {
              "Comparison": "EQUALS",
              "Value": "title-1"
            },
            {
              "Comparison": "PREFIX",
              "Value": "title-2"
            }
          ],
          "SeverityLabel": [
            {
              "Comparison": "EQUALS",
              "Value": "LOW"
            },
            {
              "Comparison": "EQUALS",
              "Value": "HIGH"
            }
          ],
          "ResourceType": [
            {
              "Comparison": "EQUALS",
              "Value": "AwsEc2Instance"
            }
          ],
          "ResourcePartition": [
            {
              "Comparison": "EQUALS",
              "Value": "aws"
            }
          ],
          "ResourceId": [
            {
              "Comparison": "PREFIX",
              "Value": "i-1234567890"
            }
          ],
          "ResourceRegion": [
            {
              "Comparison": "PREFIX",
              "Value": "us-west"
            }
          ],
          "ComplianceStatus": [
            {
              "Comparison": "EQUALS",
              "Value": "FAILED"
            }
          ],
          "ComplianceSecurityControlId": [
            {
              "Comparison": "EQUALS",
              "Value": "EC2.3"
            }
          ],
          "ComplianceAssociatedStandardsId": [
            {
              "Comparison": "EQUALS",
              "Value": "ruleset/cis-aws-foundations-benchmark/v/1.2.0"
            }
          ],
          "VerificationState": [
            {
              "Comparison": "EQUALS",
              "Value": "BENIGN_POSITIVE"
            }
          ],
          "RecordState": [
            {
              "Comparison": "EQUALS",
              "Value": "ACTIVE"
            }
          ],
          "RelatedFindingsProductArn": [
            {
              "Comparison": "EQUALS",
              "Value": "arn:aws:securityhub:eu-central-1::product/aws/securityhub"
            }
          ],
          "RelatedFindingsId": [
            {
              "Comparison": "EQUALS",
              "Value": "example-finding-id-2"
            }
          ],
          "NoteText": [
            {
              "Comparison": "EQUALS",
              "Value": "example-note-text"
            }
          ],
          "NoteUpdatedAt": [
            {
              "DateRange": {
                "Unit": "DAYS",
                "Value": 5
              }
            }
          ],
          "NoteUpdatedBy": [
            {
              "Comparison": "PREFIX",
              "Value": "sechub"
            }
          ],
          "WorkflowStatus": [
            {
              "Comparison": "EQUALS",
              "Value": "NEW"
            }
          ],
          "FirstObservedAt": [
            {
              "DateRange": {
                "Unit": "DAYS",
                "Value": 5
              }
            }
          ],
          "LastObservedAt": [
            {
              "DateRange": {
                "Unit": "DAYS",
                "Value": 5
              }
            }
          ],
          "CreatedAt": [
            {
              "DateRange": {
                "Unit": "DAYS",
                "Value": 5
              }
            }
          ],
          "UpdatedAt": [
            {
              "Start": "2023-04-25T17:05:54.832Z",
              "End": "2023-05-25T17:05:54.832Z"
            }
          ],
          "ResourceTags": [
            {
              "Comparison": "NOT_EQUALS",
              "Key": "department",
              "Value": "security"
            },
            {
              "Comparison": "NOT_EQUALS",
              "Key": "department",
              "Value": "operations"
            }
          ],
          "UserDefinedFields": [
            {
              "Comparison": "EQUALS",
              "Key": "key1",
              "Value": "security"
            },
            {
              "Comparison": "EQUALS",
              "Key": "key2",
              "Value": "operations"
            }
          ],
          "ResourceDetailsOther": [
            {
              "Comparison": "NOT_EQUALS",
              "Key": "area",
              "Value": "na"
            },
            {
              "Comparison": "NOT_EQUALS",
              "Key": "department",
              "Value": "sales"
            }
          ],
          "Confidence": [
            {
              "Gte": 50,
              "Lte": 95
            }
          ],
          "Criticality": [
            {
              "Gte": 50,
              "Lte": 95
            }
          ]
        },
        "Actions": [
          {
            "Type": "FINDING_FIELDS_UPDATE",
            "FindingFieldsUpdate": {
              "Severity": {
                "Product": 50,
                "Label": "MEDIUM",
                "Normalized": 60
              },
              "Types": [
                "Software and Configuration Checks/Industry and Regulatory Standards/AWS-Foundational-Security-Best-Practices",
                "Industry Compliance"
              ],
              "Confidence": 98,
              "Criticality": 95,
              "UserDefinedFields": {
                "key1": "value1",
                "key2": "value2"
              },
              "RelatedFindings": [
                {
                  "ProductArn": "arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default",
                  "Id": "sample-finding-id-1"
                },
                {
                  "ProductArn": "arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default",
                  "Id": "sample-finding-id-2"
                }
              ],
              "Note": {
                "Text": "sample-note-text",
                "UpdatedBy": "sechub"
              },
              "VerificationState": "TRUE_POSITIVE",
              "Workflow": {
                "Status": "NOTIFIED"
              }
            }
          }
        ],
        "Tags": {
          "sampleTag": "sampleValue",
          "organizationUnit": "pnw"
        }
      }
    }
  }
}
```

#### YAML

```yaml

Description: Example template to create a Security Hub automation rule
Resources:
  RuleWithCriteriaActionsTags:
    Type: "AWS::SecurityHub::AutomationRule"
    Properties:
      RuleName: "Example rule name"
      RuleOrder: 5
      Description: "Example rule description."
      IsTerminal: false
      RuleStatus: "ENABLED"
      Criteria:
        ProductName:
          - Comparison: EQUALS
            Value: GuardDuty
          - Comparison: PREFIX
            Value: SecurityHub
        CompanyName:
          - Comparison: EQUALS
            Value: AWS
          - Comparison: PREFIX
            Value: Private
        ProductArn:
          - Comparison: EQUALS
            Value: arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default
          - Comparison: PREFIX
            Value: arn:aws:securityhub:us-west-2:123456789012:product/aws
        AwsAccountId:
          - Comparison: EQUALS
            Value: 123456789012
        Id:
          - Comparison: EQUALS
            Value: example-finding-id
        GeneratorId:
          - Comparison: EQUALS
            Value: example-generator-id
        Type:
          - Comparison: EQUALS
            Value: type-1
          - Comparison: EQUALS
            Value: type-2
        Description:
          - Comparison: EQUALS
            Value: description1
          - Comparison: EQUALS
            Value: description2
        SourceUrl:
          - Comparison: PREFIX
            Value: https
          - Comparison: PREFIX
            Value: ftp
        Title:
          - Comparison: EQUALS
            Value: title-1
          - Comparison: PREFIX
            Value: title-2
        SeverityLabel:
          - Comparison: EQUALS
            Value: LOW
          - Comparison: EQUALS
            Value: HIGH
        ResourceType:
          - Comparison: EQUALS
            Value: AwsEc2Instance
        ResourcePartition:
          - Comparison: EQUALS
            Value: aws
        ResourceId:
          - Comparison: PREFIX
            Value: i-1234567890
        ResourceRegion:
          - Comparison: PREFIX
            Value: us-west
        ComplianceStatus:
          - Comparison: EQUALS
            Value: FAILED
        ComplianceSecurityControlId:
          - Comparison: EQUALS
            Value: EC2.3
        ComplianceAssociatedStandardsId:
          - Comparison: EQUALS
            Value: ruleset/cis-aws-foundations-benchmark/v/1.2.0
        VerificationState:
          - Comparison: EQUALS
            Value: BENIGN_POSITIVE
        RecordState:
          - Comparison: EQUALS
            Value: ACTIVE
        RelatedFindingsProductArn:
          - Comparison: EQUALS
            Value: arn:aws:securityhub:eu-central-1::product/aws/securityhub
        RelatedFindingsId:
          - Comparison: EQUALS
            Value: example-finding-id-2
        NoteText:
          - Comparison: EQUALS
            Value: example-note-text
        NoteUpdatedAt:
          - DateRange:
              Unit: DAYS
              Value: 5
        NoteUpdatedBy:
          - Comparison: PREFIX
            Value: sechub
        WorkflowStatus:
          - Comparison: EQUALS
            Value: NEW
        FirstObservedAt:
          - DateRange:
              Unit: DAYS
              Value: 5
        LastObservedAt:
          - DateRange:
              Unit: DAYS
              Value: 5
        CreatedAt:
          - DateRange:
              Unit: DAYS
              Value: 5
        UpdatedAt:
          - Start: "2023-04-25T17:05:54.832Z"
            End: "2023-05-25T17:05:54.832Z"
        ResourceTags:
          - Comparison: NOT_EQUALS
            Key: department
            Value: security
          - Comparison: NOT_EQUALS
            Key: department
            Value: operations
        UserDefinedFields:
          - Comparison: EQUALS
            Key: key1
            Value: security
          - Comparison: EQUALS
            Key: key2
            Value: operations
        ResourceDetailsOther:
          - Comparison: NOT_EQUALS
            Key: area
            Value: na
          - Comparison: NOT_EQUALS
            Key: department
            Value: sales
        Confidence:
          - Gte: 50
            Lte: 95
        Criticality:
          - Gte: 50
            Lte: 95
      Actions:
        - Type: FINDING_FIELDS_UPDATE
          FindingFieldsUpdate:
            Severity:
              Product: 50
              Label: MEDIUM
              Normalized: 60
            Types:
              - Software and Configuration Checks/Industry and Regulatory Standards/AWS-Foundational-Security-Best-Practices
              - Industry Compliance
            Confidence: 98
            Criticality: 95
            UserDefinedFields:
              key1: value1
              key2: value2
            RelatedFindings:
              - ProductArn: arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default
                Id: sample-finding-id-1
              - ProductArn: arn:aws:securityhub:us-west-2:123456789012:product/123456789012/default
                Id: sample-finding-id-2
            Note:
              Text: sample-note-text
              UpdatedBy: sechub
            VerificationState: TRUE_POSITIVE
            Workflow:
              Status: NOTIFIED
      Tags:
        sampleTag: sampleValue
        organizationUnit: pnw
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::AggregatorV2

AutomationRulesAction

All content copied from https://docs.aws.amazon.com/.
