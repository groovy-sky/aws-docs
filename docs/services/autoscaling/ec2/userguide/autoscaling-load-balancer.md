# Use Elastic Load Balancing to distribute incoming application traffic in your Auto Scaling group

Elastic Load Balancing automatically distributes your incoming application traffic across all the EC2
instances that you are running. Elastic Load Balancing helps to manage incoming requests by optimally routing
traffic so that no one instance is overwhelmed. To use Elastic Load Balancing with your Auto Scaling group, [attach the\
load balancer to your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html). This registers the group with the load balancer, which acts as a single point of contact for all incoming web traffic to your Auto Scaling group.

When you use Elastic Load Balancing with your Auto Scaling group, it's not necessary to register individual EC2
instances with the load balancer. Instances that are launched by your Auto Scaling group are
automatically registered with the load balancer. Likewise, instances that are terminated by
your Auto Scaling group are automatically deregistered from the load balancer.

After attaching a load balancer to your Auto Scaling group, you can configure your Auto Scaling group to
use Elastic Load Balancing metrics (such as the Application Load Balancer request count per target) to scale the number of
instances in the group as demand fluctuates.

Optionally, you can add Elastic Load Balancing health checks to your Auto Scaling group so that Amazon EC2 Auto Scaling can
identify and replace unhealthy instances based on these additional health checks. Otherwise,
you can create a CloudWatch alarm that notifies you if the healthy host count of the target group
is lower than allowed.

###### Contents

- [Elastic Load Balancing types](#integrations-aws-elastic-load-balancing-types)

- [Prepare to attach a\
load balancer](https://docs.aws.amazon.com/autoscaling/ec2/userguide/getting-started-elastic-load-balancing.html)

- [Attach a load\
balancer](https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)

- [Configure a load balancer](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-create-load-balancer-console.html)

- [Verify the attachment\
status](https://docs.aws.amazon.com/autoscaling/ec2/userguide/load-balancer-status.html)

- [Add an Availability Zone](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-add-az-console.html)

- [Remove an Availability Zone](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-remove-az-console.html)

- [Detach a load balancer](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-remove-load-balancer.html)

- [AWS CLI examples for working with Elastic Load Balancing](https://docs.aws.amazon.com/autoscaling/ec2/userguide/examples-elastic-load-balancing-aws-cli.html)

## Elastic Load Balancing types

Elastic Load Balancing provides four types of load balancers that can be used with your Auto Scaling group:
Application Load Balancers, Network Load Balancers, Gateway Load Balancers, and Classic Load Balancers.

There is a key difference in how the load balancer types are configured. With Application Load Balancers,
Network Load Balancers, and Gateway Load Balancers, instances are registered as targets with a target group, and you
route traffic to the target group. With Classic Load Balancers, instances are registered directly with
the load balancer.

Application Load Balancer

Routes and load balances at the application layer (HTTP/HTTPS), and
supports path-based routing. An Application Load Balancer can route requests to ports on one or
more registered targets, such as EC2 instances, in your virtual private
cloud (VPC).

Network Load Balancer

Routes and load balances at the transport layer (TCP/UDP Layer-4), based
on address information extracted from the Layer-4 header. Network Load Balancers can handle
traffic bursts, retain the source IP of the client, and use a fixed IP for
the life of the load balancer.

Gateway Load Balancer

Distributes traffic to a fleet of appliance instances. Provides scale,
availability, and simplicity for third-party virtual appliances, such as
firewalls, intrusion detection and prevention systems, and other appliances.
Gateway Load Balancers work with virtual appliances that support the GENEVE protocol.
Additional technical integration is required, so make sure to consult the
user guide before choosing a Gateway Load Balancer.

Classic Load Balancer

Routes and load balances either at the transport layer (TCP/SSL), or at
the application layer (HTTP/HTTPS).

To gain a deeper understanding of the different types of load balancers available, see
the following resources:

- [What is Elastic Load Balancing?](https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/what-is-load-balancing.html)

- [What is an Application Load Balancer?](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/introduction.html)

- [What is a Network Load Balancer?](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html)

- [What is a Gateway Load Balancer?](https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/introduction.html)

- [What is a Classic Load Balancer?](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/introduction.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Compute Optimizer

Prepare to attach a
load balancer
