This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImageRecipe AdditionalInstanceConfiguration

In addition to your infrastructure configuration, these settings provide an extra
layer of control over your build instances. You can also specify commands to run on
launch for all of your build instances.

Image Builder does not automatically install the Systems Manager agent on Windows instances. If your base
image includes the Systems Manager agent, then the AMI that you create will also include the
agent. For Linux instances, if the base image does not already include the Systems Manager agent,
Image Builder installs it. For Linux instances where Image Builder installs the Systems Manager agent, you can
choose whether to keep it for the AMI that you create.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SystemsManagerAgent" : SystemsManagerAgent,
  "UserDataOverride" : String
}

```

### YAML

```yaml

  SystemsManagerAgent:
    SystemsManagerAgent
  UserDataOverride: String

```

## Properties

`SystemsManagerAgent`

Contains settings for the Systems Manager agent on your build instance.

_Required_: No

_Type_: [SystemsManagerAgent](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-systemsmanageragent.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserDataOverride`

Use this property to provide commands or a command script to run when you launch your
build instance.

The userDataOverride property replaces any commands that Image Builder might have added to
ensure that Systems Manager is installed on your Linux build instance. If you override the user
data, make sure that you add commands to install Systems Manager, if it is not pre-installed on
your base image.

###### Note

The user data is always base 64 encoded. For example, the following commands are
encoded as
`IyEvYmluL2Jhc2gKbWtkaXIgLXAgL3Zhci9iYi8KdG91Y2ggL3Zhci$`:

_#!/bin/bash_

mkdir -p /var/bb/

touch /var

_Required_: No

_Type_: String

_Pattern_: `^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`

_Minimum_: `1`

_Maximum_: `21847`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ImageBuilder::ImageRecipe

ComponentConfiguration
