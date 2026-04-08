# Additional network path metadata included for Amazon EKS

When Network Flow Monitor gathers performance metrics for network flows between Amazon EKS components,
it includes additional metadata information about the network path,
to help you better understand how the network paths for your workload are performing.

You can view detailed information about Amazon EKS network flow performance
by creating a monitor for the network flows that you're interested in,
and then viewing details on the **Historical explorer** tab.

With Network Flow Monitor, you can measure network performance between the following Amazon EKS
components, to better understand how your workload is performing with your
Amazon EKS configuration and determine where there are bottlenecks or impairments.

- Pod to pod on the same node

- Node to node on the same cluster

- Pod to pod on a different cluster

- Node to node on different clusters

- With and without Network Load Balancer

The following table lists the information that Network Flow Monitor returns for each network flow scenario.

**Connection information****Metadata information****Local****Remote****Scenario****Initiated by****Local****Remote****Pod name****Service****Namespace****Pod name****Service****Namespace**Local pod connecting to cluster IP of another internal cluster serviceLocalLocal pod IP address

Remote pod IP address

(through cluster IP address)

✓✓✓✓ ¹✓✓Local pod in a node network namespace connecting to cluster IP of another internal cluster serviceLocalLocal node IP address

Remote pod IP address

(through cluster IP address)

✓ ²✓ ²✓ ²✓ ¹✓✓Local pod connecting to individual pod IP address of another pod (headless service)LocalLocal pod IP addressRemote pod IP address✓✓✓✓✓✓Local pod connecting to individual pod IP address of another pod in node network namespace (headless service)LocalLocal pod IP addressRemote node IP address✓✓✓✓✓✓Local pod connecting to remote pod in another clusterLocalLocal pod IP address

Remote pod IP address

(another cluster)

✓✓✓✗✗✗Local pod connecting to an external network addressLocalLocal pod IP addressExternal IP address✓✓✓N/AN/AN/ALocal pod operating in a node network namespace connecting to an external network IP addressLocalLocal node IP addressExternal IP address✓ ²✓ ²✓ ²N/AN/AN/ARemote pod connecting to local pod through cluster IP addressRemote

Local pod IP address

(through cluster IP address)

Remote pod IP address✓✓✓✓✓✓Remote pod in a node network namespace connecting to local podRemote

Local pod IP address

(through cluster IP address)

Remote node IP address✓✓✓✓ ³✓ ³✓ ³Remote pod connecting to local pod (headless service)RemoteLocal pod IP addressRemote pod IP address✓✓✓✓✓✓External pod connecting to a local podRemoteLocal pod IP addressRemote pod IP address✓✓✓✗✗✗External resource connecting through NodePort or a Load Balancer to a local podRemoteLocal pod IP addressExternal IP address ⁴✓✓✓N/AN/AN/AExternal resource connecting through NodePort or a Load Balancer to a local pod operating in a node network namespaceRemoteLocal node IP addressExternal IP address ⁴✓✓✓N/AN/AN/A

Be aware of the following additional information corresponding to the items marked with footnotes in the preceding table.

1. Pod name is not visible in this scenario for pods with other owners, such as a Kubernetes service managed by the EKS control plane.

2. Local pod name, service, and namespace are not resolved if other pods are present in node network namespace.

3. Remote pod name, service, and namespace are not resolved if other pods are present in node network namespace.

4. If service is using NodePort or LoadBalancer in instance mode, and `ExternalTrafficPolicy` is set to `Cluster`,
    then this IP address will be reported as the IP address of the node that receives the NodePort connection.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configure for third party tools

Install EC2 agents

All content copied from https://docs.aws.amazon.com/.
