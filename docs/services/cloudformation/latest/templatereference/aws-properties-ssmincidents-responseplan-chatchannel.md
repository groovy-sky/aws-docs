This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan ChatChannel

The Amazon Q Developer in chat applications chat channel used for collaboration during an
incident.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChatbotSns" : [ String, ... ]
}

```

### YAML

```yaml

  ChatbotSns:
    - String

```

## Properties

`ChatbotSns`

The Amazon SNS targets that Amazon Q Developer in chat applications uses to notify the chat channel of updates
to an incident. You can also make updates to the incident through the chat channel by
using the Amazon SNS topics

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Action

DynamicSsmParameter
