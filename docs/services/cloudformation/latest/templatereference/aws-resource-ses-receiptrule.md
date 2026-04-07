This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRule

Specifies a receipt rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ReceiptRule",
  "Properties" : {
      "After" : String,
      "Rule" : Rule,
      "RuleSetName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::ReceiptRule
Properties:
  After: String
  Rule:
    Rule
  RuleSetName: String

```

## Properties

`After`

The name of an existing rule after which the new rule is placed. If this parameter is
null, the new rule is inserted at the beginning of the rule list.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rule`

A data structure that contains the specified rule's name, actions, recipients,
domains, enabled status, scan status, and TLS policy.

_Required_: Yes

_Type_: [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-receiptrule-rule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleSetName`

The name of the rule set where the receipt rule is added.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

Specifies a receipt rule for incoming email.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES ReceiptRule Sample Template",
    "Parameters": {
        "RuleSetName": {
            "Type": "String"
        },
        "ReceiptRuleName1": {
            "Type": "String"
        },
        "ReceiptRuleName2": {
            "Type": "String"
        },
        "TlsPolicy": {
            "Type": "String"
        },
        "HeaderName": {
            "Type": "String"
        },
        "HeaderValue": {
            "Type": "String"
        }
    },
    "Resources": {
        "ReceiptRule1": {
            "Type": "AWS::SES::ReceiptRule",
            "Properties": {
                "RuleSetName": {
                    "Ref": "RuleSetName"
                },
                "Rule": {
                    "Name": {
                        "Ref": "ReceiptRuleName1"
                    },
                    "Enabled": true,
                    "ScanEnabled": true,
                    "TlsPolicy": {
                        "Ref": "TlsPolicy"
                    },
                    "Actions": [
                        {
                            "AddHeaderAction": {
                                "HeaderName": {
                                    "Ref": "HeaderName"
                                },
                                "HeaderValue": {
                                    "Ref": "HeaderValue"
                                }
                            }
                        }
                    ]
                }
            }
        },
        "ReceiptRule2": {
            "Type": "AWS::SES::ReceiptRule",
            "Properties": {
                "RuleSetName": {
                    "Ref": "RuleSetName"
                },
                "After": {
                    "Ref": "ReceiptRule1"
                },
                "Rule": {
                    "Name": {
                        "Ref": "ReceiptRuleName2"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES ReceiptRule Sample Template
Parameters:
  RuleSetName:
    Type: String
  ReceiptRuleName1:
    Type: String
  ReceiptRuleName2:
    Type: String
  TlsPolicy:
    Type: String
  HeaderName:
    Type: String
  HeaderValue:
    Type: String
Resources:
  ReceiptRule1:
    Type: 'AWS::SES::ReceiptRule'
    Properties:
      RuleSetName: !Ref RuleSetName
      Rule:
        Name: !Ref ReceiptRuleName1
        Enabled: true
        ScanEnabled: true
        TlsPolicy: !Ref TlsPolicy
        Actions:
          - AddHeaderAction:
              HeaderName: !Ref HeaderName
              HeaderValue: !Ref HeaderValue
  ReceiptRule2:
    Type: 'AWS::SES::ReceiptRule'
    Properties:
      RuleSetName: !Ref RuleSetName
      After: !Ref ReceiptRule1
      Rule:
        Name: !Ref ReceiptRuleName2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IpFilter

Action
