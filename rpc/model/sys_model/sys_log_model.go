package sys_model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SysLogModel = (*customSysLogModel)(nil)

type (
	// SysLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLogModel.
	SysLogModel interface {
		sysLogModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customSysLogModel struct {
		*defaultSysLogModel
	}
)

// NewSysLogModel returns a model for the database table.
func NewSysLogModel(conn sqlx.SqlConn) SysLogModel {
	return &customSysLogModel{
		defaultSysLogModel: newSysLogModel(conn),
	}
}

func (m *customSysLogModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]SysLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", sysLogRows, m.table)
	var resp []SysLog
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysLogModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customSysLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
