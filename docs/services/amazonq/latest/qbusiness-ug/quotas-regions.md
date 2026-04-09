# Service quotas for Amazon Q Business

The following are the service endpoints and service quotas for Amazon Q Business To
connect programmatically to Amazon Q Business, you use an endpoint. For more
information, see [AWS\
service endpoints](../../../../general/latest/gr/rande.md) in the _AWS General Reference_.
Service quotas, also referred to as limits, are the maximum number of service resources or
operations for your AWS account. For more information, see [AWS service\
quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.

## Supported Regions

The following table shows the AWS Regions and endpoints currently supported by
Amazon Q Business.

Region nameRegionEndpointProtocolUS East (N. Virginia)us-east-1

qbusiness.us-east-1.api.aws

qbusiness-fips.us-east-1.api.aws

qbusiness-websocket.us-east-1.api.aws

qbusiness-websocket-fips.us-east-1.api.aws

HTTPSUS West (Oregon)us-west-2

qbusiness.us-west-2.api.aws

qbusiness-fips.us-west-2.api.aws

qbusiness-websocket.us-west-2.api.aws

qbusiness-websocket-fips.us-west-2.api.aws

HTTPSEurope (Ireland)eu-west-1

qbusiness-eu-west-1.api.aws

qbusiness-websocket.eu-west-1.api.aws

HTTPSAsia Pacific (Sydney)ap-southeast-2

qbusiness.ap-southeast-2.api.aws

qbusiness-websocket.ap-southeast-2.api.aws

HTTPS

###### Important

Amazon Q Business Pro tier subscriptions in Europe (Ireland)
(eu-west-1) and Asia Pacific (Sydney) (ap-southeast-2) regions are available with a limited set of features.

###### Note

The Europe (Ireland) and the Asia Pacific (Sydney) regions don't currently
support all features available in the US regions, such as Q App, Q Actions, and
Audio/Video files. While these features will become available soon, Amazon Q Business
customers in this region can do the following:

- Get answers to questions submitted to the enterprise retrieval augmented
generation system.

- Generate content through Amazon Q Business assistant

- Access capabilities such as embedded images in files.

- Perform tabular search on small tables.

- Ingest data from scanned PDFs.

- Answer questions from data in scanned PDFs.

- Respond to queries to LLM knowledge.

For a list of AWS regions where Amazon Q Business is available,
see [Amazon Q Business regions and endpoints](../../../../general/latest/gr/amazonq.md) in the _Amazon Web_
_Services General Reference_.

## Quotas

Your AWS account has default quotas, formerly referred to as limits, for each
AWS service. Unless otherwise noted, each quota is Region-specific. You can request
increases for some quotas, and other quotas can't be increased. To see whether a quota
can be adjusted, refer to the **Adjustable** column in the
following table.

To view the quotas for Amazon Q Business, or request a quota increase, open the
[Service Quotas console](https://console.aws.amazon.com/servicequotas/home). In the
navigation pane, choose **AWS services** and select **Amazon Q**.

The following table shows the quotas that are related to Amazon Q Business for
your AWS account.

NameDefaultAdjustableMaximum number of data accessors per Amazon Q Business
application environment10NoMaximum number of applications per account50NoMaximum number of data sources per application50NoMaximum number of plugins per application25NoMaximum number of actions per plugin20NoMaximum number of transactions per second (TPS) to list actions
configured for a specific plugin10YesMaximum number of transactions per second (TPS) to list actions and
metadata configured for all plugins5YesMaximum number of pages per PDF for optical character recognition
(OCR) or insight extraction from embedded visual elements.3000NoMaximum number of pages per Word document for insight extraction from
embedded visual elements.3000NoMaximum number of pages per Powerpoint document for insight
extraction from embedded visual elements.3000NoMaximum number of images per PDF, Word or Powerpoint document for
insight extraction from embedded visual elements.500NoMaximum number of queries per second (QPS) per index

This limit
applies to the underlying index for your application across all user
queries. The SearchRelevantContent, Chat, and ChatSync APIs are
throttled differently. For more information, see [Chat and conversation\
management](conversation-api.md).

1YesMaximum number of documents that can be uploaded during a
conversation or chat session5NoMaximum file size per document upload during a conversation or chat
session10 MBNoMaximum number of total Amazon Q Apps created within an
application environment1000NoMaximum number of Amazon Q Apps that a single web experience user can
create100NoMaximum number of cards used per Amazon Q Apps 35NoMaximum character length for an Amazon Q Apps Creator prompt10,000NoMaximum number of file upload cards per Amazon Q Apps20NoMaximum file size per file upload card 50 MBNoMaximum frame rate for video files60 FPSNoMaximum file size limit for video files extraction10 GBNoMaximum file size limit for audio files extraction 2 GBNoFile size limit in Amazon Q Business File Upload console (all types) 50 MBNoMaximum duration limit for video files extraction 4 hoursNoMaximum duration limit for audio files extraction 4 hoursNoMaximum number of groups per user supported by Amazon Q Business1000No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Maintenance updates

API reference

All content copied from https://docs.aws.amazon.com/.
