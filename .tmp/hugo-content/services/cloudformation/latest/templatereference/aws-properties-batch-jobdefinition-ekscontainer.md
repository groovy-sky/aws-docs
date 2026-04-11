This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksContainer

EKS container properties are used in job definitions for Amazon EKS based job definitions to
describe the properties for a container node in the pod that's launched as part of a job. This
can't be specified for Amazon ECS based job definitions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Args" : [ String, ... ],
  "Command" : [ String, ... ],
  "Env" : [ EksContainerEnvironmentVariable, ... ],
  "Image" : String,
  "ImagePullPolicy" : String,
  "Name" : String,
  "Resources" : EksContainerResourceRequirements,
  "SecurityContext" : EksContainerSecurityContext,
  "VolumeMounts" : [ EksContainerVolumeMount, ... ]
}

```

### YAML

```yaml

  Args:
    - String
  Command:
    - String
  Env:
    - EksContainerEnvironmentVariable
  Image: String
  ImagePullPolicy: String
  Name: String
  Resources:
    EksContainerResourceRequirements
  SecurityContext:
    EksContainerSecurityContext
  VolumeMounts:
    - EksContainerVolumeMount

```

## Properties

`Args`

An array of arguments to the entrypoint. If this isn't specified, the `CMD` of
the container image is used. This corresponds to the `args` member in the [Entrypoint](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1) portion of the [Pod](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1)
in Kubernetes. Environment variable references are expanded using the container's environment.

If the referenced environment variable doesn't exist, the reference in the command isn't
changed. For example, if the reference is to " `$(NAME1)`" and the `NAME1`
environment variable doesn't exist, the command string will remain " `$(NAME1)`."
`$$` is replaced with `$`, and the resulting string isn't expanded. For
example, `$$(VAR_NAME)` is passed as `$(VAR_NAME)` whether or not the
`VAR_NAME` environment variable exists. For more information, see [Dockerfile reference: CMD](https://docs.docker.com/engine/reference/builder)
and [Define a command and arguments for a pod](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Command`

The entrypoint for the container. This isn't run within a shell. If this isn't specified,
the `ENTRYPOINT` of the container image is used. Environment variable references are
expanded using the container's environment.

If the referenced environment variable doesn't exist, the reference in the command isn't
changed. For example, if the reference is to " `$(NAME1)`" and the `NAME1`
environment variable doesn't exist, the command string will remain " `$(NAME1)`."
`$$` is replaced with `$` and the resulting string isn't expanded. For
example, `$$(VAR_NAME)` will be passed as `$(VAR_NAME)` whether or not the
`VAR_NAME` environment variable exists. The entrypoint can't be updated. For more
information, see [ENTRYPOINT](https://docs.docker.com/engine/reference/builder) in the _Dockerfile reference_ and [Define a command and arguments for a container](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container) and [Entrypoint](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1) in the _Kubernetes documentation_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Env`

The environment variables to pass to a container.

###### Note

Environment variables cannot start with " `AWS_BATCH`". This naming
convention is reserved for variables that AWS Batch sets.

_Required_: No

_Type_: Array of [EksContainerEnvironmentVariable](aws-properties-batch-jobdefinition-ekscontainerenvironmentvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The Docker image used to start the container.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImagePullPolicy`

The image pull policy for the container. Supported values are `Always`,
`IfNotPresent`, and `Never`. This parameter defaults to
`IfNotPresent`. However, if the `:latest` tag is specified, it defaults to
`Always`. For more information, see [Updating\
images](https://kubernetes.io/docs/concepts/containers/images) in the _Kubernetes documentation_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the container. If the name isn't specified, the default name
" `Default`" is used. Each container in a pod must have a unique name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

The type and amount of resources to assign to a container. The supported resources include
`memory`, `cpu`, and `nvidia.com/gpu`. For more information,
see [Resource management for pods and containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: [EksContainerResourceRequirements](aws-properties-batch-jobdefinition-ekscontainerresourcerequirements.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityContext`

The security context for a job. For more information, see [Configure a\
security context for a pod or container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: [EksContainerSecurityContext](aws-properties-batch-jobdefinition-ekscontainersecuritycontext.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeMounts`

The volume mounts for the container. AWS Batch supports `emptyDir`,
`hostPath`, and `secret` volume types. For more information about volumes
and volume mounts in Kubernetes, see [Volumes](https://kubernetes.io/docs/concepts/storage/volumes) in the _Kubernetes documentation_.

_Required_: No

_Type_: Array of [EksContainerVolumeMount](aws-properties-batch-jobdefinition-ekscontainervolumemount.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EFSVolumeConfiguration

EksContainerEnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
