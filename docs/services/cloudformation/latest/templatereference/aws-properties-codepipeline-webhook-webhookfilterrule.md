This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Webhook WebhookFilterRule

The event criteria that specify when a webhook notification is sent to your
URL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonPath" : String,
  "MatchEquals" : String
}

```

### YAML

```yaml

  JsonPath: String
  MatchEquals: String

```

## Properties

`JsonPath`

A JsonPath expression that is applied to the body/payload of the webhook. The value
selected by the JsonPath expression must match the value specified in the
`MatchEquals` field. Otherwise, the request is ignored. For more
information, see [Java JsonPath\
implementation](https://github.com/json-path/JsonPath) in GitHub.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchEquals`

The value selected by the `JsonPath` expression must match what is
supplied in the `MatchEquals` field. Otherwise, the request is ignored.
Properties from the target action configuration can be included as placeholders in this
value by surrounding the action configuration key with curly brackets. For example, if
the value supplied here is "refs/heads/{Branch}" and the target action has an action
configuration property called "Branch" with a value of "main", the
`MatchEquals` value is evaluated as "refs/heads/main". For a list of
action configuration properties for built-in action types, see [Pipeline Structure Reference Action Requirements](https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WebhookAuthConfiguration

Next
