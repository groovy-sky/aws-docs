This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::ReplicationConfiguration

The `AWS::ECR::ReplicationConfiguration` resource creates or updates the
replication configuration for a private registry. The first time a replication
configuration is applied to a private registry, a service-linked IAM role
is created in your account for the replication process. For more information, see [Using\
Service-Linked Roles for Amazon ECR](../../../amazonecr/latest/userguide/using-service-linked-roles.md) in the _Amazon Elastic_
_Container Registry User Guide_.

###### Note

When configuring cross-account replication, the destination account must grant the
source account permission to replicate. This permission is controlled using a
private registry permissions policy. For more information, see
`AWS::ECR::RegistryPolicy`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::ReplicationConfiguration",
  "Properties" : {
      "ReplicationConfiguration" : ReplicationConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::ECR::ReplicationConfiguration
Properties:
  ReplicationConfiguration:
    ReplicationConfiguration

```

## Properties

`ReplicationConfiguration`

The replication configuration for a registry.

_Required_: Yes

_Type_: [ReplicationConfiguration](aws-properties-ecr-replicationconfiguration-replicationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RegistryId`

The account ID of the destination registry.

## Examples

### Specify a replication configuration for a private registry

The following example specifies a replication configuration in a source Region
for a private registry to replicate the contents to the `us-east-2`
and `us-west-1` Regions within the same account.

#### JSON

```json

"TestReplicationConfiguration": {
  "Type": "AWS::ECR::ReplicationConfiguration",
  "Properties": {
     "ReplicationConfiguration": {
        "Rules": [
           {
              "Destinations": [
                 {
                    "Region": "us-east-2",
                    "RegistryId": "123456789012"
                 },
                 {
                     "Region": "us-west-1",
                     "RegistryId": "123456789012"
                 }
               ]
             }
          ]
       }
    }
}
```

#### YAML

```yaml

Resources:
  MyReplicationConfig:
    Type: AWS::ECR::ReplicationConfiguration
    Properties:
      ReplicationConfiguration:
          Rules:
            - Destinations:
                - Region: "us-east-2"
                  RegistryId: "123456789012"
                - Region: "us-west-1"
                  RegistryId: "123456789012"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScanningRule

ReplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
