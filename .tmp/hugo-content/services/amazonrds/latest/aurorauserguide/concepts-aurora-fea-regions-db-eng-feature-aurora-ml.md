# Supported Regions and DB engines for Aurora machine learning

By using Amazon Aurora machine learning, you can integrate your Aurora DB cluster with
one of the following AWS machine learning services, depending on your needs. They each
support specific machine learning use cases.

Amazon Bedrock is a fully managed service that makes leading foundation models
from AI companies available through an API, along with developer tooling to help build
and scale generative AI applications.

Amazon Comprehend is a _natural language processing_ (NLP)
service that's used to extract insights from documents. By using Aurora machine learning with Amazon Comprehend,
you can determine the sentiment of text in your database tables.

SageMaker AI is a full-featured _machine learning_ service. Data scientists
use Amazon SageMaker AI to build, train, and test machine learning models for a variety of
inference tasks, such as fraud detection. By using Aurora machine learning with SageMaker AI, database developers
can invoke the SageMaker AI functionality in SQL code.

Not all AWS Regions support all machine learning services. Only certain
AWS Regions support Aurora machine learning and thus provide access to these services from an Aurora DB
cluster. The integration process for Aurora machine learning also differs by database engine. For more
information, see [Using Amazon Aurora machine learning](aurora-ml.md).

###### Topics

- [Aurora machine learning with Aurora MySQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.Aurora_ML.amy)

- [Aurora machine learning with Aurora PostgreSQL](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.Aurora_ML.apg)

## Aurora machine learning with Aurora MySQL

Amazon Bedrock is supported only on Aurora MySQL version 3.06 and higher. For
information on Region availability for Amazon Bedrock, see [Model support by\
AWS Region](../../../bedrock/latest/userguide/models-regions.md) in the _Amazon Bedrock User Guide_.

Aurora machine learning with Amazon Comprehend and Amazon SageMaker AI is supported for
Aurora MySQL in the AWS Regions listed in the table. In addition to having your
version of Aurora MySQL available, the AWS Region must also support the service that
you want to use. For a list of AWS Regions where Amazon SageMaker AI is available, see [Amazon SageMaker AI endpoints\
and quotas](../../../../general/latest/gr/sagemaker.md) in the _Amazon Web Services General Reference_. For a list of AWS Regions where Amazon Comprehend is available, see [Amazon Comprehend endpoints\
and quotas](../../../../general/latest/gr/comprehend.md) in the
_Amazon Web Services General Reference_.

RegionAurora MySQL version 3Aurora MySQL version 2US East (N. Virginia)Version 3.01.0 and higherVersion 2.07 and higherUS East (Ohio)Version 3.01.0 and higherVersion 2.07 and higherUS West (N. California)Version 3.01.0 and higherVersion 2.07 and higherUS West (Oregon)Version 3.01.0 and higherVersion 2.07 and higherAfrica (Cape Town)Not availableNot availableAsia Pacific (Hong Kong)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Hyderabad)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Jakarta)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Malaysia)Version 3.04.0 and higherNot availableAsia Pacific (Melbourne)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Mumbai)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (New Zealand)Not availableNot availableAsia Pacific (Osaka)Version 3.01.0 and higherVersion 2.07.3 and higherAsia Pacific (Seoul)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Singapore)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Sydney)Version 3.01.0 and higherVersion 2.07 and higherAsia Pacific (Taipei)Not availableNot availableAsia Pacific (Thailand)Not availableNot availableAsia Pacific (Tokyo)Version 3.01.0 and higherVersion 2.07 and higherCanada (Central)Version 3.01.0 and higherVersion 2.07 and higherCanada West (Calgary)Version 3.01.0 and higherVersion 2.07 and higherChina (Beijing)Version 3.01.0 and higherVersion 2.07 and higherChina (Ningxia)Version 3.01.0 and higherVersion 2.07 and higherEurope (Frankfurt)Version 3.01.0 and higherVersion 2.07 and higherEurope (Ireland)Version 3.01.0 and higherVersion 2.07 and higherEurope (London)Version 3.01.0 and higherVersion 2.07 and higherEurope (Milan)Not availableNot availableEurope (Paris)Version 3.01.0 and higherVersion 2.07 and higherEurope (Spain)Version 3.01.0 and higherVersion 2.07 and higherEurope (Stockholm)Version 3.01.0 and higherVersion 2.07 and higherEurope (Zurich)Version 3.01.0 and higherVersion 2.07 and higherIsrael (Tel Aviv)Version 3.01.0 and higherVersion 2.07 and higherMexico (Central)Not availableNot availableMiddle East (Bahrain)Version 3.01.0 and higherVersion 2.07 and higherMiddle East (UAE)Version 3.01.0 and higherVersion 2.07 and higherSouth America (São Paulo)Version 3.01.0 and higherVersion 2.07 and higherAWS GovCloud (US-East)Version 3.01.0 and higherVersion 2.07 and higherAWS GovCloud (US-West)Version 3.01.0 and higherVersion 2.07 and higher

## Aurora machine learning with Aurora PostgreSQL

For information on version support for Amazon Bedrock on Aurora PostgreSQL, see
[Using Aurora PostgreSQL as a Knowledge Base for Amazon Bedrock](aurorapostgresql-vectordb.md).

Aurora machine learning with Amazon Comprehend and Amazon SageMaker AI is supported for
Aurora PostgreSQL in the AWS Regions listed in the table. In addition to having your
version of Aurora PostgreSQL available, the AWS Region must also support the service
that you want to use. For a list of AWS Regions where Amazon SageMaker AI is available, see
[Amazon SageMaker AI\
endpoints and quotas](../../../../general/latest/gr/sagemaker.md) in the _Amazon Web Services General Reference_.
For a list of AWS Regions where Amazon Comprehend is available,
see [Amazon Comprehend\
endpoints and quotas](../../../../general/latest/gr/comprehend.md) in the
_Amazon Web Services General Reference_.

The following Regions and engine versions are available for Aurora machine learning with
Aurora PostgreSQL.

RegionAurora PostgreSQL 17Aurora PostgreSQL 16Aurora PostgreSQL 15Aurora PostgreSQL 14Aurora PostgreSQL 13Aurora PostgreSQL 12Aurora PostgreSQL 11US East (N. Virginia)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherUS East (Ohio)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherUS West (N. California)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherUS West (Oregon)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAfrica (Cape Town)Not availableNot availableNot availableNot availableNot availableNot availableNot availableAsia Pacific (Hong Kong)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Hyderabad)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Jakarta)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Malaysia)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Melbourne)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Mumbai)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (New Zealand)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Osaka)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Seoul)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Singapore)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Sydney)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Taipei)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Thailand)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAsia Pacific (Tokyo)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherCanada (Central)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherCanada West (Calgary)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherChina (Beijing)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherChina (Ningxia)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Frankfurt)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Ireland)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (London)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Milan)Not availableNot availableNot availableNot availableNot availableNot availableNot availableEurope (Paris)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Spain)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Stockholm)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherEurope (Zurich)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherIsrael (Tel Aviv)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherMexico (Central)Not availableNot availableNot availableNot availableNot availableNot availableNot availableMiddle East (Bahrain)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherMiddle East (UAE)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherSouth America (São Paulo)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAWS GovCloud (US-East)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higherAWS GovCloud (US-West)Version 17.4 and higherVersion 16.1 and higherVersion 15.2 and higherVersion 14.3 and higherVersion 13.3 and higherVersion 12.4 and higherVersion 11.9, 11.12 and higher

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Kerberos authentication

Performance Insights

All content copied from https://docs.aws.amazon.com/.
