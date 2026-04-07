# IAM roles for Amazon ECS

An IAM role is an IAM identity that you can create in your account that has specific
permissions. In Amazon ECS, you can create roles to grant permissions to Amazon ECS resource such as
containers or services.

The roles Amazon ECS requires depend on the task definition launch type and the features that
you use. Use the following table to determine which IAM roles you need for Amazon ECS.

RoleDefinitionWhen requiredMore informationTask execution roleThis role allows Amazon ECS to use other AWS services on your behalf.

Your task is hosted on AWS Fargate
or on external instances and:

- pulls a container image from an Amazon ECR private
repository.

- pulls a container image from an Amazon ECR private
repository in a different account from the account that
runs the task.

- sends container logs to CloudWatch Logs using the
`awslogs` log driver.

Your task is hosted on either AWS Fargate or Amazon EC2 instances
and:

- uses private registry authentication.

- uses Runtime Monitoring.

- the task definition references sensitive data using
Secrets Manager secrets or AWS Systems Manager Parameter Store
parameters.

[Amazon ECS task execution IAM role](task-execution-iam-role.md)Task roleThis role allows your application code (on the container) to use other AWS services.Your application accesses other AWS services, such as Amazon S3.[Amazon ECS task IAM role](task-iam-roles.md)Container instance roleThis role allows your EC2 instances or external instances to register
with the cluster.Your task is hosted on Amazon EC2 instances or an external instance.[Amazon ECS container instance IAM role](instance-iam-role.md)Amazon ECS Anywhere roleThis role allows your external instances to access AWS APIs.Your task is hosted on external instances.[Amazon ECS Anywhere IAM role](iam-role-ecsanywhere.md)Amazon ECS infrastructure for load balancers roleThis role allows Amazon ECS to manage load balancer resources in your clusters on your behalf for blue/green deployments.You want to use Amazon ECS blue/green deployments.[Amazon ECS infrastructure IAM role for load balancers](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AmazonECSInfrastructureRolePolicyForLoadBalancers.html)Amazon ECS CodeDeploy roleThis role allows CodeDeploy to make updates to your services.You use the CodeDeploy blue/green deployment type to deploy services.[Amazon ECS CodeDeploy IAM Role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/codedeploy_IAM_role.html)Amazon ECS EventBridge roleThis role allows EventBridge to make updates to your services.You use the EventBridge rules and targets to schedule your tasks.[Amazon ECS EventBridge IAM Role](cwe-iam-role.md)Amazon ECS infrastructure roleThis role allows Amazon ECS to manage infrastructure resources in your
clusters.

- You want to use Amazon ECS Managed Instances for your capacity.

- You want to attach Amazon EBS volumes to your Fargate or EC2 launch
type Amazon ECS tasks. The infrastructure role allows Amazon ECS to manage Amazon EBS volumes for
your tasks.

- You want to use Transport Layer Security (TLS) to encrypt traffic between your
Amazon ECS Service Connect services.

- You want to create VPC Lattice target groups.

[Amazon ECS infrastructure IAM role](infrastructure-iam-role.md)Instance profileThis role allows allows Amazon ECS Managed Instances to assume the infrastructure role securely.You use Amazon ECS Managed Instances in your clusters.[Amazon ECS Managed Instances instance profile](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/managed-instances-instance-profile.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Allow Amazon ECS to manage Amazon ECS Managed Instances

Best practices for IAM roles
