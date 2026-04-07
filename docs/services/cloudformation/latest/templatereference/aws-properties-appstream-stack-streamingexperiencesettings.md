This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Stack StreamingExperienceSettings

The streaming protocol that you want your stack to prefer. This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PreferredProtocol" : String
}

```

### YAML

```yaml

  PreferredProtocol: String

```

## Properties

`PreferredProtocol`

The preferred protocol that you want to use while streaming your application.

_Required_: No

_Type_: String

_Allowed values_: `TCP | UDP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StorageConnector

Tag
