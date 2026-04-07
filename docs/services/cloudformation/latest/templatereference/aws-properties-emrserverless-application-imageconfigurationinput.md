This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application ImageConfigurationInput

The image configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImageUri" : String
}

```

### YAML

```yaml

  ImageUri: String

```

## Properties

`ImageUri`

The URI of an image in the Amazon ECR registry. This field is required when you create a
new application. If you leave this field blank in an update, Amazon EMR will remove the
image configuration.

_Required_: No

_Type_: String

_Pattern_: `^([a-z0-9]+[a-z0-9-.]*)\/((?:[a-z0-9]+(?:[._-][a-z0-9]+)*\/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)(?:\:([a-zA-Z0-9_][a-zA-Z0-9-._]{0,299})|@(sha256:[0-9a-f]{64}))$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IdentityCenterConfiguration

InitialCapacityConfig
