package auth

const (
	// TableAppSpec - apps_spec table
	TableAppSpec = "apps_spec"
	// TablePodData - pod_data table
	TablePodData = "pod_data"
	// TableServiceData - service_data table
	TableServiceData = "service_data"
	// TableNodeData - node_data table
	TableNodeData = "node_data"
	// TableSessionData - session_data table
	TableSessionData = "session_data"
	// TableDataNodeMapData - data_node_map table
	TableDataNodeMapData = "data_node_map"
)

// NodeInfo struct maps to the node_data table
type NodeInfo struct {
	NodeIP   string `json:"nodeIP"`
	NodeName string `json:"nodeName"`
}

// NodeQueryRequest is the parameter for GET request
type NodeQueryRequest struct {
	NodeName string `json:"nodeName"`
}

// AppSpec struct maps to the app_spec table
type AppSpec struct {
	AppName           string `json:"appName"`
	IsStateless       bool   `json:"isStateless"`
	MaxSessionPerNode int    `json:"maxSessionPerNode"`
	Label             string `json:"label"`
}

//PodInfo data struct
type PodInfo struct {
	PodName       string `json:"podname,omitempty"`
	PodIP         string `json:"podip,omitempty"`
	AppName       string `json:"appname,omitempty"`
	IsFree        bool   `json:"isfree,omitempty"`
	ServiceID     string `json:"serviceid,omitempty"`
	NodeName      string `json:"nodename,omitempty"`
	NodeIP        string `json:"nodeip,omitempty"`
	ContainerID   string `json:"containerid,omitempty"`
	ContainerName string `json:"containername,omitempty"`
	Priority      string `json:"priority,omitempty"`
}

//PodQueryRequest data struct
type PodQueryRequest struct {
	PodName string `json:"podName"`
}

// PodData -
type PodData struct {
	Pods map[string]PodInfo `json:"podinfo"` // Format is of <PodID:PodInfo>
	// Contains keys as Appname, Grouplabel, Namespace
	AppName string `json:"appname"` // Contains Appname
}

//ServiceInfo -data struct
type ServiceInfo struct {
	ServiceID string `json:"serviceid,omitempty"`
	AppName   string `json:"appname,omitempty"`
	Selector  string `json:"selector,omitempty"`
	PodID     string `json:"podid,omitempty"`
	LinkFlag  bool   `json:"linkflag,omitempty"`
}

// ServiceQueryRequest data struct
type ServiceQueryRequest struct {
	ServiceID string `json:"serviceId"`
}

// ServiceData -
type ServiceData struct {
	Services map[string]ServiceInfo `json:"services"` // Format is of <ServiceID:ServiceInfo>
	// appData  map[string]string       // Contains keys as Appname, Grouplabel, Namespace
	AppName string `json:"appname"`
}

// QueryAppSpec is the parameter for GET request
type QueryAppSpec struct {
	AppName string `json:"appName"`
}

// NodeData -
type NodeData struct {
	Node         string `json:"node"`
	AccessedTime int    `json:"accessedtime"` // not used
	Timestamp    int64  `json:"timestamp"`
}

// DataNodeItem holds the node values against dataiod
type DataNodeItem struct {
	DataUID      string     `json:"dataUID"`
	NodeDataList []NodeData `json:"nodelist"`
}

// DataNodeCollection holds DataNodeItem list
type DataNodeCollection struct {
	DataNodeList []DataNodeItem `json:"dataNodeList"`
}

// SessionData - session details
type SessionData struct {
	SessionID string `json:"sessionId"`
	AppName   string `json:"appName"`
	Pod       string `json:"pod"`
	Service   string `json:"service"`
	Node      string `json:"node"`
}

// SessionCollection  -
type SessionCollection struct {
	SessionDataList []SessionData `json:"sessionDataList"`
}

// DeleteSessionReq - Delete session request
type DeleteSessionReq struct {
	SessionID string `json:"sessionid"`
}

// AppSpecCollection holds delatils of all apps
type AppSpecCollection struct {
	AppSpecList []AppSpec `json:"appSpecList"`
}

// AppSpecDeleteReq struct app spec delete request
type AppSpecDeleteReq struct {
	AppName string `json:"appName"`
}

// AppServiceInfo holds Service-Pod detail aganist the app
type AppServiceInfo struct {
	ServiceInfoList []ServiceInfo `json:"serviceInfoList"`
}

// AppPodInfo holds Service-Pod detail aganist the app
type AppPodInfo struct {
	PodInfoList []PodInfo `json:"podInfoList"`
}
