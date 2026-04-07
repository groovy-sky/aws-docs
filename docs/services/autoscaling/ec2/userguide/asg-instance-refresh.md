# Use an instance refresh to update instances in an Auto Scaling group

You can use an instance refresh to update the instances in your Auto Scaling group. This
feature can be useful when a configuration change requires you to replace instances or their root volumes,
especially if your Auto Scaling group contains a large number of instances.

Some situations where an instance refresh can help include:

- Deploying a new Amazon Machine Image (AMI) or user data script across your
Auto Scaling group. You can create a new launch template with the changes and then use
an instance refresh to roll out the updates immediately.

- Migrating your instances to new instance types to take advantage of the latest
improvements and optimizations.

- Switching your Auto Scaling groups from using a launch configuration to using a launch
template. You can copy your launch configurations to launch templates and then
use an instance refresh to update your instances to the new templates. For more
information about migrating to launch templates, see [Migrate your Auto Scaling groups to launch templates](migrate-to-launch-templates.md).

- Applying security patches or software updates while preserving long-running instance
state and avoiding capacity constraints with specialized instance types like GPU or
Mac instances.

###### Contents

- [How an instance refresh works](https://docs.aws.amazon.com/autoscaling/ec2/userguide/instance-refresh-overview.html)

- [Understand the default values](https://docs.aws.amazon.com/autoscaling/ec2/userguide/understand-instance-refresh-default-values.html)

- [Start an instance refresh](https://docs.aws.amazon.com/autoscaling/ec2/userguide/start-instance-refresh.html)

- [Monitor an instance refresh](https://docs.aws.amazon.com/autoscaling/ec2/userguide/check-status-instance-refresh.html)

- [Replace root volumes](https://docs.aws.amazon.com/autoscaling/ec2/userguide/replace-root-volume.html)

- [Cancel an instance refresh](https://docs.aws.amazon.com/autoscaling/ec2/userguide/cancel-instance-refresh.html)

- [Undo changes with a rollback](https://docs.aws.amazon.com/autoscaling/ec2/userguide/instance-refresh-rollback.html)

- [Use skip matching](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh-skip-matching.html)

- [Add checkpoints](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-adding-checkpoints-instance-refresh.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Replace your instances

How an instance refresh works
