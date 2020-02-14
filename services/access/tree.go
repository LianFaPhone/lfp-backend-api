package access

import (
	"LianFaPhone/lfp-backend-api/models"
)

func (this *AccessList) Tree(list []*models.Access) []*models.Access {
	data := this.buildData(list)
	result := this.makeTreeCore(-1, data)//-1难道不是一个bug么

	return result
}

func (this *AccessList) Tree2(list []*models.Access) []*models.Access {
	data,min := this.buildData2(list)
	result := this.makeTreeCore(min, data)

	return result
}

func (this *AccessList) buildData2(list []*models.Access) (map[int64]map[int64]*models.Access, int64) {
	var data = make(map[int64]map[int64]*models.Access) //parent_id,id,*Access
	min := int64(-2)
	for i, v := range list {
		id := v.Id
		parentId := v.ParentId
		if _, ok := data[parentId]; !ok {
			data[parentId] = make(map[int64]*models.Access)
		}
		if i == 0{
			min = parentId
		}else if parentId <min {
			min = parentId
		}
		data[parentId][id] = v
	}

	return data,min
}

func (this *AccessList) buildData(list []*models.Access) map[int64]map[int64]*models.Access {
	var data = make(map[int64]map[int64]*models.Access) //parent_id,id,*Access
	for _, v := range list {
		id := v.Id
		parentId := v.ParentId
		if _, ok := data[parentId]; !ok {
			data[parentId] = make(map[int64]*models.Access)
		}

		data[parentId][id] = v
	}

	return data
}

func (this *AccessList) makeTreeCore(index int64, data map[int64]map[int64]*models.Access) []*models.Access {
	tmp := make([]*models.Access, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = this.makeTreeCore(id, data)
		}

		tmp = append(tmp, item)
	}
	//sort.Sort(AccessListSort(tmp))
	return tmp
}
