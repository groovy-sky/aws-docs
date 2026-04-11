This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup EC2TagSet

The `EC2TagSet` property type specifies information about groups of tags
applied to Amazon EC2 instances. The deployment group includes only Amazon EC2 instances identified by all the tag groups. `EC2TagSet` cannot be
used in the same template as `EC2TagFilter`.

For information about using tags and tag groups to help manage your Amazon EC2
instances and on-premises instances, see [Tagging Instances for Deployment\
Groups in AWS CodeDeploy](../../../codedeploy/latest/userguide/instances-tagging.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ec2TagSetList" : [ EC2TagSetListObject, ... ]
}

```

### YAML

```yaml

  Ec2TagSetList:
    - EC2TagSetListObject

```

## Properties

`Ec2TagSetList`

The Amazon EC2 tags that are already applied to Amazon EC2 instances
that you want to include in the deployment group. CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group.

Duplicates are not allowed.

_Required_: No

_Type_: Array of [EC2TagSetListObject](aws-properties-codedeploy-deploymentgroup-ec2tagsetlistobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [EC2TagSet](../../../../reference/codedeploy/latest/apireference/api-ec2tagset.md) in the _AWS CodeDeploy API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EC2TagFilter

EC2TagSetListObject

All content copied from https://docs.aws.amazon.com/.
