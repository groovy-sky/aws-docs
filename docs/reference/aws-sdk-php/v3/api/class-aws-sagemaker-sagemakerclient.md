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

- [**2017-07-24**](api-sagemaker-2017-07-24.md)

  - [AddAssociation](api-sagemaker-2017-07-24-addassociation.md)
  - [AddTags](api-sagemaker-2017-07-24-addtags.md)
  - [AssociateTrialComponent](api-sagemaker-2017-07-24-associatetrialcomponent.md)
  - [AttachClusterNodeVolume](api-sagemaker-2017-07-24-attachclusternodevolume.md)
  - [BatchAddClusterNodes](api-sagemaker-2017-07-24-batchaddclusternodes.md)
  - [BatchDeleteClusterNodes](api-sagemaker-2017-07-24-batchdeleteclusternodes.md)
  - [BatchDescribeModelPackage](api-sagemaker-2017-07-24-batchdescribemodelpackage.md)
  - [BatchRebootClusterNodes](api-sagemaker-2017-07-24-batchrebootclusternodes.md)
  - [BatchReplaceClusterNodes](api-sagemaker-2017-07-24-batchreplaceclusternodes.md)
  - [CreateAction](api-sagemaker-2017-07-24-createaction.md)
  - [CreateAlgorithm](api-sagemaker-2017-07-24-createalgorithm.md)
  - [CreateApp](api-sagemaker-2017-07-24-createapp.md)
  - [CreateAppImageConfig](api-sagemaker-2017-07-24-createappimageconfig.md)
  - [CreateArtifact](api-sagemaker-2017-07-24-createartifact.md)
  - [CreateAutoMLJob](api-sagemaker-2017-07-24-createautomljob.md)
  - [CreateAutoMLJobV2](api-sagemaker-2017-07-24-createautomljobv2.md)
  - [CreateCluster](api-sagemaker-2017-07-24-createcluster.md)
  - [CreateClusterSchedulerConfig](api-sagemaker-2017-07-24-createclusterschedulerconfig.md)
  - [CreateCodeRepository](api-sagemaker-2017-07-24-createcoderepository.md)
  - [CreateCompilationJob](api-sagemaker-2017-07-24-createcompilationjob.md)
  - [CreateComputeQuota](api-sagemaker-2017-07-24-createcomputequota.md)
  - [CreateContext](api-sagemaker-2017-07-24-createcontext.md)
  - [CreateDataQualityJobDefinition](api-sagemaker-2017-07-24-createdataqualityjobdefinition.md)
  - [CreateDeviceFleet](api-sagemaker-2017-07-24-createdevicefleet.md)
  - [CreateDomain](api-sagemaker-2017-07-24-createdomain.md)
  - [CreateEdgeDeploymentPlan](api-sagemaker-2017-07-24-createedgedeploymentplan.md)
  - [CreateEdgeDeploymentStage](api-sagemaker-2017-07-24-createedgedeploymentstage.md)
  - [CreateEdgePackagingJob](api-sagemaker-2017-07-24-createedgepackagingjob.md)
  - [CreateEndpoint](api-sagemaker-2017-07-24-createendpoint.md)
  - [CreateEndpointConfig](api-sagemaker-2017-07-24-createendpointconfig.md)
  - [CreateExperiment](api-sagemaker-2017-07-24-createexperiment.md)
  - [CreateFeatureGroup](api-sagemaker-2017-07-24-createfeaturegroup.md)
  - [CreateFlowDefinition](api-sagemaker-2017-07-24-createflowdefinition.md)
  - [CreateHub](api-sagemaker-2017-07-24-createhub.md)
  - [CreateHubContentPresignedUrls](api-sagemaker-2017-07-24-createhubcontentpresignedurls.md)
  - [CreateHubContentReference](api-sagemaker-2017-07-24-createhubcontentreference.md)
  - [CreateHumanTaskUi](api-sagemaker-2017-07-24-createhumantaskui.md)
  - [CreateHyperParameterTuningJob](api-sagemaker-2017-07-24-createhyperparametertuningjob.md)
  - [CreateImage](api-sagemaker-2017-07-24-createimage.md)
  - [CreateImageVersion](api-sagemaker-2017-07-24-createimageversion.md)
  - [CreateInferenceComponent](api-sagemaker-2017-07-24-createinferencecomponent.md)
  - [CreateInferenceExperiment](api-sagemaker-2017-07-24-createinferenceexperiment.md)
  - [CreateInferenceRecommendationsJob](api-sagemaker-2017-07-24-createinferencerecommendationsjob.md)
  - [CreateLabelingJob](api-sagemaker-2017-07-24-createlabelingjob.md)
  - [CreateMlflowApp](api-sagemaker-2017-07-24-createmlflowapp.md)
  - [CreateMlflowTrackingServer](api-sagemaker-2017-07-24-createmlflowtrackingserver.md)
  - [CreateModel](api-sagemaker-2017-07-24-createmodel.md)
  - [CreateModelBiasJobDefinition](api-sagemaker-2017-07-24-createmodelbiasjobdefinition.md)
  - [CreateModelCard](api-sagemaker-2017-07-24-createmodelcard.md)
  - [CreateModelCardExportJob](api-sagemaker-2017-07-24-createmodelcardexportjob.md)
  - [CreateModelExplainabilityJobDefinition](api-sagemaker-2017-07-24-createmodelexplainabilityjobdefinition.md)
  - [CreateModelPackage](api-sagemaker-2017-07-24-createmodelpackage.md)
  - [CreateModelPackageGroup](api-sagemaker-2017-07-24-createmodelpackagegroup.md)
  - [CreateModelQualityJobDefinition](api-sagemaker-2017-07-24-createmodelqualityjobdefinition.md)
  - [CreateMonitoringSchedule](api-sagemaker-2017-07-24-createmonitoringschedule.md)
  - [CreateNotebookInstance](api-sagemaker-2017-07-24-createnotebookinstance.md)
  - [CreateNotebookInstanceLifecycleConfig](api-sagemaker-2017-07-24-createnotebookinstancelifecycleconfig.md)
  - [CreateOptimizationJob](api-sagemaker-2017-07-24-createoptimizationjob.md)
  - [CreatePartnerApp](api-sagemaker-2017-07-24-createpartnerapp.md)
  - [CreatePartnerAppPresignedUrl](api-sagemaker-2017-07-24-createpartnerapppresignedurl.md)
  - [CreatePipeline](api-sagemaker-2017-07-24-createpipeline.md)
  - [CreatePresignedDomainUrl](api-sagemaker-2017-07-24-createpresigneddomainurl.md)
  - [CreatePresignedMlflowAppUrl](api-sagemaker-2017-07-24-createpresignedmlflowappurl.md)
  - [CreatePresignedMlflowTrackingServerUrl](api-sagemaker-2017-07-24-createpresignedmlflowtrackingserverurl.md)
  - [CreatePresignedNotebookInstanceUrl](api-sagemaker-2017-07-24-createpresignednotebookinstanceurl.md)
  - [CreateProcessingJob](api-sagemaker-2017-07-24-createprocessingjob.md)
  - [CreateProject](api-sagemaker-2017-07-24-createproject.md)
  - [CreateSpace](api-sagemaker-2017-07-24-createspace.md)
  - [CreateStudioLifecycleConfig](api-sagemaker-2017-07-24-createstudiolifecycleconfig.md)
  - [CreateTrainingJob](api-sagemaker-2017-07-24-createtrainingjob.md)
  - [CreateTrainingPlan](api-sagemaker-2017-07-24-createtrainingplan.md)
  - [CreateTransformJob](api-sagemaker-2017-07-24-createtransformjob.md)
  - [CreateTrial](api-sagemaker-2017-07-24-createtrial.md)
  - [CreateTrialComponent](api-sagemaker-2017-07-24-createtrialcomponent.md)
  - [CreateUserProfile](api-sagemaker-2017-07-24-createuserprofile.md)
  - [CreateWorkforce](api-sagemaker-2017-07-24-createworkforce.md)
  - [CreateWorkteam](api-sagemaker-2017-07-24-createworkteam.md)
  - [DeleteAction](api-sagemaker-2017-07-24-deleteaction.md)
  - [DeleteAlgorithm](api-sagemaker-2017-07-24-deletealgorithm.md)
  - [DeleteApp](api-sagemaker-2017-07-24-deleteapp.md)
  - [DeleteAppImageConfig](api-sagemaker-2017-07-24-deleteappimageconfig.md)
  - [DeleteArtifact](api-sagemaker-2017-07-24-deleteartifact.md)
  - [DeleteAssociation](api-sagemaker-2017-07-24-deleteassociation.md)
  - [DeleteCluster](api-sagemaker-2017-07-24-deletecluster.md)
  - [DeleteClusterSchedulerConfig](api-sagemaker-2017-07-24-deleteclusterschedulerconfig.md)
  - [DeleteCodeRepository](api-sagemaker-2017-07-24-deletecoderepository.md)
  - [DeleteCompilationJob](api-sagemaker-2017-07-24-deletecompilationjob.md)
  - [DeleteComputeQuota](api-sagemaker-2017-07-24-deletecomputequota.md)
  - [DeleteContext](api-sagemaker-2017-07-24-deletecontext.md)
  - [DeleteDataQualityJobDefinition](api-sagemaker-2017-07-24-deletedataqualityjobdefinition.md)
  - [DeleteDeviceFleet](api-sagemaker-2017-07-24-deletedevicefleet.md)
  - [DeleteDomain](api-sagemaker-2017-07-24-deletedomain.md)
  - [DeleteEdgeDeploymentPlan](api-sagemaker-2017-07-24-deleteedgedeploymentplan.md)
  - [DeleteEdgeDeploymentStage](api-sagemaker-2017-07-24-deleteedgedeploymentstage.md)
  - [DeleteEndpoint](api-sagemaker-2017-07-24-deleteendpoint.md)
  - [DeleteEndpointConfig](api-sagemaker-2017-07-24-deleteendpointconfig.md)
  - [DeleteExperiment](api-sagemaker-2017-07-24-deleteexperiment.md)
  - [DeleteFeatureGroup](api-sagemaker-2017-07-24-deletefeaturegroup.md)
  - [DeleteFlowDefinition](api-sagemaker-2017-07-24-deleteflowdefinition.md)
  - [DeleteHub](api-sagemaker-2017-07-24-deletehub.md)
  - [DeleteHubContent](api-sagemaker-2017-07-24-deletehubcontent.md)
  - [DeleteHubContentReference](api-sagemaker-2017-07-24-deletehubcontentreference.md)
  - [DeleteHumanTaskUi](api-sagemaker-2017-07-24-deletehumantaskui.md)
  - [DeleteHyperParameterTuningJob](api-sagemaker-2017-07-24-deletehyperparametertuningjob.md)
  - [DeleteImage](api-sagemaker-2017-07-24-deleteimage.md)
  - [DeleteImageVersion](api-sagemaker-2017-07-24-deleteimageversion.md)
  - [DeleteInferenceComponent](api-sagemaker-2017-07-24-deleteinferencecomponent.md)
  - [DeleteInferenceExperiment](api-sagemaker-2017-07-24-deleteinferenceexperiment.md)
  - [DeleteMlflowApp](api-sagemaker-2017-07-24-deletemlflowapp.md)
  - [DeleteMlflowTrackingServer](api-sagemaker-2017-07-24-deletemlflowtrackingserver.md)
  - [DeleteModel](api-sagemaker-2017-07-24-deletemodel.md)
  - [DeleteModelBiasJobDefinition](api-sagemaker-2017-07-24-deletemodelbiasjobdefinition.md)
  - [DeleteModelCard](api-sagemaker-2017-07-24-deletemodelcard.md)
  - [DeleteModelExplainabilityJobDefinition](api-sagemaker-2017-07-24-deletemodelexplainabilityjobdefinition.md)
  - [DeleteModelPackage](api-sagemaker-2017-07-24-deletemodelpackage.md)
  - [DeleteModelPackageGroup](api-sagemaker-2017-07-24-deletemodelpackagegroup.md)
  - [DeleteModelPackageGroupPolicy](api-sagemaker-2017-07-24-deletemodelpackagegrouppolicy.md)
  - [DeleteModelQualityJobDefinition](api-sagemaker-2017-07-24-deletemodelqualityjobdefinition.md)
  - [DeleteMonitoringSchedule](api-sagemaker-2017-07-24-deletemonitoringschedule.md)
  - [DeleteNotebookInstance](api-sagemaker-2017-07-24-deletenotebookinstance.md)
  - [DeleteNotebookInstanceLifecycleConfig](api-sagemaker-2017-07-24-deletenotebookinstancelifecycleconfig.md)
  - [DeleteOptimizationJob](api-sagemaker-2017-07-24-deleteoptimizationjob.md)
  - [DeletePartnerApp](api-sagemaker-2017-07-24-deletepartnerapp.md)
  - [DeletePipeline](api-sagemaker-2017-07-24-deletepipeline.md)
  - [DeleteProcessingJob](api-sagemaker-2017-07-24-deleteprocessingjob.md)
  - [DeleteProject](api-sagemaker-2017-07-24-deleteproject.md)
  - [DeleteSpace](api-sagemaker-2017-07-24-deletespace.md)
  - [DeleteStudioLifecycleConfig](api-sagemaker-2017-07-24-deletestudiolifecycleconfig.md)
  - [DeleteTags](api-sagemaker-2017-07-24-deletetags.md)
  - [DeleteTrainingJob](api-sagemaker-2017-07-24-deletetrainingjob.md)
  - [DeleteTrial](api-sagemaker-2017-07-24-deletetrial.md)
  - [DeleteTrialComponent](api-sagemaker-2017-07-24-deletetrialcomponent.md)
  - [DeleteUserProfile](api-sagemaker-2017-07-24-deleteuserprofile.md)
  - [DeleteWorkforce](api-sagemaker-2017-07-24-deleteworkforce.md)
  - [DeleteWorkteam](api-sagemaker-2017-07-24-deleteworkteam.md)
  - [DeregisterDevices](api-sagemaker-2017-07-24-deregisterdevices.md)
  - [DescribeAction](api-sagemaker-2017-07-24-describeaction.md)
  - [DescribeAlgorithm](api-sagemaker-2017-07-24-describealgorithm.md)
  - [DescribeApp](api-sagemaker-2017-07-24-describeapp.md)
  - [DescribeAppImageConfig](api-sagemaker-2017-07-24-describeappimageconfig.md)
  - [DescribeArtifact](api-sagemaker-2017-07-24-describeartifact.md)
  - [DescribeAutoMLJob](api-sagemaker-2017-07-24-describeautomljob.md)
  - [DescribeAutoMLJobV2](api-sagemaker-2017-07-24-describeautomljobv2.md)
  - [DescribeCluster](api-sagemaker-2017-07-24-describecluster.md)
  - [DescribeClusterEvent](api-sagemaker-2017-07-24-describeclusterevent.md)
  - [DescribeClusterNode](api-sagemaker-2017-07-24-describeclusternode.md)
  - [DescribeClusterSchedulerConfig](api-sagemaker-2017-07-24-describeclusterschedulerconfig.md)
  - [DescribeCodeRepository](api-sagemaker-2017-07-24-describecoderepository.md)
  - [DescribeCompilationJob](api-sagemaker-2017-07-24-describecompilationjob.md)
  - [DescribeComputeQuota](api-sagemaker-2017-07-24-describecomputequota.md)
  - [DescribeContext](api-sagemaker-2017-07-24-describecontext.md)
  - [DescribeDataQualityJobDefinition](api-sagemaker-2017-07-24-describedataqualityjobdefinition.md)
  - [DescribeDevice](api-sagemaker-2017-07-24-describedevice.md)
  - [DescribeDeviceFleet](api-sagemaker-2017-07-24-describedevicefleet.md)
  - [DescribeDomain](api-sagemaker-2017-07-24-describedomain.md)
  - [DescribeEdgeDeploymentPlan](api-sagemaker-2017-07-24-describeedgedeploymentplan.md)
  - [DescribeEdgePackagingJob](api-sagemaker-2017-07-24-describeedgepackagingjob.md)
  - [DescribeEndpoint](api-sagemaker-2017-07-24-describeendpoint.md)
  - [DescribeEndpointConfig](api-sagemaker-2017-07-24-describeendpointconfig.md)
  - [DescribeExperiment](api-sagemaker-2017-07-24-describeexperiment.md)
  - [DescribeFeatureGroup](api-sagemaker-2017-07-24-describefeaturegroup.md)
  - [DescribeFeatureMetadata](api-sagemaker-2017-07-24-describefeaturemetadata.md)
  - [DescribeFlowDefinition](api-sagemaker-2017-07-24-describeflowdefinition.md)
  - [DescribeHub](api-sagemaker-2017-07-24-describehub.md)
  - [DescribeHubContent](api-sagemaker-2017-07-24-describehubcontent.md)
  - [DescribeHumanTaskUi](api-sagemaker-2017-07-24-describehumantaskui.md)
  - [DescribeHyperParameterTuningJob](api-sagemaker-2017-07-24-describehyperparametertuningjob.md)
  - [DescribeImage](api-sagemaker-2017-07-24-describeimage.md)
  - [DescribeImageVersion](api-sagemaker-2017-07-24-describeimageversion.md)
  - [DescribeInferenceComponent](api-sagemaker-2017-07-24-describeinferencecomponent.md)
  - [DescribeInferenceExperiment](api-sagemaker-2017-07-24-describeinferenceexperiment.md)
  - [DescribeInferenceRecommendationsJob](api-sagemaker-2017-07-24-describeinferencerecommendationsjob.md)
  - [DescribeLabelingJob](api-sagemaker-2017-07-24-describelabelingjob.md)
  - [DescribeLineageGroup](api-sagemaker-2017-07-24-describelineagegroup.md)
  - [DescribeMlflowApp](api-sagemaker-2017-07-24-describemlflowapp.md)
  - [DescribeMlflowTrackingServer](api-sagemaker-2017-07-24-describemlflowtrackingserver.md)
  - [DescribeModel](api-sagemaker-2017-07-24-describemodel.md)
  - [DescribeModelBiasJobDefinition](api-sagemaker-2017-07-24-describemodelbiasjobdefinition.md)
  - [DescribeModelCard](api-sagemaker-2017-07-24-describemodelcard.md)
  - [DescribeModelCardExportJob](api-sagemaker-2017-07-24-describemodelcardexportjob.md)
  - [DescribeModelExplainabilityJobDefinition](api-sagemaker-2017-07-24-describemodelexplainabilityjobdefinition.md)
  - [DescribeModelPackage](api-sagemaker-2017-07-24-describemodelpackage.md)
  - [DescribeModelPackageGroup](api-sagemaker-2017-07-24-describemodelpackagegroup.md)
  - [DescribeModelQualityJobDefinition](api-sagemaker-2017-07-24-describemodelqualityjobdefinition.md)
  - [DescribeMonitoringSchedule](api-sagemaker-2017-07-24-describemonitoringschedule.md)
  - [DescribeNotebookInstance](api-sagemaker-2017-07-24-describenotebookinstance.md)
  - [DescribeNotebookInstanceLifecycleConfig](api-sagemaker-2017-07-24-describenotebookinstancelifecycleconfig.md)
  - [DescribeOptimizationJob](api-sagemaker-2017-07-24-describeoptimizationjob.md)
  - [DescribePartnerApp](api-sagemaker-2017-07-24-describepartnerapp.md)
  - [DescribePipeline](api-sagemaker-2017-07-24-describepipeline.md)
  - [DescribePipelineDefinitionForExecution](api-sagemaker-2017-07-24-describepipelinedefinitionforexecution.md)
  - [DescribePipelineExecution](api-sagemaker-2017-07-24-describepipelineexecution.md)
  - [DescribeProcessingJob](api-sagemaker-2017-07-24-describeprocessingjob.md)
  - [DescribeProject](api-sagemaker-2017-07-24-describeproject.md)
  - [DescribeReservedCapacity](api-sagemaker-2017-07-24-describereservedcapacity.md)
  - [DescribeSpace](api-sagemaker-2017-07-24-describespace.md)
  - [DescribeStudioLifecycleConfig](api-sagemaker-2017-07-24-describestudiolifecycleconfig.md)
  - [DescribeSubscribedWorkteam](api-sagemaker-2017-07-24-describesubscribedworkteam.md)
  - [DescribeTrainingJob](api-sagemaker-2017-07-24-describetrainingjob.md)
  - [DescribeTrainingPlan](api-sagemaker-2017-07-24-describetrainingplan.md)
  - [DescribeTrainingPlanExtensionHistory](api-sagemaker-2017-07-24-describetrainingplanextensionhistory.md)
  - [DescribeTransformJob](api-sagemaker-2017-07-24-describetransformjob.md)
  - [DescribeTrial](api-sagemaker-2017-07-24-describetrial.md)
  - [DescribeTrialComponent](api-sagemaker-2017-07-24-describetrialcomponent.md)
  - [DescribeUserProfile](api-sagemaker-2017-07-24-describeuserprofile.md)
  - [DescribeWorkforce](api-sagemaker-2017-07-24-describeworkforce.md)
  - [DescribeWorkteam](api-sagemaker-2017-07-24-describeworkteam.md)
  - [DetachClusterNodeVolume](api-sagemaker-2017-07-24-detachclusternodevolume.md)
  - [DisableSagemakerServicecatalogPortfolio](api-sagemaker-2017-07-24-disablesagemakerservicecatalogportfolio.md)
  - [DisassociateTrialComponent](api-sagemaker-2017-07-24-disassociatetrialcomponent.md)
  - [EnableSagemakerServicecatalogPortfolio](api-sagemaker-2017-07-24-enablesagemakerservicecatalogportfolio.md)
  - [ExtendTrainingPlan](api-sagemaker-2017-07-24-extendtrainingplan.md)
  - [GetDeviceFleetReport](api-sagemaker-2017-07-24-getdevicefleetreport.md)
  - [GetLineageGroupPolicy](api-sagemaker-2017-07-24-getlineagegrouppolicy.md)
  - [GetModelPackageGroupPolicy](api-sagemaker-2017-07-24-getmodelpackagegrouppolicy.md)
  - [GetSagemakerServicecatalogPortfolioStatus](api-sagemaker-2017-07-24-getsagemakerservicecatalogportfoliostatus.md)
  - [GetScalingConfigurationRecommendation](api-sagemaker-2017-07-24-getscalingconfigurationrecommendation.md)
  - [GetSearchSuggestions](api-sagemaker-2017-07-24-getsearchsuggestions.md)
  - [ImportHubContent](api-sagemaker-2017-07-24-importhubcontent.md)
  - [ListActions](api-sagemaker-2017-07-24-listactions.md)
  - [ListAlgorithms](api-sagemaker-2017-07-24-listalgorithms.md)
  - [ListAliases](api-sagemaker-2017-07-24-listaliases.md)
  - [ListAppImageConfigs](api-sagemaker-2017-07-24-listappimageconfigs.md)
  - [ListApps](api-sagemaker-2017-07-24-listapps.md)
  - [ListArtifacts](api-sagemaker-2017-07-24-listartifacts.md)
  - [ListAssociations](api-sagemaker-2017-07-24-listassociations.md)
  - [ListAutoMLJobs](api-sagemaker-2017-07-24-listautomljobs.md)
  - [ListCandidatesForAutoMLJob](api-sagemaker-2017-07-24-listcandidatesforautomljob.md)
  - [ListClusterEvents](api-sagemaker-2017-07-24-listclusterevents.md)
  - [ListClusterNodes](api-sagemaker-2017-07-24-listclusternodes.md)
  - [ListClusterSchedulerConfigs](api-sagemaker-2017-07-24-listclusterschedulerconfigs.md)
  - [ListClusters](api-sagemaker-2017-07-24-listclusters.md)
  - [ListCodeRepositories](api-sagemaker-2017-07-24-listcoderepositories.md)
  - [ListCompilationJobs](api-sagemaker-2017-07-24-listcompilationjobs.md)
  - [ListComputeQuotas](api-sagemaker-2017-07-24-listcomputequotas.md)
  - [ListContexts](api-sagemaker-2017-07-24-listcontexts.md)
  - [ListDataQualityJobDefinitions](api-sagemaker-2017-07-24-listdataqualityjobdefinitions.md)
  - [ListDeviceFleets](api-sagemaker-2017-07-24-listdevicefleets.md)
  - [ListDevices](api-sagemaker-2017-07-24-listdevices.md)
  - [ListDomains](api-sagemaker-2017-07-24-listdomains.md)
  - [ListEdgeDeploymentPlans](api-sagemaker-2017-07-24-listedgedeploymentplans.md)
  - [ListEdgePackagingJobs](api-sagemaker-2017-07-24-listedgepackagingjobs.md)
  - [ListEndpointConfigs](api-sagemaker-2017-07-24-listendpointconfigs.md)
  - [ListEndpoints](api-sagemaker-2017-07-24-listendpoints.md)
  - [ListExperiments](api-sagemaker-2017-07-24-listexperiments.md)
  - [ListFeatureGroups](api-sagemaker-2017-07-24-listfeaturegroups.md)
  - [ListFlowDefinitions](api-sagemaker-2017-07-24-listflowdefinitions.md)
  - [ListHubContentVersions](api-sagemaker-2017-07-24-listhubcontentversions.md)
  - [ListHubContents](api-sagemaker-2017-07-24-listhubcontents.md)
  - [ListHubs](api-sagemaker-2017-07-24-listhubs.md)
  - [ListHumanTaskUis](api-sagemaker-2017-07-24-listhumantaskuis.md)
  - [ListHyperParameterTuningJobs](api-sagemaker-2017-07-24-listhyperparametertuningjobs.md)
  - [ListImageVersions](api-sagemaker-2017-07-24-listimageversions.md)
  - [ListImages](api-sagemaker-2017-07-24-listimages.md)
  - [ListInferenceComponents](api-sagemaker-2017-07-24-listinferencecomponents.md)
  - [ListInferenceExperiments](api-sagemaker-2017-07-24-listinferenceexperiments.md)
  - [ListInferenceRecommendationsJobSteps](api-sagemaker-2017-07-24-listinferencerecommendationsjobsteps.md)
  - [ListInferenceRecommendationsJobs](api-sagemaker-2017-07-24-listinferencerecommendationsjobs.md)
  - [ListLabelingJobs](api-sagemaker-2017-07-24-listlabelingjobs.md)
  - [ListLabelingJobsForWorkteam](api-sagemaker-2017-07-24-listlabelingjobsforworkteam.md)
  - [ListLineageGroups](api-sagemaker-2017-07-24-listlineagegroups.md)
  - [ListMlflowApps](api-sagemaker-2017-07-24-listmlflowapps.md)
  - [ListMlflowTrackingServers](api-sagemaker-2017-07-24-listmlflowtrackingservers.md)
  - [ListModelBiasJobDefinitions](api-sagemaker-2017-07-24-listmodelbiasjobdefinitions.md)
  - [ListModelCardExportJobs](api-sagemaker-2017-07-24-listmodelcardexportjobs.md)
  - [ListModelCardVersions](api-sagemaker-2017-07-24-listmodelcardversions.md)
  - [ListModelCards](api-sagemaker-2017-07-24-listmodelcards.md)
  - [ListModelExplainabilityJobDefinitions](api-sagemaker-2017-07-24-listmodelexplainabilityjobdefinitions.md)
  - [ListModelMetadata](api-sagemaker-2017-07-24-listmodelmetadata.md)
  - [ListModelPackageGroups](api-sagemaker-2017-07-24-listmodelpackagegroups.md)
  - [ListModelPackages](api-sagemaker-2017-07-24-listmodelpackages.md)
  - [ListModelQualityJobDefinitions](api-sagemaker-2017-07-24-listmodelqualityjobdefinitions.md)
  - [ListModels](api-sagemaker-2017-07-24-listmodels.md)
  - [ListMonitoringAlertHistory](api-sagemaker-2017-07-24-listmonitoringalerthistory.md)
  - [ListMonitoringAlerts](api-sagemaker-2017-07-24-listmonitoringalerts.md)
  - [ListMonitoringExecutions](api-sagemaker-2017-07-24-listmonitoringexecutions.md)
  - [ListMonitoringSchedules](api-sagemaker-2017-07-24-listmonitoringschedules.md)
  - [ListNotebookInstanceLifecycleConfigs](api-sagemaker-2017-07-24-listnotebookinstancelifecycleconfigs.md)
  - [ListNotebookInstances](api-sagemaker-2017-07-24-listnotebookinstances.md)
  - [ListOptimizationJobs](api-sagemaker-2017-07-24-listoptimizationjobs.md)
  - [ListPartnerApps](api-sagemaker-2017-07-24-listpartnerapps.md)
  - [ListPipelineExecutionSteps](api-sagemaker-2017-07-24-listpipelineexecutionsteps.md)
  - [ListPipelineExecutions](api-sagemaker-2017-07-24-listpipelineexecutions.md)
  - [ListPipelineParametersForExecution](api-sagemaker-2017-07-24-listpipelineparametersforexecution.md)
  - [ListPipelineVersions](api-sagemaker-2017-07-24-listpipelineversions.md)
  - [ListPipelines](api-sagemaker-2017-07-24-listpipelines.md)
  - [ListProcessingJobs](api-sagemaker-2017-07-24-listprocessingjobs.md)
  - [ListProjects](api-sagemaker-2017-07-24-listprojects.md)
  - [ListResourceCatalogs](api-sagemaker-2017-07-24-listresourcecatalogs.md)
  - [ListSpaces](api-sagemaker-2017-07-24-listspaces.md)
  - [ListStageDevices](api-sagemaker-2017-07-24-liststagedevices.md)
  - [ListStudioLifecycleConfigs](api-sagemaker-2017-07-24-liststudiolifecycleconfigs.md)
  - [ListSubscribedWorkteams](api-sagemaker-2017-07-24-listsubscribedworkteams.md)
  - [ListTags](api-sagemaker-2017-07-24-listtags.md)
  - [ListTrainingJobs](api-sagemaker-2017-07-24-listtrainingjobs.md)
  - [ListTrainingJobsForHyperParameterTuningJob](api-sagemaker-2017-07-24-listtrainingjobsforhyperparametertuningjob.md)
  - [ListTrainingPlans](api-sagemaker-2017-07-24-listtrainingplans.md)
  - [ListTransformJobs](api-sagemaker-2017-07-24-listtransformjobs.md)
  - [ListTrialComponents](api-sagemaker-2017-07-24-listtrialcomponents.md)
  - [ListTrials](api-sagemaker-2017-07-24-listtrials.md)
  - [ListUltraServersByReservedCapacity](api-sagemaker-2017-07-24-listultraserversbyreservedcapacity.md)
  - [ListUserProfiles](api-sagemaker-2017-07-24-listuserprofiles.md)
  - [ListWorkforces](api-sagemaker-2017-07-24-listworkforces.md)
  - [ListWorkteams](api-sagemaker-2017-07-24-listworkteams.md)
  - [PutModelPackageGroupPolicy](api-sagemaker-2017-07-24-putmodelpackagegrouppolicy.md)
  - [QueryLineage](api-sagemaker-2017-07-24-querylineage.md)
  - [RegisterDevices](api-sagemaker-2017-07-24-registerdevices.md)
  - [RenderUiTemplate](api-sagemaker-2017-07-24-renderuitemplate.md)
  - [RetryPipelineExecution](api-sagemaker-2017-07-24-retrypipelineexecution.md)
  - [Search](api-sagemaker-2017-07-24-search.md)
  - [SearchTrainingPlanOfferings](api-sagemaker-2017-07-24-searchtrainingplanofferings.md)
  - [SendPipelineExecutionStepFailure](api-sagemaker-2017-07-24-sendpipelineexecutionstepfailure.md)
  - [SendPipelineExecutionStepSuccess](api-sagemaker-2017-07-24-sendpipelineexecutionstepsuccess.md)
  - [StartEdgeDeploymentStage](api-sagemaker-2017-07-24-startedgedeploymentstage.md)
  - [StartInferenceExperiment](api-sagemaker-2017-07-24-startinferenceexperiment.md)
  - [StartMlflowTrackingServer](api-sagemaker-2017-07-24-startmlflowtrackingserver.md)
  - [StartMonitoringSchedule](api-sagemaker-2017-07-24-startmonitoringschedule.md)
  - [StartNotebookInstance](api-sagemaker-2017-07-24-startnotebookinstance.md)
  - [StartPipelineExecution](api-sagemaker-2017-07-24-startpipelineexecution.md)
  - [StartSession](api-sagemaker-2017-07-24-startsession.md)
  - [StopAutoMLJob](api-sagemaker-2017-07-24-stopautomljob.md)
  - [StopCompilationJob](api-sagemaker-2017-07-24-stopcompilationjob.md)
  - [StopEdgeDeploymentStage](api-sagemaker-2017-07-24-stopedgedeploymentstage.md)
  - [StopEdgePackagingJob](api-sagemaker-2017-07-24-stopedgepackagingjob.md)
  - [StopHyperParameterTuningJob](api-sagemaker-2017-07-24-stophyperparametertuningjob.md)
  - [StopInferenceExperiment](api-sagemaker-2017-07-24-stopinferenceexperiment.md)
  - [StopInferenceRecommendationsJob](api-sagemaker-2017-07-24-stopinferencerecommendationsjob.md)
  - [StopLabelingJob](api-sagemaker-2017-07-24-stoplabelingjob.md)
  - [StopMlflowTrackingServer](api-sagemaker-2017-07-24-stopmlflowtrackingserver.md)
  - [StopMonitoringSchedule](api-sagemaker-2017-07-24-stopmonitoringschedule.md)
  - [StopNotebookInstance](api-sagemaker-2017-07-24-stopnotebookinstance.md)
  - [StopOptimizationJob](api-sagemaker-2017-07-24-stopoptimizationjob.md)
  - [StopPipelineExecution](api-sagemaker-2017-07-24-stoppipelineexecution.md)
  - [StopProcessingJob](api-sagemaker-2017-07-24-stopprocessingjob.md)
  - [StopTrainingJob](api-sagemaker-2017-07-24-stoptrainingjob.md)
  - [StopTransformJob](api-sagemaker-2017-07-24-stoptransformjob.md)
  - [UpdateAction](api-sagemaker-2017-07-24-updateaction.md)
  - [UpdateAppImageConfig](api-sagemaker-2017-07-24-updateappimageconfig.md)
  - [UpdateArtifact](api-sagemaker-2017-07-24-updateartifact.md)
  - [UpdateCluster](api-sagemaker-2017-07-24-updatecluster.md)
  - [UpdateClusterSchedulerConfig](api-sagemaker-2017-07-24-updateclusterschedulerconfig.md)
  - [UpdateClusterSoftware](api-sagemaker-2017-07-24-updateclustersoftware.md)
  - [UpdateCodeRepository](api-sagemaker-2017-07-24-updatecoderepository.md)
  - [UpdateComputeQuota](api-sagemaker-2017-07-24-updatecomputequota.md)
  - [UpdateContext](api-sagemaker-2017-07-24-updatecontext.md)
  - [UpdateDeviceFleet](api-sagemaker-2017-07-24-updatedevicefleet.md)
  - [UpdateDevices](api-sagemaker-2017-07-24-updatedevices.md)
  - [UpdateDomain](api-sagemaker-2017-07-24-updatedomain.md)
  - [UpdateEndpoint](api-sagemaker-2017-07-24-updateendpoint.md)
  - [UpdateEndpointWeightsAndCapacities](api-sagemaker-2017-07-24-updateendpointweightsandcapacities.md)
  - [UpdateExperiment](api-sagemaker-2017-07-24-updateexperiment.md)
  - [UpdateFeatureGroup](api-sagemaker-2017-07-24-updatefeaturegroup.md)
  - [UpdateFeatureMetadata](api-sagemaker-2017-07-24-updatefeaturemetadata.md)
  - [UpdateHub](api-sagemaker-2017-07-24-updatehub.md)
  - [UpdateHubContent](api-sagemaker-2017-07-24-updatehubcontent.md)
  - [UpdateHubContentReference](api-sagemaker-2017-07-24-updatehubcontentreference.md)
  - [UpdateImage](api-sagemaker-2017-07-24-updateimage.md)
  - [UpdateImageVersion](api-sagemaker-2017-07-24-updateimageversion.md)
  - [UpdateInferenceComponent](api-sagemaker-2017-07-24-updateinferencecomponent.md)
  - [UpdateInferenceComponentRuntimeConfig](api-sagemaker-2017-07-24-updateinferencecomponentruntimeconfig.md)
  - [UpdateInferenceExperiment](api-sagemaker-2017-07-24-updateinferenceexperiment.md)
  - [UpdateMlflowApp](api-sagemaker-2017-07-24-updatemlflowapp.md)
  - [UpdateMlflowTrackingServer](api-sagemaker-2017-07-24-updatemlflowtrackingserver.md)
  - [UpdateModelCard](api-sagemaker-2017-07-24-updatemodelcard.md)
  - [UpdateModelPackage](api-sagemaker-2017-07-24-updatemodelpackage.md)
  - [UpdateMonitoringAlert](api-sagemaker-2017-07-24-updatemonitoringalert.md)
  - [UpdateMonitoringSchedule](api-sagemaker-2017-07-24-updatemonitoringschedule.md)
  - [UpdateNotebookInstance](api-sagemaker-2017-07-24-updatenotebookinstance.md)
  - [UpdateNotebookInstanceLifecycleConfig](api-sagemaker-2017-07-24-updatenotebookinstancelifecycleconfig.md)
  - [UpdatePartnerApp](api-sagemaker-2017-07-24-updatepartnerapp.md)
  - [UpdatePipeline](api-sagemaker-2017-07-24-updatepipeline.md)
  - [UpdatePipelineExecution](api-sagemaker-2017-07-24-updatepipelineexecution.md)
  - [UpdatePipelineVersion](api-sagemaker-2017-07-24-updatepipelineversion.md)
  - [UpdateProject](api-sagemaker-2017-07-24-updateproject.md)
  - [UpdateSpace](api-sagemaker-2017-07-24-updatespace.md)
  - [UpdateTrainingJob](api-sagemaker-2017-07-24-updatetrainingjob.md)
  - [UpdateTrial](api-sagemaker-2017-07-24-updatetrial.md)
  - [UpdateTrialComponent](api-sagemaker-2017-07-24-updatetrialcomponent.md)
  - [UpdateUserProfile](api-sagemaker-2017-07-24-updateuserprofile.md)
  - [UpdateWorkforce](api-sagemaker-2017-07-24-updateworkforce.md)
  - [UpdateWorkteam](api-sagemaker-2017-07-24-updateworkteam.md)

### Table of Contents  [header link](class-aws-sagemaker-sagemakerclient-toc.md)

#### Methods  [header link](class-aws-sagemaker-sagemakerclient-toc-methods.md)

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

### Methods  [header link](class-aws-sagemaker-sagemakerclient-methods.md)

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
  - [Methods](class-aws-sagemaker-sagemakerclient-toc-methods.md)
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

[Back To Top](class-aws-sagemaker-sagemakerclient-top.md)
