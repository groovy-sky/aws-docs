This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptRuleSet

Creates an empty receipt rule set.

For information about setting up receipt rule sets, see the [Amazon SES Developer Guide](https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules).

You can execute this operation no more than once per second.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ReceiptRuleSet",
  "Properties" : {
      "RuleSetName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::ReceiptRuleSet
Properties:
  RuleSetName: String

```

## Properties

`RuleSetName`

The name of the receipt rule set to make active. Setting this value to null disables
all email receiving.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

Specifies a collection of receipt rules that are applied to incoming email.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS SES ReceiptRuleSet Sample Template",
    "Parameters": {
        "ReceiptRuleSetName": {
            "Type": "String"
        }
    },
    "Resources": {
        "ReceiptRuleSet": {
            "Type": "AWS::SES::ReceiptRuleSet",
            "Properties": {
                "RuleSetName": {
                    "Ref": "ReceiptRuleSetName"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: AWS SES ReceiptRuleSet Sample Template
Parameters:
  ReceiptRuleSetName:
    Type: String
Resources:
  ReceiptRuleSet:
    Type: 'AWS::SES::ReceiptRuleSet'
    Properties:
      RuleSetName: !Ref ReceiptRuleSetName
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkmailAction

AWS::SES::Template
