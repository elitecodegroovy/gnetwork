package model

type ReleaseMessage struct {
	Model
	AppName       string `gorm:"column:app_name;        not null; size:50" json:"appName"`
	ClusterName   string `gorm:"column:cluster_name;    not null; size:10" json:"clusterName"`
	NamespaceName string `gorm:"column:namespace_name;  not null; size:10" json:"namespaceName"`
	Format        string `gorm:"column:format;          not null;"          json:"format"`
	Value         string `gorm:"column:value;           type:longtext"      json:"value"`
}
