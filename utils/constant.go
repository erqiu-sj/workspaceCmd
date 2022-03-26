package utils

const (
	CmdName                = "work"
	Version                = "0.0.1"
	IniConfigurationFolder = "./ini"
)

// workspace group 工作组配置
const (
	WorkgroupConfigurationNameInIni = "workspaceGroup" // 在 ini 文件中的 工作组 section name
	WorkGroupName                   = "groupName"      // 工作组key groupName
	WorkGroupWithRemarks            = "groupRemarks"   // 工作组key group remarks
	WorkGroupPath                   = "groupPath"      // 工作组路径
)

// workspace 工作组配置

const (
	WorkSpacePath    = "path"    // 工作区路径
	WorkSpaceRemarks = "remarks" // 工作区备注
)

// Error line
const (
	NotAFile                  = `the parameter address is not a file or folder address`              // 不是文件夹或者文件地址
	RemarksNotEmpty           = "remarks cannot be empty"                                            // 备注不能为空
	CreateErrorFile           = "error creating file"                                                // 创建文件出错
	CheckWorkSpaceOrWorkGroup = "check whether the corresponding workgroup or workspace is created?" // 检查是否创建对应的工作区或工作组
	NotExitsWorkGroup         = "workgroup does not exist"                                           // 工作组不存在
	CannotBeEmpty             = "cannot be empty"                                                    // 不能为空
)
