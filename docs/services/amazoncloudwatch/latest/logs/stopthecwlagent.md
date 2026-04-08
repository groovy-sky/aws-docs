# Stop the CloudWatch Logs agent

Use the following procedure to stop the CloudWatch Logs agent on your EC2 instance.

###### To stop the agent

1. Connect to your EC2 instance. For more information, see [Connect to Your\
    Instance](../../../ec2/latest/userguide/ec2-connect-to-instance-linux.md) in the _Amazon EC2 User Guide_.

For more information about connection issues, see [Troubleshooting Connecting to Your Instance](../../../ec2/latest/userguide/troubleshootinginstancesconnecting.md) in the
    _Amazon EC2 User Guide_.

2. At a command prompt, type the following command:

```nohighlight

sudo service awslogs stop
```

If you are running Amazon Linux 2, type the following command:

```nohighlight

sudo service awslogsd stop
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Start the CloudWatch Logs agent

CloudWatch Logs agent reference

All content copied from https://docs.aws.amazon.com/.
