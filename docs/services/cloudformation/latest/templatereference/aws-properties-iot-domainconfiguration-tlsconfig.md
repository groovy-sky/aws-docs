This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::DomainConfiguration TlsConfig

An object that specifies the TLS configuration for a domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityPolicy" : String
}

```

### YAML

```yaml

  SecurityPolicy: String

```

## Properties

`SecurityPolicy`

The security policy for a domain configuration. For more information, see [Security\
policies](https://docs.aws.amazon.com/iot/latest/developerguide/transport-security.html#tls-policy-table) in the _AWS IoT Core developer guide_.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::IoT::EncryptionConfiguration
