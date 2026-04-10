This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AWSManagedRulesACFPRuleSet

Details for your use of the account creation fraud prevention managed rule group, `AWSManagedRulesACFPRuleSet`. This configuration is used in `ManagedRuleGroupConfig`.

For additional information about this and the other intelligent threat mitigation rule groups,
see [Intelligent threat mitigation in AWS WAF](../../../waf/latest/developerguide/waf-managed-protections.md)
and [AWS Managed Rules rule groups list](../../../waf/latest/developerguide/aws-managed-rule-groups-list.md)
in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreationPath" : String,
  "EnableRegexInPath" : Boolean,
  "RegistrationPagePath" : String,
  "RequestInspection" : RequestInspectionACFP,
  "ResponseInspection" : ResponseInspection
}

```

### YAML

```yaml

  CreationPath: String
  EnableRegexInPath: Boolean
  RegistrationPagePath: String
  RequestInspection:
    RequestInspectionACFP
  ResponseInspection:
    ResponseInspection

```

## Properties

`CreationPath`

The path of the account creation endpoint for your application. This is the page on your website that accepts the completed registration form for a new user. This page must accept `POST` requests.

For example, for the URL `https://example.com/web/newaccount`, you would provide
the path `/web/newaccount`. Account creation page paths that
start with the path that you provide are considered a match. For example
`/web/newaccount` matches the account creation paths
`/web/newaccount`, `/web/newaccount/`,
`/web/newaccountPage`, and
`/web/newaccount/thisPage`, but doesn't match the path
`/home/web/newaccount` or
`/website/newaccount`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableRegexInPath`

Allow the use of regular expressions in the registration page path and the account creation path.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistrationPagePath`

The path of the account registration endpoint for your application. This is the page on your website that presents the registration form to new users.

###### Note

This page must accept `GET` text/html requests.

For example, for the URL `https://example.com/web/registration`, you would provide
the path `/web/registration`. Registration page paths that
start with the path that you provide are considered a match. For example
`/web/registration` matches the registration paths
`/web/registration`, `/web/registration/`,
`/web/registrationPage`, and
`/web/registration/thisPage`, but doesn't match the path
`/home/web/registration` or
`/website/registration`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestInspection`

The criteria for inspecting account creation requests, used by the ACFP rule group to validate and track account creation attempts.

_Required_: Yes

_Type_: [RequestInspectionACFP](aws-properties-wafv2-webacl-requestinspectionacfp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseInspection`

The criteria for inspecting responses to account creation requests, used by the ACFP rule group to track account creation success rates.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

The ACFP rule group evaluates the responses that your protected resources send back to client account creation attempts, keeping count of successful and failed attempts from each IP address and client session. Using this information, the rule group labels
and mitigates requests from client sessions and IP addresses that have had too many successful account creation attempts in a short amount of time.

_Required_: No

_Type_: [ResponseInspection](aws-properties-wafv2-webacl-responseinspection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssociationConfig

AWSManagedRulesAntiDDoSRuleSet

All content copied from https://docs.aws.amazon.com/.
