This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Container PrivateRegistryAccess

Describes the configuration for an Amazon Lightsail container service to access private
container image repositories, such as Amazon Elastic Container Registry (Amazon ECR)
private repositories.

For more information, see [Configuring access to an Amazon ECR private repository for an Amazon Lightsail\
container service](../../../lightsail/latest/userguide/amazon-lightsail-container-service-ecr-private-repo-access.md) in the _Amazon Lightsail Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EcrImagePullerRole" : EcrImagePullerRole
}

```

### YAML

```yaml

  EcrImagePullerRole:
    EcrImagePullerRole

```

## Properties

`EcrImagePullerRole`

An object that describes the activation status of the role that you can use to grant a
Lightsail container service access to Amazon ECR private repositories. If the role is
activated, the Amazon Resource Name (ARN) of the role is also listed.

_Required_: No

_Type_: [EcrImagePullerRole](aws-properties-lightsail-container-ecrimagepullerrole.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PortInfo

PublicDomainName

All content copied from https://docs.aws.amazon.com/.
