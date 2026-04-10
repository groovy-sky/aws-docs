This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker ConfigurationId

A list of information about the configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Revision" : Integer
}

```

### YAML

```yaml

  Id: String
  Revision: Integer

```

## Properties

`Id`

Required. The unique ID that Amazon MQ generates for the configuration.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Revision`

The revision number of the configuration.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AmazonMQ::Broker

EncryptionOptions

All content copied from https://docs.aws.amazon.com/.
