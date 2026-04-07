This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ContactList Topic

An interest group, theme, or label within a list. Lists can have multiple
topics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultSubscriptionStatus" : String,
  "Description" : String,
  "DisplayName" : String,
  "TopicName" : String
}

```

### YAML

```yaml

  DefaultSubscriptionStatus: String
  Description: String
  DisplayName: String
  TopicName: String

```

## Properties

`DefaultSubscriptionStatus`

The default subscription status to be applied to a contact if the contact has not
noted their preference for subscribing to a topic.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of what the topic is about, which the contact will see.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the topic the contact will see.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopicName`

The name of the topic.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,64}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::SES::CustomVerificationEmailTemplate
