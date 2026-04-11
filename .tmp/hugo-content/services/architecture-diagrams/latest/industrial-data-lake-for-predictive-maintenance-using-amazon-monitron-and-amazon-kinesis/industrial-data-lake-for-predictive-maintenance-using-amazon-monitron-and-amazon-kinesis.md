# Industrial Data Lake for Predictive Maintenance using Amazon Monitron and Amazon Kinesis

Publication date: **September 28, 2022 ( [Diagram history](#diagram-history))**

This architecture diagram shows you how to build a data lake using AWS IoT sensors, real-time data streams, alerts, visualization, and integrated workflow with Enterprise Resource Planning (ERP) to analyze factory data for predictive maintenance and improve equipment uptime.

## Industrial Data Lake for Predictive Maintenance using Amazon Monitron and Amazon Kinesis Diagram

![Reference architecture diagram showing how to build a data lake using AWS IoT sensors, real-time data streams, alerts, visualization, and integrated workflow with Enterprise Resource Planning (ERP) to analyze factory data for predictive maintenance and improve equipment uptime.](https://docs.aws.amazon.com/images/architecture-diagrams/latest/industrial-data-lake-for-predictive-maintenance-using-amazon-monitron-and-amazon-kinesis/images/industrial-data-lake-for-predictive-maintenance.png)

01. Install **Amazon Monitron** sensors and gateway in a factory.

02. Create **Amazon Kinesis Data Streams** using **Amazon Monitron** as the data source.

03. Configure **Amazon Kinesis Data Streams** from **Amazon Monitron** managed account to customer account.

04. Configure **Amazon Simple Storage Service** (Amazon S3) bucket as delivery
     destination of **Amazon Data Firehose**. **Amazon S3** serves as storage foundation for industrial data lake.

05. Configure **Amazon S3** notifications to send events to the
     **Amazon EventBridge** destination.

06. Configure an **AWS Lambda** function as the target of
     **Amazon EventBridge** destination rules. The
     **Lambda** function processes the **Amazon S3** event and sends it to an **AWS IoT Events** state machine.

07. **AWS IoT Events** responds to sensor warning
     state and creates ERP work order using **AWS Lambda**.

08. **AWS IoT Events** responds to the sensor
     warning state and notifies personnel using **Amazon Simple Notification Service**
     (Amazon SNS) topic via SMS, mobile push, and email.

09. Connect **AWS Glue** data pipeline to
     **Amazon S3** bucket and schedule Glue job via **Amazon EventBridge**. **Amazon Athena** then queries S3 data as reports and metrics.

10. Visualize IoT metrics and state from **Athena** queries
     using **Amazon Managed Grafana**.

## Download editable diagram

To customize this reference architecture diagram based on your business needs, [download the ZIP file](https://docs.aws.amazon.com/architecture-diagrams/latest/industrial-data-lake-for-predictive-maintenance-using-amazon-monitron-and-amazon-kinesis/samples/industrial-data-lake-for-predictive-maintenance.zip) which contains an editable PowerPoint.

## Create a free AWS account

[![Sign up for a free AWS account](https://docs.aws.amazon.com/images/architecture-diagrams/latest/industrial-data-lake-for-predictive-maintenance-using-amazon-monitron-and-amazon-kinesis/images/signup.png)](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)

Sign up for an AWS account. New accounts include 12 months of [AWS Free Tier](https://aws.amazon.com/free) access, including the use of Amazon EC2, Amazon S3, and
Amazon DynamoDB.

## Further reading

For additional information, refer to

- [AWS Architecture\
Icons](https://aws.amazon.com/architecture/icons)

- [AWS Architecture Center](https://aws.amazon.com/architecture)

- [AWS Well-Architected](https://aws.amazon.com/architecture/well-architected)

## Diagram history

To be notified about updates to this reference architecture diagram, subscribe to the RSS feed.

ChangeDescriptionDate

Initial publication

Reference architecture diagram first published.

September 22, 2021

###### Note

To subscribe to RSS updates, you must have an RSS plugin enabled for the browser you are using.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

All content copied from https://docs.aws.amazon.com/.
