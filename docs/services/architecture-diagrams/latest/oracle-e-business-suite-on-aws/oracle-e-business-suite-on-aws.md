# Oracle E-Business Suite on AWS

Publication date: **August 3, 2023 ( [Diagram history](#diagram-history))**

This reference architecture can be used to deploy Oracle E-Business Suite on AWS by showing how to configure high availability for the database and application tier.

## Oracle E-Business Suite on AWS Diagram

![Reference architecture diagram showing how to deploy Oracle E-Business Suite on AWS by showing how to configure high availability for the database and application tier.](https://docs.aws.amazon.com/images/architecture-diagrams/latest/oracle-e-business-suite-on-aws/images/oracle-e-business-suite-on-aws.png)

01. The architecture starts with a single Region and single Virtual
     Private Cloud (VPC) on-par with the on-premises data center.

02.
     Multiple Availability Zones (AZs) provide resilience and high
     availability for the production workload.

03. **Application Load Balancer** distributes network traffic to improve
     the scalability and availability of your applications across multiple AZs.

04. **AWS WAF** is the web application firewall
     that protects the Oracle E-Business Suite applications against common web exploits.

05. **Amazon Route 53** provides domain name service (DNS) configuration.

06. **Amazon WorkSpaces** provides a desktop experience in the cloud.
     Use **AWS Directory Service** to enable user authentication.

07. **Amazon Simple Storage Service** (Amazon S3) is used for storing backups, files, and objects.

08. **Amazon CloudWatch** is used for application logging, monitoring, and alarms.

09. **AWS Systems Manager** provides bastion-less access to instances in
     private subnet, along with management and monitoring capabilities.

10. **AWS Backup** is a fully managed service that
     enables you to centralize and automate data protection across on-premise and
     AWS services.

## Download editable diagram

To customize this reference architecture diagram based on your business needs, [download the ZIP file](https://docs.aws.amazon.com/architecture-diagrams/latest/oracle-e-business-suite-on-aws/samples/oracle-e-business-suite-on-aws.zip) which contains an editable PowerPoint.

## Create a free AWS account

[![Sign up for a free AWS account](https://docs.aws.amazon.com/images/architecture-diagrams/latest/oracle-e-business-suite-on-aws/images/signup.png)](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)

Sign up for an AWS account. New accounts include 12 months of [AWS Free Tier](https://aws.amazon.com/free) access, including the use of Amazon EC2, Amazon S3, and
Amazon DynamoDB.

## Further reading

For additional information, refer to

- [AWS Architecture\
Icons](https://aws.amazon.com/architecture/icons)

- [AWS Architecture Center](https://aws.amazon.com/architecture)

- [AWS Well-Architected](https://aws.amazon.com/architecture/well-architected)

## Contributors

Contributors to this reference architecture diagram include:

- Joyjeet Banerjee, Senior Partner Solutions Architect, Amazon Web Services

## Diagram history

To be notified about updates to this reference architecture diagram, subscribe to the RSS feed.

ChangeDescriptionDate

Initial publication

Reference architecture diagram first published.

August 3, 2023

###### Note

To subscribe to RSS updates, you must have an RSS plugin enabled for the browser you are using.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

All content copied from https://docs.aws.amazon.com/.
