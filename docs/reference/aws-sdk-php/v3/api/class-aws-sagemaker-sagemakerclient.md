Menu

- [Aws](namespace-aws.md)
- [SageMaker](namespace-aws-sagemaker.md)

## SageMakerClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon SageMaker Service** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2017-07-24**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html)

  - [AddAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#addassociation)
  - [AddTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#addtags)
  - [AssociateTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#associatetrialcomponent)
  - [AttachClusterNodeVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#attachclusternodevolume)
  - [BatchAddClusterNodes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#batchaddclusternodes)
  - [BatchDeleteClusterNodes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#batchdeleteclusternodes)
  - [BatchDescribeModelPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#batchdescribemodelpackage)
  - [BatchRebootClusterNodes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#batchrebootclusternodes)
  - [BatchReplaceClusterNodes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#batchreplaceclusternodes)
  - [CreateAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createaction)
  - [CreateAlgorithm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createalgorithm)
  - [CreateApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createapp)
  - [CreateAppImageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createappimageconfig)
  - [CreateArtifact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createartifact)
  - [CreateAutoMLJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createautomljob)
  - [CreateAutoMLJobV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createautomljobv2)
  - [CreateCluster](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createcluster)
  - [CreateClusterSchedulerConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createclusterschedulerconfig)
  - [CreateCodeRepository](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createcoderepository)
  - [CreateCompilationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createcompilationjob)
  - [CreateComputeQuota](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createcomputequota)
  - [CreateContext](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createcontext)
  - [CreateDataQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createdataqualityjobdefinition)
  - [CreateDeviceFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createdevicefleet)
  - [CreateDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createdomain)
  - [CreateEdgeDeploymentPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createedgedeploymentplan)
  - [CreateEdgeDeploymentStage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createedgedeploymentstage)
  - [CreateEdgePackagingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createedgepackagingjob)
  - [CreateEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createendpoint)
  - [CreateEndpointConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createendpointconfig)
  - [CreateExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createexperiment)
  - [CreateFeatureGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createfeaturegroup)
  - [CreateFlowDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createflowdefinition)
  - [CreateHub](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createhub)
  - [CreateHubContentPresignedUrls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createhubcontentpresignedurls)
  - [CreateHubContentReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createhubcontentreference)
  - [CreateHumanTaskUi](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createhumantaskui)
  - [CreateHyperParameterTuningJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createhyperparametertuningjob)
  - [CreateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createimage)
  - [CreateImageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createimageversion)
  - [CreateInferenceComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createinferencecomponent)
  - [CreateInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createinferenceexperiment)
  - [CreateInferenceRecommendationsJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createinferencerecommendationsjob)
  - [CreateLabelingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createlabelingjob)
  - [CreateMlflowApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmlflowapp)
  - [CreateMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmlflowtrackingserver)
  - [CreateModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodel)
  - [CreateModelBiasJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelbiasjobdefinition)
  - [CreateModelCard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelcard)
  - [CreateModelCardExportJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelcardexportjob)
  - [CreateModelExplainabilityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelexplainabilityjobdefinition)
  - [CreateModelPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelpackage)
  - [CreateModelPackageGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelpackagegroup)
  - [CreateModelQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmodelqualityjobdefinition)
  - [CreateMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createmonitoringschedule)
  - [CreateNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createnotebookinstance)
  - [CreateNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createnotebookinstancelifecycleconfig)
  - [CreateOptimizationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createoptimizationjob)
  - [CreatePartnerApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpartnerapp)
  - [CreatePartnerAppPresignedUrl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpartnerapppresignedurl)
  - [CreatePipeline](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpipeline)
  - [CreatePresignedDomainUrl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpresigneddomainurl)
  - [CreatePresignedMlflowAppUrl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpresignedmlflowappurl)
  - [CreatePresignedMlflowTrackingServerUrl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpresignedmlflowtrackingserverurl)
  - [CreatePresignedNotebookInstanceUrl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createpresignednotebookinstanceurl)
  - [CreateProcessingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createprocessingjob)
  - [CreateProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createproject)
  - [CreateSpace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createspace)
  - [CreateStudioLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createstudiolifecycleconfig)
  - [CreateTrainingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createtrainingjob)
  - [CreateTrainingPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createtrainingplan)
  - [CreateTransformJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createtransformjob)
  - [CreateTrial](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createtrial)
  - [CreateTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createtrialcomponent)
  - [CreateUserProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createuserprofile)
  - [CreateWorkforce](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createworkforce)
  - [CreateWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#createworkteam)
  - [DeleteAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteaction)
  - [DeleteAlgorithm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletealgorithm)
  - [DeleteApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteapp)
  - [DeleteAppImageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteappimageconfig)
  - [DeleteArtifact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteartifact)
  - [DeleteAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteassociation)
  - [DeleteCluster](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletecluster)
  - [DeleteClusterSchedulerConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteclusterschedulerconfig)
  - [DeleteCodeRepository](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletecoderepository)
  - [DeleteCompilationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletecompilationjob)
  - [DeleteComputeQuota](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletecomputequota)
  - [DeleteContext](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletecontext)
  - [DeleteDataQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletedataqualityjobdefinition)
  - [DeleteDeviceFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletedevicefleet)
  - [DeleteDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletedomain)
  - [DeleteEdgeDeploymentPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteedgedeploymentplan)
  - [DeleteEdgeDeploymentStage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteedgedeploymentstage)
  - [DeleteEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteendpoint)
  - [DeleteEndpointConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteendpointconfig)
  - [DeleteExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteexperiment)
  - [DeleteFeatureGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletefeaturegroup)
  - [DeleteFlowDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteflowdefinition)
  - [DeleteHub](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletehub)
  - [DeleteHubContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletehubcontent)
  - [DeleteHubContentReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletehubcontentreference)
  - [DeleteHumanTaskUi](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletehumantaskui)
  - [DeleteHyperParameterTuningJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletehyperparametertuningjob)
  - [DeleteImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteimage)
  - [DeleteImageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteimageversion)
  - [DeleteInferenceComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteinferencecomponent)
  - [DeleteInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteinferenceexperiment)
  - [DeleteMlflowApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemlflowapp)
  - [DeleteMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemlflowtrackingserver)
  - [DeleteModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodel)
  - [DeleteModelBiasJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelbiasjobdefinition)
  - [DeleteModelCard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelcard)
  - [DeleteModelExplainabilityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelexplainabilityjobdefinition)
  - [DeleteModelPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelpackage)
  - [DeleteModelPackageGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelpackagegroup)
  - [DeleteModelPackageGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelpackagegrouppolicy)
  - [DeleteModelQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemodelqualityjobdefinition)
  - [DeleteMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletemonitoringschedule)
  - [DeleteNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletenotebookinstance)
  - [DeleteNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletenotebookinstancelifecycleconfig)
  - [DeleteOptimizationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteoptimizationjob)
  - [DeletePartnerApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletepartnerapp)
  - [DeletePipeline](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletepipeline)
  - [DeleteProcessingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteprocessingjob)
  - [DeleteProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteproject)
  - [DeleteSpace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletespace)
  - [DeleteStudioLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletestudiolifecycleconfig)
  - [DeleteTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletetags)
  - [DeleteTrainingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletetrainingjob)
  - [DeleteTrial](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletetrial)
  - [DeleteTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deletetrialcomponent)
  - [DeleteUserProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteuserprofile)
  - [DeleteWorkforce](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteworkforce)
  - [DeleteWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deleteworkteam)
  - [DeregisterDevices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#deregisterdevices)
  - [DescribeAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeaction)
  - [DescribeAlgorithm](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describealgorithm)
  - [DescribeApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeapp)
  - [DescribeAppImageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeappimageconfig)
  - [DescribeArtifact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeartifact)
  - [DescribeAutoMLJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeautomljob)
  - [DescribeAutoMLJobV2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeautomljobv2)
  - [DescribeCluster](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describecluster)
  - [DescribeClusterEvent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeclusterevent)
  - [DescribeClusterNode](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeclusternode)
  - [DescribeClusterSchedulerConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeclusterschedulerconfig)
  - [DescribeCodeRepository](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describecoderepository)
  - [DescribeCompilationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describecompilationjob)
  - [DescribeComputeQuota](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describecomputequota)
  - [DescribeContext](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describecontext)
  - [DescribeDataQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describedataqualityjobdefinition)
  - [DescribeDevice](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describedevice)
  - [DescribeDeviceFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describedevicefleet)
  - [DescribeDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describedomain)
  - [DescribeEdgeDeploymentPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeedgedeploymentplan)
  - [DescribeEdgePackagingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeedgepackagingjob)
  - [DescribeEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeendpoint)
  - [DescribeEndpointConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeendpointconfig)
  - [DescribeExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeexperiment)
  - [DescribeFeatureGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describefeaturegroup)
  - [DescribeFeatureMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describefeaturemetadata)
  - [DescribeFlowDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeflowdefinition)
  - [DescribeHub](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describehub)
  - [DescribeHubContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describehubcontent)
  - [DescribeHumanTaskUi](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describehumantaskui)
  - [DescribeHyperParameterTuningJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describehyperparametertuningjob)
  - [DescribeImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeimage)
  - [DescribeImageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeimageversion)
  - [DescribeInferenceComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeinferencecomponent)
  - [DescribeInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeinferenceexperiment)
  - [DescribeInferenceRecommendationsJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeinferencerecommendationsjob)
  - [DescribeLabelingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describelabelingjob)
  - [DescribeLineageGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describelineagegroup)
  - [DescribeMlflowApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemlflowapp)
  - [DescribeMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemlflowtrackingserver)
  - [DescribeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodel)
  - [DescribeModelBiasJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelbiasjobdefinition)
  - [DescribeModelCard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelcard)
  - [DescribeModelCardExportJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelcardexportjob)
  - [DescribeModelExplainabilityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelexplainabilityjobdefinition)
  - [DescribeModelPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelpackage)
  - [DescribeModelPackageGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelpackagegroup)
  - [DescribeModelQualityJobDefinition](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemodelqualityjobdefinition)
  - [DescribeMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describemonitoringschedule)
  - [DescribeNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describenotebookinstance)
  - [DescribeNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describenotebookinstancelifecycleconfig)
  - [DescribeOptimizationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeoptimizationjob)
  - [DescribePartnerApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describepartnerapp)
  - [DescribePipeline](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describepipeline)
  - [DescribePipelineDefinitionForExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describepipelinedefinitionforexecution)
  - [DescribePipelineExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describepipelineexecution)
  - [DescribeProcessingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeprocessingjob)
  - [DescribeProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeproject)
  - [DescribeReservedCapacity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describereservedcapacity)
  - [DescribeSpace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describespace)
  - [DescribeStudioLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describestudiolifecycleconfig)
  - [DescribeSubscribedWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describesubscribedworkteam)
  - [DescribeTrainingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetrainingjob)
  - [DescribeTrainingPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetrainingplan)
  - [DescribeTrainingPlanExtensionHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetrainingplanextensionhistory)
  - [DescribeTransformJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetransformjob)
  - [DescribeTrial](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetrial)
  - [DescribeTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describetrialcomponent)
  - [DescribeUserProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeuserprofile)
  - [DescribeWorkforce](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeworkforce)
  - [DescribeWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#describeworkteam)
  - [DetachClusterNodeVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#detachclusternodevolume)
  - [DisableSagemakerServicecatalogPortfolio](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#disablesagemakerservicecatalogportfolio)
  - [DisassociateTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#disassociatetrialcomponent)
  - [EnableSagemakerServicecatalogPortfolio](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#enablesagemakerservicecatalogportfolio)
  - [ExtendTrainingPlan](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#extendtrainingplan)
  - [GetDeviceFleetReport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getdevicefleetreport)
  - [GetLineageGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getlineagegrouppolicy)
  - [GetModelPackageGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getmodelpackagegrouppolicy)
  - [GetSagemakerServicecatalogPortfolioStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getsagemakerservicecatalogportfoliostatus)
  - [GetScalingConfigurationRecommendation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getscalingconfigurationrecommendation)
  - [GetSearchSuggestions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#getsearchsuggestions)
  - [ImportHubContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#importhubcontent)
  - [ListActions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listactions)
  - [ListAlgorithms](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listalgorithms)
  - [ListAliases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listaliases)
  - [ListAppImageConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listappimageconfigs)
  - [ListApps](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listapps)
  - [ListArtifacts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listartifacts)
  - [ListAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listassociations)
  - [ListAutoMLJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listautomljobs)
  - [ListCandidatesForAutoMLJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listcandidatesforautomljob)
  - [ListClusterEvents](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listclusterevents)
  - [ListClusterNodes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listclusternodes)
  - [ListClusterSchedulerConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listclusterschedulerconfigs)
  - [ListClusters](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listclusters)
  - [ListCodeRepositories](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listcoderepositories)
  - [ListCompilationJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listcompilationjobs)
  - [ListComputeQuotas](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listcomputequotas)
  - [ListContexts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listcontexts)
  - [ListDataQualityJobDefinitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listdataqualityjobdefinitions)
  - [ListDeviceFleets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listdevicefleets)
  - [ListDevices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listdevices)
  - [ListDomains](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listdomains)
  - [ListEdgeDeploymentPlans](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listedgedeploymentplans)
  - [ListEdgePackagingJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listedgepackagingjobs)
  - [ListEndpointConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listendpointconfigs)
  - [ListEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listendpoints)
  - [ListExperiments](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listexperiments)
  - [ListFeatureGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listfeaturegroups)
  - [ListFlowDefinitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listflowdefinitions)
  - [ListHubContentVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listhubcontentversions)
  - [ListHubContents](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listhubcontents)
  - [ListHubs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listhubs)
  - [ListHumanTaskUis](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listhumantaskuis)
  - [ListHyperParameterTuningJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listhyperparametertuningjobs)
  - [ListImageVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listimageversions)
  - [ListImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listimages)
  - [ListInferenceComponents](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listinferencecomponents)
  - [ListInferenceExperiments](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listinferenceexperiments)
  - [ListInferenceRecommendationsJobSteps](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listinferencerecommendationsjobsteps)
  - [ListInferenceRecommendationsJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listinferencerecommendationsjobs)
  - [ListLabelingJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listlabelingjobs)
  - [ListLabelingJobsForWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listlabelingjobsforworkteam)
  - [ListLineageGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listlineagegroups)
  - [ListMlflowApps](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmlflowapps)
  - [ListMlflowTrackingServers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmlflowtrackingservers)
  - [ListModelBiasJobDefinitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelbiasjobdefinitions)
  - [ListModelCardExportJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelcardexportjobs)
  - [ListModelCardVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelcardversions)
  - [ListModelCards](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelcards)
  - [ListModelExplainabilityJobDefinitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelexplainabilityjobdefinitions)
  - [ListModelMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelmetadata)
  - [ListModelPackageGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelpackagegroups)
  - [ListModelPackages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelpackages)
  - [ListModelQualityJobDefinitions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodelqualityjobdefinitions)
  - [ListModels](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmodels)
  - [ListMonitoringAlertHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmonitoringalerthistory)
  - [ListMonitoringAlerts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmonitoringalerts)
  - [ListMonitoringExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmonitoringexecutions)
  - [ListMonitoringSchedules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listmonitoringschedules)
  - [ListNotebookInstanceLifecycleConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listnotebookinstancelifecycleconfigs)
  - [ListNotebookInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listnotebookinstances)
  - [ListOptimizationJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listoptimizationjobs)
  - [ListPartnerApps](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpartnerapps)
  - [ListPipelineExecutionSteps](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpipelineexecutionsteps)
  - [ListPipelineExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpipelineexecutions)
  - [ListPipelineParametersForExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpipelineparametersforexecution)
  - [ListPipelineVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpipelineversions)
  - [ListPipelines](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listpipelines)
  - [ListProcessingJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listprocessingjobs)
  - [ListProjects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listprojects)
  - [ListResourceCatalogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listresourcecatalogs)
  - [ListSpaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listspaces)
  - [ListStageDevices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#liststagedevices)
  - [ListStudioLifecycleConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#liststudiolifecycleconfigs)
  - [ListSubscribedWorkteams](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listsubscribedworkteams)
  - [ListTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtags)
  - [ListTrainingJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtrainingjobs)
  - [ListTrainingJobsForHyperParameterTuningJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtrainingjobsforhyperparametertuningjob)
  - [ListTrainingPlans](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtrainingplans)
  - [ListTransformJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtransformjobs)
  - [ListTrialComponents](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtrialcomponents)
  - [ListTrials](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listtrials)
  - [ListUltraServersByReservedCapacity](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listultraserversbyreservedcapacity)
  - [ListUserProfiles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listuserprofiles)
  - [ListWorkforces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listworkforces)
  - [ListWorkteams](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#listworkteams)
  - [PutModelPackageGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#putmodelpackagegrouppolicy)
  - [QueryLineage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#querylineage)
  - [RegisterDevices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#registerdevices)
  - [RenderUiTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#renderuitemplate)
  - [RetryPipelineExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#retrypipelineexecution)
  - [Search](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#search)
  - [SearchTrainingPlanOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#searchtrainingplanofferings)
  - [SendPipelineExecutionStepFailure](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#sendpipelineexecutionstepfailure)
  - [SendPipelineExecutionStepSuccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#sendpipelineexecutionstepsuccess)
  - [StartEdgeDeploymentStage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startedgedeploymentstage)
  - [StartInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startinferenceexperiment)
  - [StartMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startmlflowtrackingserver)
  - [StartMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startmonitoringschedule)
  - [StartNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startnotebookinstance)
  - [StartPipelineExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startpipelineexecution)
  - [StartSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#startsession)
  - [StopAutoMLJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopautomljob)
  - [StopCompilationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopcompilationjob)
  - [StopEdgeDeploymentStage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopedgedeploymentstage)
  - [StopEdgePackagingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopedgepackagingjob)
  - [StopHyperParameterTuningJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stophyperparametertuningjob)
  - [StopInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopinferenceexperiment)
  - [StopInferenceRecommendationsJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopinferencerecommendationsjob)
  - [StopLabelingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stoplabelingjob)
  - [StopMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopmlflowtrackingserver)
  - [StopMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopmonitoringschedule)
  - [StopNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopnotebookinstance)
  - [StopOptimizationJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopoptimizationjob)
  - [StopPipelineExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stoppipelineexecution)
  - [StopProcessingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stopprocessingjob)
  - [StopTrainingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stoptrainingjob)
  - [StopTransformJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#stoptransformjob)
  - [UpdateAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateaction)
  - [UpdateAppImageConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateappimageconfig)
  - [UpdateArtifact](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateartifact)
  - [UpdateCluster](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatecluster)
  - [UpdateClusterSchedulerConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateclusterschedulerconfig)
  - [UpdateClusterSoftware](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateclustersoftware)
  - [UpdateCodeRepository](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatecoderepository)
  - [UpdateComputeQuota](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatecomputequota)
  - [UpdateContext](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatecontext)
  - [UpdateDeviceFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatedevicefleet)
  - [UpdateDevices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatedevices)
  - [UpdateDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatedomain)
  - [UpdateEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateendpoint)
  - [UpdateEndpointWeightsAndCapacities](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateendpointweightsandcapacities)
  - [UpdateExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateexperiment)
  - [UpdateFeatureGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatefeaturegroup)
  - [UpdateFeatureMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatefeaturemetadata)
  - [UpdateHub](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatehub)
  - [UpdateHubContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatehubcontent)
  - [UpdateHubContentReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatehubcontentreference)
  - [UpdateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateimage)
  - [UpdateImageVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateimageversion)
  - [UpdateInferenceComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateinferencecomponent)
  - [UpdateInferenceComponentRuntimeConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateinferencecomponentruntimeconfig)
  - [UpdateInferenceExperiment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateinferenceexperiment)
  - [UpdateMlflowApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemlflowapp)
  - [UpdateMlflowTrackingServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemlflowtrackingserver)
  - [UpdateModelCard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemodelcard)
  - [UpdateModelPackage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemodelpackage)
  - [UpdateMonitoringAlert](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemonitoringalert)
  - [UpdateMonitoringSchedule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatemonitoringschedule)
  - [UpdateNotebookInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatenotebookinstance)
  - [UpdateNotebookInstanceLifecycleConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatenotebookinstancelifecycleconfig)
  - [UpdatePartnerApp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatepartnerapp)
  - [UpdatePipeline](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatepipeline)
  - [UpdatePipelineExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatepipelineexecution)
  - [UpdatePipelineVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatepipelineversion)
  - [UpdateProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateproject)
  - [UpdateSpace](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatespace)
  - [UpdateTrainingJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatetrainingjob)
  - [UpdateTrial](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatetrial)
  - [UpdateTrialComponent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updatetrialcomponent)
  - [UpdateUserProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateuserprofile)
  - [UpdateWorkforce](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateworkforce)
  - [UpdateWorkteam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-sagemaker-2017-07-24.html#updateworkteam)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.SageMaker.SageMakerClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.SageMaker.SageMakerClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-awsclient.md#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.SageMaker.SageMakerClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-awsclient.md\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

##### Tags  [header link](class-aws-awsclient.md\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing or
the service is not supported.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.SageMaker.SageMakerClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-awsclient.md#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.SageMaker.SageMakerClient.html#top)
