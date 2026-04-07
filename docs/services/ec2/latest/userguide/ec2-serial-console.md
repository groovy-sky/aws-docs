# EC2 Serial Console for instances

With the EC2 serial console, you have access to your Amazon EC2 instance's serial port, which you
can use to troubleshoot boot, network configuration, and other issues. The serial console does
not require your instance to have any networking capabilities. With the serial console, you can
enter commands to an instance as if your keyboard and monitor are directly attached to the
instance's serial port. The serial console session lasts during instance reboot and stop. During
reboot, you can view all of the boot messages from the start.

Access to the serial console is not available by default. Your organization must grant
account access to the serial console and configure IAM policies to grant your users access to
the serial console. Serial console access can be controlled at a granular level by using
instance IDs, resource tags, and other IAM levers. For more information, see [Configure access to the EC2 Serial Console](configure-access-to-serial-console.md).

The serial console can be accessed by using the EC2 console or the AWS CLI.

The serial console is available at no additional cost.

###### Topics

- [Prerequisites for the EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-serial-console-prerequisites.html)

- [Configure access to the EC2 Serial Console](configure-access-to-serial-console.md)

- [Connect to the EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/connect-to-serial-console.html)

- [Disconnect from the EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/disconnect-serial-console-session.html)

- [Troubleshoot your Amazon EC2 instance using the EC2 Serial Console](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/troubleshoot-using-serial-console.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshoot using EC2Rescue and Systems Manager

Prerequisites
