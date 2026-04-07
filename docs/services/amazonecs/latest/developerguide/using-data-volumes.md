# Storage options for Amazon ECS tasks

Amazon ECS provides you with flexible, cost effective, and easy-to-use data storage options
depending on your needs. Amazon ECS supports the following data volume options for
containers:

Data volumeSupported capacitySupported operating systemsStorage persistenceUse casesAmazon Elastic Block Store (Amazon EBS)Fargate, Amazon EC2, Amazon ECS Managed InstancesLinux, Windows (on Amazon EC2 only)Can be persisted when attached to a standalone task. Ephemeral when
attached to a task maintained by a service.Amazon EBS volumes provide cost-effective, durable, high-performance block
storage for data-intensive containerized workloads. Common use cases include
transactional workloads such as databases, virtual desktops and root
volumes, and throughput intensive workloads such as log processing and ETL
workloads. For more information, see [Use Amazon EBS volumes with Amazon ECS](ebs-volumes.md).Amazon Elastic File System (Amazon EFS)Fargate, Amazon EC2, Amazon ECS Managed InstancesLinuxPersistentAmazon EFS volumes provide simple, scalable, and persistent shared file
storage for use with your Amazon ECS tasks that grows and shrinks automatically
as you add and remove files. Amazon EFS volumes support concurrency and are
useful for containerized applications that scale horizontally and need
storage functionalities like low latency, high throughput, and
read-after-write consistency. Common use cases include workloads such as
data analytics, media processing, content management, and web serving. For
more information, see [Use Amazon EFS volumes with Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html).Amazon FSx for Windows File ServerAmazon EC2WindowsPersistentFSx for Windows File Server volumes provide fully managed Windows file servers that you
can use to provision your Windows tasks that need persistent, distributed,
shared, and static file storage. Common use cases include .NET applications
that might require local folders as persistent storage to save application
outputs. Amazon FSx for Windows File Server offers a local folder in the container which
allows for multiple containers to read-write on the same file system that's
backed by a SMB Share. For more information, see [Use FSx for Windows File Server volumes with Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/wfsx-volumes.html).Amazon FSx for NetApp ONTAPAmazon EC2LinuxPersistentAmazon FSx for NetApp ONTAP volumes provide fully managed NetApp ONTAP file systems that you
can use to provision your Linux tasks that need persistent, high-performance,
and feature-rich shared file storage. Amazon FSx for NetApp ONTAP supports NFS and SMB protocols
and provides enterprise-grade features like snapshots, cloning, and data
deduplication. Common use cases include high-performance computing workloads,
content repositories, and applications requiring POSIX-compliant shared storage.
For more information, see [Mounting Amazon FSx for NetApp ONTAP file systems from Amazon ECS containers](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/mount-ontap-ecs-containers.html).Docker volumesAmazon EC2Windows, LinuxPersistentDocker volumes are a feature of the Docker container runtime that allow
containers to persist data by mounting a directory from the file system of
the host. Docker volume drivers (also referred to as plugins) are used to
integrate container volumes with external storage systems. Docker volumes
can be managed by third-party drivers or by the built in `local`
driver. Common use cases for Docker volumes include providing persistent data volumes or
sharing volumes at different locations on different containers on the same
container instance. For more information, see [Use Docker volumes with Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-volumes.html).Bind mountsFargate, Amazon EC2, Amazon ECS Managed InstancesWindows, LinuxEphemeralBind mounts consist of a file or directory on the host, such as an Amazon EC2
instance or AWS Fargate, that is mounted onto a container. Common use cases for bind mounts
include sharing a volume from a source container with other containers in
the same task, or mounting a host volume or an empty volume in one or more
containers. For more information, see [Use bind mounts with Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/bind-mounts.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Task networking for Fargate

Amazon EBS
