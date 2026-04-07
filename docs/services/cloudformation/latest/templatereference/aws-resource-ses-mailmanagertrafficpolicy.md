This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy

Resource to create a traffic policy for a Mail Manager ingress endpoint which contains
policy statements used to evaluate whether incoming emails should be allowed or
denied.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerTrafficPolicy",
  "Properties" : {
      "DefaultAction" : String,
      "MaxMessageSizeBytes" : Number,
      "PolicyStatements" : [ PolicyStatement, ... ],
      "Tags" : [ Tag, ... ],
      "TrafficPolicyName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerTrafficPolicy
Properties:
  DefaultAction: String
  MaxMessageSizeBytes: Number
  PolicyStatements:
    - PolicyStatement
  Tags:
    - Tag
  TrafficPolicyName: String

```

## Properties

`DefaultAction`

Default action instructs the traﬃc policy to either Allow or Deny (block) messages that fall outside of (or not addressed by) the conditions of your policy statements

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxMessageSizeBytes`

The maximum message size in bytes of email which is allowed in by this traffic
policy—anything larger will be blocked.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyStatements`

Conditional statements for filtering email traffic.

_Required_: Yes

_Type_: Array of [PolicyStatement](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanagertrafficpolicy-policystatement.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanagertrafficpolicy-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficPolicyName`

The name of the policy.

The policy name cannot exceed 64 characters and can only include alphanumeric
characters, dashes, and underscores.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_\-]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`TrafficPolicyArn`

The Amazon Resource Name (ARN) of the traffic policy resource.

`TrafficPolicyId`

The identifier of the traffic policy resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

IngressAnalysis
