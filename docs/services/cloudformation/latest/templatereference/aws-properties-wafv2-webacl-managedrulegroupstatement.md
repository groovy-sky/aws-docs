This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ManagedRuleGroupStatement

A rule statement used to run the rules that are defined in a managed rule group. To use
this, provide the vendor name and the name of the rule group in this statement. You can
retrieve the required names through the API call
`ListAvailableManagedRuleGroups`.

You cannot nest a `ManagedRuleGroupStatement`, for example for use inside a
`NotStatement` or `OrStatement`. You cannot use a managed rule
group statement inside another rule group. You can only use a managed rule group statement
as a top-level statement in a rule that you define in a web ACL.

###### Note

You are charged additional fees when you use the AWS WAF Bot Control
managed rule group `AWSManagedRulesBotControlRuleSet`, the AWS WAF Fraud Control account takeover prevention (ATP) managed rule group
`AWSManagedRulesATPRuleSet`, or the AWS WAF Fraud Control
account creation fraud prevention (ACFP) managed rule group
`AWSManagedRulesACFPRuleSet`. For more information, see [AWS WAF Pricing](https://aws.amazon.com/waf/pricing).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludedRules" : [ ExcludedRule, ... ],
  "ManagedRuleGroupConfigs" : [ ManagedRuleGroupConfig, ... ],
  "Name" : String,
  "RuleActionOverrides" : [ RuleActionOverride, ... ],
  "ScopeDownStatement" : Statement,
  "VendorName" : String,
  "Version" : String
}

```

### YAML

```yaml

  ExcludedRules:
    - ExcludedRule
  ManagedRuleGroupConfigs:
    - ManagedRuleGroupConfig
  Name: String
  RuleActionOverrides:
    - RuleActionOverride
  ScopeDownStatement:
    Statement
  VendorName: String
  Version: String

```

## Properties

`ExcludedRules`

Rules in the referenced rule group whose actions are set to `Count`.

###### Note

Instead of this option, use `RuleActionOverrides`. It accepts any valid action setting, including `Count`.

_Required_: No

_Type_: Array of [ExcludedRule](aws-properties-wafv2-webacl-excludedrule.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedRuleGroupConfigs`

Additional information that's used by a managed rule group. Many managed rule groups don't require this.

The rule groups used for intelligent threat mitigation require additional configuration:

- Use the `AWSManagedRulesACFPRuleSet` configuration object to configure the account creation fraud prevention managed rule group. The configuration includes the registration and sign-up pages of your application and the locations in the account creation request payload of data, such as the user email and phone number fields.

- Use the `AWSManagedRulesAntiDDoSRuleSet` configuration object to configure the anti-DDoS managed rule group. The configuration includes the sensitivity levels to use in the rules that typically block and challenge requests that might be participating in DDoS attacks and the specification to use to indicate whether a request can handle a silent browser challenge.

- Use the `AWSManagedRulesATPRuleSet` configuration object to configure the account takeover prevention managed rule group. The configuration includes the sign-in page of your application and the locations in the login request payload of data such as the username and password.

- Use the `AWSManagedRulesBotControlRuleSet` configuration object to configure the
protection level that you want the Bot Control rule group to use.

_Required_: No

_Type_: Array of [ManagedRuleGroupConfig](aws-properties-wafv2-webacl-managedrulegroupconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the managed rule group. You use this, along with the vendor name, to identify the rule group.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleActionOverrides`

Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change.

###### Note

Verify the rule names in your overrides carefully. With managed rule groups, AWS WAF silently ignores any override that uses an invalid rule name. With customer-owned rule groups, invalid rule names in your overrides will cause web ACL updates to fail. An invalid rule name is any name that doesn't exactly match the case-sensitive name of an existing rule in the rule group.

You can use overrides for testing, for example you can override all of rule actions to `Count` and then monitor the resulting count metrics to understand how the rule group would handle your web traffic. You can also permanently override some or all actions, to modify how the rule group manages your web traffic.

_Required_: No

_Type_: Array of [RuleActionOverride](aws-properties-wafv2-webacl-ruleactionoverride.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScopeDownStatement`

An optional nested statement that narrows the scope of the web requests that are
evaluated by the managed rule group. Requests are only evaluated by the rule group if they
match the scope-down statement. You can use any nestable [Statement](../userguide/aws-properties-wafv2-webacl-notstatement.md#cfn-wafv2-webacl-notstatement-statement) in the
scope-down statement, and you can nest statements at any level, the same as you can for a
rule statement.

_Required_: No

_Type_: [Statement](aws-properties-wafv2-webacl-statement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VendorName`

The name of the managed rule group vendor. You use this, along with the rule group name, to identify a rule group.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the managed rule group to use. If you specify this, the version setting
is fixed until you change it. If you don't specify this, AWS WAF uses the vendor's
default version, and then keeps the version at the vendor's default when the vendor updates
the managed rule group settings.

_Required_: No

_Type_: String

_Pattern_: `^[\w#:\.\-/]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Configure the managed rule group statement for AWSManagedRulesATPRuleSet](#aws-properties-wafv2-webacl-managedrulegroupstatement--examples--Configure_the_managed_rule_group_statement_for_AWSManagedRulesATPRuleSet)

- [Configure a standard managed rule group statement](#aws-properties-wafv2-webacl-managedrulegroupstatement--examples--Configure_a_standard_managed_rule_group_statement)

### Configure the managed rule group statement for `AWSManagedRulesATPRuleSet`

The following shows an example `ManagedRuleGroupStatement` for the
AWS WAF ATP managed rule group. The
`ManagedRuleGroupConfigs` settings are provided as a number of
individual `ManagedRuleGroupConfig` settings.

#### YAML

```yaml

ManagedRuleGroupStatement:
  VendorName: AWS
  Name: AWSManagedRulesATPRuleSet
  ManagedRuleGroupConfigs:
    - LoginPath: /api/accounts/login
    - PayloadType: JSON
    - PasswordField:
        Identifier: /form/password
    - UsernameField:
        Identifier: /form/username
```

#### JSON

```json

{
  "ManagedRuleGroupStatement": {
    "VendorName": "AWS",
    "Name": "AWSManagedRulesATPRuleSet",
    "ManagedRuleGroupConfigs": [
      {
        "LoginPath": "/api/accounts/login"
      },
      {
        "PayloadType": "JSON"
      },
      {
        "PasswordField": {
          "Identifier": "/form/password"
        }
      },
      {
        "UsernameField": {
          "Identifier": "/form/username"
        }
      }
    ]
  }
}
```

### Configure a standard managed rule group statement

The following shows an example `ManagedRuleGroupStatement` for a
managed rule group that doesn't require additional configuration.

#### YAML

```yaml

ManagedRuleGroupStatement:
  VendorName: AWS
  Name: AWSManagedRulesCommonRuleSet
```

#### JSON

```json

{
  "ManagedRuleGroupStatement": {
    "VendorName": "AWS",
    "Name": "AWSManagedRulesCommonRuleSet"
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedRuleGroupConfig

NotStatement

All content copied from https://docs.aws.amazon.com/.
