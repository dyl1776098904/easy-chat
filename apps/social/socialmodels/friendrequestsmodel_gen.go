// Code generated by goctl. DO NOT EDIT.

package socialmodels

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	friendRequestsFieldNames          = builder.RawFieldNames(&FriendRequests{})
	friendRequestsRows                = strings.Join(friendRequestsFieldNames, ",")
	friendRequestsRowsExpectAutoSet   = strings.Join(stringx.Remove(friendRequestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	friendRequestsRowsWithPlaceHolder = strings.Join(stringx.Remove(friendRequestsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFriendRequestsIdPrefix = "cache:friendRequests:id:"
)

type (
	friendRequestsModel interface {
		Insert(ctx context.Context, data *FriendRequests) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*FriendRequests, error)
		FindByReqUidAndUserId(ctx context.Context, rid, uid string) (*FriendRequests, error)
		Update(ctx context.Context, session sqlx.Session, data *FriendRequests) error
		Delete(ctx context.Context, id uint64) error
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	defaultFriendRequestsModel struct {
		sqlc.CachedConn
		table string
	}

	FriendRequests struct {
		Id           uint64         `db:"id"`
		UserId       string         `db:"user_id"`
		ReqUid       string         `db:"req_uid"`
		ReqMsg       sql.NullString `db:"req_msg"`
		ReqTime      time.Time      `db:"req_time"`
		HandleResult sql.NullInt64  `db:"handle_result"`
		HandleMsg    sql.NullString `db:"handle_msg"`
		HandledAt    sql.NullTime   `db:"handled_at"`
	}
)

func newFriendRequestsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFriendRequestsModel {
	return &defaultFriendRequestsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`friend_requests`",
	}
}

func (m *defaultFriendRequestsModel) Delete(ctx context.Context, id uint64) error {
	friendRequestsIdKey := fmt.Sprintf("%s%v", cacheFriendRequestsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, friendRequestsIdKey)
	return err
}

func (m *defaultFriendRequestsModel) FindOne(ctx context.Context, id uint64) (*FriendRequests, error) {
	friendRequestsIdKey := fmt.Sprintf("%s%v", cacheFriendRequestsIdPrefix, id)
	var resp FriendRequests
	err := m.QueryRowCtx(ctx, &resp, friendRequestsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendRequestsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
// 利用请求id和用户id查询申请的状态
func (m *defaultFriendRequestsModel) FindByReqUidAndUserId(ctx context.Context, rid, uid string) (*FriendRequests, error) {
	query := fmt.Sprintf("select %s from %s where `req_uid` = ? and `user_id` = ?", friendRequestsRows, m.table)

	var resp FriendRequests
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, rid, uid)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}


func (m *defaultFriendRequestsModel) Trans(ctx context.Context, fn func(ctx context.Context,
	session sqlx.Session) error) error {
	//开始一个事务，事务具体的细节由自己提供，最后的结果是提交还是RollBack由TransactCtx决定
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}



func (m *defaultFriendRequestsModel) Insert(ctx context.Context, data *FriendRequests) (sql.Result, error) {
	friendRequestsIdKey := fmt.Sprintf("%s%v", cacheFriendRequestsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, friendRequestsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ReqUid, data.ReqMsg, data.ReqTime, data.HandleResult, data.HandleMsg, data.HandledAt)
	}, friendRequestsIdKey)
	return ret, err
}

func (m *defaultFriendRequestsModel) Update(ctx context.Context, session sqlx.Session, data *FriendRequests) error {
	friendRequestsIdKey := fmt.Sprintf("%s%v", cacheFriendRequestsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendRequestsRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.ReqUid, data.ReqMsg, data.ReqTime, data.HandleResult, data.HandleMsg, data.HandledAt, data.Id)
	}, friendRequestsIdKey)
	return err
}

func (m *defaultFriendRequestsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFriendRequestsIdPrefix, primary)
}

func (m *defaultFriendRequestsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendRequestsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFriendRequestsModel) tableName() string {
	return m.table
}
