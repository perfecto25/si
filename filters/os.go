package filters

//type OperatingSystem struct {
//Kernel        string `json:"kernel"`
//KernelRelease string `json:"kernel release"`
//KernelVersion string `json:"kernel version"`
//LSBRelease    string `json:"lsb release,omitempty"`
//	Kernel, KernelRelease string/
//}

// comment
//func ReturnOS() map[string]string {
// comment

//var m map[string]string

//osinfo := make(map[string]string)

//osinfo["Kernel"] = "centos 7"
//osinfo["kernel_release"] = "6.34"

//	return osinfo

//}

func ReturnOS2(data map[string]map[string]string) map[string]map[string]string {

	data["operating system"]["kernel"] = "centos 7"

	return data
}
