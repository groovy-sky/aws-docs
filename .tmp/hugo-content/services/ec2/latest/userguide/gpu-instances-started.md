# Get started with GPU accelerated instances

The latest generation of GPU accelerated instance types, such as those shown in the following list
deliver the highest performance capabilities for deep learning and high performance computing (HPC)
applications. Select the instance type link to learn more about its capabilities.

- [P6 family](https://aws.amazon.com/ec2/instance-types/p6)

- [P6 family](https://aws.amazon.com/ec2/instance-types/p6)

- [P5 family](https://aws.amazon.com/ec2/instance-types/p5)

For a complete list of instance type specifications for accelerated instance types, see
[Accelerated computing](../instancetypes/ac.md)
in the _Amazon EC2 Instance Types_ reference.

###### Software configuration

The easiest way to get started with the latest generation GPU accelerated instance types is
to launch an instance from an AWS Deep Learning AMI that's preconfigured with all of the
required software. For the latest AWS Deep Learning AMIs for use with GPU accelerated instance types, see
[P6 Supported DLAMI](../../../dlami/latest/devguide/p6-support-dlami.md) in the _AWS Deep Learning AMIs Developer Guide_.

If you need to build a custom AMI to launch instances that host deep learning or HPC applications,
we recommend that you install the following minimum software versions on top of your base image.

Instance typeNVIDIA driverCUDANVIDIA GDRCopyEFA installerNCCLEFA K8s ¹G7e57512.92.51.45.02.28.30.5.10P553012.12.31.24.12.18.30.4.4P5.4xlarge53012.12.31.43.1 ²2.18.30.4.4P5e55012.12.31.24.12.18.30.5.5P5en55012.12.31.24.12.18.30.5.6P6-B20057012.82.51.41.02.26.2-10.5.10P6e-GB20057012.82.51.41.02.26.2-10.5.10P6-B30058013.02.51.44.02.28.30.5.10

**¹** The **EFA K8s** column contains the minimum recommended
version for `aws-efa-k8s-device-plugin`.

**²** There is compatibility issue that affects `P5.4xlarge` instances
when GPU-to-GPU communication uses Elastic Fabric Adapter (EFA) and the NVIDIA Collective Communications Library (NCCL). To mitigate
the issue, set the environment variable `FI_HMEM_DISABLE_P2P` to `1`, and ensure that you install
EFA version 1.43.1 or newer.

###### Note

If you use version 1.41.0 of the EFA installer, the `aws-ofi-nccl plugin` comes with it.
For earlier versions of the EFA installer, use `aws-ofi-nccl plugin` version `1.7.2-aws`
or later.

We also recommend that you configure the instance to not use deeper C-states. For more
information, see [High performance and low \
latency by limiting deeper C-states](../../../linux/al2/ug/processor-state-control.md#c-states) in the _Amazon Linux 2 User Guide_.
The latest AWS Deep Learning Base GPU AMIs are preconfigured to not use deeper C-states.

For networking and Elastic Fabric Adapter (EFA) configuration see [Maximize network bandwidth on Amazon EC2 instances with multiple network cards](efa-acc-inst-types.md).

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Set up dual 4K displays on G4ad

Mac instances
