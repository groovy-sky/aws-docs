This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker LogList

The list of information about logs to be enabled for the specified broker.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Audit" : Boolean,
  "General" : Boolean
}

```

### YAML

```yaml

  Audit: Boolean
  General: Boolean

```

## Properties

`Audit`

Enables audit logging. Every user management action made using JMX or the ActiveMQ
Web Console is logged. Does not apply to RabbitMQ brokers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`General`

Enables general logging.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LdapServerMetadata

MaintenanceWindow
