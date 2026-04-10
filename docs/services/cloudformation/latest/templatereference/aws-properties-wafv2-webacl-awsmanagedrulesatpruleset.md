This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AWSManagedRulesATPRuleSet

Details for your use of the account takeover prevention managed rule group, `AWSManagedRulesATPRuleSet`. This configuration is used in `ManagedRuleGroupConfig`.

For additional information about this and the other intelligent threat mitigation rule groups,
see [Intelligent threat mitigation in AWS WAF](../../../waf/latest/developerguide/waf-managed-protections.md)
and [AWS Managed Rules rule groups list](../../../waf/latest/developerguide/aws-managed-rule-groups-list.md)
in the _AWS WAF Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableRegexInPath" : Boolean,
  "LoginPath" : String,
  "RequestInspection" : RequestInspection,
  "ResponseInspection" : ResponseInspection
}

```

### YAML

```yaml

  EnableRegexInPath: Boolean
  LoginPath: String
  RequestInspection:
    RequestInspection
  ResponseInspection:
    ResponseInspection

```

## Properties

`EnableRegexInPath`

Allow the use of regular expressions in the login page path.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginPath`

The path of the login endpoint for your application. For example, for the URL
`https://example.com/web/login`, you would provide the path
`/web/login`. Login paths that start with the path that you provide are considered a match. For example `/web/login` matches the login paths `/web/login`, `/web/login/`, `/web/loginPage`, and `/web/login/thisPage`, but doesn't match the login path `/home/web/login` or `/website/login`.

The rule group inspects only HTTP `POST` requests to your specified login endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RequestInspection`

The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage.

_Required_: No

_Type_: [RequestInspection](aws-properties-wafv2-webacl-requestinspection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseInspection`

The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates.

###### Note

Response inspection is available only in web ACLs that protect Amazon CloudFront distributions.

The ATP rule group evaluates the responses that your protected resources send back to client login attempts, keeping count of successful and failed attempts for each IP address and client session. Using this information, the rule group labels
and mitigates requests from client sessions and IP addresses that have had too many failed login attempts in a short amount of time.

_Required_: No

_Type_: [ResponseInspection](aws-properties-wafv2-webacl-responseinspection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSManagedRulesAntiDDoSRuleSet

AWSManagedRulesBotControlRuleSet

All content copied from https://docs.aws.amazon.com/.
