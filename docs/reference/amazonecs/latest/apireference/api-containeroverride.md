# ContainerOverride

The overrides that are sent to a container. An empty container override can be passed
in. An example of an empty container override is `{"containerOverrides": [ ]
				}`. If a non-empty container override is specified, the `name`
parameter must be included.

You can use Secrets Manager or AWS Systems Manager Parameter Store to store the
sensitive data. For more information, see [Retrieve secrets through\
environment variables](../../../../services/amazonecs/latest/developerguide/secrets-envvar.md) in the Amazon ECS Developer Guide.

## Contents

**command**

The command to send to the container that overrides the default command from the
Docker image or the task definition. You must also specify a container name.

Type: Array of strings

Required: No

**cpu**

The number of `cpu` units reserved for the container, instead of the
default value from the task definition. You must also specify a container name.

Type: Integer

Required: No

**environment**

The environment variables to send to the container. You can add new environment
variables, which are added to the container at launch, or you can override the existing
environment variables from the Docker image or the task definition. You must also
specify a container name.

Type: Array of [KeyValuePair](api-keyvaluepair.md) objects

Required: No

**environmentFiles**

A list of files containing the environment variables to pass to a container, instead
of the value from the container definition.

Type: Array of [EnvironmentFile](api-environmentfile.md) objects

Required: No

**memory**

The hard limit (in MiB) of memory to present to the container, instead of the default
value from the task definition. If your container attempts to exceed the memory
specified here, the container is killed. You must also specify a container name.

Type: Integer

Required: No

**memoryReservation**

The soft limit (in MiB) of memory to reserve for the container, instead of the default
value from the task definition. You must also specify a container name.

Type: Integer

Required: No

**name**

The name of the container that receives the override. This parameter is required if
any override is specified.

Type: String

Required: No

**resourceRequirements**

The type and amount of a resource to assign to a container, instead of the default
value from the task definition. The only supported resource is a GPU.

Type: Array of [ResourceRequirement](api-resourcerequirement.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/containeroverride.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/containeroverride.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/containeroverride.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ContainerInstanceHealthStatus

ContainerRestartPolicy
