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
policies](../../../iot/latest/developerguide/transport-security.md#tls-policy-table) in the _AWS IoT Core developer guide_.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoT::EncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
