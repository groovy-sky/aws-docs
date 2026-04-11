This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ServiceConnectDefaults

Use this parameter to set a default Service Connect namespace. After you set a default
Service Connect namespace, any new services with Service Connect turned on that are
created in the cluster are added as client services in the namespace. This setting only
applies to new services that set the `enabled` parameter to `true`
in the `ServiceConnectConfiguration`. You can set the namespace of each
service individually in the `ServiceConnectConfiguration` to override this
default parameter.

Tasks that run in a namespace can use short names to connect to services in the
namespace. Tasks can connect to services across all of the clusters in the namespace.
Tasks connect through a managed proxy container that collects logs and metrics for
increased visibility. Only the tasks that Amazon ECS services create are supported with
Service Connect. For more information, see [Service Connect](../../../amazonecs/latest/developerguide/service-connect.md)
in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespace" : String
}

```

### YAML

```yaml

  Namespace: String

```

## Properties

`Namespace`

The namespace name or full Amazon Resource Name (ARN) of the AWS Cloud Map namespace that's
used when you create a service and don't specify a Service Connect configuration. The
namespace name can include up to 1024 characters. The name is case-sensitive. The name
can't include greater than (>), less than (<), double quotation marks ("), or
slash (/).

If you enter an existing namespace name or ARN, then that namespace will be used. Any
namespace type is supported. The namespace must be in this account and this AWS
Region.

If you enter a new name, a AWS Cloud Map namespace will be created. Amazon ECS creates a AWS Cloud Map namespace
with the "API calls" method of instance discovery only. This instance discovery method
is the "HTTP" namespace type in the AWS Command Line Interface. Other types of
instance discovery aren't used by Service Connect.

If you update the cluster with an empty string `""` for the namespace name,
the cluster configuration for Service Connect is removed. Note that the namespace will
remain in AWS Cloud Map and must be deleted separately.

For more information about AWS Cloud Map, see [Working with Services](../../../cloud-map/latest/dg/working-with-services.md)
in the _AWS Cloud Map Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedStorageConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
