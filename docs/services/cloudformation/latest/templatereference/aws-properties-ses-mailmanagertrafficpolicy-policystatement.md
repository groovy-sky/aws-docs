This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy PolicyStatement

The structure containing traffic policy conditions and actions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "Conditions" : [ PolicyCondition, ... ]
}

```

### YAML

```yaml

  Action: String
  Conditions:
    - PolicyCondition

```

## Properties

`Action`

The action that informs a traffic policy resource to either allow or block the email
if it matches a condition in the policy statement.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Conditions`

The list of conditions to apply to incoming messages for filtering email
traffic.

_Required_: Yes

_Type_: Array of [PolicyCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanagertrafficpolicy-policycondition.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PolicyCondition

Tag
