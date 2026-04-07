This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Application AutoSubscriptionConfiguration

Subscription configuration information for an Amazon Q Business application
using IAM identity federation for user management.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoSubscribe" : String,
  "DefaultSubscriptionType" : String
}

```

### YAML

```yaml

  AutoSubscribe: String
  DefaultSubscriptionType: String

```

## Properties

`AutoSubscribe`

Describes whether automatic subscriptions are enabled for an Amazon Q Business
application using IAM identity federation for user management.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSubscriptionType`

Describes the default subscription type assigned to an Amazon Q Business
application using IAM identity federation for user management. If the
value for `autoSubscribe` is set to `ENABLED` you must select a
value for this field.

_Required_: No

_Type_: String

_Allowed values_: `Q_LITE | Q_BUSINESS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttachmentsConfiguration

EncryptionConfiguration
