This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet

Resource to create a rule set for a Mail Manager ingress endpoint which contains a
list of rules that are evaluated sequentially for each email.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerRuleSet",
  "Properties" : {
      "Rules" : [ Rule, ... ],
      "RuleSetName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerRuleSet
Properties:
  Rules:
    - Rule
  RuleSetName: String
  Tags:
    - Tag

```

## Properties

`Rules`

Conditional rules that are evaluated for determining actions on email.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-ses-mailmanagerruleset-rule.md)

_Minimum_: `0`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleSetName`

A user-friendly name for the rule set.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-mailmanagerruleset-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RuleSetArn`

The Amazon Resource Name (ARN) of the rule set resource.

`RuleSetId`

The identifier of the rule set.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AddHeaderAction

All content copied from https://docs.aws.amazon.com/.
