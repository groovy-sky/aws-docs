This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksSecret

Specifies the configuration of a Kubernetes `secret` volume. For more information, see
[secret](https://kubernetes.io/docs/concepts/storage/volumes) in the
_Kubernetes documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Optional" : Boolean,
  "SecretName" : String
}

```

### YAML

```yaml

  Optional: Boolean
  SecretName: String

```

## Properties

`Optional`

Specifies whether the secret or the secret's keys must be defined.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretName`

The name of the secret. The name must be allowed as a DNS subdomain name. For more
information, see [DNS subdomain names](https://kubernetes.io/docs/concepts/overview/working-with-objects/names) in the _Kubernetes documentation_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksProperties

EksVolume

All content copied from https://docs.aws.amazon.com/.
