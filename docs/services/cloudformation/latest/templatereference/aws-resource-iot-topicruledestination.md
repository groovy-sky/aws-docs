This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRuleDestination

A topic rule destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::TopicRuleDestination",
  "Properties" : {
      "HttpUrlProperties" : HttpUrlDestinationSummary,
      "Status" : String,
      "VpcProperties" : VpcDestinationProperties
    }
}

```

### YAML

```yaml

Type: AWS::IoT::TopicRuleDestination
Properties:
  HttpUrlProperties:
    HttpUrlDestinationSummary
  Status: String
  VpcProperties:
    VpcDestinationProperties

```

## Properties

`HttpUrlProperties`

Properties of the HTTP URL.

_Required_: No

_Type_: [HttpUrlDestinationSummary](aws-properties-iot-topicruledestination-httpurldestinationsummary.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

IN\_PROGRESS

A topic rule destination was created but has not been confirmed. You can set
status to `IN_PROGRESS` by calling
`UpdateTopicRuleDestination`. Calling
`UpdateTopicRuleDestination` causes a new confirmation challenge to
be sent to your confirmation endpoint.

ENABLED

Confirmation was completed, and traffic to this destination is allowed. You can
set status to `DISABLED` by calling
`UpdateTopicRuleDestination`.

DISABLED

Confirmation was completed, and traffic to this destination is not allowed. You
can set status to `ENABLED` by calling
`UpdateTopicRuleDestination`.

ERROR

Confirmation could not be completed; for example, if the confirmation timed
out. You can call `GetTopicRuleDestination` for details about the
error. You can set status to `IN_PROGRESS` by calling
`UpdateTopicRuleDestination`. Calling
`UpdateTopicRuleDestination` causes a new confirmation challenge to
be sent to your confirmation endpoint.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | IN_PROGRESS | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcProperties`

Properties of the virtual private cloud (VPC) connection.

_Required_: No

_Type_: [VpcDestinationProperties](aws-properties-iot-topicruledestination-vpcdestinationproperties.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the topic rule destination. For example:

`{ "Ref": "TopicRuleDestination" }`

A value similar to the following is returned:

`a1234567b89c012d3e4fg567hij8k9l01mno1p23q45678901rs234567890t1u2`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The topic rule destination URL.

`StatusReason`

Additional details or reason why the topic rule destination is in the current
status.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserProperty

HttpUrlDestinationSummary

All content copied from https://docs.aws.amazon.com/.
