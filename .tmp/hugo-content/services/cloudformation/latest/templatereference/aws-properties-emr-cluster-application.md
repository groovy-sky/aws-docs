This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster Application

`Application` is a property of `AWS::EMR::Cluster`. The `Application` property type defines the open-source big data applications for EMR to install and configure when a cluster is created.

With Amazon EMR release version 4.0 and later, the only accepted parameter is the application `Name`. To pass arguments to these applications, you use configuration classifications specified using JSON objects in a `Configuration` property. For more information, see [Configuring Applications](../../../emr/latest/releaseguide/emr-configure-apps.md).

With earlier Amazon EMR releases, the application is any AWS or third-party software that you can add to the cluster. You can specify the version of the application and arguments to pass to it. Amazon EMR accepts and forwards the argument list to the corresponding installation script as a bootstrap action argument.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalInfo" : {Key: Value, ...},
  "Args" : [ String, ... ],
  "Name" : String,
  "Version" : String
}

```

### YAML

```yaml

  AdditionalInfo:
    Key: Value
  Args:
    - String
  Name: String
  Version: String

```

## Properties

`AdditionalInfo`

This option is for advanced users only. This is meta information about clusters and applications that are used for testing and troubleshooting.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Args`

Arguments for Amazon EMR to pass to the application.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the application.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The version of the application.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::Cluster

AutoScalingPolicy

All content copied from https://docs.aws.amazon.com/.
