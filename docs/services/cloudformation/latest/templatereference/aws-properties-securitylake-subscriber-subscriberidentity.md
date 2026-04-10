This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber SubscriberIdentity

Specify the AWS account ID and external ID that the subscriber will use to access source data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExternalId" : String,
  "Principal" : String
}

```

### YAML

```yaml

  ExternalId: String
  Principal: String

```

## Properties

`ExternalId`

The external ID is a unique identifier that the subscriber provides to you.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w+=,.@:/-]*$`

_Minimum_: `2`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

Principals can include accounts, users, roles, federated users, or AWS services.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9]{12}|[a-z0-9\.\-]*\.(amazonaws|amazon)\.com)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Source

Tag

All content copied from https://docs.aws.amazon.com/.
