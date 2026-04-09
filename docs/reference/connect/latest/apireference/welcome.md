# Welcome

## Amazon Connect Service

- [Amazon Connect actions](api-operations-amazon-connect-service.md)

- [Amazon Connect data types](api-types-amazon-connect-service.md)

Amazon Connect is a cloud-based contact center solution that you use to set up and manage a customer contact
center and provide reliable customer engagement at any scale.

Amazon Connect provides metrics and real-time reporting that enable you to optimize contact routing. You can
also resolve customer issues more efficiently by getting customers in touch with the appropriate agents.

There are limits to the number of Amazon Connect resources that you can create. There are also limits to the
number of requests that you can make per second. For more information, see [Amazon Connect Service Quotas](../../../../services/connect/latest/adminguide/amazon-connect-service-limits.md) in the
_Amazon Connect Administrator Guide_.

You can use an endpoint to connect programmatically to an AWS service. For a list of Amazon Connect endpoints, see [Amazon Connect\
Endpoints](../../../../general/latest/gr/connect-region.md).

## Amazon AppIntegrations Service

- [Amazon AppIntegrations\
actions](api-operations-amazon-appintegrations-service.md)

- [Amazon AppIntegrations\
data types](api-types-amazon-appintegrations-service.md)

The Amazon AppIntegrations service enables you to configure and reuse connections to external
applications.

For information about how you can use external applications with Amazon Connect, see
the following topics in the _Amazon Connect Administrator_
_Guide_:

- [Third-party\
applications (3p apps) in the agent workspace](../../../../services/connect/latest/adminguide/3p-apps.md)

- [Use\
Amazon Q in Connect for generative AI–powered agent assistance in\
real-time](../../../../services/connect/latest/adminguide/amazon-q-connect.md)

## Amazon Connect Contact Lens

- [Contact Lens actions](api-operations-amazon-connect-contact-lens.md)

- [Contact Lens data types](api-types-amazon-connect-contact-lens.md)

Amazon Connect Contact Lens enables you to analyze conversations between customer and agents, by using
speech transcription, natural language processing, and intelligent search capabilities.
It performs sentiment analysis, detects issues, and enables you to automatically
categorize contacts.

Amazon Connect Contact Lens provides both real-time and post-call analytics of customer-agent
conversations. For more information, see [Analyze conversations\
using speech analytics](../../../../services/connect/latest/adminguide/analyze-conversations.md) in the _Amazon Connect Administrator_
_Guide_.

## Amazon Connect Outbound Campaigns

- [Outbound campaigns actions](api-operations-amazon-connect-outbound-campaigns.md)

- [Outbound campaigns data types](api-types-amazon-connect-outbound-campaigns.md)

With the outbound campaigns feature of Amazon Connect, you can create high-volume outbound
campaigns. For example, you can use this feature for appointment reminders, telemarketing,
subscription renewals, or debt collection. For more information, see [Set up\
outbound communications](../../../../services/connect/latest/adminguide/outbound-communications.md) in the _Amazon Connect Administrator_
_Guide_.

###### Note

Outbound campaigns version 1 APIs are not supported in the af-south-1 Africa (Cape
Town) region. You can use the Outbound campaigns version 2 API [actions](api-operations-amazon-connect-outbound-campaigns-v2.md) and [data types](api-types-amazon-connect-outbound-campaigns-v2.md) in this region.

## Amazon Connect Outbound Campaigns V2

- [Outbound campaignsV2 actions](api-operations-amazon-connect-outbound-campaigns-v2.md)

- [Outbound campaignsV2 data types](api-types-amazon-connect-outbound-campaigns-v2.md)

With the outbound campaignsV2 feature of Amazon Connect, you can create high-volume outbound
campaigns. For example, you can use this feature for appointment reminders, telemarketing,
subscription renewals, or debt collection. For more information, see [Set up\
outbound communications](../../../../services/connect/latest/adminguide/outbound-communications.md) in the _Amazon Connect Administrator_
_Guide_.

## Amazon Connect Cases

- [Cases\
actions](api-operations-amazon-connect-cases.md)

- [Cases data\
types](api-types-amazon-connect-cases.md)

With Amazon Connect Cases, your agents can track and manage customer issues that require
multiple interactions, follow-up tasks, and teams in your contact center. A case represents a
customer issue. It records the issue, the steps and interactions taken to resolve the issue,
and the outcome. For more information, see [Amazon Connect Cases](../../../../services/connect/latest/adminguide/cases.md) in the
_Amazon Connect Administrator Guide_.

## Amazon Connect Participant Service

- [Participant Service actions](api-operations-amazon-connect-participant-service.md)

- [Participant Service data types](api-types-amazon-connect-participant-service.md)

Amazon Connect is an easy-to-use omnichannel cloud contact center service that
enables companies of any size to deliver superior customer service at a lower cost.
Amazon Connect communications capabilities make it easy for companies to deliver
personalized interactions across communication channels, including chat.

Use the Amazon Connect Participant Service to manage participants (for example,
agents, customers, and managers listening in), and to send messages and events within a
chat contact. The APIs in the service enable the following: sending chat messages,
attachment sharing, managing a participant's connection state and message events, and
retrieving chat transcripts.

## Amazon Connect Customer Profiles

- [Customer Profiles actions](api-operations-amazon-connect-customer-profiles.md)

- [Customer Profiles data types](api-types-amazon-connect-customer-profiles.md)

Amazon Connect Customer Profiles is a unified customer profile for your contact
center that has pre-built connectors powered by AppFlow that make it easy to combine
customer information from third party applications, such as Salesforce (CRM), ServiceNow
(ITSM), and your enterprise resource planning (ERP), with contact history from your Amazon Connect contact center.

For more information about the Amazon Connect Customer Profiles feature, see [Use Customer\
Profiles](../../../../services/connect/latest/adminguide/customer-profiles.md) in the _Amazon Connect Administrator's Guide_.

## Amazon Q Connect

- [Amazon Q\
actions](api-operations-amazon-q-connect.md)

- [Amazon Q data\
types](api-types-amazon-q-connect.md)

###### Note

**Powered by Amazon Bedrock**: AWS implements [automated abuse detection](../../../../services/bedrock/latest/userguide/abuse-detection.md). Because Amazon Q in Connect is
built on Amazon Bedrock, users can take full advantage of the controls implemented in Amazon Bedrock to enforce safety, security, and
the responsible use of artificial intelligence (AI).

Amazon Q in Connect is a generative AI customer service assistant. It is an LLM-enhanced evolution of Amazon Connect Wisdom that delivers real-time recommendations to help contact center agents resolve customer issues quickly and
accurately.

Amazon Q in Connect automatically detects customer intent during calls and chats using conversational analytics and natural
language understanding (NLU). It then provides agents with immediate, real-time generative responses and suggested
actions, and links to relevant documents and articles. Agents can also query Amazon Q in Connect directly using natural language
or keywords to answer customer requests.

Use the Amazon Q in Connect APIs to create an assistant and a knowledge base, for example, or manage content by
uploading custom files.

For more information, see [Use\
Amazon Q in Connect for generative AI powered agent assistance in real-time](../../../../services/connect/latest/adminguide/amazon-q-connect.md) in the _Amazon Connect_
_Administrator Guide_.

## Amazon Voice ID

###### Important

End of support notice: On May 20, 2026, AWS will end support for
Amazon Connect Voice ID. After May 20, 2026, you will no longer be able to access Voice ID on the Amazon Connect
console, access Voice ID features on the Amazon Connect admin website or Contact Control Panel, or access
Voice ID resources. For more information, visit
[Amazon Connect Voice ID end of support](../../../../services/connect/latest/adminguide/amazonconnect-voiceid-end-of-support.md).

- [Voice ID actions](api-operations-amazon-voice-id.md)

- [Voice ID data types](api-types-amazon-voice-id.md)

Amazon Connect Voice ID provides real-time caller authentication and fraud risk detection, which
make voice interactions in contact centers more secure and efficient.

For more information about the Voice ID feature, see [Use real-time caller authentication with\
Voice ID](../../../../services/connect/latest/adminguide/voice-id.md) in the _Amazon Connect Administrator Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Programming with the Amazon Connect APIs

All content copied from https://docs.aws.amazon.com/.
