This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReplicaStreamSpecification

Represents the DynamoDB Streams configuration for a global table
replica.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourcePolicy" : ResourcePolicy
}

```

### YAML

```yaml

  ResourcePolicy:
    ResourcePolicy

```

## Properties

`ResourcePolicy`

A resource-based policy document that contains the permissions for the specified stream
of a DynamoDB global table replica. Resource-based policies let you define access
permissions by specifying who has access to each resource, and the actions they are allowed
to perform on each resource.

In a CloudFormation template, you can provide the policy in JSON or YAML format
because CloudFormation converts YAML to JSON before submitting it to DynamoDB. For more information about resource-based policies, see [Using\
resource-based policies for DynamoDB](../../../dynamodb/latest/developerguide/access-control-resource-based.md) and [Resource-based policy\
examples](../../../dynamodb/latest/developerguide/rbac-examples.md).

You can update the `ResourcePolicy` property if you've specified more than
one table using the [AWS::DynamoDB::GlobalTable](../userguide/aws-resource-dynamodb-globaltable.md) resource.

_Required_: Yes

_Type_: [ResourcePolicy](aws-properties-dynamodb-globaltable-resourcepolicy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaSSESpecification

ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
