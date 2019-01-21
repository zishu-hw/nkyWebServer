package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Node struct {
	ID   int    `orm:"column(f_nodeID)" json:"NodeID"`
	Name string `orm:"column(f_nodeName)" json:"NodeName"`
}

type Data struct {
	EnvValue float64 `orm:"column(f_value)" json:"EnvValue"`
	EnvName  string  `orm:"column(f_envName)" json:"EnvName"`
	EnvUnit  string  `orm:"column(f_envUnit)" json:"EnvUnit"`
}

type NodeData struct {
	Node    `json:"Node"`
	EnvData []Data `json:"EnvData"`
}

func RegisterDB() {
	orm.RegisterModel(new(Node))
}

func CheckAccount(uname, pwd string) bool {
	var dbpwd string
	o := orm.NewOrm()
	err := o.Raw("select f_userPwd from tb_manager where f_userName=?", uname).QueryRow(&dbpwd)
	if err != nil {
		beego.Info(err.Error())
		return false
	}
	if pwd == dbpwd {
		return true
	}
	return false
}

func GetAllNodes() (nodes *[]Node, err error) {
	nodes = new([]Node)
	o := orm.NewOrm()

	_, err = o.Raw("select f_nodeID,f_nodeName from tb_node").QueryRows(nodes)
	return
}

func GetAllNodeData() (nodeDatas []*NodeData, err error) {
	nodes, err := GetAllNodes()
	nodeDatas = make([]*NodeData, len(*nodes))
	o := orm.NewOrm()

	for index, _ := range nodeDatas {
		nodeDatas[index] = &NodeData{Node: (*nodes)[index]}
		_, err = o.Raw(`select tb_vr.f_value,tb_ep.f_envName,tb_ep.f_envUnit 
		from tb_valueRealTime tb_vr, tb_envParameter tb_ep 
		WHERE tb_vr.f_envID=tb_ep.f_envID and tb_vr.f_nodeID=?;`, nodeDatas[index].ID).QueryRows(&(nodeDatas[index].EnvData))
	}
	return
}
