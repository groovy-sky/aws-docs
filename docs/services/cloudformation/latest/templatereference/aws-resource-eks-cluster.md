This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster

Creates an Amazon EKS control plane.

The Amazon EKS control plane consists of control plane instances that run the Kubernetes
software, such as `etcd` and the API server. The control plane runs in an
account managed by AWS, and the Kubernetes API is exposed by the Amazon EKS API server endpoint.
Each Amazon EKS cluster control plane is single tenant and unique. It runs on its own set of
Amazon EC2 instances.

The cluster control plane is provisioned across multiple Availability Zones and fronted by an Elastic Load Balancing
Network Load Balancer. Amazon EKS also provisions elastic network interfaces in your VPC subnets to provide
connectivity from the control plane instances to the nodes (for example, to support
`kubectl exec`, `logs`, and `proxy` data
flows).

Amazon EKS nodes run in your AWS account and connect to your cluster's control plane over
the Kubernetes API server endpoint and a certificate file that is created for your
cluster.

You can use the `endpointPublicAccess` and
`endpointPrivateAccess` parameters to enable or disable public and
private access to your cluster's Kubernetes API server endpoint. By default, public access is
enabled, and private access is disabled. The
endpoint domain name and IP address family depends on the value of the
`ipFamily` for the cluster. For more information, see [Amazon EKS Cluster\
Endpoint Access Control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the _Amazon EKS User Guide_.

You can use the `logging` parameter to enable or disable exporting the
Kubernetes control plane logs for your cluster to CloudWatch Logs. By default, cluster control plane
logs aren't exported to CloudWatch Logs. For more information, see [Amazon EKS\
Cluster Control Plane Logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the
_Amazon EKS User Guide_.

###### Note

CloudWatch Logs ingestion, archive storage, and data scanning rates apply to exported
control plane logs. For more information, see [CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

In most cases, it takes several minutes to create a cluster. After you create an Amazon EKS
cluster, you must configure your Kubernetes tooling to communicate with the API server and
launch nodes into your cluster. For more information, see [Allowing users to\
access your cluster](https://docs.aws.amazon.com/eks/latest/userguide/cluster-auth.html) and [Launching Amazon EKS\
nodes](https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html) in the _Amazon EKS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::Cluster",
  "Properties" : {
      "AccessConfig" : AccessConfig,
      "BootstrapSelfManagedAddons" : Boolean,
      "ComputeConfig" : ComputeConfig,
      "ControlPlaneScalingConfig" : ControlPlaneScalingConfig,
      "DeletionProtection" : Boolean,
      "EncryptionConfig" : [ EncryptionConfig, ... ],
      "Force" : Boolean,
      "KubernetesNetworkConfig" : KubernetesNetworkConfig,
      "Logging" : Logging,
      "Name" : String,
      "OutpostConfig" : OutpostConfig,
      "RemoteNetworkConfig" : RemoteNetworkConfig,
      "ResourcesVpcConfig" : ResourcesVpcConfig,
      "RoleArn" : String,
      "StorageConfig" : StorageConfig,
      "Tags" : [ Tag, ... ],
      "UpgradePolicy" : UpgradePolicy,
      "Version" : String,
      "ZonalShiftConfig" : ZonalShiftConfig
    }
}

```

### YAML

```yaml

Type: AWS::EKS::Cluster
Properties:
  AccessConfig:
    AccessConfig
  BootstrapSelfManagedAddons: Boolean
  ComputeConfig:
    ComputeConfig
  ControlPlaneScalingConfig:
    ControlPlaneScalingConfig
  DeletionProtection: Boolean
  EncryptionConfig:
    - EncryptionConfig
  Force: Boolean
  KubernetesNetworkConfig:
    KubernetesNetworkConfig
  Logging:
    Logging
  Name: String
  OutpostConfig:
    OutpostConfig
  RemoteNetworkConfig:
    RemoteNetworkConfig
  ResourcesVpcConfig:
    ResourcesVpcConfig
  RoleArn: String
  StorageConfig:
    StorageConfig
  Tags:
    - Tag
  UpgradePolicy:
    UpgradePolicy
  Version: String
  ZonalShiftConfig:
    ZonalShiftConfig

```

## Properties

`AccessConfig`

The access configuration for the cluster.

_Required_: No

_Type_: [AccessConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-accessconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BootstrapSelfManagedAddons`

If you set this value to `False` when creating a cluster, the default
networking add-ons will not be installed.

The default networking add-ons include `vpc-cni`, `coredns`, and
`kube-proxy`.

Use this option when you plan to install third-party alternative add-ons or
self-manage the default networking add-ons.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ComputeConfig`

Indicates the current configuration of the compute capability on your EKS Auto Mode
cluster. For example, if the capability is enabled or disabled. If the compute
capability is enabled, EKS Auto Mode will create and delete EC2 Managed Instances in
your AWS account. For more information, see EKS Auto Mode compute capability in the
_Amazon EKS User Guide_.

_Required_: No

_Type_: [ComputeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-computeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlPlaneScalingConfig`

The control plane scaling tier configuration. For more information, see EKS Provisioned Control Plane in the Amazon EKS User Guide.

_Required_: No

_Type_: [ControlPlaneScalingConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-controlplanescalingconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtection`

The current deletion protection setting for the cluster. When `true`,
deletion protection is enabled and the cluster cannot be deleted until protection is
disabled. When `false`, the cluster can be deleted normally. This setting
only applies to clusters in an active state.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfig`

The encryption configuration for the cluster.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-encryptionconfig.html) of [EncryptionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-encryptionconfig.html)

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Force`

Set this value to `true` to override upgrade-blocking readiness checks when
updating a cluster.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KubernetesNetworkConfig`

The Kubernetes network configuration for the cluster.

_Required_: No

_Type_: [KubernetesNetworkConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-kubernetesnetworkconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

The logging configuration for your cluster.

_Required_: No

_Type_: [Logging](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-logging.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name to give to your cluster. The name can contain only alphanumeric characters (case-sensitive) and
hyphens. It must start with an alphanumeric character and can't be longer than 100 characters.
The name must be unique within the AWS Region and AWS account that you're creating the cluster in. Note that underscores can't be used in CloudFormation.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z][A-Za-z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutpostConfig`

An object representing the configuration of your local Amazon EKS cluster on an AWS
Outpost. This object isn't available for clusters on the AWS cloud.

_Required_: No

_Type_: [OutpostConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-outpostconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RemoteNetworkConfig`

The configuration in the cluster for EKS Hybrid Nodes. You can add, change, or remove this
configuration after the cluster is created.

_Required_: No

_Type_: [RemoteNetworkConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-remotenetworkconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcesVpcConfig`

The VPC configuration that's used by the cluster control plane. Amazon EKS VPC
resources have specific requirements to work properly with Kubernetes. For more
information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and
[Cluster\
Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the _Amazon EKS User_
_Guide_. You must specify at least two subnets. You can specify up to five
security groups, but we recommend that you use a dedicated security group for your
cluster control plane.

_Required_: Yes

_Type_: [ResourcesVpcConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-resourcesvpcconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane
to make calls to AWS API operations on your behalf. For more information, see [Amazon EKS Service\
IAM Role](https://docs.aws.amazon.com/eks/latest/userguide/service_IAM_role.html) in the _Amazon EKS User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageConfig`

Indicates the current configuration of the block storage capability on your EKS Auto
Mode cluster. For example, if the capability is enabled or disabled. If the block
storage capability is enabled, EKS Auto Mode will create and delete EBS volumes in your
AWS account. For more information, see EKS Auto Mode block storage capability in the
_Amazon EKS User Guide_.

_Required_: No

_Type_: [StorageConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-storageconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The metadata that you apply to the cluster to assist with categorization and
organization. Each tag consists of a key and an optional value, both of which you
define. Cluster tags don't propagate to any other resources associated with the
cluster.

###### Note

You must have the `eks:TagResource` and `eks:UntagResource`
permissions for your [IAM\
principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CloudFormation stack. If you don't have
these permissions, there might be unexpected behavior with stack-level tags
propagating to the resource during resource creation and update.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpgradePolicy`

This value indicates if extended support is enabled or disabled for the
cluster.

[Learn more about EKS Extended Support in the _Amazon EKS User Guide_.](https://docs.aws.amazon.com/eks/latest/userguide/extended-support-control.html)

_Required_: No

_Type_: [UpgradePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-upgradepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The desired Kubernetes version for your cluster. If you don't specify a value here, the
default version available in Amazon EKS is used.

###### Note

The default version might not be the latest version available.

_Required_: No

_Type_: String

_Pattern_: `1\.\d\d`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZonalShiftConfig`

The configuration for zonal shift for the cluster.

_Required_: No

_Type_: [ZonalShiftConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-zonalshiftconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myCluster" }`

For the Amazon EKS cluster `myCluster`, `Ref` returns
the name of the cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the cluster, such as
`arn:aws:eks:us-west-2:666666666666:cluster/prod`.

`CertificateAuthorityData`

The `certificate-authority-data` for your cluster.

`ClusterSecurityGroupId`

The cluster security group that was created by Amazon EKS for the cluster.
Managed node groups use this security group for control plane to data plane
communication.

This parameter is only returned by Amazon EKS clusters that support managed
node groups. For more information, see [Managed node groups](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html) in
the _Amazon EKS User Guide_.

`EncryptionConfigKeyArn`

Amazon Resource Name (ARN) or alias of the customer master key (CMK).

`Endpoint`

The endpoint for your Kubernetes API server, such as
`https://5E1D0CEXAMPLEA591B746AFC5AB30262.yl4.us-west-2.eks.amazonaws.com`.

`Id`

The ID of your local Amazon EKS cluster on an AWS Outpost. This property isn't available
for an Amazon EKS cluster on the AWS cloud.

`KubernetesNetworkConfig.ServiceIpv6Cidr`

The CIDR block that Kubernetes Service IP addresses are assigned from if you created a
`1.21` or later cluster with version `>1.10.1` or later of the Amazon VPC CNI add-on and
specified `ipv6` for **ipFamily** when you
created the cluster. Kubernetes assigns Service addresses from the unique local address
range ( `fc00::/7`) because you can't specify a custom `IPv6` CIDR block when
you create the cluster.

`OpenIdConnectIssuerUrl`

The issuer URL for the OIDC identity provider of the cluster, such as
`https://oidc.eks.us-west-2.amazonaws.com/id/EXAMPLED539D4633E53DE1B716D3041E`.
If you need to remove `https://` from this output value, you can include the
following code in your template.

`!Select [1, !Split ["//", !GetAtt
            EKSCluster.OpenIdConnectIssuerUrl]]`

## Examples

### Create a cluster

The following example creates an Amazon EKS cluster named
`Prod`.

#### JSON

```json

{
    "EKSCluster": {
       "Type": "AWS::EKS::Cluster",
       "Properties": {
          "Name": "Prod",
          "Version": "1.20",
          "RoleArn": "arn:aws:iam::012345678910:role/eks-service-role-AWSServiceRoleForAmazonEKS-EXAMPLEBQ4PI",
          "ResourcesVpcConfig": {
             "SecurityGroupIds": [
                "sg-6979fe18"
             ],
             "SubnetIds": [
                "subnet-6782e71e",
                "subnet-e7e761ac"
             ],
             "EndpointPublicAccess": false,
             "EndpointPrivateAccess": true,
             "PublicAccessCidrs": [
                "1.1.1.2/32"
             ]
          },
          "Logging": {
             "ClusterLogging": {
                "EnabledTypes": [
                   {
                      "Type": "api"
                   },
                   {
                      "Type": "audit"
                   }
                ]
             }
          },
          "Tags": [
             {
                "Key": "key",
                "Value": "val"
             }
          ]
       }
    }
 }
```

#### YAML

```yaml

EKSCluster:
    Type: AWS::EKS::Cluster
    Properties:
      Name: Prod
      Version: "1.20"
      RoleArn: "arn:aws:iam::012345678910:role/eks-service-role-AWSServiceRoleForAmazonEKS-EXAMPLEBQ4PI"
      ResourcesVpcConfig:
        SecurityGroupIds:
          - sg-6979fe18
        SubnetIds:
          - subnet-6782e71e
          - subnet-e7e761ac
        EndpointPublicAccess: false
        EndpointPrivateAccess: true
        PublicAccessCidrs: [ "1.1.1.2/32" ]
      Logging:
        ClusterLogging:
          EnabledTypes:
            - Type: api
            - Type: audit
      Tags:
        - Key: "key"
          Value: "val"
```

## See also

- [Clusters](https://docs.aws.amazon.com/eks/latest/userguide/clusters.html) in the _Amazon EKS User_
_Guide_.

- [`CreateCluster`](https://docs.aws.amazon.com/eks/latest/APIReference/API_CreateCluster.html) in the _Amazon EKS API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AccessConfig
