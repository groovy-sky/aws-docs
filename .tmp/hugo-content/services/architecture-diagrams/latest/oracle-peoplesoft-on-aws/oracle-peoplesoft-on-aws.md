# Oracle PeopleSoft on AWS

Publication date: **March 21, 2023 ( [Diagram history](#diagram-history))**

This architecture shows how to deploy a highly available and resilient Oracle PeopleSoft production environment on AWS.

## Oracle PeopleSoft on AWS Diagram

![Reference architecture diagram showing how to deploy a highly available and resilient Oracle PeopleSoft production environment on AWS.](https://docs.aws.amazon.com/images/architecture-diagrams/latest/oracle-peoplesoft-on-aws/images/oracle-peoplesoft-on-aws.png)

01. A single Region and single Virtual Private Cloud (VPC) on-par with the on-premises data centre.

02. Multiple Availability Zones (AZs) provide resilience and high availability for the production workload.

03. Application Load Balancer (ALB) distributes network traffic to improve the scalability and availability of your applications across multiple AZs.

04. **AWS WAF** is the web application firewall that protects the PeopleSoft against common web exploits.

05. **Amazon Route 53** provides domain name service (DNS) configuration.

06. **Amazon WorkSpaces** provides users with a desktop experience in the cloud. Use **AWS Directory Service** to enable user authentication.

07. **Amazon Simple Storage Service** (Amazon S3) for storing backups, files, objects.

08. **Amazon CloudWatch** is used for application logging, monitoring, and alarms.

09. **AWS Systems Manager** provides bastion-less access to instances in private subnet along with management and monitoring capabilities.

10. **AWS Backup** is a fully managed service that enables you to centralize and automate data protection across on-premises and AWS services.

## Download editable diagram

To customize this reference architecture diagram based on your business needs, [download the ZIP file](https://docs.aws.amazon.com/architecture-diagrams/latest/oracle-peoplesoft-on-aws/samples/oracle-peoplesoft-on-aws.zip) which contains an editable PowerPoint.

## Create a free AWS account

[![Sign up for a free AWS account](https://docs.aws.amazon.com/images/architecture-diagrams/latest/oracle-peoplesoft-on-aws/images/signup.png)](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)

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

- Gaurav Gupta, Senior Partner Solutions Architect, Amazon Web Services

## Diagram history

To be notified about updates to this reference architecture diagram, subscribe to the RSS feed.

ChangeDescriptionDate

Initial publication

Reference architecture diagram first published.

March 21, 2023

###### Note

To subscribe to RSS updates, you must have an RSS plugin enabled for the browser you are using.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

All content copied from https://docs.aws.amazon.com/.
