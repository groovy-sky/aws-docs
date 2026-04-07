This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::DeliveryDestination DestinationPolicy

An IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account
to a specified destination in this account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeliveryDestinationName" : String,
  "DeliveryDestinationPolicy" : Json
}

```

### YAML

```yaml

  DeliveryDestinationName: String
  DeliveryDestinationPolicy: Json

```

## Properties

`DeliveryDestinationName`

A name for an existing destination.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryDestinationPolicy`

Creates or updates an access policy associated with an existing destination. An access
policy is an [IAM\
policy document](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) that is used to authorize claims to register a subscription filter
against a given destination.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Logs::DeliveryDestination

Tag
