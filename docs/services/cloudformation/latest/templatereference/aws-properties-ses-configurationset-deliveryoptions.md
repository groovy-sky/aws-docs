This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet DeliveryOptions

Specifies the name of the dedicated IP pool to associate with the configuration set
and whether messages that use the configuration set are required to use Transport Layer
Security (TLS).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxDeliverySeconds" : Number,
  "SendingPoolName" : String,
  "TlsPolicy" : String
}

```

### YAML

```yaml

  MaxDeliverySeconds: Number
  SendingPoolName: String
  TlsPolicy: String

```

## Properties

`MaxDeliverySeconds`

The name of the configuration set used when sent through a configuration set with archiving
enabled.

_Required_: No

_Type_: Number

_Minimum_: `300`

_Maximum_: `50400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendingPoolName`

The name of the dedicated IP pool to associate with the configuration set.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TlsPolicy`

Specifies whether messages that use the configuration set are required to use
Transport Layer Security (TLS). If the value is `REQUIRE`, messages are only
delivered if a TLS connection can be established. If the value is `OPTIONAL`,
messages can be delivered in plain text if a TLS connection can't be established.

Valid Values: `REQUIRE | OPTIONAL`

_Required_: No

_Type_: String

_Pattern_: `REQUIRE|OPTIONAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DashboardOptions

GuardianOptions
