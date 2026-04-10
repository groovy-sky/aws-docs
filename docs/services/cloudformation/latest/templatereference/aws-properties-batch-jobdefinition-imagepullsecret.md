This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition ImagePullSecret

References a Kubernetes secret resource. This name of the secret must start and end with an
alphanumeric character, is required to be lowercase, can include periods (.) and hyphens (-), and
can't contain more than 253 characters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

Provides a unique identifier for the `ImagePullSecret`. This object is required
when `EksPodProperties$imagePullSecrets` is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Host

JobTimeout

All content copied from https://docs.aws.amazon.com/.
