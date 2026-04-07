This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule TopicRulePayload

Describes a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ Action, ... ],
  "AwsIotSqlVersion" : String,
  "Description" : String,
  "ErrorAction" : Action,
  "RuleDisabled" : Boolean,
  "Sql" : String
}

```

### YAML

```yaml

  Actions:
    - Action
  AwsIotSqlVersion: String
  Description: String
  ErrorAction:
    Action
  RuleDisabled: Boolean
  Sql: String

```

## Properties

`Actions`

The actions associated with the rule.

_Required_: Yes

_Type_: Array of [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-topicrule-action.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsIotSqlVersion`

The version of the SQL rules engine to use when evaluating the rule.

The default value is 2015-10-08.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the rule.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorAction`

The action to take when an error occurs.

_Required_: No

_Type_: [Action](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-topicrule-action.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleDisabled`

Specifies whether the rule is disabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sql`

The SQL statement used to query the topic. For more information, see [AWS IoT SQL\
Reference](https://docs.aws.amazon.com/iot/latest/developerguide/iot-sql-reference.html) in the _AWS IoT Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TimestreamTimestamp

UserProperty
