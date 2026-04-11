This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster EncryptionInTransit

The settings for encrypting data in transit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientBroker" : String,
  "InCluster" : Boolean
}

```

### YAML

```yaml

  ClientBroker: String
  InCluster: Boolean

```

## Properties

`ClientBroker`

Indicates the encryption setting for data in transit between clients and brokers. You must set it to one of the following values.

- `TLS`: Indicates that client-broker communication is enabled with TLS only.

- `TLS_PLAINTEXT`: Indicates that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.

- `PLAINTEXT`: Indicates that client-broker communication is enabled in plaintext only.

The default value is `TLS`.

_Required_: No

_Type_: String

_Allowed values_: `TLS | TLS_PLAINTEXT | PLAINTEXT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InCluster`

When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted. When set to false, the communication happens in plaintext.

The default value is true.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionInfo

Firehose

All content copied from https://docs.aws.amazon.com/.
