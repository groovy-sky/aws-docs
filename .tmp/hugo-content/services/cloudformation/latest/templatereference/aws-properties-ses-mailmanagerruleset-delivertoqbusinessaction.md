This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet DeliverToQBusinessAction

The action to deliver incoming emails to an Amazon Q Business application for indexing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionFailurePolicy" : String,
  "ApplicationId" : String,
  "IndexId" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  ActionFailurePolicy: String
  ApplicationId: String
  IndexId: String
  RoleArn: String

```

## Properties

`ActionFailurePolicy`

A policy that states what to do in the case of failure. The action will fail if there are
configuration errors. For example, the specified application has been deleted or the role lacks
necessary permissions to call the `qbusiness:BatchPutDocument` API.

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | DROP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplicationId`

The unique identifier of the Amazon Q Business application instance where the email
content will be delivered.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexId`

The identifier of the knowledge base index within the Amazon Q Business application
where the email content will be stored and indexed.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9-]+$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM Role to use while delivering to Amazon Q Business. This role must have access
to the `qbusiness:BatchPutDocument` API for the given application and index.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9:_/+=,@.#-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeliverToMailboxAction

RelayAction

All content copied from https://docs.aws.amazon.com/.
