// Copyright 2018 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package project

import (
    "time"

    "github.com/dreamans/syncd/model"
)

type Project struct {
    ID              int     `gorm:"primary_key"`
    Name            string  `gorm:"type:varchar(100);not null;default:''"`
    Description     string  `gorm:"type:varchar(100);not null;default:''"`
    SpaceId         int     `gorm:"type:int(11);not null;default:0"`
    RepoUrl         string  `gorm:"type:varchar(200);not null;default:''"`
    RepoMode        int     `gorm:"type:int(11);not null;default:0"`
    RepoBranch      string  `gorm:"type:varchar(20);not null;default:''"`
    ExcludeFiles    string  `gorm:"type:varchar(1000);not null;default:''"`
    DeployServer    string  `gorm:"type:varchar(2000);not null;default:''"`
    DeployUser      string  `gorm:"type:varchar(20);not null;default:''"`
    DeployPath      string  `gorm:"type:varchar(100);not null;default:''"`
    DeployTimeout   int     `gorm:"type:int(11);not null;default:0"`
    AuditNoticeEmail    string  `gorm:"type:varchar(1000);not null;default:''"`
    DeployNoticeEmail   string  `gorm:"type:varchar(1000);not null;default:''"`
    PreDeployCmd    string  `gorm:"type:varchar(2000);not null;default:''"`
    PostDeployCmd   string  `gorm:"type:varchar(2000);not null;default:''"`
    NeedAudit       int     `gorm:"type:int(11);not null;default:0"`
    Status          int     `gorm:"type:int(11);not null;default:0"`
    Ctime           int     `gorm:"type:int(11);not null;default:0"`
}

const (
    TableName = "project"
)

func Create(data *Project) bool {
    data.Ctime = int(time.Now().Unix())
    return model.Create(TableName, data)
}

func List(query model.QueryParam) ([]Project, bool) {
    var p []Project
    ok := model.GetMulti(TableName, &p, query)
    return p, ok
}

func Total(query model.QueryParam) (int, bool) {
    var count int
    ok := model.Count(TableName, &count, query)
    return count, ok
}

func Get(id int) (Project, bool){
    var data Project
    ok := model.GetByPk(TableName, &data, id)
    return data, ok
}

func GetOne(query model.QueryParam) (Project, bool) {
    var data Project
    ok := model.GetOne(TableName, &data, query)
    return data, ok
}

func Update(id int, data Project) bool {
    updateFields := map[string]interface{}{
        "name": data.Name,
        "description": data.Description,
        "repo_url": data.RepoUrl,
        "deploy_server": data.DeployServer,
        "deploy_user": data.DeployUser,
        "deploy_path": data.DeployPath,
        "deploy_timeout": data.DeployTimeout,
        "audit_notice_email": data.AuditNoticeEmail,
        "deploy_notice_email": data.DeployNoticeEmail,
        "pre_deploy_cmd": data.PreDeployCmd,
        "post_deploy_cmd": data.PostDeployCmd,
        "need_audit": data.NeedAudit,
        "repo_mode": data.RepoMode,
        "repo_branch": data.RepoBranch,
        "exclude_files": data.ExcludeFiles,
        "ctime": int(time.Now().Unix()),
        "status": data.Status,
    }
    ok := model.Update(TableName, updateFields, model.QueryParam{
        Where: []model.WhereParam{
            model.WhereParam{
                Field: "id",
                Prepare: id,
            },
        },
    })
    return ok
}

func Delete(id int) bool {
    ok := model.DeleteByPk(TableName, Project{ID: id})
    return ok
}

func UpdateFields(id int, data map[string]interface{}) bool {
    ok := model.Update(TableName, data, model.QueryParam{
        Where: []model.WhereParam{
            model.WhereParam{
                Field: "id",
                Prepare: id,
            },
        },
    })
    return ok
}
