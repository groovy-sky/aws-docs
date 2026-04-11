This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration ContainerInfo

The information about the container used for a job run or a managed endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EksInfo" : EksInfo
}

```

### YAML

```yaml

  EksInfo:
    EksInfo

```

## Properties

`EksInfo`

The information about the Amazon EKS cluster.

_Required_: No

_Type_: [EksInfo](aws-properties-emrcontainers-securityconfiguration-eksinfo.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationConfiguration

ContainerProvider

All content copied from https://docs.aws.amazon.com/.
