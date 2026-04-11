This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup OnPremisesTagSetListObject

The `OnPremisesTagSetListObject` property type specifies lists of on-premises
instance tag groups. In order for an instance to be included in the deployment group, it must
be identified by all the tag groups in the list.

`OnPremisesTagSetListObject` is a property of the [CodeDeploy DeploymentGroup OnPremisesTagSet](../userguide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnPremisesTagGroup" : [ TagFilter, ... ]
}

```

### YAML

```yaml

  OnPremisesTagGroup:
    - TagFilter

```

## Properties

`OnPremisesTagGroup`

Information about groups of on-premises instance tags.

_Required_: No

_Type_: Array of [TagFilter](aws-properties-codedeploy-deploymentgroup-tagfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnPremisesTagSet

RevisionLocation

All content copied from https://docs.aws.amazon.com/.
