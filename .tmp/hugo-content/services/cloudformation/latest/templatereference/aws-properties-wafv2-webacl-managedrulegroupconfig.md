This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ManagedRuleGroupConfig

Additional information that's used by a managed rule group. Many managed rule groups
don't require this.

The rule groups used for intelligent threat mitigation require additional configuration:

- Use the `AWSManagedRulesACFPRuleSet` configuration object to configure
the account creation fraud prevention managed rule group. The configuration includes
the registration and sign-up pages of your application and the locations in the
account creation request payload of data, such as the user email and phone number
fields.

- Use the `AWSManagedRulesATPRuleSet` configuration object to configure
the account takeover prevention managed rule group. The configuration includes the
sign-in page of your application and the locations in the login request payload of
data such as the username and password.

- Use the `AWSManagedRulesBotControlRuleSet` configuration object to
configure the protection level that you want the Bot Control rule group to use.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AWSManagedRulesACFPRuleSet" : AWSManagedRulesACFPRuleSet,
  "AWSManagedRulesAntiDDoSRuleSet" : AWSManagedRulesAntiDDoSRuleSet,
  "AWSManagedRulesATPRuleSet" : AWSManagedRulesATPRuleSet,
  "AWSManagedRulesBotControlRuleSet" : AWSManagedRulesBotControlRuleSet,
  "LoginPath" : String,
  "PasswordField" : FieldIdentifier,
  "PayloadType" : String,
  "UsernameField" : FieldIdentifier
}

```

### YAML

```yaml

  AWSManagedRulesACFPRuleSet:
    AWSManagedRulesACFPRuleSet
  AWSManagedRulesAntiDDoSRuleSet:
    AWSManagedRulesAntiDDoSRuleSet
  AWSManagedRulesATPRuleSet:
    AWSManagedRulesATPRuleSet
  AWSManagedRulesBotControlRuleSet:
    AWSManagedRulesBotControlRuleSet
  LoginPath: String
  PasswordField:
    FieldIdentifier
  PayloadType: String
  UsernameField:
    FieldIdentifier

```

## Properties

`AWSManagedRulesACFPRuleSet`

Additional configuration for using the account creation fraud prevention (ACFP) managed rule group, `AWSManagedRulesACFPRuleSet`.
Use this to provide account creation request information to the rule group. For web ACLs that protect CloudFront distributions, use this to also provide
the information about how your distribution responds to account creation requests.

For information
about using the ACFP managed rule group, see [AWS WAF Fraud Control account creation fraud prevention (ACFP) rule group](../../../waf/latest/developerguide/aws-managed-rule-groups-acfp.md)
and [AWS WAF Fraud Control account creation fraud prevention (ACFP)](../../../waf/latest/developerguide/waf-acfp.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [AWSManagedRulesACFPRuleSet](aws-properties-wafv2-webacl-awsmanagedrulesacfpruleset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AWSManagedRulesAntiDDoSRuleSet`

Additional configuration for using the anti-DDoS managed rule group, `AWSManagedRulesAntiDDoSRuleSet`.
Use this to configure anti-DDoS behavior for the rule group.

For information
about using the anti-DDoS managed rule group, see [AWS WAF Anti-DDoS rule group](../../../waf/latest/developerguide/aws-managed-rule-groups-anti-ddos.md)
and [Distributed Denial of Service (DDoS) prevention](../../../waf/latest/developerguide/waf-anti-ddos.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [AWSManagedRulesAntiDDoSRuleSet](aws-properties-wafv2-webacl-awsmanagedrulesantiddosruleset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AWSManagedRulesATPRuleSet`

Additional configuration for using the account takeover prevention (ATP) managed rule group, `AWSManagedRulesATPRuleSet`.
Use this to provide login request information to the rule group. For web ACLs that protect CloudFront distributions, use this to also provide
the information about how your distribution responds to login requests.

This configuration replaces the individual configuration fields in `ManagedRuleGroupConfig` and provides additional feature configuration.

For information
about using the ATP managed rule group, see [AWS WAF Fraud Control account takeover prevention (ATP) rule group](../../../waf/latest/developerguide/aws-managed-rule-groups-atp.md)
and [AWS WAF Fraud Control account takeover prevention (ATP)](../../../waf/latest/developerguide/waf-atp.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [AWSManagedRulesATPRuleSet](aws-properties-wafv2-webacl-awsmanagedrulesatpruleset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AWSManagedRulesBotControlRuleSet`

Additional configuration for using the Bot Control managed rule group. Use this to specify the
inspection level that you want to use. For information
about using the Bot Control managed rule group, see [AWS WAF Bot Control rule group](../../../waf/latest/developerguide/aws-managed-rule-groups-bot.md)
and [AWS WAF Bot Control](../../../waf/latest/developerguide/waf-bot-control.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [AWSManagedRulesBotControlRuleSet](aws-properties-wafv2-webacl-awsmanagedrulesbotcontrolruleset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginPath`

###### Note

Instead of this setting, provide your configuration under `AWSManagedRulesATPRuleSet`.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PasswordField`

###### Note

Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet`.

_Required_: No

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadType`

###### Note

Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet`.

_Required_: No

_Type_: String

_Allowed values_: `JSON | FORM_ENCODED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsernameField`

###### Note

Instead of this setting, provide your configuration under the request inspection configuration for `AWSManagedRulesATPRuleSet` or `AWSManagedRulesACFPRuleSet`.

_Required_: No

_Type_: [FieldIdentifier](aws-properties-wafv2-webacl-fieldidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelMatchStatement

ManagedRuleGroupStatement

All content copied from https://docs.aws.amazon.com/.
