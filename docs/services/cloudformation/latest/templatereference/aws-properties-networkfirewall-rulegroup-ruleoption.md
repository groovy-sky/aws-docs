This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup RuleOption

Additional settings for a stateful rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Keyword" : String,
  "Settings" : [ String, ... ]
}

```

### YAML

```yaml

  Keyword: String
  Settings:
    - String

```

## Properties

`Keyword`

The Suricata rule option keywords. For Network Firewall, the keyword signature ID (sid) is required in the format `sid:112233`. The sid must be unique within the rule group. For information about Suricata rule option keywords, see [Rule options](https://suricata.readthedocs.io/en/suricata-6.0.9/rules/intro.html).

_Required_: Yes

_Type_: String

_Pattern_: `^.*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

The Suricata rule option settings. Settings have zero or more values, and the number of possible settings and required settings depends on the keyword. The format for Settings is `number`. For information about Suricata rule option settings, see [Rule options](https://suricata.readthedocs.io/en/suricata-6.0.9/rules/intro.html).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleGroup

RulesSource

All content copied from https://docs.aws.amazon.com/.
