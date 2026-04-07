This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::OrganizationConfigRule

Adds or updates an AWS Config rule for your entire organization to evaluate if your AWS resources comply with your
desired configurations. For information on how many organization AWS Config rules you can have per account,
see [**Service Limits**](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the _AWS Config Developer Guide_.

Only a management account and a delegated administrator can create or update an organization AWS Config rule.
When calling the `OrganizationConfigRule` resource with a delegated administrator, you must ensure AWS Organizations `ListDelegatedAdministrator` permissions are added. An organization can have up to 3 delegated administrators.

The `OrganizationConfigRule` resource enables organization service access through the `EnableAWSServiceAccess` action and creates a service-linked
role `AWSServiceRoleForConfigMultiAccountSetup` in the management or delegated administrator account of your organization.
The service-linked role is created only when the role does not exist in the caller account.
AWS Config verifies the existence of role with `GetRole` action.

To use the `OrganizationConfigRule` resource with delegated administrator, register a delegated administrator by calling AWS Organization
`register-delegated-administrator` for `config-multiaccountsetup.amazonaws.com`.

There are two types of rules: _AWS Config Managed Rules_ and _AWS Config Custom Rules_.
You can use `PutOrganizationConfigRule` to create both AWS Config Managed Rules and AWS Config Custom Rules.

AWS Config Managed Rules are predefined,
customizable rules created by AWS Config. For a list of managed rules, see
[List of AWS Config\
Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an AWS Config managed rule, you must specify the rule's identifier for the `RuleIdentifier` key.

AWS Config Custom Rules are rules that you create from scratch. There are two ways to create AWS Config custom rules: with Lambda functions
( [AWS Lambda Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function)) and with Guard ( [Guard GitHub\
Repository](https://github.com/aws-cloudformation/cloudformation-guard)), a policy-as-code language.

AWS Config custom rules created with AWS Lambda
are called _AWS Config Custom Lambda Rules_ and AWS Config custom rules created with
Guard are called _AWS Config Custom Policy Rules_.

If you are adding a new AWS Config Custom Lambda rule, you first need to create an AWS Lambda function in the management account or a delegated
administrator that the rule invokes to evaluate your resources. You also need to create an IAM role in the managed account that can be assumed by the Lambda function.
When you use `PutOrganizationConfigRule` to add a Custom Lambda rule to AWS Config, you must
specify the Amazon Resource Name (ARN) that AWS Lambda assigns to the function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::OrganizationConfigRule",
  "Properties" : {
      "ExcludedAccounts" : [ String, ... ],
      "OrganizationConfigRuleName" : String,
      "OrganizationCustomPolicyRuleMetadata" : OrganizationCustomPolicyRuleMetadata,
      "OrganizationCustomRuleMetadata" : OrganizationCustomRuleMetadata,
      "OrganizationManagedRuleMetadata" : OrganizationManagedRuleMetadata
    }
}

```

### YAML

```yaml

Type: AWS::Config::OrganizationConfigRule
Properties:
  ExcludedAccounts:
    - String
  OrganizationConfigRuleName: String
  OrganizationCustomPolicyRuleMetadata:
    OrganizationCustomPolicyRuleMetadata
  OrganizationCustomRuleMetadata:
    OrganizationCustomRuleMetadata
  OrganizationManagedRuleMetadata:
    OrganizationManagedRuleMetadata

```

## Properties

`ExcludedAccounts`

A comma-separated list of accounts excluded from organization AWS Config rule.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationConfigRuleName`

The name that you assign to organization AWS Config rule.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9-_]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrganizationCustomPolicyRuleMetadata`

An
object that specifies metadata for your organization's AWS Config Custom Policy rule. The metadata includes the runtime system in use, which accounts have
debug logging enabled, and other custom rule metadata, such as resource type, resource
ID of AWS resource, and organization trigger types that initiate AWS Config to evaluate AWS resources against a rule.

_Required_: No

_Type_: [OrganizationCustomPolicyRuleMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-organizationconfigrule-organizationcustompolicyrulemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationCustomRuleMetadata`

An `OrganizationCustomRuleMetadata` object.

_Required_: No

_Type_: [OrganizationCustomRuleMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-organizationconfigrule-organizationcustomrulemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationManagedRuleMetadata`

An `OrganizationManagedRuleMetadata` object.

_Required_: No

_Type_: [OrganizationManagedRuleMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the OrganizationConfigRuleName.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Managed Rule](#aws-resource-config-organizationconfigrule--examples--Managed_Rule)

- [Custom Rule](#aws-resource-config-organizationconfigrule--examples--Custom_Rule)

### Managed Rule

The following example creates a managed organization config rule.

#### JSON

```json

{
    "BasicOrganizationConfigRule": {
        "Type": "AWS::Config::OrganizationConfigRule",
        "Properties": {
            "OrganizationConfigRuleName": "OrganizationConfigRuleName",
            "OrganizationManagedRuleMetadata": {
                "RuleIdentifier": "CLOUD_TRAIL_ENABLED",
                "Description": "Cloudtrail enabled rule"
            },
            "ExcludedAccounts": [
                "accountId"
            ]
        }
    }
}
```

#### YAML

```yaml

BasicOrganizationConfigRule:
    Type: "AWS::Config::OrganizationConfigRule"
    Properties:
        OrganizationConfigRuleName: "OrganizationConfigRuleName"
        OrganizationManagedRuleMetadata:
            RuleIdentifier: "CLOUD_TRAIL_ENABLED"
            Description: "Cloudtrail enabled rule"
        ExcludedAccounts:
        - "accountId"
```

### Custom Rule

The following example creates a custom organization config rule.

#### JSON

```json

{
    "BasicOrganizationConfigRule": {
        "Type": "AWS::Config::OrganizationConfigRule",
        "Properties": {
            "OrganizationConfigRuleName": "OrganizationConfigRuleName",
            "OrganizationCustomRuleMetadata": {
                "LambdaFunctionArn": "CustomRuleLambdaArn",
                "OrganizationConfigRuleTriggerTypes": [
                    "ScheduledNotification"
                ]
            },
            "ExcludedAccounts": [
                "accountId"
            ]
        }
    }
}
```

#### YAML

```yaml

BasicOrganizationConfigRule:
    Type: "AWS::Config::OrganizationConfigRule"
    Properties:
        OrganizationConfigRuleName: "OrganizationConfigRuleName"
        OrganizationCustomRuleMetadata:
            LambdaFunctionArn: "CustomRuleLambdaArn"
            OrganizationConfigRuleTriggerTypes:
                - "ScheduledNotification"
            ExcludedAccounts:
            - "accountId"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigSnapshotDeliveryProperties

OrganizationCustomPolicyRuleMetadata
