package db

import (
	//"LianFaPhone/lfp-backend-api/config"
	"errors"
	"github.com/bluele/gcache"
	"time"
)

var GCache Cache

type Cache struct {
	RoleAccessCache gcache.Cache
	IgnoreAccessCache gcache.Cache
}

func (this *Cache) Init() {
}

//**************************************************/
func (this *Cache) GetRoleAccess(rid int) (interface{}, error) {
	if this.RoleAccessCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.RoleAccessCache.Get(rid)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetRoleAccess(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.RoleAccessCache = gcache.New(100).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveRoleAccess(tp int) {
	if this.RoleAccessCache == nil {
		return
	}
	this.RoleAccessCache.Remove(tp)
}

//**************************************************/
func (this *Cache) GetIgnoreAccess(code string) (interface{}, error) {
	if this.IgnoreAccessCache == nil {
		return nil, errors.New("not init")
	}
	value, err := this.IgnoreAccessCache.Get(0)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (this *Cache) SetIgnoreAccess(f func(key interface{}) (interface{}, *time.Duration, error)) {
	this.IgnoreAccessCache = gcache.New(2).LRU().LoaderExpireFunc(f).Build()
}

func (this *Cache) RemoveIgnoreAccess(tp string) {
	if this.IgnoreAccessCache == nil {
		return
	}
	this.IgnoreAccessCache.Remove(0)
}
