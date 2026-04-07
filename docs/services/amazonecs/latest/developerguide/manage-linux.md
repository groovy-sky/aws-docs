# Amazon ECS Linux container instance management

When you use EC2 instances for your Amazon ECS workloads, you are responsible for maintaining the instances

###### Management procedures

- [Launching a container instance](launch-container-instance.md)

- [Bootstrapping Linux container instances](bootstrap-container-instance.md)

- [Configuring container instances to receive Spot Instance notices](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/spot-instance-draining-linux-container.html)

- [Running a script when you launch a container instance](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/start_task_at_launch.html)

- [Increasing Amazon ECS Linux container instance network interfaces](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/container-instance-eni.html)

- [Reserving container instance memory](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/memory-management.html)

- [Manage container instances remotely](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ec2-run-command.html)

- [Using an HTTP proxy for Linux container instances](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/http_proxy_config.html)

- [Configuring pre-initialized instances for your Auto Scaling\
group](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-warm-pool.html)

- [Updating the Amazon ECS container agent](ecs-agent-update.md)

Each Amazon ECS container agent version supports a different feature set and provides bug
fixes from previous versions. When possible, we always recommend using the latest version
of the Amazon ECS container agent. To update your container agent to the latest version, see
[Updating the Amazon ECS container agent](ecs-agent-update.md).

To see which features and enhancements are included with each agent release, see [https://github.com/aws/amazon-ecs-agent/releases](https://github.com/aws/amazon-ecs-agent/releases).

###### Important

The minimum Docker version for reliable metrics is Docker version `v20.10.13` and
newer, which is included in Amazon ECS-optimized AMI `20220607` and newer.

Amazon ECS agent versions `1.20.0` and newer have deprecated support for
Docker versions older than `18.01.0`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Launching a Bottlerocket
instance

Launching a container instance
