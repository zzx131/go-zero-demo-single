package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysUserFieldNames          = builder.RawFieldNames(&SysUser{})
	sysUserRows                = strings.Join(sysUserFieldNames, ",")
	sysUserRowsExpectAutoSet   = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysUserRowsWithPlaceHolder = strings.Join(stringx.Remove(sysUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	SysUserModel interface {
		Insert(ctx context.Context, data *SysUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysUser, error)
		Update(ctx context.Context, data *SysUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysUser struct {
		Id        int64          `db:"id"`
		Username  string         `db:"username"` // 用户名
		RealName  string         `db:"real_name"`
		Password  string         `db:"password"`
		OrgName   sql.NullString `db:"org_name"` // 组织名称
		OrgId     sql.NullInt64  `db:"org_id"`
		LockFlag  sql.NullInt64  `db:"lock_flag"`
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
	}
)

func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &defaultSysUserModel{
		conn:  conn,
		table: "`sys_user`",
	}
}

func (m *defaultSysUserModel) Insert(ctx context.Context, data *SysUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.RealName, data.Password, data.OrgName, data.OrgId, data.LockFlag, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultSysUserModel) FindOne(ctx context.Context, id int64) (*SysUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysUserRows, m.table)
	var resp SysUser
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

func (m *defaultSysUserModel) Update(ctx context.Context, data *SysUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Username, data.RealName, data.Password, data.OrgName, data.OrgId, data.LockFlag, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultSysUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}
