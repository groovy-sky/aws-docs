This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EksPodProperties

The properties for the pod.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Containers" : [ EksContainer, ... ],
  "DnsPolicy" : String,
  "HostNetwork" : Boolean,
  "ImagePullSecrets" : [ ImagePullSecret, ... ],
  "InitContainers" : [ EksContainer, ... ],
  "Metadata" : EksMetadata,
  "ServiceAccountName" : String,
  "ShareProcessNamespace" : Boolean,
  "Volumes" : [ EksVolume, ... ]
}

```

### YAML

```yaml

  Containers:
    - EksContainer
  DnsPolicy: String
  HostNetwork: Boolean
  ImagePullSecrets:
    - ImagePullSecret
  InitContainers:
    - EksContainer
  Metadata:
    EksMetadata
  ServiceAccountName: String
  ShareProcessNamespace: Boolean
  Volumes:
    - EksVolume

```

## Properties

`Containers`

The properties of the container that's used on the Amazon EKS pod.

###### Note

This object is limited to 10 elements.

_Required_: No

_Type_: Array of [EksContainer](aws-properties-batch-jobdefinition-ekscontainer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DnsPolicy`

The DNS policy for the pod. The default value is `ClusterFirst`. If the
`hostNetwork` parameter is not specified, the default is
`ClusterFirstWithHostNet`. `ClusterFirst` indicates that any DNS query
that does not match the configured cluster domain suffix is forwarded to the upstream nameserver
inherited from the node. For more information, see [Pod's DNS policy](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service) in the _Kubernetes documentation_.

Valid values: `Default` \| `ClusterFirst` \|
`ClusterFirstWithHostNet`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostNetwork`

Indicates if the pod uses the hosts' network IP address. The default value is
`true`. Setting this to `false` enables the Kubernetes pod networking model.
Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each
pod for incoming connections. For more information, see [Host\
namespaces](https://kubernetes.io/docs/concepts/security/pod-security-policy) and [Pod networking](https://kubernetes.io/docs/concepts/workloads/pods)
in the _Kubernetes documentation_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImagePullSecrets`

References a Kubernetes secret resource. It holds a list of secrets. These secrets help to gain
access to pull an images from a private registry.

`ImagePullSecret$name` is required when this object is used.

_Required_: No

_Type_: Array of [ImagePullSecret](aws-properties-batch-jobdefinition-imagepullsecret.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InitContainers`

These containers run before application containers, always runs to completion, and must
complete successfully before the next container starts. These containers are registered with the
Amazon EKS Connector agent and persists the registration information in the Kubernetes backend data store.
For more information, see [Init\
Containers](https://kubernetes.io/docs/concepts/workloads/pods/init-containers) in the _Kubernetes documentation_.

###### Note

This object is limited to 10 elements.

_Required_: No

_Type_: Array of [EksContainer](aws-properties-batch-jobdefinition-ekscontainer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metadata`

Metadata about the Kubernetes pod. For more information, see [Understanding Kubernetes Objects](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: [EksMetadata](aws-properties-batch-jobdefinition-eksmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceAccountName`

The name of the service account that's used to run the pod. For more information, see
[Kubernetes service\
accounts](../../../eks/latest/userguide/service-accounts.md) and [Configure a Kubernetes service account\
to assume an IAM role](../../../eks/latest/userguide/associate-service-account-role.md) in the _Amazon EKS User Guide_ and [Configure service accounts for pods](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account) in the _Kubernetes_
_documentation_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShareProcessNamespace`

Indicates if the processes in a container are shared, or visible, to other containers in the
same pod. For more information, see [Share\
Process Namespace between Containers in a Pod](https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Volumes`

Specifies the volumes for a job definition that uses Amazon EKS resources.

_Required_: No

_Type_: Array of [EksVolume](aws-properties-batch-jobdefinition-eksvolume.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EksPersistentVolumeClaim

EksProperties

All content copied from https://docs.aws.amazon.com/.
