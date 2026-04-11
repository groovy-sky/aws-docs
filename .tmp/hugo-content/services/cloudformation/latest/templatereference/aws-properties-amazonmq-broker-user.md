This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker User

The list of broker users (persons or applications) who can access queues and topics.
For Amazon MQ for RabbitMQ brokers, one and only one administrative user is
accepted and created when a broker is first provisioned. All subsequent broker users are created by making
RabbitMQ API calls directly to brokers or via the RabbitMQ web console.

When OAuth 2.0 is enabled, the broker accepts one or no users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConsoleAccess" : Boolean,
  "Groups" : [ String, ... ],
  "Password" : String,
  "ReplicationUser" : Boolean,
  "Username" : String
}

```

### YAML

```yaml

  ConsoleAccess: Boolean
  Groups:
    - String
  Password: String
  ReplicationUser: Boolean
  Username: String

```

## Properties

`ConsoleAccess`

Enables access to the ActiveMQ Web Console for the ActiveMQ user. Does not apply to RabbitMQ brokers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Groups`

The list of groups (20 maximum) to which the ActiveMQ user belongs. This value can
contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_
~). This value must be 2-100 characters long. Does not apply to RabbitMQ brokers.

_Required_: No

_Type_: Array of String

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Password`

Required. The password of the user. This value must be at least 12 characters
long, must contain at least 4 unique characters, and must not contain commas, colons, or equal signs (,:=).

_Required_: Yes

_Type_: String

_Pattern_: `^[^,:=]+$`

_Minimum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationUser`

Defines if this user is intended for CRDR replication purposes.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The username of the broker user. The following restrictions apply to broker usernames:

- For Amazon MQ for ActiveMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores, and tildes (- . \_ ~).
This value must be 2-100 characters long.

- For Amazon MQ for RabbitMQ brokers, this value can contain only alphanumeric characters, dashes, periods, underscores (- . \_).
This value must not contain a tilde (~) character. Amazon MQ prohibts using `guest` as a valid usename.
This value must be 2-100 characters long.

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames.
Broker usernames are accessible to other AWS services, including CloudWatch Logs. Broker usernames are not intended to be
used for private or sensitive data.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9_.~-]{2,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagsEntry

AWS::AmazonMQ::Configuration

All content copied from https://docs.aws.amazon.com/.
