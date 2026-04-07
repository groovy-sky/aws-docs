This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSet DeliveryOptions

Used to associate a configuration set with a dedicated IP pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SendingPoolName" : String
}

```

### YAML

```yaml

  SendingPoolName: String

```

## Properties

`SendingPoolName`

The name of the dedicated IP pool that you want to associate with the configuration
set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::PinpointEmail::ConfigurationSet

ReputationOptions
