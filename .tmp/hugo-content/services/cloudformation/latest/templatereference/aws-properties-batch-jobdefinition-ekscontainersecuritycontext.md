This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksContainerSecurityContext

The security context for a job. For more information, see [Configure a\
security context for a pod or container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context) in the _Kubernetes_
_documentation_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowPrivilegeEscalation" : Boolean,
  "Privileged" : Boolean,
  "ReadOnlyRootFilesystem" : Boolean,
  "RunAsGroup" : Integer,
  "RunAsNonRoot" : Boolean,
  "RunAsUser" : Integer
}

```

### YAML

```yaml

  AllowPrivilegeEscalation: Boolean
  Privileged: Boolean
  ReadOnlyRootFilesystem: Boolean
  RunAsGroup: Integer
  RunAsNonRoot: Boolean
  RunAsUser: Integer

```

## Properties

`AllowPrivilegeEscalation`

Whether or not a container or a Kubernetes pod is allowed to gain more privileges than its parent
process. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Privileged`

When this parameter is `true`, the container is given elevated permissions on the
host container instance. The level of permissions are similar to the `root` user
permissions. The default value is `false`. This parameter maps to
`privileged` policy in the [Privileged\
pod security policies](https://kubernetes.io/docs/concepts/security/pod-security-policy) in the _Kubernetes documentation_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadOnlyRootFilesystem`

When this parameter is `true`, the container is given read-only access to its
root file system. The default value is `false`. This parameter maps to
`ReadOnlyRootFilesystem` policy in the [Volumes and file systems pod security policies](https://kubernetes.io/docs/concepts/security/pod-security-policy) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunAsGroup`

When this parameter is specified, the container is run as the specified group ID
( `gid`). If this parameter isn't specified, the default is the group that's specified
in the image metadata. This parameter maps to `RunAsGroup` and `MustRunAs`
policy in the [Users\
and groups pod security policies](https://kubernetes.io/docs/concepts/security/pod-security-policy) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunAsNonRoot`

When this parameter is specified, the container is run as a user with a `uid`
other than 0. If this parameter isn't specified, so such rule is enforced. This parameter maps to
`RunAsUser` and `MustRunAsNonRoot` policy in the [Users\
and groups pod security policies](https://kubernetes.io/docs/concepts/security/pod-security-policy) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunAsUser`

When this parameter is specified, the container is run as the specified user ID
( `uid`). If this parameter isn't specified, the default is the user that's specified
in the image metadata. This parameter maps to `RunAsUser` and `MustRanAs`
policy in the [Users\
and groups pod security policies](https://kubernetes.io/docs/concepts/security/pod-security-policy) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksContainerResourceRequirements

EksContainerVolumeMount

All content copied from https://docs.aws.amazon.com/.
