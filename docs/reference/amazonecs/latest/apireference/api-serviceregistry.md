# ServiceRegistry

The details for the service registry.

Each service may be associated with one service registry. Multiple service registries
for each service are not supported.

When you add, update, or remove the service registries configuration, Amazon ECS
starts a new deployment. New tasks are registered and deregistered to the updated
service registry configuration.

## Contents

**containerName**

The container name value to be used for your service discovery service. It's already
specified in the task definition. If the task definition that your service task
specifies uses the `bridge` or `host` network mode, you must
specify a `containerName` and `containerPort` combination from the
task definition. If the task definition that your service task specifies uses the
`awsvpc` network mode and a type SRV DNS record is used, you must specify
either a `containerName` and `containerPort` combination or a
`port` value. However, you can't specify both.

Type: String

Required: No

**containerPort**

The port value to be used for your service discovery service. It's already specified
in the task definition. If the task definition your service task specifies uses the
`bridge` or `host` network mode, you must specify a
`containerName` and `containerPort` combination from the task
definition. If the task definition your service task specifies uses the
`awsvpc` network mode and a type SRV DNS record is used, you must specify
either a `containerName` and `containerPort` combination or a
`port` value. However, you can't specify both.

Type: Integer

Required: No

**port**

The port value used if your service discovery service specified an SRV record. This
field might be used if both the `awsvpc` network mode and SRV records are
used.

Type: Integer

Required: No

**registryArn**

The Amazon Resource Name (ARN) of the service registry. The currently supported
service registry is AWS Cloud Map. For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/ServiceRegistry)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/ServiceRegistry)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/ServiceRegistry)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceManagedEBSVolumeConfiguration

ServiceRevision
