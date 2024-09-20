// Code generated by goctl. DO NOT EDIT.

package socialmodels

import (
	"context"
	"database/sql"
	"fmt"

	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	friendsFieldNames          = builder.RawFieldNames(&Friends{})
	friendsRows                = strings.Join(friendsFieldNames, ",")
	//friendsRowsExpectAutoSet   = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	//friendsRowsWithPlaceHolder = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
	friendsRowsExpectAutoSet   = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	friendsRowsWithPlaceHolder = strings.Join(stringx.Remove(friendsFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheFriendsIdPrefix = "cache:friends:id:"
)

type (
	friendsModel interface {
		Insert(ctx context.Context, data *Friends) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Friends, error)
		FindByUidAndFid(ctx context.Context, uid, fid string)(*Friends, error)
		Update(ctx context.Context, data *Friends) error
		Delete(ctx context.Context, id uint64) error
		Inserts(ctx context.Context, session sqlx.Session, data ...*Friends) (sql.Result, error)
		ListByUserid(ctx context.Context, userId string) ([]*Friends, error)
	}

	defaultFriendsModel struct {
		sqlc.CachedConn
		table string
	}

	Friends struct {
		Id        uint64         `db:"id"`
		UserId    string         `db:"user_id"`
		FriendUid string         `db:"friend_uid"`
		Remark    sql.NullString `db:"remark"`
		AddSource sql.NullInt64  `db:"add_source"`
		CreatedAt sql.NullTime   `db:"created_at"`
	}
)

func newFriendsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFriendsModel {
	return &defaultFriendsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`friends`",
	}
}

func (m *defaultFriendsModel) Delete(ctx context.Context, id uint64) error {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, friendsIdKey)
	return err
}

func (m *defaultFriendsModel) ListByUserid(ctx context.Context, userId string) ([]*Friends, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", friendsRows, m.table)

	var resp []*Friends
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}


func (m *defaultFriendsModel) FindByUidAndFid(ctx context.Context, uid, fid string)(*Friends, error){
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `friend_uid` = ?", friendsRows, m.table)
	// 业务中不常使用，所以不使用缓存的方式查询

	var resp Friends
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, uid, fid)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendsModel) FindOne(ctx context.Context, id uint64) (*Friends, error) {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, id)
	var resp Friends
	err := m.QueryRowCtx(ctx, &resp, friendsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendsRows, m.table)
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



//func (m *defaultFriendsModel) Insert(ctx context.Context, data *Friends) (sql.Result, error) {
//	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, data.Id)
//	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
//		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, friendsRowsExpectAutoSet)
//		return conn.ExecCtx(ctx, query, data.UserId, data.FriendUid, data.Remark, data.AddSource)
//	}, friendsIdKey)
//	return ret, err
//}
// 增加字段CreatedAt
func (m *defaultFriendsModel) Insert(ctx context.Context, data *Friends) (sql.Result, error) {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, friendsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FriendUid, data.Remark, data.AddSource, data.CreatedAt)
	}, friendsIdKey)
	return ret, err
}

func (m *defaultFriendsModel) Inserts(ctx context.Context, session sqlx.Session, data ...*Friends) (sql.Result, error) {
	var (
		sql  strings.Builder
		args []any
	)

	if len(data) == 0 {
		return nil, nil
	}

	// insert into tablename values(数据), (数据)
	sql.WriteString(fmt.Sprintf("insert into %s (%s) values ", m.table, friendsRowsExpectAutoSet))
	// 最后的语句是INSERT INTO table_name (column1, column2, column3) VALUES (?, ?, ?, ?, ?), (?, ?, ?, ?, ?), ...;
	for i, v := range data {
		sql.WriteString("(?, ?, ?, ?, ?)")
		//sql.WriteString("(?, ?, ?, ?)")
		args = append(args, v.UserId, v.FriendUid, v.Remark, v.AddSource, v.CreatedAt)
		//args = append(args, v.UserId, v.FriendUid, v.Remark, v.AddSource)

		if i == len(data)-1 {
			break
		}

		sql.WriteString(",")
	}

	fmt.Println(sql.String())
	return session.ExecCtx(ctx, sql.String(), args...)

}


func (m *defaultFriendsModel) Update(ctx context.Context, data *Friends) error {
	friendsIdKey := fmt.Sprintf("%s%v", cacheFriendsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FriendUid, data.Remark, data.AddSource, data.Id)
	}, friendsIdKey)
	return err
}

func (m *defaultFriendsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFriendsIdPrefix, primary)
}

func (m *defaultFriendsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFriendsModel) tableName() string {
	return m.table
}
