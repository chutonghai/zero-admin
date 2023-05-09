// Code generated by goctl. DO NOT EDIT.

package sysmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysConfigFieldNames          = builder.RawFieldNames(&SysConfig{})
	sysConfigRows                = strings.Join(sysConfigFieldNames, ",")
	sysConfigRowsExpectAutoSet   = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysConfigRowsWithPlaceHolder = strings.Join(stringx.Remove(sysConfigFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	sysConfigModel interface {
		Insert(ctx context.Context, data *SysConfig) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysConfig, error)
		Update(ctx context.Context, data *SysConfig) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysConfigModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysConfig struct {
		Id          int64          `db:"id"`          // 编号
		Value       string         `db:"value"`       // 数据值
		Label       string         `db:"label"`       // 标签名
		Type        string         `db:"type"`        // 类型
		Description string         `db:"description"` // 描述
		Sort        float64        `db:"sort"`        // 排序（升序）
		CreateBy    sql.NullString `db:"create_by"`   // 创建人
		CreateTime  time.Time      `db:"create_time"` // 创建时间
		UpdateBy    sql.NullString `db:"update_by"`   // 更新人
		UpdateTime  sql.NullTime   `db:"update_time"` // 更新时间
		Remarks     sql.NullString `db:"remarks"`     // 备注信息
		DelFlag     int64          `db:"del_flag"`    // 是否删除  -1：已删除  0：正常
	}
)

func newSysConfigModel(conn sqlx.SqlConn) *defaultSysConfigModel {
	return &defaultSysConfigModel{
		conn:  conn,
		table: "`sys_config`",
	}
}

func (m *defaultSysConfigModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysConfigModel) FindOne(ctx context.Context, id int64) (*SysConfig, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysConfigRows, m.table)
	var resp SysConfig
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysConfigModel) Insert(ctx context.Context, data *SysConfig) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysConfigRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.UpdateBy, data.Remarks, data.DelFlag)
	return ret, err
}

func (m *defaultSysConfigModel) Update(ctx context.Context, data *SysConfig) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysConfigRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Value, data.Label, data.Type, data.Description, data.Sort, data.CreateBy, data.UpdateBy, data.Remarks, data.DelFlag, data.Id)
	return err
}

func (m *defaultSysConfigModel) tableName() string {
	return m.table
}