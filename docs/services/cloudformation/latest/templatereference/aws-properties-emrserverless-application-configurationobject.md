This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRServerless::Application ConfigurationObject

A configuration specification to be used when provisioning an application. A
configuration consists of a classification, properties, and optional nested configurations.
A classification refers to an application-specific configuration file. Properties are the
settings you want to change in that file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Classification" : String,
  "Configurations" : [ ConfigurationObject, ... ],
  "Properties" : {Key: Value, ...}
}

```

### YAML

```yaml

  Classification: String
  Configurations:
    - ConfigurationObject
  Properties:
    Key: Value

```

## Properties

`Classification`

The classification within a configuration.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Configurations`

A list of additional configurations to apply within a configuration object.

_Required_: No

_Type_: Array of [ConfigurationObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-emrserverless-application-configurationobject.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Properties`

A set of properties specified within a configuration classification.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z]+[-a-zA-Z0-9_.]*$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchLoggingConfiguration

IdentityCenterConfiguration
