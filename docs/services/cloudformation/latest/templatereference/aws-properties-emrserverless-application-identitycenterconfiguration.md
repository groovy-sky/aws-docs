This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application IdentityCenterConfiguration

The IAM Identity Center Configuration accepts the Identity Center instance parameter required to enable trusted identity
propagation. This configuration allows identity propagation between integrated services and the Identity Center instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdentityCenterInstanceArn" : String
}

```

### YAML

```yaml

  IdentityCenterInstanceArn: String

```

## Properties

`IdentityCenterInstanceArn`

The ARN of the IAM Identity Center instance.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationObject

ImageConfigurationInput
