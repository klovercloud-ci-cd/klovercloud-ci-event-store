package enums

// STEP_TYPE steps type
type STEP_TYPE string

const (
	// BUILD build step
	BUILD = STEP_TYPE("BUILD")
	// DEPLOY deploy step
	DEPLOY = STEP_TYPE("DEPLOY")
	// INTERMEDIARY step that runs custom jobs
	INTERMEDIARY = STEP_TYPE("INTERMEDIARY")
	// JENKIN step that runs jenkin jobs
	JENKIN = STEP_TYPE("JENKINS_JOB")
)
const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

const (
	// DEFAULT_PAGE_LIMIT default page limit for rest api
	DEFAULT_PAGE_LIMIT = 10
	// DEFAULT_PAGE default page for rest api
	DEFAULT_PAGE = 0
)

// PIPELINE_PURGING pipeline process purging policy
type PIPELINE_PURGING string

const (
	// PIPELINE_PURGING_ENABLE pipeline process purging is enabled
	PIPELINE_PURGING_ENABLE = PIPELINE_PURGING("ENABLE")
	// PIPELINE_PURGING_DISABLE pipeline process purging is disabled
	PIPELINE_PURGING_DISABLE = PIPELINE_PURGING("DISABLE")
)

// TRIGGER pipeline trigger options
type TRIGGER string

const (
	// AUTO pipeline trigger options is auto
	AUTO = TRIGGER("AUTO")
	// MANUAL pipeline trigger options is MANUAL
	MANUAL = TRIGGER("MANUAL")
)

// PARAMS pipeline parameters
type PARAMS string

const (
	// REPOSITORY_TYPE repository type key for pipeline step param
	REPOSITORY_TYPE = PARAMS("repository_type")
	// REVISION resource revision key for  pipeline step param
	REVISION = PARAMS("revision")
	// SERVICE_ACCOUNT k8s service account key that contains registry and repository secret as pipeline step param
	SERVICE_ACCOUNT = PARAMS("service_account")
	// IMAGES key for container images as pipeline step param
	IMAGES = PARAMS("images")
	// ARGS_FROM_CONFIGMAPS key for build and other arguments via configmaps as pipeline step param
	ARGS_FROM_CONFIGMAPS = PARAMS("args_from_configmaps")
	// ARGS_FROM_SECRETS key for build and other arguments via secrets as pipeline step param
	ARGS_FROM_SECRETS = PARAMS("args_from_secrets")
	// ENVS_FROM_CONFIGMAPS key for env via configmaps as pipeline step param
	ENVS_FROM_CONFIGMAPS = PARAMS("envs_from_configmaps")
	// ENVS_FROM_SECRETS key for env via secrets as pipeline step param
	ENVS_FROM_SECRETS = PARAMS("envs_from_secrets")
	// ARGS key for build and other arguments as pipeline step param
	ARGS = PARAMS("arg")
	// ENVS key for env as pipeline step param
	ENVS = PARAMS("envs")
	// AGENT key for agent name as pipeline step param
	AGENT = PARAMS("agent")
	// RESOURCE_NAME key for k8s resource name as pipeline step param
	RESOURCE_NAME = PARAMS("name")
	// RESOURCE_NAMESPACE key for k8s resource namespace as pipeline step param
	RESOURCE_NAMESPACE = PARAMS("namespace")
	// IMAGE_URL key for image url as pipeline step param
	IMAGE_URL = PARAMS("url")
)

// PROCESS_STATUS pipeline steps status
type PROCESS_STATUS string

const (
	// QUEUED pipeline steps status queued
	QUEUED=PROCESS_STATUS("queued")
	// NON_INITIALIZED pipeline steps status non_initialized
	NON_INITIALIZED = PROCESS_STATUS("non_initialized")
	// ACTIVE pipeline steps status active
	ACTIVE = PROCESS_STATUS("active")
	// COMPLETED pipeline steps status completed
	COMPLETED = PROCESS_STATUS("completed")
	// FAILED pipeline steps status failed
	FAILED = PROCESS_STATUS("failed")
	// PAUSED pipeline steps status paused
	PAUSED = PROCESS_STATUS("paused")
)

// PROCESS_EVENT_STATUS process event status
type PROCESS_EVENT_STATUS string

const (
	// PROCESS_EVENT_FAILED process event has been FAILED
	PROCESS_EVENT_FAILED = PROCESS_EVENT_STATUS("FAILED")
	// PROCESS_EVENT_PROCESSING process event has been PROCESSING
	PROCESS_EVENT_PROCESSING = PROCESS_EVENT_STATUS("PROCESSING")
	// PROCESS_EVENT_TERMINATING process event has been TERMINATING
	PROCESS_EVENT_TERMINATING = PROCESS_EVENT_STATUS("TERMINATING")
	// PROCESS_EVENT_INITIALIZING process event has been INITIALIZING
	PROCESS_EVENT_INITIALIZING = PROCESS_EVENT_STATUS("INITIALIZING")
	// PROCESS_EVENT_SUCCESSFUL process event has been SUCCESSFUL
	PROCESS_EVENT_SUCCESSFUL = PROCESS_EVENT_STATUS("SUCCESSFUL")
	// PROCESS_EVENT_ERROR process event has been ERROR
	PROCESS_EVENT_ERROR = PROCESS_EVENT_STATUS("ERROR")
)

// PIPELINE_RESOURCE_TYPE pipeline resource types
type PIPELINE_RESOURCE_TYPE string

const (
	// GIT git as resource
	GIT = PIPELINE_RESOURCE_TYPE("git")
	// IMAGE docker image as resource
	IMAGE = PIPELINE_RESOURCE_TYPE("image")
	// DEPLOYMENT k8s deployment as resource
	DEPLOYMENT = PIPELINE_RESOURCE_TYPE("deployment")
	// STATEFULSET k8s statefulset as resource
	STATEFULSET = PIPELINE_RESOURCE_TYPE("statefulset")
	// DAEMONSET k8s daemonset as resource
	DAEMONSET = PIPELINE_RESOURCE_TYPE("daemonset")
	// POD k8s pod as resource
	POD = PIPELINE_RESOURCE_TYPE("pod")
	// REPLICASET k8s replicaset as resource
	REPLICASET = PIPELINE_RESOURCE_TYPE("replicaset")
)

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

// FOOTMARK process footmark (step breakdown)
type FOOTMARK string

const (
	// INIT_BUILD_JOB FOOTMARK name
	INIT_BUILD_JOB = FOOTMARK("init_build_job")
	// POST_BUILD_JOB FOOTMARK name
	POST_BUILD_JOB = FOOTMARK("post_build_job")
	// BUILD_AND_PUSH FOOTMARK name
	BUILD_AND_PUSH = FOOTMARK("build_and_push")
	// GIT_CLONE FOOTMARK name
	GIT_CLONE = FOOTMARK("git_clone")

	// INIT_INTERMEDIARY_JOB FOOTMARK name
	INIT_INTERMEDIARY_JOB = FOOTMARK("init_intermediary_job")
	// POST_INTERMEDIARY_JOB FOOTMARK name
	POST_INTERMEDIARY_JOB = FOOTMARK("post_intermediary_job")

	// INIT_JENKINS_JOB FOOTMARK name
	INIT_JENKINS_JOB = FOOTMARK("init_jenkins_job")
	// POST_JENKINS_JOB FOOTMARK name
	POST_JENKINS_JOB = FOOTMARK("post_jenkins_job")

	POST_DEPLOY_JOB = FOOTMARK("post_deploy_job")
)
