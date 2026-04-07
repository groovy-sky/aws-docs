This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ResourceDefinition ResourceDefinitionVersion

A resource definition version contains a list of resources. (In CloudFormation,
resources are named _resource instances_.)

###### Note

After you create a resource definition version that contains the resources you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In an CloudFormation template, `ResourceDefinitionVersion` is the property
type of the `InitialVersion` property in the [`AWS::Greengrass::ResourceDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinition.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Resources" : [ ResourceInstance, ... ]
}

```

### YAML

```yaml

  Resources:
    - ResourceInstance

```

## Properties

`Resources`

The resources in this version.

_Required_: Yes

_Type_: Array of [ResourceInstance](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-resourcedefinition-resourceinstance.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ResourceDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-resourcedefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceDataContainer

ResourceDownloadOwnerSetting
