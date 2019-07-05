/*
   @Time : 2019-07-05 14:55
   @Author : frozenChen
   @File : home
   @Software: blog-api
*/
package dao

import "blog/model"

func (d Dao) GetHomeList(page *model.Page) ([]*model.Article, error) {

	var list = make([]*model.Article, 0)

	err := d.Db.Limit(page.Size, page.GetOffSet()).Find(&list)

	return list, err
}

func (d Dao) GetHomePageIndex(page *model.Page) int {

	count, err := d.Db.Count(&model.Article{})
	if err != nil {
		return 0
	}
	p := int(count) / page.Size
	if int(count)%page.Size != 0 {
		p += 1
	}
	return p

}
