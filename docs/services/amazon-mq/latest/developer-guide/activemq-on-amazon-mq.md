# ActiveMQ tutorials

The following tutorials show how you can create and connect to your ActiveMQ brokers. To use the ActiveMQ Java example code, you must install the
[Java Standard Edition Development Kit](https://www.oracle.com/technetwork/java/javase/downloads/index.html) and make some
changes to the code

###### Topics

- [Creating and configuring an Amazon MQ network of brokers](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-creating-configuring-network-of-brokers.html)

- [Connecting a Java application to your Amazon MQ broker](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-connecting-application.html)

- [Integrating ActiveMQ brokers with LDAP](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/security-authentication-authorization.html)

- [Step 3: (Optional) Connect to an AWS Lambda function](#activemq-connect-to-lambda)

- [Creating an ActiveMQ broker user](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-listing-managing-users.html)

- [Edit an ActiveMQ broker user](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/edit-existing-user-console.html)

- [Delete an ActiveMQ broker user](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/delete-existing-user-console.html)

- [Working examples of using Java Message Service (JMS) with ActiveMQ](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-working-java-example.html)

## Step 3: (Optional) Connect to an AWS Lambda function

AWS Lambda can connect to and consume messages from your Amazon MQ broker.
When you connect a broker to Lambda, you create an [event source mapping](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html)
that reads messages from a queue and invokes the function [synchronously](https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html). The event
source mapping you create reads messages from your broker in batches and converts them into a Lambda payload in the form of a JSON object.

###### To connect your broker to a Lambda function

1. Add the following IAM role permissions to your Lambda function [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html).

- [mq:DescribeBroker](https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id.html#brokers-broker-id-http-methods)

- [ec2:CreateNetworkInterface](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html)

- [ec2:DeleteNetworkInterface](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteNetworkInterface.html)

- [ec2:DescribeNetworkInterfaces](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeNetworkInterfaces.html)

- [ec2:DescribeSecurityGroups](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html)

- [ec2:DescribeSubnets](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSubnets.html)

- [ec2:DescribeVpcs](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVpcs.html)

- [logs:CreateLogGroup](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogGroup.html)

- [logs:CreateLogStream](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogStream.html)

- [logs:PutLogEvents](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html)

- [secretsmanager:GetSecretValue](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html)

###### Note

Without the necessary IAM permissions, your function will not be able to successfully read records from Amazon MQ resources.

2. (Optional) If you have created a broker without public accessibility, you must do one of the following to allow Lambda to connect to your broker:

- Configure one NAT gateway per public subnet. For more information, see
[Internet and service access for VPC-connected functions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html#vpc-internet) in the _AWS Lambda Developer Guide_.

- Create a connection between your Amazon Virtual Private Cloud (Amazon VPC) and Lambda using a VPC endpoint. Your Amazon VPC must also connect to AWS Security Token Service (AWS STS) and Secrets Manager endpoints.
For more information, see [Configuring interface VPC endpoints for Lambda](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc-endpoints.html)
in the _AWS Lambda Developer Guide_.

3. [Configure your broker as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping) for a Lambda function using the AWS Management Console. You can also use the
    [`create-event-source-mapping`](https://docs.aws.amazon.com/cli/latest/reference/lambda/create-event-source-mapping.html) AWS Command Line Interface command.

4. Write some code for your Lambda function to process the messages consumed from your broker. The Lambda payload that
    retrieved by your event source mapping depends on the engine type of the broker. The following is an example of a Lambda payload for an Amazon MQ for ActiveMQ queue.

###### Note

In the example, `testQueue` is the name of the queue.

```json

{
     "eventSource": "aws:amq",
     "eventSourceArn": "arn:aws:mq:us-west-2:112556298976:broker:test:b-9bcfa592-423a-4942-879d-eb284b418fc8",
     "messages": {
       [
         {
           "messageID": "ID:b-9bcfa592-423a-4942-879d-eb284b418fc8-1.mq.us-west-2.amazonaws.com-37557-1234520418293-4:1:1:1:1",
           "messageType": "jms/text-message",
           "data": "QUJDOkFBQUE=",
           "connectionId": "myJMSCoID",
           "redelivered": false,
           "destination": {
             "physicalname": "testQueue"
           },
           "timestamp": 1598827811958,
           "brokerInTime": 1598827811958,
           "brokerOutTime": 1598827811959
         },
         {
           "messageID": "ID:b-9bcfa592-423a-4942-879d-eb284b418fc8-1.mq.us-west-2.amazonaws.com-37557-1234520418293-4:1:1:1:1",
           "messageType":"jms/bytes-message",
           "data": "3DTOOW7crj51prgVLQaGQ82S48k=",
           "connectionId": "myJMSCoID1",
           "persistent": false,
           "destination": {
             "physicalname": "testQueue"
           },
           "timestamp": 1598827811958,
           "brokerInTime": 1598827811958,
           "brokerOutTime": 1598827811959
         }
       ]
     }
}
```

For more information about connecting Amazon MQ to Lambda, the options Lambda supports for an Amazon MQ event source, and event source mapping errors, see
[Using Lambda with Amazon MQ](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html) in the _AWS Lambda Developer Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Metrics

Creating
and configuring a network of brokers
