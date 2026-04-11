# Flight Information Management System (FIMS)

Publication date: **February 2, 2022 ( [Diagram history](#diagram-history))**

This architecture enables you to manage and operate services required for unmanned aerial
systems on AWS.

## Flight Information Management System (FIMS) Diagram

![Reference architecture diagram showing how you can manage and operate services required for unmanned aerial systems on AWS.](https://docs.aws.amazon.com/images/architecture-diagrams/latest/flight-information-management-system/images/flight-information-management-system.png)

01. Operators access the FIMS services through a web-based portal for flight planning and
     additional services.

02. Flight plans are sent to appropriate microservices and **Amazon Elastic Container Service** (Amazon ECS) is leveraged for processing approvals.

03. Persistent relational data such as registration and flight plan are stored in
     **Amazon Aurora**.

04. Real-time data (such as telemetry, position, track, and velocity) is stored in
     **DynamoDB** .

05. Cached data is stored through **Amazon ElastiCache** for quick
     retrieval.

06. Processed data is encrypted and stored in data lake, and managed by **AWS Lake Formation**, in accordance with regulatory compliance.

07. Virtualization tools to view real-time positioning, flight paths, potential conflicts,
     etc. can be deployed using **Quick** and **Amazon Managed Grafana**.

08. Flight plan approvals automatically provided to UAS operators by FIMS.

09. UAS operators request/receive manual approval from FIMS Admins for flight plans when
     conflicts arise.

10. UAS provider data stored in **Amazon Aurora** and real-time
     data stored in **DynamoDB** databases.

11. Data lake built on **Amazon S3** is used for secure data
     archival, analytics, and visualization.

12. Supplemental data from weather services is ingested into FIMS for smart
     decisioning.

## Download editable diagram

To customize this reference architecture diagram based on your business needs, [download the ZIP file](https://docs.aws.amazon.com/architecture-diagrams/latest/flight-information-management-system/samples/flight-information-management-system.zip) which contains an editable PowerPoint.

## Create a free AWS account

[![Sign up for a free AWS account](https://docs.aws.amazon.com/images/architecture-diagrams/latest/flight-information-management-system/images/signup.png)](https://portal.aws.amazon.com/gp/aws/developer/registration/index.html)

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

February 2, 2022

###### Note

To subscribe to RSS updates, you must have an RSS plugin enabled for the browser you are using.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

All content copied from https://docs.aws.amazon.com/.
