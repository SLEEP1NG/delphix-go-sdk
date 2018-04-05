package delphix

// OracleVirtualSource - A virtual Oracle source.
// extends OracleSource
type OracleVirtualSource struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Online Redo Log size in MB.
	// create = optional
	// update = readonly
	// minimum = 4
	RedoLogSizeInMB int `json:"redoLogSizeInMB,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// create = optional
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Number of Online Redo Log Groups.
	// default = 3
	// create = optional
	// update = readonly
	// minimum = 2
	RedoLogGroups int `json:"redoLogGroups,omitempty"`
	// A list of object references to Oracle Node Listeners selected for this
	// Virtual Database. Delphix picks one default listener from the target
	// environment if this list is empty at virtual database provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 256
	// create = required
	MountBase string `json:"mountBase,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the configuration for the source.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	Config string `json:"config,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	// format = objectReference
	ConfigTemplate string `json:"configTemplate,omitempty"`
}

// SourceStartParameters - The parameters to use as input to start a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceStartParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEVirtualSource - A virtual SAP ASE source.
// extends ASESource
type ASEVirtualSource struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Database file mapping rules.
	// create = optional
	// maxLength = 4096
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The base mount point for the NFS mounts.
	// maxLength = 80
	// create = optional
	// update = optional
	MountBase string `json:"mountBase,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// SAP ASE database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// User-specified operation hooks for this source.
	// update = optional
	// create = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntime `json:"runtime,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
}

// ConnectorConnectivity - Mechanism to test Connector connectivity of arbitrary hosts.
// extends TypedObject
type ConnectorConnectivity struct {
	// Host to use as a proxy for credential validation.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	Proxy string `json:"proxy,omitempty"`
	// Connector port on remote server.
	// default = 9100
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// User name.
	// required = true
	Username string `json:"username,omitempty"`
	// User credentials.
	// required = true
	Credentials *PasswordCredential `json:"credentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Target host name or IP address.
	// format = host
	// required = true
	Address string `json:"address,omitempty"`
}

// MySQLSourceConnectionInfo - Contains information that can be used to connect to a MySQL source.
// extends SourceConnectionInfo
type MySQLSourceConnectionInfo struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The port on which the MySQL server for the data directory is
	// listening.
	Port int `json:"port,omitempty"`
	// The username of the database user.
	User string `json:"user,omitempty"`
	// The data directory for the MySQL server.
	DataDirectory string `json:"dataDirectory,omitempty"`
	// The JDBC string used to connect to the MySQL server instance.
	JdbcString string `json:"jdbcString,omitempty"`
}

// SupportAccessState - The state of the access to the support shell.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SupportAccessState struct {
	// If ENABLED_WITH_TOKEN, the token that must be supplied to login.
	// update = optional
	Token string `json:"token,omitempty"`
	// How the support shell can be accessed.
	// update = required
	// enum = [DISABLED ENABLED_NO_TOKEN ENABLED_WITH_TOKEN]
	AccessType string `json:"accessType,omitempty"`
	// If ENABLED_WITH_TOKEN, the time that the token will be valid.
	// format = date
	// update = optional
	StartTime string `json:"startTime,omitempty"`
	// If ENABLED_WITH_TOKEN, time that the token will no longer be valid.
	// format = date
	// update = optional
	EndTime string `json:"endTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDBContainerRuntime - Runtime properties of an Oracle database container.
// extends DBContainerRuntime
type OracleDBContainerRuntime struct {
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// Indicates whether or not the given container is cross-platform
	// eligible or not.
	// default = false
	CrossPlatformEligible *bool `json:"crossPlatformEligible,omitempty"`
	// Indicates whether or not the given container has a cross-platform user
	// script uploaded.
	// default = false
	CrossPlatformScriptUploaded *bool `json:"crossPlatformScriptUploaded,omitempty"`
	// Indicates whether or not a LiveSource can be added to the given
	// container.
	// default = false
	LiveSourceEligible *bool `json:"liveSourceEligible,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
}

// NetworkInterfaceUtilDatapoint - An analytics datapoint generated by the NETWORK_INTERFACE_UTIL
// statistic type.
// extends Datapoint
type NetworkInterfaceUtilDatapoint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Bytes received on the interface.
	InBytes int `json:"inBytes,omitempty"`
	// Packets received on the interface.
	InPackets int `json:"inPackets,omitempty"`
	// Bytes transmitted on the interface.
	OutBytes int `json:"outBytes,omitempty"`
	// Packets transmitted on the interface.
	OutPackets int `json:"outPackets,omitempty"`
}

// ChangeCurrentPasswordPolicyParameters - The parameters to use as input when changing the currently active
// password policy.
// extends TypedObject
type ChangeCurrentPasswordPolicyParameters struct {
	// Password policy reference.
	// format = objectReference
	// referenceTo = /delphix-password-policy.json
	// required = true
	CurrentPasswordPolicy string `json:"currentPasswordPolicy,omitempty"`
	// Type of user.
	// required = true
	// enum = [SYSTEM DOMAIN]
	UserType string `json:"userType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEAttachData - Represents the SAP ASE specific parameters of an attach request.
// extends AttachData
type ASEAttachData struct {
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-ase-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The SAP ASE instance on the staging environment that we want to use
	// for validated sync.
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 80
	// create = optional
	MountBase string `json:"mountBase,omitempty"`
	// The user name for the source DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// Information about the host OS user on the staging environment to use
	// for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// The credential for the source DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with the
	// source database.
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// A user-provided shell script or executable to run after restoring from
	// a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// Source database backup location.
	// maxLength = 1024
	// required = true
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
	// The credential for the source DB user.
	// create = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// Information about the host OS user on the source to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// A user-provided shell script or executable to run prior to restoring
	// from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPreScript string `json:"stagingPreScript,omitempty"`
}

// SNMPConfig - SNMP configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SNMPConfig struct {
	// The network which is authorized to query this SNMP server, in CIDR
	// notation. Toallow any client, then leave unset or set to 0.0.0.0/0. To
	// block all clients, set to 127.0.0.1/8.
	// format = cidrAddress
	// update = optional
	// default = 0.0.0.0/0
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty"`
	// The physical location of this Delphix Engine (OID 1.3.6.1.2.1.1.6 -
	// sysLocation).
	// update = optional
	Location string `json:"location,omitempty"`
	// SNMP trap severity. SNMP managers are only notified of events at or
	// above this level.
	// update = optional
	// enum = [CRITICAL WARNING INFORMATIONAL]
	// default = WARNING
	Severity string `json:"severity,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the SNMP service is enabled.
	// update = optional
	// default = true
	Enabled *bool `json:"enabled,omitempty"`
	// The community string that clients must provide when querying this
	// server.
	// update = optional
	// default = public
	Community string `json:"community,omitempty"`
}

// RegistrationInfo - The information required to register the Delphix Engine.
// extends TypedObject
// cliVisibility = [SYSTEM]
type RegistrationInfo struct {
	// The UUID for this Delphix Engine.
	Uuid string `json:"uuid,omitempty"`
	// The registration code for this Delphix Engine.
	Code string `json:"code,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The registration portal hostname.
	RegistrationPortalHostname string `json:"registrationPortalHostname,omitempty"`
}

// OrFilter - A container filter that combines other filters together using OR
// logic.
// extends AlertFilter
type OrFilter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Filters which are combined together using OR logic.
	// create = required
	// update = optional
	// minItems = 2
	SubFilters []AlertFilter `json:"subFilters,omitempty"`
}

// CredentialUpdateParameters - Parameters to update a Delphix user's password.
// extends TypedObject
type CredentialUpdateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The old credentials.
	// create = optional
	OldCredential *PasswordCredential `json:"oldCredential,omitempty"`
	// The new credentials.
	// create = required
	NewCredential *PasswordCredential `json:"newCredential,omitempty"`
}

// ToolkitDiscoveryDefinition - Defines the discovery schemas and workflow scripts for a toolkit.
// extends TypedObject
type ToolkitDiscoveryDefinition struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script that discovers repositories on a target environment.
	// The script must return a list of repositories matching the
	// repositorySchema.
	// required = true
	RepositoryDiscovery string `json:"repositoryDiscovery,omitempty"`
	// A workflow script that discovers source configs on a target
	// environment. The script must return a list of source configs matching
	// the sourceConfigSchema.
	// required = true
	SourceConfigDiscovery string `json:"sourceConfigDiscovery,omitempty"`
	// The field of the sourceConfigSchema to display to the end user for
	// naming this source config.
	// required = true
	SourceConfigNameField string `json:"sourceConfigNameField,omitempty"`
	// A user defined schema to represent the repository.
	// required = true
	RepositorySchema *SchemaDraftV4 `json:"repositorySchema,omitempty"`
	// True if this toolkit supports manual discovery of source configs.
	// required = false
	ManualSourceConfigDiscovery *bool `json:"manualSourceConfigDiscovery,omitempty"`
	// A list of fields in the repositorySchema that collectively identify
	// each discovered repository.
	// required = true
	RepositoryIdentityFields []string `json:"repositoryIdentityFields,omitempty"`
	// A list of fields in the sourceConfigSchema that collectively identify
	// each discovered source config.
	// required = true
	SourceConfigIdentityFields []string `json:"sourceConfigIdentityFields,omitempty"`
	// The field of the repositorySchema to display to the end user for
	// naming this repository.
	// required = true
	RepositoryNameField string `json:"repositoryNameField,omitempty"`
	// A user defined schema to represent the source config.
	// required = true
	SourceConfigSchema *SchemaDraftV4 `json:"sourceConfigSchema,omitempty"`
}

// JDBCConnectivity - Mechanism to test JDBC connectivity of arbitrary databases.
// extends TypedObject
type JDBCConnectivity struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// JDBC connection URL.
	// format = oracleJDBCConnectionString
	// required = true
	Url string `json:"url,omitempty"`
	// Database username.
	// required = true
	User string `json:"user,omitempty"`
	// Database password.
	// required = true
	// format = password
	Password string `json:"password,omitempty"`
}

// MySQLNewMySQLDumpSyncParameters - The parameters to use as input to sync requests for MySQL databases
// using a new MySQL dump taken by Delphix.
// extends MySQLSyncParameters
type MySQLNewMySQLDumpSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleInstance - Representation of an Oracle instance configuration.
// extends TypedObject
type OracleInstance struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The number of the instance.
	// create = required
	// update = optional
	// minimum = 1
	InstanceNumber float64 `json:"instanceNumber,omitempty"`
	// The name of the instance.
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	// maxLength = 15
	InstanceName string `json:"instanceName,omitempty"`
}

// StorageTestResult - The test results of one storage test.
// extends TypedObject
type StorageTestResult struct {
	// Standard deviation of latency in milliseconds.
	StddevLatency float64 `json:"stddevLatency,omitempty"`
	// Average latency in milliseconds.
	AverageLatency float64 `json:"averageLatency,omitempty"`
	// 95th percentile latency in milliseconds.
	Latency95thPercentile float64 `json:"latency95thPercentile,omitempty"`
	// Grade assigned to the test for load scaling.
	LoadScalingGrade string `json:"loadScalingGrade,omitempty"`
	// Load scaling.
	LoadScaling float64 `json:"loadScaling,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Throughput.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Grade assigned to the test for latency.
	LatencyGrade string `json:"latencyGrade,omitempty"`
	// Block size used for the test.
	// units = B
	// base = 1024
	BlockSize int `json:"blockSize,omitempty"`
	// Minimum latency in milliseconds.
	MinLatency float64 `json:"minLatency,omitempty"`
	// Name of the test for which the grade is assigned.
	TestName string `json:"testName,omitempty"`
	// Maximum latency in milliseconds.
	MaxLatency float64 `json:"maxLatency,omitempty"`
	// IO operations per second.
	Iops int `json:"iops,omitempty"`
	// The test type.
	// enum = [READ WRITE RANDREAD RANDWRITE]
	TestType string `json:"testType,omitempty"`
	// No of jobs/threads used.
	Jobs int `json:"jobs,omitempty"`
}

// StatisticAxis - The attributes of a statistic axis.
// extends TypedObject
type StatisticAxis struct {
	// A deeper explanation of the data this corresponds to.
	Explanation string `json:"explanation,omitempty"`
	// The type of constraint that can be applied to this axis.
	ConstraintType string `json:"constraintType,omitempty"`
	// The type of value this axis will have for collected data.
	// enum = [INTEGER BOOLEAN STRING HISTOGRAM]
	ValueType string `json:"valueType,omitempty"`
	// Whether this axis appears as an attribute of a datapoint stream or of
	// datapoints themselves.
	StreamAttribute *bool `json:"streamAttribute,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name for this axis.
	AxisName string `json:"axisName,omitempty"`
}

// JSOperationEndpoint - The first and last JSOperation for a given data layout or branch.
// extends TypedObject
type JSOperationEndpoint struct {
	// The first JSOperation.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	First string `json:"first,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The last JSOperation.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	Last string `json:"last,omitempty"`
}

// WindowsHostCreateParameters - The parameters used for the add Windows host operation.
// extends HostCreateParameters
type WindowsHostCreateParameters struct {
	// Password for the certificate in the Java Keystore.
	// format = password
	// create = optional
	ConnectorCertificatePassword string `json:"connectorCertificatePassword,omitempty"`
	// Byte array of the Java Keystore data.
	// create = optional
	ConnectorKeystore string `json:"connectorKeystore,omitempty"`
	// Password for the Java Keystore data.
	// format = password
	// create = optional
	ConnectorKeystorePassword string `json:"connectorKeystorePassword,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The host object.
	// create = required
	Host Host `json:"host,omitempty"`
}

// SshVerifyRawKey - SSH verification strategy based on a known per-host key.
// extends SshVerifyBase
type SshVerifyRawKey struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of ssh key.
	// enum = [RSA DSA ECDSA ED25519]
	// required = true
	KeyType string `json:"keyType,omitempty"`
	// Base64-encoded ssh key of the host.
	// required = true
	// format = hostKey
	RawKey string `json:"rawKey,omitempty"`
}

// NamespaceFailoverParameters - The parameters to use as input to Namespace.failover.
// extends TypedObject
type NamespaceFailoverParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Enable automatic conflict resolution during failover.
	// create = optional
	// default = false
	SmartFailover *bool `json:"smartFailover,omitempty"`
}

// OracleDatabaseContainer - Data container for Oracle databases, both linked and virtual.
// extends DatabaseContainer
type OracleDatabaseContainer struct {
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Indicates whether or not this container has an associated LiveSource.
	// default = false
	LiveSource *bool `json:"liveSource,omitempty"`
	// If true, NOLOGGING operations on this container are treated as faults
	// and cannot be resolved manually. Otherwise, these operations are
	// ignored.
	// default = true
	// create = optional
	// update = optional
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
	// Optional user-provided description for the container.
	// update = optional
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Runtime properties of this container.
	Runtime *OracleDBContainerRuntime `json:"runtime,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = optional
	// update = optional
	PerformanceMode string `json:"performanceMode,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Indicates whether or not the database in this container is a physical
	// standby.
	// default = false
	PhysicalStandby *bool `json:"physicalStandby,omitempty"`
	// Indicates whether or not this container is ready for cross-platform
	// provisioning.
	// default = false
	CrossPlatformReady *bool `json:"crossPlatformReady,omitempty"`
	// Indicates whether this container is for a PDB, CDB root, auxiliary
	// CDB, or non-CDB.
	// enum = [PDB ROOT_CDB AUX_CDB NON_CDB]
	ContentType string `json:"contentType,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Indicates whether or not the database in this container consists only
	// of transportable tablespaces.
	// featureFlag = CONSPRO
	// default = false
	DatabaseFraction *bool `json:"databaseFraction,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// update = optional
	// create = optional
	SourcingPolicy *OracleSourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// update = optional
	// default = false
	// create = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
}

// MySQLDatabaseContainer - A MySQL Database Container.
// extends DatabaseContainer
type MySQLDatabaseContainer struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// Optional user-provided description for the container.
	// update = optional
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// A reference to the group containing this container.
	// referenceTo = /delphix-group.json
	// create = required
	// format = objectReference
	Group string `json:"group,omitempty"`
	// Runtime properties of this container.
	Runtime *MySQLDBContainerRuntime `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// Variant of the MySQL installation.
	// enum = [CommunityServer MariaDB]
	Variant string `json:"variant,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
}

// AppDataSnapshotRuntime - Runtime (non-persistent) properties of AppData TimeFlow snapshots.
// extends SnapshotRuntime
type AppDataSnapshotRuntime struct {
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MSSqlFailoverClusterListener - The representation of a SQL Server Failover Cluster Listener.
// extends MSSqlBaseClusterListener
type MSSqlFailoverClusterListener struct {
	// The name of the listener.
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The port for the listener.
	Port int `json:"port,omitempty"`
	// The address of the listener.
	Address string `json:"address,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ToolkitUpgradeDefinition - The toolkit upgrade logic to upgrade metadata from the previous
// version of the toolkit.
// extends TypedObject
type ToolkitUpgradeDefinition struct {
	// A workflow script to upgrade the linked source parameters.
	// required = true
	UpgradeLinkedSource string `json:"upgradeLinkedSource,omitempty"`
	// A workflow script to upgrade the manually-discovered source config
	// parameters.
	// required = false
	UpgradeManualSourceConfig string `json:"upgradeManualSourceConfig,omitempty"`
	// The version of the toolkit that we are upgrading from.
	// format = toolkitVersion
	// required = true
	FromVersion string `json:"fromVersion,omitempty"`
	// A workflow script to upgrade the snapshot metadata.
	// required = true
	UpgradeSnapshot string `json:"upgradeSnapshot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script to upgrade the virtual source parameters.
	// required = true
	UpgradeVirtualSource string `json:"upgradeVirtualSource,omitempty"`
}

// TimeflowPointBookmarkTag - TimeFlow point based on a TimeFlow bookmark tag.
// extends TimeflowPointParameters
type TimeflowPointBookmarkTag struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The name of the tag.
	// required = true
	Tag string `json:"tag,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	// required = true
	Container string `json:"container,omitempty"`
}

// CPUInfo - Describes a processor available to the system.
// extends TypedObject
type CPUInfo struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Speed of the processor, in hertz.
	// units = Hz
	// base = 1000
	Speed float64 `json:"speed,omitempty"`
	// Number of cores in the processor.
	Cores int `json:"cores,omitempty"`
}

// Domain - Represents the root container of all objects on the system.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type Domain struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
}

// ASEStagingSource - An SAP ASE staging source used for validated sync..
// extends ASESource
type ASEStagingSource struct {
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntime `json:"runtime,omitempty"`
	// The base mount point for the NFS mounts.
	// maxLength = 80
	// update = optional
	MountBase string `json:"mountBase,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The path to a user-provided shell script or executable to run on the
	// staging host after restoring from a backup during validated sync.
	// create = optional
	// update = optional
	// maxLength = 1024
	PostScript string `json:"postScript,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The path to a user-provided script or executable to run on the staging
	// host prior to restoring from a backup during validated sync.
	// create = optional
	// update = optional
	// maxLength = 1024
	PreScript string `json:"preScript,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
}

// MySQLExportParameters - The parameters to use as input to export MySQL databases.
// extends DbExportParameters
type MySQLExportParameters struct {
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The source config to use when creating the exported database.
	// required = true
	SourceConfig *MySQLServerConfig `json:"sourceConfig,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayout `json:"filesystemLayout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleDatabaseStatsSection - Oracle database performance statistics for a specific section.
// extends TypedObject
type OracleDatabaseStatsSection struct {
	// List of statistic column headers.
	ColumnHeaders []string `json:"columnHeaders,omitempty"`
	// List of statistic rows corresponding to column headers.
	RowValues []*OracleDatabaseStatistic `json:"rowValues,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Database statistic section name.
	SectionName string `json:"sectionName,omitempty"`
}

// OraclePDBAttachData - Represents parameters to attach an Oracle pluggable database.
// extends OracleBaseAttachData
type OraclePDBAttachData struct {
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// True if SnapSync data from the source should be compressed over the
	// network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially over
	// slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// True if attach should succeed even if the resetlogs of the original
	// database does not match the resetlogs of the new database and the
	// resetlogs information of the original database is not a parent
	// incarnation of the current resetlogs. This can happen when the
	// controlfile has been recreated and the incarnation table is
	// incomplete. Use this option with extreme caution. Attached database
	// must be the same database to avoid data corruption later on.
	// default = false
	// create = optional
	Force *bool `json:"force,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels int `json:"rmanChannels,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet int `json:"filesPerSet,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// True if SnapSync data from the source should be retrieved through an
	// encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Reference to the configuration for the PDB source.
	// format = objectReference
	// referenceTo = /delphix-oracle-pdb-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// maximum = 16
	// create = optional
	// default = 1
	// minimum = 1
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
}

// JSBookmarkDataParent - The bookmark data parent of a RESTORE or CREATE_BRANCH operation.
// extends JSDataParent
type JSBookmarkDataParent struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The bookmark that this operation's data came from. This will be null
	// if the bookmark has been deleted.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// This will always contain the name of the bookmark, even if it has been
	// deleted.
	// maxLength = 256
	BookmarkName string `json:"bookmarkName,omitempty"`
}

// TimeflowBookmark - A TimeFlow bookmark is a user defined name for a TimeFlow point
// (location or timestamp within a TimeFlow).
// extends NamedUserObject
type TimeflowBookmark struct {
	// The logical time corresponding to the TimeFlow location.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Indicates whether retention should be allowed to clean up the TimeFlow
	// bookmark and associated data.
	// update = optional
	RetentionProof *bool `json:"retentionProof,omitempty"`
	// Reference to the TimeFlow for this bookmark.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	// required = true
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The TimeFlow location.
	Location string `json:"location,omitempty"`
	// A tag for the bookmark that can be used to group TimeFlow bookmarks
	// together or qualify the type of the bookmark.
	// maxLength = 64
	Tag string `json:"tag,omitempty"`
}

// Namespace - Object namespace for target-side replication.
// extends NamedUserObject
type Namespace struct {
	// Type of object namespace.
	// enum = [REPLICATION]
	NamespaceType string `json:"namespaceType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Description of this namespace.
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// A unique identifier for the source data stream.
	Tag string `json:"tag,omitempty"`
	// True if the source data stream was generated from a secure replication
	// spec.
	SecureNamespace *bool `json:"secureNamespace,omitempty"`
	// Indicates the namespace has been failed over into the live
	// environment.
	// default = false
	FailedOver *bool `json:"failedOver,omitempty"`
	// If failedOver is true, contains a report concern objects affected
	// during failover.
	FailoverReport string `json:"failoverReport,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// For replication, the source host IP address.
	Source string `json:"source,omitempty"`
}

// RunExpectOnSourceOperation - A user-specifiable operation that runs an expect script on the target
// host.
// extends SourceOperation
type RunExpectOnSourceOperation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The expect command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
}

// OraclePDBLinkData - Represents parameters to link a Oracle pluggable database.
// extends OracleBaseLinkData
type OraclePDBLinkData struct {
	// True if SnapSync data from the source should be retrieved through an
	// encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// minimum = 1
	// maximum = 16
	// create = optional
	// default = 1
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Skip check that tests if there is enough space available to store the
	// database in the Delphix Engine. The Delphix Engine estimates how much
	// space a database will occupy after compression and prevents SnapSync
	// if insufficient space is available. This safeguard can be overridden
	// using this option. This may be useful when linking highly compressible
	// databases.
	// default = false
	// create = optional
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// default = false
	// create = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *OracleSourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// default = 5
	// create = optional
	// minimum = 1
	// maximum = 64
	FilesPerSet int `json:"filesPerSet,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// Number of parallel channels to use.
	// maximum = 32
	// default = 2
	// create = optional
	// minimum = 1
	RmanChannels int `json:"rmanChannels,omitempty"`
	// True if SnapSync data from the source should be compressed over the
	// network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially over
	// slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-pdb-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession to
	// reduce the number of logs required to provision the snapshot. This may
	// significantly reduce the time necessary to provision from a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// Information about the OS user to use for linking.
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// If true, NOLOGGING operations on this container are treated as faults
	// and cannot be resolved manually. Otherwise, these operations are
	// ignored.
	// default = true
	// create = optional
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
}

// OracleEnableParameters - The parameters to use as input to enable Oracle sources.
// extends SourceEnableParameters
type OracleEnableParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether to attempt a startup of the source after the enable.
	// default = true
	AttemptStart *bool `json:"attemptStart,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	Credential Credential `json:"credential,omitempty"`
}

// XppValidateStatus - Cross-platform provisioning validation status of the container.
// extends ChecklistItemDetail
type XppValidateStatus struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
}

// WindowsHostEnvironment - The representation of a windows host environment object.
// extends HostEnvironment
type WindowsHostEnvironment struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// A reference to the primary user for this environment.
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	// format = objectReference
	PrimaryUser string `json:"primaryUser,omitempty"`
	// The reference to the associated host.
	// referenceTo = /delphix-host.json
	// format = objectReference
	Host string `json:"host,omitempty"`
	// The reference to the proxy associated with the host.
	// referenceTo = /delphix-host.json
	// create = optional
	// update = optional
	// format = objectReference
	Proxy string `json:"proxy,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
}

// LinkParameters - Represents the parameters of a link request.
// extends TypedObject
type LinkParameters struct {
	// Database specific data required for linking.
	// required = true
	// create = required
	LinkData LinkData `json:"linkData,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// DSource name.
	// format = objectName
	// required = true
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the group containing this container.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	Description string `json:"description,omitempty"`
}

// WorkflowFunctionDefinition - The definition of a Workflow Function.
// extends TypedObject
type WorkflowFunctionDefinition struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of this function.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The input schema for this function according to DRAFTV4.
	// update = optional
	// create = required
	InputSchema *SchemaDraftV4 `json:"inputSchema,omitempty"`
	// The output schema for this function according to DRAFTV4.
	// create = required
	// update = optional
	OutputSchema *SchemaDraftV4 `json:"outputSchema,omitempty"`
}

// ASEDBContainerRuntime - Runtime properties of an SAP ASE database container.
// extends DBContainerRuntime
type ASEDBContainerRuntime struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
	// The source database backupset that was last restored in this
	// container.
	// format = date
	LastRestoredBackupDate string `json:"lastRestoredBackupDate,omitempty"`
	// The timezone for the last restored source database backupset in this
	// container.
	LastRestoredBackupTimeZone string `json:"lastRestoredBackupTimeZone,omitempty"`
}

// AppDataSnapshot - Snapshot of an AppData TimeFlow.
// extends TimeflowSnapshot
type AppDataSnapshot struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *AppDataTimeflowPoint `json:"firstChangePoint,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *AppDataSnapshotRuntime `json:"runtime,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *AppDataTimeflowPoint `json:"latestChangePoint,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	Metadata *Json `json:"metadata,omitempty"`
	// The toolkit associated with this snapshot.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
}

// JSDataSource - The data source used for Jet Stream data layouts.
// extends NamedUserObject
type JSDataSource struct {
	// Flag indicating whether the source is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Runtime properties of this data source.
	Runtime SourceConnectionInfo `json:"runtime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Dictates order of operations on data sources. Operations can be
	// performed in parallel for all sources or sequentially. Below are
	// possible valid and invalid orderings given an example data template
	// with 3 sources (A, B, and C).<br>Valid:<br>A B C<br>1 1 1
	// (parallel)<br>1 2 3 (sequential)<br>Invalid:<br>A B C<br>2 2 2<br>0 1
	// 2<br>2 3 4<br>1 2 2<br>In the sequential case the data source with
	// priority 1 is the first to be started and the last to be stopped. This
	// value is set on creation of the template's data sources and copied to
	// the data container's data sources.
	// create = optional
	// minimum = 1
	// default = 1
	Priority int `json:"priority,omitempty"`
	// Key/value pairs used to specify attributes for this data source.
	// create = optional
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// A description of this data source.
	// create = optional
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// A reference to the Jet Stream data layout to which this source
	// belongs.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
	// A reference to the underlying container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Flag indicating whether the source is masked.
	Masked *bool `json:"masked,omitempty"`
}

// ReplicationList - List of objects that are to be replicated.
// extends ReplicationObjectSpecification
type ReplicationList struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Objects to replicate, in canonical object reference form.
	// minItems = 1
	// create = required
	// update = optional
	Objects []string `json:"objects,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// AppDataTimeflowPoint - A unique point within an AppData TimeFlow.
// extends TimeflowPoint
type AppDataTimeflowPoint struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
}

// PgSQLTimeflow - TimeFlow representing historical data for a particular timeline within
// a PostgreSQL container.
// extends Timeflow
type PgSQLTimeflow struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *PgSQLTimeflowPoint `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// MySQLServerConfig - Configuration information for a MySQL server instance.
// extends SourceConfig
type MySQLServerConfig struct {
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The data directory for the MySQL server instance.
	// create = optional
	DataDirectory string `json:"dataDirectory,omitempty"`
	// The port on which the MySQL server instance is listening.
	// create = required
	// update = optional
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// The password of the server instance user.
	// create = optional
	// update = optional
	Credentials *PasswordCredential `json:"credentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The username of the server instance user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// PgSQLDBConfig - Configuration information for a PostgreSQL database in a cluster.
// extends TypedObject
type PgSQLDBConfig struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the database.
	// create = optional
	// update = optional
	// maxLength = 63
	DatabaseName string `json:"databaseName,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The PostgreSQL cluster this database is part of.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	DatabaseCluster string `json:"databaseCluster,omitempty"`
}

// StatisticSlice - Collects a slice of a multidimensional analytics statistic.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type StatisticSlice struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Axis constraints act as per-axis filters on data that is being
	// collected.
	// create = optional
	AxisConstraints []AxisConstraint `json:"axisConstraints,omitempty"`
	// The type name for the data this can collect.
	// create = required
	StatisticType string `json:"statisticType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Collection state of the slice.
	// enum = [INITIALIZED RUNNING PAUSED FAILED]
	State string `json:"state,omitempty"`
	// The set of axes to collect (usually these are not constrained axes).
	// create = required
	CollectionAxes []string `json:"collectionAxes,omitempty"`
	// The minimum interval between each reading for this statistic.
	// create = optional
	// units = sec
	CollectionInterval int `json:"collectionInterval,omitempty"`
}

// UnixHost - The representation of a Unix host object.
// extends Host
type UnixHost struct {
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Profile for escalating user privileges.
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// format = objectReference
	// create = optional
	// update = optional
	PrivilegeElevationProfile string `json:"privilegeElevationProfile,omitempty"`
	// Mechanism to use for ssh host verification.
	// create = optional
	// update = optional
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// The password for the user managed truststore.
	// format = password
	// create = optional
	// update = optional
	// minLength = 1
	UserTruststorePassword string `json:"userTruststorePassword,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The address associated with the host.
	// format = host
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// The port number used to connect to the host via SSH.
	// default = 22
	// minimum = 1
	// maximum = 65535
	// create = optional
	// update = optional
	SshPort int `json:"sshPort,omitempty"`
	// The date the host was added.
	DateAdded string `json:"dateAdded,omitempty"`
	// Runtime properties for this host.
	HostRuntime *HostRuntime `json:"hostRuntime,omitempty"`
	// The path for the toolkit that resides on the host.
	// create = required
	// update = optional
	ToolkitPath string `json:"toolkitPath,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The X.509 certificate chain for the host. Can be modified via
	// importCertificateChain.
	// update = readonly
	// create = readonly
	CertificateChain []string `json:"certificateChain,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The host configuration object associated with the host.
	HostConfiguration *HostConfiguration `json:"hostConfiguration,omitempty"`
}

// JSDataContainerActiveBranchParameters - Input parameters for the API that given a point in time, returns the
// active branch of the data container.
// extends TypedObject
type JSDataContainerActiveBranchParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The time that will be used to find which branch was active in the data
	// layout.
	// format = date
	// required = true
	Time string `json:"time,omitempty"`
}

// ASEBackupLocation - SAP ASE backup location.
// extends TypedObject
type ASEBackupLocation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Name of the backup server instance.
	// create = required
	// update = optional
	BackupServerName string `json:"backupServerName,omitempty"`
	// OS user for the host where the backup server is located.
	// create = required
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	BackupHostUser string `json:"backupHostUser,omitempty"`
	// Host environment where the backup server is located.
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	// update = optional
	BackupHost string `json:"backupHost,omitempty"`
}

// Job - Represents a job object.
// extends NamedUserObject
// cliVisibility = [SYSTEM DOMAIN]
type Job struct {
	// Time the job was last updated.
	// format = date
	UpdateTime string `json:"updateTime,omitempty"`
	// User that initiated the action.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// Action type of the Job.
	ActionType string `json:"actionType,omitempty"`
	// Object type of the target.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Title of the job.
	Title string `json:"title,omitempty"`
	// Time the job was created. Note that this is not the time when the job
	// started executing.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Completion percentage. This value is a copy of the last event's
	// percentComplete. It will be 0 if there are no job events or if the
	// events field is not populated while fetching the job.
	// units = %
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A description of why the job was canceled.
	CancelReason string `json:"cancelReason,omitempty"`
	// Whether this job is waiting for resources to be available for its
	// execution.
	Queued *bool `json:"queued,omitempty"`
	// A cached copy of the target object name.
	TargetName string `json:"targetName,omitempty"`
	// This job's parent action.
	// format = objectReference
	// referenceTo = /delphix-action.json
	ParentAction string `json:"parentAction,omitempty"`
	// State of this job's parent action. This value is populated only if the
	// job is fetched via the plain get API call.
	// enum = [EXECUTING WAITING COMPLETED FAILED CANCELED]
	ParentActionState string `json:"parentActionState,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// State of the job.
	// enum = [RUNNING SUSPENDED CANCELED COMPLETED FAILED]
	JobState string `json:"jobState,omitempty"`
	// Object reference of the target.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Whether this job can be canceled.
	Cancelable *bool `json:"cancelable,omitempty"`
	// A list of time-sorted past JobEvent objects associated with this job.
	Events []*JobEvent `json:"events,omitempty"`
	// Whether this job can be suspended.
	Suspendable *bool `json:"suspendable,omitempty"`
	// Email addresses to be notified on job notification alerts.
	// update = required
	EmailAddresses []string `json:"emailAddresses,omitempty"`
}

// RunCommandOnSourceOperation - A user-specifiable operation that runs a shell command on the target
// host.
// extends SourceOperation
type RunCommandOnSourceOperation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The shell command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
}

// KerberosConfig - Kerberos Client Configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type KerberosConfig struct {
	// Kerberos Realm name.
	// create = required
	// update = optional
	Realm string `json:"realm,omitempty"`
	// One of more KDC servers.
	// create = required
	// update = optional
	// minItems = 1
	Kdcs []*KerberosKDC `json:"kdcs,omitempty"`
	// Kerberos keytab file data in base64 encoding.
	// format = password
	// create = required
	// update = optional
	Keytab string `json:"keytab,omitempty"`
	// Kerberos principal name.
	// create = required
	// update = optional
	Principal string `json:"principal,omitempty"`
	// Indicates whether kerberos has been configured or not.
	Enabled *bool `json:"enabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlAttachData - Represents the MSSQL specific parameters of an attach request.
// extends AttachData
type MSSqlAttachData struct {
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mssql-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The SQL Server instance on the staging environment to use for
	// pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
	// Specifies the backup types validated sync will use to synchronize the
	// dSource with the source database.
	// enum = [TRANSACTION_LOG FULL_OR_DIFFERENTIAL FULL NONE]
	// default = TRANSACTION_LOG
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// The user for accessing the shared backup location.
	// create = optional
	// maxLength = 256
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// The password for accessing the shared backup location.
	// create = optional
	BackupLocationCredentials *PasswordCredential `json:"backupLocationCredentials,omitempty"`
	// The SQL Server login password for the source DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// External file path.
	// create = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// The path to a user-provided PowerShell script or executable to run
	// prior to restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// The path to a user-provided PowerShell script or executable to run
	// after restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// OS user on the PPT host to use for linking.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// format = objectReference
	PptHostUser string `json:"pptHostUser,omitempty"`
	// OS user on the source to use for linking.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// format = objectReference
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// The SQL Server login username for the source DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Shared source database backup location.
	// create = optional
	// maxLength = 260
	SharedBackupLocation string `json:"sharedBackupLocation,omitempty"`
}

// MSSqlFailoverClusterDriveLetter - This represents a logical volume with a drive letter that resides on a
// Physical Disk cluster resource that is part of a SQL Server Failover
// Cluster Instance.
// extends TypedObject
type MSSqlFailoverClusterDriveLetter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The drive letter.
	// minLength = 1
	// maxLength = 1
	DriveLetter string `json:"driveLetter,omitempty"`
	// The drive letter label.
	// maxLength = 32
	Label string `json:"label,omitempty"`
}

// JSTimestampDataParent - The timestamp data parent of a REFRESH, RESTORE, UNDO or CREATE_BRANCH
// operation.
// extends JSDataParent
type JSTimestampDataParent struct {
	// This will always contain the name of the branch, even if it has been
	// deleted.
	// maxLength = 256
	BranchName string `json:"branchName,omitempty"`
	// The data time on the branch that this operation's data came from.
	// format = date
	Time string `json:"time,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The branch this operation's data came from. This will be null if the
	// branch has been deleted.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
}

// ASESourceRuntime - Runtime (non-persistent) properties of a SAP ASE source.
// extends SourceRuntime
type ASESourceRuntime struct {
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// SAP ASE database durability level.
	// enum = [FULL AT_SHUTDOWN NO_RECOVERY]
	DurabilityLevel string `json:"durabilityLevel,omitempty"`
	// True if configured to truncate log on checkpoint.
	TruncateLogOnCheckpoint *bool `json:"truncateLogOnCheckpoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
}

// OracleNodeListener - An Oracle node listener.
// extends OracleListener
type OracleNodeListener struct {
	// The list of protocol addresses for this listener. These are used for
	// the local_listener parameter when provisioning VDBs.
	// minItems = 1
	// create = required
	// update = optional
	ProtocolAddresses []string `json:"protocolAddresses,omitempty"`
	// The list of client endpoints for this listener of the format
	// hostname:port. These are used when constructing the JDBC connection
	// string.
	// create = readonly
	// update = readonly
	ClientEndpoints []string `json:"clientEndpoints,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the environment this listener is associated with.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Whether this listener was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Reference to the host this listener is associated with.
	// referenceTo = /delphix-host.json
	// create = required
	// format = objectReference
	Host string `json:"host,omitempty"`
}

// TracerouteInfo - Trace route info from target host to Delphix Engine.
// extends TypedObject
type TracerouteInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Latency of network hops from host to Delphix Engine.
	NetworkHops string `json:"networkHops,omitempty"`
}

// JSBookmarkUsageData - The space usage information for a Jet Stream bookmark.
// extends TypedObject
type JSBookmarkUsageData struct {
	// The amount of space that will be freed if this bookmark is deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space referenced by this bookmark that cannot be freed
	// up by deleting this bookmark because it is also referenced by
	// neighboring bookmarks or branches that have been created or restored
	// from this bookmark.
	// base = 1024
	// units = B
	Shared float64 `json:"shared,omitempty"`
	// The amount of space referenced by this bookmark that cannot be freed
	// up by deleting this bookmark because it is also being referenced
	// outside of Jet Stream (e.g. by retention policy).
	// base = 1024
	// units = B
	ExternallyReferenced float64 `json:"externallyReferenced,omitempty"`
	// The Jet Stream bookmark that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
	// The data layout that this bookmark belongs to.
	DataLayout string `json:"dataLayout,omitempty"`
}

// ASEHostEnvironmentParameters - SAP ASE host environment parameters.
// extends TypedObject
type ASEHostEnvironmentParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// create = optional
	// update = optional
	// maxLength = 256
	DbUser string `json:"dbUser,omitempty"`
	// The credentials of the database user.
	// create = required
	// update = optional
	// properties = map[type:map[format:type default:PasswordCredential type:string description:Object type. required:true]]
	Credentials Credential `json:"credentials,omitempty"`
}

// NetworkThroughputTest - Bi-directional throughput tests to a target system.
// extends NetworkThroughputTestBase
// cliVisibility = [DOMAIN SYSTEM]
type NetworkThroughputTest struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Number of connections used to achieve maximum sustained throughput.
	NumConnections int `json:"numConnections,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkThroughputTestParameters `json:"parameters,omitempty"`
}

// OracleCharacterSet - Represents an Oracle character set.
// extends TypedObject
type OracleCharacterSet struct {
	// Name of character set.
	CharacterSet string `json:"characterSet,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASEPlatformParameters - ASE platform-specific parameters that are stored on a transformation.
// extends BasePlatformParameters
type ASEPlatformParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ToolkitLinkedDirectSource - A linked source definition for toolkits that perform linking using
// direct file sync.
// extends ToolkitLinkedSource
type ToolkitLinkedDirectSource struct {
	// A workflow script to run just prior to snapshotting the staged source.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// A workflow script to run immediately after snapshotting the staged
	// source.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A user defined schema for the linking parameters.
	// required = true
	Parameters *SchemaDraftV4 `json:"parameters,omitempty"`
}

// OracleAddLiveSourceParameters - The parameters to use as input to convert an Oracle dSource to an
// Oracle LiveSource.
// extends AddLiveSourceParameters
type OracleAddLiveSourceParameters struct {
	// The name of the privileged user to run the LiveSource creation
	// operation as.
	// create = optional
	Username string `json:"username,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The source that describes the LiveSource.
	// required = true
	Source *OracleLiveSource `json:"source,omitempty"`
	// The source config of the LiveSource.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// The security credential of the privileged user to run the LiveSource
	// creation operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
}

// MSSqlSnapshotRuntime - Runtime (non-persistent) properties of a MSSQL TimeFlow snapshot.
// extends SnapshotRuntime
type MSSqlSnapshotRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
}

// TimeflowPointTimestamp - TimeFlow point based on a timestamp.
// extends TimeflowPointParameters
type TimeflowPointTimestamp struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// required = true
	Timestamp string `json:"timestamp,omitempty"`
}

// MSSqlStagingSource - An MSSQL staging source used for pre-provisioning.
// extends MSSqlSource
type MSSqlStagingSource struct {
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntime `json:"runtime,omitempty"`
	// The base mount point for the iSCSI LUN mounts.
	MountBase string `json:"mountBase,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A user-provided PowerShell script or executable to run after restoring
	// from a backup during pre-provisioning.
	// create = optional
	// update = optional
	// maxLength = 1024
	PostScript string `json:"postScript,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// restoring from a backup during pre-provisioning.
	// maxLength = 1024
	// create = optional
	// update = optional
	PreScript string `json:"preScript,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
}

// MSSqlProvisionParameters - The parameters to use as input to provision MSSQL databases.
// extends ProvisionParameters
type MSSqlProvisionParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *MSSqlDatabaseContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *MSSqlVirtualSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	MaskingJob string `json:"maskingJob,omitempty"`
}

// PathDescendantConstraint - Constraint placed on a filesystem path axis of a particular analytics
// slice.
// extends PathConstraint
type PathDescendantConstraint struct {
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be a descendant of this path.
	// format = unixpath
	// create = required
	DescendantOf string `json:"descendantOf,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ASESourceConnectionInfo - Contains information that can be used to connect to an SAP ASE source.
// extends SourceConnectionInfo
type ASESourceConnectionInfo struct {
	// User to use for connecting to the source.
	User string `json:"user,omitempty"`
	// The JDBC string used to connect to the source.
	JdbcString string `json:"jdbcString,omitempty"`
	// The database name of the source.
	DatabaseName string `json:"databaseName,omitempty"`
	// Host to use for connecting to the source.
	Host string `json:"host,omitempty"`
	// Port to use for connecting to the source.
	Port int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// TCPStatsDatapointStream - A stream of datapoints from a TCP_STATS analytics slice.
// extends DatapointStream
type TCPStatsDatapointStream struct {
	// The local Delphix Engine IP address.
	// format = ipAddress
	LocalAddress string `json:"localAddress,omitempty"`
	// The remote IP address.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// The local TCP port number.
	LocalPort int `json:"localPort,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// The remote TCP port number.
	RemotePort int `json:"remotePort,omitempty"`
}

// SourceDisableParameters - The parameters to use as input to disable a MSSQL, PostgreSQL,
// AppData, ASE or MySQL source.
// extends TypedObject
type SourceDisableParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether to attempt a cleanup of the database from the environment
	// before the disable.
	// default = true
	AttemptCleanup *bool `json:"attemptCleanup,omitempty"`
}

// StatisticEnumAxis - The attributes of a statistic axis which is an enum type.
// extends StatisticAxis
type StatisticEnumAxis struct {
	// Whether this axis appears as an attribute of a datapoint stream or of
	// datapoints themselves.
	StreamAttribute *bool `json:"streamAttribute,omitempty"`
	// The set of values that are allowed for this axis.
	EnumValues []string `json:"enumValues,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name for this axis.
	AxisName string `json:"axisName,omitempty"`
	// A deeper explanation of the data this corresponds to.
	Explanation string `json:"explanation,omitempty"`
	// The type of constraint that can be applied to this axis.
	ConstraintType string `json:"constraintType,omitempty"`
	// The type of value this axis will have for collected data.
	// enum = [INTEGER BOOLEAN STRING HISTOGRAM]
	ValueType string `json:"valueType,omitempty"`
}

// OracleCluster - The representation of an oracle cluster environment object.
// extends SourceEnvironment
type OracleCluster struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The name of the cluster.
	// create = optional
	// maxLength = 15
	CrsClusterName string `json:"crsClusterName,omitempty"`
	// The version of the cluster.
	// create = optional
	// maxLength = 14
	Version string `json:"version,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// Indicates whether the Single Client Access Name of the cluster is
	// manually configured.
	ScanManual *bool `json:"scanManual,omitempty"`
	// A reference to the cluster user.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	ClusterUser string `json:"clusterUser,omitempty"`
	// The default remote_listener parameter to be used for databases on the
	// cluster.
	// create = optional
	// update = optional
	// maxLength = 256
	RemoteListener string `json:"remoteListener,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The Single Client Access Name of the cluster (11.2 and greater
	// clusters only).
	// maxLength = 256
	// update = optional
	Scan string `json:"scan,omitempty"`
	// The location of the cluster installation.
	// create = required
	// update = optional
	// maxLength = 256
	CrsClusterHome string `json:"crsClusterHome,omitempty"`
}

// MaskingJob - The Delphix Engine record of an existing Masking Job.
// extends UserObject
type MaskingJob struct {
	// The application id from the Delphix Masking Engine instance that
	// specifies the desired Masking Job.
	// create = required
	// update = readonly
	ApplicationId string `json:"applicationId,omitempty"`
	// The masking job id from the Delphix Masking Engine instance that
	// specifies the desired Masking Job.
	// create = required
	// update = readonly
	MaskingJobId string `json:"maskingJobId,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = required
	// update = readonly
	Name string `json:"name,omitempty"`
	// A reference to the container that the Masking Job is intended to
	// operate on.
	// referenceTo = /delphix-container.json
	// create = optional
	// update = optional
	// format = objectReference
	AssociatedContainer string `json:"associatedContainer,omitempty"`
}

// NetworkDSPTestParameters - Parameters used to execute a network throughput test using the Delphix
// Session Protocol.
// extends NetworkThroughputTestBaseParameters
type NetworkDSPTestParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The number of connections to use for the test. The special value 0
	// (the default) causes the test to automatically discover and use the
	// optimal number of connections to use for maximum throughput.
	// maximum = 32
	// default = 0
	// create = optional
	// minimum = 0
	NumConnections int `json:"numConnections,omitempty"`
	// The size of the send socket buffer in bytes.
	// default = 262144
	// minimum = 0
	// maximum = 1.6777216e+07
	// base = 1024
	// units = B
	// create = optional
	SendSocketBuffer int `json:"sendSocketBuffer,omitempty"`
	// The size of the receive socket buffer in bytes.
	// default = 262144
	// minimum = 0
	// maximum = 1.6777216e+07
	// base = 1024
	// units = B
	// create = optional
	ReceiveSocketBuffer int `json:"receiveSocketBuffer,omitempty"`
	// Whether or not compression is used for the test.
	// default = false
	// create = optional
	Compression *bool `json:"compression,omitempty"`
	// The size of each transmit request in bytes.
	// maximum = 1.048576e+06
	// base = 1024
	// units = B
	// create = optional
	// default = 65536
	// minimum = 0
	BlockSize int `json:"blockSize,omitempty"`
	// The queue depth used for the DSP throughput test.
	// default = 32
	// minimum = 0
	// maximum = 4096
	// create = optional
	QueueDepth int `json:"queueDepth,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// referenceTo = /delphix-host.json
	// create = optional
	// format = objectReference
	RemoteHost string `json:"remoteHost,omitempty"`
	// The duration of the test in seconds. Note that when numConnections is
	// 0, an initial period of time will be spent calculating the optimal
	// number of connections, and that time does not count toward the
	// duration of the test.
	// maximum = 3600
	// default = 30
	// create = optional
	// minimum = 1
	Duration int `json:"duration,omitempty"`
	// Whether the test is a transmit or receive test.
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	// create = optional
	Direction string `json:"direction,omitempty"`
	// Whether the test is testing connectivity to a Delphix Engine or remote
	// host.
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	// default = REMOTE_HOST
	// create = optional
	DestinationType string `json:"destinationType,omitempty"`
	// Address, username and password used when running a test to another
	// Delphix Engine.
	// create = optional
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfo `json:"remoteDelphixEngineInfo,omitempty"`
	// Whether or not encryption is used for the test.
	// default = false
	// create = optional
	Encryption *bool `json:"encryption,omitempty"`
}

// MSSqlLinkedSourceUpgradeParameters - The parameters to use as input to upgrade an MSSQL linked source.
// extends SourceUpgradeParameters
type MSSqlLinkedSourceUpgradeParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The source config that the source database upgrades to.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// The SQL instance on the PPT environment that we want to use for
	// pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
}

// TCPStatsDatapoint - An analytics datapoint generated by the TCP_STATS statistic type.
// extends Datapoint
type TCPStatsDatapoint struct {
	// The size of the peer's receive window.
	// units = B
	// base = 1024
	SendWindowSize int `json:"sendWindowSize,omitempty"`
	// The size of the local receive window.
	// units = B
	// base = 1024
	ReceiveWindowSize int `json:"receiveWindowSize,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Data bytes transmitted.
	// units = B
	// base = 1024
	OutBytes int `json:"outBytes,omitempty"`
	// Number of bytes sent but unacknowledged.
	// units = B
	// base = 1024
	UnacknowledgedBytes int `json:"unacknowledgedBytes,omitempty"`
	// Bytes retransmitted.
	// units = B
	// base = 1024
	RetransmittedBytes int `json:"retransmittedBytes,omitempty"`
	// The size of the local congestion window.
	// units = B
	// base = 1024
	CongestionWindowSize int `json:"congestionWindowSize,omitempty"`
	// The smoothed average round trip time for this connection (us).
	// units = usec
	RoundTripTime int `json:"roundTripTime,omitempty"`
	// Data bytes received.
	// units = B
	// base = 1024
	InBytes int `json:"inBytes,omitempty"`
	// Number of bytes received out of order. This is a subset of the
	// 'inBytes' value.
	// units = B
	// base = 1024
	InUnorderedBytes int `json:"inUnorderedBytes,omitempty"`
}

// ProvisionCompatibilityParameters - The criteria necessary to select valid repositories for provisioning.
// extends CompatibleRepositoriesParameters
type ProvisionCompatibilityParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The TimeFlow point to use as a source of compatibility information.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// ASETimeflow - TimeFlow representing historical data for a particular timeline within
// a SAP ASE data container.
// extends Timeflow
type ASETimeflow struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *ASETimeflowPoint `json:"parentPoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
}

// ErrorResult - Result of a failed API call.
// extends CallResult
type ErrorResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
	// Specifics of the error that occurred during API call execution.
	Error *APIError `json:"error,omitempty"`
}

// RunBashOnSourceOperation - A user-specifiable operation that runs a shell command on the target
// host using the Delphix supplied Bash shell.
// extends SourceOperation
type RunBashOnSourceOperation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The shell command to execute on the target host.
	// create = required
	// update = required
	Command string `json:"command,omitempty"`
}

// Statistic - Multidimensional analytics statistics which can be queried for data.
// extends TypedObject
// cliVisibility = [DOMAIN]
type Statistic struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A deeper explanation of the data this can collect.
	Explanation string `json:"explanation,omitempty"`
	// The smallest unit of time this statistic can measure on.
	// units = sec
	MinCollectionInterval int `json:"minCollectionInterval,omitempty"`
	// The set of axes this statistic has.
	Axes []*StatisticAxis `json:"axes,omitempty"`
	// The type name for the data this can collect.
	StatisticType string `json:"statisticType,omitempty"`
}

// SystemInfo - Retrieve system-wide properties and manage the state of the system.
// extends PublicSystemInfo
// cliVisibility = [DOMAIN SYSTEM]
type SystemInfo struct {
	// The list of enabled features on this Delphix Engine.
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
	// Indicates whether the server has gone through initial setup or not.
	Configured *bool `json:"configured,omitempty"`
	// The current system locale.
	// format = locale
	CurrentLocale string `json:"currentLocale,omitempty"`
	// Product type.
	ProductType string `json:"productType,omitempty"`
	// List of available locales.
	Locales []string `json:"locales,omitempty"`
	// Name of the product that the system is running.
	ProductName string `json:"productName,omitempty"`
	// Amount of raw storage used by dSources, VDBs and system metadata.
	// base = 1024
	// units = B
	StorageUsed float64 `json:"storageUsed,omitempty"`
	// Processors on the system.
	Processors []*CPUInfo `json:"processors,omitempty"`
	// The date and time that the Delphix Engine was installed.
	// format = date
	InstallationTime string `json:"installationTime,omitempty"`
	// Maximum supported API version of the current system software.
	ApiVersion *APIVersion `json:"apiVersion,omitempty"`
	// Time at which the current system software was built.
	// format = date
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	// SSH public key to be added to SSH authorized_keys for environment
	// users using the SystemKeyCredential authorization mechanism.
	SshPublicKey string `json:"sshPublicKey,omitempty"`
	// Description of the current system software.
	BuildTitle string `json:"buildTitle,omitempty"`
	// Globally unique identifier for this software installation.
	Uuid string `json:"uuid,omitempty"`
	// Delphix version of the current system software.
	BuildVersion *VersionInfo `json:"buildVersion,omitempty"`
	// Description of the current system platform.
	Platform string `json:"platform,omitempty"`
	// Total amount of raw storage allocated for dSources, VDBs, and system
	// metadata. Zero if storage has not yet been configured.
	// units = B
	// base = 1024
	StorageTotal float64 `json:"storageTotal,omitempty"`
	// Total memory on the system, in bytes.
	// units = B
	// base = 1024
	MemorySize float64 `json:"memorySize,omitempty"`
	// System hostname.
	// format = hostname
	// update = optional
	Hostname string `json:"hostname,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Security banner to display prior to login.
	Banner string `json:"banner,omitempty"`
}

// VirtualSourceOperations - Describes operations which are performed on virtual sources at various
// times.
// extends TypedObject
type VirtualSourceOperations struct {
	// Operations to perform before snapshotting a virtual source. These
	// operations can quiesce any data prior to snapshotting.
	// update = optional
	// create = optional
	PreSnapshot []SourceOperation `json:"preSnapshot,omitempty"`
	// Operations to perform after snapshotting a virtual source.
	// update = optional
	// create = optional
	PostSnapshot []SourceOperation `json:"postSnapshot,omitempty"`
	// Operations to perform before stopping a virtual source.
	// create = optional
	// update = optional
	PreStop []SourceOperation `json:"preStop,omitempty"`
	// Operations to perform after stopping a virtual source.
	// update = optional
	// create = optional
	PostStop []SourceOperation `json:"postStop,omitempty"`
	// Operations to perform before rewinding a virtual source. These
	// operations can backup any data or configuration from the running
	// source prior to rewinding.
	// create = optional
	// update = optional
	PreRollback []SourceOperation `json:"preRollback,omitempty"`
	// Operations to perform after refreshing a virtual source. These
	// operations can be used to restore any data or configuration backed up
	// in the preRefresh operations.
	// update = optional
	// create = optional
	PostRefresh []SourceOperation `json:"postRefresh,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Operations to perform before refreshing a virtual source. These
	// operations can backup any data or configuration from the running
	// source before doing the refresh.
	// create = optional
	// update = optional
	PreRefresh []SourceOperation `json:"preRefresh,omitempty"`
	// Operations to perform before starting a virtual source.
	// create = optional
	// update = optional
	PreStart []SourceOperation `json:"preStart,omitempty"`
	// Operations to perform when initially creating the virtual source and
	// every time it is refreshed.
	// create = optional
	// update = optional
	ConfigureClone []SourceOperation `json:"configureClone,omitempty"`
	// Operations to perform after rewinding a virtual source. These
	// operations can be used to automate processes once the rewind is
	// complete.
	// create = optional
	// update = optional
	PostRollback []SourceOperation `json:"postRollback,omitempty"`
	// Operations to perform after starting a virtual source.
	// update = optional
	// create = optional
	PostStart []SourceOperation `json:"postStart,omitempty"`
}

// NotFilter - A container filter that inverts the logic of another filter.
// extends AlertFilter
type NotFilter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Filter whose logic is to be inverted.
	// create = required
	// update = optional
	// properties = map[type:map[default:SeverityFilter]]
	SubFilter AlertFilter `json:"subFilter,omitempty"`
}

// PopulateCompatibilityParameters - The criteria necessary to select valid repositories to populate into a
// warehouse.
// extends CompatibleRepositoriesParameters
type PopulateCompatibilityParameters struct {
	// Restrict returned repositories to this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = optional
	// update = optional
	Environment string `json:"environment,omitempty"`
	// The warehouse repository to use as a source of compatibility
	// information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleMissingLog - An Oracle missing log file.
// extends OracleTimeflowLog
type OracleMissingLog struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the database to which this log belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-container.json
	Container string `json:"container,omitempty"`
	// Instance number associated with the log file.
	InstanceNum int `json:"instanceNum,omitempty"`
	// Sequence number for the log file.
	Sequence int `json:"sequence,omitempty"`
	// End SCN for the log file.
	EndScn int `json:"endScn,omitempty"`
	// Start timestamp for the log file.
	// format = date
	StartTimestamp string `json:"startTimestamp,omitempty"`
	// End timestamp for the log file.
	// format = date
	EndTimestamp string `json:"endTimestamp,omitempty"`
	// Reference to the TimeFlow of which this log is a part.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Start SCN for the log file.
	StartScn int `json:"startScn,omitempty"`
}

// OracleLog - Oracle log file.
// extends TypedObject
type OracleLog struct {
	// Sequence number for the log file.
	Sequence int `json:"sequence,omitempty"`
	// Instance number associated with the log file.
	InstanceNum int `json:"instanceNum,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleCustomEnvVarRACPair - Dictates a single environment variable name and value to be set when
// the Delphix Engine administers an Oracle virtual database. For a RAC
// environment, the cluster node where the target pair is valid must also
// be specified.
// extends OracleCustomEnvVarPair
type OracleCustomEnvVarRACPair struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the environment variable.
	// format = envvarIdentifier
	// required = true
	VarName string `json:"varName,omitempty"`
	// The value of the environment variable.
	// required = true
	VarValue string `json:"varValue,omitempty"`
	// The cluster node on which the environment variable is relevant.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// required = true
	ClusterNode string `json:"clusterNode,omitempty"`
}

// AttachSourceParameters - Represents the parameters of an attach request.
// extends TypedObject
type AttachSourceParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The database-specific parameters of an attach request.
	// required = true
	AttachData AttachData `json:"attachData,omitempty"`
}

// MySQLBinlogCoordinates - The current position in the master binary logs when a MySQL backup was
// taken.
// extends MySQLReplicationCoordinates
type MySQLBinlogCoordinates struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Name of the master log file to start replication from.
	// required = true
	MasterLogName string `json:"masterLogName,omitempty"`
	// Position within the master log file to start replication from.
	// required = true
	MasterLogPos int `json:"masterLogPos,omitempty"`
}

// AppDataCachedMountPoint - Specified information about an active mount of an AppData container.
// extends TypedObject
type AppDataCachedMountPoint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Absolute path on the target environment were the filesystem is
	// mounted.
	// update = optional
	// format = unixpath
	// create = required
	MountPath string `json:"mountPath,omitempty"`
	// Reference to the environment on which the file system is mounted.
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	// update = optional
	Environment string `json:"environment,omitempty"`
	// Order in mount sequence.
	// create = required
	// update = optional
	Ordinal int `json:"ordinal,omitempty"`
	// Relative path within the container of the directory that is mounted.
	// format = unixpath
	// create = optional
	// update = optional
	SharedPath string `json:"sharedPath,omitempty"`
}

// SystemVersion - Describes a Delphix software revision.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type SystemVersion struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Delphix version number.
	Version string `json:"version,omitempty"`
	// The minimum DelphixOS version supported by this Delphix version.
	MinOsVersion string `json:"minOsVersion,omitempty"`
	// DelphixOS is running from this version.
	OsRunning *bool `json:"osRunning,omitempty"`
	// The minimum required Delphix version in order to upgrade.
	MinVersion string `json:"minVersion,omitempty"`
	// Date on which this version was last verified.
	// format = date
	VerifyDate string `json:"verifyDate,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The DelphixOS version number.
	OsVersion string `json:"osVersion,omitempty"`
	// Date on which the version was built.
	// format = date
	BuildDate string `json:"buildDate,omitempty"`
	// Date on which this version was installed.
	// format = date
	InstallDate string `json:"installDate,omitempty"`
	// The state of the version.
	// enum = [PREVIOUS CURRENTLY_RUNNING DEFERRED UPLOADED UNPACKING DELETING VERIFYING VERIFIED APPLYING UNKNOWN DISABLE_FAILED]
	Status string `json:"status,omitempty"`
}

// MSSqlExistingSpecificBackupSyncParameters - The parameters to use as input to sync MSSQL databases using an
// existing specific full or differential backup.
// extends MSSqlExistingBackupSyncParameters
type MSSqlExistingSpecificBackupSyncParameters struct {
	// The UUID of the full or differential backup of the source database to
	// restore from. The backup should be present in the shared backup
	// location for the source database.
	// required = true
	BackupUUID string `json:"backupUUID,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleService - The representation of an oracle service object.
// extends TypedObject
type OracleService struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The connection string used to connect to JDBC.
	// format = oracleJDBCConnectionString
	// create = required
	// update = optional
	JdbcConnectionString string `json:"jdbcConnectionString,omitempty"`
	// Whether this service was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// HostConfiguration - The representation of the host configuration properties.
// extends TypedObject
type HostConfiguration struct {
	// The host machine information.
	Machine *HostMachine `json:"machine,omitempty"`
	// The timestamp when the host was last updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// Indicates whether the host configuration properties were discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The timestamp when the host was last refreshed.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
	// The host operating system information.
	OperatingSystem *HostOS `json:"operatingSystem,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLInstall - A MySQL installation.
// extends SourceRepository
type MySQLInstall struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// update = optional
	// default = false
	Staging *bool `json:"staging,omitempty"`
	// Version string for the repository.
	// create = readonly
	// update = readonly
	Version string `json:"version,omitempty"`
	// Version of the MySQL installation.
	InternalVersion *MySQLVersion `json:"internalVersion,omitempty"`
	// Directory path where the MySQL installation is located.
	// create = required
	InstallationPath string `json:"installationPath,omitempty"`
	// Flag indicating whether the MySQL installation was discovered or
	// manually entered.
	Discovered *bool `json:"discovered,omitempty"`
}

// DiagnosisResult - Details from a diagnosis check that was run due to a failed operation.
// extends TypedObject
type DiagnosisResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Message code associated with the event.
	MessageCode string `json:"messageCode,omitempty"`
	// Localized message.
	Message string `json:"message,omitempty"`
	// True if this was a check that did not pass.
	Failure *bool `json:"failure,omitempty"`
}

// PhoneHomeService - Phone home service configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type PhoneHomeService struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// True if the phone home service is enabled.
	// required = true
	Enabled *bool `json:"enabled,omitempty"`
}

// CurrentGroupCapacityData - Capacity data aggregated over a group.
// extends BaseGroupCapacityData
// cliVisibility = [DOMAIN]
type CurrentGroupCapacityData struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdown `json:"source,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdown `json:"virtual,omitempty"`
	// Which group these stats represent.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
}

// SnapshotSpaceResult - Result of the operation to determine how much space is used by a set
// of snapshots.
// extends TypedObject
type SnapshotSpaceResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Total amount of space, in bytes, that would be freed by deleting the
	// input snapshots.
	// units = B
	// base = 1024
	TotalSize float64 `json:"totalSize,omitempty"`
}

// TimeflowLogFetchParameters - Parameters to fetch log files within a TimeFlow.
// extends LogFetchSSH
type TimeflowLogFetchParameters struct {
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// Reference to the TimeFlow for which to fetch logs.
	// referenceTo = /delphix-oracle-timeflow.json
	// required = true
	// format = objectReference
	Timeflow string `json:"timeflow,omitempty"`
	// The ending SCN of the range of log files to fetch.
	// required = true
	EndLocation string `json:"endLocation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port int `json:"port,omitempty"`
	// The starting SCN of the range of log files to fetch.
	// required = true
	StartLocation string `json:"startLocation,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
	// User credentials. If not provided will use environment credentials for
	// 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
}

// MySQLSnapshot - Provisionable snapshot of a MySQL TimeFlow.
// extends TimeflowSnapshot
type MySQLSnapshot struct {
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *MySQLTimeflowPoint `json:"latestChangePoint,omitempty"`
	// Version of the source database at the time the snapshot was taken.
	InternalVersion *MySQLVersion `json:"internalVersion,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *MySQLSnapshotRuntime `json:"runtime,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *MySQLTimeflowPoint `json:"firstChangePoint,omitempty"`
}

// StorageTest - Test the performance of storage devices.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type StorageTest struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// The results assigned to various tests.
	TestResults []*StorageTestResult `json:"testResults,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The state of the test.
	// enum = [WAITING RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// The parameters used to execute the test.
	Parameters *StorageTestParameters `json:"parameters,omitempty"`
}

// UnixHostEnvironment - The representation of a unix host environment object.
// extends HostEnvironment
type UnixHostEnvironment struct {
	// A reference to the primary user for this environment.
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	// format = objectReference
	PrimaryUser string `json:"primaryUser,omitempty"`
	// The reference to the associated host.
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Parameters for an environment with SAP ASE instances.
	// create = optional
	// update = optional
	AseHostEnvironmentParameters *ASEHostEnvironmentParameters `json:"aseHostEnvironmentParameters,omitempty"`
}

// AppDataContainerRuntime - Runtime properties of an AppData container.
// extends DBContainerRuntime
type AppDataContainerRuntime struct {
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SecurityConfig - System wide security configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SecurityConfig struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Banner displayed prior to login.
	// update = optional
	Banner string `json:"banner,omitempty"`
	// Whether or not CORS is enabled. Changing this value requires a stack
	// restart for it to take effect.
	// update = optional
	IsCORSEnabled *bool `json:"isCORSEnabled,omitempty"`
	// Allowed origin domains for CORS. Should be a comma separated list. Use
	// * for all domains. Defaults to none. Changing this value requires a
	// stack restart for it to take effect.
	// update = optional
	AllowedCORSOrigins string `json:"allowedCORSOrigins,omitempty"`
}

// DxFsIoQueueOpsDatapointStream - A stream of datapoints from a DxFS_IO_QUEUE_OPS analytics slice.
// extends DatapointStream
type DxFsIoQueueOpsDatapointStream struct {
	// Priority of the I/O.
	// enum = [sync cache/agg asyncw asyncr resilver scrub ddt_prefetch]
	Priority string `json:"priority,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write claim free ioctl null]
	Op string `json:"op,omitempty"`
}

// BatchSnapshotDeleteParameters - The parameters to use as input to TimeFlow snapshots batch delete
// requests.
// extends TypedObject
type BatchSnapshotDeleteParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// TimeFlow snapshots to delete.
	// minItems = 1
	// required = true
	Snapshots []string `json:"snapshots,omitempty"`
}

// ASEProvisionParameters - The parameters to use as input to provision SAP ASE databases.
// extends ProvisionParameters
type ASEProvisionParameters struct {
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// Set the "trunc log on chkpt" database option.
	// default = true
	// create = required
	TruncateLogOnCheckpoint *bool `json:"truncateLogOnCheckpoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *ASEDBContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *ASEVirtualSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig ASEDBConfig `json:"sourceConfig,omitempty"`
}

// TimeRangeParameters - The parameters to use as input for methods requiring a time range.
// extends TypedObject
type TimeRangeParameters struct {
	// The date at the end of the period.
	// format = date
	// create = required
	EndTime string `json:"endTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The date at the beginning of the period.
	// create = required
	// format = date
	StartTime string `json:"startTime,omitempty"`
}

// MSSqlSnapshot - Provisionable snapshot of a MSSQL TimeFlow.
// extends TimeflowSnapshot
type MSSqlSnapshot struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *MSSqlSnapshotRuntime `json:"runtime,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Internal version of the source database at the time the snapshot was
	// taken.
	InternalVersion int `json:"internalVersion,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *MSSqlTimeflowPoint `json:"latestChangePoint,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *MSSqlTimeflowPoint `json:"firstChangePoint,omitempty"`
	// UUID of the source database backup that was restored for this
	// snapshot.
	BackupSetUUID string `json:"backupSetUUID,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// Alert - An alert describing an event for a given object.
// extends PersistentObject
type Alert struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Event title.
	EventTitle string `json:"eventTitle,omitempty"`
	// Type of target object.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Reference to target object.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Event class.
	Event string `json:"event,omitempty"`
	// Event severity.
	// enum = [INFORMATIONAL WARNING CRITICAL AUDIT]
	EventSeverity string `json:"eventSeverity,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Time at which event occurred.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Event recommended action.
	EventAction string `json:"eventAction,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Name of target object.
	TargetName string `json:"targetName,omitempty"`
	// Event description.
	EventDescription string `json:"eventDescription,omitempty"`
	// Event response.
	EventResponse string `json:"eventResponse,omitempty"`
	// Additional text associated with the event. This text is not localized
	// and is only provided for certain alerts. For example, if an alert is
	// caused by a post script failure, the output of the post script may be
	// included here to assist with debugging the failure.
	EventCommandOutput string `json:"eventCommandOutput,omitempty"`
}

// OraclePlatformParameters - Oracle platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type OraclePlatformParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleCreateTransformationParameters - Represents the parameters of a createTransformation request for an
// Oracle container.
// extends CreateTransformationParameters
type OracleCreateTransformationParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *OracleDatabaseContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
}

// SingletonUpdate - An event indicating an update to a singleton object on the system.
// extends Notification
type SingletonUpdate struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of target object.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
}

// JSDataContainer - A container represents a data template provisioned for a specific
// user.
// extends JSDataLayout
type JSDataContainer struct {
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// Name of the user who locked this data container.
	// update = readonly
	LockUserName string `json:"lockUserName,omitempty"`
	// The data template that this data container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	Template string `json:"template,omitempty"`
	// The state of the data container.
	// enum = [ONLINE OFFLINE INCONSISTENT]
	State string `json:"state,omitempty"`
	// The active branch of the data layout.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	ActiveBranch string `json:"activeBranch,omitempty"`
	// The first JSOperation on this data layout by data time.
	// referenceTo = /delphix-js-operation.json
	// format = objectReference
	FirstOperation string `json:"firstOperation,omitempty"`
	// The reference to the user who locked this data container.
	// format = objectReference
	// referenceTo = /delphix-user.json
	// update = readonly
	LockUserReference string `json:"lockUserReference,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Timestamp of the last update to the application.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The last JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	LastOperation string `json:"lastOperation,omitempty"`
	// The number of operations performed on this data container.
	OperationCount int `json:"operationCount,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Notes for this data layout.
	// maxLength = 4096
	// update = optional
	Notes string `json:"notes,omitempty"`
	// For backward compatibility. The owner of the data container.
	// referenceTo = /delphix-user.json
	// update = optional
	// internal = true
	// format = objectReference
	Owner string `json:"owner,omitempty"`
}

// InterfaceAddress - IP address assigned to a network interface.
// extends TypedObject
type InterfaceAddress struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this address should accept incoming SSH connections.
	// default = true
	// create = optional
	// update = optional
	EnableSSH *bool `json:"enableSSH,omitempty"`
	// True if the API session is established over this address. This
	// property helps a client make informative decisions about which address
	// should not be modified without affecting the session over which it is
	// connected.
	SessionInUse *bool `json:"sessionInUse,omitempty"`
	// The type of address (STATIC or DHCP).
	// default = STATIC
	// create = required
	// update = optional
	// enum = [STATIC DHCP]
	AddressType string `json:"addressType,omitempty"`
	// The address in Classless Inter-Domain Routing (CIDR) notation.
	// update = optional
	// format = cidrAddress
	// create = optional
	Address string `json:"address,omitempty"`
	// The state of the address.
	// enum = [OK TENTATIVE DUPLICATE INACCESSIBLE]
	State string `json:"state,omitempty"`
}

// ASENewBackupSyncParameters - The parameters to use as input to sync a SAP ASE database by taking a
// new full backup.
// extends ASESyncParameters
type ASENewBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleRedoLogFileSpecification - Describes an Oracle redo log file.
// extends OracleFileSpecification
type OracleRedoLogFileSpecification struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// The size of the log file in MB.
	// create = optional
	// update = optional
	// minimum = 100
	// default = 100
	Size int `json:"size,omitempty"`
}

// PreProvisioningRuntime - Runtime properties for pre-provisioning of a MSSQL database container.
// extends TypedObject
type PreProvisioningRuntime struct {
	// Indicates the current state of pre-provisioning for the database.
	// enum = [ACTIVE INACTIVE FAULTED UNKNOWN]
	PreProvisioningState string `json:"preProvisioningState,omitempty"`
	// Timestamp of the last update to the status.
	LastUpdateTimestamp string `json:"lastUpdateTimestamp,omitempty"`
	// The status of the pre-provisioning run.
	Status string `json:"status,omitempty"`
	// Response taken based on the status of the pre-provisioning run.
	Response string `json:"response,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User action required to resolve any error that the pre-provisioning
	// run encountered.
	PendingAction string `json:"pendingAction,omitempty"`
}

// PolicyApplyTargetParameters - Specifies the target to which a policy is applied.
// extends TypedObject
type PolicyApplyTargetParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object reference of the target.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// Object references of the targets.
	Targets []string `json:"targets,omitempty"`
}

// JSUserUsageData - The space usage information for a Jet Stream user.
// extends TypedObject
type JSUserUsageData struct {
	// The amount of space referenced by the data containers owned by this
	// user.
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
	// The number of containers owned by this user.
	NumContainers int `json:"numContainers,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The user.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
}

// ASESpecificBackupSyncParameters - The parameters to use as input to sync a SAP ASE database using a
// specific existing backup.
// extends ASESyncParameters
type ASESpecificBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The location of the full backup of the source database to restore
	// from. The backup should be present in the shared backup location for
	// the source database.
	// create = required
	// uniqueItems = true
	// minItems = 1
	BackupFiles []string `json:"backupFiles,omitempty"`
}

// MSSqlExistingMostRecentBackupSyncParameters - The parameters to use as input to sync MSSQL databases using an
// existing most recent full or differential backup.
// extends MSSqlExistingBackupSyncParameters
type MSSqlExistingMostRecentBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ScrubStatus - The status of a scrub of the storage in the system.
// extends TypedObject
// cliVisibility = [SYSTEM]
type ScrubStatus struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Scrub state.
	// enum = [NONE ACTIVE COMPLETED CANCELED]
	State string `json:"state,omitempty"`
	// Amount of data scrubbed, in bytes.
	// units = B
	// base = 1024
	Completed float64 `json:"completed,omitempty"`
	// Total amount of data to scrub (including completed), in bytes.
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
	// Number of errors encountered during scrub.
	Errors float64 `json:"errors,omitempty"`
	// Time scrub was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Time scrub ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
}

// KerberosCredential - Kerberos based security credential.
// extends Credential
type KerberosCredential struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// TargetFilter - An alert filter that specifies which targets to match against.
// extends AlertFilter
type TargetFilter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// List of object references. Only alerts related to one of the targets
	// or its children are included.
	// create = optional
	// update = optional
	// minItems = 1
	Targets []string `json:"targets,omitempty"`
}

// AppDataPlatformParameters - AppData platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type AppDataPlatformParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// create = required
	Payload *Json `json:"payload,omitempty"`
}

// PgSQLDBContainerRuntime - Runtime properties of a PostgreSQL database container.
// extends DBContainerRuntime
type PgSQLDBContainerRuntime struct {
	// The ID of the WAL segment that was last restored in this container (if
	// applicable).
	LastRestoredWALSegment string `json:"lastRestoredWALSegment,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
}

// TimeflowPointSnapshot - TimeFlow point based on a snapshot reference.
// extends TimeflowPointParameters
type TimeflowPointSnapshot struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Reference to the snapshot.
	// referenceTo = /delphix-timeflow-snapshot.json
	// required = true
	// format = objectReference
	Snapshot string `json:"snapshot,omitempty"`
}

// MSSqlAvailabilityGroupListener - The representation of a SQL Server Availability Group Listener.
// extends MSSqlBaseClusterListener
type MSSqlAvailabilityGroupListener struct {
	// The address of the listener.
	Address string `json:"address,omitempty"`
	// The name of the listener.
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The port for the listener.
	Port int `json:"port,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataLinkedStagedSource - An AppData linked source with a staging source.
// extends AppDataLinkedSource
type AppDataLinkedStagedSource struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-appdata-source-config.json
	// create = required
	// update = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// The toolkit associated with this source.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntime `json:"runtime,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// create = required
	// update = optional
	Parameters *Json `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// The base mount point for the NFS mount on the staging environment.
	// update = optional
	// maxLength = 256
	// create = required
	StagingMountBase string `json:"stagingMountBase,omitempty"`
	// The environment used as an intermediate stage to pull data into
	// Delphix.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	// update = optional
	StagingEnvironment string `json:"stagingEnvironment,omitempty"`
	// The environment user used to access the staging environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = required
	// update = optional
	StagingEnvironmentUser string `json:"stagingEnvironmentUser,omitempty"`
}

// SourceConfigConnectivity - Mechanism to test JDBC connectivity to source configs.
// extends TypedObject
type SourceConfigConnectivity struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Database username.
	// required = true
	Username string `json:"username,omitempty"`
	// Database password.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
}

// XppLastRunStatus - Status of the last cross-platform provision of the container.
// extends ChecklistItem
type XppLastRunStatus struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Last cross-platform provision job run on the container.
	Job *Job `json:"job,omitempty"`
	// Error message associated with the last run, if any.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// ASESIConfig - A SAP ASE single instance database config.
// extends ASEDBConfig
type ASESIConfig struct {
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	// maxLength = 30
	DatabaseName string `json:"databaseName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
}

// Permission - Describes a permission to perform an operation on an object in the
// Delphix Engine.
// extends ReadonlyNamedUserObject
type Permission struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Name of the action governed by the permission.
	ActionType string `json:"actionType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// TimeConfig - Get and set the current time configuration.
// extends TypedObject
// cliVisibility = [SYSTEM DOMAIN]
type TimeConfig struct {
	// System time zone offset as a String. For instance 'UTC -5:00'.
	SystemTimeZoneOffsetString string `json:"systemTimeZoneOffsetString,omitempty"`
	// NTP configuration.
	// update = optional
	NtpConfig *NTPConfig `json:"ntpConfig,omitempty"`
	// Current system time. This value can only be set if NTP is disabled.
	// The management service is automatically restarted if the time is
	// changed.
	// format = date
	// update = optional
	CurrentTime string `json:"currentTime,omitempty"`
	// Default time zone for system wide policies and schedules. The
	// management service is automatically restarted if the timezone is
	// changed.
	// update = optional
	// default = Etc/UTC
	SystemTimeZone string `json:"systemTimeZone,omitempty"`
	// The difference, in minutes, between UTC and local time. For example,
	// if your time zone is UTC -5:00 (Eastern Standard Time), 300 will be
	// returned. Daylight saving time prevents this value from being a
	// constant even for a given locale.
	SystemTimeZoneOffset int `json:"systemTimeZoneOffset,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// MySQLAttachData - Represents the MySQL parameters of an attach request.
// extends AttachData
type MySQLAttachData struct {
	// OS user on the staging host to use for attaching.
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	// format = objectReference
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// The database username.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The credentials for the database user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The MySQL installation on the staging environment to use for validated
	// sync.
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Reference to the configuration for the source. Must be an existing
	// linked source config.
	// referenceTo = /delphix-mysql-server-config.json
	// required = true
	// format = objectReference
	Config string `json:"config,omitempty"`
	// The port on which the MySQL staging instance will listen.
	// maximum = 65535
	// required = true
	// minimum = 1
	StagingPort int `json:"stagingPort,omitempty"`
}

// AppDataLinkedDirectSource - An AppData linked source directly replicated into the Delphix Engine.
// extends AppDataLinkedSource
type AppDataLinkedDirectSource struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// List of subdirectories in the source to exclude when syncing data.
	// These paths are relative to the root of the source directory.
	// create = required
	// update = optional
	Excludes []string `json:"excludes,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// The toolkit associated with this source.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// create = required
	// update = optional
	Parameters *Json `json:"parameters,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// List of symlinks in the source to follow when syncing data. These
	// paths are relative to the root of the source directory. All other
	// symlinks are preserved.
	// create = required
	// update = optional
	FollowSymlinks []string `json:"followSymlinks,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntime `json:"runtime,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
}

// AppDataSourceConnectionInfo - Contains information that can be used to connect to the application
// source.
// extends SourceConnectionInfo
type AppDataSourceConnectionInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The path where the application data is located on the host.
	Path string `json:"path,omitempty"`
}

// OracleSISourceConnectionInfo - Contains information that can be used to connect to a single instance
// Oracle source.
// extends OracleSourceConnectionInfo
type OracleSISourceConnectionInfo struct {
	// The database name.
	DatabaseName string `json:"databaseName,omitempty"`
	// The Oracle installation home.
	OracleHome string `json:"oracleHome,omitempty"`
	// The JDBC strings used to connect to the source.
	JdbcStrings []string `json:"jdbcStrings,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// MSSqlAvailabilityGroupDBConfig - Database for a SQL Server Availability Group.
// extends MSSqlDBConfig
type MSSqlAvailabilityGroupDBConfig struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Specifies the current recovery model of the source database.
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	// create = optional
	// update = readonly
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The name of the database.
	// update = optional
	// maxLength = 128
	// create = required
	DatabaseName string `json:"databaseName,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
}

// PgSQLInstall - A PostgreSQL installation.
// extends SourceRepository
type PgSQLInstall struct {
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Variant of the repository.
	// enum = [PostgreSQL EnterpriseDB]
	Variant string `json:"variant,omitempty"`
	// 32 or 64 bit installation.
	// enum = [32 64]
	Bits int `json:"bits,omitempty"`
	// Size of the WAL segments (in bytes) generated by PostgreSQL binaries.
	// units = B
	// enum = [1.048576e+06 2.097152e+06 4.194304e+06 8.388608e+06 1.6777216e+07 3.3554432e+07 6.7108864e+07]
	// base = 1024
	SegmentSize int `json:"segmentSize,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Directory path where the installation is located.
	// create = required
	// maxLength = 1024
	InstallationPath string `json:"installationPath,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Version of the repository.
	// format = pgsqlVersion
	Version string `json:"version,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Reference to the environment containing this repository.
	// referenceTo = /delphix-source-environment.json
	// create = required
	// format = objectReference
	Environment string `json:"environment,omitempty"`
	// Flag indicating whether the installation was discovered or manually
	// entered.
	Discovered *bool `json:"discovered,omitempty"`
}

// LocaleSettings - Global locale settings.
// extends UserObject
// cliVisibility = [SYSTEM]
type LocaleSettings struct {
	// System default locale as an IETF BCP 47 language tag, defaults to
	// 'en-US'.
	// update = optional
	// format = locale
	// default = en-US
	// create = required
	Locale string `json:"locale,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
}

// OracleLinkData - Represents parameters to link non-pluggable Oracle databases.
// extends OracleBaseLinkData
type OracleLinkData struct {
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// True if SnapSync data from the source should be retrieved through an
	// encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Number of parallel channels to use.
	// maximum = 32
	// default = 2
	// create = optional
	// minimum = 1
	RmanChannels int `json:"rmanChannels,omitempty"`
	// If true, NOLOGGING operations on this container are treated as faults
	// and cannot be resolved manually. Otherwise, these operations are
	// ignored.
	// default = true
	// create = optional
	DiagnoseNoLoggingFaults *bool `json:"diagnoseNoLoggingFaults,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// Non-SYS database credentials to access this database.
	// create = optional
	NonSysCredentials *PasswordCredential `json:"nonSysCredentials,omitempty"`
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// Skip check that tests if there is enough space available to store the
	// database in the Delphix Engine. The Delphix Engine estimates how much
	// space a database will occupy after compression and prevents SnapSync
	// if insufficient space is available. This safeguard can be overridden
	// using this option. This may be useful when linking highly compressible
	// databases.
	// default = false
	// create = optional
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
	// True if SnapSync data from the source should be compressed over the
	// network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially over
	// slow network.
	// default = true
	// create = optional
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet int `json:"filesPerSet,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession to
	// reduce the number of logs required to provision the snapshot. This may
	// significantly reduce the time necessary to provision from a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// Non-SYS database user to access this database.
	// maxLength = 30
	// create = optional
	NonSysUser string `json:"nonSysUser,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *OracleSourcingPolicy `json:"sourcingPolicy,omitempty"`
	// If true, pre-provisioning will be performed after every sync.
	// default = false
	// create = optional
	PreProvisioningEnabled *bool `json:"preProvisioningEnabled,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// minimum = 1
	// maximum = 16
	// create = optional
	// default = 1
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
}

// MSSqlSIConfig - Configuration information for a single instance MSSQL Source.
// extends MSSqlDBConfig
type MSSqlSIConfig struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Specifies the current recovery model of the source database.
	// update = readonly
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	// create = optional
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// update = optional
	// default = true
	// create = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// maxLength = 128
	DatabaseName string `json:"databaseName,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 256
	User string `json:"user,omitempty"`
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
}

// OracleVirtualIP - The parameters used for virtual IP operations.
// extends TypedObject
type OracleVirtualIP struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The virtual IP address.
	// format = ipv4Address
	// required = true
	Ip string `json:"ip,omitempty"`
	// The name of the domain where the cluster is residing.
	// required = true
	DomainName string `json:"domainName,omitempty"`
	// A boolean indicating whether this VIP was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// JSDataContainerRestoreParameters - The parameters used to restore a data container.
// extends TypedObject
type JSDataContainerRestoreParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Jet Stream data timeline point to restore from.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
}

// PgSQLHBAEntry - An entry in the PostgreSQL host-based authentication file
// (pg_hba.conf).
// extends TypedObject
type PgSQLHBAEntry struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The authentication method to use when connecting via this entry.
	// enum = [trust reject md5 password gss sspi krb5 ident peer ldap radius cert pam]
	// create = required
	AuthMethod string `json:"authMethod,omitempty"`
	// Options for the authentication method.
	// create = optional
	AuthOptions string `json:"authOptions,omitempty"`
	// The connection type of this entry.
	// create = required
	// enum = [local host hostssl hostnossl]
	EntryType string `json:"entryType,omitempty"`
	// The database name this entry matches.
	// default = all
	// create = required
	// maxLength = 63
	Database string `json:"database,omitempty"`
	// The database username this entry matches.
	// default = all
	// create = required
	// maxLength = 63
	User string `json:"user,omitempty"`
	// The client machine address that this entry matches.
	// create = optional
	Address string `json:"address,omitempty"`
}

// SourceStopParameters - The parameters to use as input to stop a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceStopParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowRangeParameters - The parameters to use as input to fetch TimeFlow ranges.
// extends TypedObject
type TimeflowRangeParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The starting TimeFlow point of the time period to search for TimeFlow
	// ranges.
	StartPoint TimeflowPoint `json:"startPoint,omitempty"`
	// The ending TimeFlow point of the time period to search for TimeFlow
	// ranges.
	EndPoint TimeflowPoint `json:"endPoint,omitempty"`
}

// OracleSIConfig - The representation of a single-instance Oracle DB configuration.
// extends OracleDBConfig
type OracleSIConfig struct {
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The Oracle instance.
	// create = required
	// update = optional
	Instance *OracleInstance `json:"instance,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-oracle-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The password of the database user. This must be a PasswordCredential
	// instance.
	// update = optional
	Credentials *PasswordCredential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 8
	DatabaseName string `json:"databaseName,omitempty"`
	// The container type of this database.
	// create = readonly
	// update = readonly
	// enum = [UNKNOWN ROOT_CDB NON_CDB AUX_CDB]
	CdbType string `json:"cdbType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 30
	User string `json:"user,omitempty"`
	// The unique name.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 30
	UniqueName string `json:"uniqueName,omitempty"`
	// The username of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	// maxLength = 30
	NonSysUser string `json:"nonSysUser,omitempty"`
	// The password of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	NonSysCredentials string `json:"nonSysCredentials,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleService `json:"services,omitempty"`
}

// SnapshotPolicy - This policy creates snapshots of a container with externally managed
// sources (virtual databases) according to a schedule.
// extends SchedulePolicy
type SnapshotPolicy struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// List of Schedule objects representing when the policy will execute.
	// create = required
	// update = optional
	ScheduleList []*Schedule `json:"scheduleList,omitempty"`
	// The timezone of this policy. If not specified, defaults to the Delphix
	// Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZone `json:"timezone,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// True if this is the default policy created when the system is setup.
	// Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See the
	// effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
}

// OracleProvisionParameters - The parameters to use as input to provision Oracle (non-multitenant)
// databases.
// extends OracleBaseProvisionParameters
type OracleProvisionParameters struct {
	// The source that describes an external database instance.
	// required = true
	Source *OracleVirtualSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The name of the privileged user to run the provision operation as.
	// create = optional
	Username string `json:"username,omitempty"`
	// Flag indicating whether to open the database after provision.
	// create = optional
	// default = true
	OpenResetlogs *bool `json:"openResetlogs,omitempty"`
	// Flag indicating whether the virtual database is provisioned as a
	// physical standby database.
	// create = optional
	// default = false
	PhysicalStandby *bool `json:"physicalStandby,omitempty"`
	// Flag indicating whether to generate a new DBID for the provisioned
	// database.
	// create = optional
	// default = false
	NewDBID *bool `json:"newDBID,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *OracleDatabaseContainer `json:"container,omitempty"`
}

// PgSQLDatabaseContainer - A PostgreSQL Database Container.
// extends DatabaseContainer
type PgSQLDatabaseContainer struct {
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// referenceTo = /delphix-timeflow.json
	// format = objectReference
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Whether to enable high performance mode.
	// update = readonly
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// referenceTo = /delphix-timeflow.json
	// format = objectReference
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Runtime properties of this container.
	Runtime *PgSQLDBContainerRuntime `json:"runtime,omitempty"`
}

// IntegerLessThanConstraint - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerLessThanConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be less than than this value.
	// create = required
	LessThan int `json:"lessThan,omitempty"`
}

// MySQLCompatibilityCriteria - The compatibility criteria to use for selecting compatible MySQL
// repositories.
// extends CompatibilityCriteria
type MySQLCompatibilityCriteria struct {
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Version of the MySQL installation.
	Version *MySQLVersion `json:"version,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
}

// OracleExportParameters - The parameters to use as input to export Oracle databases.
// extends DbExportParameters
type OracleExportParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayout `json:"filesystemLayout,omitempty"`
	// Open the database after recovery. This can have a true value only if
	// 'recoverDatabase' is true.
	// create = optional
	// default = true
	OpenDatabase *bool `json:"openDatabase,omitempty"`
	// Number of files to stream in parallel across the network.
	// minimum = 1
	// maximum = 64
	// create = optional
	// default = 3
	FileParallelism int `json:"fileParallelism,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// DSP options for export.
	// create = optional
	DspOptions *DSPOptions `json:"dspOptions,omitempty"`
}

// SystemKeyCredential - The system public key based security credential.
// extends PublicKeyCredential
type SystemKeyCredential struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// UpgradeCompatibilityParameters - The criteria necessary to select valid repositories for upgrading.
// extends CompatibleRepositoriesParameters
type UpgradeCompatibilityParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
}

// OracleDeleteParameters - The parameters passed in for an Oracle database delete operation.
// extends DeleteParameters
type OracleDeleteParameters struct {
	// The security credential of the privileged user to run the delete
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
	// The name of the privileged user to run the delete operation as.
	Username string `json:"username,omitempty"`
}

// AppDataTimeflow - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends AppDataBaseTimeflow
type AppDataTimeflow struct {
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *AppDataTimeflowPoint `json:"parentPoint,omitempty"`
}

// TimeflowSnapshotDayRange - Count of TimeFlow snapshots aggregated by day.
// extends TypedObject
type TimeflowSnapshotDayRange struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Number of TimeFlow snapshots on that day.
	Count float64 `json:"count,omitempty"`
	// Date for which TimeFlow snapshots have been aggregated.
	Date string `json:"date,omitempty"`
	// Start of day of this range in the time zone used for computation.
	// format = date
	StartOfDay string `json:"startOfDay,omitempty"`
	// End of day of this range in the time zone used for computation.
	// format = date
	EndOfDay string `json:"endOfDay,omitempty"`
}

// MSSqlSourceRuntime - Runtime (non-persistent) properties of a MSSQL source.
// extends SourceRuntime
type MSSqlSourceRuntime struct {
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// base = 1024
	// units = B
	DatabaseSize float64 `json:"databaseSize,omitempty"`
}

// FeatureFlagParameters - Feature Flags for the Delphix Engine.
// extends TypedObject
type FeatureFlagParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Name of the feature flag.
	// required = true
	Name string `json:"name,omitempty"`
}

// CompatibilityCriteria - The compatibility criteria to use for selecting compatible
// repositories. Parameters with a value of null are not considered when
// selecting compatible repositories.
// extends TypedObject
type CompatibilityCriteria struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
}

// OracleVirtualCdbProvisionParameters - The parameters to use as input to provision Oracle virtual container
// databases.
// extends TypedObject
type OracleVirtualCdbProvisionParameters struct {
	// The new container for the created database.
	// required = true
	Container *OracleDatabaseContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *OracleVirtualCdbSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig OracleDBConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ReplicationSourceState - State of a replication spec.
// extends UserObject
type ReplicationSourceState struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// The last serialization point sent. This can be null prior to the first
	// replication run.
	// format = objectReference
	// referenceTo = /delphix-serialization-point.json
	LastPoint string `json:"lastPoint,omitempty"`
	// The active serialization point, currently being sent or about to be
	// sent to replication targets.
	// format = objectReference
	// referenceTo = /delphix-serialization-point.json
	ActivePoint string `json:"activePoint,omitempty"`
	// A reference to the replication specification responsible for the
	// current state.
	// format = objectReference
	// referenceTo = /delphix-replicationspec.json
	Spec string `json:"spec,omitempty"`
}

// OraclePDBSourceRuntime - Runtime (non-persistent) properties of an Oracle PDB source.
// extends OracleBaseSourceRuntime
type OraclePDBSourceRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// The current role of the database.
	// default = UNKNOWN
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstance `json:"activeInstances,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Operating mode of the database.
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	// default = UNKNOWN
	DatabaseMode string `json:"databaseMode,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSection `json:"databaseStats,omitempty"`
}

// MySQLTimeflowPoint - A unique point within a MySQL database TimeFlow.
// extends TimeflowPoint
type MySQLTimeflowPoint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Reference to TimeFlow containing this point.
	// format = objectReference
	// referenceTo = /delphix-mysql-timeflow.json
	// required = true
	Timeflow string `json:"timeflow,omitempty"`
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
}

// OracleSystemDatafileSpecification - Describes an Oracle datafile in the SYSTEM tablespace.
// extends OracleDatafileTempfileSpecification
type OracleSystemDatafileSpecification struct {
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space that
	// Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize int `json:"maxSize,omitempty"`
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the size
	// of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement int `json:"autoExtendIncrement,omitempty"`
	// The size of the file in MB.
	// minimum = 700
	// default = 700
	// create = optional
	// update = optional
	Size int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
}

// DSPOptions - Options commonly used by apps that use DSP.
// extends TypedObject
type DSPOptions struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Compress the data stream over the network.
	// create = optional
	// default = false
	Compression *bool `json:"compression,omitempty"`
	// Encrypt the data stream over the network.
	// create = optional
	// default = false
	Encryption *bool `json:"encryption,omitempty"`
	// Bandwidth limit (MB/s) for network traffic. A value of 0 means no
	// limit.
	// create = optional
	// default = 0
	// minimum = 0
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// Total number of transport connections to use.
	// create = optional
	// default = 1
	// minimum = 1
	// maximum = 16
	NumConnections int `json:"numConnections,omitempty"`
}

// OracleSnapshotRuntime - Runtime (non-persistent) properties of an Oracle TimeFlow snapshot.
// extends SnapshotRuntime
type OracleSnapshotRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// List of missing log files for this snapshot, if any.
	MissingLogs []*OracleLog `json:"missingLogs,omitempty"`
}

// TimeflowPointLocation - TimeFlow point based on a database-specific identifier (SCN, LSN,
// etc).
// extends TimeflowPointParameters
type TimeflowPointLocation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow location.
	// required = true
	Location string `json:"location,omitempty"`
	// Reference to TimeFlow containing this location.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
}

// DetachSourceParameters - The parameters passed in for a database detach source operation.
// extends TypedObject
type DetachSourceParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Reference to the source to be removed. This must be a linked source
	// attached to the target database.
	// format = objectReference
	// referenceTo = /delphix-source.json
	// required = true
	Source string `json:"source,omitempty"`
}

// DxFsOpsDatapointStream - A stream of datapoints from a DxFS_OPS analytics slice.
// extends DatapointStream
type DxFsOpsDatapointStream struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
}

// SourceUpgradeParameters - The parameters to use as input to upgrade a source.
// extends TypedObject
type SourceUpgradeParameters struct {
	// The source config that the source database upgrades to.
	// required = true
	SourceConfig SourceConfig `json:"sourceConfig,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// PgSQLSnapshot - Provisionable snapshot of a PostgreSQL TimeFlow.
// extends TimeflowSnapshot
type PgSQLSnapshot struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *PgSQLTimeflowPoint `json:"latestChangePoint,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *PgSQLTimeflowPoint `json:"firstChangePoint,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *SnapshotRuntime `json:"runtime,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
}

// XPPCompatibilityParameters - The criteria necessary to select valid repositories for cross-platform
// provisioning.
// extends CompatibleRepositoriesParameters
type XPPCompatibilityParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The TimeFlow point to use as a source of compatibility information.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// Fault - A representation of a fault, with associated user object.
// extends PersistentObject
type Fault struct {
	// The user-visible Delphix object that is faulted.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// The name of the faulted object at the time the fault was diagnosed.
	TargetName string `json:"targetName,omitempty"`
	// Summary of the fault.
	Title string `json:"title,omitempty"`
	// A unique dot delimited identifier associated with the fault.
	BundleID string `json:"bundleID,omitempty"`
	// A suggested user action.
	Action string `json:"action,omitempty"`
	// The status of the fault. This can be ACTIVE, RESOLVED or IGNORED.
	// enum = [ACTIVE RESOLVED IGNORED]
	Status string `json:"status,omitempty"`
	// The user-visible Delphix object that is faulted.
	// format = type
	TargetObjectType string `json:"targetObjectType,omitempty"`
	// Full description of the fault.
	Description string `json:"description,omitempty"`
	// The date when the fault was diagnosed.
	// format = date
	DateDiagnosed string `json:"dateDiagnosed,omitempty"`
	// A comment that describes the fault resolution.
	ResolutionComments string `json:"resolutionComments,omitempty"`
	// The automated response taken by the system.
	Response string `json:"response,omitempty"`
	// The severity of the fault event. This can either be CRITICAL or
	// WARNING.
	// enum = [CRITICAL WARNING]
	Severity string `json:"severity,omitempty"`
	// The date when the fault was resolved.
	// format = date
	DateResolved string `json:"dateResolved,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// PgSQLDBClusterConfigConnectivity - Mechanism to test JDBC connectivity to PostgreSQL DB clusters.
// extends SourceConfigConnectivity
type PgSQLDBClusterConfigConnectivity struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Database username.
	// required = true
	Username string `json:"username,omitempty"`
	// Database password.
	// format = password
	// required = true
	Password string `json:"password,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// maxLength = 256
	// create = optional
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
}

// SupportBundleGenerateParameters - Parameters to be used when generating a support bundle.
// extends BaseSupportBundleParameters
type SupportBundleGenerateParameters struct {
	// Type of support bundle to generate. Reserved for Delphix support use.
	// default = ALL
	// enum = [PHONEHOME MDS OS CORE LOG DROPBOX STORAGE_TEST MASKING ALL]
	BundleType string `json:"bundleType,omitempty"`
	// Whether or not to include the analytics data in the support bundle
	// which is generated. Including analytics data may significantly
	// increase the support bundle size and upload time, but enables analysis
	// of performance characteristics of the Delphix Engine.
	// default = false
	IncludeAnalyticsData *bool `json:"includeAnalyticsData,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleSourceRuntime - Runtime (non-persistent) properties of an Oracle source.
// extends OracleBaseSourceRuntime
type OracleSourceRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstance `json:"activeInstances,omitempty"`
	// True for a RAC source database.
	RacEnabled *bool `json:"racEnabled,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// Operating mode of the database.
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	// default = UNKNOWN
	DatabaseMode string `json:"databaseMode,omitempty"`
	// True if block change tracking is enabled.
	BctEnabled *bool `json:"bctEnabled,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSection `json:"databaseStats,omitempty"`
	// The current role of the database.
	// default = UNKNOWN
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// True if the database has Oracle Direct NFS client enabled.
	DnfsEnabled *bool `json:"dnfsEnabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if the database is running in ARCHIVELOG mode.
	ArchivelogEnabled *bool `json:"archivelogEnabled,omitempty"`
}

// AppDataCreateTransformationParameters - Represents the parameters of a createTransformation request for an
// AppData container.
// extends CreateTransformationParameters
type AppDataCreateTransformationParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *AppDataContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	Repository string `json:"repository,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// create = required
	Payload *Json `json:"payload,omitempty"`
}

// OracleLinkedSource - A linked Oracle source.
// extends OracleSource
type OracleLinkedSource struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// True if SnapSync data from the source should be retrieved through an
	// encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// create = optional
	// update = optional
	// default = false
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// True if SnapSync data from the source should be compressed over the
	// network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially over
	// slow network.
	// create = optional
	// update = optional
	// default = true
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// Defines whether backup level is enabled.
	// create = optional
	// update = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	// update = optional
	RmanChannels int `json:"rmanChannels,omitempty"`
	// User-specified operation hooks for this source.
	// update = optional
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-base-db-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	// update = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// create = optional
	// update = optional
	// default = 1
	// minimum = 1
	// maximum = 16
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// default = 0
	// create = optional
	// update = optional
	// minimum = 0
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	// update = optional
	FilesPerSet int `json:"filesPerSet,omitempty"`
}

// SnapshotSpaceParameters - Input to the operation to determine how much space is used by a set of
// snapshots.
// extends TypedObject
type SnapshotSpaceParameters struct {
	// FilesystemObjects to query, in canonical object reference form.
	// required = true
	ObjectReferences []string `json:"objectReferences,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLStagingSource - A MySQL staging source used for validated sync.
// extends MySQLSource
type MySQLStagingSource struct {
	// A user-provided script to run prior to taking a snapshot.
	// create = optional
	// update = optional
	PreScript string `json:"preScript,omitempty"`
	// A user-provided script to run after taking a snapshot.
	// create = optional
	// update = optional
	PostScript string `json:"postScript,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntime `json:"runtime,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// The base mount point for the NFS mounts on the staging host.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
}

// OracleTempfileSpecification - Describes an Oracle temporary file.
// extends OracleDatafileTempfileSpecification
type OracleTempfileSpecification struct {
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// update = optional
	// default = true
	// create = optional
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the size
	// of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement int `json:"autoExtendIncrement,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space that
	// Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// create = optional
	// update = optional
	// minimum = 300
	// default = 300
	Size int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DeleteParameters - The parameters to use as input to delete requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type DeleteParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
}

// JSDataTemplate - A data template is a collection of data sources and configuration
// representing a data layout that can be provisioned to Jet Stream
// users.
// extends JSDataLayout
type JSDataTemplate struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Notes for this data layout.
	// update = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// The last JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	LastOperation string `json:"lastOperation,omitempty"`
	// The first JSOperation on this data layout by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	FirstOperation string `json:"firstOperation,omitempty"`
	// A client should consider warning the user before performing an
	// operation which may take a long time, if this is true.
	// default = true
	// create = optional
	// update = optional
	ConfirmTimeConsumingOperations *bool `json:"confirmTimeConsumingOperations,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// update = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The active branch of the data layout.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	ActiveBranch string `json:"activeBranch,omitempty"`
	// Timestamp of the last update to the application.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// OracleActiveInstance - Active instance information for an Oracle database.
// extends TypedObject
type OracleActiveInstance struct {
	// The name of the Oracle instance.
	InstanceName string `json:"instanceName,omitempty"`
	// The name of the host the instance runs on.
	HostName string `json:"hostName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The number of the Oracle instance.
	InstanceNumber int `json:"instanceNumber,omitempty"`
}

// Transformation - A data platform agnostic transformation object.
// extends UserObject
type Transformation struct {
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Operations to perform when this transformation is applied.
	Operations []SourceOperation `json:"operations,omitempty"`
	// Platform-specific parameters that are stored on a transformation.
	PlatformParams BasePlatformParameters `json:"platformParams,omitempty"`
	// A reference to the container which is a transformed version of the
	// parent container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// TargetOwnerFilter - An alert filter which matches when an alert's target is owned by one
// of the specified users.
// extends AlertFilter
type TargetOwnerFilter struct {
	// Target owners to match against.
	// update = optional
	// minItems = 1
	// create = required
	Owners []string `json:"owners,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// LoginRequest - Represents a Delphix user authentication request.
// extends TypedObject
type LoginRequest struct {
	// The username of the user to authenticate.
	// required = true
	Username string `json:"username,omitempty"`
	// The password of the user to authenticate.
	// required = true
	// format = password
	Password string `json:"password,omitempty"`
	// Whether to keep session alive for all requests or only via
	// 'KeepSessionAlive' request headers. Defaults to ALL_REQUESTS if
	// omitted.
	// enum = [ALL_REQUESTS KEEP_ALIVE_HEADER_ONLY]
	// default = ALL_REQUESTS
	KeepAliveMode string `json:"keepAliveMode,omitempty"`
	// The authentication domain.
	// enum = [DOMAIN SYSTEM]
	Target string `json:"target,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RunPowerShellOnSourceOperation - A user-specifiable operation that runs a PowerShell command on the
// target host.
// extends SourceOperation
type RunPowerShellOnSourceOperation struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A name for the source operation.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The PowerShell command to execute on the target host.
	// create = required
	// update = optional
	Command string `json:"command,omitempty"`
}

// MSSqlSourceConnectionInfo - Contains information that can be used to connect to a SQL server
// source.
// extends SourceConnectionInfo
type MSSqlSourceConnectionInfo struct {
	// The name of the database.
	DatabaseName string `json:"databaseName,omitempty"`
	// The JDBC string used to connect to the SQL server instance.
	InstanceJDBCString string `json:"instanceJDBCString,omitempty"`
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
	// The port number used to connect to the source.
	Port int `json:"port,omitempty"`
	// The name of the instance.
	InstanceName string `json:"instanceName,omitempty"`
}

// OracleMultitenantProvisionParameters - The parameters to use as input to provision Oracle multitenant
// databases.
// extends OracleBaseProvisionParameters
type OracleMultitenantProvisionParameters struct {
	// The new container for the provisioned database.
	// required = true
	Container *OracleDatabaseContainer `json:"container,omitempty"`
	// The pluggable database source that describes an external database
	// instance.
	// required = true
	Source *OracleVirtualPdbSource `json:"source,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// The name of the privileged user to run the provision operation as.
	// create = optional
	Username string `json:"username,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	// create = optional
	Credential Credential `json:"credential,omitempty"`
	// The new container for the created dataset.
	// create = optional
	VirtualCdb *OracleVirtualCdbProvisionParameters `json:"virtualCdb,omitempty"`
	// The pluggable database source config including dynamically discovered
	// attributes of the source.
	// required = true
	SourceConfig *OraclePDBConfig `json:"sourceConfig,omitempty"`
}

// AppDataFilesystemLayout - A filesystem layout that matches the filesystem of a Delphix TimeFlow.
// extends FilesystemLayout
type AppDataFilesystemLayout struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The base directory to use for the exported database.
	// required = true
	TargetDirectory string `json:"targetDirectory,omitempty"`
}

// CpuUtilDatapointStream - A stream of datapoints from a CPU_UTIL analytics slice.
// extends DatapointStream
type CpuUtilDatapointStream struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Which CPU was utilized.
	Cpu int `json:"cpu,omitempty"`
}

// ConfiguredStorageDevice - A storage device configured as usable storage.
// extends StorageDevice
type ConfiguredStorageDevice struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Model ID of the device.
	Model string `json:"model,omitempty"`
	// Physical size of the device, in bytes.
	// base = 1024
	// units = B
	Size float64 `json:"size,omitempty"`
	// Serial number of the device.
	Serial string `json:"serial,omitempty"`
	// Boolean value indicating if this is a boot device.
	BootDevice *bool `json:"bootDevice,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Vendor ID of the device.
	Vendor string `json:"vendor,omitempty"`
	// True if the device is currently configured in the system.
	Configured *bool `json:"configured,omitempty"`
	// Amount of additional space that would be made available, if the device
	// is expanded.
	// base = 1024
	// units = B
	ExpandableSize float64 `json:"expandableSize,omitempty"`
	// Size of allocated space on the device.
	// base = 1024
	// units = B
	UsedSize float64 `json:"usedSize,omitempty"`
}

// FaultEffect - An error affecting a user object whose root cause is a fault. A fault
// effect can only be resolved by resolving the fault which is its root
// cause.
// extends PersistentObject
type FaultEffect struct {
	// The cause of the fault effect, in case there is a chain of fault
	// effects originating from the root cause which resulted in this effect.
	// format = objectReference
	// referenceTo = /delphix-fault-effect.json
	CausedBy string `json:"causedBy,omitempty"`
	// The date when the root cause fault was diagnosed.
	// format = date
	DateDiagnosed string `json:"dateDiagnosed,omitempty"`
	// The name of the user-visible Delphix object that has a fault effect.
	TargetName string `json:"targetName,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The automated response taken by the system.
	Response string `json:"response,omitempty"`
	// The user-visible Delphix object that has a fault effect.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The root cause of this fault effect. Resolving the fault effect can
	// only occur by resolving its root cause.
	// format = objectReference
	// referenceTo = /delphix-fault.json
	RootCause string `json:"rootCause,omitempty"`
	// The severity of the fault effect. This can either be CRITICAL or
	// WARNING.
	// enum = [CRITICAL WARNING]
	Severity string `json:"severity,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A unique dot delimited identifier associated with the fault effect.
	BundleID string `json:"bundleID,omitempty"`
	// Summary of the fault effect.
	Title string `json:"title,omitempty"`
	// A suggested user action.
	Action string `json:"action,omitempty"`
	// Full description of the fault effect.
	Description string `json:"description,omitempty"`
}

// PgSQLIdentEntry - An entry in the PostgreSQL username map file (pg_ident.conf).
// extends TypedObject
type PgSQLIdentEntry struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the map to which this entry belongs (used to refer to the
	// map in pg_hba.conf).
	// create = required
	MapName string `json:"mapName,omitempty"`
	// The operating system username this entry matches.
	// create = required
	SystemUsername string `json:"systemUsername,omitempty"`
	// The database username this entry matches.
	// maxLength = 63
	// create = required
	DatabaseUsername string `json:"databaseUsername,omitempty"`
}

// JSTemplateUsageData - The space usage information for a data template.
// extends TypedObject
type JSTemplateUsageData struct {
	// The amount of space consumed by the bookmarks on this data template.
	// This is the space that will be freed up if all bookmarks on the
	// template were deleted. This presumes that all of child data containers
	// are purged first.
	// base = 1024
	// units = B
	Bookmarks float64 `json:"bookmarks,omitempty"`
	// The amount of space that would be consumed by the data in this
	// template (and child containers) without Delphix.
	// base = 1024
	// units = B
	Unvirtualized float64 `json:"unvirtualized,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The data template that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	Template string `json:"template,omitempty"`
	// The space that will be freed up if this template (and all of its child
	// data containers are deleted).
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
	// The amount of space consumed by data containers that were provisioned
	// from this data template. This is the space that will be freed up if
	// all of those data containers are deleted or purged. This assumes that
	// the data containers are deleted along with underlying data sources.
	// units = B
	// base = 1024
	Containers float64 `json:"containers,omitempty"`
}

// OracleInstall - The Oracle source repository.
// extends SourceRepository
type OracleInstall struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Reference to the environment containing this repository.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// 32 or 64 bits.
	// enum = [32 64]
	// create = required
	// update = optional
	Bits int `json:"bits,omitempty"`
	// Flag indicating whether this repository can use LogSync.
	LogsyncPossible *bool `json:"logsyncPossible,omitempty"`
	// The Oracle install home.
	// create = required
	// maxLength = 256
	InstallationHome string `json:"installationHome,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Flag indicating whether the install was discovered or manually
	// entered.
	Discovered *bool `json:"discovered,omitempty"`
	// Group name of the user that owns the install.
	GroupName string `json:"groupName,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Version of the repository.
	// create = required
	// update = optional
	// format = oracleVersion
	Version string `json:"version,omitempty"`
	// Flag indicating whether the install supports Oracle RAC.
	Rac *bool `json:"rac,omitempty"`
	// User name of the user that owns the install.
	UserName string `json:"userName,omitempty"`
	// List of Oracle patches that have been applied to this Oracle Home.
	// update = optional
	// create = optional
	AppliedPatches []int `json:"appliedPatches,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Group ID of the user that owns the install.
	GroupId int `json:"groupId,omitempty"`
	// User ID of the user that owns the install.
	UserId int `json:"userId,omitempty"`
	// The Oracle base where database binaries are located.
	// create = optional
	// update = optional
	// maxLength = 256
	OracleBase string `json:"oracleBase,omitempty"`
}

// AppDataProvisionParameters - The parameters to use as input to provision AppData.
// extends ProvisionParameters
type AppDataProvisionParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *AppDataContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *AppDataVirtualSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
}

// JSOperationEndpointDataLayoutParameters - The data layout to fetch the first and last events from.
// extends JSOperationEndpointParameters
type JSOperationEndpointDataLayoutParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The data layout to search.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
}

// BooleanEqualConstraint - Constraints placed on a boolean axis of a particular analytics slice.
// extends BooleanConstraint
type BooleanEqualConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be equal to the boolean argument.
	// create = required
	Equals *bool `json:"equals,omitempty"`
}

// JSBookmarkCreateParameters - The parameters used to create a Jet Stream bookmark.
// extends TypedObject
type JSBookmarkCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Jet Stream bookmark object.
	// required = true
	Bookmark *JSBookmark `json:"bookmark,omitempty"`
	// The Jet Stream data timeline point at which the bookmark will be
	// created.
	// required = true
	TimelinePointParameters JSTimelinePointTimeParameters `json:"timelinePointParameters,omitempty"`
}

// PgSQLStagingSource - A PostgreSQL staging source used for pre-provisioning.
// extends PgSQLSource
type PgSQLStagingSource struct {
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// The base mount point for the NFS mounts on the pre-provisioning host.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntime `json:"runtime,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowPointSemantic - TimeFlow point based on a semantic reference.
// extends TimeflowPointParameters
type TimeflowPointSemantic struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	// create = optional
	// update = optional
	Container string `json:"container,omitempty"`
	// A semantic description of a TimeFlow location.
	// enum = [LATEST_POINT LATEST_SNAPSHOT]
	// default = LATEST_POINT
	// create = optional
	// update = optional
	Location string `json:"location,omitempty"`
}

// SnapshotCapacityData - Capacity metrics for a single snapshot.
// extends TypedObject
// cliVisibility = [DOMAIN]
type SnapshotCapacityData struct {
	// Reference to the container to which this snapshot belongs.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Space used by the snapshot.
	// base = 1024
	// units = B
	Space float64 `json:"space,omitempty"`
	// Reference to the snapshot.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	Snapshot string `json:"snapshot,omitempty"`
	// Time at which this snapshot was taken.
	// format = date
	SnapshotTimestamp string `json:"snapshotTimestamp,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether this snapshot is currently being retained due to policy
	// settings.
	PolicyRetention *bool `json:"policyRetention,omitempty"`
	// The manual retention setting on this snapshot, in days.
	ManualRetention int `json:"manualRetention,omitempty"`
	// List of VDBs that have been provisioned from this snapshot.
	DescendantVDBs []string `json:"descendantVDBs,omitempty"`
}

// ASESnapshotRuntime - Runtime (non-persistent) properties of a SAP ASE TimeFlow snapshot.
// extends SnapshotRuntime
type ASESnapshotRuntime struct {
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// APISession - Describes a Delphix web service session and is the result of an
// initial handshake.
// extends TypedObject
type APISession struct {
	// Locale as an IETF BCP 47 language tag, defaults to 'en-US'.
	// format = locale
	// required = false
	Locale string `json:"locale,omitempty"`
	// Client software identification token.
	// required = false
	// maxLength = 64
	Client string `json:"client,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Version of the API to use.
	// required = true
	Version *APIVersion `json:"version,omitempty"`
}

// ToolkitLinkedStagedSource - A linked staged source definition for toolkits.
// extends ToolkitLinkedSource
type ToolkitLinkedStagedSource struct {
	// A workflow script to run just prior to snapshotting the staged source.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// A workflow script that stop the staged source. The staged files will
	// be mounted and available. Upon completion of this workflow, the staged
	// files will be unmounted.
	// required = true
	StopStaging string `json:"stopStaging,omitempty"`
	// A workflow script run periodically to monitor the health of the data
	// source and staging environment. This script will be run every 10
	// seconds.
	// required = false
	Worker string `json:"worker,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A user defined schema for the linking parameters.
	// required = true
	Parameters *SchemaDraftV4 `json:"parameters,omitempty"`
	// A workflow script to run immediately after snapshotting the staged
	// source.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// A workflow script that specifies where the virtual copy of the
	// application should be mounted.
	// required = false
	MountSpec string `json:"mountSpec,omitempty"`
	// A workflow script that builds the staging instance from production.
	// required = true
	Resync string `json:"resync,omitempty"`
	// A workflow script that start the staged source. The staged files will
	// be mounted and available.
	// required = true
	StartStaging string `json:"startStaging,omitempty"`
	// A workflow script that returns whether or not the data source is
	// active/inactive. The script should exit with an exit status of ACTIVE
	// if the data source is available. The script should exit with an exit
	// status of INACTIVE if the data source is unavailable. An exit status
	// of UNKNOWN implies the script encountered an unexpected state or
	// error. If no status script is supplied, the dSource will always be in
	// an active state while enabled.
	// required = false
	Status string `json:"status,omitempty"`
}

// RegistrationStatus - Information on the status of the Delphix Engine's registration.
// extends UserObject
// cliVisibility = [SYSTEM]
type RegistrationStatus struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The time at which the registration status was last updated.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// The status of the Delphix Engine's registration. It may be unknown,
	// unregistered, in-progress, or registered.
	// update = optional
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
}

// QuotaPolicy - This policy limits the maximum amount of space an object (group or
// database) can use.
// extends Policy
type QuotaPolicy struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// Whether this policy has been directly applied or inherited. See the
	// effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Last time a critical alert was generated.
	// format = date
	CritAlertTime string `json:"critAlertTime,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// True if this is the default policy created when the system is setup.
	// Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Size of the quota, in bytes.
	// units = B
	// base = 1024
	// create = required
	// update = optional
	// minimum = 1
	Size float64 `json:"size,omitempty"`
	// Last time a warning alert was generated.
	// format = date
	WarnAlertTime string `json:"warnAlertTime,omitempty"`
}

// SMTPConfig - SMTP configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SMTPConfig struct {
	// True if TLS (transport layer security) should be used.
	// default = false
	// update = optional
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// If authentication is enabled, username to use when authenticating to
	// the server.
	// update = optional
	Username string `json:"username,omitempty"`
	// From address to use when sending mail. If unspecified,
	// 'noreply@delphix.com' is used.
	// format = email
	// update = optional
	FromAddress string `json:"fromAddress,omitempty"`
	// True if outbound email is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// IP address or hostname of SMTP relay server.
	// format = host
	// update = optional
	Server string `json:"server,omitempty"`
	// Port number to use. A value of -1 indicates the default (25 or 587 for
	// TLS).
	// default = -1
	// update = optional
	// minimum = -1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// Maximum timeout to wait, in seconds, when sending mail.
	// minimum = 1
	// update = optional
	SendTimeout int `json:"sendTimeout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if username/password authentication should be used.
	// default = false
	// update = optional
	AuthenticationEnabled *bool `json:"authenticationEnabled,omitempty"`
	// If authentication is enabled, password to use when authenticating to
	// the server.
	// format = password
	// update = optional
	Password string `json:"password,omitempty"`
}

// SchemaDraftV4 - A dummy schema that is used to represent JSON that is a valid Draft v4
// schema.
type SchemaDraftV4 struct {
}

// PurgeLogsParameters - Represents the parameters of a purgeLogs request.
// extends TypedObject
type PurgeLogsParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of space in bytes to reclaim as part of purgeLogs process.
	// minimum = 1
	// units = B
	// base = 1024
	// required = true
	StorageSpaceToReclaim float64 `json:"storageSpaceToReclaim,omitempty"`
	// If this is set to true, this operation does not actually delete logs.
	// It returns the affected snapshots and truncated timeline as if the
	// logs were deleted.
	// default = true
	// required = true
	DryRun *bool `json:"dryRun,omitempty"`
	// Delete expired logs which have been retained to make snapshots
	// consistent.
	// default = false
	// required = true
	DeleteSnapshotLogs *bool `json:"deleteSnapshotLogs,omitempty"`
}

// HostRuntime - Runtime, non-persistent properties for a host machine.
// extends TypedObject
type HostRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the host is up and a connection can be established.
	Available *bool `json:"available,omitempty"`
	// The time that the 'available' propery was last checked.
	// format = date
	AvailableTimestamp string `json:"availableTimestamp,omitempty"`
	// The reason why the host is not available.
	NotAvailableReason string `json:"notAvailableReason,omitempty"`
	// Traceroute network hops from host to Delphix Engine.
	TraceRouteInfo *TracerouteInfo `json:"traceRouteInfo,omitempty"`
}

// OracleRACInstance - The representation of an Oracle Database RAC Instance Configuration.
// extends OracleInstance
type OracleRACInstance struct {
	// Reference to the Oracle cluster node that the RAC instance is running
	// on.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// create = required
	Node string `json:"node,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The number of the instance.
	// create = required
	// update = optional
	// minimum = 1
	InstanceNumber float64 `json:"instanceNumber,omitempty"`
	// The name of the instance.
	// create = required
	// update = optional
	// pattern = ^[a-zA-Z0-9_]+$
	// maxLength = 15
	InstanceName string `json:"instanceName,omitempty"`
}

// SeverityFilter - An alert filter that specifies which severity levels to match against.
// extends AlertFilter
type SeverityFilter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// List of severity levels. Only alerts matching one of the given
	// severity levels are included.
	// uniqueItems = true
	// minItems = 1
	// create = optional
	// update = optional
	SeverityLevels []string `json:"severityLevels,omitempty"`
}

// OperationTemplate - Template for commonly used operations.
// extends NamedUserObject
type OperationTemplate struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The name clients should use when setting the parameter's value.
	// create = required
	// update = optional
	// minLength = 1
	Name string `json:"name,omitempty"`
	// Template contents.
	// create = required
	// update = optional
	Operation SourceOperation `json:"operation,omitempty"`
	// Most recently modified time.
	// format = date
	LastUpdated string `json:"lastUpdated,omitempty"`
	// User provided description for this template.
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// PgSQLTimeflowPoint - A unique point within a PostgreSQL database TimeFlow.
// extends TimeflowPoint
type PgSQLTimeflowPoint struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-pgsql-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ObjectNotification - An event indicating a change to an object on the system.
// extends Notification
type ObjectNotification struct {
	// Target object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Object string `json:"object,omitempty"`
	// Type of target object.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
	// Type of operation on the object.
	// enum = [CREATE UPDATE DELETE]
	EventType string `json:"eventType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ContainerUtilizationInterval - Represents a represents an interval of time with which container
// utilization statistics are associated.
// extends TypedObject
type ContainerUtilizationInterval struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The start time of the interval.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// The average throughput for the container throughout the duration of
	// the sampling interval, measured in kilobits per second.
	AverageThroughput float64 `json:"averageThroughput,omitempty"`
}

// EnumEqualConstraint - Constraints placed on an enumeration axis of a particular analytics
// slice.
// extends EnumConstraint
type EnumEqualConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be equal to the specified value.
	// create = required
	Equals string `json:"equals,omitempty"`
}

// CapacityBreakdown - Storage stats breakdown.
// extends TypedObject
type CapacityBreakdown struct {
	// Unvirtualized space used by the TimeFlow.
	// units = B
	// base = 1024
	TimeflowUnvirtualizedSpace float64 `json:"timeflowUnvirtualizedSpace,omitempty"`
	// Amount of space used for snapshots from which VDBs have been
	// provisioned.
	// base = 1024
	// units = B
	DescendantSpace float64 `json:"descendantSpace,omitempty"`
	// Amount of space used for snapshots part of held space.
	// units = B
	// base = 1024
	UnownedSnapshotSpace float64 `json:"unownedSnapshotSpace,omitempty"`
	// Amount of space used by snapshots.
	// units = B
	// base = 1024
	SyncSpace float64 `json:"syncSpace,omitempty"`
	// Amount of space used for snapshots held by policy settings.
	// units = B
	// base = 1024
	PolicySpace float64 `json:"policySpace,omitempty"`
	// Actual space used by the container.
	// units = B
	// base = 1024
	ActualSpace float64 `json:"actualSpace,omitempty"`
	// Unvirtualized space used by the container.
	// units = B
	// base = 1024
	UnvirtualizedSpace float64 `json:"unvirtualizedSpace,omitempty"`
	// Amount of space used for the active copy of the container.
	// units = B
	// base = 1024
	ActiveSpace float64 `json:"activeSpace,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of space used by logs.
	// units = B
	// base = 1024
	LogSpace float64 `json:"logSpace,omitempty"`
	// Amount of space used for snapshots held by manual retention settings.
	// units = B
	// base = 1024
	ManualSpace float64 `json:"manualSpace,omitempty"`
}

// JSConfig - Jet Stream configuration.
// extends TypedObject
type JSConfig struct {
	// The number of times to retry failed sources during Jet Stream data
	// operations.
	// update = optional
	// minimum = 0
	RetryAttempts int `json:"retryAttempts,omitempty"`
	// Default expiration for bookmarks created through the GUI, in days. If
	// value is 0, bookmarks will default to no expiration.
	// update = optional
	// minimum = 0
	DefaultBookmarkExpiration int `json:"defaultBookmarkExpiration,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleAttachData - Represents parameters to attach non-pluggable Oracle databases.
// extends OracleBaseAttachData
type OracleAttachData struct {
	// True if initial load should be done immediately.
	// default = false
	// create = optional
	LinkNow *bool `json:"linkNow,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession to
	// reduce the number of logs required to provision the snapshot. This may
	// significantly reduce the time necessary to provision from a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// Total number of transport connections to use during SnapSync.
	// default = 1
	// minimum = 1
	// maximum = 16
	// create = optional
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
	// True if SnapSync data from the source should be retrieved through an
	// encrypted connection. Enabling this feature can decrease the
	// performance of SnapSync from the source but has no impact on the
	// performance of VDBs created from the retrieved data.
	// default = false
	// create = optional
	EncryptedLinkingEnabled *bool `json:"encryptedLinkingEnabled,omitempty"`
	// True if SnapSync data from the source should be compressed over the
	// network. Enabling this feature will reduce network bandwidth
	// consumption and may significantly improve throughput, especially over
	// slow network.
	// create = optional
	// default = true
	CompressedLinkingEnabled *bool `json:"compressedLinkingEnabled,omitempty"`
	// The name of the DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// Number of data files to include in each RMAN backup set.
	// minimum = 1
	// maximum = 64
	// default = 5
	// create = optional
	FilesPerSet int `json:"filesPerSet,omitempty"`
	// The password for the DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// True if extended block checking should be used for this linked
	// database.
	// default = false
	// create = optional
	CheckLogical *bool `json:"checkLogical,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Number of parallel channels to use.
	// minimum = 1
	// maximum = 32
	// default = 2
	// create = optional
	RmanChannels int `json:"rmanChannels,omitempty"`
	// True if attach should succeed even if the resetlogs of the original
	// database does not match the resetlogs of the new database and the
	// resetlogs information of the original database is not a parent
	// incarnation of the current resetlogs. This can happen when the
	// controlfile has been recreated and the incarnation table is
	// incomplete. Use this option with extreme caution. Attached database
	// must be the same database to avoid data corruption later on.
	// default = false
	// create = optional
	Force *bool `json:"force,omitempty"`
	// Bandwidth limit (MB/s) for SnapSync and LogSync network traffic. A
	// value of 0 means no limit.
	// minimum = 0
	// default = 0
	// create = optional
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// Information about the OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Defines whether backup level is enabled.
	// create = optional
	BackupLevelEnabled *bool `json:"backupLevelEnabled,omitempty"`
}

// MSSqlFailoverClusterDBConfig - Database for a SQL Server Failover Cluster.
// extends MSSqlDBConfig
type MSSqlFailoverClusterDBConfig struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The name of the database.
	// create = required
	// update = optional
	// maxLength = 128
	DatabaseName string `json:"databaseName,omitempty"`
	// Base drive letter location for mount points.
	// create = required
	// minLength = 1
	// maxLength = 1
	DriveLetter string `json:"driveLetter,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Specifies the current recovery model of the source database.
	// enum = [FULL SIMPLE BULK_LOGGED]
	// default = SIMPLE
	// create = optional
	// update = readonly
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The username of the database user.
	// maxLength = 256
	// update = optional
	User string `json:"user,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-mssql-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// Whether this source should be used for linking.
	// create = readonly
	// update = readonly
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The password of the database user.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
}

// PgSQLExportParameters - The parameters to use as input to export PostgreSQL databases.
// extends DbExportParameters
type PgSQLExportParameters struct {
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayout `json:"filesystemLayout,omitempty"`
	// Entries in the PostgreSQL host-based authentication file
	// (pg_hba.conf).
	// create = optional
	HbaEntries []*PgSQLHBAEntry `json:"hbaEntries,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The source config to use when creating the exported database.
	// required = true
	SourceConfig *PgSQLDBClusterConfig `json:"sourceConfig,omitempty"`
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Entries in the PostgreSQL username map file (pg_ident.conf).
	// create = optional
	IdentEntries []*PgSQLIdentEntry `json:"identEntries,omitempty"`
}

// UserAuthInfo - Summary authorization information about the current user.
// extends TypedObject
type UserAuthInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The currently logged in user.
	User *User `json:"user,omitempty"`
	// A reference to the system-defined owner role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	OwnerRole string `json:"ownerRole,omitempty"`
	// A reference to the system-defined provisioner role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	ProvisionerRole string `json:"provisionerRole,omitempty"`
	// A reference to the system-defined Jet Stream user role.
	// format = objectReference
	// referenceTo = /delphix-role.json
	JetStreamUserRole string `json:"jetStreamUserRole,omitempty"`
	// The list of authorizations granted to the current user.
	Authorizations []*Authorization `json:"authorizations,omitempty"`
}

// FileMappingParameters - Input parameters to test file mapping rules.
// extends TypedObject
type FileMappingParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The list of TimeFlow point, bookmark, or semantic location to use for
	// source files to be mapped.
	// required = true
	// minItems = 1
	TimeflowPointParameters []TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// Database file mapping rules.
	// required = true
	MappingRules string `json:"mappingRules,omitempty"`
}

// JSTimelinePointBookmarkInput - Specifies the Jet Stream timeline point using a reference to the Jet
// Stream bookmark.
// extends JSTimelinePointParameters
type JSTimelinePointBookmarkInput struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Jet Stream bookmark.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-js-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
}

// TlsCaCert - Trusted CA certificate.
// extends NamedUserObject
// cliVisibility = [SYSTEM DOMAIN]
type TlsCaCert struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Certificate serial number.
	SerialNumber string `json:"serialNumber,omitempty"`
	// Start of validity.
	ValidFrom string `json:"validFrom,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// Distinguished name of subject of this certificate.
	IssuedToDN string `json:"issuedToDN,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The alias for this certificate.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// End of validity.
	ValidTo string `json:"validTo,omitempty"`
	// SHA-1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// Issuer of this certificate.
	IssuedByDN string `json:"issuedByDN,omitempty"`
}

// CertificateChainImportParameters - Parameters for importing a certificate chain.
// extends TypedObject
type CertificateChainImportParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A sequence of X.509 certificates in PEM format.
	// create = required
	CertificateChain string `json:"certificateChain,omitempty"`
}

// TimeflowPointBookmark - TimeFlow point based on a TimeFlow bookmark.
// extends TimeflowPointParameters
type TimeflowPointBookmark struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the bookmark.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow-bookmark.json
	Bookmark string `json:"bookmark,omitempty"`
}

// HistoricalGroupCapacityData - Historical capacity data aggregated over a group.
// extends BaseGroupCapacityData
// cliVisibility = [DOMAIN]
type HistoricalGroupCapacityData struct {
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdown `json:"source,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdown `json:"virtual,omitempty"`
	// Which group these stats represent.
	// referenceTo = /delphix-group.json
	// format = objectReference
	Group string `json:"group,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSTimelinePointLatestTimeInput - Specifies the use of the latest available data from the given data
// layout.
// extends JSTimelinePointTimeParameters
type JSTimelinePointLatestTimeInput struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The reference to the data layout used for this operation.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	// required = true
	SourceDataLayout string `json:"sourceDataLayout,omitempty"`
}

// RetentionPolicy - This policy controls what data (log and snapshot) is kept.
// extends Policy
type RetentionPolicy struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Amount of time (in dataUnit units) to keep source data.
	// create = optional
	// update = optional
	DataDuration int `json:"dataDuration,omitempty"`
	// Time unit for logDuration.
	// create = optional
	// update = optional
	// enum = [DAY WEEK MONTH QUARTER YEAR]
	LogUnit string `json:"logUnit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Day of week upon which to enforce weekly snapshot retention.
	// create = optional
	// update = optional
	// enum = [MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY]
	DayOfWeek string `json:"dayOfWeek,omitempty"`
	// Number of monthly snapshots to keep.
	// create = optional
	// update = optional
	NumOfMonthly int `json:"numOfMonthly,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Day of year upon which to enforce yearly snapshot retention, expressed
	// a month / day string (e.g., "Jan 1").
	// create = optional
	// update = optional
	// maxLength = 32
	DayOfYear string `json:"dayOfYear,omitempty"`
	// Time unit for dataDuration.
	// update = optional
	// enum = [DAY WEEK MONTH QUARTER YEAR]
	// create = optional
	DataUnit string `json:"dataUnit,omitempty"`
	// Number of daily snapshots to keep.
	// create = optional
	// update = optional
	NumOfDaily int `json:"numOfDaily,omitempty"`
	// Day of month upon which to enforce monthly snapshot retention.
	// create = optional
	// update = optional
	DayOfMonth int `json:"dayOfMonth,omitempty"`
	// True if this is the default policy created when the system is setup.
	// Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// Amount of time (in logUnit units) to keep log data.
	// create = optional
	// update = optional
	LogDuration int `json:"logDuration,omitempty"`
	// Number of yearly snapshots to keep.
	// create = optional
	// update = optional
	NumOfYearly int `json:"numOfYearly,omitempty"`
	// Whether this policy has been directly applied or inherited. See the
	// effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// Number of weekly snapshots to keep.
	// create = optional
	// update = optional
	NumOfWeekly int `json:"numOfWeekly,omitempty"`
}

// RunMaskingJobOperation - An operation that runs a Masking Job on the local Delphix Masking
// Engine instance.
// extends Operation
type RunMaskingJobOperation struct {
	// The application ID of the Masking Job to be executed.
	// create = readonly
	// update = readonly
	ApplicationId string `json:"applicationId,omitempty"`
	// The Masking Job ID of the Masking Job to be executed.
	// create = readonly
	// update = readonly
	MaskingJobId string `json:"maskingJobId,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The location this Masking Job will be executed on.
	// create = readonly
	// update = readonly
	Host string `json:"host,omitempty"`
}

// ToolkitVirtualSource - A virtual source definition for toolkits.
// extends TypedObject
type ToolkitVirtualSource struct {
	// A workflow script to run when stopping a virtual copy of the
	// application.
	// required = true
	Stop string `json:"stop,omitempty"`
	// A workflow script run when returning a virtual copy of the appliction
	// to an environment that it was previously removed from.
	// required = true
	Reconfigure string `json:"reconfigure,omitempty"`
	// A user defined schema for the provisioning parameters.
	// required = true
	Parameters *SchemaDraftV4 `json:"parameters,omitempty"`
	// A workflow script to run when starting a virtual copy of the
	// application.
	// required = true
	Start string `json:"start,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A workflow script to run after taking a snapshot of a virtual copy of
	// the application.
	// required = true
	PostSnapshot string `json:"postSnapshot,omitempty"`
	// A workflow script that specifies where the virtual copy of the
	// application should be mounted.
	// required = false
	MountSpec string `json:"mountSpec,omitempty"`
	// A workflow script run when configuring a virtual copy of the
	// application in a new environment.
	// required = true
	Configure string `json:"configure,omitempty"`
	// A workflow script run when removing a virtual copy of the application
	// from an environment (e.g. on delete, disable, or refresh).
	// required = true
	Unconfigure string `json:"unconfigure,omitempty"`
	// A workflow script to run before taking a snapshot of a virtual copy of
	// the application.
	// required = true
	PreSnapshot string `json:"preSnapshot,omitempty"`
	// The workflow script to run to determine if a virtual copy of the
	// application is running. The script should output 'ACTIVE' if the
	// application is running, 'INACTIVE' if the application is not running,
	// and 'UNKNOWN' if the script encounters an unexpected problem.
	// required = false
	Status string `json:"status,omitempty"`
	// A workflow script to run when creating an empty application.
	// required = false
	Initialize string `json:"initialize,omitempty"`
}

// NetworkInterface - Configuration of an IP interface.
// extends UserObject
// cliVisibility = [DOMAIN SYSTEM]
type NetworkInterface struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The name of the device over which this interface is configured.
	Device string `json:"device,omitempty"`
	// The maximum transmission unit for this interface.
	// update = optional
	Mtu int `json:"mtu,omitempty"`
	// The range of possible values for the mtu property.
	MtuRange string `json:"mtuRange,omitempty"`
	// The MAC address associated with this interface.
	// format = macAddress
	MacAddress string `json:"macAddress,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The state of the interface.
	// enum = [OK DOWN FAILED]
	State string `json:"state,omitempty"`
	// List of IP addresses assigned to the interface.
	// update = optional
	Addresses []*InterfaceAddress `json:"addresses,omitempty"`
}

// HostPrivilegeElevationSettings - Settings for elevating user privileges on a host.
// extends TypedObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationSettings struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The default privilege elevation profile for new environments.
	// format = objectReference
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// update = required
	DefaultProfile string `json:"defaultProfile,omitempty"`
}

// ASEDBContainer - An SAP ASE Database Container.
// extends DatabaseContainer
type ASEDBContainer struct {
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// The operating system for the source database.
	Os string `json:"os,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// The processor type for the source database.
	Processor string `json:"processor,omitempty"`
	// Policies for managing LogSync and SnapSync across sources for an SAP
	// ASE container.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// Runtime properties of this container.
	Runtime *ASEDBContainerRuntime `json:"runtime,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
}

// MySQLTimeflow - TimeFlow representing historical data for a particular timeline within
// a MySQL container.
// extends Timeflow
type MySQLTimeflow struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *MySQLTimeflowPoint `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// referenceTo = /delphix-timeflow-snapshot.json
	// format = objectReference
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
}

// AppDataRepository - An AppData repository.
// extends SourceRepository
type AppDataRepository struct {
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
	// The toolkit associated with this repository.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	// create = required
	// update = optional
	Toolkit string `json:"toolkit,omitempty"`
	// The list of parameters specified by the repository schema in the
	// toolkit. If no schema is specified, this list is empty.
	// create = required
	Parameters *Json `json:"parameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeflowFilesystemLayout - A filesystem layout that matches the filesystem of a Delphix TimeFlow.
// extends FilesystemLayout
type TimeflowFilesystemLayout struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The directory for data files.
	// create = optional
	DataDirectory string `json:"dataDirectory,omitempty"`
	// The directory for archive files.
	// create = optional
	ArchiveDirectory string `json:"archiveDirectory,omitempty"`
	// The directory for external files.
	// create = optional
	ExternalDirectory string `json:"externalDirectory,omitempty"`
	// The directory for temporary files.
	// create = optional
	TempDirectory string `json:"tempDirectory,omitempty"`
	// The directory for script files.
	// create = optional
	ScriptDirectory string `json:"scriptDirectory,omitempty"`
	// The base directory to use for the exported database.
	// create = optional
	TargetDirectory string `json:"targetDirectory,omitempty"`
}

// MySQLXtraBackupSyncParameters - The parameters to use as input to sync requests for MySQL databases
// using an existing XtraBackup backup.
// extends MySQLExistingBackupSyncParameters
type MySQLXtraBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Path to the existing backup to be loaded.
	// required = true
	BackupLocation string `json:"backupLocation,omitempty"`
	// The coordinates corresponding to the MySQL backup to start replication
	// from.
	// create = optional
	ReplicationCoordinates MySQLReplicationCoordinates `json:"replicationCoordinates,omitempty"`
}

// SyncParameters - The parameters to use as input to sync requests.
// extends TypedObject
type SyncParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// PgSQLSyncParameters - The parameters to use as input to sync PostgreSQL databases.
// extends SyncParameters
type PgSQLSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether or not to take another full backup of the source database.
	// default = false
	RedoBaseBackup *bool `json:"redoBaseBackup,omitempty"`
}

// OracleTimeflowPoint - A unique point within an Oracle database TimeFlow.
// extends TimeflowPoint
type OracleTimeflowPoint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
}

// JSDataContainerDeleteParameters - The parameters used to delete a data container.
// extends TypedObject
type JSDataContainerDeleteParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If this value is true, then delete the underlying data from all data
	// sources.
	// required = true
	// default = true
	DeleteDataSources *bool `json:"deleteDataSources,omitempty"`
}

// OracleCustomEnvVarRACFile - Dictates an environment file to be sourced when the Delphix Engine
// administers an Oracle virtual database. This environment file must be
// available on the target environment. This type also includes
// parameters which will be passed to the environment file when it is
// sourced. For a RAC environment, the cluster node where the target
// environment file exists must also be specified.
// extends OracleCustomEnvVarFile
type OracleCustomEnvVarRACFile struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A string of whitespace-separated parameters to be passed to the source
	// command. The first parameter must be an absolute path to a file that
	// exists on the target environment. Every subsequent parameter will be
	// treated as an argument interpreted by the environment file.
	// required = true
	PathParameters string `json:"pathParameters,omitempty"`
	// The cluster node on which the target environment file exists.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster-node.json
	// required = true
	ClusterNode string `json:"clusterNode,omitempty"`
}

// LdapInfo - Global LDAP information.
// extends TypedObject
// cliVisibility = [SYSTEM DOMAIN]
type LdapInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether LDAP authentication is configured and enabled or not for this
	// Delphix Engine.
	Enabled *bool `json:"enabled,omitempty"`
}

// PgSQLSourceConnectionInfo - Contains information that can be used to connect to a PostgreSQL
// source.
// extends SourceConnectionInfo
type PgSQLSourceConnectionInfo struct {
	// The hostname or IP address of the host where the source resides.
	Host string `json:"host,omitempty"`
	// The port on which the PostgresSQL server for the cluster is listening.
	Port int `json:"port,omitempty"`
	// The username of the database cluster user.
	User string `json:"user,omitempty"`
	// The data directory for the PostgreSQL cluster.
	PgDataDirectory string `json:"pgDataDirectory,omitempty"`
	// The JDBC string used to connect to the PostgreSQL server instance.
	JdbcString string `json:"jdbcString,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
}

// ChecklistItem - Generic checklist item.
// extends TypedObject
type ChecklistItem struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CertificateImportParameters - Parameters for importing a certificate.
// extends TypedObject
type CertificateImportParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The unique name for this certificate.
	// required = true
	Alias string `json:"alias,omitempty"`
	// Certificate contents.
	// required = true
	Certificate string `json:"certificate,omitempty"`
}

// Role - Describes a role as applied to a user on an object.
// extends UserObject
type Role struct {
	// Determines if the role can be modified or not. Some roles are shipped
	// with the Delphix Engine and cannot be changed.
	Immutable *bool `json:"immutable,omitempty"`
	// List of permissions contained in the role.
	// create = required
	// update = optional
	Permissions []string `json:"permissions,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
}

// OracleClusterCreateParameters - The parameters used for the oracle cluster create operation.
// extends SourceEnvironmentCreateParameters
type OracleClusterCreateParameters struct {
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUser `json:"primaryUser,omitempty"`
	// The representation of the cluster object.
	// create = required
	Cluster *OracleCluster `json:"cluster,omitempty"`
	// The list of nodes in the cluster.
	// create = required
	Nodes []*OracleClusterNodeCreateParameters `json:"nodes,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// APIError - Description of an error encountered during an API call.
// extends TypedObject
type APIError struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A stable identifier for the class of error encountered.
	Id string `json:"id,omitempty"`
	// Extra output, often from a script or other external process, that may
	// give more insight into the cause of this error.
	CommandOutput string `json:"commandOutput,omitempty"`
	// Results of diagnostic checks run, if any, if the job failed.
	Diagnoses []*DiagnosisResult `json:"diagnoses,omitempty"`
	// For validation errors, a map of fields to APIError objects. For all
	// other errors, a string with further details of the error.
	Details string `json:"details,omitempty"`
	// Action to be taken by the user, if any, to fix the underlying problem.
	Action string `json:"action,omitempty"`
}

// KerberosKDC - Kerberos Client Configuration.
// extends TypedObject
type KerberosKDC struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// KDC Server hostname.
	// create = required
	// update = optional
	// format = host
	Hostname string `json:"hostname,omitempty"`
	// KDC Server port number.
	// create = required
	// update = optional
	// minimum = 0
	// maximum = 65535
	// default = 88
	Port int `json:"port,omitempty"`
}

// ReplicationSpecRuntime - Runtime properties for a replication spec.
// extends TypedObject
type ReplicationSpecRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Date of the next execution of this replication spec according to the
	// schedule.
	// format = date
	NextScheduledExecution string `json:"nextScheduledExecution,omitempty"`
}

// JSDataTemplateCreateParameters - The parameters used to create a data template.
// extends JSDataLayoutCreateParameters
type JSDataTemplateCreateParameters struct {
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParameters `json:"dataSources,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A description of this data layout to define what it is used for.
	// create = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
}

// OracleVirtualPdbSource - A virtual Oracle multitenant pluggable database source.
// extends OracleVirtualSource
type OracleVirtualPdbSource struct {
	// Online Redo Log size in MB. This is not applicable to pluggable
	// databases.
	// create = readonly
	// update = readonly
	RedoLogSizeInMB int `json:"redoLogSizeInMB,omitempty"`
	// Number of Online Redo Log Groups. This is not applicable to pluggable
	// databases.
	// update = readonly
	// create = readonly
	RedoLogGroups int `json:"redoLogGroups,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh. This is
	// currently not supported for pluggable databases.
	// create = readonly
	// update = readonly
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Archive Log Mode of the Oracle virtual database. This is not
	// applicable to pluggable databases.
	// create = readonly
	// update = readonly
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Oracle database configuration parameter overrides. This is currently
	// not supported for pluggable databases.
	// create = readonly
	// update = readonly
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-source-config.json
	// create = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// A list of object references to Oracle Node Listeners selected for this
	// Virtual Database. Delphix picks one default listener from the target
	// environment if this list is empty at virtual database provision time.
	// This is not applicable to pluggable databases.
	// create = readonly
	// update = readonly
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
}

// OracleFetchedLog - An Oracle log file fetched by LogSync.
// extends OracleTimeflowLog
type OracleFetchedLog struct {
	// Sequence number for the log file.
	Sequence int `json:"sequence,omitempty"`
	// Start SCN for the log file.
	StartScn int `json:"startScn,omitempty"`
	// End timestamp for the log file.
	// format = date
	EndTimestamp string `json:"endTimestamp,omitempty"`
	// Reference to the database to which this log belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-container.json
	Container string `json:"container,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the TimeFlow of which this log is a part.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Instance number associated with the log file.
	InstanceNum int `json:"instanceNum,omitempty"`
	// End SCN for the log file.
	EndScn int `json:"endScn,omitempty"`
	// Start timestamp for the log file.
	// format = date
	StartTimestamp string `json:"startTimestamp,omitempty"`
}

// PasswordCredential - The password based security credential.
// extends Credential
type PasswordCredential struct {
	// The password.
	// format = password
	// required = true
	// minLength = 1
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSBookmarkTagUsageData - The space usage information for a Jet Stream bookmark tag.
// extends TypedObject
type JSBookmarkTagUsageData struct {
	// The total amount of space referenced by bookmarks with this tag. This
	// is the sum of the bookmarks' unique, shared, and externallyReferenced
	// space.
	// base = 1024
	// units = B
	Referenced float64 `json:"referenced,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The amount of space that will be freed if bookmarks with this tag are
	// deleted.
	BookmarkTag string `json:"bookmarkTag,omitempty"`
	// The space that is being consumed by the set of bookmarks with the
	// given tag. This represents the minimum amount of space that will be
	// freed if all of the bookmarks are deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
}

// RemoteDelphixEngineInfo - Parameters for logging into another Delphix Engine when running a
// network throughput test.
// extends TypedObject
type RemoteDelphixEngineInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Address of other Delphix Engine.
	// format = host
	// create = required
	Address string `json:"address,omitempty"`
	// Username for the other Delphix Engine.
	// create = required
	Principal string `json:"principal,omitempty"`
	// Password for the other Delphix Engine.
	// create = required
	Credential *PasswordCredential `json:"credential,omitempty"`
}

// PgSQLLinkedSource - A linked PostgreSQL source.
// extends PgSQLSource
type PgSQLLinkedSource struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntime `json:"runtime,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// The staging source for pre-provisioning of the database.
	// format = objectReference
	// referenceTo = /delphix-pgsql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The external file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
}

// JSDailyOperationDuration - Information about the durations of a specific operation type for a
// data container over the past week.
// extends JSUsageData
type JSDailyOperationDuration struct {
	// The operation performed.
	// enum = [REFRESH RESET CREATE_BRANCH RESTORE UNDO]
	Operation string `json:"operation,omitempty"`
	// The number of times the specified operation was run in the past day.
	DailyCount int `json:"dailyCount,omitempty"`
	// The minimum duration in seconds of running the specified operation in
	// the past day.
	// units = sec
	DailyMinDuration int `json:"dailyMinDuration,omitempty"`
	// The object the usage data is centered around.
	// format = objectReference
	// referenceTo = /delphix-named-user-object.json
	UsageObject string `json:"usageObject,omitempty"`
	// The date at the beginning of the time period this datapoint
	// corresponds to. The time period itself varies between datapoint types.
	// format = date
	StartDate string `json:"startDate,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The average duration in seconds of running the specified operation in
	// the past day.
	// units = sec
	DailyAverageDuration int `json:"dailyAverageDuration,omitempty"`
	// The maximum duration in seconds of running the specified operation in
	// the past day.
	// units = sec
	DailyMaxDuration int `json:"dailyMaxDuration,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// NotificationDrop - An object to track dropped notifications.
// extends Notification
type NotificationDrop struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The number of notifications which were dropped since the last
	// notifications were pulled. If this is greater than zero, you may want
	// to refresh your view of the data to ensure everything is up to date.
	DropCount int `json:"dropCount,omitempty"`
}

// JSBranchUsageData - The space usage information for a Jet Stream branch.
// extends TypedObject
type JSBranchUsageData struct {
	// The Jet Stream branch that this usage information is for.
	// referenceTo = /delphix-js-branch.json
	// format = objectReference
	Branch string `json:"branch,omitempty"`
	// The name of the data container that this branch resides on.
	DataContainer string `json:"dataContainer,omitempty"`
	// The amount of space that will be freed if this branch is deleted.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
	// The amount of space that cannot be freed on the parent data template
	// (or sibling data containers) because it is also being referenced by
	// this branch due to restore or create branch operations.
	// base = 1024
	// units = B
	SharedOthers float64 `json:"sharedOthers,omitempty"`
	// The amount of space that cannot be freed up on this branch because it
	// is also being referenced by sibling data containers due to restore or
	// create branch operations.
	// base = 1024
	// units = B
	SharedSelf float64 `json:"sharedSelf,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// IntegerGreaterThanConstraint - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerGreaterThanConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must be greater than this value.
	// create = required
	GreaterThan int `json:"greaterThan,omitempty"`
}

// MySQLExistingMySQLDumpSyncParameters - The parameters to use as input to sync requests for MySQL databases
// using an existing MySQL dump.
// extends MySQLExistingBackupSyncParameters
type MySQLExistingMySQLDumpSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Path to the existing backup to be loaded.
	// required = true
	BackupLocation string `json:"backupLocation,omitempty"`
	// The coordinates corresponding to the MySQL backup to start replication
	// from.
	// create = optional
	ReplicationCoordinates MySQLReplicationCoordinates `json:"replicationCoordinates,omitempty"`
}

// StorageDeviceRemovalStatus - The status of a device removal of the storage in the system.
// extends TypedObject
// cliVisibility = [SYSTEM]
type StorageDeviceRemovalStatus struct {
	// Time removal was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Memory used to account for removed devices, in bytes.
	// units = B
	// base = 1024
	MappingMemory float64 `json:"mappingMemory,omitempty"`
	// Removal state.
	// enum = [NONE ACTIVE COMPLETED CANCELED]
	State string `json:"state,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of data removed, in bytes.
	// base = 1024
	// units = B
	Copied float64 `json:"copied,omitempty"`
	// Total amount of data to remove (including completed), in bytes.
	// base = 1024
	// units = B
	Total float64 `json:"total,omitempty"`
}

// OracleWarehouseSource - A warehouse represents a housing that accepts databases.
// extends OracleVirtualSource
type OracleWarehouseSource struct {
	// Custom environment variables for Oracle databases.
	// update = optional
	// create = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Number of Online Redo Log Groups.
	// create = optional
	// update = readonly
	// minimum = 2
	// default = 3
	RedoLogGroups int `json:"redoLogGroups,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = optional
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// A list of object references to Oracle Node Listeners selected for this
	// Virtual Database. Delphix picks one default listener from the target
	// environment if this list is empty at virtual database provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// update = optional
	// create = required
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Online Redo Log size in MB.
	// update = readonly
	// minimum = 4
	// create = optional
	RedoLogSizeInMB int `json:"redoLogSizeInMB,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// update = readonly
	// default = true
	// create = optional
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
}

// UnixHostCreateParameters - The parameters used for the add Unix host operation.
// extends HostCreateParameters
type UnixHostCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The host object.
	// create = required
	Host Host `json:"host,omitempty"`
}

// AppDataWindowsTimeflow - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends AppDataBaseTimeflow
type AppDataWindowsTimeflow struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *AppDataTimeflowPoint `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
}

// JSBranch - A branch represents a distinct timeline for data sources in a data
// layout.
// extends NamedUserObject
type JSBranch struct {
	// The last JSOperation on this branch by data time.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	LastOperation string `json:"lastOperation,omitempty"`
	// A reference to the data layout this branch was created on.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
	// The first JSOperation on this branch by data time.
	// referenceTo = /delphix-js-operation.json
	// format = objectReference
	FirstOperation string `json:"firstOperation,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
}

// NetworkLatencyTest - Round-trip latency tests to a target system.
// extends NetworkTest
// cliVisibility = [DOMAIN SYSTEM]
type NetworkLatencyTest struct {
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Percentage of requests or replies lost.
	Loss int `json:"loss,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkLatencyTestParameters `json:"parameters,omitempty"`
	// Maximum measured round-trip time (usec).
	// units = usec
	Maximum int `json:"maximum,omitempty"`
	// Standard deviation (usec).
	// units = usec
	Stddev int `json:"stddev,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// Minimum measured round-trip time (usec).
	// units = usec
	Minimum int `json:"minimum,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Average measured round-trip time (usec).
	// units = usec
	Average int `json:"average,omitempty"`
}

// ChecklistItemDetail - Fields to indicate detailed status for a specific checklist item.
// extends ChecklistItem
type ChecklistItemDetail struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
}

// OracleUndoDatafileSpecification - Describes an Oracle datafile that stores undo data.
// extends OracleDatafileTempfileSpecification
type OracleUndoDatafileSpecification struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the size
	// of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement int `json:"autoExtendIncrement,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space that
	// Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// create = optional
	// update = optional
	// minimum = 300
	// default = 300
	Size int `json:"size,omitempty"`
}

// MigrateCompatibilityParameters - The criteria necessary to select valid repositories for migration.
// extends CompatibleRepositoriesParameters
type MigrateCompatibilityParameters struct {
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// OracleDatabaseCreationParameters - The parameters to use as input when creating a new Oracle database.
// extends EmptyDatasetCreationParameters
type OracleDatabaseCreationParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the created database.
	// required = true
	Container *OracleDatabaseContainer `json:"container,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig OracleBaseDBConfig `json:"sourceConfig,omitempty"`
	// The password for the SYS user.
	// format = password
	// required = true
	SysPassword string `json:"sysPassword,omitempty"`
	// The maximum number of archived redo log files for automatic media
	// recovery of Oracle RAC.
	// default = 100
	// maximum = 65535
	// create = optional
	// update = optional
	// minimum = 0
	MaxLogHistory int `json:"maxLogHistory,omitempty"`
	// The source that describes the created database instance.
	// required = true
	Source *OracleWarehouseSource `json:"source,omitempty"`
	// The national character set used to store data in columns specifically
	// defined as NCHAR, NCLOB, or NVARCHAR2.
	// create = optional
	// update = optional
	// enum = [AL16UTF16 UTF8]
	// default = AL16UTF16
	NationalCharacterSet string `json:"nationalCharacterSet,omitempty"`
	// The password for the Delphix database user.
	// format = password
	// required = true
	DelphixPassword string `json:"delphixPassword,omitempty"`
	// The maximum number of redo log files that can ever be created for the
	// database.
	// minimum = 3
	// default = 64
	// maximum = 255
	// create = optional
	// update = optional
	MaxLogFiles int `json:"maxLogFiles,omitempty"`
	// The redo log files. If no filename is provided, Oracle-managed files
	// will be used.
	// create = optional
	// update = optional
	// minItems = 3
	RedoLogs []*OracleRedoLogFileSpecification `json:"redoLogs,omitempty"`
	// The password for the SYSTEM user.
	// format = password
	// required = true
	SystemPassword string `json:"systemPassword,omitempty"`
	// The datafile for the SYSTEM tablespace. If no filename is provided,
	// Oracle-managed files will be used.
	// create = optional
	// update = optional
	SysDatafile *OracleSystemDatafileSpecification `json:"sysDatafile,omitempty"`
	// The name of the Delphix database user.
	// required = true
	// pattern = ^[a-zA-Z][_a-zA-Z0-9]*$
	DelphixUsername string `json:"delphixUsername,omitempty"`
	// Grants the SELECT ANY DICTIONARY system privilege to the Delphix
	// database user. If disabled, the Delphix database user will only have
	// SELECT access to a limited set of views.
	// create = optional
	// update = optional
	// default = true
	GrantSelectAnyDictionary *bool `json:"grantSelectAnyDictionary,omitempty"`
	// The datafile to be used for undo data. If no filename is provided,
	// Oracle-managed files will be used.
	// create = optional
	// update = optional
	UndoTablespace *OracleUndoDatafileSpecification `json:"undoTablespace,omitempty"`
	// The initial sizing of the data files section of the control file at
	// CREATE DATABASE or CREATE CONTROLFILE time.
	// minimum = 2
	// default = 32
	// maximum = 65535
	// create = optional
	// update = optional
	MaxDataFiles int `json:"maxDataFiles,omitempty"`
	// The maximum number of instances that can simultaneously have this
	// database mounted and open.
	// create = optional
	// update = optional
	// minimum = 1
	// default = 32
	// maximum = 1055
	MaxInstances int `json:"maxInstances,omitempty"`
	// The datafile for the SYSAUX tablespace. If no filename is provided,
	// Oracle-managed files will be used.
	// update = optional
	// create = optional
	SysauxDatafile *OracleSysauxDatafileSpecification `json:"sysauxDatafile,omitempty"`
	// The character set the database uses to store data.
	// update = optional
	// pattern = ^UTF8|UTFE|[A-Z]+[0-9]+[A-Z0-9]+$
	// default = AL32UTF8
	// create = optional
	CharacterSet string `json:"characterSet,omitempty"`
	// Indicates the timezone file version that will be used to create the
	// database.
	// create = optional
	// update = optional
	// minimum = 1
	// maximum = 22
	TimezoneFileVersion int `json:"timezoneFileVersion,omitempty"`
	// Puts the database into FORCE LOGGING mode. Oracle Database will log
	// all changes in the database except for changes in temporary
	// tablespaces and temporary segments.
	// create = optional
	// update = optional
	// default = false
	ForceLogging *bool `json:"forceLogging,omitempty"`
	// The tempfile for the database. If no filename is provided,
	// Oracle-managed files will be used.
	// create = optional
	// update = optional
	TempTablespace *OracleTempfileSpecification `json:"tempTablespace,omitempty"`
}

// AuthFilterResult - Result of an auth filter request.
// extends TypedObject
type AuthFilterResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The list of objects that have been filtered.
	Objects []string `json:"objects,omitempty"`
}

// JSDataSourceCreateParameters - The parameters used to create the Jet Stream data sources.
// extends TypedObject
type JSDataSourceCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Jet Stream data source object.
	// required = true
	Source *JSDataSource `json:"source,omitempty"`
	// A reference to the underlying container object.
	// referenceTo = /delphix-container.json
	// required = true
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Key/value pairs used to specify attributes for this data source.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
}

// SupportBundleUploadParameters - Parameters to be used when uploading a support bundle.
// extends BaseSupportBundleParameters
type SupportBundleUploadParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether or not to include the analytics data in the support bundle
	// which is generated. Including analytics data may significantly
	// increase the support bundle size and upload time, but enables analysis
	// of performance characteristics of the Delphix Engine.
	// default = false
	IncludeAnalyticsData *bool `json:"includeAnalyticsData,omitempty"`
	// Type of support bundle to generate. Reserved for Delphix support use.
	// default = ALL
	// enum = [PHONEHOME MDS OS CORE LOG DROPBOX STORAGE_TEST MASKING ALL]
	BundleType string `json:"bundleType,omitempty"`
	// The Delphix support case number.
	CaseNumber int `json:"caseNumber,omitempty"`
}

// StringEqualConstraint - Constraint placed on a string axis of a particular analytics slice.
// extends StringConstraint
type StringEqualConstraint struct {
	// The axis values must match this string.
	// create = required
	Equals string `json:"equals,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
}

// OracleRACSourceConnectionInfo - Contains information that can be used to connect to a single instance
// Oracle source.
// extends OracleSourceConnectionInfo
type OracleRACSourceConnectionInfo struct {
	// The JDBC strings used to connect to the source.
	JdbcStrings []string `json:"jdbcStrings,omitempty"`
	// The default remote_listener parameter to be used for databases on the
	// cluster.
	RemoteListener string `json:"remoteListener,omitempty"`
	// The addresses for the nodes on which the source resides.
	Nodes []string `json:"nodes,omitempty"`
	// The Single Client Access Name of the cluster (11.2 and greater
	// clusters only).
	Scan string `json:"scan,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The database version string.
	Version string `json:"version,omitempty"`
	// The database name.
	DatabaseName string `json:"databaseName,omitempty"`
	// The Oracle installation home.
	OracleHome string `json:"oracleHome,omitempty"`
	// The location of the cluster installation.
	CrsClusterHome string `json:"crsClusterHome,omitempty"`
}

// GlobalLinkingSettings - System-wide linking settings.
// extends TypedObject
type GlobalLinkingSettings struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if encrypted linking should be enabled by default on new
	// dSources.
	// update = optional
	EncryptedLinkingEnabledByDefault *bool `json:"encryptedLinkingEnabledByDefault,omitempty"`
}

// OracleRACConfig - Representation of the properties specific to a RAC Oracle DB
// configuration.
// extends OracleDBConfig
type OracleRACConfig struct {
	// The username of the database user.
	// maxLength = 30
	// update = optional
	User string `json:"user,omitempty"`
	// The name of the database.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 8
	DatabaseName string `json:"databaseName,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The username of a database user that does not have administrative
	// privileges.
	// update = optional
	// maxLength = 30
	// create = optional
	NonSysUser string `json:"nonSysUser,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference of the source repository.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-oracle-install.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// The unique name.
	// create = required
	// pattern = ^[a-zA-Z0-9_$#]+$
	// maxLength = 30
	UniqueName string `json:"uniqueName,omitempty"`
	// The list of RAC instances for this RAC configuration.
	// create = required
	// update = optional
	Instances []*OracleRACInstance `json:"instances,omitempty"`
	// The Oracle Clusterware database name.
	// create = optional
	// maxLength = 30
	CrsDatabaseName string `json:"crsDatabaseName,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleService `json:"services,omitempty"`
	// The password of the database user. This must be a PasswordCredential
	// instance.
	// update = optional
	Credentials *PasswordCredential `json:"credentials,omitempty"`
	// The password of a database user that does not have administrative
	// privileges.
	// create = optional
	// update = optional
	NonSysCredentials string `json:"nonSysCredentials,omitempty"`
	// The container type of this database.
	// enum = [UNKNOWN ROOT_CDB NON_CDB AUX_CDB]
	// create = readonly
	// update = readonly
	CdbType string `json:"cdbType,omitempty"`
}

// PgSQLAttachData - Represents the PostgreSQL specific parameters of an attach request.
// extends AttachData
type PgSQLAttachData struct {
	// The external file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// The Postgres installation on the PPT environment that will be used for
	// pre-provisioning.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-pgsql-install.json
	PptRepository string `json:"pptRepository,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// maxLength = 256
	// default = postgres
	// create = optional
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// Information about the host OS user on the PPT host to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	PptHostUser string `json:"pptHostUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The username of the database cluster user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The credential of the database cluster user.
	// create = optional
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
}

// AppDataExportParameters - The parameters to use as input to export AppData.
// extends ExportParameters
type AppDataExportParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
	// The filesystem configuration of the exported database.
	// create = optional
	FilesystemLayout *AppDataFilesystemLayout `json:"filesystemLayout,omitempty"`
}

// DatabaseTemplateConfig - Static template configuration information for a given source type.
// extends TypedObject
type DatabaseTemplateConfig struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object type for sources to which this template is applicable.
	// format = type
	SourceType string `json:"sourceType,omitempty"`
	// A list of reserved parameters names that cannot be used in the
	// template.
	Reserved []string `json:"reserved,omitempty"`
}

// HostEnvironmentCreateParameters - The parameters used for the host environment create operation.
// extends SourceEnvironmentCreateParameters
type HostEnvironmentCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUser `json:"primaryUser,omitempty"`
	// The host environment.
	// create = required
	HostEnvironment HostEnvironment `json:"hostEnvironment,omitempty"`
	// The host parameters used to add a host.
	// create = required
	HostParameters HostCreateParameters `json:"hostParameters,omitempty"`
}

// WindowsHost - The representation of a Windows host object.
// extends Host
type WindowsHost struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// The host configuration object associated with the host.
	HostConfiguration *HostConfiguration `json:"hostConfiguration,omitempty"`
	// Profile for escalating user privileges.
	// format = objectReference
	// create = optional
	// update = optional
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	PrivilegeElevationProfile string `json:"privilegeElevationProfile,omitempty"`
	// The port number used to connect to the host via SSH.
	// maximum = 65535
	// create = optional
	// update = optional
	// default = 22
	// minimum = 1
	SshPort int `json:"sshPort,omitempty"`
	// The date the host was added.
	DateAdded string `json:"dateAdded,omitempty"`
	// The port that the connector connects on.
	// create = optional
	// update = optional
	ConnectorPort float64 `json:"connectorPort,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Runtime properties for this host.
	HostRuntime *HostRuntime `json:"hostRuntime,omitempty"`
	// The address associated with the host.
	// format = host
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// Unique per Delphix key used to authenticate with the remote Delphix
	// Connector.
	// create = optional
	// update = optional
	ConnectorAuthenticationKey string `json:"connectorAuthenticationKey,omitempty"`
	// The path for the toolkit that resides on the host.
	// create = readonly
	// update = readonly
	ToolkitPath string `json:"toolkitPath,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// PgSQLProvisionParameters - The parameters to use as input to provision PostgreSQL databases.
// extends ProvisionParameters
type PgSQLProvisionParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *PgSQLDatabaseContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *PgSQLVirtualSource `json:"source,omitempty"`
	// The source config for the source.
	// required = true
	SourceConfig *PgSQLDBClusterConfig `json:"sourceConfig,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
}

// Schedule - Represent a schedule in the Delphix system.
// extends TypedObject
type Schedule struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Schedule cron string. See CronTrigger documentation at
	// http://quartz-scheduler.org/ for details.
	// maxLength = 120
	// create = required
	// update = required
	CronString string `json:"cronString,omitempty"`
	// Cutoff time in seconds. The policy job will suspend if not completed
	// within the given time limit.
	// units = sec
	// create = optional
	// update = optional
	CutoffTime int `json:"cutoffTime,omitempty"`
}

// TimeflowRange - Time range within a TimeFlow.
// extends TypedObject
type TimeflowRange struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The starting TimeFlow point of this range.
	StartPoint TimeflowPoint `json:"startPoint,omitempty"`
	// The ending TimeFlow point of this range.
	EndPoint TimeflowPoint `json:"endPoint,omitempty"`
	// Whether or not this TimeFlow range is provisionable.
	Provisionable *bool `json:"provisionable,omitempty"`
}

// JSBookmarkCheckoutCount - The number of times a bookmark has been checked out. This means it was
// used as input to a RESTORE, CREATE_BRANCH, or RESET operation. The
// bookmark checkout count is kept separately on replicated templates.
// extends PersistentObject
type JSBookmarkCheckoutCount struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The bookmark that this checkout count is associated with.
	// referenceTo = /delphix-js-bookmark.json
	// format = objectReference
	Bookmark string `json:"bookmark,omitempty"`
	// The number of times the bookmark has been checked out. This means it
	// was used as input to a RESTORE, CREATE_BRANCH, or RESET operation.
	// This should not be replicated.
	CheckoutCount int `json:"checkoutCount,omitempty"`
}

// SourceRepositoryTemplate - The representation of a repository template object.
// extends NamedUserObject
type SourceRepositoryTemplate struct {
	// The reference to the associated template.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-database-template.json
	Template string `json:"template,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The reference to the target repository.
	// referenceTo = /delphix-source-repository.json
	// format = objectReference
	// create = required
	Repository string `json:"repository,omitempty"`
	// The reference to the database container.
	// referenceTo = /delphix-container.json
	// format = objectReference
	// create = required
	Container string `json:"container,omitempty"`
}

// MSSqlTimeflowPoint - A unique point within an MSSql database TimeFlow.
// extends TimeflowPoint
type MSSqlTimeflowPoint struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-mssql-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
}

// JSDataContainerCreateWithRefreshParameters - The parameters used to create a data container when refreshing data
// sources.
// extends JSDataContainerCreateParameters
type JSDataContainerCreateWithRefreshParameters struct {
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParameters `json:"dataSources,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// A reference to the list of users that own this data container.
	// create = optional
	Owners []string `json:"owners,omitempty"`
	// A reference to the template that this data container is provisioned
	// from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	// required = true
	Template string `json:"template,omitempty"`
	// Create the data container with initial data specified by this Jet
	// Stream timeline point.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
	// A description of this data layout to define what it is used for.
	// create = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
}

// NetworkThroughputTestParameters - Parameters used to execute a network throughput test.
// extends NetworkThroughputTestBaseParameters
type NetworkThroughputTestParameters struct {
	// The number of connections to use for the test. The special value 0
	// (the default) causes the test to automatically discover and use the
	// optimal number of connections to use for maximum throughput.
	// minimum = 0
	// maximum = 32
	// default = 0
	// create = optional
	NumConnections int `json:"numConnections,omitempty"`
	// The TCP port number that the server (the receiver) will be listening
	// on.
	// maximum = 65535
	// create = optional
	// minimum = 1
	Port int `json:"port,omitempty"`
	// The size of each transmit request in bytes.
	// maximum = 1.048576e+06
	// base = 1024
	// units = B
	// create = optional
	// default = 131072
	// minimum = 0
	BlockSize int `json:"blockSize,omitempty"`
	// The size of the send socket buffer in bytes.
	// create = optional
	// default = 4.194304e+06
	// minimum = 0
	// maximum = 1.6777216e+07
	// base = 1024
	// units = B
	SendSocketBuffer int `json:"sendSocketBuffer,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = required
	RemoteHost string `json:"remoteHost,omitempty"`
	// The duration of the test in seconds. Note that when numConnections is
	// 0, an initial period of time will be spent calculating the optimal
	// number of connections, and that time does not count toward the
	// duration of the test.
	// default = 30
	// create = optional
	// minimum = 1
	// maximum = 3600
	Duration int `json:"duration,omitempty"`
	// Whether the test is a transmit or receive test.
	// create = optional
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	Direction string `json:"direction,omitempty"`
}

// ASELinkedSource - A linked SAP ASE source.
// extends ASESource
type ASELinkedSource struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The credential for the source DB user.
	// create = optional
	// update = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Runtime properties of this source.
	Runtime *ASESourceRuntime `json:"runtime,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with the
	// source database.
	// update = optional
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	// update = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-ase-db-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// The staging source for validated sync of the database.
	// format = objectReference
	// referenceTo = /delphix-ase-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// Source database backup location.
	// create = required
	// update = optional
	// maxLength = 1024
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
}

// ASECompatibilityCriteria - The compatibility criteria to use for filtering the list of available
// SAP ASE repositories.
// extends CompatibilityCriteria
type ASECompatibilityCriteria struct {
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
}

// BatchContainerDeleteParameters - The parameters to use as input to batch container delete requests.
// extends TypedObject
type BatchContainerDeleteParameters struct {
	// Optional parameters to the delete operations.
	// required = false
	DeleteParameters *DeleteParameters `json:"deleteParameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Containers to delete.
	// minItems = 1
	// required = true
	Containers []string `json:"containers,omitempty"`
}

// PasswordPolicy - Password policies for Delphix users.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type PasswordPolicy struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Name of password policy.
	// create = required
	// update = optional
	// minLength = 1
	// maxLength = 64
	Name string `json:"name,omitempty"`
	// True if password must contain at least one symbol.
	// create = required
	// update = optional
	Symbol *bool `json:"symbol,omitempty"`
	// True to disallow password containing username.
	// create = required
	// update = optional
	DisallowUsernameAsPassword *bool `json:"disallowUsernameAsPassword,omitempty"`
	// The password may not be the same as any of previous n passwords.
	// minimum = 0
	// maximum = 1
	// create = required
	// update = optional
	ReuseDisallowLimit int `json:"reuseDisallowLimit,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// True if password must contain at least one uppercase letter.
	// create = required
	// update = optional
	UppercaseLetter *bool `json:"uppercaseLetter,omitempty"`
	// True if password must contain at least one lowercase letter.
	// create = required
	// update = optional
	LowercaseLetter *bool `json:"lowercaseLetter,omitempty"`
	// True if password must contain at least one digit.
	// create = required
	// update = optional
	Digit *bool `json:"digit,omitempty"`
	// Minimum length for the password.
	// minimum = 1
	// maximum = 128
	// create = required
	// update = optional
	MinLength int `json:"minLength,omitempty"`
}

// JSDataContainerLockParameters - The parameters used to lock a data container.
// extends TypedObject
type JSDataContainerLockParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the user object who locks the container.
	// referenceTo = /delphix-user.json
	// required = true
	// format = objectReference
	LockUser string `json:"lockUser,omitempty"`
}

// WindowsClusterNode - A node in a Windows Cluster.
// extends NamedUserObject
type WindowsClusterNode struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the Windows cluster environment this node belongs to.
	// referenceTo = /delphix-windows-cluster.json
	// format = objectReference
	Cluster string `json:"cluster,omitempty"`
	// A reference to the physical host.
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Indicates whether the node is discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// JSDataChild - A branch with data from a specific bookmark or PIT (point in time).
// extends TypedObject
type JSDataChild struct {
	// Reference to the branch. This will be null if the user does not have
	// permission to see it.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// The name of the branch.
	// maxLength = 256
	BranchName string `json:"branchName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the container.
	// maxLength = 256
	ContainerName string `json:"containerName,omitempty"`
	// The operation performed.
	// enum = [REFRESH RESTORE CREATE_BRANCH RESET]
	Operation string `json:"operation,omitempty"`
	// The username of the owner of the branch. This will be null if there is
	// no owner.
	// maxLength = 256
	UserName string `json:"userName,omitempty"`
}

// AlertActionEmailUser - Alert action that sends email to the email address associated with the
// user.
// extends AlertActionEmail
type AlertActionEmailUser struct {
	// Email format to use. The HTML format will generate a multipart message
	// containing both HTML and plain text. The TEXT format will explicitly
	// generate text-only mail. The JSON format will generate a JSON object
	// identical to the $Alert format returned through the web services API.
	// create = optional
	// update = optional
	// enum = [HTML TEXT JSON]
	// default = HTML
	Format string `json:"format,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSSourceDataTimestamp - The association between a Jet Stream data source and a point in time
// that's provisionable.
// extends TypedObject
type JSSourceDataTimestamp struct {
	// A reference to the Jet Stream data source.
	// format = objectReference
	// referenceTo = /delphix-js-data-source.json
	Source string `json:"source,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the Jet Stream data source.
	Name string `json:"name,omitempty"`
	// The priority of the Jet Stream data source.
	Priority int `json:"priority,omitempty"`
	// The point in the source's dataset time which is provisionable.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// A reference to the Jet Stream branch.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
}

// DataResult - Result of a successful API call containing a reference to a
// downloadable resource.
// extends OKResult
type DataResult struct {
	// Reference to the action associated with the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-action.json
	Action string `json:"action,omitempty"`
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
}

// MySQLGtidCoordinates - The GTID coordinates on the master when a MySQL backup was taken.
// extends MySQLReplicationCoordinates
type MySQLGtidCoordinates struct {
	// The GTID coordinates on the master to start replication from.
	// minItems = 1
	// required = true
	Gtids []string `json:"gtids,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleTimeflow - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends Timeflow
type OracleTimeflow struct {
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *OracleTimeflowPoint `json:"parentPoint,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// Reference to the TimeFlow of the warehouse that this TimeFlow is a
	// part of.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	// featureFlag = CONSPRO
	WarehouseTimeflow string `json:"warehouseTimeflow,omitempty"`
	// Set to true if the TimeFlow represents a warehouse.
	// featureFlag = CONSPRO
	Warehouse *bool `json:"warehouse,omitempty"`
	// Oracle-specific incarnation identifier for this TimeFlow.
	IncarnationID string `json:"incarnationID,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// referenceTo = /delphix-timeflow-snapshot.json
	// format = objectReference
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// Reference to the mirror CDB TimeFlow if this is a PDB TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-oracle-timeflow.json
	CdbTimeflow string `json:"cdbTimeflow,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
}

// SshVerifyFingerprint - SSH verification strategy based on a known per-host fingerprint.
// extends SshVerifyBase
type SshVerifyFingerprint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of ssh key.
	// enum = [RSA DSA ECDSA ED25519]
	// required = true
	KeyType string `json:"keyType,omitempty"`
	// Hash function for the fingerprint.
	// enum = [SHA256 SHA512]
	// required = true
	FingerprintType string `json:"fingerprintType,omitempty"`
	// Base-64 encoded fingerprint of the ssh key of the host.
	// format = hostFingerprint
	// required = true
	Fingerprint string `json:"fingerprint,omitempty"`
}

// VerifyVersionParameters - The parameters to use as input to verifying an upgrade.
// extends TypedObject
type VerifyVersionParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If true, validates the success of upgrading the Delphix Engine without
	// updating the OSsoftware. The operation will detect a failure if the
	// upgrade version requires a version of the OSthat is newer than what is
	// currently running.
	// create = optional
	// default = false
	Defer *bool `json:"defer,omitempty"`
}

// MSSqlClusterInstance - The representation of a SQL Server Instance on a clustered node for
// Availability Groups.
// extends MSSqlBaseClusterInstance
type MSSqlClusterInstance struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// A reference to the Windows Cluster Node for this instance.
	// format = objectReference
	// referenceTo = /delphix-windows-cluster-node.json
	Node string `json:"node,omitempty"`
	// The port to connect to the SQL Server Instance.
	Port int `json:"port,omitempty"`
	// The name of the SQL Server Instance.
	// maxLength = 16
	Name string `json:"name,omitempty"`
	// The version of the SQL Server Instance.
	Version string `json:"version,omitempty"`
	// The owner of the SQL Server Instance.
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// The Servername of the SQL Server Instance.
	ServerName string `json:"serverName,omitempty"`
}

// AppDataVirtualSource - A virtual AppData source.
// extends AppDataManagedSource
type AppDataVirtualSource struct {
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Locations to mount subdirectories of the AppData in addition to the
	// normal target mount point. These paths will be mounted and unmounted
	// as part of enabling and disabling this source.
	// update = optional
	// create = optional
	AdditionalMountPoints []*AppDataAdditionalMountPoint `json:"additionalMountPoints,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Runtime properties of this source.
	Runtime *AppDataSourceRuntime `json:"runtime,omitempty"`
	// Reference to the configuration for the source.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	Config string `json:"config,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// The toolkit associated with this source.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// create = required
	// update = optional
	Parameters *Json `json:"parameters,omitempty"`
}

// RefreshPolicy - This policy refreshes a container according to a schedule.
// extends SchedulePolicy
type RefreshPolicy struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// List of Schedule objects representing when the policy will execute.
	// create = required
	// update = optional
	ScheduleList []*Schedule `json:"scheduleList,omitempty"`
	// Provision source, either the latest time or latest snapshot.
	// enum = [LATEST_SNAPSHOT LATEST_TIME_FLOW_LOG]
	// create = required
	// update = optional
	ProvisionSource string `json:"provisionSource,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// True if this is the default policy created when the system is setup.
	// Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// Whether this policy has been directly applied or inherited. See the
	// effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// The timezone of this policy. If not specified, defaults to the Delphix
	// Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZone `json:"timezone,omitempty"`
}

// AppDataSourceRuntime - Runtime (non-persistent) properties of an AppData source.
// extends SourceRuntime
type AppDataSourceRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
}

// Authorization - Describes a role as applied to a user on an object.
// extends ReadonlyNamedUserObject
type Authorization struct {
	// Applied role.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-role.json
	Role string `json:"role,omitempty"`
	// Reference to the object that the authorization applies to.
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	// create = required
	Target string `json:"target,omitempty"`
	// Reference to the user that the authorization applies to.
	// format = objectReference
	// create = required
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
}

// ASETimeflowPoint - A unique point within a SAP ASE database TimeFlow.
// extends TimeflowPoint
type ASETimeflowPoint struct {
	// The TimeFlow location.
	// create = optional
	Location string `json:"location,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The logical time corresponding to the TimeFlow location.
	// format = date
	// create = optional
	Timestamp string `json:"timestamp,omitempty"`
	// Reference to TimeFlow containing this point.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-ase-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
}

// JSTimelinePointTimeInput - Specifies a point in time on the Jet Stream timeline for a specific
// branch. Latest provisionable points before the specified time will be
// used.
// extends JSTimelinePointTimeParameters
type JSTimelinePointTimeInput struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A point in time on the given branch.
	// format = date
	// required = true
	Time string `json:"time,omitempty"`
	// The reference to the branch used for this operation.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// required = true
	Branch string `json:"branch,omitempty"`
}

// JobEvent - Represents a job event object. This can either be a state change or a
// progress update.
// extends TypedObject
type JobEvent struct {
	// Command output associated with the event, if applicable.
	MessageCommandOutput string `json:"messageCommandOutput,omitempty"`
	// Localized message action.
	MessageAction string `json:"messageAction,omitempty"`
	// New state of the job.
	// enum = [INITIAL RUNNING SUSPENDED CANCELED COMPLETED FAILED RETRYABLE]
	State string `json:"state,omitempty"`
	// Message ID associated with the event.
	MessageCode string `json:"messageCode,omitempty"`
	// Results of diagnostic checks run, if any, if the job failed.
	Diagnoses []*DiagnosisResult `json:"diagnoses,omitempty"`
	// Time the event occurred.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Completion percentage.
	// units = %
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// Localized message details.
	MessageDetails string `json:"messageDetails,omitempty"`
	// Type of event.
	// enum = [INFO WARNING ERROR]
	EventType string `json:"eventType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CreateMaskingJobTransformationParameters - Represents the parameters to create a transformation for a provided
// Masking Job.
// extends TypedObject
type CreateMaskingJobTransformationParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// The container that will contain the masked data associated with the
	// newly created transformation; the "transformation container".
	// create = required
	Container Container `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
}

// MSSqlCreateTransformationParameters - Represents the parameters of a createTransformation request for a
// MSSql container.
// extends CreateTransformationParameters
type MSSqlCreateTransformationParameters struct {
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *MSSqlDatabaseContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleStartParameters - The parameters to use as input to start oracle sources.
// extends SourceStartParameters
type OracleStartParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// List of specific Oracle instances to start.
	Instances []float64 `json:"instances,omitempty"`
}

// PgSQLSourceRuntime - Runtime (non-persistent) properties of a PostgreSQL source.
// extends SourceRuntime
type PgSQLSourceRuntime struct {
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
}

// JSDataContainerRefreshParameters - The parameters used to refresh a data container.
// extends TypedObject
type JSDataContainerRefreshParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
}

// JSOperationEndpointBranchParameters - The branch to fetch the first and last event from.
// extends JSOperationEndpointParameters
type JSOperationEndpointBranchParameters struct {
	// The branch to search.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// create = required
	Branch string `json:"branch,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// DSPAutotunerParameters - Network information required by the DSP autotuner.
// extends TypedObject
type DSPAutotunerParameters struct {
	// Whether the test is a transmit or receive test.
	// create = optional
	// enum = [TRANSMIT RECEIVE]
	// default = TRANSMIT
	Direction string `json:"direction,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	RemoteHost string `json:"remoteHost,omitempty"`
	// Address, username and password used when running a test to another
	// Delphix Engine.
	// create = optional
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfo `json:"remoteDelphixEngineInfo,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether the test is testing connectivity to a Delphix Engine or a
	// remote host.
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	// create = required
	DestinationType string `json:"destinationType,omitempty"`
}

// EnvironmentUser - The representation of an environment user object.
// extends UserObject
type EnvironmentUser struct {
	// User ID of the user.
	// create = optional
	// update = optional
	// minimum = 0
	// maximum = 4.294967295e+09
	UserId int `json:"userId,omitempty"`
	// The credential for the environment user.
	// properties = map[type:map[description:Object type. required:true format:type default:PasswordCredential type:string]]
	// create = required
	// update = optional
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the associated environment.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Group ID of the user.
	// minimum = 0
	// maximum = 4.294967295e+09
	// create = optional
	// update = optional
	GroupId int `json:"groupId,omitempty"`
}

// PgSQLPlatformParameters - PgSQL platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type PgSQLPlatformParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The port number used for provisioning a PgSQL container during
	// transformation application. This port must be available when the
	// transformation is applied so that the temporary VDB created during the
	// transformation process can start and listen to this port.
	// maximum = 65535
	// minimum = 1
	Port int `json:"port,omitempty"`
}

// AlertProfile - A profile that describes a set of actions to take in response to an
// alert being generated.
// extends PersistentObject
type AlertProfile struct {
	// List of actions to take. Only alerts visible to the user and matching
	// the optional filters are included. If there are multiple actions with
	// the same result (such as emailing a user), only one result is acted
	// upon.
	// create = required
	// update = optional
	Actions []AlertAction `json:"actions,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Specifies which alerts should be matched by this profile.
	// update = optional
	// properties = map[type:map[default:SeverityFilter]]
	// create = optional
	FilterSpec AlertFilter `json:"filterSpec,omitempty"`
}

// FaultResolveParameters - The parameters to use as input when marking a fault as resolved.
// extends TypedObject
type FaultResolveParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether to ignore this fault if it is detected on the
	// same object in the future.
	// create = optional
	// default = false
	Ignore *bool `json:"ignore,omitempty"`
	// The comments describing the steps taken to resolve a fault.
	// create = optional
	// default =
	Comments string `json:"comments,omitempty"`
}

// NetworkDSPTest - DSP throughput tests to a target system.
// extends NetworkThroughputTestBase
// cliVisibility = [DOMAIN SYSTEM]
type NetworkDSPTest struct {
	// Time when the test ended.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Number of connections used to achieve maximum sustained throughput.
	NumConnections int `json:"numConnections,omitempty"`
	// Average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The state of the test.
	// enum = [RUNNING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// Time when the test was started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The parameters used to execute the test.
	Parameters *NetworkDSPTestParameters `json:"parameters,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The remote IP address used for the test.
	// format = ipAddress
	RemoteAddress string `json:"remoteAddress,omitempty"`
}

// MySQLCreateTransformationParameters - Represents the parameters of a createTransformation request for a
// MySQL container.
// extends CreateTransformationParameters
type MySQLCreateTransformationParameters struct {
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *MySQLDatabaseContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The port number used for provisioning a MySQL container during
	// transformation application.
	// create = required
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleClusterNode - The representation of an oracle cluster node object.
// extends NamedUserObject
type OracleClusterNode struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The reference to the parent cluster environment.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster.json
	Cluster string `json:"cluster,omitempty"`
	// The list of virtual IPs belonging to this node.
	// update = optional
	VirtualIPs []*OracleVirtualIP `json:"virtualIPs,omitempty"`
	// Indicates whether the node is discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The reference to the associated host object.
	// format = objectReference
	// referenceTo = /delphix-host.json
	Host string `json:"host,omitempty"`
	// Indicates whether the node is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
}

// MySQLProvisionParameters - The parameters to use as input to provision requests for MySQL
// databases.
// extends ProvisionParameters
type MySQLProvisionParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the provisioned database.
	// required = true
	Container *MySQLDatabaseContainer `json:"container,omitempty"`
	// The source that describes an external database instance.
	// required = true
	Source *MySQLVirtualSource `json:"source,omitempty"`
	// The source config for the source.
	// required = true
	SourceConfig *MySQLServerConfig `json:"sourceConfig,omitempty"`
	// The Masking Job to be run when this dataset is provisioned or
	// refreshed.
	// format = objectReference
	// referenceTo = /delphix-masking-job.json
	// create = optional
	// update = readonly
	MaskingJob string `json:"maskingJob,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// TimeflowRepairSSH - Parameters to repair log files within a TimeFlow.
// extends TimeflowRepairParameters
type TimeflowRepairSSH struct {
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// User credentials. If not provided will use environment credentials for
	// 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port int `json:"port,omitempty"`
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// The starting point of the range of log files to fetch.
	// required = true
	StartLocation string `json:"startLocation,omitempty"`
	// The ending point of the range of log files to fetch.
	// required = true
	EndLocation string `json:"endLocation,omitempty"`
}

// Toolkit - An installed toolkit.
// extends NamedUserObject
type Toolkit struct {
	// A unique and descriptive name for the toolkit.
	// required = true
	// maxLength = 256
	// pattern = ^[a-z0-9_:-]+$
	Name string `json:"name,omitempty"`
	// The version of the toolkit that is of the form: 'major.minor.patch'.
	// format = toolkitVersion
	// required = true
	Version string `json:"version,omitempty"`
	// A list of host types compatible with this toolkit.
	// required = true
	HostTypes []string `json:"hostTypes,omitempty"`
	// A human readable name for the toolkit.
	// required = true
	// maxLength = 256
	PrettyName string `json:"prettyName,omitempty"`
	// Definition of how to discover sources of this type.
	// required = false
	DiscoveryDefinition *ToolkitDiscoveryDefinition `json:"discoveryDefinition,omitempty"`
	// Definition of how to link sources of this type.
	// required = true
	LinkedSourceDefinition ToolkitLinkedSource `json:"linkedSourceDefinition,omitempty"`
	// The Delphix API version that the toolkit was built against.
	// required = true
	BuildApi *APIVersion `json:"buildApi,omitempty"`
	// Resources for use by workflows in this toolkit.
	// required = true
	Resources map[string]string `json:"resources,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Definition of how to provision virtual sources of this type.
	// required = true
	VirtualSourceDefinition *ToolkitVirtualSource `json:"virtualSourceDefinition,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Definition of how to upgrade sources of this type.
	// required = false
	UpgradeDefinition *ToolkitUpgradeDefinition `json:"upgradeDefinition,omitempty"`
	// Schema for metadata collected during snapshotting.
	// required = true
	SnapshotSchema *SchemaDraftV4 `json:"snapshotSchema,omitempty"`
	// The default locale for this toolkit. This locale defines the set of
	// all message IDs for the toolkit and serves as the fallback locale when
	// messages cannot be localized in a particular locale. If no messages
	// are specified for the toolkit, the defaultLocale may be any locale.
	// format = locale
	// required = true
	DefaultLocale string `json:"defaultLocale,omitempty"`
	// The set of localizable messages for this toolkit.
	// required = false
	Messages []*ToolkitLocale `json:"messages,omitempty"`
}

// SnapshotRuntime - Runtime properties of a TimeFlow snapshot.
// extends TypedObject
type SnapshotRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
}

// DatapointSet - A set of datapoints from a particular analytics slice.
// extends TypedObject
type DatapointSet struct {
	// The amount of time each datapoint spans.
	// units = sec
	Resolution int `json:"resolution,omitempty"`
	// The set of datapoint streams in the result.
	DatapointStreams []DatapointStream `json:"datapointStreams,omitempty"`
	// True if the number of datapoints to be included exceeded the maximum
	// allowable for a single datapoint set. As a result, not all datapoints
	// could be included in the datapointStreams array.
	Overflow *bool `json:"overflow,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AuthGetByPropertiesParameters - The parameters to use as input to get an authorization with given
// user, target, and role.
// extends TypedObject
type AuthGetByPropertiesParameters struct {
	// The role type authorizations are applied to.
	// referenceTo = /delphix-role.json
	// format = objectReference
	// required = true
	Role string `json:"role,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The user authorizations are applied to.
	// format = objectReference
	// required = true
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// The target authorizations are applied to.
	// referenceTo = /delphix-user-object.json
	// format = objectReference
	// required = true
	Target string `json:"target,omitempty"`
}

// RefreshParameters - The parameters to use as input to refresh requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type RefreshParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to refresh the
	// database to.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// GetCurrentPasswordPolicyParameters - The parameters to use as input when getting the currently active
// password policy.
// extends TypedObject
type GetCurrentPasswordPolicyParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Type of user.
	// required = true
	// enum = [SYSTEM DOMAIN]
	UserType string `json:"userType,omitempty"`
}

// NetworkInterfaceUtilDatapointStream - A stream of datapoints from a NETWORK_INTERFACE_UTIL analytics slice.
// extends DatapointStream
type NetworkInterfaceUtilDatapointStream struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Which network interface was utilized.
	NetworkInterface string `json:"networkInterface,omitempty"`
}

// OracleDatabaseStatistic - A row in the database performance statistic table.
// extends TypedObject
type OracleDatabaseStatistic struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A single performance statistic row.
	StatisticValues []string `json:"statisticValues,omitempty"`
}

// ListResult - Result of a successful API call returning a list.
// extends OKResult
type ListResult struct {
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// Reference to the action associated with the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-action.json
	Action string `json:"action,omitempty"`
	// The number of items in the entire result set, regardless of the
	// requested page size. For some operations, this value is null.
	Total int `json:"total,omitempty"`
	// True if the total number of matching items is too large to be
	// calculated.
	Overflow *bool `json:"overflow,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
}

// PurgeLogsResult - Represents the result of a purgeLogs operation.
// extends TypedObject
type PurgeLogsResult struct {
	// TimeFlow point after the last snapshot beyond which TimeFlow will be
	// lost as a result of purging logs.
	TruncatePoint *OracleTimeflowPoint `json:"truncatePoint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// List of snapshots which have been rendered unprovisionable because
	// logs needed to make them consistent have been deleted.
	AffectedSnapshots []TimeflowSnapshot `json:"affectedSnapshots,omitempty"`
}

// JSDataContainerUndoParameters - The parameters used to undo an operation on a data container.
// extends TypedObject
type JSDataContainerUndoParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The operation to undo. This is only valid for RESET, RESTORE, UNDO,
	// and REFRESH operations.
	// format = objectReference
	// referenceTo = /delphix-js-operation.json
	// required = true
	Operation string `json:"operation,omitempty"`
}

// NetworkRouteLookupParameters - Parameters used for a routing table lookup.
// extends TypedObject
type NetworkRouteLookupParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Destination address or hostname.
	// format = host
	// create = required
	Destination string `json:"destination,omitempty"`
}

// PolicyCreateAndApplyParameters - The parameters used for creating and applying policies.
// extends TypedObject
type PolicyCreateAndApplyParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Policy to create.
	// required = true
	Policy Policy `json:"policy,omitempty"`
	// Object reference of the target.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-user-object.json
	Target string `json:"target,omitempty"`
}

// OracleDisableParameters - The parameters to use as input to disable oracle sources.
// extends SourceDisableParameters
type OracleDisableParameters struct {
	// Whether to attempt a cleanup of the database from the environment
	// before the disable.
	// default = true
	AttemptCleanup *bool `json:"attemptCleanup,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// TimeZone - This represents a time zone offset.
// extends TypedObject
type TimeZone struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The difference, in minutes, between UTC and local time. For example,
	// if your time zone is UTC -5:00 (Eastern Standard Time), 300 will be
	// returned. Daylight saving time prevents this value from being a
	// constant even for a given locale.
	Offset int `json:"offset,omitempty"`
	// The Offset as a String. For instance 'UTC -5:00'.
	OffsetString string `json:"offsetString,omitempty"`
	// The ID of this time zone.
	// required = true
	// enum = [ACT AET AGT ART AST Africa/Abidjan Africa/Accra Africa/Addis_Ababa Africa/Algiers Africa/Asmara Africa/Asmera Africa/Bamako Africa/Bangui Africa/Banjul Africa/Bissau Africa/Blantyre Africa/Brazzaville Africa/Bujumbura Africa/Cairo Africa/Casablanca Africa/Ceuta Africa/Conakry Africa/Dakar Africa/Dar_es_Salaam Africa/Djibouti Africa/Douala Africa/El_Aaiun Africa/Freetown Africa/Gaborone Africa/Harare Africa/Johannesburg Africa/Juba Africa/Kampala Africa/Khartoum Africa/Kigali Africa/Kinshasa Africa/Lagos Africa/Libreville Africa/Lome Africa/Luanda Africa/Lubumbashi Africa/Lusaka Africa/Malabo Africa/Maputo Africa/Maseru Africa/Mbabane Africa/Mogadishu Africa/Monrovia Africa/Nairobi Africa/Ndjamena Africa/Niamey Africa/Nouakchott Africa/Ouagadougou Africa/Porto-Novo Africa/Sao_Tome Africa/Timbuktu Africa/Tripoli Africa/Tunis Africa/Windhoek America/Adak America/Anchorage America/Anguilla America/Antigua America/Araguaina America/Argentina/Buenos_Aires America/Argentina/Catamarca America/Argentina/ComodRivadavia America/Argentina/Cordoba America/Argentina/Jujuy America/Argentina/La_Rioja America/Argentina/Mendoza America/Argentina/Rio_Gallegos America/Argentina/Salta America/Argentina/San_Juan America/Argentina/San_Luis America/Argentina/Tucuman America/Argentina/Ushuaia America/Aruba America/Asuncion America/Atikokan America/Atka America/Bahia America/Bahia_Banderas America/Barbados America/Belem America/Belize America/Blanc-Sablon America/Boa_Vista America/Bogota America/Boise America/Buenos_Aires America/Cambridge_Bay America/Campo_Grande America/Cancun America/Caracas America/Catamarca America/Cayenne America/Cayman America/Chicago America/Chihuahua America/Coral_Harbour America/Cordoba America/Costa_Rica America/Creston America/Cuiaba America/Curacao America/Danmarkshavn America/Dawson America/Dawson_Creek America/Denver America/Detroit America/Dominica America/Edmonton America/Eirunepe America/El_Salvador America/Ensenada America/Fort_Nelson America/Fort_Wayne America/Fortaleza America/Glace_Bay America/Godthab America/Goose_Bay America/Grand_Turk America/Grenada America/Guadeloupe America/Guatemala America/Guayaquil America/Guyana America/Halifax America/Havana America/Hermosillo America/Indiana/Indianapolis America/Indiana/Knox America/Indiana/Marengo America/Indiana/Petersburg America/Indiana/Tell_City America/Indiana/Vevay America/Indiana/Vincennes America/Indiana/Winamac America/Indianapolis America/Inuvik America/Iqaluit America/Jamaica America/Jujuy America/Juneau America/Kentucky/Louisville America/Kentucky/Monticello America/Knox_IN America/Kralendijk America/La_Paz America/Lima America/Los_Angeles America/Louisville America/Lower_Princes America/Maceio America/Managua America/Manaus America/Marigot America/Martinique America/Matamoros America/Mazatlan America/Mendoza America/Menominee America/Merida America/Metlakatla America/Mexico_City America/Miquelon America/Moncton America/Monterrey America/Montevideo America/Montreal America/Montserrat America/Nassau America/New_York America/Nipigon America/Nome America/Noronha America/North_Dakota/Beulah America/North_Dakota/Center America/North_Dakota/New_Salem America/Ojinaga America/Panama America/Pangnirtung America/Paramaribo America/Phoenix America/Port-au-Prince America/Port_of_Spain America/Porto_Acre America/Porto_Velho America/Puerto_Rico America/Punta_Arenas America/Rainy_River America/Rankin_Inlet America/Recife America/Regina America/Resolute America/Rio_Branco America/Rosario America/Santa_Isabel America/Santarem America/Santiago America/Santo_Domingo America/Sao_Paulo America/Scoresbysund America/Shiprock America/Sitka America/St_Barthelemy America/St_Johns America/St_Kitts America/St_Lucia America/St_Thomas America/St_Vincent America/Swift_Current America/Tegucigalpa America/Thule America/Thunder_Bay America/Tijuana America/Toronto America/Tortola America/Vancouver America/Virgin America/Whitehorse America/Winnipeg America/Yakutat America/Yellowknife Antarctica/Casey Antarctica/Davis Antarctica/DumontDUrville Antarctica/Macquarie Antarctica/Mawson Antarctica/McMurdo Antarctica/Palmer Antarctica/Rothera Antarctica/South_Pole Antarctica/Syowa Antarctica/Troll Antarctica/Vostok Arctic/Longyearbyen Asia/Aden Asia/Almaty Asia/Amman Asia/Anadyr Asia/Aqtau Asia/Aqtobe Asia/Ashgabat Asia/Ashkhabad Asia/Atyrau Asia/Baghdad Asia/Bahrain Asia/Baku Asia/Bangkok Asia/Barnaul Asia/Beirut Asia/Bishkek Asia/Brunei Asia/Calcutta Asia/Chita Asia/Choibalsan Asia/Chongqing Asia/Chungking Asia/Colombo Asia/Dacca Asia/Damascus Asia/Dhaka Asia/Dili Asia/Dubai Asia/Dushanbe Asia/Famagusta Asia/Gaza Asia/Harbin Asia/Hebron Asia/Ho_Chi_Minh Asia/Hong_Kong Asia/Hovd Asia/Irkutsk Asia/Istanbul Asia/Jakarta Asia/Jayapura Asia/Jerusalem Asia/Kabul Asia/Kamchatka Asia/Karachi Asia/Kashgar Asia/Kathmandu Asia/Katmandu Asia/Khandyga Asia/Kolkata Asia/Krasnoyarsk Asia/Kuala_Lumpur Asia/Kuching Asia/Kuwait Asia/Macao Asia/Macau Asia/Magadan Asia/Makassar Asia/Manila Asia/Muscat Asia/Nicosia Asia/Novokuznetsk Asia/Novosibirsk Asia/Omsk Asia/Oral Asia/Phnom_Penh Asia/Pontianak Asia/Pyongyang Asia/Qatar Asia/Qyzylorda Asia/Rangoon Asia/Riyadh Asia/Saigon Asia/Sakhalin Asia/Samarkand Asia/Seoul Asia/Shanghai Asia/Singapore Asia/Srednekolymsk Asia/Taipei Asia/Tashkent Asia/Tbilisi Asia/Tehran Asia/Tel_Aviv Asia/Thimbu Asia/Thimphu Asia/Tomsk Asia/Tokyo Asia/Ujung_Pandang Asia/Ulaanbaatar Asia/Ulan_Bator Asia/Urumqi Asia/Ust-Nera Asia/Vientiane Asia/Vladivostok Asia/Yakutsk Asia/Yangon Asia/Yekaterinburg Asia/Yerevan Atlantic/Azores Atlantic/Bermuda Atlantic/Canary Atlantic/Cape_Verde Atlantic/Faeroe Atlantic/Faroe Atlantic/Jan_Mayen Atlantic/Madeira Atlantic/Reykjavik Atlantic/South_Georgia Atlantic/St_Helena Atlantic/Stanley Australia/ACT Australia/Adelaide Australia/Brisbane Australia/Broken_Hill Australia/Canberra Australia/Currie Australia/Darwin Australia/Eucla Australia/Hobart Australia/LHI Australia/Lindeman Australia/Lord_Howe Australia/Melbourne Australia/NSW Australia/North Australia/Perth Australia/Queensland Australia/South Australia/Sydney Australia/Tasmania Australia/Victoria Australia/West Australia/Yancowinna BET BST Brazil/Acre Brazil/DeNoronha Brazil/East Brazil/West CAT CET CNT CST CST6CDT CTT Canada/Atlantic Canada/Central Canada/East-Saskatchewan Canada/Eastern Canada/Mountain Canada/Newfoundland Canada/Pacific Canada/Saskatchewan Canada/Yukon Chile/Continental Chile/EasterIsland Cuba EAT ECT EET EST EST5EDT Egypt Eire Etc/GMT Etc/GMT+0 Etc/GMT+1 Etc/GMT+10 Etc/GMT+11 Etc/GMT+12 Etc/GMT+2 Etc/GMT+3 Etc/GMT+4 Etc/GMT+5 Etc/GMT+6 Etc/GMT+7 Etc/GMT+8 Etc/GMT+9 Etc/GMT-0 Etc/GMT-1 Etc/GMT-10 Etc/GMT-11 Etc/GMT-12 Etc/GMT-13 Etc/GMT-14 Etc/GMT-2 Etc/GMT-3 Etc/GMT-4 Etc/GMT-5 Etc/GMT-6 Etc/GMT-7 Etc/GMT-8 Etc/GMT-9 Etc/GMT0 Etc/Greenwich Etc/UCT Etc/UTC Etc/Universal Etc/Zulu Europe/Amsterdam Europe/Andorra Europe/Astrakhan Europe/Athens Europe/Belfast Europe/Belgrade Europe/Berlin Europe/Bratislava Europe/Brussels Europe/Bucharest Europe/Budapest Europe/Busingen Europe/Chisinau Europe/Copenhagen Europe/Dublin Europe/Gibraltar Europe/Guernsey Europe/Helsinki Europe/Isle_of_Man Europe/Istanbul Europe/Jersey Europe/Kaliningrad Europe/Kiev Europe/Kirov Europe/Lisbon Europe/Ljubljana Europe/London Europe/Luxembourg Europe/Madrid Europe/Malta Europe/Mariehamn Europe/Minsk Europe/Monaco Europe/Moscow Europe/Nicosia Europe/Oslo Europe/Paris Europe/Podgorica Europe/Prague Europe/Riga Europe/Rome Europe/Samara Europe/San_Marino Europe/Sarajevo Europe/Saratov Europe/Simferopol Europe/Skopje Europe/Sofia Europe/Stockholm Europe/Tallinn Europe/Tirane Europe/Tiraspol Europe/Ulyanovsk Europe/Uzhgorod Europe/Vaduz Europe/Vatican Europe/Vienna Europe/Vilnius Europe/Volgograd Europe/Warsaw Europe/Zagreb Europe/Zaporozhye Europe/Zurich GB GB-Eire GMT GMT0 Greenwich HST Hongkong IET IST Iceland Indian/Antananarivo Indian/Chagos Indian/Christmas Indian/Cocos Indian/Comoro Indian/Kerguelen Indian/Mahe Indian/Maldives Indian/Mauritius Indian/Mayotte Indian/Reunion Iran Israel JST Jamaica Japan Kwajalein Libya MET MIT MST MST7MDT Mexico/BajaNorte Mexico/BajaSur Mexico/General NET NST NZ NZ-CHAT Navajo PLT PNT PRC PRT PST PST8PDT Pacific/Apia Pacific/Auckland Pacific/Bougainville Pacific/Chatham Pacific/Chuuk Pacific/Easter Pacific/Efate Pacific/Enderbury Pacific/Fakaofo Pacific/Fiji Pacific/Funafuti Pacific/Galapagos Pacific/Gambier Pacific/Guadalcanal Pacific/Guam Pacific/Honolulu Pacific/Johnston Pacific/Kiritimati Pacific/Kosrae Pacific/Kwajalein Pacific/Majuro Pacific/Marquesas Pacific/Midway Pacific/Nauru Pacific/Niue Pacific/Norfolk Pacific/Noumea Pacific/Pago_Pago Pacific/Palau Pacific/Pitcairn Pacific/Pohnpei Pacific/Ponape Pacific/Port_Moresby Pacific/Rarotonga Pacific/Saipan Pacific/Samoa Pacific/Tahiti Pacific/Tarawa Pacific/Tongatapu Pacific/Truk Pacific/Wake Pacific/Wallis Pacific/Yap Poland Portugal ROK SST Singapore SystemV/AST4 SystemV/AST4ADT SystemV/CST6 SystemV/CST6CDT SystemV/EST5 SystemV/EST5EDT SystemV/HST10 SystemV/MST7 SystemV/MST7MDT SystemV/PST8 SystemV/PST8PDT SystemV/YST9 SystemV/YST9YDT Turkey UCT US/Alaska US/Aleutian US/Arizona US/Central US/East-Indiana US/Eastern US/Hawaii US/Indiana-Starke US/Michigan US/Mountain US/Pacific US/Pacific-New US/Samoa UTC Universal VST W-SU WET Zulu]
	Id string `json:"id,omitempty"`
}

// OracleStopParameters - The parameters to use as input to stop oracle sources.
// extends SourceStopParameters
type OracleStopParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// List of specific Oracle instances to stop.
	Instances []float64 `json:"instances,omitempty"`
	// The name of the privileged user to run the provision operation as.
	Username string `json:"username,omitempty"`
	// The security credential of the privileged user to run the provision
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// Whether to issue 'shutdown abort' to shutdown Oracle instances.
	// default = false
	Abort *bool `json:"abort,omitempty"`
}

// WindowsClusterCreateParameters - The parameters used to create a Windows Cluster.
// extends SourceEnvironmentCreateParameters
type WindowsClusterCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The primary user associated with the environment.
	// create = required
	PrimaryUser *EnvironmentUser `json:"primaryUser,omitempty"`
	// The representation of the cluster object.
	// create = required
	Cluster *WindowsCluster `json:"cluster,omitempty"`
	// Create as a target environment.
	// create = optional
	Target *bool `json:"target,omitempty"`
}

// OracleCustomEnvVarSIFile - Dictates an environment file to be sourced when the Delphix Engine
// administers an Oracle virtual database. This environment file must be
// available on the target environment. This type also includes
// parameters which will be passed to the environment file when it is
// sourced.
// extends OracleCustomEnvVarFile
type OracleCustomEnvVarSIFile struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A string of whitespace-separated parameters to be passed to the source
	// command. The first parameter must be an absolute path to a file that
	// exists on the target environment. Every subsequent parameter will be
	// treated as an argument interpreted by the environment file.
	// required = true
	PathParameters string `json:"pathParameters,omitempty"`
}

// IScsiOpsDatapointStream - A stream of datapoints from a iSCSI_OPS analytics slice.
// extends DatapointStream
type IScsiOpsDatapointStream struct {
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Address of the client.
	// format = host
	Client string `json:"client,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
}

// SyncPolicy - This policy syncs a container (runs SnapSync) according to the given
// schedule.
// extends SchedulePolicy
type SyncPolicy struct {
	// The timezone of this policy. If not specified, defaults to the Delphix
	// Engine's timezone.
	// create = optional
	// update = optional
	Timezone *TimeZone `json:"timezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this policy is customized specifically for one object.
	// Customized policies cannot be shared between objects.
	// create = optional
	// default = false
	Customized *bool `json:"customized,omitempty"`
	// Whether this policy has been directly applied or inherited. See the
	// effectivePolicies parameter of the list call for details.
	// enum = [DIRECT_APPLIED INHERITED]
	EffectiveType string `json:"effectiveType,omitempty"`
	// True if this is the default policy created when the system is setup.
	// Default policies cannot be deleted.
	Default *bool `json:"default,omitempty"`
	// List of Schedule objects representing when the policy will execute.
	// create = required
	// update = optional
	ScheduleList []*Schedule `json:"scheduleList,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
}

// ASECreateTransformationParameters - Represents the parameters of a createTransformation request for an ASE
// container.
// extends CreateTransformationParameters
type ASECreateTransformationParameters struct {
	// Reference to the repository used during application of the
	// transformation.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = required
	Repository string `json:"repository,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *ASEDBContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// create = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
}

// SourcingPolicy - Database policies for managing SnapSync and LogSync across sources for
// a MSSQL container.
// extends TypedObject
type SourcingPolicy struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// True if LogSync should run for this database.
	// default = false
	// create = optional
	// update = optional
	LogsyncEnabled *bool `json:"logsyncEnabled,omitempty"`
}

// OracleSourcingPolicy - Database policies for managing SnapSync and LogSync across sources for
// an Oracle container.
// extends SourcingPolicy
type OracleSourcingPolicy struct {
	// LogSync operation mode for this database.
	// update = optional
	// default = UNDEFINED
	// enum = [ARCHIVE_ONLY_MODE ARCHIVE_REDO_MODE UNDEFINED]
	// create = optional
	LogsyncMode string `json:"logsyncMode,omitempty"`
	// Interval between LogSync requests, in seconds.
	// create = optional
	// update = optional
	// units = sec
	// default = 5
	LogsyncInterval int `json:"logsyncInterval,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if LogSync should run for this database.
	// create = optional
	// update = optional
	// default = false
	LogsyncEnabled *bool `json:"logsyncEnabled,omitempty"`
}

// SamlAuthParameters - The parameter to use as input to determine whether to encode a SAML
// authentication request.
// extends TypedObject
type SamlAuthParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Set to true to encode SAML authentication requests.
	// required = true
	// default = true
	EncodeRequest *bool `json:"encodeRequest,omitempty"`
}

// Action - Represents an action, a permanent record of activity on the server.
// extends PersistentObject
type Action struct {
	// Name of user or policy that initiated the action.
	WorkSourceName string `json:"workSourceName,omitempty"`
	// Name of client software used to initiate the action.
	UserAgent string `json:"userAgent,omitempty"`
	// State of the action.
	// enum = [EXECUTING WAITING COMPLETED FAILED CANCELED]
	State string `json:"state,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Action type.
	ActionType string `json:"actionType,omitempty"`
	// The time the action completed.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Network address used to initiate the action.
	OriginIp string `json:"originIp,omitempty"`
	// Action to be taken to resolve the failure.
	FailureAction string `json:"failureAction,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Action title.
	Title string `json:"title,omitempty"`
	// Plain text description of the action.
	Details string `json:"details,omitempty"`
	// The time the action occurred. For long running processes, this
	// represents the starting time.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// User who initiated the action.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// Message ID associated with the event.
	FailureMessageCode string `json:"failureMessageCode,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The parent action of this action.
	// format = objectReference
	// referenceTo = /delphix-action.json
	ParentAction string `json:"parentAction,omitempty"`
	// Origin of the work that caused the action.
	// enum = [WEBSERVICE POLICY SYSTEM]
	WorkSource string `json:"workSource,omitempty"`
	// Report of progress and warnings for some actions.
	Report string `json:"report,omitempty"`
	// Details of the action failure.
	FailureDescription string `json:"failureDescription,omitempty"`
}

// SamlServiceProvider - SAML service provider configuration.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type SamlServiceProvider struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Algorithm used for creation of digital signature on metadata.
	// required = true
	HashAlgUrl string `json:"hashAlgUrl,omitempty"`
	// URL to which the SAML authentication request will be sent.
	// required = true
	Destination string `json:"destination,omitempty"`
	// A unique identifier that is provided by the identity provider to
	// ensure we are a valid service provider.
	// required = true
	IssuerId string `json:"issuerId,omitempty"`
	// The signing (public) key that will be used to verify SAML signatures.
	// Leave empty if responses will not be signed.
	// create = optional
	// update = optional
	SigningKey string `json:"signingKey,omitempty"`
	// The decryption (private) key that will be used to decrypt SAML
	// assertions. Leave empty if response will not be encrypted. This key
	// MUST be a PKCS8 key.
	// create = optional
	// update = optional
	DecryptingKey string `json:"decryptingKey,omitempty"`
	// Unique name of service provider.
	// required = true
	EntityId string `json:"entityId,omitempty"`
	// The public URL the Delphix Engine will be accessed at.
	// required = true
	BaseUrl string `json:"baseUrl,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// OracleLiveSource - An Oracle LiveSource.
// extends OracleVirtualSource
type OracleLiveSource struct {
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// update = optional
	// create = required
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Amount of tolerable delay for this Oracle LiveSource in seconds.
	// units = sec
	// create = optional
	// update = optional
	// default = 900
	DataAgeWarningThreshold int `json:"dataAgeWarningThreshold,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	// create = optional
	// update = optional
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// A list of object references to Oracle Node Listeners selected for this
	// Virtual Database. Delphix picks one default listener from the target
	// environment if this list is empty at virtual database provision time.
	// create = optional
	// update = optional
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// Online Redo Log size in MB.
	// create = optional
	// update = readonly
	// minimum = 4
	RedoLogSizeInMB int `json:"redoLogSizeInMB,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Resync status for this Oracle LiveSource.
	// enum = [RESYNC_NOT_REQUIRED RESYNC_NEEDED RESYNC_IN_PROGRESS APPLY_READY APPLY_IN_PROGRESS APPLY_FAILED]
	ResyncStatus string `json:"resyncStatus,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Archive Log Mode of the Oracle virtual database.
	// create = optional
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Custom environment variables for Oracle databases.
	// create = optional
	// update = optional
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// Number of Online Redo Log Groups.
	// minimum = 2
	// default = 3
	// create = optional
	// update = readonly
	RedoLogGroups int `json:"redoLogGroups,omitempty"`
}

// StorageTestParameters - Parameters used to execute a storage test.
// extends TypedObject
type StorageTestParameters struct {
	// The list of devices to be used for the test.
	// create = optional
	Devices []string `json:"devices,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Run time of each test, in seconds.
	// minimum = 1
	// maximum = 3600
	// default = 120
	// create = optional
	Duration int `json:"duration,omitempty"`
	// The tests that are to be run.
	// default = ALL
	// create = required
	// enum = [ALL MINIMAL READ WRITE RANDREAD]
	Tests string `json:"tests,omitempty"`
	// Total disk space, spread over all devices, used by the test (in
	// bytes).
	// create = optional
	// units = B
	// base = 1024
	// minimum = 1.048576e+06
	// default = 5.49755813888e+11
	TestRegion float64 `json:"testRegion,omitempty"`
	// True if the disks should be initialized prior to running the
	// benchmark.
	// default = true
	// create = optional
	InitializeDevices *bool `json:"initializeDevices,omitempty"`
	// True if the entire disk should be initialized prior to running the
	// benchmark.
	// default = false
	// create = optional
	InitializeEntireDevice *bool `json:"initializeEntireDevice,omitempty"`
}

// ASELinkData - SAP ASE specific parameters for a link request.
// extends LinkData
type ASELinkData struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The user name for the source DB user.
	// create = optional
	DbUser string `json:"dbUser,omitempty"`
	// Specifies the validated sync mode to synchronize the dSource with the
	// source database.
	// create = optional
	// enum = [ENABLED DISABLED]
	// default = ENABLED
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// The credential for the source DB user.
	// create = optional
	DumpCredentials string `json:"dumpCredentials,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-ase-db-config.json
	// required = true
	// format = objectReference
	Config string `json:"config,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// External file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 80
	// create = optional
	MountBase string `json:"mountBase,omitempty"`
	// A user-provided shell script or executable to run after restoring from
	// a backup during validated sync.
	// create = optional
	// maxLength = 1024
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// The credentials of the database user.
	// required = true
	// properties = map[type:map[description:Object type. required:true format:type default:PasswordCredential type:string]]
	DbCredentials Credential `json:"dbCredentials,omitempty"`
	// Sync parameters for the container.
	// required = true
	// properties = map[type:map[type:string description:Object type. required:true format:type default:ASELatestBackupSyncParameters]]
	SyncParameters ASESyncParameters `json:"syncParameters,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// A user-provided shell script or executable to run prior to restoring
	// from a backup during validated sync.
	// maxLength = 1024
	// create = optional
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// Source database backup location.
	// maxLength = 1024
	// required = true
	LoadBackupPath string `json:"loadBackupPath,omitempty"`
	// Backup location to use for loading backups from the source.
	// create = optional
	LoadLocation string `json:"loadLocation,omitempty"`
	// Information about the host OS user on the staging environment to use
	// for linking.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// Information about the host OS user on the source to use for linking.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	SourceHostUser string `json:"sourceHostUser,omitempty"`
	// The SAP ASE instance on the staging environment that we want to use
	// for validated sync.
	// format = objectReference
	// referenceTo = /delphix-ase-instance.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
}

// OracleCompatibilityCriteria - The compatibility criteria to use for selecting compatible Oracle
// repositories.
// extends CompatibilityCriteria
type OracleCompatibilityCriteria struct {
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are this database version. In case of upgrade,
	// selected repositories are strictly greater than this database version.
	Version string `json:"version,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
}

// MSSqlExportParameters - The parameters to use as input to export MSSQL databases.
// extends DbExportParameters
type MSSqlExportParameters struct {
	// Recovery model of the database.
	// create = optional
	// enum = [SIMPLE BULK_LOGGED FULL]
	// default = FULL
	RecoveryModel string `json:"recoveryModel,omitempty"`
	// The filesystem configuration of the exported database.
	// required = true
	FilesystemLayout *TimeflowFilesystemLayout `json:"filesystemLayout,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig MSSqlDBConfig `json:"sourceConfig,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// create = optional
	// default = true
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
}

// FractionPlugParameters - The parameters to use as input when transporting a transportable
// tablespace.
// extends TypedObject
type FractionPlugParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Optional prefix to add to schemas being moved into warehouse.
	// create = optional
	// update = optional
	// pattern = ^([A-Za-z][A-Za-z0-9_]+)|("[A-Za-z][A-Za-z0-9_]+")$
	// maxLength = 28
	SchemasPrefix string `json:"schemasPrefix,omitempty"`
	// Optional prefix to add to tablespaces being moved into warehouse.
	// create = optional
	// update = optional
	// pattern = ^([A-Za-z][A-Za-z0-9_]+)|("[A-Za-z][A-Za-z0-9_]+")$
	// maxLength = 28
	TablespacesPrefix string `json:"tablespacesPrefix,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base
	// provisioning on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// OracleVirtualCdbSource - A virtual Oracle multitenant container database source. Certain fields
// of the Oracle virtual source are not applicable to the virtual
// container database.
// extends OracleVirtualSource
type OracleVirtualCdbSource struct {
	// Optional database template to use for provisioning and refresh. If
	// set, configParams will be ignored on provision or refresh.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-database-template.json
	ConfigTemplate string `json:"configTemplate,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Runtime properties for this Oracle source.
	Runtime OracleBaseSourceRuntime `json:"runtime,omitempty"`
	// Database file mapping rules. These can only be set at the PDB level.
	// Delphix applies the PDB setting to the virtual CDB.
	// create = readonly
	// update = readonly
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// A list of object references to Oracle Node Listeners selected for this
	// Virtual Database. Delphix picks one default listener from the target
	// environment if this list is empty at virtual database provision time.
	// This is currently not supported for virtual container databases.
	// create = readonly
	// update = readonly
	NodeListeners []string `json:"nodeListeners,omitempty"`
	// Archive Log Mode of the Oracle virtual database. NOARCHIVELOG mode is
	// currently not supported for virtual container databases.
	// create = readonly
	// update = readonly
	// default = true
	ArchivelogMode *bool `json:"archivelogMode,omitempty"`
	// Custom environment variables for Oracle databases. These can only be
	// set at the PDB level. Delphix applies the PDB setting to the virtual
	// CDB.
	// create = readonly
	// update = readonly
	CustomEnvVars []OracleCustomEnvVar `json:"customEnvVars,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// User-specified operation hooks for this source. This is currently not
	// supported for virtual container databases.
	// create = readonly
	// update = readonly
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Online Redo Log size in MB. Customizing Online Redo Log size is
	// currently not supported for virtual container databases.
	// create = readonly
	// update = readonly
	RedoLogSizeInMB int `json:"redoLogSizeInMB,omitempty"`
	// Oracle database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Number of Online Redo Log Groups. Customizing number of Online Redo
	// Log Groups is currently not supported for virtual container databases.
	// create = readonly
	// update = readonly
	RedoLogGroups int `json:"redoLogGroups,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The base mount point to use for the NFS mounts.
	// create = required
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
}

// CurrentConsumerCapacityData - Current data about a particular capacity consumer.
// extends BaseConsumerCapacityData
// cliVisibility = [DOMAIN]
type CurrentConsumerCapacityData struct {
	// Statistics for this consumer.
	Breakdown *CapacityBreakdown `json:"breakdown,omitempty"`
	// Internal unique identifier for this consumer.
	StorageContainer string `json:"storageContainer,omitempty"`
	// Name of this container's group.
	GroupName string `json:"groupName,omitempty"`
	// Name of the container.
	Name string `json:"name,omitempty"`
	// Container from which this TimeFlow was provisioned.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Parent string `json:"parent,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to this container's group.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// AppDataDirectLinkData - Represents the AppData specific parameters of a link request for a
// source directly replicated into the Delphix Engine.
// extends AppDataLinkData
type AppDataDirectLinkData struct {
	// List of symlinks in the source to follow when syncing data. These
	// paths are relative to the root of the source directory. All other
	// symlinks are preserved.
	// required = true
	FollowSymlinks []string `json:"followSymlinks,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-direct-source-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// required = true
	Parameters *Json `json:"parameters,omitempty"`
	// The OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// List of subdirectories in the source to exclude when syncing data.
	// These paths are relative to the root of the source directory.
	// required = true
	Excludes []string `json:"excludes,omitempty"`
}

// MSSqlPlatformParameters - MSSql platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type MSSqlPlatformParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SourceEnableParameters - The parameters to use as input to enable a MSSQL, PostgreSQL, AppData,
// ASE or MySQL source.
// extends TypedObject
type SourceEnableParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether to attempt a startup of the source after the enable.
	// default = true
	AttemptStart *bool `json:"attemptStart,omitempty"`
}

// PgSQLLinkData - PostgreSQL specific parameters for a link request.
// extends LinkData
type PgSQLLinkData struct {
	// Information about the host OS user on the PPT host to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	PptHostUser string `json:"pptHostUser,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// maxLength = 256
	// default = postgres
	// create = optional
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// The username of the database cluster user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The credential of the database cluster user.
	// create = optional
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-pgsql-db-cluster-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// The Postgres installation on the PPT environment that will be used for
	// pre-provisioning.
	// referenceTo = /delphix-pgsql-install.json
	// required = true
	// format = objectReference
	PptRepository string `json:"pptRepository,omitempty"`
	// The external file path.
	// maxLength = 1024
	// create = optional
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// Schema - Schema object.
// extends TypedObject
type Schema struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// JSON representation of the schema based on the locale and API Session
	// version.
	Schema string `json:"schema,omitempty"`
}

// OracleSnapshot - Provisionable snapshot of an Oracle TimeFlow.
// extends TimeflowSnapshot
type OracleSnapshot struct {
	// Runtime properties of the snapshot.
	Runtime *OracleSnapshotRuntime `json:"runtime,omitempty"`
	// Online redo log size in bytes when this snapshot was taken.
	// create = optional
	RedoLogSizeInBytes float64 `json:"redoLogSizeInBytes,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *OracleTimeflowPoint `json:"firstChangePoint,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// True if this snapshot was taken of a standby database.
	// default = false
	// create = optional
	FromPhysicalStandbyVdb *bool `json:"fromPhysicalStandbyVdb,omitempty"`
	// Auxiliary TimeFlows with snapshots controlled by this master snapshot.
	FractionTimeflows []string `json:"fractionTimeflows,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *OracleTimeflowPoint `json:"latestChangePoint,omitempty"`
}

// SyslogConfig - Syslog configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SyslogConfig struct {
	// True if the syslog service is enabled.
	// default = true
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Syslog logging severity. Only events at or above this severity will be
	// logged.
	// update = optional
	// enum = [EMERGENCY ALERT CRITICAL ERROR WARNING NOTICE INFORMATIONAL DEBUG]
	// default = WARNING
	Severity string `json:"severity,omitempty"`
	// Syslog logging pattern. Events will be logged in the pattern as
	// specified.
	// update = optional
	// default = %-5p delphix : %m%n
	Pattern string `json:"pattern,omitempty"`
	// Syslog message format.
	// enum = [TEXT JSON]
	// default = TEXT
	// update = optional
	Format string `json:"format,omitempty"`
	// List of syslog servers. At least one syslog server must be specified.
	// update = required
	Servers []*SyslogServer `json:"servers,omitempty"`
}

// HostMachine - The representation of the host machine.
// extends TypedObject
type HostMachine struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The platform for the host machine.
	Platform string `json:"platform,omitempty"`
	// The amount of RAM on the host machine.
	// units = B
	// base = 1024
	MemorySize float64 `json:"memorySize,omitempty"`
}

// MSSqlDatabaseContainer - A MSSQL Database Container.
// extends DatabaseContainer
type MSSqlDatabaseContainer struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Runtime properties of this container.
	Runtime *MSSqlDBContainerRuntime `json:"runtime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// update = optional
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// Specifies whether Delphix should manage the backups for this container
	// and if so, specify whether they are compressed or uncompressed.
	// enum = [DELPHIX_MANAGED_UNCOMPRESSED DELPHIX_MANAGED_COMPRESSED NOT_DELPHIX_MANAGED]
	// default = NOT_DELPHIX_MANAGED
	// create = optional
	// update = optional
	DelphixManagedStatus string `json:"delphixManagedStatus,omitempty"`
}

// HttpConnectorConfig - Configuration for the HTTP and HTTPS connector of this application.
// extends TypedObject
// cliVisibility = [SYSTEM]
type HttpConnectorConfig struct {
	// Controls the HTTP(s) protocol configuration of this appliance.
	// enum = [HTTP_ONLY HTTPS_ONLY HTTP_REDIRECT BOTH]
	// default = BOTH
	// update = optional
	HttpMode string `json:"httpMode,omitempty"`
	// Version of TLS (transport layer security) enabled on this appliance.
	// minItems = 1
	// update = optional
	TlsVersions []string `json:"tlsVersions,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// ProxyConfiguration - Proxy configuration for a specific protocol.
// extends TypedObject
type ProxyConfiguration struct {
	// Port to use when connecting to the proxy host.
	// minimum = 1
	// maximum = 65535
	// update = optional
	Port int `json:"port,omitempty"`
	// If authentication is required, the username to use when connecting to
	// the proxy.
	// update = optional
	Username string `json:"username,omitempty"`
	// If authentication is required, the password to use when connecting to
	// the proxy.
	// format = password
	// update = optional
	Password string `json:"password,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the proxy is enabled.
	// update = optional
	Enabled *bool `json:"enabled,omitempty"`
	// Host or IP address to use as proxy.
	// format = host
	// update = optional
	Host string `json:"host,omitempty"`
}

// AppDataContainer - Data container for AppData.
// extends DatabaseContainer
type AppDataContainer struct {
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the currently active TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	CurrentTimeflow string `json:"currentTimeflow,omitempty"`
	// Whether to enable high performance mode.
	// enum = [TEMPORARILY_ENABLED ENABLED DISABLED]
	// default = DISABLED
	// create = readonly
	// update = readonly
	PerformanceMode string `json:"performanceMode,omitempty"`
	// A global identifier for this container, including across Delphix
	// Engines.
	Guid string `json:"guid,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Optional user-provided description for the container.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// True if this container is a masked container.
	Masked *bool `json:"masked,omitempty"`
	// The date this container was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	// update = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// The toolkit managing the data in the container.
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// A reference to the container this container was provisioned from.
	// format = objectReference
	// referenceTo = /delphix-container.json
	ProvisionContainer string `json:"provisionContainer,omitempty"`
	// Runtime properties of this container.
	Runtime *AppDataContainerRuntime `json:"runtime,omitempty"`
	// Native processor type of the original database source system.
	Processor string `json:"processor,omitempty"`
	// Native operating system of the original database source system.
	Os string `json:"os,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A reference to the previous TimeFlow for this container.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	PreviousTimeflow string `json:"previousTimeflow,omitempty"`
	// True if this container is a transformation container.
	Transformation *bool `json:"transformation,omitempty"`
	// A reference to the group containing this container.
	// format = objectReference
	// referenceTo = /delphix-group.json
	// create = required
	Group string `json:"group,omitempty"`
}

// JSBranchCreateParameters - The parameters used to create a Jet Stream branch.
// extends TypedObject
type JSBranchCreateParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the branch.
	// required = true
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A reference to the data container to create this branch on.
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	// required = true
	DataContainer string `json:"dataContainer,omitempty"`
	// The Jet Stream data timeline point from which the branch will be
	// created.
	// required = true
	TimelinePointParameters JSTimelinePointParameters `json:"timelinePointParameters,omitempty"`
}

// SyslogServer - Syslog server configuration.
// extends TypedObject
type SyslogServer struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Syslog transport protocol.
	// required = true
	// enum = [udp tcp]
	// default = udp
	Protocol string `json:"protocol,omitempty"`
	// Syslog host name or IP address.
	// required = true
	// format = host
	Address string `json:"address,omitempty"`
	// Syslog port number.
	// required = true
	// minimum = 0
	// maximum = 65535
	// default = 514
	Port int `json:"port,omitempty"`
}

// MSSqlLinkedSource - A linked MSSQL source.
// extends MSSqlSource
type MSSqlLinkedSource struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mssql-db-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntime `json:"runtime,omitempty"`
	// External file path.
	// create = optional
	// update = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// The user for accessing the shared backup location.
	// create = optional
	// update = optional
	// maxLength = 256
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// The staging source for pre-provisioning of the database.
	// format = objectReference
	// referenceTo = /delphix-mssql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	// update = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Specifies the backup types validated sync will use to synchronize the
	// dSource with the source database.
	// update = optional
	// enum = [TRANSACTION_LOG FULL_OR_DIFFERENTIAL FULL NONE]
	// default = TRANSACTION_LOG
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Shared source database backup location.
	// update = optional
	// maxLength = 260
	// create = optional
	SharedBackupLocation string `json:"sharedBackupLocation,omitempty"`
	// The password for accessing the shared backup location.
	// create = optional
	// update = optional
	BackupLocationCredentials *PasswordCredential `json:"backupLocationCredentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
}

// Checklist - Generic checklist object.
// extends TypedObject
type Checklist struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// WindowsCluster - A Windows cluster environment.
// extends SourceEnvironment
type WindowsCluster struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Indicates whether the source environment is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The environment description.
	// create = optional
	// update = optional
	// maxLength = 1024
	Description string `json:"description,omitempty"`
	// The address that will be used to perform discovery on the cluster.
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// A reference to the primary user for this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// update = optional
	PrimaryUser string `json:"primaryUser,omitempty"`
	// A reference to the proxy that will be used to discover the cluster.
	// create = required
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-host.json
	Proxy string `json:"proxy,omitempty"`
	// A reference to the proxy host that will be used for cluster support
	// operations.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// update = optional
	TargetProxy string `json:"targetProxy,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// OracleCustomEnvVarSIPair - Dictates a single environment variable name and value to be set when
// the Delphix Engine administers an Oracle virtual database.
// extends OracleCustomEnvVarPair
type OracleCustomEnvVarSIPair struct {
	// The value of the environment variable.
	// required = true
	VarValue string `json:"varValue,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the environment variable.
	// required = true
	// format = envvarIdentifier
	VarName string `json:"varName,omitempty"`
}

// AppDataStagedLinkData - Represents the AppData specific parameters of a link request for a
// source with a staging source.
// extends AppDataLinkData
type AppDataStagedLinkData struct {
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// The OS user to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// The environment user used to access the staging environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// required = true
	StagingEnvironmentUser string `json:"stagingEnvironmentUser,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-appdata-staged-source-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The JSON payload conforming to the DraftV4 schema based on the type of
	// application data being manipulated.
	// required = true
	Parameters *Json `json:"parameters,omitempty"`
	// The base mount point for the NFS mount on the staging environment.
	// maxLength = 256
	// required = false
	StagingMountBase string `json:"stagingMountBase,omitempty"`
	// The environment used as an intermediate stage to pull data into
	// Delphix.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// required = true
	StagingEnvironment string `json:"stagingEnvironment,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// RollbackParameters - The parameters to use as input to rollback requests for MSSQL,
// PostgreSQL, AppData, ASE or MySQL.
// extends TypedObject
type RollbackParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to roll the
	// database back to.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// MSSqlInstance - A SQL Server Instance.
// extends MSSqlRepository
type MSSqlInstance struct {
	// Account the SQL Server instance is running as.
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The version of the SQL Server instance.
	Version string `json:"version,omitempty"`
	// The name of the SQL Server instance.
	InstanceName string `json:"instanceName,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The Servername of the SQL Server instance.
	ServerName string `json:"serverName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the environment containing this repository.
	// create = required
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
	// The SQL Server instance home.
	InstallationPath string `json:"installationPath,omitempty"`
	// The network port for connecting to the SQL Server instance.
	Port int `json:"port,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Internal version of the SQL Server instance.
	InternalVersion int `json:"internalVersion,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
}

// NetworkRoute - IP routing table.
// extends TypedObject
// cliVisibility = [DOMAIN SYSTEM]
type NetworkRoute struct {
	// Output interface to use for the route.
	// format = objectReference
	// referenceTo = /delphix-network-interface.json
	// create = optional
	OutInterface string `json:"outInterface,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Destination for the route in Classless Inter-Domain Routing (CIDR)
	// notation or the keyword 'default'.
	// format = routeDestination
	// create = required
	Destination string `json:"destination,omitempty"`
	// Next hop for the route.
	// format = ipAddress
	// create = required
	Gateway string `json:"gateway,omitempty"`
}

// XppStatus - The current cross-platform provisioning status of a container.
// extends Checklist
type XppStatus struct {
	// Status of the cross-platform provisioning staging environment.
	StagingStatus *XppStagingStatus `json:"stagingStatus,omitempty"`
	// Status of the cross-platform provisioning target environment.
	TargetStatus *XppTargetStatus `json:"targetStatus,omitempty"`
	// Cross-platform provisioning validation status of the container.
	ValidateStatus *XppValidateStatus `json:"validateStatus,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Status of the last cross-platform provision of the container.
	LastRunStatus *XppLastRunStatus `json:"lastRunStatus,omitempty"`
}

// OracleRefreshParameters - The parameters to use as input to refresh Oracle databases.
// extends RefreshParameters
type OracleRefreshParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to refresh the
	// database to.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The security credential of the privileged user to run the refresh
	// operation as.
	Credential Credential `json:"credential,omitempty"`
	// The name of the privileged user to run the refresh operation as.
	Username string `json:"username,omitempty"`
}

// ReplicationSpec - Replication setup.
// extends UserObject
type ReplicationSpec struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Connect to the replication target host via the system-wide SOCKS
	// proxy.
	// create = optional
	// update = optional
	// default = false
	UseSystemSocksSetting *bool `json:"useSystemSocksSetting,omitempty"`
	// Target TCP port number for the Delphix Session Protocol.
	// maximum = 65535
	// default = 8415
	// create = optional
	// update = optional
	// minimum = 0
	TargetPort int `json:"targetPort,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// format = objectName
	// create = optional
	Name string `json:"name,omitempty"`
	// Replication schedule in the form of a quartz-formatted string.
	// minLength = 1
	// maxLength = 256
	// create = optional
	// update = optional
	Schedule string `json:"schedule,omitempty"`
	// Globally unique identifier for this replication spec.
	// minLength = 1
	// maxLength = 256
	Tag string `json:"tag,omitempty"`
	// Runtime properties of this replication spec.
	Runtime *ReplicationSpecRuntime `json:"runtime,omitempty"`
	// Specification of the objects to replicate.
	// create = required
	// update = optional
	ObjectSpecification ReplicationObjectSpecification `json:"objectSpecification,omitempty"`
	// Replication target host address.
	// format = host
	// create = required
	// update = optional
	TargetHost string `json:"targetHost,omitempty"`
	// Total number of transport connections to use.
	// create = optional
	// update = optional
	// default = 1
	// minimum = 1
	// maximum = 16
	NumberOfConnections int `json:"numberOfConnections,omitempty"`
	// Principal name used to authenticate to the replication target host.
	// create = required
	// update = optional
	TargetPrincipal string `json:"targetPrincipal,omitempty"`
	// Description of this replication spec.
	// create = optional
	// update = optional
	// maxLength = 4096
	Description string `json:"description,omitempty"`
	// Encrypt replication network traffic.
	// update = optional
	// default = false
	// create = optional
	Encrypted *bool `json:"encrypted,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Indication whether the replication spec schedule is enabled or not.
	// create = optional
	// update = optional
	// default = false
	AutomaticReplication *bool `json:"automaticReplication,omitempty"`
	// Credential used to authenticate to the replication target host.
	// create = required
	// update = optional
	TargetCredential *PasswordCredential `json:"targetCredential,omitempty"`
	// Bandwidth limit (MB/s) for replication network traffic. A value of 0
	// means no limit.
	// create = optional
	// update = optional
	// default = 0
	// minimum = 0
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
}

// MSSqlCompatibilityCriteria - The compatibility criteria to use for selecting compatible MSSql
// repositories.
// extends CompatibilityCriteria
type MSSqlCompatibilityCriteria struct {
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories are this database version. In case of upgrade,
	// selected repositories are strictly greater than this database version.
	Version string `json:"version,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
}

// JSDataContainerResetParameters - The parameters used to reset a data container..
// extends TypedObject
type JSDataContainerResetParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If this value is true, then do the operation without taking a
	// snapshot.
	// required = true
	// default = false
	ForceOption *bool `json:"forceOption,omitempty"`
}

// ReplicationTargetState - State of a replication at the target.
// extends UserObject
type ReplicationTargetState struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
}

// UpgradeCheckResult - Describes unsatisfied upgrade requirements.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type UpgradeCheckResult struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// A localized, textual description of the action the user should take to
	// overcome the error.
	// create = readonly
	// update = readonly
	Action string `json:"action,omitempty"`
	// A localized string that is the broad category of this check.
	// create = readonly
	// update = readonly
	Title string `json:"title,omitempty"`
	// The status of the upgrade check result.
	// enum = [ACTIVE IGNORED RESOLVED]
	// create = readonly
	// update = readonly
	Status string `json:"status,omitempty"`
	// A unique identifier for the type of the upgrade check result.
	// create = readonly
	// update = readonly
	BundleId string `json:"bundleId,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// A localized, textual description of the error.
	// create = readonly
	// update = readonly
	Description string `json:"description,omitempty"`
	// Script output related to the check result to assist in resolving the
	// issue.
	// create = readonly
	// update = readonly
	Output string `json:"output,omitempty"`
	// A localized, textual description of the impact the error might have on
	// the system.
	// create = readonly
	// update = readonly
	Impact string `json:"impact,omitempty"`
	// The severity of the missing upgrade requirement. CRITICAL check
	// results block the upgrade.
	// enum = [WARNING CRITICAL INFORMATIONAL]
	// create = readonly
	// update = readonly
	Severity string `json:"severity,omitempty"`
	// A reference to the upgrade version that generated this check result.
	// referenceTo = /delphix-upgrade-version.json
	// create = readonly
	// update = readonly
	// format = objectReference
	Version string `json:"version,omitempty"`
}

// JSBookmark - A named entity that represents a point in time for all of the data
// sources in a data layout.
// extends NamedUserObject
type JSBookmark struct {
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The data template this bookmark was created on or the template of the
	// data container this bookmark was created on.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	Template string `json:"template,omitempty"`
	// A set of user-defined labels for this bookmark.
	// update = optional
	// create = optional
	Tags []string `json:"tags,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// True if this bookmark is shared.
	// create = optional
	Shared *bool `json:"shared,omitempty"`
	// True if this bookmark is usable as input to a data operation (e.g.,
	// CREATE_BRANCH or RESTORE).
	Usable *bool `json:"usable,omitempty"`
	// The time at which the bookmark was created.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// A reference to the branch this bookmark applies to.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// create = required
	Branch string `json:"branch,omitempty"`
	// Denotes whether or not this bookmark was created on a data container
	// or a data template.
	// enum = [DATA_CONTAINER DATA_TEMPLATE]
	BookmarkType string `json:"bookmarkType,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The timestamp for the data that the bookmark refers to.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// The name of the data template this bookmark was created on or the
	// template of the data container this bookmark was created on.
	TemplateName string `json:"templateName,omitempty"`
	// The number of times this bookmark has been checked out. This means it
	// was used as input to a RESTORE, CREATE_BRANCH, or RESET operation.
	CheckoutCount int `json:"checkoutCount,omitempty"`
	// Description of this bookmark.
	// maxLength = 4096
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// A policy will automatically delete this bookmark at this time. If the
	// value is null, then the bookmark will be kept until manually deleted.
	// format = date
	// create = optional
	// update = optional
	Expiration string `json:"expiration,omitempty"`
	// The data container this bookmark was created on. This will be null if
	// the bookmark was created on a data template.
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	Container string `json:"container,omitempty"`
	// The name of the data container this bookmark was created on. This will
	// be null if the bookmark was created on a data template.
	ContainerName string `json:"containerName,omitempty"`
}

// MaskingServiceConfig - Configuration for the Masking Service this Engine communicates with.
// extends UserObject
type MaskingServiceConfig struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// update = readonly
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Protocol scheme for use when communicating with server.
	// enum = [HTTP HTTPS]
	// update = optional
	Scheme string `json:"scheme,omitempty"`
	// Password to use when authenticating to the server.
	// update = optional
	Credentials Credential `json:"credentials,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// IP address or hostname of server hosting Masking Service.
	// format = host
	// update = optional
	Server string `json:"server,omitempty"`
	// Port number to use.
	// minimum = 1
	// maximum = 65535
	// update = optional
	Port int `json:"port,omitempty"`
	// Username to use when authenticating to the server.
	// update = optional
	Username string `json:"username,omitempty"`
}

// DiskOpsDatapointStream - A stream of datapoints from a DISK_OPS analytics slice.
// extends DatapointStream
type DiskOpsDatapointStream struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Device the I/Os were issued to.
	Device string `json:"device,omitempty"`
	// Whether the I/Os resulted in errors.
	Error *bool `json:"error,omitempty"`
}

// RegistrationParameters - Credentials used to register the Delphix Engine.
// extends TypedObject
type RegistrationParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Username to send to registration portal.
	Username string `json:"username,omitempty"`
	// Password to send to registration portal.
	// format = password
	Password string `json:"password,omitempty"`
}

// LdapServer - LDAP Server Configuration.
// extends ReadonlyNamedUserObject
// cliVisibility = [SYSTEM]
type LdapServer struct {
	// LDAP server port.
	// minimum = 1
	// maximum = 65535
	// required = true
	Port int `json:"port,omitempty"`
	// LDAP authentication method.
	// enum = [SIMPLE DIGEST_MD5]
	// required = true
	AuthMethod string `json:"authMethod,omitempty"`
	// Authenticate using SSL.
	// required = true
	UseSSL *bool `json:"useSSL,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// LDAP server host name.
	// format = host
	// required = true
	Host string `json:"host,omitempty"`
}

// OracleLiveSourceRuntime - Runtime (non-persistent) properties of an Oracle LiveSource.
// extends OracleBaseSourceRuntime
type OracleLiveSourceRuntime struct {
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// List of active database instances for the source.
	ActiveInstances []*OracleActiveInstance `json:"activeInstances,omitempty"`
	// Table of key database performance statistics.
	DatabaseStats []*OracleDatabaseStatsSection `json:"databaseStats,omitempty"`
	// Operating mode of the database.
	// default = UNKNOWN
	// enum = [READ_WRITE READ_ONLY STANDBY_READ_ONLY MOUNTED_ONLY UNKNOWN]
	DatabaseMode string `json:"databaseMode,omitempty"`
	// Time zone of the source database at the time the runtime data was
	// updated.
	SourceDatabaseTimezone string `json:"sourceDatabaseTimezone,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether the incarnation ID changed on the primary database.
	SourceResetlogsIDChangeDetected *bool `json:"sourceResetlogsIDChangeDetected,omitempty"`
	// The current role of the database.
	// default = UNKNOWN
	// enum = [PHYSICAL_STANDBY LOGICAL_STANDBY SNAPSHOT_STANDBY FAR_SYNC PRIMARY UNKNOWN]
	DatabaseRole string `json:"databaseRole,omitempty"`
	// Highest SCN at which non-logged changes were generated.
	LastNonLoggedLocation string `json:"lastNonLoggedLocation,omitempty"`
	// Current data lag between LiveSource and source database in seconds.
	// units = sec
	CurrentDataAge int `json:"currentDataAge,omitempty"`
	// Indicates whether there is non-logged data on the standby.
	NonLoggedDataDetected *bool `json:"nonLoggedDataDetected,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Redo log transport status from the source database to the LiveSource.
	// enum = [UNKNOWN WORKING NO_INITIAL_DATA NO_NEW_DATA]
	TransportStatus string `json:"transportStatus,omitempty"`
	// MRP apply status for the standby database associated with the
	// LiveSource.
	// enum = [UNKNOWN WORKING APPLY_FAILED APPLY_ON_WRONG_INCARNATION UNRESOLVABLE_GAP_DETECTED]
	ApplyStatus string `json:"applyStatus,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
	// The time at which this runtime data was updated.
	// format = date
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Has data age exceeded the user specified threshold.
	IsDataAgeWarningExceeded *bool `json:"isDataAgeWarningExceeded,omitempty"`
	// Indicates whether the LiveSource is not in standby mode.
	UnexpectedRoleChangeDetected *bool `json:"unexpectedRoleChangeDetected,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
}

// ASELatestBackupSyncParameters - The parameters to use as input to sync a SAP ASE database using the
// latest backup.
// extends ASESyncParameters
type ASELatestBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OraclePDBConfig - Representation of properties for an Oracle pluggable database
// configuration.
// extends OracleBaseDBConfig
type OraclePDBConfig struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The username of the database user.
	// update = optional
	// maxLength = 30
	User string `json:"user,omitempty"`
	// The password of the database user. This must be a PasswordCredential
	// instance.
	// update = optional
	Credentials *PasswordCredential `json:"credentials,omitempty"`
	// The name of the database.
	// create = required
	// pattern = ^[a-zA-Z0-9][a-zA-Z0-9_]*$
	// maxLength = 30
	DatabaseName string `json:"databaseName,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	// create = optional
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The list of database services.
	// create = optional
	// update = optional
	Services []*OracleService `json:"services,omitempty"`
	// The DB config of an Oracle multitenant database this pluggable
	// database belongs to.
	// format = objectReference
	// referenceTo = /delphix-oracle-db-config.json
	// create = optional
	// update = optional
	CdbConfig string `json:"cdbConfig,omitempty"`
}

// ReplicationSecureList - List of containers that are to be securely replicated.
// extends ReplicationObjectSpecification
type ReplicationSecureList struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Containers to replicate, in canonical object reference form.
	// update = optional
	// minItems = 1
	// create = required
	Containers []string `json:"containers,omitempty"`
}

// ApplyVersionParameters - The parameters to use as input to upgrade.
// extends TypedObject
type ApplyVersionParameters struct {
	// If true, the system reboots immediately after upgrade. If false, the
	// system is shutdown.
	// create = optional
	// default = true
	Reboot *bool `json:"reboot,omitempty"`
	// If true, the Delphix Engine is upgraded without updating the OS
	// software. The operation will fail gracefully if the upgrade version
	// requires a version of the OS that is newer than what is currently
	// running. The OS software can subsequently be upgraded by applying any
	// version and setting defer to false. It is possible to catch up to the
	// current OS version on a previously deferred upgrade by re-applying the
	// running version with a defer setting of false.
	// default = false
	// create = optional
	Defer *bool `json:"defer,omitempty"`
	// If true, all data sources (VDBs and dSources) are automatically
	// disabled prior to upgrade, and re-enabled after upgrade. If any source
	// cannot be disabled, the recovery semantics are governed by the
	// "ignoreQuiesceSourcesFailures" and "enableSourcesOnFailure"
	// properties.
	// create = optional
	// default = true
	QuiesceSources *bool `json:"quiesceSources,omitempty"`
	// This property governs whether or not data sources are re-enabled or
	// left disabled in the event that upgrade fails before the Delphix
	// Engine is restarted.
	// create = optional
	// default = true
	EnableSourcesOnFailure *bool `json:"enableSourcesOnFailure,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If set to false, disables verification before applying the upgrade.
	// This will only disable verification if a successful verification has
	// been run in the past hour.
	// create = optional
	// default = true
	Verify *bool `json:"verify,omitempty"`
	// If true, a failure to quiesce sources will not block the upgrade.
	// create = optional
	// default = false
	IgnoreQuiesceSourcesFailures *bool `json:"ignoreQuiesceSourcesFailures,omitempty"`
}

// AppDataDirectSourceConfig - Source config for directly linked AppData sources.
// extends AppDataSourceConfig
type AppDataDirectSourceConfig struct {
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The toolkit associated with this source config.
	// referenceTo = /delphix-toolkit.json
	// format = objectReference
	Toolkit string `json:"toolkit,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// default = true
	// create = optional
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The name of the config.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The list of parameters specified by the source config schema in the
	// toolkit. If no schema is specified, this list is empty.
	// create = optional
	// update = optional
	Parameters *Json `json:"parameters,omitempty"`
	// The path to the data to be synced.
	// create = optional
	// update = optional
	// maxLength = 1024
	Path string `json:"path,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
}

// OracleSysauxDatafileSpecification - Describes an Oracle datafile in the SYSAUX tablespace.
// extends OracleDatafileTempfileSpecification
type OracleSysauxDatafileSpecification struct {
	// Enable or disable the automatic extension of a new or existing
	// datafile or tempfile.
	// create = optional
	// update = optional
	// default = true
	AutoExtend *bool `json:"autoExtend,omitempty"`
	// The size in MB of the next increment of disk space to be allocated
	// automatically when more extents are required. The default is the size
	// of one data block.
	// create = optional
	// update = optional
	AutoExtendIncrement int `json:"autoExtendIncrement,omitempty"`
	// The maximum disk space allowed for automatic extension of the
	// datafile. Omit this if you do not want to limit the disk space that
	// Oracle can allocate to the datafile or tempfile.
	// create = optional
	// update = optional
	MaxSize int `json:"maxSize,omitempty"`
	// The size of the file in MB.
	// create = optional
	// update = optional
	// minimum = 900
	// default = 900
	Size int `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data file, temporary file, or redo log.
	// pattern = ^[a-zA-Z0-9_\-\.]+$
	// create = optional
	// update = optional
	Filename string `json:"filename,omitempty"`
}

// StorageDevice - A storage device on the system.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type StorageDevice struct {
	// True if the device is currently configured in the system.
	Configured *bool `json:"configured,omitempty"`
	// Physical size of the device, in bytes.
	// base = 1024
	// units = B
	Size float64 `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Vendor ID of the device.
	Vendor string `json:"vendor,omitempty"`
	// Model ID of the device.
	Model string `json:"model,omitempty"`
	// Object name.
	// maxLength = 256
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// Serial number of the device.
	Serial string `json:"serial,omitempty"`
}

// XppTargetStatus - Status of the cross-platform provisioning target environment.
// extends ChecklistItemDetail
type XppTargetStatus struct {
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Description of this status item.
	Description string `json:"description,omitempty"`
}

// JSSourceDataTimestampParameters - Input parameters for the API that given a point in time, returns the
// timestamps of the latest provisionable points, before the specified
// time and from the given branch, for each data source in the branch's
// data layout.
// extends TypedObject
type JSSourceDataTimestampParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The time that will be used to find provisionable timestamps for the
	// sources in the branch's data layout.
	// format = date
	// required = true
	Time string `json:"time,omitempty"`
	// A reference to the Jet Stream branch.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	// required = true
	Branch string `json:"branch,omitempty"`
}

// LinkedSourceOperations - Describes operations which are performed on linked sources at various
// times.
// extends TypedObject
type LinkedSourceOperations struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Operations to perform after syncing a linked source.
	// create = optional
	// update = optional
	PostSync []SourceOperation `json:"postSync,omitempty"`
	// Operations to perform before syncing from a linked source. These
	// operations can quiesce any data prior to syncing.
	// create = optional
	// update = optional
	PreSync []SourceOperation `json:"preSync,omitempty"`
}

// FileDownloadResult - Result of a file download request.
// extends FileProcessingResult
type FileDownloadResult struct {
	// URL to download from or upload to.
	Url string `json:"url,omitempty"`
	// Token to pass as parameter to identify the file.
	Token string `json:"token,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// JSContainerUsageData - The space usage information for a data container.
// extends TypedObject
type JSContainerUsageData struct {
	// The data container that this usage information is for.
	// format = objectReference
	// referenceTo = /delphix-js-data-container.json
	DataContainer string `json:"dataContainer,omitempty"`
	// The amount of space that will be freed if this data container is
	// deleted or purged. This assumes that the data container is deleted
	// along with underlying data sources.
	// base = 1024
	// units = B
	Unique float64 `json:"unique,omitempty"`
	// The amount of space that cannot be freed on the parent data template
	// (or sibling data containers) because it is also being referenced by
	// this data container due to restore or create branch operations.
	// base = 1024
	// units = B
	SharedOthers float64 `json:"sharedOthers,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The amount of space that cannot be freed on this data container
	// because it is also being referenced by sibling data containers due to
	// restore or create branch operations.
	// units = B
	// base = 1024
	SharedSelf float64 `json:"sharedSelf,omitempty"`
	// The amount of space that would be consumed by the data in this
	// container without Delphix.
	// base = 1024
	// units = B
	Unvirtualized float64 `json:"unvirtualized,omitempty"`
}

// X509Certificate - X509 Certificate.
// extends UserObject
// cliVisibility = [SYSTEM]
type X509Certificate struct {
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Delphix trusts this certificate .
	Accepted *bool `json:"accepted,omitempty"`
	// Certificate serial number.
	SerialNumber string `json:"serialNumber,omitempty"`
	// End of validity.
	ValidTo string `json:"validTo,omitempty"`
	// SHA-1 fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
	// MD5 fingerprint.
	Md5Fingerprint string `json:"md5Fingerprint,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Issuer of this certificate.
	IssuedByDN string `json:"issuedByDN,omitempty"`
	// Distinguished name of subject of this certificate.
	IssuedToDN string `json:"issuedToDN,omitempty"`
	// Start of validity.
	ValidFrom string `json:"validFrom,omitempty"`
}

// MSSqlDBContainerRuntime - Runtime properties of a MSSQL database container.
// extends DBContainerRuntime
type MSSqlDBContainerRuntime struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
	// The UUID of the source database backupset that was last restored in
	// this container.
	LastRestoredBackupSetUUID string `json:"lastRestoredBackupSetUUID,omitempty"`
}

// JSDataContainerCreateWithoutRefreshParameters - The parameters used to create a data container when not refreshing
// data sources.
// extends JSDataContainerCreateParameters
type JSDataContainerCreateWithoutRefreshParameters struct {
	// A description of this data layout to define what it is used for.
	// create = optional
	// maxLength = 4096
	Notes string `json:"notes,omitempty"`
	// The set of data sources that belong to this data layout.
	// required = true
	DataSources []*JSDataSourceCreateParameters `json:"dataSources,omitempty"`
	// Key/value pairs used to specify attributes for this data layout.
	// create = optional
	Properties map[string]string `json:"properties,omitempty"`
	// A reference to the template that this data container is provisioned
	// from.
	// format = objectReference
	// referenceTo = /delphix-js-data-template.json
	// required = true
	Template string `json:"template,omitempty"`
	// A reference to the list of users that own this data container.
	// create = optional
	Owners []string `json:"owners,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the data layout.
	// maxLength = 256
	// required = true
	Name string `json:"name,omitempty"`
}

// MSSqlFailoverClusterRepository - The representation of a SQL Server Failover Cluster repository.
// extends MSSqlBaseClusterRepository
type MSSqlFailoverClusterRepository struct {
	// The list of MSSQL Cluster instances belonging to this repository.
	Instances []MSSqlBaseClusterInstance `json:"instances,omitempty"`
	// The list of listeners belonging to this repository.
	Listeners []MSSqlBaseClusterListener `json:"listeners,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// update = optional
	// default = true
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Internal version of the SQL Server instance.
	InternalVersion int `json:"internalVersion,omitempty"`
	// The list of drive letters belonging to this Failover Cluster
	// repository.
	Drives []*MSSqlFailoverClusterDriveLetter `json:"drives,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
}

// AppDataEmptyVFilesCreationParameters - The parameters to use as input when creating a new empty vFiles
// dataset.
// extends EmptyDatasetCreationParameters
type AppDataEmptyVFilesCreationParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The new container for the created dataset.
	// required = true
	Container *AppDataContainer `json:"container,omitempty"`
	// The source that describes an external dataset instance.
	// required = true
	Source *AppDataVirtualSource `json:"source,omitempty"`
	// The source config including dynamically discovered attributes of the
	// source.
	// required = true
	SourceConfig AppDataSourceConfig `json:"sourceConfig,omitempty"`
}

// PgSQLCompatibilityCriteria - The compatibility criteria to use for selecting compatible PgSQL
// repositories.
// extends CompatibilityCriteria
type PgSQLCompatibilityCriteria struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Selected repositories are installed on a host with this OS.
	// enum = [Linux AIX HPUX SunOS Windows]
	Os string `json:"os,omitempty"`
	// Selected repositories are installed on a host with this type of
	// processor.
	// enum = [x86 ia64 powerpc sparc]
	Processor string `json:"processor,omitempty"`
	// Selected repositories are installed on a host with this architecture
	// (32-bit or 64-bit).
	Architecture int `json:"architecture,omitempty"`
	// If true, selected repositories have staging enabled.
	StagingEnabled *bool `json:"stagingEnabled,omitempty"`
	// Selected repositories are installed on this environment.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// Selected repositories have this size WAL segments.
	// units = B
	// base = 1024
	SegmentSize int `json:"segmentSize,omitempty"`
	// Selected repositories will match this variant of the PostgreSQL
	// distribution.
	// enum = [PostgreSQL EnterpriseDB]
	Variant string `json:"variant,omitempty"`
	// Selected repositories are this database version. In case of upgrade,
	// selected repositories are strictly greater than this database version.
	// format = pgsqlVersion
	Version string `json:"version,omitempty"`
}

// ToolkitLocale - Contains a mapping from message IDs to messages for a locale.
// extends TypedObject
type ToolkitLocale struct {
	// The name of this locale.
	// required = true
	// format = locale
	LocaleName string `json:"localeName,omitempty"`
	// A mapping of message IDs to messages for this locale.
	// required = true
	Messages map[string]string `json:"messages,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLVersion - Version of a MySQL installation.
// extends TypedObject
type MySQLVersion struct {
	// Version of the MySQL installation.
	// format = mysqlVersion
	Version string `json:"version,omitempty"`
	// Variant of the MySQL installation.
	// enum = [CommunityServer MariaDB]
	Variant string `json:"variant,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostOS - The operating system information for the host.
// extends TypedObject
type HostOS struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The OS version.
	Version string `json:"version,omitempty"`
	// The OS timezone.
	Timezone string `json:"timezone,omitempty"`
	// The OS distribution.
	Distribution string `json:"distribution,omitempty"`
	// The OS name.
	Name string `json:"name,omitempty"`
	// The OS kernel.
	Kernel string `json:"kernel,omitempty"`
	// The OS release.
	Release string `json:"release,omitempty"`
}

// HostPrivilegeElevationProfileScript - A script that is part of a profile for elevating user privileges on a
// host.
// extends PersistentObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationProfileScript struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The privilege elevation profile script name.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// The contents of the privilege elevation profile script.
	// create = required
	// update = optional
	Contents string `json:"contents,omitempty"`
	// The privilege elevation profile to which this script belongs.
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-host-privilege-elevation-profile.json
	// create = required
	Profile string `json:"profile,omitempty"`
}

// DNSConfig - DNS Client Configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type DNSConfig struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// One of more DNS domain names.
	// update = optional
	Domain []string `json:"domain,omitempty"`
	// List of DNS servers.
	// update = optional
	Servers []string `json:"servers,omitempty"`
}

// Json - A dummy schema that is used to represent JSON.
type Json struct {
}

// CertificateFetchParameters - Parameters for fetching a certificate.
// extends TypedObject
type CertificateFetchParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Port number.
	// minimum = 1
	// maximum = 65535
	// required = true
	Port int `json:"port,omitempty"`
	// Hostname or IP address.
	// format = host
	// required = true
	Host string `json:"host,omitempty"`
}

// PublicSystemInfo - Retrieve static system-wide properties.
// extends TypedObject
// cliVisibility = [DOMAIN SYSTEM]
type PublicSystemInfo struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Delphix version of the current system software.
	BuildVersion *VersionInfo `json:"buildVersion,omitempty"`
	// Maximum supported API version of the current system software.
	ApiVersion *APIVersion `json:"apiVersion,omitempty"`
	// Product type.
	ProductType string `json:"productType,omitempty"`
	// Description of the current system software.
	BuildTitle string `json:"buildTitle,omitempty"`
	// Time at which the current system software was built.
	// format = date
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	// List of available locales.
	Locales []string `json:"locales,omitempty"`
	// Name of the product that the system is running.
	ProductName string `json:"productName,omitempty"`
	// The list of enabled features on this Delphix Engine.
	EnabledFeatures []string `json:"enabledFeatures,omitempty"`
	// Security banner to display prior to login.
	Banner string `json:"banner,omitempty"`
	// Indicates whether the server has gone through initial setup or not.
	Configured *bool `json:"configured,omitempty"`
	// The current system locale.
	// format = locale
	CurrentLocale string `json:"currentLocale,omitempty"`
}

// PgSQLCreateTransformationParameters - Represents the parameters of a createTransformation request for a
// PgSQL container.
// extends CreateTransformationParameters
type PgSQLCreateTransformationParameters struct {
	// The port number used for provisioning a PgSQL container during
	// transformation application.
	// create = required
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Operations to perform when this transformation is applied.
	// create = required
	Operations []SourceOperation `json:"operations,omitempty"`
	// The container that will contain the transformed data associated with
	// the newly created transformation; the "transformation container".
	// create = required
	Container *PgSQLDatabaseContainer `json:"container,omitempty"`
	// Reference to the user used during application of the transformation.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// Reference to the repository used during application of the
	// transformation.
	// referenceTo = /delphix-source-repository.json
	// create = required
	// format = objectReference
	Repository string `json:"repository,omitempty"`
}

// PgSQLVirtualSource - A virtual PostgreSQL source.
// extends PgSQLSource
type PgSQLVirtualSource struct {
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// PostgreSQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// The base mount point to use for the NFS mounts.
	// maxLength = 256
	// create = required
	MountBase string `json:"mountBase,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-source-config.json
	// create = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// User-specified operation hooks for this source.
	// update = optional
	// create = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Entries in the PostgreSQL username map file (pg_ident.conf).
	// create = optional
	IdentEntries []*PgSQLIdentEntry `json:"identEntries,omitempty"`
	// Entries in the PostgreSQL host-based authentication file
	// (pg_hba.conf).
	// create = optional
	HbaEntries []*PgSQLHBAEntry `json:"hbaEntries,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Object name.
	// maxLength = 256
	// format = objectName
	// create = optional
	// update = optional
	Name string `json:"name,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Runtime properties of this source.
	Runtime *PgSQLSourceRuntime `json:"runtime,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
}

// VersionInfo - Representation of a Delphix software revision.
// extends TypedObject
type VersionInfo struct {
	// Minor version number.
	Minor int `json:"minor,omitempty"`
	// Micro version number.
	Micro int `json:"micro,omitempty"`
	// Patch version number.
	Patch int `json:"patch,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Major version number.
	Major int `json:"major,omitempty"`
}

// EventFilter - An event filter that specifies which event types to match against.
// extends AlertFilter
type EventFilter struct {
	// List of event types. Only alerts of the given event type are included.
	// Each event type is a string representing the event class of the
	// corresponding alerts. Wildcards are supported to include classes of
	// events.
	// create = optional
	// update = optional
	// minItems = 1
	EventTypes []string `json:"eventTypes,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// StorageDeviceInitializeStatus - The status of an initialize operation of a storage device in the
// system.
// extends TypedObject
type StorageDeviceInitializeStatus struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of data initialized, in bytes.
	// base = 1024
	// units = B
	BytesDone float64 `json:"bytesDone,omitempty"`
	// Total amount of data to initialize (including data already
	// initialized), in bytes.
	// base = 1024
	// units = B
	BytesEst float64 `json:"bytesEst,omitempty"`
	// Initialization state.
	// enum = [NONE ACTIVE CANCELED SUSPENDED COMPLETED]
	State string `json:"state,omitempty"`
}

// ASESnapshot - Provisionable snapshot of a SAP ASE TimeFlow.
// extends TimeflowSnapshot
type ASESnapshot struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// update = readonly
	// create = readonly
	Name string `json:"name,omitempty"`
	// Reference to the database of which this TimeFlow is a part.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Runtime properties of the snapshot.
	Runtime *ASESnapshotRuntime `json:"runtime,omitempty"`
	// Version of database source repository at the time the snapshot was
	// taken.
	Version string `json:"version,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The location of the snapshot within the parent TimeFlow represented by
	// this snapshot.
	LatestChangePoint *ASETimeflowPoint `json:"latestChangePoint,omitempty"`
	// Retention policy, in days. A value of -1 indicates the snapshot should
	// be kept forever.
	// update = optional
	Retention int `json:"retention,omitempty"`
	// A value in the set {CONSISTENT, INCONSISTENT, CRASH_CONSISTENT}
	// indicating what type of recovery strategies must be invoked when
	// provisioning from this snapshot.
	Consistency string `json:"consistency,omitempty"`
	// Point in time at which this snapshot was created. This may be
	// different from the time corresponding to the TimeFlow.
	// format = date
	CreationTime string `json:"creationTime,omitempty"`
	// Time zone of the source database at the time the snapshot was taken.
	Timezone string `json:"timezone,omitempty"`
	// Boolean value indicating if a virtual database provisioned from this
	// snapshot will be missing nologging changes.
	MissingNonLoggedData *bool `json:"missingNonLoggedData,omitempty"`
	// Boolean value indicating that this snapshot is in a transient state
	// and should not be user visible.
	Temporary *bool `json:"temporary,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The location within the parent TimeFlow at which this snapshot was
	// initiated.
	FirstChangePoint *ASETimeflowPoint `json:"firstChangePoint,omitempty"`
	// TimeFlow of which this snapshot is a part.
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
}

// SystemPackage - A package whose version can be changed by sysadmins.
// extends NamedUserObject
// cliVisibility = [SYSTEM]
type SystemPackage struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Package name.
	// maxLength = 256
	// format = objectName
	// update = readonly
	Name string `json:"name,omitempty"`
	// Current version of the package.
	// update = required
	Version string `json:"version,omitempty"`
	// Possible versions for this package.
	// update = readonly
	PossibleVersions []string `json:"possibleVersions,omitempty"`
}

// NullConstraint - If an axis has this type of constraint, it means that no constraints
// can be placed on this axis. This constraint type does nothing and has
// no descendent types.
// extends AxisConstraint
type NullConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
}

// NetworkLatencyTestParameters - Parameters used to execute a network latency test.
// extends NetworkTestParameters
type NetworkLatencyTestParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	// create = optional
	RemoteHost string `json:"remoteHost,omitempty"`
	// A hostname or literal IP address to test. Either remoteAddress or
	// remoteHost must be set, but not both.
	// format = host
	// create = optional
	RemoteAddress string `json:"remoteAddress,omitempty"`
	// Number of requests to send.
	// maximum = 3600
	// default = 20
	// create = optional
	// minimum = 1
	RequestCount int `json:"requestCount,omitempty"`
	// The size of requests to send (bytes).
	// units = B
	// base = 1024
	// minimum = 8
	// maximum = 65507
	// default = 8
	// create = optional
	RequestSize int `json:"requestSize,omitempty"`
}

// NfsOpsDatapointStream - A stream of datapoints from an NFS_OPS analytics slice.
// extends DatapointStream
type NfsOpsDatapointStream struct {
	// Address of the client.
	// format = host
	Client string `json:"client,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
}

// SnapshotLogFetchParameters - Parameters to fetch log files that cover a snapshot and the TimeFlow
// range up to the next snapshot.
// extends LogFetchSSH
type SnapshotLogFetchParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Remote host to connect to.
	// required = true
	Host string `json:"host,omitempty"`
	// User name to authenticate as.
	// required = true
	Username string `json:"username,omitempty"`
	// User credentials. If not provided will use environment credentials for
	// 'username' on 'host'.
	Credentials Credential `json:"credentials,omitempty"`
	// SSH port to connect to.
	// default = 22
	Port int `json:"port,omitempty"`
	// Directory on the remote server where the missing log files reside.
	// required = true
	Directory string `json:"directory,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// Reference to the snapshot for which to fetch logs.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-oracle-snapshot.json
	Snapshot string `json:"snapshot,omitempty"`
}

// UpgradeCheckResultsVersionParameters - Parameters used to modify an upgradeCheckResult. These parameters are
// mutually exclusive.
// extends TypedObject
type UpgradeCheckResultsVersionParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to a single upgrade check result.
	// format = objectReference
	// referenceTo = /delphix-upgrade-check-result.json
	// create = optional
	Reference string `json:"reference,omitempty"`
	// BundleID of upgrade check result(s).
	// create = optional
	BundleId string `json:"bundleId,omitempty"`
}

// Group - Database group.
// extends NamedUserObject
type Group struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// Optional description for the group.
	// maxLength = 1024
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
}

// MSSqlLinkData - MSSQL specific parameters for a link request.
// extends LinkData
type MSSqlLinkData struct {
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
	// Specifies the backup types validated sync will use to synchronize the
	// dSource with the source database.
	// enum = [TRANSACTION_LOG FULL_OR_DIFFERENTIAL FULL NONE]
	// default = TRANSACTION_LOG
	// create = optional
	ValidatedSyncMode string `json:"validatedSyncMode,omitempty"`
	// The credential for the source DB user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// Sync parameters for the container.
	// required = true
	SyncParameters MSSqlSyncParameters `json:"syncParameters,omitempty"`
	// Shared source database backup location.
	// create = optional
	// maxLength = 260
	SharedBackupLocation string `json:"sharedBackupLocation,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mssql-db-config.json
	// required = true
	Config string `json:"config,omitempty"`
	// The SQL instance on the PPT environment that we want to use for
	// pre-provisioning.
	// format = objectReference
	// referenceTo = /delphix-mssql-instance.json
	// required = true
	PptRepository string `json:"pptRepository,omitempty"`
	// Specifies whether Delphix should manage the backups for this container
	// and if so, specify whether the backups are compressed or uncompressed.
	// default = NOT_DELPHIX_MANAGED
	// create = optional
	// update = optional
	// enum = [DELPHIX_MANAGED_UNCOMPRESSED DELPHIX_MANAGED_COMPRESSED NOT_DELPHIX_MANAGED]
	DelphixManagedStatus string `json:"delphixManagedStatus,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// restoring from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPreScript string `json:"stagingPreScript,omitempty"`
	// The user name for the source DB user.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// The user for accessing the shared backup location.
	// create = optional
	// maxLength = 256
	BackupLocationUser string `json:"backupLocationUser,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// External file path.
	// create = optional
	// maxLength = 1024
	ExternalFilePath string `json:"externalFilePath,omitempty"`
	// The password for accessing the shared backup location.
	// create = optional
	BackupLocationCredentials *PasswordCredential `json:"backupLocationCredentials,omitempty"`
	// The encryption key to use when restoring encrypted backups.
	// create = optional
	EncryptionKey string `json:"encryptionKey,omitempty"`
	// Information about the host OS user on the PPT host to use for linking.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// format = objectReference
	PptHostUser string `json:"pptHostUser,omitempty"`
	// A user-provided PowerShell script or executable to run after restoring
	// from a backup during pre-provisioning.
	// create = optional
	// maxLength = 1024
	StagingPostScript string `json:"stagingPostScript,omitempty"`
	// Information about the host OS user on the source to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	SourceHostUser string `json:"sourceHostUser,omitempty"`
}

// JSOperation - An operation that occurred on a Jet Stream data layout.
// extends NamedUserObject
type JSOperation struct {
	// The time the operation finished. It will be null if the operation is
	// in progress.
	// format = date
	EndTime string `json:"endTime,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The operation performed.
	// enum = [REFRESH RESET CREATE_BRANCH DELETE_BRANCH CREATE_BOOKMARK DELETE_BOOKMARK ENABLE DISABLE ACTIVATE DEACTIVATE RECOVER RESTORE UNDO LOCK UNLOCK]
	Operation string `json:"operation,omitempty"`
	// The root action that spawned this operation.
	// referenceTo = /delphix-action.json
	// format = objectReference
	RootAction string `json:"rootAction,omitempty"`
	// The time that the data represented by this operation was active. It
	// will be null if the operation is in progress.
	// format = date
	DataTime string `json:"dataTime,omitempty"`
	// Was this operation perfomed with the force option, which means no
	// pre-operation snapshot was taken.
	ForceOption *bool `json:"forceOption,omitempty"`
	// The time the operation started.
	// format = date
	StartTime string `json:"startTime,omitempty"`
	// The data parent of the operation.
	DataParent JSDataParent `json:"dataParent,omitempty"`
	// The data layout that this operation was performed on.
	// format = objectReference
	// referenceTo = /delphix-js-data-layout.json
	DataLayout string `json:"dataLayout,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The bookmark that was created.
	// referenceTo = /delphix-js-bookmark.json
	// format = objectReference
	Bookmark string `json:"bookmark,omitempty"`
	// The user who performed the operation.
	// format = objectReference
	// referenceTo = /delphix-user.json
	User string `json:"user,omitempty"`
	// The branch that this operation was performed on.
	// format = objectReference
	// referenceTo = /delphix-js-branch.json
	Branch string `json:"branch,omitempty"`
	// Plain text description of the operation.
	Description string `json:"description,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OKResult - Result of a successful API call.
// extends CallResult
type OKResult struct {
	// Result of the operation. This will be specific to the API being
	// invoked.
	Result string `json:"result,omitempty"`
	// Reference to the job started by the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-job.json
	Job string `json:"job,omitempty"`
	// Reference to the action associated with the operation, if any.
	// format = objectReference
	// referenceTo = /delphix-action.json
	Action string `json:"action,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether an error occurred during the call.
	// enum = [OK ERROR]
	Status string `json:"status,omitempty"`
}

// AlertActionEmailList - Alert action to email a list of users in response to an alert.
// extends AlertActionEmail
type AlertActionEmailList struct {
	// Email format to use. The HTML format will generate a multipart message
	// containing both HTML and plain text. The TEXT format will explicitly
	// generate text-only mail. The JSON format will generate a JSON object
	// identical to the $Alert format returned through the web services API.
	// create = optional
	// update = optional
	// enum = [HTML TEXT JSON]
	// default = HTML
	Format string `json:"format,omitempty"`
	// List of email addresses to send mail to.
	// required = true
	Addresses []string `json:"addresses,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// OracleClusterNodeCreateParameters - The parameters used for oracle cluster node operations.
// extends TypedObject
type OracleClusterNodeCreateParameters struct {
	// The host object associated with the cluster node.
	// create = required
	HostParameters HostCreateParameters `json:"hostParameters,omitempty"`
	// The list of virtual IPs belonging to this node.
	// create = optional
	VirtualIPs []*OracleVirtualIP `json:"virtualIPs,omitempty"`
	// The cluster to which the node belongs.
	// format = objectReference
	// referenceTo = /delphix-oracle-cluster.json
	// create = optional
	Cluster string `json:"cluster,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the cluster node.
	// create = optional
	Name string `json:"name,omitempty"`
}

// AppDataStagedSourceConfig - An AppData source config with a staging source.
// extends AppDataSourceConfig
type AppDataStagedSourceConfig struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// The name of the config.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// The user used to create and manage the configuration.
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	// format = objectReference
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-appdata-source-repository.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The toolkit associated with this source config.
	// format = objectReference
	// referenceTo = /delphix-toolkit.json
	Toolkit string `json:"toolkit,omitempty"`
	// The list of parameters specified by the source config schema in the
	// toolkit. If no schema is specified, this list is empty.
	// create = optional
	// update = optional
	Parameters *Json `json:"parameters,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Whether this source should be used for linking.
	// create = optional
	// update = optional
	// default = true
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
}

// ValidateSMTPParameters - Validate SMTP configuration without committing it by sending mail to
// the specified address(es).
// extends TypedObject
type ValidateSMTPParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// SMTP configuration to use for validation.
	// required = true
	Config *SMTPConfig `json:"config,omitempty"`
	// List of email addresses to send test email to.
	// required = true
	Addresses []string `json:"addresses,omitempty"`
}

// CompatibleRepositoriesResult - Result of a compatible repositories request.
// extends TypedObject
type CompatibleRepositoriesResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The criteria matched to select compatible repositories.
	// required = true
	Criteria *CompatibilityCriteria `json:"criteria,omitempty"`
	// The list of compatible repositories.
	// required = true
	Repositories []SourceRepository `json:"repositories,omitempty"`
}

// MySQLLinkedSource - A linked MySQL source.
// extends MySQLSource
type MySQLLinkedSource struct {
	// MySQL database configuration parameter overrides.
	// create = optional
	// update = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// The staging source for validated sync of the database.
	// format = objectReference
	// referenceTo = /delphix-mysql-staging-source.json
	StagingSource string `json:"stagingSource,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-mysql-server-config.json
	// create = required
	// update = optional
	Config string `json:"config,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
}

// MSSqlVirtualSource - A virtual MSSQL source.
// extends MSSqlSource
type MSSqlVirtualSource struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// A user-provided PowerShell script or executable to run prior to
	// provisioning.
	// create = optional
	// update = optional
	// maxLength = 256
	PreScript string `json:"preScript,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
	// Reference to the container being fed by this source, if any.
	// referenceTo = /delphix-container.json
	// format = objectReference
	Container string `json:"container,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Reference to the configuration for the source.
	// format = objectReference
	// referenceTo = /delphix-source-config.json
	// create = optional
	Config string `json:"config,omitempty"`
	// MSSQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Runtime properties of this source.
	Runtime *MSSqlSourceRuntime `json:"runtime,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// The base mount point for the iSCSI LUN mounts.
	// maxLength = 256
	MountBase string `json:"mountBase,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	// default = true
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// A user-provided PowerShell script or executable to run after
	// provisioning.
	// update = optional
	// maxLength = 256
	// create = optional
	PostScript string `json:"postScript,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
}

// AppDataSyncParameters - The parameters to use as input to sync an AppData source.
// extends SyncParameters
type AppDataSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether or not to force a non-incremental load of data prior to taking
	// a snapshot.
	// default = false
	Resync *bool `json:"resync,omitempty"`
}

// NTPConfig - NTP (Network Time Protocol) configuration.
// extends TypedObject
type NTPConfig struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If true, then time is synchronized with the configured NTP servers.
	// The management service is automatically restarted if this value is
	// changed.
	// update = optional
	// default = false
	Enabled *bool `json:"enabled,omitempty"`
	// A list of NTP servers to use for synchronization. At least one server
	// must be specified if multicast is not being used.
	// update = optional
	Servers []string `json:"servers,omitempty"`
	// If true, discover NTP servers using multicast.
	// default = false
	// update = optional
	UseMulticast *bool `json:"useMulticast,omitempty"`
	// Address to use for multicast NTP discovery. This is only valid when
	// 'useMulticast' is set.
	// format = ipv4Address
	// default = 224.0.1.1
	// update = optional
	MulticastAddress string `json:"multicastAddress,omitempty"`
}

// JSWeeklyOperationCount - Information about the number of operations on a data container each
// week for up to 30 weeks.
// extends JSUsageData
type JSWeeklyOperationCount struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object the usage data is centered around.
	// format = objectReference
	// referenceTo = /delphix-named-user-object.json
	UsageObject string `json:"usageObject,omitempty"`
	// The date at the beginning of the time period this datapoint
	// corresponds to. The time period itself varies between datapoint types.
	// format = date
	StartDate string `json:"startDate,omitempty"`
	// The number of operations run against a data container in the specified
	// week.
	WeeklyCount int `json:"weeklyCount,omitempty"`
	// The total time spent in seconds running all operations during the
	// specified week.
	// units = sec
	WeeklyDuration int `json:"weeklyDuration,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// ASEExportParameters - The parameters to use as input to export SAP ASE databases.
// extends DbExportParameters
type ASEExportParameters struct {
	// The source config to use when creating the exported DB.
	// required = true
	SourceConfig ASEDBConfig `json:"sourceConfig,omitempty"`
	// If specified, then take the exported database through recovery
	// procedures, if necessary, to reach a consistent point.
	// default = true
	// create = optional
	RecoverDatabase *bool `json:"recoverDatabase,omitempty"`
	// Database-specific configuration parameters.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to base export on.
	// required = true
	// properties = map[type:map[default:TimeflowPointSemantic]]
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
}

// MySQLLinkData - MySQL specific parameters for a link request.
// extends LinkData
type MySQLLinkData struct {
	// OS user on the staging host to use for linking.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	StagingHostUser string `json:"stagingHostUser,omitempty"`
	// Sync parameters for the container.
	// required = true
	SyncParameters MySQLSyncParameters `json:"syncParameters,omitempty"`
	// The database username.
	// required = true
	DbUser string `json:"dbUser,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	Operations *LinkedSourceOperations `json:"operations,omitempty"`
	// MySQL database configuration parameter overrides.
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// The MySQL installation on the staging environment that will be used
	// for validated sync.
	// format = objectReference
	// referenceTo = /delphix-mysql-install.json
	// required = true
	StagingRepository string `json:"stagingRepository,omitempty"`
	// The credentials for the database user.
	// required = true
	DbCredentials *PasswordCredential `json:"dbCredentials,omitempty"`
	// The port on the staging host that the MySQL staging server can listen
	// on for TCP/IP connections.
	// required = true
	// minimum = 1
	// maximum = 65535
	StagingPort int `json:"stagingPort,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-mysql-server-config.json
	// required = true
	// format = objectReference
	Config string `json:"config,omitempty"`
	// Policies for managing LogSync and SnapSync across sources.
	// create = optional
	SourcingPolicy *SourcingPolicy `json:"sourcingPolicy,omitempty"`
}

// OracleExport - The mutable state of an Oracle database export.
// extends TypedObject
type OracleExport struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// DSP options for export.
	DspOptions *DSPOptions `json:"dspOptions,omitempty"`
	// Number of files to stream in parallel across the network.
	// minimum = 1
	// default = 3
	FileParallelism int `json:"fileParallelism,omitempty"`
}

// SnapshotSpaceMap - Mapping of containers and snapshots to their respective space usage.
// extends TypedObject
type SnapshotSpaceMap struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of space, per object, in bytes, that it uses.
	SizeMap map[string]string `json:"sizeMap,omitempty"`
}

// PgSQLDBClusterConfig - Configuration information for a PostgreSQL database cluster.
// extends SourceConfig
type PgSQLDBClusterConfig struct {
	// The port on which the PostgresSQL server for the cluster is listening.
	// update = optional
	// minimum = 1
	// maximum = 65535
	// create = required
	Port int `json:"port,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Whether this source should be used for linking.
	// update = optional
	// default = true
	// create = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The password of the database cluster user.
	// create = optional
	// update = optional
	Credentials string `json:"credentials,omitempty"`
	// The database that must be used to run SQL queries against this
	// cluster.
	// create = optional
	// update = optional
	// maxLength = 256
	ConnectionDatabase string `json:"connectionDatabase,omitempty"`
	// The data directory for the PostgreSQL cluster.
	// create = optional
	ClusterDataDirectory string `json:"clusterDataDirectory,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Whether this source was discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// The user used to create and manage the configuration.
	// format = objectReference
	// referenceTo = /delphix-source-environment-user.json
	// create = optional
	// update = optional
	EnvironmentUser string `json:"environmentUser,omitempty"`
	// The object reference of the source repository.
	// format = objectReference
	// referenceTo = /delphix-pgsql-install.json
	// create = required
	// update = optional
	Repository string `json:"repository,omitempty"`
	// The username of the database cluster user.
	// maxLength = 256
	// update = optional
	User string `json:"user,omitempty"`
}

// HistoricalConsumerCapacityData - Historical data about a particular capacity consumer.
// extends BaseConsumerCapacityData
// cliVisibility = [DOMAIN]
type HistoricalConsumerCapacityData struct {
	// Name of this container's group.
	GroupName string `json:"groupName,omitempty"`
	// Reference to the container.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// Reference to this container's group.
	// format = objectReference
	// referenceTo = /delphix-group.json
	Group string `json:"group,omitempty"`
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Statistics for this consumer.
	Breakdown *CapacityBreakdown `json:"breakdown,omitempty"`
	// Name of the container.
	Name string `json:"name,omitempty"`
	// Container from which this TimeFlow was provisioned.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Parent string `json:"parent,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// SystemInitializationParameters - Parameters used for intializing an engine.
// extends TypedObject
type SystemInitializationParameters struct {
	// List of storage devices to use.
	// create = required
	Devices []string `json:"devices,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Name of the default domain administrator to create.
	// default = delphix_admin
	// create = optional
	// pattern = ^[a-zA-Z][-_.a-zA-Z0-9]*$
	// minLength = 1
	// maxLength = 256
	DefaultUser string `json:"defaultUser,omitempty"`
	// Password to use for the default domain administrator.
	// format = password
	// default = delphix
	// create = optional
	DefaultPassword string `json:"defaultPassword,omitempty"`
}

// ResetIgnoredFaultsParameters - The parameters to use as input when marking selected ignored faults as
// resolved.
// extends TypedObject
type ResetIgnoredFaultsParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Input list of ignored faults which will be marked as resolved.
	// required = true
	FaultReferences []string `json:"faultReferences,omitempty"`
}

// SNMPManager - SNMP manager configuration.
// extends PersistentObject
// cliVisibility = [SYSTEM]
type SNMPManager struct {
	// SNMP manager host.
	// format = host
	// create = required
	// update = optional
	Address string `json:"address,omitempty"`
	// SNMP manager port number.
	// maximum = 65535
	// default = 162
	// create = optional
	// update = optional
	// minimum = 1
	Port int `json:"port,omitempty"`
	// SNMP manager community string.
	// update = optional
	// create = required
	CommunityString string `json:"communityString,omitempty"`
	// Describes if the most recent attempt to send a trap succeeded or
	// failed.
	// enum = [FAILED SUCCEEDED PENDING UNCHECKED]
	// default = PENDING
	// create = readonly
	// update = readonly
	LastSendStatus string `json:"lastSendStatus,omitempty"`
	// True if INFORM messages are to be sent to this manager, false for TRAP
	// messages.
	// update = optional
	// default = false
	// create = optional
	UseInform *bool `json:"useInform,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
}

// ProxyService - Proxy service configuration.
// extends TypedObject
// cliVisibility = [SYSTEM]
type ProxyService struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// HTTPS proxy configuration.
	// required = false
	// update = optional
	Https *ProxyConfiguration `json:"https,omitempty"`
	// SOCKS proxy configuration.
	// update = optional
	// required = false
	Socks *ProxyConfiguration `json:"socks,omitempty"`
}

// ASEInstance - The SAP ASE source repository.
// extends SourceRepository
type ASEInstance struct {
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Database page size for the SAP ASE instance.
	PageSize int `json:"pageSize,omitempty"`
	// The Kerberos SPN of the database.
	// create = optional
	// update = optional
	ServicePrincipalName string `json:"servicePrincipalName,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// The name of the SAP ASE instance.
	// create = required
	InstanceName string `json:"instanceName,omitempty"`
	// The network ports for connecting to the SAP ASE instance.
	// create = required
	// update = optional
	Ports []int `json:"ports,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// The username of the database user.
	// create = optional
	// update = optional
	// maxLength = 256
	DbUser string `json:"dbUser,omitempty"`
	// The uid of the account the SAP ASE instance is running as.
	// create = readonly
	// update = readonly
	InstanceOwnerUid int `json:"instanceOwnerUid,omitempty"`
	// The credentials of the database user.
	// create = optional
	// update = optional
	// properties = map[type:map[default:PasswordCredential type:string description:Object type. required:true format:type]]
	Credentials Credential `json:"credentials,omitempty"`
	// True if the SAP ASE instance was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The gid of the account the SAP ASE instance is running as.
	// create = readonly
	// update = readonly
	InstanceOwnerGid int `json:"instanceOwnerGid,omitempty"`
	// The SAP ASE instance home.
	// create = required
	// update = optional
	InstallationPath string `json:"installationPath,omitempty"`
	// The username of the account the SAP ASE instance is running as.
	// update = optional
	// create = required
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// The path to the isql binary to use for this SAP ASE instance.
	// create = optional
	// update = optional
	IsqlPath string `json:"isqlPath,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// default = false
	// update = optional
	Staging *bool `json:"staging,omitempty"`
}

// AppDataAdditionalMountPoint - Specifies an additional location on which to mount a subdirectory of
// an AppData container.
// extends TypedObject
type AppDataAdditionalMountPoint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Absolute path on the target environment were the filesystem should be
	// mounted.
	// create = required
	// update = optional
	// format = unixpath
	MountPath string `json:"mountPath,omitempty"`
	// Reference to the environment on which the file system will be mounted.
	// format = objectReference
	// referenceTo = /delphix-host-environment.json
	// create = required
	// update = optional
	Environment string `json:"environment,omitempty"`
	// Relative path within the container of the directory that should be
	// mounted.
	// format = unixpath
	// create = optional
	// update = optional
	SharedPath string `json:"sharedPath,omitempty"`
}

// SwitchTimeflowParameters - The input parameters used for the TimeFlow switch operation.
// extends TypedObject
type SwitchTimeflowParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The reference to the target TimeFlow that should be made the current
	// TimeFlow.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-timeflow.json
	Timeflow string `json:"timeflow,omitempty"`
}

// UpgradeNotification - An object to track when an upgrade is happening to trigger UI
// response.
// extends Notification
type UpgradeNotification struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// VfsOpsDatapointStream - A stream of datapoints from a VFS_OPS analytics slice.
// extends DatapointStream
type VfsOpsDatapointStream struct {
	// Whether reads were cached.
	Cached *bool `json:"cached,omitempty"`
	// I/O operation type.
	// enum = [read write]
	Op string `json:"op,omitempty"`
	// Path of the affected file.
	// format = unixpath
	Path string `json:"path,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The set of datapoints in the stream.
	Datapoints []Datapoint `json:"datapoints,omitempty"`
	// Whether writes were synchronous.
	Sync *bool `json:"sync,omitempty"`
}

// CurrentSystemCapacityData - Capacity data for the entire system.
// extends BaseSystemCapacityData
// cliVisibility = [DOMAIN]
type CurrentSystemCapacityData struct {
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdown `json:"source,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdown `json:"virtual,omitempty"`
	// Total storage space (used and unused).
	// units = B
	// base = 1024
	TotalSpace float64 `json:"totalSpace,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// HostPrivilegeElevationProfile - Profile for elevating user privileges on a host.
// extends PersistentObject
// cliVisibility = [DOMAIN]
type HostPrivilegeElevationProfile struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Privilege elevation profile version.
	// update = optional
	// create = required
	Version string `json:"version,omitempty"`
	// The privilege elevation profile name.
	// create = required
	// update = optional
	Name string `json:"name,omitempty"`
	// True if this is the default privilege elevation profile for new
	// environments.
	// default = false
	// update = readonly
	IsDefault *bool `json:"isDefault,omitempty"`
}

// AndFilter - A container filter that combines other filters together using AND
// logic.
// extends AlertFilter
type AndFilter struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Filters which are combined together using AND logic.
	// update = optional
	// minItems = 2
	// create = required
	SubFilters []AlertFilter `json:"subFilters,omitempty"`
}

// MySQLPlatformParameters - MySQL platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type MySQLPlatformParameters struct {
	// The port number used for provisioning a MySQL container during
	// transformation application. This port must be available when the
	// transformation is applied so that the temporary VDB created during the
	// transformation process can start and listen to this port.
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// BatchContainerRefreshParameters - The parameters to use as input to batch container refresh requests.
// extends TypedObject
type BatchContainerRefreshParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// A semantic description of a TimeFlow location.
	// enum = [LATEST_POINT LATEST_SNAPSHOT]
	// default = LATEST_POINT
	// required = false
	Location string `json:"location,omitempty"`
	// Containers to refresh.
	// required = true
	// minItems = 1
	Containers []string `json:"containers,omitempty"`
}

// VMwarePlatformParameters - VMware platform-specific parameters that are stored on a
// transformation.
// extends BasePlatformParameters
type VMwarePlatformParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// AuthFilterParameters - The parameters to use as input to filter a list of objects or object
// type by a permission.
// extends TypedObject
type AuthFilterParameters struct {
	// The object type on which to perform filtering. This option is mutually
	// exclusive with the "objects" field.
	// format = type
	ObjectType string `json:"objectType,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The permission to filter by.
	// format = objectReference
	// referenceTo = /delphix-permission.json
	Permission string `json:"permission,omitempty"`
	// The list of objects to filter. This option is mutually exclusive with
	// the "objectType" field.
	Objects []string `json:"objects,omitempty"`
}

// HistoricalSystemCapacityData - Capacity data for the entire system.
// extends BaseSystemCapacityData
// cliVisibility = [DOMAIN]
type HistoricalSystemCapacityData struct {
	// Time at which this information was sampled.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// Statistics for dSources in this aggregation.
	Source *CapacityBreakdown `json:"source,omitempty"`
	// Statistics for VDBs in this aggregation.
	Virtual *CapacityBreakdown `json:"virtual,omitempty"`
	// Total storage space (used and unused).
	// units = B
	// base = 1024
	TotalSpace float64 `json:"totalSpace,omitempty"`
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
}

// User - Delphix users.
// extends NamedUserObject
// cliVisibility = [DOMAIN SYSTEM]
type User struct {
	// Last name of user.
	// create = optional
	// update = optional
	// maxLength = 64
	LastName string `json:"lastName,omitempty"`
	// Session timeout in minutes.
	// minimum = 1
	// default = 30
	// units = min
	// create = optional
	// update = optional
	SessionTimeout int `json:"sessionTimeout,omitempty"`
	// Mobile phone number of user.
	// maxLength = 32
	// create = optional
	// update = optional
	MobilePhoneNumber string `json:"mobilePhoneNumber,omitempty"`
	// Credential used for authentication.
	// create = optional
	Credential *PasswordCredential `json:"credential,omitempty"`
	// Email address for the user.
	// format = email
	// maxLength = 256
	// create = optional
	// update = optional
	EmailAddress string `json:"emailAddress,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Work phone number of user.
	// maxLength = 32
	// create = optional
	// update = optional
	WorkPhoneNumber string `json:"workPhoneNumber,omitempty"`
	// User authentication type.
	// enum = [LDAP NATIVE SAML]
	// create = optional
	AuthenticationType string `json:"authenticationType,omitempty"`
	// Type of user.
	// enum = [SYSTEM DOMAIN]
	// default = DOMAIN
	// create = optional
	UserType string `json:"userType,omitempty"`
	// Public key used for authentication.
	// create = optional
	// update = optional
	PublicKey string `json:"publicKey,omitempty"`
	// Home phone number of user.
	// maxLength = 32
	// create = optional
	// update = optional
	HomePhoneNumber string `json:"homePhoneNumber,omitempty"`
	// Principal name used for authentication.
	// create = optional
	// update = optional
	Principal string `json:"principal,omitempty"`
	// True if the user is currently enabled and can log into the system.
	// default = true
	// create = optional
	Enabled *bool `json:"enabled,omitempty"`
	// First name of user.
	// maxLength = 64
	// create = optional
	// update = optional
	FirstName string `json:"firstName,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// referenceTo = /delphix-namespace.json
	// format = objectReference
	Namespace string `json:"namespace,omitempty"`
	// Name of user.
	// create = required
	Name string `json:"name,omitempty"`
	// Preferred locale as an IETF BCP 47 language tag, defaults to 'en-US'.
	// maxLength = 16
	// create = optional
	// update = optional
	// format = locale
	// default = en-US
	Locale string `json:"locale,omitempty"`
	// Whether the user's password should be updated and why.
	// enum = [NONE FIRST_LOGIN ADMIN_REQUEST PASSWORD_POLICY]
	// default = NONE
	// update = optional
	PasswordUpdateRequest string `json:"passwordUpdateRequest,omitempty"`
	// True if this is the default user and cannot be deleted.
	// create = optional
	IsDefault *bool `json:"isDefault,omitempty"`
}

// IoOpsDatapoint - An analytics datapoint generated by the DISK_OPS, DxFS_OPS,
// DxFS_IO_QUEUE_OPS, iSCSI_OPS, NFS_OPS, or VFS_OPS statistic types.
// extends Datapoint
type IoOpsDatapoint struct {
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// I/O throughput in bytes.
	// base = 1024
	// units = B
	Throughput int `json:"throughput,omitempty"`
	// Number of I/O operations.
	Count int `json:"count,omitempty"`
	// I/O latencies in nanoseconds.
	Latency map[string]string `json:"latency,omitempty"`
	// Average I/O latency in nanoseconds.
	AvgLatency int `json:"avgLatency,omitempty"`
	// I/O sizes in bytes.
	Size map[string]string `json:"size,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// IntegerEqualConstraint - Constraint placed on a numerical axis of a particular analytics slice.
// extends IntegerConstraint
type IntegerEqualConstraint struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The name of the axis being constrained.
	// create = required
	AxisName string `json:"axisName,omitempty"`
	// The axis values must equal this value.
	// create = required
	Equals int `json:"equals,omitempty"`
}

// DatabaseTemplate - Parameter configuration for virtual databases.
// extends NamedUserObject
type DatabaseTemplate struct {
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// User provided description for this template.
	// create = optional
	// update = optional
	Description string `json:"description,omitempty"`
	// The type of the source associated with the template.
	// create = required
	// update = optional
	// enum = [MySQLVirtualSource OracleVirtualSource MSSqlVirtualSource PgSQLVirtualSource]
	// format = type
	SourceType string `json:"sourceType,omitempty"`
	// A name/value map of string configuration parameters.
	// create = required
	// update = optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
}

// SerializationPoint - A serialization point represents the data and metadata associated with
// a replication spec at a point in time.
// extends UserObject
type SerializationPoint struct {
	// Object name.
	// create = optional
	// update = optional
	// maxLength = 256
	// format = objectName
	Name string `json:"name,omitempty"`
	// An arbitrary string used to map the serialization point to a
	// corresponding replication spec or namespace.
	Tag string `json:"tag,omitempty"`
	// Timestamp of the data being stored in the serialization point.
	// format = date
	DataTimestamp string `json:"dataTimestamp,omitempty"`
	// Average throughput of the transfer of the serialization point
	// (bytes/s).
	// base = 1024
	// units = B/s
	AverageThroughput float64 `json:"averageThroughput,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Bytes of the serialization point which have been transferred.
	// units = B
	// base = 1024
	BytesTransferred int `json:"bytesTransferred,omitempty"`
	// The elapsed time spent sending the serialization point (nanoseconds).
	// base = 1024
	// units = nsec
	ElapsedTimeNanos float64 `json:"elapsedTimeNanos,omitempty"`
}

// ContainerUtilization - Represents the utilization of all containers during a particular
// period of time.
// extends TypedObject
// cliVisibility = []
type ContainerUtilization struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Reference to the container whose utilization we are describing.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// A list of container utilization statistics corresponding to this
	// period of time, one for each sampling interval.
	Utilization []*ContainerUtilizationInterval `json:"utilization,omitempty"`
	// True if this container has been deleted.
	Deleted *bool `json:"deleted,omitempty"`
	// True if the current user does not have access to this container.
	Hidden *bool `json:"hidden,omitempty"`
}

// FileMappingResult - Result of a file mapping request.
// extends TypedObject
type FileMappingResult struct {
	// Mapped files.
	MappedFiles map[string]string `json:"mappedFiles,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// CpuUtilDatapoint - An analytics datapoint generated by the CPU_UTIL statistic type.
// extends Datapoint
type CpuUtilDatapoint struct {
	// Kernel time in milliseconds.
	Kernel int `json:"kernel,omitempty"`
	// User time in milliseconds.
	User int `json:"user,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The time this datapoint was collected.
	// format = date
	Timestamp string `json:"timestamp,omitempty"`
	// DTrace time in milliseconds (subset of time in kernel).
	Dtrace int `json:"dtrace,omitempty"`
	// Idle time in milliseconds.
	Idle int `json:"idle,omitempty"`
}

// MySQLDBContainerRuntime - Runtime properties of a MySQL database container.
// extends DBContainerRuntime
type MySQLDBContainerRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The pre-provisioning runtime for the container.
	PreProvisioningStatus *PreProvisioningRuntime `json:"preProvisioningStatus,omitempty"`
	// True if the LogSync is enabled and running for this container.
	LogSyncActive *bool `json:"logSyncActive,omitempty"`
}

// MSSqlAvailabilityGroup - The representation of a SQL Server Availability Group.
// extends MSSqlBaseClusterRepository
type MSSqlAvailabilityGroup struct {
	// Object name.
	// create = readonly
	// update = readonly
	Name string `json:"name,omitempty"`
	// Version of the repository.
	// create = optional
	// update = optional
	Version string `json:"version,omitempty"`
	// Flag indicating whether the repository should be used for
	// provisioning.
	// default = true
	// update = optional
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Internal version of the SQL Server instance.
	InternalVersion int `json:"internalVersion,omitempty"`
	// The list of listeners belonging to this repository.
	Listeners []MSSqlBaseClusterListener `json:"listeners,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// referenceTo = /delphix-persistent-object.json
	// format = objectReference
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Reference to the environment containing this repository.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
	// Flag indicating whether this repository can be used by the Delphix
	// Engine for internal processing.
	// update = optional
	// default = false
	Staging *bool `json:"staging,omitempty"`
	// Flag indicating whether the repository should be used for linking.
	// default = true
	// update = optional
	LinkingEnabled *bool `json:"linkingEnabled,omitempty"`
	// The list of MSSQL Cluster instances belonging to this repository.
	Instances []MSSqlBaseClusterInstance `json:"instances,omitempty"`
}

// MSSqlFailoverClusterInstance - The representation of a SQL Server Instance on a clustered node for
// Failover Clusters.
// extends MSSqlBaseClusterInstance
type MSSqlFailoverClusterInstance struct {
	// The version of the SQL Server Instance.
	Version string `json:"version,omitempty"`
	// The owner of the SQL Server Instance.
	InstanceOwner string `json:"instanceOwner,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The Servername of the SQL Server Instance.
	ServerName string `json:"serverName,omitempty"`
	// A reference to the Windows Cluster Node for this instance.
	// format = objectReference
	// referenceTo = /delphix-windows-cluster-node.json
	Node string `json:"node,omitempty"`
	// The port to connect to the SQL Server Instance.
	Port int `json:"port,omitempty"`
	// The name of the SQL Server Instance.
	// maxLength = 16
	Name string `json:"name,omitempty"`
}

// OracleScanListener - An Oracle scan listener.
// extends OracleListener
type OracleScanListener struct {
	// The list of protocol addresses for this listener. These are used for
	// the local_listener parameter when provisioning VDBs.
	// minItems = 1
	// create = required
	// update = optional
	ProtocolAddresses []string `json:"protocolAddresses,omitempty"`
	// The list of client endpoints for this listener of the format
	// hostname:port. These are used when constructing the JDBC connection
	// string.
	// create = readonly
	// update = readonly
	ClientEndpoints []string `json:"clientEndpoints,omitempty"`
	// Whether this listener was automatically discovered.
	Discovered *bool `json:"discovered,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Object name.
	// create = required
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Reference to the environment this listener is associated with.
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	// create = required
	Environment string `json:"environment,omitempty"`
}

// DSPBestParameters - DSP parameters, found by autotuner, that give the highest throughput
// for a certain target.
// extends PersistentObject
type DSPBestParameters struct {
	// The number of connections used to achieve maximum throughput.
	NumConnections int `json:"numConnections,omitempty"`
	// Whether the test is testing connectivity to a Delphix Engine or remote
	// host.
	// enum = [REMOTE_HOST DELPHIX_ENGINE]
	DestinationType string `json:"destinationType,omitempty"`
	// The size of the send and receive socket buffers in bytes used to
	// achieve maximum throughput.
	// base = 1024
	// units = B
	BufferSize int `json:"bufferSize,omitempty"`
	// Information used when running a test to another Delphix Engine.
	RemoteDelphixEngineInfo *RemoteDelphixEngineInfo `json:"remoteDelphixEngineInfo,omitempty"`
	// The remote host for the test. The host must be part of an existing
	// environment.
	// format = objectReference
	// referenceTo = /delphix-host.json
	RemoteHost string `json:"remoteHost,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// The queue depth used to achieve maximum throughput.
	QueueDepth int `json:"queueDepth,omitempty"`
	// The average throughput measured.
	// base = 1024
	// units = bps
	Throughput float64 `json:"throughput,omitempty"`
}

// StagingCompatibilityParameters - The criteria necessary to select valid repositories for staging.
// extends CompatibleRepositoriesParameters
type StagingCompatibilityParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Restrict returned repositories to this environment.
	// create = optional
	// update = optional
	// format = objectReference
	// referenceTo = /delphix-source-environment.json
	Environment string `json:"environment,omitempty"`
	// The repository to use as a source of compatibility information.
	// required = true
	// format = objectReference
	// referenceTo = /delphix-source-repository.json
	SourceRepository string `json:"sourceRepository,omitempty"`
}

// TimeflowBookmarkCreateParameters - The parameters to use as input to create TimeFlow bookmarks.
// extends TypedObject
type TimeflowBookmarkCreateParameters struct {
	// The TimeFlow point which is referenced by this bookmark.
	// required = true
	TimeflowPoint TimeflowPoint `json:"timeflowPoint,omitempty"`
	// The bookmark name.
	// required = true
	Name string `json:"name,omitempty"`
	// A tag for the bookmark that can be used to group bookmarks together or
	// qualify the type of the bookmark.
	// create = optional
	// maxLength = 64
	Tag string `json:"tag,omitempty"`
	// Indicates whether retention should be allowed to clean up the TimeFlow
	// bookmark and associated data.
	// create = optional
	RetentionProof *bool `json:"retentionProof,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MySQLSourceRuntime - Non-persistent runtime properties of a MySQL source.
// extends SourceRuntime
type MySQLSourceRuntime struct {
	// The reason why the source is not JDBC accessible.
	NotAccessibleReason string `json:"notAccessibleReason,omitempty"`
	// Status indicating whether the source is enabled. A source has a
	// 'PARTIAL' status if its sub-sources are not all enabled.
	// enum = [ENABLED PARTIAL DISABLED]
	Enabled string `json:"enabled,omitempty"`
	// Status of the source. 'Unknown' if all attempts to connect to the
	// source failed.
	// enum = [RUNNING INACTIVE PENDING CANCELED FAILED CHECKING UNKNOWN]
	Status string `json:"status,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if the source is JDBC accessible. If false then no properties can
	// be retrieved.
	Accessible *bool `json:"accessible,omitempty"`
	// The time that the 'accessible' propery was last checked.
	// format = date
	AccessibleTimestamp string `json:"accessibleTimestamp,omitempty"`
	// Size of the database in bytes.
	// units = B
	// base = 1024
	DatabaseSize float64 `json:"databaseSize,omitempty"`
}

// SamlInfo - Global SAML information.
// extends TypedObject
// cliVisibility = [SYSTEM]
type SamlInfo struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Whether or not SAML authentication is configured and enabled for this
	// Delphix Engine.
	Enabled *bool `json:"enabled,omitempty"`
}

// MSSqlTimeflow - TimeFlow representing historical data for a particular timeline within
// a data container.
// extends Timeflow
type MSSqlTimeflow struct {
	// Object name.
	// update = optional
	// maxLength = 256
	// create = required
	Name string `json:"name,omitempty"`
	// The origin point on the parent TimeFlow from which this TimeFlow was
	// provisioned. This will not be present for TimeFlows derived from
	// linked sources.
	ParentPoint *MSSqlTimeflowPoint `json:"parentPoint,omitempty"`
	// Reference to the parent snapshot that serves as the provisioning base
	// for this object. This may be different from the snapshot within the
	// parent point, and is only present for virtual TimeFlows.
	// format = objectReference
	// referenceTo = /delphix-timeflow-snapshot.json
	ParentSnapshot string `json:"parentSnapshot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// The source action that created the TimeFlow.
	// enum = [INITIAL INDETERMINATE REFRESH ROLLBACK TEMPORARY TRANSFORMATION V2P PDB_PLUG WAREHOUSE ORACLE_LIVE_SOURCE_RESYNC]
	CreationType string `json:"creationType,omitempty"`
	// MSSQL-specific recovery branch identifier for this TimeFlow.
	DatabaseGuid string `json:"databaseGuid,omitempty"`
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Reference to the data container (database) for this TimeFlow.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
}

// FileUploadResult - Result of a file upload request.
// extends FileProcessingResult
type FileUploadResult struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// URL to download from or upload to.
	Url string `json:"url,omitempty"`
	// Token to pass as parameter to identify the file.
	Token string `json:"token,omitempty"`
}

// SSHConnectivity - Mechanism to test SSH connectivity of arbitrary hosts.
// extends TypedObject
type SSHConnectivity struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Mechanism to use for ssh host verification.
	// required = false
	SshVerificationStrategy SshVerificationStrategy `json:"sshVerificationStrategy,omitempty"`
	// Target host name or IP address.
	// required = true
	Address string `json:"address,omitempty"`
	// SSH port on remote server.
	// default = 22
	// minimum = 1
	// maximum = 65535
	Port int `json:"port,omitempty"`
	// User name.
	// required = true
	Username string `json:"username,omitempty"`
	// User credentials.
	// required = true
	Credentials Credential `json:"credentials,omitempty"`
}

// MySQLSnapshotRuntime - Runtime (non-persistent) properties of a MySQL TimeFlow snapshot.
// extends SnapshotRuntime
type MySQLSnapshotRuntime struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// True if this snapshot can be used as the basis for provisioning a new
	// TimeFlow.
	Provisionable *bool `json:"provisionable,omitempty"`
}

// APIVersion - Describes an API version.
// extends TypedObject
type APIVersion struct {
	// Object type.
	// required = true
	// format = type
	Type *string `json:"type,omitempty"`
	// Micro API version number.
	// minimum = 0
	// required = true
	Micro *int `json:"micro,omitempty"`
	// Major API version number.
	// minimum = 0
	// required = true
	Major *int `json:"major,omitempty"`
	// Minor API version number.
	// minimum = 0
	// required = true
	Minor *int `json:"minor,omitempty"`
}

// StorageDeviceRemovalVerifyResult - The .
// extends TypedObject
type StorageDeviceRemovalVerifyResult struct {
	// Amount of memory to be used by mappings if this device is removed, in
	// bytes.
	// base = 1024
	// units = B
	NewMappingMemory float64 `json:"newMappingMemory,omitempty"`
	// Free space of the pool before this device is removed, in bytes.
	// base = 1024
	// units = B
	OldFreeBytes float64 `json:"oldFreeBytes,omitempty"`
	// Free space of the pool if this device is removed, in bytes.
	// base = 1024
	// units = B
	NewFreeBytes float64 `json:"newFreeBytes,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Amount of memory used by removal mappings before this device is
	// removed, in bytes.
	// base = 1024
	// units = B
	OldMappingMemory float64 `json:"oldMappingMemory,omitempty"`
}

// OracleRollbackParameters - The parameters to use as input to roll back Oracle databases.
// extends RollbackParameters
type OracleRollbackParameters struct {
	// The security credential of the user who has the required privileges to
	// run the rollback operation.
	Credential Credential `json:"credential,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The TimeFlow point, bookmark, or semantic location to roll the
	// database back to.
	// properties = map[type:map[default:TimeflowPointSemantic]]
	// required = true
	TimeflowPointParameters TimeflowPointParameters `json:"timeflowPointParameters,omitempty"`
	// The name of the user who has the required privileges to perform the
	// rollback operation.
	Username string `json:"username,omitempty"`
}

// XppStagingStatus - Status of the cross-platform provisioning staging environment.
// extends ChecklistItemDetail
type XppStagingStatus struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The status item name.
	Name string `json:"name,omitempty"`
	// Description of this status item.
	Description string `json:"description,omitempty"`
	// Status of this item.
	// enum = [SUCCESS WARNING ERROR]
	Status string `json:"status,omitempty"`
	// Status message, if applicable.
	Message string `json:"message,omitempty"`
}

// KeyPairCredential - The public key based security credential consisting of a user
// specified key pair.
// extends PublicKeyCredential
type KeyPairCredential struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// The private key in the key pair.
	// format = password
	// required = true
	PrivateKey string `json:"privateKey,omitempty"`
	// The public key in the key pair.
	// format = password
	// required = true
	PublicKey string `json:"publicKey,omitempty"`
}

// ResolveOrIgnoreSelectedFaultsParameters - The parameters to use as input when marking selected faults as
// resolved or ignored.
// extends TypedObject
type ResolveOrIgnoreSelectedFaultsParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// Input list of faults which will be marked as resolved or ignored.
	// required = true
	FaultReferences []string `json:"faultReferences,omitempty"`
	// Flag indicating whether to ignore the selected faults if they are
	// detected on the same objects in the future.
	// required = true
	// default = false
	Ignore *bool `json:"ignore,omitempty"`
}

// SshAcceptAlways - Key-verification strategy that always accepts the host's key.
// extends SshVerificationStrategy
type SshAcceptAlways struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
}

// MSSqlNewCopyOnlyFullBackupSyncParameters - The parameters to use as input to sync MSSQL databases using a new
// copy-only full backup taken by Delphix.
// extends MSSqlSyncParameters
type MSSqlNewCopyOnlyFullBackupSyncParameters struct {
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// If this parameter is set to true, Delphix will take a compressed copy
	// only full backup of the source database. When set to false, Delphix
	// will use the SQL Server instance's compression default.
	// default = false
	// required = true
	CompressionEnabled *bool `json:"compressionEnabled,omitempty"`
}

// MySQLVirtualSource - A virtual MySQL source.
// extends MySQLSource
type MySQLVirtualSource struct {
	// The object reference.
	// format = objectReference
	// referenceTo = /delphix-persistent-object.json
	Reference string `json:"reference,omitempty"`
	// Hosts that might affect operations on this source. Property will be
	// null unless the includeHosts parameter is set when listing sources.
	Hosts []string `json:"hosts,omitempty"`
	// User-specified operation hooks for this source.
	// create = optional
	// update = optional
	Operations *VirtualSourceOperations `json:"operations,omitempty"`
	// A user-provided description of the source.
	Description string `json:"description,omitempty"`
	// Reference to the container being fed by this source, if any.
	// format = objectReference
	// referenceTo = /delphix-container.json
	Container string `json:"container,omitempty"`
	// The base mount point for the NFS mounts.
	// maxLength = 256
	// create = required
	// update = optional
	MountBase string `json:"mountBase,omitempty"`
	// Indicates whether Delphix should automatically restart this virtual
	// source when target host reboot is detected.
	// create = required
	// update = optional
	AllowAutoVDBRestartOnHostReboot *bool `json:"allowAutoVDBRestartOnHostReboot,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Alternate namespace for this object, for replicated and restored
	// objects.
	// format = objectReference
	// referenceTo = /delphix-namespace.json
	Namespace string `json:"namespace,omitempty"`
	// Flag indicating whether the source is used as a staging source for
	// pre-provisioning. Staging sources are managed by the Delphix system.
	Staging *bool `json:"staging,omitempty"`
	// Reference to the configuration for the source.
	// referenceTo = /delphix-source-config.json
	// create = optional
	// format = objectReference
	Config string `json:"config,omitempty"`
	// Database file mapping rules.
	// create = optional
	FileMappingRules string `json:"fileMappingRules,omitempty"`
	// MySQL database configuration parameter overrides.
	// update = optional
	// create = optional
	ConfigParams map[string]string `json:"configParams,omitempty"`
	// Object name.
	// format = objectName
	// create = optional
	// update = optional
	// maxLength = 256
	Name string `json:"name,omitempty"`
	// Runtime properties of this source.
	Runtime *MySQLSourceRuntime `json:"runtime,omitempty"`
	// Flag indicating whether the source is a linked source in the Delphix
	// system.
	Linked *bool `json:"linked,omitempty"`
	// Flag indicating whether the source is a virtual source in the Delphix
	// system.
	Virtual *bool `json:"virtual,omitempty"`
	// Status of this source.
	// enum = [DEFAULT PENDING_UPGRADE]
	Status string `json:"status,omitempty"`
}

// JSDataContainerModifyOwnerParameters - Input parameters for addOwner or removeOwner for a data container.
// extends TypedObject
type JSDataContainerModifyOwnerParameters struct {
	// Object type.
	// format = type
	// required = true
	Type string `json:"type,omitempty"`
	// A reference to the user object for whom to add or remove
	// authorizations.
	// referenceTo = /delphix-user.json
	// required = true
	// format = objectReference
	Owner string `json:"owner,omitempty"`
}

// OracleSyncParameters - The parameters to use as input to sync Oracle databases.
// extends SyncParameters
type OracleSyncParameters struct {
	// Indicates whether a fresh SnapSync must be started regardless if it
	// was possible to resume the current SnapSync. If true, we will not
	// resume but instead ignore previous progress and backup all datafiles
	// even if already completed from previous failed SnapSync. This does not
	// force a full backup, if an incremental was in progress this will start
	// a new incremental snapshot.
	// default = false
	// create = optional
	DoNotResume *bool `json:"doNotResume,omitempty"`
	// Object type.
	// required = true
	// format = type
	Type string `json:"type,omitempty"`
	// Whether or not to take another full backup of the source database.
	// default = false
	ForceFullBackup *bool `json:"forceFullBackup,omitempty"`
	// True if two SnapSyncs should be performed in immediate succession to
	// reduce the number of logs required to provision the snapshot. This may
	// significantly reduce the time necessary to provision from a snapshot.
	// default = false
	// create = optional
	DoubleSync *bool `json:"doubleSync,omitempty"`
	// Skip check that tests if there is enough space available to store the
	// database in the Delphix Engine. The Delphix Engine estimates how much
	// space a database will occupy after compression and prevents SnapSync
	// if insufficient space is available. This safeguard can be overridden
	// using this option. This may be useful when linking highly compressible
	// databases.
	// default = false
	// create = optional
	SkipSpaceCheck *bool `json:"skipSpaceCheck,omitempty"`
}